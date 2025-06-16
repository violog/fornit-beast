/* функции пассивного режима мышления

 */

package psychic

/////////////////////////////////////////////////////
/*Найти кадр ветки эп.памяти для заданных условий (conditions[]), включая стимул
с наиболее экстремальной значимостью ИЛИ правил ИЛИ значимостью стимула (PARAMS[2]).

например, conditions := []int{
CurrentCommonBadNormalWell,
CurrentEmotionReception.ID,
detectedActiveLastProblemNodID,
extremImportanceObject.objID }
с данной значимостью effect.
Учительские правила не интересуют, нужна свой кадр, который можно модифицировать.
Но там могут быть самые разные EpisodicTreeNode.Action с данным эффектом (хотя и редко),
так что выдаем первую попавшуюся.

В conditions могут быть все условия обнулены!
*/
var extremMemArr []*EpisodicTreeNode

func findEpisodicBrangeFromObject(conditions []int) *EpisodicTreeNode {
	if conditions[3] == 0 {
		return nil // нужно искать данный стимул при любом состоянии условий!
	}
	extremMemArr = nil
	curLevelSimplify := 0
	// прямо проходим до ObjectID
	for extremMemArr == nil { // с каждым неудачным проходом упрошаем условия
		findEpisodicBrangeCondition(0, conditions, &EpisodicTree, conditions[3])
		if extremMemArr == nil {
			if curLevelSimplify < 3 {
				curLevelSimplify--
				conditions = toSimplifyCondition(conditions, curLevelSimplify)
			} else { // уже итак предельно упрощены условия
				return nil
			}
		} else {
			break // нашли
		}
	}

	extr := selectExtremalFrame(extremMemArr)
	if extr != nil {
		return extr
	}

	return nil
}
func findEpisodicBrangeCondition(level int, cond []int, root *EpisodicTreeNode, trigger int) {

	// Поиск узла с ID из списка cond в дочерних узлах текущего узла
	for _, child := range root.Children { // пока есть дочки
		if isEquivalentCondition(level, &child, cond) {
			if len(child.PARAMS) > 0 { // узел с прописанным PARAMS
				if child.PARAMS[0] < 100 && child.Trigger == trigger { // не учительские правила
					extremMemArr = append(extremMemArr, &child)
					return
				}
			}
			// Если узел найден, продолжим рекурсивно искать в нем далее
			findEpisodicBrangeCondition(level+1, cond, &child, trigger)
			return
		}
		//  пусть перебирает дочки! return 0,level // не найдено совпадение на данном уровне
	}

	// не найден узел на данном уровне
	return
}

// упостить условия поиска
func toSimplifyCondition(cond []int, levelSimplify int) []int {
	if levelSimplify == 1 {
		cond[3] = 0
	}
	if levelSimplify == 2 {
		cond[2] = 0
		cond[3] = 0
	}
	if levelSimplify == 3 {
		cond[1] = 0
		cond[2] = 0
		cond[3] = 0
	}
	//
	newPassiveEnvironment.mood = cond[0]
	newPassiveEnvironment.emotionID = cond[1]
	newPassiveEnvironment.problemNodID = cond[2]
	return cond
}

/////////////////////////////////////////////////////
/* в массиве eArr []*EpisodicTreeNode кадров выбрать один наиболее экстремальный

 */
func selectExtremalFrame(eArr []*EpisodicTreeNode) *EpisodicTreeNode {
	if eArr != nil && len(eArr) > 0 {
		var extrMin *EpisodicTreeNode
		var extrMax *EpisodicTreeNode
		for _, em := range eArr {
			if em.PARAMS[0] < extrMin.PARAMS[0] { // значимость парвила
				extrMin = em
			}
			// пишем все равно в em.PARAMS[0] т.к. это - просто буфер для значимостей
			if em.PARAMS[0] < extrMin.PARAMS[2] { // значимость стимула
				extrMin = em
			}

			if em.PARAMS[0] > extrMax.PARAMS[0] { // значимость парвила
				extrMax = em
			}
			// пишем все равно в em.PARAMS[0] т.к. это - просто буфер для значимостей
			if em.PARAMS[0] > extrMax.PARAMS[2] { // значимость стимула
				extrMax = em
			}
		}
		if extrMax.PARAMS[0] > -extrMin.PARAMS[0] {
			return extrMax
		} else {
			return extrMin
		}
	}
	return nil
}

///////////////////////////////////////////////////////

/*
	выбрать наиболее значимое из еще не угасших циклов кроме главного

Это - начало любого пассивного режима: о чем думать
Заполняется objectsOfPassiveToughtArr
*/

var objectsForPassiveToughtArr []int

func objectsFromPhoneCyclesArr() {
	objectsForPassiveToughtArr = nil
	// перебор фоновых циклов и выявление экстремальных объектов в них
	for _, v := range cyclesArr {
		if v == nil {
			continue
		}
		if v.isMainCycle {
			continue
		}
		if v.impObjID != 0 {
			objectsForPassiveToughtArr = append(objectsForPassiveToughtArr, v.impObjID)
		}
	}
}

//////////////////////////////////////////////////////////////////
/*находим недооцененные по значимости стимула неучительские кадры, т.е. с PARAMS[2]==0
первые 3 объекта с конца, они заполнятся в эпиз.памяти и не будут здесь возникать
*/
var unknownObjectsFrameArr []*EpisodicTreeNode

func unknownObjectsFromMemoryArr() {
	unknownObjectsFrameArr = nil
	fCount := 0 // число найденных
	from := len(EpisodicHistoryArr) - 1
	for i := from; i > 0; i-- {
		node, ok := ReadeEpisodicTreeNodeFromID(EpisodicHistoryArr[i].ID)
		if !ok {
			continue
		}
		if fCount == 3 {
			break
		}
		//неучительские node.PARAMS[0] < 100
		if node.PARAMS != nil && node.PARAMS[0] < 100 && node.PARAMS[2] == 0 {
			unknownObjectsFrameArr = append(unknownObjectsFrameArr, node)
			fCount++
		}
	}
}

////////////////////////////////////////////////////////////////

/*
Один шаг пассивного мышления.

Начиная с опорного кадра эпиз.памяти (*EpisodicTreeNode), а не кадра исторической.

Это - специализированная инфо-функция для мышления в цикле
с выдачей промежуточных результатов в инфо-картину.
Т.к. она не имеет цели и развивается сама по себе, то не вынесена в список инфо-функций, доступных для произвольного выбора.
Но она может быть активирована из какой-то инфо-функции (м.б. для пассивного мышления в фоне).

При проходе не ставится никакой цели (не требуется достичь позитива и т.п.).

Алгоритм действий с memFr *EpisodicTreeNode:
т.к. memFr *EpisodicTreeNode - уже найденный экстремальный кадр (func findEpisodicBrangeFromObject)
то для него просто смотрятся все следующие по истории кадры (в цепочке до прерывания),
из всех находится самый экстремальный,
и его стимул сохраняется в newPassiveEnvironment.extremID,
что и становится исходным для следующей итерации.

Если не находит ничего, то вернуть true
*/
var saveBeginEpisodicTreeNode *EpisodicTreeNode // сохранить кадр начала мышления для модификации после вспомианния
func gotoPassiveMaind(cycle *cycleInfo, memFr *EpisodicTreeNode) bool {
	saveBeginEpisodicTreeNode = memFr

	getNextHistoryEpisodicFrameArr(memFr)
	if fantasmArr != nil && len(fantasmArr) > 0 {
		extr := selectExtremalFrame(extremMemArr)
		if extr == nil {
			return true
		}
		// результаты итерации ЗАПИСАТЬ в newPassiveEnvironment

		/* первые задются в func toSimplifyCondition
		newPassiveEnvironment.mood=
		newPassiveEnvironment.emotionID=
		newPassiveEnvironment.problemNodID=
		*/
		if extr.PARAMS[0] < -5 || extr.PARAMS[2] < -8 {
			newPassiveEnvironment.danger = true
		}
		newPassiveEnvironment.stimulID = extr.Trigger
		newPassiveEnvironment.answerID = extr.Action
		if needOptimisationFrame != nil {
			newPassiveEnvironment.effect = extr.PARAMS[2]
			newPassiveEnvironment.extremID = extr.Trigger
			if extr.PARAMS[0] == 100 { // в учителском стимулом является Action
				newPassiveEnvironment.extremID = extr.Action
			}
		} else {
			newPassiveEnvironment.effect = extr.PARAMS[0]
			newPassiveEnvironment.extremID = extr.Trigger
		}
		newPassiveEnvironment.episodFrameID = extr.ID

		return false
	}

	return true // конец цепочки, исчерпание значимостей, закончить итерации
}

///////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////
/* найти следующие после memFr кадры эпизодов в исторической цепочке до прерывания (пустой кадр)
БЕЗ УСЛОВИЙ кроме настроения(что и дает диапазон фантазирования).
*/
var fantasmArr []*EpisodicTreeNode    // результирующие последующие кадры
var foundFrameArr []*EpisodicTreeNode // промежуточные исходные кадры
func getNextHistoryEpisodicFrameArr(memFr *EpisodicTreeNode) {
	if memFr == nil {
		return
	}
	foundFrameArr = nil
	fantasmArr = nil
	сondintion := []int{
		newPassiveEnvironment.mood,
		0,
		0,
		0}
	findFantasmArr(0, сondintion, &EpisodicTree, memFr.Trigger, memFr.Action)
	if foundFrameArr != nil && len(foundFrameArr) > 1 { //len(foundFrameArr) > 1 нужно чтобы находилось не только для исходного memFr!
		// выбираем последущие кадры в цепочках
		for _, frm := range foundFrameArr {
			nextArr := getNextHistiryEpisodicFrame(frm)
			if nextArr != nil && len(nextArr) > 0 {
				for _, f := range nextArr {
					fantasmArr = append(fantasmArr, f)
				}
			}
		}
	}

	return
}
func findFantasmArr(level int, cond []int, root *EpisodicTreeNode, trigger int, action int) {

	// Поиск узла с ID из списка cond в дочерних узлах текущего узла
	for _, child := range root.Children { // пока есть дочки
		if isEquivalentCondition(level, &child, cond) {
			if len(child.PARAMS) > 0 { // узел с прописанным PARAMS
				if child.PARAMS[0] < 100 && child.Trigger == trigger && child.Action == action { // не учительские правила
					foundFrameArr = append(foundFrameArr, &child)
					return
				}
			}
			// Если узел найден, продолжим рекурсивно искать в нем далее
			findFantasmArr(level+1, cond, &child, trigger, action)
			return
		}
		//  пусть перебирает дочки! return 0,level // не найдено совпадение на данном уровне
	}

	// не найден узел на данном уровне
	return
}

////////////////////////////////////////////////////////////////

/*
	дерево ментальных кадров эпизодов памяти

Носитель ментальных правил:
какая цепочка инфо-функций в условиях данной проблемы и темы привела к данному эффекту.
Выполняет роль Ментальных автоматизмов при поиске мент.правила (последователньости инфо-функций) для уверенного запуска в данных условиях.

В файл пишется:
ID|ParentID|NodePID|ThemeID|PurposeID|info1,info2#Effect|Count

Для каждой новой цели должен начинаться свой кадр ментальной эпиз.памяти: clinerFuncSequence()
это делается в func infoFunc8

После отработки wasRunPurposeActionFunc нужно прекратить набор кадров и ждать ответа: func setCurIfoFuncID

Вроде логично, что если не запускалась функция с wasRunPurposeActionFunc=true (id==14,17,26)
- значит не цикл мышления привел к ответу и он не должен учитываться как источник эффекта.
Но даже при запуске автоматизма продолжается осмысление, в том числе и в фоновых циклах и еще есть доминанта.
Есть CurrentInformationEnvironment.needThinkingAboutAutomatizm - думать о создании автоматизма.
Так что пусть пишутся любые цепочки, даже если не было вызова func infoFunc17.
Вообще TODO нужно продолжить поиск более адаптиной реализации ментальных правил...

EpisodicMentalHistoryArr имеет связь с EpisodicHistoryArr через EpisodicMentalHistoryArr.lastEpisodicMemID
*/
package psychic

import (
	"BOT/lib"
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////

type EpisodicMentalTreeNode struct { // узел дерева Эпизодической памяти
	ID int
	// Узлы дерева:
	/* субъективный узел активной ветки дерева проблем (последний активировавшийся)
	По нему можно найти:
	node,ok:=ReadeProblemTreeNodeFromID()
	if ok{
	NodeAID:=node.autTreeID //конечный узел активной ветки дерева моторных автоматизмов
	NodeSID:=node.situationTreeID //конечный узел активной ветки дерева ситуации
	}*/
	NodePID int

	/* Темы мышления
	 */
	ThemeID int

	/* Текущая цель
	 */
	PurposeID int
	// последовательность ID инфо функций, которая применялось после активации NodePID
	InfoArr []int // разделитель - запятая!
	// Свойства данного узла - массив PARAMS:
	PARAMS []int
	/* [2] иммет два значения:
	Effect int // эффект от действий: (меньше 0) или 0 или (больше 0)
	Count int // всегда>0 - число усреднений Effect - уверенность в Правиле
	*/

	//Связь узлов дерева:
	Children   []EpisodicMentalTreeNode // дочерние узлы (ветвление) НЕ АДРЕСА, А РЕАЛЬНЫЕ ОБЪЕКТЫ
	ParentID   int                      // ID родителя
	ParentNode *EpisodicMentalTreeNode  // адрес родителя
}

// ////////////////////////////////////////////////////
var EpisodicMentalTree EpisodicMentalTreeNode

// var EpisodicMentalTreeNodeFromID=make(map[int]*EpisodicMentalTreeNode)
var EpisodicMentalTreeNodeFromID []*EpisodicMentalTreeNode // сам массив
// запись члена
func WriteEpisodicMentalTreeNodeFromID(index int, value *EpisodicMentalTreeNode) {
	addEpisodicMentalTreeNodeFromID(index)
	EpisodicMentalTreeNodeFromID[index] = value
}
func addEpisodicMentalTreeNodeFromID(index int) {
	if index >= len(EpisodicMentalTreeNodeFromID) {
		newSlice := make([]*EpisodicMentalTreeNode, index+1)
		copy(newSlice, EpisodicMentalTreeNodeFromID)
		EpisodicMentalTreeNodeFromID = newSlice
	}
}

// считывание члена
func ReadeEpisodicMentalTreeNodeFromID(index int) (*EpisodicMentalTreeNode, bool) {
	if index >= len(EpisodicMentalTreeNodeFromID) || EpisodicMentalTreeNodeFromID[index] == nil {
		return nil, false
	}
	return EpisodicMentalTreeNodeFromID[index], true
}

//////////////////////////////////////////////////////////////////////

/*
Связь узлов историческая - ID конечных узлов веток EpisodicMentalTree подряд по порядку возникновения новых.
Записывается в файл "/memory_psy/episodic_history.txt" - сплошная строка EpisodicMentalTreeNode.ID с разделителем "|"
В случе добавления newPARAMS в существующий EpisodicMentalTreeNode пишется его EpisodicMentalTreeNode.ID
так что в EpisodicMentalHistoryArr могут появляться несколько EpisodicMentalTreeNode.ID.

Поиск в EpisodicMentalHistoryArr нужного EpisodicMentalTreeNode.ID
При оптимизации во сне порядок сохраняется, но могут вырезаться отдельные составляющие.
При вырезании из EpisodicMentalHistoryArr пишется в файл episodic_history.txt новая строка
и т.о. EpisodicMentalHistoryArr []int остается с непрерывной цепочкой индексов.

Найти в массиве EpisodicMentalHistoryArr[] индексы всех членов по значению:
iArr:=lib.FindIndexes(EpisodicMentalHistoryArr,ID)

Запись в файле: ID1,LifeTime1|ID2,LifeTime2 ...

EpisodicMentalHistoryArr имеет связь с EpisodicHistoryArr через EpisodicMentalHistoryArr.lastEpisodicMemID
*/
type HistoryMental struct {
	ID       int // ID EpisodicMentalTreeNode
	LifeTime int // время появления узла в пульсах от рождения
	// сопутствующий кадр моторной эпиз.памяти
	lastEpisodicMemID int //Запись saveNewEpisodic всегда предшествует с последним кадром lastEpisodicMemID
}

var EpisodicMentalHistoryArr []HistoryMental

///////////////////////////////////////////////////////

// усреднить эффект в PARAMS узла
func averageMentalEffect(node *EpisodicMentalTreeNode, effect int) {
	if node.PARAMS == nil {
		return
	}
	// вес нового эффекта
	count := node.PARAMS[1]
	if count == 0 {
		lib.TodoPanic("func averageEffect НУЛЕВОЕ ЗНАЧЕНИЕ count")
		return
	}
	w := effect + int(effect/(count+1))
	if w > 10 {
		w = 10
	}
	if w < -10 {
		w = -10
	}
	node.PARAMS[0] = w
	node.PARAMS[1] = count + 1
}

//////////////////////////////////////////

/*
	ЗАПИСАТЬ В ДЕРЕВО НОВЫЙ ЭПИЗОД

По каждому срабатыванию инфо-фукнций заполняется буфер infoFuncSequence для набора инфоID
Буфер очищается clinerFuncSequence() и в конце func saveNewMentalEpisodic
сбрасывает буфер в func clinerAutomatizmRunning()

Если при осмыслении проходились только уровни до 3-го,
то и нет обработки, нет записи в Эпизод.память ментальных кадров (определяет ощущение субъективного времени).

Только есди в цепочке ID есть 14,17,26(запуска автоматизма) можно судить, что именно этот запуск и привел к эффекту,

Запись saveNewEpisodic всегда предшествует с последним кадром lastEpisodicMemID
*/
func saveNewMentalEpisodic(effect int) {
	// последовательность ID инфо-функций, вызываемых после очередной активации
	if infoFuncSequence == nil || len(infoFuncSequence) == 0 { // так не должно быть при effect!=0
		return
	}
	// не сохранять кадры, в InfoArr которых нет id==17 - запуска автоматизма
	/*НЕТ!  if !lib.ExistsValInArr(infoFuncSequence, 17){
		return
	}*/
	infoFuncSequence = lib.RemoveDuplicates(infoFuncSequence) // удалить повторяющиеся номера

	NodePID := detectedActiveLastProblemNodID // ID проблемы
	ThemeID := problemTreeInfo.themeID
	PurposeID := mentalInfoStruct.mentalPurposeID

	params := []int{effect, 1}
	var condArr = []int{NodePID, ThemeID, PurposeID}

	//проверяются дубликаты
	idOld, nodeOld := checMentalEpisodicBranchFromCondition(NodePID, ThemeID, PurposeID)
	if idOld > 0 {
		EpisodicMentalHistoryArr = append(EpisodicMentalHistoryArr, HistoryMental{idOld, LifeTime, lastEpisodicMemID})
		// усреднить эффект и добавить уверенность
		averageMentalEffect(nodeOld, effect)
		params = nodeOld.PARAMS
	}

	if params != nil {
		if params[0] > 10 {
			params[0] = 10
		}
		if params[0] < -10 {
			params[0] = -10
		}
	}

	//добавляем новый кадр
	lastNodeID := addMentalEpisodicFromNodeIDsToBrange(0, 0, condArr, infoFuncSequence, params)

	if idOld == 0 {
		EpisodicMentalHistoryArr = append(EpisodicMentalHistoryArr, HistoryMental{lastNodeID, LifeTime, lastEpisodicMemID})
	}

	clinerFuncSequence() // сброс infoFuncSequence
}

////////////////////////////////////////////////////////////

/*
	Для создания нового узла ветки дерева

Возвращает ID узла и указатель на него - в любом случае, создан ли новый узел или найден такой же уже имеющийся.
*/
var lastIDNodeMentalTree = 0 // счетчик ID узлов, хранит ID хранит ID последнего созданного узла
func createEpisodicNodeMentalTree(parent *EpisodicMentalTreeNode, id int, NodePID int, ThemeID int, PurposeID int, InfoArr []int, PARAMS []int) (int, *EpisodicMentalTreeNode) {
	if parent == nil {
		return 0, nil
	}

	if id == 0 {
		lastIDNodeMentalTree++
		id = lastIDNodeMentalTree
	} else {
		//		newW.ID=id
		if lastIDNodeMentalTree < id {
			lastIDNodeMentalTree = id
		}
	}

	var node EpisodicMentalTreeNode
	node.ID = id
	node.ParentNode = parent
	node.ParentID = parent.ID
	node.NodePID = NodePID
	node.ThemeID = ThemeID
	node.PurposeID = PurposeID
	node.InfoArr = InfoArr
	node.PARAMS = PARAMS

	parent.Children = append(parent.Children, node)
	// четко находим новый вставленный член (а не &parent.Children[count-1])
	var newN *EpisodicMentalTreeNode
	for i := 0; i < len(parent.Children); i++ {
		if parent.Children[i].ID == node.ID {
			newN = &parent.Children[i]
			//		newN.Children=append(newN.Children, node)
		}
	}

	WriteEpisodicMentalTreeNodeFromID(node.ID, &node)

	// т.к. append меняет длину массива, перетусовывая адреса, то нужно обновить адреса в EpisodicMentalTreeNodeFromID:
	// ДЕЛАТЬ ТОЛЬКО ДЛЯ ДЕРЕВЬЕВ!
	updatingEpisodicMentalTreeNodeFromID(parent) // здесь потому, что при загрузке из файла нужно на лету получать адреса
	return id, newN
}

//////////////////////////////////////////////////////////

// ////////////////////////////////
func checMentalEpisodicBranchFromCondition(NodePID int, ThemeID int, PurposeID int) (int, *EpisodicMentalTreeNode) {
	cond := []int{NodePID, ThemeID, PurposeID}
	//maxLev:=5
	id, level := findEpisodicMentalBrange(0, cond, &EpisodicMentalTree)
	if id > 0 {
		// найти уровень lev0 ненулевых параметров
		lev0 := getMentalTrueLevel(cond)
		if lev0 > level {
			return 0, nil
		}
		//if maxLevel==maxLev{// есть такая ветка, просто возвращаем ее значения

		//node,ok:=EpisodicMentalTreeNodeFromID[id]
		node, ok := ReadeEpisodicMentalTreeNodeFromID(id)
		if ok {
			return id, node
		} else { // такого не должно быть
			lib.TodoPanic("В func FindEpisodicMentalTreeNodeFromCondition должно быть значение карты EpisodicMentalTreeNodeFromID.") //вызвать панику
		}

		//}else{// нужно дорастить ветку
		//	return 0,nil
		//}
	} else { // ID==0 вообще нет совпадений, нужно наращивать с основы
		return 0, nil
	}

	return 0, nil
}

// ////////
// найти уровень lev0 ненулевых параметров из массива условий
func getMentalTrueLevel(cond []int) int {
	lev0 := 0
	for i := 0; i < len(cond); i++ {
		if cond[i] == 0 {
			break
		}
		lev0++
	}
	return lev0
}

// //////////
// рекурсивно корректируем адреса всех узлов
// ДЕЛАТЬ ТОЛЬКО ДЛЯ ДЕРЕВЬЕВ!
func updatingEpisodicMentalTreeNodeFromID(rt *EpisodicMentalTreeNode) {
	if rt.ID > 0 {
		rt.ParentNode = EpisodicMentalTreeNodeFromID[rt.ParentID] // wr.ParentNode адрес меняется из=за corretsParent(,
		_, ok := ReadeEpisodicMentalTreeNodeFromID(rt.ParentID)
		if ok {
			WriteEpisodicMentalTreeNodeFromID(rt.ID, rt)
		}
	}
	if rt.Children == nil { // конец ветки
		return
	}
	for i := 0; i < len(rt.Children); i++ {
		updatingEpisodicMentalTreeNodeFromID(&rt.Children[i])
	}
}

///////////////////////////////////////////////////////////

// //////////////////////////////////////////////
// получить массив PARAMS[2] последнего узла дерева
// Только узел c ActionID имеет PARAMS[][]
func getMentalEpisodicPARAMS(nodeID int) []int {

	//var actionNode,ok=EpisodicMentalTreeNodeFromID[nodeID]
	actionNode, ok := ReadeEpisodicMentalTreeNodeFromID(nodeID)
	if !ok {
		return nil
	}

	if len(actionNode.InfoArr) == 0 {
		return nil
	}

	return actionNode.PARAMS //[]int{}
}

////

/////////////////////////////////////////////////////////////////////////

/*
	Доращивание ветки, начиная с заданного узла fromID по массиву всех значений для ветки

newPARAMS - добавить к ветке, если она существует новый массив Effect|LifeTime
Возвращает ID конечного узла ветки.
*/
func addMentalEpisodicFromNodeIDsToBrange(fromID int, lastLevel int, condArr []int, InfoArr []int, newPARAMS []int) int {

	var bNode *EpisodicMentalTreeNode
	if fromID > 0 {
		//node, ok := EpisodicMentalTreeNodeFromID[fromID]
		node, ok := ReadeEpisodicMentalTreeNodeFromID(fromID)
		if !ok {
			return 0
		}
		bNode = node
	} else { // не определен начальный узел, хотя он может и быть
		bNode = &EpisodicMentalTree
	}

	lastNodeID, _ := addMentalEpisodicNodesToBrange(bNode, lastLevel, condArr, InfoArr, newPARAMS)

	return lastNodeID
}

/*
рекурсивно создать все недостающие уровни ветки по значениям condArr []int
*/
func addMentalEpisodicNodesToBrange(fromNode *EpisodicMentalTreeNode, level int, condArr []int, InfoArr []int, newPARAMS []int) (int, *EpisodicMentalTreeNode) {
	if fromNode == nil {
		return 0, nil
	}
	if level >= len(condArr) {
		return fromNode.ID, fromNode
	}
	//vArr := make([]int, len(condArr))// все vArr[n] имеют нулевые значения
	vArr := []int{0, 0, 0}
	for i := 0; i <= level; i++ { // заполнить vArr до текущего уровня включительно, остальные - оставить нулями
		vArr[i] = condArr[i]
	}
	// с каждым уровнем vArr[n] добавляется новое значение для создания узла ветки следующего уровня.
	// если такой узел есть (checkUnicum==true), то просто вернет его node
	var node *EpisodicMentalTreeNode
	var pars []int
	var info []int

	//проверяются дубликаты потому, что новый узел нужно делать только если такого еще нет!!!
	idOld, nodeOld := checMentalEpisodicBranchFromCondition(vArr[0], vArr[1], vArr[2])

	//node.PARAMS=append(node.PARAMS,pars)
	if idOld > 0 {
		node = nodeOld
		if level == 2 {
			if newPARAMS != nil {
				// усреднить эффект
				averageMentalEffect(node, newPARAMS[0])
			}
		}
	} else { // НОВЫЙ
		if level == 2 {
			pars = newPARAMS
			info = InfoArr
		}

		if pars != nil && pars[0] != 100 {
			if pars[0] > 10 {
				pars[0] = 10
			}
			if pars[0] < -10 {
				pars[0] = -10
			}
		}

		_, node = createEpisodicNodeMentalTree(fromNode, 0, vArr[0], vArr[1], vArr[2], info, pars)
	}

	level++
	if level >= len(condArr) {
		return node.ID, node
	}
	id, node := addMentalEpisodicNodesToBrange(node, level, condArr, InfoArr, newPARAMS)
	return id, node
}

///////////////////////////////////////

/*
	поиск конечного узла ветки (lastBrangeID) дерева (root - ID начального узла == 0) по массиву ID узлов ветки (cond []int)

Поиск начинается с первого узла (по значению cond[0]).
Если первый узел найден, то findSampleBrange вызывается рекурсивно для поиска следующего узла
и так далее, пока не будет найден удел последнего члена cond.
Если для текущего cond[level] узел в редеве не существует, то в дереве доращивается ветка по значениям cond []int.

Возвращает ID последней найденной ветки и ее уровень, если конечный узел не найден (если найден, то len(cond)-1).
Если вернула уровень меньше, чем максимальное число уровней ветки,
то ветка доращивается func addNodesToBrange

Ищет до узла Action включительно, а к этому узлу могут быть прикреплены несколько PARAMS
*/
func findEpisodicMentalBrange(level int, cond []int, root *EpisodicMentalTreeNode) (int, int) {

	// Обработка случая когда мы достигли конца дерева
	if cond == nil || len(cond) <= 0 || level >= len(cond) {
		return root.ID, len(cond)
	}

	// Поиск узла с ID из списка cond в дочерних узлах текущего узла
	for _, child := range root.Children {
		if isEquivalentMentalCondition(level, &child, cond) {
			// Если узел найден, продолжим рекурсивно искать в нем далее
			id, lev := findEpisodicMentalBrange(level+1, cond, &child)
			if id == 0 {
				/*Eсли условие cond []int содержит значение 0 для какого-то уровня, то на этом уровне поиск заканчивается.
				if cond[lev]==0{// считается нормальный поиск по неполному условию

				}*/
				return child.ID, lev // поиск закончен на child.ID
				//return 0,lev // не найдено совпадение на данном уровне
			}

			return id, lev
		}
		//  пусть перебирает дочки! return 0,level // не найдено совпадение на данном уровне
	}

	// не найден узел на данном уровне
	return 0, level
}

// //////////////////////////////////////////////////
func isEquivalentMentalCondition(level int, node *EpisodicMentalTreeNode, cond []int) bool {
	// массив значений по уровням - в зависимости от числа уровней в ветке, прописывается вручную
	nArr := []int{node.NodePID, node.ThemeID, node.PurposeID}

	for i := 0; i < len(cond) && i <= level; i++ {
		if cond[i] != nArr[i] {
			return false
		}
	}
	return true
}

/////////////////////////////////////////////////////////

/* загрузить записанное дерево
 */
func loadEpisodicMentalTree() {
	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/episodic_mental_tree.txt")
	initEpisodicMentalTree(strArr)

	str := lib.ReadFileContent(lib.GetMainPathExeFile() + "/memory_psy/episodic_mental_history.txt")

	sArr := strings.Split(str, "|")
	for n := 0; n < len(sArr); n++ {
		if len(sArr[n]) == 0 {
			break
		}
		il := strings.Split(sArr[n], ",")
		id, _ := strconv.Atoi(il[0])
		lifeTime, _ := strconv.Atoi(il[1])
		lastEpisodicMemID, _ := strconv.Atoi(il[2])
		EpisodicMentalHistoryArr = append(EpisodicMentalHistoryArr, HistoryMental{id, lifeTime, lastEpisodicMemID})
	}

	// ОТЛАДКА, ТЕСТИРОВАНИЕ
	// saveNewEpisodic(67,38,107,1,3,3)
	//	saveNewEpisodic(200,39,100,1,3,3)

	// getEpisodesArrFromConditions(0,curStimulImageID,0)
	// getEpisodesArrFromConditions(0,1,0)
	// getEpisodesArrFromConditions(1,1,0)

	// тестирование getTargetEpisodicStrIdArr()
	testing := false
	//	testing=true
	if testing { // здесь еще не активированные деревья

	}
	// setMentalInterruptionEpisosde() // обязательно в конце втыкаем пустой кадр!!! Иначе после просыпания начнет дописывать новые кадры к старым цепочкам
	return
}

// В файл пишется: ID|ParentID|NodePID|ThemeID|PurposeID|info1,info2#Effect|Count
func initEpisodicMentalTree(strArr []string) {

	cunt := len(strArr)
	EpisodicMentalTreeNodeFromID = make([]*EpisodicMentalTreeNode, cunt) //задать сразу имеющиеся в файле число при загрузке из файла
	lastIDNodeMentalTree = 1
	// определить нулевой узел
	//EpisodicMentalTreeNodeFromID[0]=&EpisodicMentalTree// все по нулям по умолчанию
	WriteEpisodicMentalTreeNodeFromID(0, &EpisodicMentalTree)

	//просто проход по всем строкам файла подряд так что сначала идут дочки, потом - их родители
	for n := 0; n < cunt; n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		if len(strArr[n]) < 2 {
			panic("Сбой загрузки дерева: [" + strconv.Itoa(n) + "] " + strArr[n])
			return
		}

		par := strings.Split(strArr[n], "#")
		p := strings.Split(par[0], "|")
		var parArr []string
		if len(par) > 1 {
			parArr = strings.Split(par[1], "|")
		}

		id, _ := strconv.Atoi(p[0])
		parentID, _ := strconv.Atoi(p[1])

		NodePID, _ := strconv.Atoi(p[2])
		ThemeID, _ := strconv.Atoi(p[3])
		PurposeID, _ := strconv.Atoi(p[4])

		var InfoArr []int
		iaStr := strings.Split(p[5], ",")
		if len(iaStr) > 0 {
			for i := 0; i < len(iaStr); i++ {
				if len(iaStr[i]) == 0 {
					continue
				}
				iID, _ := strconv.Atoi(iaStr[i])
				InfoArr = append(InfoArr, iID)
			}
		}

		// считывание PARAMS
		var params []int
		if PurposeID > 0 {
			eff, _ := strconv.Atoi(parArr[0])
			count, _ := strconv.Atoi(parArr[1])
			params = []int{eff, count}
		}

		parentNode, ok := ReadeEpisodicMentalTreeNodeFromID(parentID)
		if ok {

			// Создание нового узла дерева
			node := EpisodicMentalTreeNode{
				ID:        id,
				NodePID:   NodePID,
				ThemeID:   ThemeID,
				PurposeID: PurposeID,

				Children:   []EpisodicMentalTreeNode{},
				ParentID:   parentID,
				ParentNode: parentNode,
			}
			node.PARAMS = params
			node.InfoArr = InfoArr

			node.ParentID = parentNode.ID
			node.ParentNode = parentNode
			parentNode.Children = append(parentNode.Children, node)

			// Добавление узла в массив по ID
			WriteEpisodicMentalTreeNodeFromID(id, &node)
			lastIDNodeMentalTree++

			// т.к. append меняет длину массива, перетусовывая адреса, то нужно обновить адреса в EpisodicMentalTreeNodeFromID:
			// ДЕЛАТЬ ТОЛЬКО ДЛЯ ДЕРЕВЬЕВ!
			updatingEpisodicMentalTreeNodeFromID(parentNode) // здесь потому, что при загрузке из файла нужно на лету получать адреса
		}
	}
	return
}

/////////////////////////////////////

// В файл пишется: ID|ParentID|NodePID|ThemeID|PurposeID|info1,info2#Effect|Count
func SaveEpisodicMentalTree() {
	var out = ""
	cnt := len(EpisodicMentalTree.Children)
	for n := 0; n < cnt; n++ { // чтобы записывалось по порядку родителей
		out += getEpisodicMentalTreeNode(&EpisodicMentalTree.Children[n])
	}
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/episodic_mental_tree.txt", out)

	// запись истории эпизодов EpisodicMentalHistoryArr []int
	out = ""
	for n := 0; n < len(EpisodicMentalHistoryArr); n++ {
		out += strconv.Itoa(EpisodicMentalHistoryArr[n].ID) +
			"," + strconv.Itoa(EpisodicMentalHistoryArr[n].LifeTime) +
			"," + strconv.Itoa(EpisodicMentalHistoryArr[n].lastEpisodicMemID) + "|"
	}
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/episodic_mental_history.txt", out)

	return
}
func getEpisodicMentalTreeNode(wt *EpisodicMentalTreeNode) string {
	var out = ""
	//	if wt.ParentID>0 {
	out += strconv.Itoa(wt.ID) + "|"
	out += strconv.Itoa(wt.ParentID) + "|"
	out += strconv.Itoa(wt.NodePID) + "|"
	out += strconv.Itoa(wt.ThemeID) + "|"
	out += strconv.Itoa(wt.PurposeID) + "|"
	for n := 0; n < len(wt.InfoArr); n++ {
		out += strconv.Itoa(wt.InfoArr[n]) + ","
	}

	// записать массивы PARAMS
	parsArr := getMentalEpisodicPARAMS(wt.ID)
	if parsArr != nil {
		out += "#"
		out += strconv.Itoa(parsArr[0]) + "|"
		out += strconv.Itoa(parsArr[1])
	}

	out += "\r\n"
	//	}
	if wt.Children == nil { // конец
		return out
	}
	for n := 0; n < len(wt.Children); n++ {
		out += getEpisodicMentalTreeNode(&wt.Children[n])
	}
	return out
}

/////////////////////////////////////

// Очистить эпизодическую память
func ClianMentalEpisodicMemory() {
	lib.WriteFileContentExactly(lib.GetMainPathExeFile()+"/memory_psy/episodic_mental_tree.txt", "")
	lib.WriteFileContentExactly(lib.GetMainPathExeFile()+"/memory_psy/episodic_mental_history.txt", "")

	EpisodicMentalHistoryArr = nil
	lastIDNodeMentalTree = 1
	WriteEpisodicMentalTreeNodeFromID(0, &EpisodicMentalTree)
}

/////////////////////////////////////

/* вставить кадр прерывания цепочки (темы)   ВСТАВКА ПУСТЫХ КАДРОВ НЕ НУЖНА!!!!просто условия будут разделять.
Означает, что цепочка кадров памяти данной темы закончиена (fornit.ru/67675).
Вставляется по щелчку по плащке ожидания ответа (func StopWaitingWeriodFromOperator())

func setMentalInterruptionEpisosde(){
	// первым не включать
	if len(EpisodicMentalHistoryArr)==0{
		return
	}

	// если только что уже был пустой кадр, то не вставлять лишний
	lastI:=len(EpisodicMentalHistoryArr)-1
	if EpisodicMentalHistoryArr[lastI].ID==-1{
		return
	}
	// пустой кадр с ID=-1
	EpisodicMentalHistoryArr = append(EpisodicMentalHistoryArr, HistoryMental{-1, LifeTime,lastEpisodicMemID})
}*/
///////////////////////////////////////////////////////////

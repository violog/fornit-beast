/* Функции Значимости объектов восприятия

Значимость всегда определяется в контексте всех предшествующих условий,
т.е. специфична для активностей деревьев автоматизмов и понимания.

При каждом вызове consciousnessElementary определяется текущий объект наибольшой значимости в воспринимаемом -
в функции определения текущей Цели getMentalPurpose().
*/

package psychic

import "BOT/lib"

/////////////////////////////////////////////////

///////////////////////////////////////////////////
/*
найти extremImportance по ActionsImage ID в тек.условиях detectedActiveLastProblemNodID
*/
func getExtremObjFromID(actID int) *extremImportance {

	obj, _ := getObjectsImportanceValue(actID, detectedActiveLastProblemNodID)
	if obj != nil {
		return obj
	}
	return nil
}

/////////////////////////////////////////////////

////////////////////////////////////////////////////
/*
	значимость ID объекта внимания для данного ProblemID (detectedActiveLastProblemNodID)

в условиях ProblemID (текущий: detectedActiveLastUnderstandingNodID)
Возвращает:
1 - объект значимости *extremImportance
2 - значимость
*/
var extremImportanceFind *extremImportance // найденная экстремальная значимость
func getObjectsImportanceValue(ObjectID int, ProblemID int) (*extremImportance, int) {
	if ObjectID == 0 {
		return nil, 0
	}
	extremImportanceFind = nil
	cond := []int{CurrentCommonBadNormalWell, CurrentEmotionReception.ID, ProblemID}
	findEpisodicBrangeImportance(0, cond, &EpisodicTree, ObjectID)
	if extremImportanceFind != nil {
		return extremImportanceFind, extremImportanceFind.extremVal
	}
	return nil, 0
}

// поиск конечного узла ветки (lastBrangeID) дерева (root - ID начального узла == 0) по массиву ID узлов ветки (cond []int)
func findEpisodicBrangeImportance(level int, cond []int, root *EpisodicTreeNode, ObjectID int) {

	if cond == nil || len(cond) <= 0 || level >= len(cond) { // Обработка случая когда мы достигли конца дерева
		//	return   !кроме условий тут еще 1 или 2 уровня прохода в поиске ObjectID
	}
	// Поиск узла с ID из списка cond в дочерних узлах текущего узла
	for _, child := range root.Children { // пока есть дочки
		if isEquivalentCondition(level, &child, cond) {
			if len(child.PARAMS) > 0 { // узел с прописанным PARAMS
				/* первая строка - для прямого правила,
				вторая - для учительского:
				*/
				if (child.Trigger == ObjectID && child.PARAMS[0] < 100) ||
					(child.Action == ObjectID && child.PARAMS[0] == 100) {
					//extremImportanceFind.objID = ObjectID  все равно приходится каждый раз создавать объект...
					//extremImportanceFind.extremVal = child.PARAMS[2]
					extremImportanceFind = &extremImportance{ObjectID, child.PARAMS[2]} //неизбежная утечка памяти...
					return                                                              // нашли
				}
			}
			// Если узел найден, продолжим рекурсивно искать в нем далее
			findEpisodicBrangeImportance(level+1, cond, &child, ObjectID)
			return
		}
		//  пусть перебирает дочки! return 0,level // не найдено совпадение на данном уровне
	}

	// не найден узел на данном уровне
	return
}

/////////////////////////////////////////////////////

/* для текущих условий выбрать самое значимое правило bestRule из всех найденных foundRulesArr
 */
var bestRule Rule        // результат поиска
var foundRulesArr []Rule // сбор всех найденных
func getBestRuleFromImpotrents() {
	foundRulesArr = nil
	bestRule.Count = 0 // маркер, что не найден
	cond := []int{CurrentCommonBadNormalWell, CurrentEmotionReception.ID, detectedActiveLastUnderstandingNodID}
	findBestRuleFromImpotrents(0, cond, &EpisodicTree)
	if len(foundRulesArr) > 0 { // найти наилучшую значимость
		choseBestRuleFromImpotrents()
	}

}
func findBestRuleFromImpotrents(level int, cond []int, root *EpisodicTreeNode) {
	// Поиск узла с ID из списка cond в дочерних узлах текущего узла
	for _, child := range root.Children { // пока есть дочки
		if isEquivalentCondition(level, &child, cond) {
			if len(child.PARAMS) > 0 { // узел с прописанным PARAMS
				if lib.Abs(child.PARAMS[2]) > 0 {
					foundRulesArr = append(foundRulesArr, Rule{child.Trigger, child.Action, child.PARAMS[0], child.PARAMS[1], child.PARAMS[2]})
				}
			}
			// Если узел найден, продолжим рекурсивно искать в нем далее
			findBestRuleFromImpotrents(level+1, cond, &child)
			return
		}
		//  пусть перебирает дочки! return 0,level // не найдено совпадение на данном уровне
	}

	// не найден узел на данном уровне
	return
}
func choseBestRuleFromImpotrents() {
	curI := 0
	for _, r := range foundRulesArr {
		if r.Importence > 0 {
			if r.Importence*r.Count > curI {
				curI = r.Importence * r.Count
				bestRule = r
			}
		}
	}
}

/////////////////////////////////////////////////////////

/*
знаком ли образ actID для данных условий? Новизна.
*/
func isUnknownActionsImage(actID int) bool {
	obj, _ := getObjectsImportanceValue(actID, detectedActiveLastProblemNodID)
	if obj != nil {
		return true
	}
	return false
}

//////////////////////////////////////////////////////////

/*
Найти действие в учительских правилах (стимул EpisodicTreeNode.Action),
которое имеет наивысшую значимость в данных условиях
чтобы использовать в качестве ответного действия (как элемент фантазирования).
*/
var chooseBestActionID = 0
var findPositiveActionArr []extremImportance

func findBestPositiveAction() int {
	chooseBestActionID = 0
	// прямо проходим условия
	cond := []int{CurrentCommonBadNormalWell, CurrentEmotionReception.ID, detectedActiveLastProblemNodID}
	findEpisodicBestPositiveAction(0, cond, &EpisodicTree)
	chooseBestAction()
	if chooseBestActionID > 0 {
		return chooseBestActionID
	}
	return 0
}
func findEpisodicBestPositiveAction(level int, cond []int, root *EpisodicTreeNode) {

	// Поиск узла с ID из списка cond в дочерних узлах текущего узла
	for _, child := range root.Children { // пока есть дочки
		if isEquivalentCondition(level, &child, cond) {
			if len(child.PARAMS) > 0 { // узел с прописанным PARAMS
				if (child.PARAMS[0] == 100) && (child.PARAMS[2] > 2) {
					findPositiveActionArr = append(findPositiveActionArr, extremImportance{child.Action, child.PARAMS[2] * child.PARAMS[1]})
				}
			}
			// Если узел найден, продолжим рекурсивно искать в нем далее
			findEpisodicBestPositiveAction(level+1, cond, &child)
			return
		}
		//  пусть перебирает дочки! return 0,level // не найдено совпадение на данном уровне
	}

	// не найден узел на данном уровне
	return
}
func chooseBestAction() {
	maxV := 0
	for _, ex := range findPositiveActionArr {
		if ex.extremVal > maxV {
			maxV = ex.extremVal
			chooseBestActionID = ex.objID
		}
	}
}

/////////////////////////////////////////////////////

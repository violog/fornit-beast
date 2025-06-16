/* Дерево понимания проблемы (дерево нерешенных проблем problemTree)

detectedActiveLastProblemNodID - ID образа текущей проблемы в данных условиях

формат записи: ID|ParentNode|autTreeID|situationTreeID|themeID|purposeID
активирующееся с каждым новым мент.циклом, с 4 уровнями:
ID дерева автоматизмов, ID дерева ситуации, ID темы, Id цели.
К этому дереву могут привязываться мент.автоматизмы, которые теперь становятся,
по примеру моторных, просто носителями последовательности ментальных проблем
*/

package psychic

// инициализирующий блок - в порядке последовательности инициализаций
// из psychic.go
func ProblemTreeInit() {
	if EvolushnStage < 4 { // только со стадии развития 4
		return
	}
	loadProblemTree()
	if len(ProblemTree.Children) == 0 { // еще нет никаких веток
		// создать базу
		createBasicProblemTree()
	}
}

/*
	эксклюзивное хранилище информации для активации дерева проблем

Дерево понимания активируется в func infoFunc8() только когда все переменные problemTreeInfo актуальны.
*/
type prTrInfo struct {
	autTreeID       int // устанавливается в func understandingSituation()
	situationTreeID int // устанавливается в func understandingSituation()
	themeID         int // устанавливается в func runNewTheme(). Это именно ThemeImage.ID!!!
	purposeID       int // устанавливается в getMentalPurpose()
}

var problemTreeInfo prTrInfo

/*
	// узел дерева проблем

Имеет фиксированных 4 уровней (кроме базового нулевого)
формат записи: ID|ParentNode|Mood|EmotionID|SituationID
Узлы всех уровней могут произвольно меняться на другие для переактивации Дерева.

Дерево может переактивароваться при срабатывании мент.автоматизмов с действиями
MentalActionsImages.activateBaseID и MentalActionsImages.activateEmotion
в mental_automatizm_actions.go RunMentalAutomatizm(
*/
type ProblemTreeNode struct { // узел дерева проблем
	ID              int
	autTreeID       int // образ автоматизма (узел дерева автоматизмов AutomatizmTreeFromID)
	situationTreeID int // образ понимания ситуации (узел дерева понимания ситуации UnderstandingNodeFromID)
	themeID         int // образ темы
	purposeID       int // образ цели. Это - не сознательная цель, а мотивирующая потребность, дающая направленность мышлению.

	Children   []ProblemTreeNode // дочерние узлы (ветвление) НЕ АДРЕСА, А РЕАЛЬНЫЕ ОБЪЕКТЫ
	ParentID   int               // ID родителя
	ParentNode *ProblemTreeNode  // адрес родителя
}

var ProblemTree ProblemTreeNode

// var ProblemTreeNodeFromID=make(map[int]*ProblemTreeNode)
var ProblemTreeNodeFromID []*ProblemTreeNode // сам массив
// var AFromID = make([]*aNode, len(strArr))//задать сразу имеющиеся в файле число при загрузке
// запись члена
func WriteProblemTreeNodeFromID(index int, value *ProblemTreeNode) {
	addProblemTreeNodeFromID(index)
	ProblemTreeNodeFromID[index] = value
}
func addProblemTreeNodeFromID(index int) {
	if index >= len(ProblemTreeNodeFromID) {
		newSlice := make([]*ProblemTreeNode, index+1)
		copy(newSlice, ProblemTreeNodeFromID)
		ProblemTreeNodeFromID = newSlice
	}
}

// считывание члена
func ReadeProblemTreeNodeFromID(index int) (*ProblemTreeNode, bool) {
	if index >= len(ProblemTreeNodeFromID) || ProblemTreeNodeFromID[index] == nil {
		return nil, false
	}
	return ProblemTreeNodeFromID[index], true
}

////////////////////////////////////////////

// последовательность узлов активной ветки
//var ActiveBranchProblemArr []int

/*
попытка активации дерева ментальных автоматизмов
*/
var detectedActiveLastProblemNodID = 0 // ID конечного узла активной ветки
var oldDetectedActiveLastProblemNodID = 0

// нераспознанный остаток - НОВИЗНА
var CurrentProblemTreeEnd []int
var currentProblemStepCount = 0

// массив узлов активной ветки   currentProblemActivedNodes=getcurrentProblemActivedNodes(lastID)
var currentProblemActivedNodes []*ProblemTreeNode // вначале конечный узел
// последовательность узлов активной ветки
var activeBranchProblemNodeArr []int // подряд

/*
	Активация дерева понимания проблемы ProblemTreeActivation

происходит при начале нового цикла мышления TODO
Вызов из infoFunc8()
*/
func ProblemTreeActivation() {

	if EvolushnStage < 4 { // только со стадии развития 4
		return
	}
	if PulsCount < 4 { // не активировать пока все не устаканится
		return
	}

	oldDetectedActiveLastProblemNodID = detectedActiveLastProblemNodID
	detectedActiveLastProblemNodID = 0

	activeBranchProblemNodeArr = nil
	CurrentProblemTreeEnd = nil
	currentProblemStepCount = 0
	currentProblemActivedNodes = nil
	///////// если еще не были активрованы деревья автоматизмов и ситуации (при просыпании)
	if detectedActiveLastNodID == 0 {
		automatizmTreeActivation()
	}

	// 4 уровня условий в виде ID их образов
	// определить ID проблемы
	// в infoFunc8()доблжет получаться дает mentalInfoStruct.mentalPurposeID

	var lev1 = problemTreeInfo.autTreeID
	var lev2 = problemTreeInfo.situationTreeID
	var lev3 = problemTreeInfo.themeID
	var lev4 = problemTreeInfo.purposeID

	// создать массив параметров ветки дерева
	condArr := getProblemActiveConditionsArr(lev1, lev2, lev3, lev4)
	conditionProblemFound(0, condArr, &ProblemTree)

	// результат поиска:
	if detectedActiveLastProblemNodID > 0 {
		// есть ли неучтенные условия?
		conditionsCount := getConditionsCount(condArr)
		CurrentProblemTreeEnd = condArr[currentProblemStepCount:] // НОВИЗНА
		/*Алексей: detectedActiveLastProblemNodID = formingProblemBranch(detectedActiveLastProblemNodID, currentProblemStepCount+1, condArr)
		Та же ситуация, что и с деревом автоматизмов. Почему currentProblemStepCount+1? Допустим, по этому 4-х уровневому дереву дошли до 3 уровня и остался еще один, которого там нет и его надо дорастить. Если оставить currentProblemStepCount+1 то передадим в функцию доращивания вместо 3 уровня 4 и соответственно никакого доращивания не будет. Но что то ты же имел в виду...
		Уберу этот +1 - пока не вижу в нем смысла.

		Тогда и в understanding_tree.go нужно убрать currentUnderstandingStepCount+1
		В func addNewUnderstandingBranchFromNodes( вообще можно начинать растить ветку с самого нуля,т.к. если такая есть, то createNewUnderstandingNode( просто вернет ее ID.
		Убрал в understanding_tree.go +1.
		*/
		if currentProblemStepCount < conditionsCount { // не пройдено до конца имеющихся условий
			// нарастить недостающее в ветке дерева
			detectedActiveLastProblemNodID = formingProblemBranch(detectedActiveLastProblemNodID, currentProblemStepCount, condArr)
		}
	} else { // вообще нет совпадений для данных условий
		// нарастить недостающее в ветке дерева
		detectedActiveLastProblemNodID = formingProblemBranch(detectedActiveLastProblemNodID, currentProblemStepCount, condArr)
		CurrentProblemTreeEnd = condArr // все - новизна
	}
	// все узлы активной ветки, в начале - конечный узел
	currentProblemActivedNodes = getcurrentProblemActivedNodes(detectedActiveLastProblemNodID)

	mainDreamCycle := IsDreamMainProcess()
	if mainDreamCycle != nil { //процесс пассивного размышления, не в ответ на Стимул
		//lastCommonDiffValue, _ := wasChangingMoodCondition()
		//fixDreamsRules(mainDreamCycle, lastCommonDiffValue)
	}

	return
}

func conditionProblemFound(level int, cond []int, node *ProblemTreeNode) {
	if cond == nil || len(cond) == 0 {
		return
	}

	ost := cond[1:]

	for n := 0; n < len(node.Children); n++ {
		cld := node.Children[n]
		var levID = 0
		switch level {
		case 0:
			levID = cld.autTreeID
		case 1:
			levID = cld.situationTreeID
		case 2:
			levID = cld.themeID
		case 3:
			levID = cld.purposeID
		}
		if cond[1] == levID {
			detectedActiveLastProblemNodID = cld.ID
			activeBranchProblemNodeArr = append(activeBranchProblemNodeArr, cld.ID)
		} else {
			currentProblemStepCount = level - 1
			continue
		}

		level++
		currentProblemStepCount = level
		conditionProblemFound(level, ost, &node.Children[n])
		return // раз совпало, то другие ветки не смотреть
	}
	return
}

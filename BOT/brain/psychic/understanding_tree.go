/* Дерево понимания ситуации

формат записи: ID|ParentNode|Mood|EmotionID|SituationID

При активации дерева ситуации текущая ситуация
UnderstandingNodeFromID[detectedActiveLastUnderstandingNodID].SituationID
и записывается в
curBaseStateImage.SituationID
*/

package psychic

import (
	"BOT/brain/gomeostas"
)

var noProblemTreeActivation = false // true - не активировать дерево проблем

// инициализирующий блок - в порядке последовательности инициализаций
// из psychic.go
func UnderstandingTreeInit() {
	if EvolushnStage < 4 { // только со стадии развития 4
		return
	}
	loadPurposeImageFromIdArr()
	loadUnderstandingTree()
	if len(UnderstandingTree.Children) == 0 { // еще нет никаких веток
		// создать первые три ветки базовых состояний
		createBasicUnderstandingTree()
	}
}

/*
	ДЕРЕВО понимания ситуации.

Имеет фиксированных 3 уровня (кроме базового нулевого)
формат записи: ID|ParentNode|Mood|EmotionID|SituationID
Узлы всех уровней могут произвольно меняться на другие для переактивации Дерева.

Дерево может переактивароваться при срабатывании мент.автоматизмов с действиями
MentalActionsImages.activateBaseID и MentalActionsImages.activateEmotion
в mental_automatizm_actions.go RunMentalAutomatizm(
*/
type UnderstandingNode struct { // узел дерева ситуации
	ID int
	//Mood = PsyBaseMood: -1 Плохое настроение, 0 Нормальное, 1 - хорошее настроение
	Mood      int // ощущаемое настроение МОЖЕТ ПРОИЗВОЛЬНО МЕНЯТЬСЯ
	EmotionID int // эмоция, МОЖЕТ ПРОИЗВОЛЬНО МЕНЯТЬСЯ
	/* SituationID определяет основной контекст ситуации, определяемый при вызове активации дерева понимания.
	   Если этот контекст не задан в understandingSituation(situationImageID
	   то в getCurSituationImageID() по-началу выбирается наугад (для первого приближения) более важные из существующих,
	   но потом дерево понимания может переактивироваться с произвольным заданием контекста.
	   От этого параметра зависит в каком направлении пойдет информационный поиск решений,
	   если не будет запущен штатный автоматизм ветки (ориентировочные реакции).
	*/
	SituationID int // ID объекта структуры понимания SituationImage, может произвольно меняться

	Children   []UnderstandingNode // дочерние узлы (ветвление) НЕ АДРЕСА, А РЕАЛЬНЫЕ ОБЪЕКТЫ
	ParentID   int                 // ID родителя
	ParentNode *UnderstandingNode  // адрес родителя
}

var UnderstandingTree UnderstandingNode

// var UnderstandingNodeFromID=make(map[int]*UnderstandingNode)
var UnderstandingNodeFromID []*UnderstandingNode // сам массив
// var AFromID = make([]*aNode, len(strArr))//задать сразу имеющиеся в файле число при загрузке
// запись члена
func WriteUnderstandingNodeFromID(index int, value *UnderstandingNode) {
	addUnderstandingNodeFromID(index)
	UnderstandingNodeFromID[index] = value
}
func addUnderstandingNodeFromID(index int) {
	if index >= len(UnderstandingNodeFromID) {
		newSlice := make([]*UnderstandingNode, index+1)
		copy(newSlice, UnderstandingNodeFromID)
		UnderstandingNodeFromID = newSlice
	}
}

// считывание члена
func ReadeUnderstandingNodeFromID(index int) (*UnderstandingNode, bool) {
	if index >= len(UnderstandingNodeFromID) || UnderstandingNodeFromID[index] == nil {
		return nil, false
	}
	return UnderstandingNodeFromID[index], true
}

//////////////////////////////////////////////////

// последовательность узлов активной ветки
//var ActiveBranchUnderstandingArr []int

// если в результате ментальных процессов было действие, то нужно заблокировать обработку активации дерева моторных автоматизмов
var MentalReasonBlocing = false

/*
попытка активации дерева ментальных автоматизмов
*/
var detectedActiveLastUnderstandingNodID = 0

// нераспознанный остаток - НОВИЗНА
var CurrentUnderstandingTreeEnd []int
var currentUnderstandingStepCount = 0

// массив узлов активной ветки   currentUnderstandingActivedNodes=getcurrentUnderstandingActivedNodes(lastID)
var currentUnderstandingActivedNodes []*UnderstandingNode // вначале конечный узел
// последовательность узлов активной ветки подряд
var activeUnderstandingNodeArr []int

// текущие образы  гомеостатической этиологии, колторые могут быть произвольно перекрыты ментальными образами
var newMoodID = 0
var newEmotionID = 0
var newSituationID = 0

// сохраненные образы гомеостатической этиологии, колторые могут быть произвольно перекрыты ментальными образами
var preMoodID = 0
var preEmotionID = 0
var preSituationID = 0

/*
	Активация дерева понимания ситуации происходит из:

func afterTreeActivation() - при каждой активации automatism_tree.go
и если было действия без ответа в течении 20 пульсов, то understandingSituation вызывается из
func noAutovatizmResult()
т.е. оба деревав работают совместно при EvolushnStage > 3 и по каждой активации UnderstandingTree
добавляется эпизд памяти newEpisodeMemory()

Аналогично дереву моторных автоматзмов, после активации могут быть ориентировочные рефлексы привлечения внимания.

При вызове может быть определен situationImageID или проставлен 0 и тогда образ ситуации определяется в самой функции.

Если были совершены действия, то нужно выставлять MotorTerminalBlocking=true !!!

activationType =1 - объективная активация
activationType =2 - произвольная переактивация

return true// заблокировать все низкоуровневое
*/
func understandingSituation(activationType int) {
	MentalReasonBlocing = false // true - заблокировать все рефлексы и штатные автоматизмы

	if EvolushnStage < 4 { // только со стадии развития 4
		return
	}
	if PulsCount < 4 { // не активировать пока все не устаканится
		return
	}
	// определить ID ситуации: настроение при посылке сообщения, нажатые кнопки и т.п.
	situationImageID := getCurSituationImageID()
	if situationImageID == 0 { // нет выбранной ситуации
		return
	}
	/*
		BaseStateImageMapCheck()
		if BaseStateImageArr != nil{ // перекрыть curBaseStateImage
			curBaseStateImage = *BaseStateImageArr[situationImageID]
			curBaseStateImage.SituationID = situationImageID
		}*/

	//	ps:=getPurposeGenetic() // - тут уже сохраняется savePurposeGenetic
	//	savePurposeGenetic=ps

	detectedActiveLastUnderstandingNodID = 0 // ID конечного узла активной ветки

	activeUnderstandingNodeArr = nil
	CurrentUnderstandingTreeEnd = nil
	currentUnderstandingStepCount = 0
	currentUnderstandingActivedNodes = nil

	// текушие гомео-зависимые параметы
	newMoodID = PsyBaseMood

	bsIDarr := gomeostas.GetCurContextActiveIDarr()
	newEmotionID, _ = createNewBaseStyle(0, bsIDarr, true)

	/* не просрочены ли произвольно активированные параметры
	Держатся на время, пока не изменятся генетически определенные соотвествующие параметры или
	если активация была в данном пульсе
	*/
	if mentalMoodVolitionPulsCount != PulsCount && preMoodID != newMoodID {
		mentalMoodVolitionID = 0
	}
	if mentalEmotionVolitionPulsCount != PulsCount && preEmotionID != newEmotionID {
		mentalEmotionVolitionID = 0
	}
	if mentalSituationVolitionPulsCount != PulsCount && preEmotionID != newSituationID {
		mentalSituationVolitionID = 0
	}

	// сохранять только изменившиеся значения
	if preMoodID != newMoodID {
		preMoodID = newMoodID
	}
	if preEmotionID != newEmotionID {
		preEmotionID = newEmotionID
	}
	if preSituationID != newSituationID {
		preSituationID = newSituationID
	}

	// 3 уровня условий в виде ID их образов

	var lev1 = newMoodID // PsyBaseMood: -1 Плохое настроение, 0 Нормальное, 1 - хорошее настроение
	if mentalMoodVolitionID > 0 {
		lev1 = mentalMoodVolitionID
	}

	var lev2 = newEmotionID
	if mentalEmotionVolitionID > 0 {
		lev2 = mentalEmotionVolitionID
	}

	var lev3 = situationImageID
	if mentalSituationVolitionID > 0 {
		lev3 = mentalSituationVolitionID
	}

	condArr := getUnderstandingActiveConditionsArr(lev1, lev2, lev3)
	// основа дерева
	cnt := len(UnderstandingTree.Children)
	for n := 0; n < cnt; n++ {
		node := UnderstandingTree.Children[n]
		lev1 := node.Mood
		if condArr[0] == lev1 {
			detectedActiveLastUnderstandingNodID = node.ID
			ost := condArr[1:]
			if len(ost) == 0 {

			}

			conditionUnderstandingFound(1, ost, &node)

			break // другие ветки не смотреть
		}
	}

	// результат поиска:
	if detectedActiveLastUnderstandingNodID > 0 {
		// есть ли неучтенные условия?
		conditionsCount := getConditionsCount(condArr)
		CurrentUnderstandingTreeEnd = condArr[currentUnderstandingStepCount:] // НОВИЗНА
		if currentUnderstandingStepCount < conditionsCount {                  // не пройдено до конца имеющихся условий
			// нарастить недостающее в ветке дерева
			detectedActiveLastUnderstandingNodID = formingUnderstandingBranch(detectedActiveLastUnderstandingNodID, currentUnderstandingStepCount, condArr)
		}
	} else { // вообще нет совпадений для данных условий
		// нарастить недостающее в ветке дерева
		detectedActiveLastUnderstandingNodID = formingUnderstandingBranch(detectedActiveLastUnderstandingNodID, currentUnderstandingStepCount, condArr)
		CurrentUnderstandingTreeEnd = condArr // все - новизна
	}

	//инфа для активации дерева проблем:
	problemTreeInfo.autTreeID = detectedActiveLastNodID
	problemTreeInfo.situationTreeID = situationImageID

	// все узлы активной ветки, в начале - конечный узел
	currentUnderstandingActivedNodes = getcurrentUnderstandingActivedNodes(detectedActiveLastUnderstandingNodID)

	// здесь - новые темы при активации деревьев, т.к. для дерева проблем нужен ID дерева ситуации
	if activationType == 1 && needRunNewTheme > 0 {
		if ActivationTypeSensor > 1 { // есть инфа о curActions - параметры действий оператора
			NewTeme() // тут всегда обнулять тему т.к. - новое воздействие оператора
		}
		runNewTheme(needRunNewTheme, 2) //после активации дерева ситуации
	}

	/* объективный запуск consciousnessElementary - по акативации дерева автоматизмов
	ментальный запуск в случае произвольной переактивации дерева понимания или цикле осмысления (5-я ступень)

	При вызове  infoFunc8() происходит активация дерева проблем и НЕ ДОЛЖНА БЫТЬ новая об.активация осмысления!
	*/
	if !noProblemTreeActivation {
		activationType = 1
		// перезапустить дерево проблем
		MentalReasonBlocing = consciousnessElementary() // тут перезапускается главный цикл мышления
	} else { // была переактивация дерева из func infoFunc8
		if wasRunPurposeActionFunc { // перективация запущена из func infoFunc14
			wasRunPurposeActionFunc = false
			// начать новый главный цикл мышления
			resetMineCycleAndBeginAsNew()
		}
	}

	return
}

func conditionUnderstandingFound(level int, cond []int, node *UnderstandingNode) {
	if cond == nil || len(cond) == 0 {
		return
	}

	ost := cond[1:]

	for n := 0; n < len(node.Children); n++ {
		cld := node.Children[n]
		var levID = 0
		switch level {
		case 0:
			levID = cld.Mood
		case 1:
			levID = cld.EmotionID
		case 2:
			levID = cld.SituationID
		}
		if cond[0] == levID {
			detectedActiveLastUnderstandingNodID = cld.ID
			activeUnderstandingNodeArr = append(activeUnderstandingNodeArr, cld.ID)
		} else {
			currentUnderstandingStepCount = level - 1
			continue
		}

		level++
		currentUnderstandingStepCount = level
		conditionUnderstandingFound(level, ost, &node.Children[n])
		return // раз совпало, то другие ветки не смотреть
	}
	return
}

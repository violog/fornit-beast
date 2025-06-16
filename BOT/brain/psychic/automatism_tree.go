/*  Дерево автоматизмов

Все начинается с psychic.go (atomatizmID:=automatizmTreeActivation()) -> func automatizmTreeActivation()

Это дерево активируется при:
1. Всегда при любых событиях с Пульта – так же как дерево рефлексов, но если к ветке привязан автоматизм,
то он выполняется преимущественно, блокируя рефлексы потому,
что уже было произвольностью преодолено действие рефлекса при выработке автоматизма.
Такой автоматизм обладает меткой успешности ==1. Успешность ==0 означает предположительный вариант действий,
а успешность < 0 – заблокированный вариант действия.
Так что к ветке может быть прикреплено множество неудачных и предположительных автоматизмов
и только один удачный. Более удачный результат переводит ранее удачный автоматизм в предполагаемые.
2. При произвольной активации отдельных условий.
Отсуствие подходящей для данных условий ветки дерева вызывает
Ориентировочный рефлекс привлечения внимания к активной ветке с осмыслением ситуации
и рассмотрением альтернатив действиям (4 уровня глубины рассмотрения).
При формировании нового предположительного действия создается новая ветка дерева и к ней прикрепляется автоматизм.
Т.е. новые условия не создают новой ветки, а тольно - новый автоматизм,
а пока нет автоматизма будет ориентировочный рефлекс.

У дерева фиксированных 6 уровней:
0 нулевой - основание
1 Базовое состояние - 3 вида
2 Эмоция
3 Активность с Пульта - образ ActivityFromIdArr=make(map[int]*Activity)
4 Образ контекста сообщения: сочетание Tone и Mood из структуры vrbal
5 Первый символ фразы
6 Фраза - PhraseID
До 6-го уровня - полный аналог условным рефлексам, только вместо сочетаний контекстов - эмоция.

Для оптимизации поиска по дереву перед узлом Verbal идет узел первого символа : var symbolsArr из word_tree.go

Формат записи:
ID|ParentNode|BaseID|EmotionID|ActivityID|ToneMoodID|SimbolID|PhraseID


Самоадаптация уровня Дерева автоматизмов
В результате действия автоматизма могут измениться условия и, значит,
будут запущены дерево рефлексов и опять - Дерево автоматизмов.
Возникает новая итерация адаптивности, возможно, с новым ориентировочным рефлексом второго типа.
Такой процесс может продолжаться до прихода к устойчивому состоянию.

*/

package psychic

import (
	"BOT/brain/action_sensor"
	"BOT/brain/gomeostas"
	wordSensor "BOT/brain/words_sensor"
	"BOT/lib"
)

// психика инициализирована
var StartPsichicNow = false

// инициализирующий блок - в порядке последовательности инициализаций
// из psychic.go
func automatizmTreeInit() {

	loadAutomatizmTree()
	if len(AutomatizmTree.Children) == 0 { // еще нет никаких веток
		// создать первые три ветки базовых состояний
		createBasicAutomatizmTree()
	}
	StartPsichicNow = true
}

/////////////////////////////////////////////////////////////

////////////////////////////////////////////

// //// ДЕРЕВО автоматизмов имеет фиксированных 6 уровней (кроме базового нулевого)
type AutomatizmNode struct { // узел дерева автоматизмов
	ID     int
	BaseID int // 1 - Похо, 2 - Норма, 3 - Хорошо, базовое состояние, !это еще не произвольно меняющееся PsyBaseMood
	/* эмоция (type Emotion struct) Эмоция может произвольно меняться, независимо от базовых контекстов
	   т.е., к примеру, при BaseID Плохо может быть позитивное EmotionID
	*/
	EmotionID  int
	ActivityID int // образ сочетания действия с Пульта
	/* образ контекста сообщения: сочетание Tone и Mood из структуры vrbal из automatism_tree_verbal_img.go
	   т.е. просто toneID+moodID - в виде строки, например: "922" = "Обычный, Хорошее"
	   дешифруется func getToneMoodStrFromID(id string)(string)
	*/
	ToneMoodID int
	SimbolID   int
	PhraseID   int // Verbal.ID // массив фраз. Может быть длинное сообщение из нескольких фраз.

	Children   []AutomatizmNode // дочерние узлы (ветвление) НЕ АДРЕСА, А РЕАЛЬНЫЕ ОБЪЕКТЫ
	ParentID   int              // ID родителя
	ParentNode *AutomatizmNode  // адрес родителя
}

var AutomatizmTree AutomatizmNode

//var AutomatizmTreeFromID=make(map[int]*AutomatizmNode)
//var MapGwardAutomatizmTreeFromID=lib.RegNewMapGuard()
///////////////////////////////////////

var AutomatizmTreeFromID []*AutomatizmNode // узел по его ID
// var AutomatizmTreeFromID = make([]*AutomatizmNode, 20000)//задать сразу имеющиеся в файле число
// запись члена
func WriteAutomatizmTreeFromID(index int, value *AutomatizmNode) {
	if index >= len(AutomatizmTreeFromID) {
		newSlice := make([]*AutomatizmNode, index+1)
		copy(newSlice, AutomatizmTreeFromID)
		AutomatizmTreeFromID = newSlice
	}
	AutomatizmTreeFromID[index] = value
}

// считывание члена
func ReadeAutomatizmTreeFromID(index int) (*AutomatizmNode, bool) {
	if index >= len(AutomatizmTreeFromID) || AutomatizmTreeFromID[index] == nil {
		return nil, false
	}
	return AutomatizmTreeFromID[index], true
}

// последовательность узлов активной ветки
var ActiveBranchNodeArr []int

////////////////////////////////////////////////

// временная структура действий оператора (всегда с ID=0) для формирования постоянного образа curActiveActions
var curActions ActionsImage //
// структура действий оператора при активации дерева автоматизмов
var curActiveActions *ActionsImage // зеркалит текущий ActionsImage
var curActiveActionsID = 0         // ID запускаемого образа действия

// Уже есть curActions.PhraseID var curActivePhraseID=0// текущая фраза с пульта, если она есть
var curActiveTreePulsCount = 0    // время активации дерева автоматизмов в пульсах, т.е. прихода любого вида стимула
var curActiveActionsPulsCount = 0 // время действия оператора
// образ предыдущего сосотояния Стимула ПОСЛЕ стимула Оператора (не меняется при активации изменением состояния)
var curStimulImage *ActionsImage
var curStimulImageID = 0

// Допустимое число пульсов после Стимула чтобы считать его ответом на действие
// ! НО УЖЕ ОПРЕДЕЛЕН ПЕРИОД ОЖИДАНИЯ ОТВЕТА - LastRunAutomatizmPulsCount, зачем еще этот??
var limitOfActionsAfterStimul = 5 //предел действий после стимула - не более 5 пульсов

// //////////////////////////////////////
var curActiveVerbalID = 0

// эффект от стимула Оператора
var globalStimulsEffect = 0
var prevStimulsEffect = 0 // предыдущий эффект

/*
	Нужна активаци новой Темы с ID==nn.

вызвать runNewTheme(needRunNewTheme,2)

	после активации дерева ситуации т.к. для дерева проблем нужен ID дерева ситуации
*/
var needRunNewTheme = 0

////////////////////////////////////////////////////////////////////////////////////////
/* попытка активации дерева автоматизмов, если неудачно - начать искать вариант действий
Используется активная текущая информационная среда из psychic.go:
var PsyBaseID=0 // текущее базовое состояние, может быть произвольно изменено
var PsyEmotionImg *Emotion // текущая эмоция Emotion, может быть произвольно изменена
var PsyActionImg *Activity // текущий образ сочетания действий с Пульта Activity
var PsyVerbImg *Verbal // текущий образ фразы с Пульта Verbal
*/
var detectedActiveLastNodID = 0

// последний detectedActiveLastNodID где была распознана фраза без isUnrecognizedPhraseFromAtmtzmTreeActivation==true
var lastWellPhrasedetectedActiveLastNodID = 0

// при запуске автоматизма по действию оператора, для Правил
var detectedActiveLastNodPrevID = 0
var detectedActiveLastUnderstandingNodPrevID = 0

// нераспознанный остаток - НОВИЗНА
var CurrentAutomatizTreeEnd []int
var currentStepCount = 0
var currentAutomatizmAfterTreeActivatedID = 0 //! это  не обязательно штатный автоматизм ветки, а выбранный мягким алгоритмом
var wasCurrentAutomatizmAfterTree = 0         // для активации функции consciousnessElementary: 1 этот автоматизм из текущей активации дерева, есть currentAutomatizmAfterTreeActivatedID

var curActivePhraseStr = ""
var isUnrecognizedPhraseFromAtmtzmTreeActivation = false //true - при активации была нераспознанная фраза

/*
	Детектор: активация дерева не вызвала автоматизм и не было периода ожидания.

только после действий оператора (а не активация по изменению гомео-параметров)
если ==2 - активация дерева не вызвала автоматизм и не было периода ожидания
*/
var NoautomatizmAfterStimul = 0

/////////////////////////////////

// активация дерева автоматизмов
func automatizmTreeActivation() int {

	if PulsCount < 4 { // не активировать пока все не устаканится
		return 0
	}

	if isRepressionStimulsNoise { //Подавление мешающих стимулов при серьезном поиске решений проблемы
		//контроль удержания режима isRepressionStimulsNoise
		checkRepressionStimulsNoises() // можно ли снять режим
	}

	if IsSleeping {
		return 0
	}
	/* НУЖНО, просто новый ор.рефлекс будет ждать окончания периода LastRunAutomatizmPulsCount
	if LastRunAutomatizmPulsCount >0{// не активировать в период ожидания результатов действий!
		return 0
	}
	*/

	/* ТЕПЕПЕРЬ ВСЕГДА АКТИВИРОВАТЬ потому как и по изменению состояния формируются Правила.
	   Но нужно блокировать ор.рефлексы!
	   // не активировать дерево по изменению гомеостатуса во время ожидания ответа оператора
	   //  LastRunAutomatizmPulsCount устанавливается в RumAutomatizm(
	   if LastRunAutomatizmPulsCount > 0{
	   if !WasOperatorActiveted {
	   	return 0
	   }
	*/

	curActiveTreePulsCount = PulsCount // время активации дерева автоматизмов, т.е. прихода любого вида стимула

	detectedActiveLastNodID = 0
	needRunNewTheme = 0

	ActiveBranchNodeArr = nil
	CurrentAutomatizTreeEnd = nil
	currentStepCount = 0
	currentAutomatizmAfterTreeActivatedID = 0
	isUnrecognizedPhraseFromAtmtzmTreeActivation = false

	// вытащить 3 уровня условий в виде ID их образов
	//Еще нет InformationEnvironment т.к. Дерево активруется ДО ор.рефлексов
	lev1 := gomeostas.CommonBadNormalWell
	oldCommonBadNormalWell = CurrentCommonBadNormalWell
	CurrentCommonBadNormalWell = lev1

	bsIDarr := gomeostas.GetCurContextActiveIDarr()
	lev2, emotion := createNewBaseStyle(0, bsIDarr, true)
	oldCurrentEmotionReception = CurrentEmotionReception
	CurrentEmotionReception = emotion

	/* Для дерева автоматизмов используютсЯ ID отдельных видов действий,
	но уже есть и общий ID стимула (который используется в дереве рефлексов)
	ActiveCurTriggerStimulsID - он испоьзуется в дереве эпиз.пяти как Стимул: EpisodicTreeNode.Trigger
	*/

	ActID := action_sensor.CheckCurActionsContext() //CheckCurActions()
	curActions.ActID = ActID
	lev3, _ := createNewlastActivityID(0, ActID, true) // текущий образ сочетания действий с Пульта Activity

	// дезактивировать все контексты!!!! чтобы не влияли на следующую активность
	action_sensor.DeactivationTriggersContext()
	//!!!!curActiveActionsID = 0
	//!!!!curActiveActions = nil

	var lev4 = 90 // нельзя 0, это ToneMood
	var lev5 = 0
	var lev6 = 0
	if len(wordSensor.CurrentPhrasesIDarr) > 0 {
		PhraseIDarr := wordSensor.CurrentPhrasesIDarr
		//Если CurrentPhrasesIDarr[n]==-1 - фраза есть, но она нераспознана.
		for i := 0; i < len(PhraseIDarr); i++ {
			if PhraseIDarr[i] == -1 { // фраза нераспознанна
				isUnrecognizedPhraseFromAtmtzmTreeActivation = true
			}
		}

		// нераспознанная фраза - все равно - наличие фразы, а не молчание и образуется нормальный CreateVerbalImage()

		FirstSimbolID := wordSensor.GetFirstSymbolFromPraseID(PhraseIDarr)
		ToneID := wordSensor.CurPultTone
		if ToneID == 0 && wordSensor.DetectedTone == 1 { // повышенный из-за знака "!"
			ToneID = 4 // 4- разница в кодировке тона
		}
		MoodID := wordSensor.CurPultMood
		verbID, verb := CreateVerbalImage(FirstSimbolID, PhraseIDarr, ToneID, MoodID)
		if verb != nil {
			lev4 = GetToneMoodID(verb.ToneID, verb.MoodID)
			lev5 = verb.SimbolID
			/* для дерева берется только первая фраза КАК ГЛАВНОЕ УСЛОВИЕ ВЕТКИ, остальные можно восстановить для сопоставлений из
			AutomatizmNode.PhraseID.PhraseIDarr[]
			*/
			lev6 = verb.ID
		}
		// в четвертой стадии обнулится в fixEpizMemoryTeachRules(). Если здесь обнулять, то в fixEpizMemoryTeachRules() пойдет фраза nil
		if EvolushnStage < 4 {
			// очистить фразу после использования, чтобы не влияла на следующую активность
			wordSensor.CurrentPhrasesIDarr = nil // остается еще wordSensor.CurretWordsIDarr
		}

		// сохраняем для отзеркаливания действий оператора
		curActiveVerbalID = verbID
		curActions.PhraseID = PhraseIDarr
		curActions.ToneID = ToneID
		curActions.MoodID = MoodID
	} else { // иначе при ответе ТОЛЬКО через действие, в правилах запишется фраза от предыдущего ответа, хотя на самом деле там не должно быть фразы вообще
		curActions.PhraseID = nil
		curActions.ToneID = 0
		curActions.MoodID = 0
		curActiveVerbalID = 0
	}

	if ActivationTypeSensor == 1 { // активация по изменению гомео-параметров
		/*	в understanding.go func consciousnessElementary()
			блокируются активаци при незначительных извенениях гомеопараметров if !gomeostas.GetGomeoParsDiff() {
			и если прошло меньше 10 пульсов со времени последнего стимула от оператора
		*/
		if isRepressionStimulsNoise { //Подавление мешающих стимулов при серьезном поиске решений проблемы
			// сторожевая функция определения критического состояния гомеостаза
			watchdogFunctionGomeo()
		}
	} else { // образ действий оператора

		// зафиксировать время появления Стимула и ждать запуска автоматизма 2 пульса, иначе detectedActiveLastNodID=2
		NoautomatizmAfterStimul = PulsCount // хоть и синоним curActiveActionsPulsCount
		// Внизу - после обработки clinerAutomatizmRunning() - новый симул сбрасывает ожидание ответа на старый

		// сохраняем предыдущий Стимул
		curStimulImage = curActiveActions
		curStimulImageID = curActiveActionsID

		curActiveActionsID, curActiveActions = CreateNewlastActionsImageID(0, 0, curActions.ActID, curActions.PhraseID, curActions.ToneID, curActions.MoodID, true)
		curActiveActionsPulsCount = PulsCount
		stimulCount++ // сколько раз был стимул от оператора после последнего запуска Ответа
		CurrentInformationEnvironment.ActionsImageID = curActiveActionsID

		if isRepressionStimulsNoise { //Подавление мешающих стимулов при серьезном поиске решений проблемы
			// сторожевая функция определения очень важного объекта внимания
			watchdogFunctionStimul(curActiveActions)
		}
	}

	/* Если данный curActiveActions содержит экстремальный объект с высокй значимостью,
	то это повлияет на базовые контексты, переактивировав их через func ContextActiveFromPsy()
	но не сразу, а консервативно с заданной степень консервативности.
	*/
	getExtremImportanceObject()
	if extremImportanceObject != nil && lib.Abs(extremImportanceObject.extremVal) > 5 {
		// переактивировать базовые контексты при conservatismPsyStymulEffect повторений такой значимости
		if gomeostas.ContextActiveFromPsy(extremImportanceObject.extremVal) { // была переактивация
			/* Посте такой переактивации будет опять запущена активация дерева func changingConditionsDetector()
			ну и пусть. А пока что пусть выполнится привычный автоматизм по прежнему условию.
			*/
			setInterruptionEpisosde() //вставить пустой кадр эпиз.памяти - прервать тему
			clinerAutomatizmRunning() // окончить период ожидания
		}
	}

	//Подавление мешающих стимулов при серьезном поиске решений проблемы if EvolushnStage > 4
	if isRepressionStimulsNoise {
		runIgnoreAction() // запустить действие Игнорировать
		return 0
	}

	condArr := getActiveConditionsArr(lev1, lev2, lev3, lev4, lev5, lev6)
	notAllowScanInTreeThisTime = true // защелка от повтора во время обработки
	// основа дерева
	cnt := len(AutomatizmTree.Children)
	for n := 0; n < cnt; n++ {
		node := AutomatizmTree.Children[n]
		lev1 := node.BaseID
		if condArr[0] == lev1 {
			detectedActiveLastNodID = node.ID
			ost := condArr[1:]
			if len(ost) == 0 {

			}

			conditionAutomatizmFound(1, ost, &node)

			break // другие ветки не смотреть
		}
	}

	// результат активации Дерева:
	if detectedActiveLastNodID > 0 {
		// есть ли еще неучтенные, нулевые условия? т.е. просто показаь число ненулевых значений condArr
		conditionsCount := getConditionsCount(condArr)
		CurrentAutomatizTreeEnd = condArr[currentStepCount:] // НОВИЗНА
		if currentStepCount < conditionsCount {              // не пройдено до конца имеющихся условий
			// нарастить недостающее в ветке дерева - всегда для orientation_1()
			//oldDetectedActiveLastNodID:=detectedActiveLastNodID
			detectedActiveLastNodID = formingBranch(detectedActiveLastNodID, currentStepCount, condArr)

			/* !!! если нераспознана фраза, то detectedActiveLastNodID НЕВЕРЕН его нельзя использовать как поиск по эп.памяти
			поэтому нужно использовать lastWellPhrasedetectedActiveLastNodID
			*/
			if !isUnrecognizedPhraseFromAtmtzmTreeActivation {
				// последний detectedActiveLastNodID где была распознана фраза без isUnrecognizedPhraseFromAtmtzmTreeActivation==true
				lastWellPhrasedetectedActiveLastNodID = detectedActiveLastNodID
			}
		}

	} else { // вообще нет совпадений для данных условий
		// нарастить недостающее в ветке дерева - всегда для orientation_1()
		detectedActiveLastNodID = formingBranch(detectedActiveLastNodID, currentStepCount, condArr)

		// автоматизма нет у недоделанной ветки
		CurrentAutomatizTreeEnd = condArr // все - новизна

	}

	if afterTreeActivation() {
		notAllowScanInTreeThisTime = false // снять блокировку
		return 1
	}
	notAllowScanInTreeThisTime = false // снять блокировку
	return 0
}

//////////////////////////////////////////////////////////////////

func conditionAutomatizmFound(level int, cond []int, node *AutomatizmNode) {
	if cond == nil || len(cond) == 0 {
		return
	}

	ost := cond[1:]

	for n := 0; n < len(node.Children); n++ {
		cld := node.Children[n]
		var val = 0
		switch level {
		case 0:
			val = cld.BaseID
		case 1:
			val = cld.EmotionID
		case 2:
			val = cld.ActivityID
		case 3:
			val = cld.ToneMoodID
		case 4:
			val = cld.SimbolID
		case 5:
			val = cld.PhraseID
		}
		if cond[0] == val {
			detectedActiveLastNodID = cld.ID
			ActiveBranchNodeArr = append(ActiveBranchNodeArr, cld.ID)
		} else {
			currentStepCount = level - 1
			continue
		}

		level++
		currentStepCount = level
		conditionAutomatizmFound(level, ost, &node.Children[n])
		return // раз совпало, то другие ветки не смотреть
	}

	return
}

////////////////////////////////////////////////////////

var onliOnceWasConditionsActiveted = false // т.к. опять может продолжиться изменение состояния в период ожидания
/*
	реакция после активации ветки дерева

если нет никаких действий, то возвращает false, инчае - true для блокировки более низкоуровневого
*/
func afterTreeActivation() bool {
	/* Нельзя здесь определять currentAutomatizmAfterTreeActivatedID перед if LastRunAutomatizmPulsCount >0{
	// ЕСТЬ ЛИ АВТОМАТИЗМ В ВЕТКЕ и болеее ранних? выбрать лучший автоматизм для сформированной ветки nodeID
	currentAutomatizmAfterTreeActivatedID = getAutomatizmFromNodeID(detectedActiveLastNodID)
	*/

	// по каждому действию оператора задумываться
	if ActivationTypeSensor > 1 { // есть инфа о curActions - параметры действий оператора
		needRunNewTheme = 11 //"Действие оператора"
	}

	/*

	 */

	//var wasRunUnderstandingSituation=false
	/*ПЕРИОД ОЖИДАНИЯ ОТВЕТА ОПЕРАТОРА, реагировать только на действия Оператора ActivationTypeSensor >1
	  	Был запущен моторный автоматизм (в том числе и ментальным автоматизмом)
	  Срабатывает при типе активации (ActivationTypeSensor>1) т.к. Правила записываются только
	  	со стимулом от Оператора и НЕ бывает со стимулом - по изменению состояния.
	*/

	if LastRunAutomatizmPulsCount > 0 && ActivationTypeSensor > 1 { //Обработка нового ответа оператора
		effect := 0
		// 	Контроль за изменением состояния, возвращает:
		//	lastCommonDiffValue - насколько изменилось общее состояние
		//  	lastBetterOrWorse - стали лучше или хуже: величина измнения от -10 через 0 до 10
		//  	gomeoParIdSuccesArr - стали лучше следующие г.параметры []int гоменостаза
		if WasOperatorActiveted { // оператор отреагировал
			lastCommonDiffValue, gomeoParIdSuccesArr := wasChangingMoodCondition() // здесь stimulsEffect
			/* обработать изменение состояния фиксация ПРАВИЛА, Стимул - ОТ ОПЕРАТОРА
			если эффектом действий кнопок заблокируется автоматизм, клонированный от рефлекса и повешенный на кнопку
			будет срабатывать его родитель-рефлекс и оказывать обучающее действие, через изменение состояния инфо-среды
			LastAutomatipmCorrection() отдельно отрабатывает эффект учительских кнопок Наказать/Поощрить,
			поэтому здесь это блокируем - иначе будет двойная оценка
			*/
			//if !action_sensor.IsPress3or4button {
			calcAutomatizmResult(lastCommonDiffValue, gomeoParIdSuccesArr)
			effect = lastCommonDiffValue
			/*} else {
				effect = effectPress3or4button
				effectPress3or4button = 0
			}*/
			// по результатам обработки, но до очистки 	LastRunAutomatizmPulsCount и LastAutomatizmWeiting
			if EvolushnStage > 3 {
				//	wasRunUnderstandingSituation=true
				/* Здесь не активировать Дерево пониманияситуации. Активировать только 1 раз (зачем 2 раза??)- внизу,
				   только после определения currentAutomatizmAfterTreeActivatedID
				   ведь будет вызов consciousnessElementary() - обработка ситуации с currentAutomatizmAfterTreeActivatedID !!!
				*/
				//understandingSituation(1)
				// !!!return true
			}
			//consciousnessThinking(8)

			// МЕНТАЛЬНЫЕ ПРАВИЛА после периода ожидания
			afterWaitingPeriod(effect) //Учесть последствия ментального запуска мот.автоматизма

			//и если нужно обдумать это
			/* позитивный эффект не обдумывать
			if effect >= 0 && curActiveActions.PhraseID !=nil && len(curActiveActions.PhraseID)>0{
				runNewTheme(11,1)
			}
			*/

			// закончить период ожидания после реакции оператора
			clinerAutomatizmRunning()
			WasConditionsActiveted = false // иначе сразу сработает fixRulesBaseStateImage после изменения состояния при действиях
		}

		/* Не записывать Правила по изменению состояния, а только - по стимулу от Оператора!
		   потому как состояние может улучшить только оператор, так что после его Стимула смотрим результат и оцениваем действие.
		   	if !onliOnceWasConditionsActiveted {// только один раз во время периода ожидания
		   		onliOnceWasConditionsActiveted = true
		   		if WasConditionsActiveted { // изменились условия (не действия оператора)
		   			WasConditionsActiveted = false
		   			if EvolushnStage > 3 {
		   				lastCommonDiffValue, _, _ := wasChangingMoodCondition(2)
			   // записать ПРАВИЛО типа BaseStateImage Стимул - НЕ ОТ ОПЕРАТОРА, а при активации изменением состояния
		   				fixRulesBaseStateImage(lastCommonDiffValue)// здесь корректируется успешность автоматизма - как в calcAutomatizmResult
		   				// Активировать Дерево Понимания: или запустить ментальный автоматизм или - ориентировочная реакция для осмысления
		   				understandingSituation(1)

		   				// НЕ заканчивать период ожидания после переактивации по изменившимся условиям, но не запускать ор.рефлекс:
		   				return true
		   			}
		   		}
		   	}*/

		// осмыслить результат, задав тему ThemeImage.Type=1
		if effect < 0 { // негативный эффект Стимула
			// -effect обычно большой и уже не перекрывается другой темой, так что задаем фиксированный 2
			//runNewTheme(2,2)
			needRunNewTheme = 2 //Негативный эффект ментального автоматизма
		}

		//  после обработки ожидаемой реакции Оператора - следует реакция Beast
		//		return true  поэтому нельзя здесь делать прерывание!
	}
	////////////////////////////// конец обработки ожидания ответа оператора

	/* ЕСТЬ ЛИ АВТОМАТИЗМ В ВЕТКЕ и болеее ранних? выбрать лучший автоматизм для сформированной ветки nodeID
		а если нет, то учитывать общие автомтизмы, привязанные к действиям (виртуальная ветка ID от 1000000) и словам (>2000000)

	Нужно учесть случай, что если на авторитарный ответ оператора вдруг срабатывает автоматизм,
		его и надо на пульт отправить. Иначе отправится сгенеренный зеркальный.
		Есть ведь ситуации, когда стимул новый, а ответ старый.
		Например, привет - хай, приветствую - хай. На два разных стимула может быть автоматизмы с одинаковыми действиями, но привязанные к разным узлам дерева.
		Без учета этого бот не сделает НОВЫЙ автоматизм, а просто выдаст старый.
	*/
	if oldAtmzAfterTreeActivatedID == 0 {

		currentAutomatizmAfterTreeActivatedID = getAutomatizmFromNodeID(detectedActiveLastNodID)

	} else {
		// ID автоматизма, активировашегося на стимул ДО создания зеркального в calcAutomatizmResult()
		currentAutomatizmAfterTreeActivatedID = oldAtmzAfterTreeActivatedID
		oldAtmzAfterTreeActivatedID = 0
	}
	wasCurrentAutomatizmAfterTree = 1 // для активации функции consciousnessElementary: 1 этот автоматизм из текущей активации дерева, есть currentAutomatizmAfterTreeActivatedID

	// всегда сначала активировать дерево понимания, результаты которого могут заблокировать все внизу
	//if !wasRunUnderstandingSituation && EvolushnStage > 3 {
	if EvolushnStage > 3 {

		// Активировать Дерево ситуации: или запустить ментальный автоматизм или - ориентировочная реакция для осмысления
		understandingSituation(1) //
		/*в func consciousnessElementary() был запущен моторный автоматизм, нужнозаблокировать все рефлексы и штатные автоматизмы,
		  может быть изменен штатный авмтоматизм и уже запущен другой.
		*/
		if MentalReasonBlocing && curFunc13ID == 0 {
			return true
		}
	}
	//////////////////////
	//более примитивное реагирование, EvolushnStage < 4
	if EvolushnStage < 4 {
		if currentAutomatizmAfterTreeActivatedID > 0 { //ориентировочный рефлекс 2
			// EvolushnStage < 4  проверить подходит ли автоматизм к текущим условиям, если нет, - режим нахождения альтернативы  - ориентировочный рефлекс 2
			atzm := orientation(currentAutomatizmAfterTreeActivatedID)
			// если автоматизм в orientation прошел проверку, то он сразу запущен
			if atzm > 0 { // блокировка рефлексов, если automatizmID > 0
				return true
			}
		} else {
			// автоматизма нет у недоделанной ветки
			atzm := orientation(0)
			if atzm > 0 { // блокировка рефлексов, если automatizmID > 0
				return true
			} else { // нет реакции
				if EvolushnStage < 4 {
					// это означает блокировку рефлексов, если его клон-автоматизм заблокирован. Так происходит простейшая коррекция наследственной модели поведения на уровне Да/Нет.
					// более тонкая настройка автоматизма и его разблокировка возможна на стадиях > 3
					lib.SentConfusion("Не смог сориентироваться.")
					return true
				}
			}
		}
	} //if EvolushnStage < 4 {

	if EvolushnStage > 3 {
		//- просто запустить штатный автоматизм, раз нет блокировки сознанием !MentalReasonBlocing
		if currentAutomatizmAfterTreeActivatedID > 0 {
			if curFunc13ID > 0 {
				/*  endBaseIdCycle(curFunc13ID) // подвисает
				clinerFunctionsInAllCickles(13)
				IsEndWaitPeriodFunc13 = true */
				curFunc13ID = 0 // закрываем маркер цикла func13
			}
			RumAutomatizmID(currentAutomatizmAfterTreeActivatedID)
		}
	}

	if EvolushnStage > 4 {
		// отзеркаливание авторитарных действий - как решение доминант по аналогии
		if existsBaseContext(2) || existsBaseContext(3) { // текущая эмоция - поиск или игра
			if lastAutomatizmRun != nil && lastAutomatizmRunPulsCount > PulsCount-20 {
				// оператор ответил на действия Beast не позже, чем за 20 сек
				checkRelevantAction(lastAutomatizmRun.ActionsImageID, curActions.ID, 2) // примем авторитарный эффект ==2
			}
		}
	}

	// не блокировать то, что не относится к текущей активации условий деревьями
	MentalReasonBlocing = false //true - заблокировать все низкоуровневое

	return false
}

//////////////////////////////////////////////////////////

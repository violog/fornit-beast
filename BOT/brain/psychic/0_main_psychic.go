/*
	Общий файл психики

Психика активируется PsychicInit() после активации всех предшествующих структур.

По каждому событию с Пульта или изменению состояния организма активируется дерево автоматизмов.
Если есть автоматизм - перед его выполнением вызывается ориентировочный рефлекс привлечеия внимания orientation_2()
если нет автоматизма - ориентировчный рефлекс оценки ситуации orientation_1()
Привлечение осознанного внимания выявляет конечную цель - найти автоматизм или ничего не делать.

На уровне наивности (нет уверенно решающих проблему автоматизмов)
на первый план выходит отзеркаливание чужих действий и случайные пробы-ошибки.

Понимание смысла воспринимаемого и своих действий начинается с предопределенных
генетические целей действий Beast ID гомео-параметров, которые призвано улучшить данное действие - по его ID:
TerminalActionsTargetsFromID - это – наследственно заданная цель действия, не осознаваемая при его совершении.
Но с опытом каждому действию в конкретных условиях (и к ним добавляются слова и фразы)
будет ассоциироваться смысл (осознаваемая значимость).
Таким образом, отзеркаливая чужие "зеркальные" действия и совершая свои с оценкой результата,
будет пополняться МОДЕЛЬ ПОНИМАНИЯ при данных условиях.
Эта модель, фактически, будет составлена из наборов автоматизмов, привязанных к активной ветке дерева автоматизмов,
из которых один - текущий актуальный, остальные отбракованные и предположительные.
Автоматизмы с внутренними, ментальными действиями будут обеспечивать произвольность.

Это – наследственно заданная цель действия, не осознаваемая при его совершении.
Но с опытом каждому действию в конкретных условиях (и к ним добавляются слова и фразы)
будет ассоциироваться смысл (осознаваемая значимость).

Безусловные рефлексы психики прописываются в виде функций обработки
текущей инфрмационной среды CurrentInformationEnvironment.
В этой среде активируются текущие проблемы и доминанты нерешенной проблемы.

НЕ ЗАБЫВАТЬ для всех функций произвольной активации (по актуализации текущего самоощущения)
ставить блокировку по brain.NotAllowAnyActions - if brain.NotAllowAnyActions{ return }
*/
package psychic

import (
	"BOT/lib"
	"strconv"
)

///////////////////////////////

// true - ПРИ ТЕСТИРОВАНИИ СОХРАНЯТЬ В ФАЙЛАХ ВСЕ ЭЛЕМЕНТЫ
var doWritingFile = false // true  false

// переключатель игрового режима  psychic.IsGameMode
//var IsGameMode=false
// переключатель произвольной активации игрового режима
//var IsArbitraryGameMode = false

///////////////////////////////////////////////////////////////////////

// инициализирующий блок - в порядке последовательности инициализаций
// после condition_reflex.go
func PsychicInit() {

	if EvolushnStage < 2 { // еще нет психики
		return
	}
	automatizmTreeInit()
	loadActionsImageArr()
	automatizmInit()
	loadAmtzmNextString()
	emotionsInit()
	loadActivityInit()
	verbalInit()
	cerebellumReflexInit()

	loadEpisodicTree()
	loadEpisodicMentalTree()

	initCurrentInformationEnvironment()

	loadMentalCyckleEffectArr()

	loadSituationImage()
	loadThemeImageFromIdArr()
	//уже есть  loadPurposeImageFromIdArr()
	// уже есть loadProblemTree()
	UnderstandingTreeInit()
	ProblemTreeInit()

	loadProblemDominenta()
	// saveActionImageArr()// сохранить образы сочетаний ответных действий

	// просыпание - создание базового самоощущения CurrentInformationEnvironment
	wakingUp()

	//	SensorActivation(1,1,[]int{1})
	/*
		atmzm:=findAnySympleRandActions()
		if atmzm!=nil{	}
	*/

	//	FormingMirrorAutomatizmFromList("/mirror_reflexes_basic_phrases/1_2.txt")

	//	FormingMirrorAutomatizmFromTempList("/lib/mirror_basic_phrases_common.txt")

	/*	lib.NoReflexWithAutomatizm = true
		lib.ActionsForPultStr = "0|qqqqqq||1|rrrrrrrrrrrrrr||3|wwwwwwwwwwww"
		lib.ActionsForPultStr = lib.SharedReflexWithAutomatizm()
	*/
	return
}

/////////////////////////////////////////////////////////////

// ПУЛЬС психики
var PulsCount = 0 // передача тика Пульса из brine.go
var LifeTime = 0
var EvolushnStage = 0       // стадия развития
var IsSleeping = false      // сон без сновидений
var IsSleepingDream = false // фаза сновидений сна
var WakeUppingActivation = true

func PsychicCountPuls(evolushnStage int, lifeTime int, puls int, sleepingType int) {

	if evolushnStage < 2 { // недостаточная стадия развития
		return
	}

	LifeTime = lifeTime
	EvolushnStage = evolushnStage
	PulsCount = puls // передача номера тика из более низкоуровневого пакета
	if sleepingType > 0 {
		IsSleeping = true
		if sleepingType == 2 {
			IsSleepingDream = true
		}
	}

	// тики в automatizm_result.go для удобства
	orientarionPuls()
	automatizmActionsPuls()
	moodePulse()

	if lastAutomatizmRunPulsCount > 0 && (lastAutomatizmRunPulsCount+100 < PulsCount) {
		CurrentInformationEnvironment.IsIdleness100pulse = true
	}

	// в условии NoautomatizmAfterStimul>2 чтобы не повторялось lib.WritePultConsol
	if NoautomatizmAfterStimul > 2 && (NoautomatizmAfterStimul < PulsCount-2) && PulsCount > 5 {
		// уже 2 пульса как нет автоматизма в ответ на Стимул, значит нет периода ожидания
		NoautomatizmAfterStimul = 2 // - сигнал детектора отсуствия автоматизма на Стимул оператора
		lib.WritePultConsol("<span style='color:blue;background-color:#C5FFCC'>ПРАВИЛА. Уже 2 пульса как нет автоматизма в ответ на Стимул, значит нет периода ожидания (установлено: NoautomatizmAfterStimul=2)</span>")
	}

	if IsSleeping {
		sleepingProcess()
	}

	if IsSleeping {
		IsFirstActivation = true

	} else {
		if IsFirstActivation {
			ReadiStatus = 1 //готовность Beast Для пульта:
		}

		if evolushnStage > 3 && PulsCount == 4 {

		}

		// осознание при включении и бодрствовании - один раз
		if evolushnStage > 3 && PulsCount > 4 && IsFirstActivation {
			// начать мышление
			WakeUpping() // в understanding.go
			// первый запуск при пробуждении
			ReadiStatus = 2 ////готовность Beast Для пульта

			// запустить дерево автоматизмов в первый раз, иначе не будет информации для дерева проблем.
			automatizmTreeActivation()
			//!!	MentalReasonBlocing = consciousnessElementary() - будет запущен после automatizmTreeActivation()
			beginMentalCycle() // начать первый, главный цикл мышлления с вызовом func infoFunc8()
		}

		if evolushnStage > 3 && PulsCount > 5 {
			dispetchConsciousnessThinking() // постоянная циркуляция циклов мышления

			if isNeedForCommunication() { // нужно провоцировать оператора
				// провокация в infoFunc31(c)
				mentalInfoStruct.noOperatorStimul = true
			}
		}
	}

	return

	// просыпание - создание базового самоощущения CurrentInformationEnvironment
	//	if psychicPulsCount>3 {
	//		wakingUp()
	//	}

	//		CurrentInformationEnvironment.ActionsImageID=ActivityFromIdArr[1].ID// образ бездействия
	//		CurrentInformationEnvironment.LifeTime=LifeTime

}

/*
	готовность Beast Для пульта:

0 - Beast еще не пришел в себя, общение невозможно.
1 - Психика Beast активрована, но без осознания.
2 - Beast готов к общению.
*/
var ReadiStatus = 0

func GetPsichicReady() string {
	return strconv.Itoa(ReadiStatus)
}

// ////////////////////////////////////////////////////////////
var NotAllowAnyActions = false

func SetNotAllowAnyActions(notAllow bool) {
	NotAllowAnyActions = notAllow
}

///////////////////////////////////////

// просыпание - создание базового самоощущения CurrentInformationEnvironment
func wakingUp() {

	// осознание самоощущения
	SensorActivation(1)

	// очистить всякое при просырании
	usedActIdArr = nil
	UsedPraseIdArr = nil
}

/////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////
/*  активация по событиям с Пульта - из perception.go
Для блокировки активации дерева рефлексов вернуть true
*/
var firstStadiesWarning = true // защелка от повторов
// var pevPulsCount=0// не активировать в том же пульсе (возникает паразитная повторная активация)
var ActivationTypeSensor = 0 //передача типа акетивации в психику из рефлексов
func SensorActivation(activationType int) bool {
	if PulsCount < 4 { // || pevPulsCount == PulsCount
		return false
	}
	//pevPulsCount=PulsCount
	ActivationTypeSensor = activationType

	if EvolushnStage < 2 { // недостаточная стадия развития
		if firstStadiesWarning {
			firstStadiesWarning = false
			lib.WritePultConsol("Стадия развития " + strconv.Itoa(EvolushnStage) + " НЕДОСТАТОЧНА ДЛЯ АВТОМАТИЗМОВ")
		}
		return false
	}

	//  получение текущего состояния информационной среды: отражение Базового состояния и Активных Базовых контекстов
	//!!!! GetCurrentInformationEnvironment() только при ориентировочном рефлексе - смена самоощущения !!!

	atomatizmID := automatizmTreeActivation()
	if atomatizmID > 0 {

		return true
	}

	return false
}

/////////////////////////////////////////////////////////////////

////////////////////////////////////
/* Блокировать выполнение рефлексов на время ожидания результата автоматизма
вызывается из reflex_action.go рефлексов
*/
func NotAllowReflexesAction() bool {
	if MotorTerminalBlocking {
		return true
	}
	return false
}

////////////////////////////////////

// /////////////////////////////////
func SaveAllPsihicMemory() {
	notAllowScanInTreeThisTime = true
	SaveEmotionArr()
	SaveActivityFromIdArr()
	SaveVerbalFromIdArr()
	SaveAutomatizmTree()
	SaveAutomatizm()
	SaveAmtzmNextString()
	SaveSituationImage()
	SaveActionsImageArr()
	SaveCerebellumReflex()
	SaveUnderstandingTree()
	SaveProblemTree()

	SaveEpisodicTree()
	SaveEpisodicMentalTree()

	SaveMentalCyckleEffectArr()
	SavePurposeImageFromIdArr()

	SaveThemeImageFromIdArr()
	saveInterruptMemory()

	SaveProblemDominenta()

	notAllowScanInTreeThisTime = false
}

//////////////////////////////////////

// /////////////////////////////////////
func GetExtandIndoForPult() string {
	if detectedActiveLastNodID == 0 {
		return ""
	}

	return "Инфо: <span title='ID базового состояния 1 - Похо, 2 - Норма, 3 - Хорошо'>BaseID=<b>" + strconv.Itoa(CurrentCommonBadNormalWell) + "</b></span>," +
		"<span title='ID эмоции'>EmotionID=<b>" + strconv.Itoa(CurrentEmotionReception.ID) + "</b></span>, " +
		"<span title='ID дерева автоматизмов'>atmzmID=<b>" + strconv.Itoa(detectedActiveLastNodID) + "</b></span>, " +
		"<span title='ID дерева ситуации'>situatID=<b>" + strconv.Itoa(detectedActiveLastUnderstandingNodID) + "</b></span>, " +
		"<span title='ID дерева проблем'>problemID=<b>" + strconv.Itoa(detectedActiveLastProblemNodID) + "</b></span> " +
		// доп.инфа:
		//"<span title='XXXXX'>problemID=<b>"+XXXX+"</b></span> " +
		"<span style='color:#999999'>(Чтобы добавить любую нужную инфу вставь ее в 0_main_psychic.go func GetExtandIndoForPult )</span>"

} /////////////////////////////////////

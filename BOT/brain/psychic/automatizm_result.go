/* Ожидание результата запущенного автоматизма и его обработка

В BAD_detector.go в самом низу есть func BetterOrWorseNow() с комментариями по делу. Я ее отрабатывал как раз для того, чтобы фиксировать любые улучшения или ухудшения для определения эффекта автоматизма.
Она вызывается (через трансформатор против цицличности wasChangingMoodCondition()) 2 раза: в момент запуска автоматизма и как только совершится любое действие оператора на пульте. Таким образом в automatizm_result.go получается дифферент:
oldlastBetterOrWorse,oldBetterOrWorse,oldParIdSuccesArr = wasChangingMoodCondition()
Т.е. если ты поставишь точку прерывания на
oldlastBetterOrWorse,oldBetterOrWorse,oldParIdSuccesArr = wasChangingMoodCondition()
то и получишь эффект автоматизма.
*/

package psychic

import (
	"BOT/brain/gomeostas"
	"BOT/lib"
	"strconv"
)

//////////////////////////////////////////////////////////////////
/* Длина группового правила.
Накручивается при каждом стимуле в activeGameMode(), сбрасывается в automatizmActionsPuls() и StopWaitingWeriodFromOperator() - принудительно по пашке на пульте или по истечении периода ожидания.
то есть глубина группового правила зависит от времени удержания образа. Это позволяет записывать в группу только события, которые были в периоде удержания, а не все подряд.
В функциях GPT лимит устанавливается отдельно.

ПРАКТИЧЕСКИ НЕ ИСПОЛЗУЕТСЯ...
*/
//var LimitGroupRules int

////////////////////////////////////////////////////////////////

/*
	Это используется для определения момента реакция оператора Пульта на действия автоматизма.

За 20 сек г.параметры могли бы просто натечь и вызывать сработавание при ожидании ответной реакции.
Флаг сбрасывается через пульс после запуска автоматизма.
*/
var WasOperatorActiveted = false

var WasConditionsActiveted = false

var savePsyBaseMood = 0 // -1 Плохое настроение, 0 Нормальное, 1 - хорошее настроение
// для более точной оценки
var savePsyMood = 0 //сила Плохо -10 ... 0 ...+10 Хорошо
// НОВИЗНА СИТУАЦИИ сохраняется значение CurrentAutomatizTreeEnd[] для решений
var savedNoveltySituation []int

// отслеживание запущенных автоматизмов
// структура примитивных целей, создающих контекст ситуации НЕ СБРАСЫВАЕТСЯ после ожидания
var savePurposeGenetic *PurposeGenetic

/* При запуске автоматизма определяются:
// момент запуска автоматизма в числе пульсов
var LastRunAutomatizmPulsCount =0 //сбрасывать ожидание результата автоматизма если прошло 20 пульсов
// ожидается результат запущенного MotAutomatizm
var LastAutomatizmWeiting *Automatizm //сбрасывается указатель автоматизма
*/

func setAutomatizmRunning(am *Automatizm, ps *PurposeGenetic) {
	lib.WritePultConsol("<span style='color:blue'>Ожидание ответа оператора.</span>")

	// при срабатывании автоматизма - блокируются все рефлексторные действия
	//MotorTerminalBlocking=true
	notAllowReflexRuning = true //уже есть, но на всякий случай :)

	LastAutomatizmWeiting = am                            // уже есть, но для надежности :)
	LastDetectedActiveLastNodID = detectedActiveLastNodID // уже есть, но для надежности :)

	savePsyBaseMood = PsyBaseMood
	savePsyMood = PsyMood
	savedNoveltySituation = NoveltySituation
	if ps != nil {
		savePurposeGenetic = ps
	}
	WasOperatorActiveted = false // ждем ответа оператора
	// зафиксировать текущее состояние на момент срабатывания автоматизма
	//!!!! oldlastBetterOrWorse,oldParIdSuccesArr = wasChangingMoodCondition()
	gomeostas.BetterOrWorseInit()

	CurrentInformationEnvironment.IsWaitingPeriod = true
	notAllowReflexRuning = false
}

func clinerAutomatizmRunning() {
	//MotorTerminalBlocking=false
	notAllowReflexRuning = false

	LastAutomatizmWeiting = nil // func RumAutomatizm()
	lastRunVolutionAction = nil // func showVolutionAction

	LastRunAutomatizmPulsCount = 0
	WasOperatorActiveted = false
	onliOnceWasConditionsActiveted = false
	// !!!! НЕ СБРАСЫВАТЬ savePurposeGenetic=nil - он может определяться независимо от запуска автоматизма

	CurrentInformationEnvironment.IsWaitingPeriod = false

	wasRunProvocationFunc = false

	// только в func saveNewMentalEpisodic очищать!   clinerFuncSequence()
}

func StopWaitingWeriodFromOperator() {
	setInterruptionEpisosde() //вставить пустой кадр эпиз.памяти - прервать тему
	clinerAutomatizmRunning()
	/*
		if !IsArbitraryGameMode {
			transfer.IsPsychicGameMode = false
			LimitGroupRules = -1
		}	*/
}

// ////////////////////// ПУЛЬС
// ПУЛЬС срабатывания по каждому Пульсу здесь для удобства
var oldBetterOrWorse = 0     //- стали лучше или хуже: величина измнения от -10 через 0 до 10
var oldParIdSuccesArr []int  //стали лучше следующие г.параметры []int гоменостаза
var oldlastBetterOrWorse = 0 // насколько изменилось общее состояние, значение от  -10(максимально Плохо) через 0 до 10(максимально Хорошо)
func automatizmActionsPuls() {

	if LastRunAutomatizmPulsCount == 0 {
		return
	}
	// вышло время ожидания реакции
	if (LastRunAutomatizmPulsCount + WaitingPeriodForActionsVal) < PulsCount {
		/*
			if !IsArbitraryGameMode {
				transfer.IsPsychicGameMode = false
				LimitGroupRules = -1
			}*/
		// если просто позволить отработать noAutovatizmResult() при закрытии func13(), то принудительное обнуление времени ожидания в
		// main/getParams := r.FormValue("get_params") приведет к тому, что тут же сработает noAutovatizmResult() и автоматизм запустится 2 раза подряд
		// поэтому блокируем это дело - не позволяем двойную активацию после отработки func13()
		if IsEndWaitPeriodFunc13 {
			IsEndWaitPeriodFunc13 = false
		} else {
			// отреагировать на отсуствие реакции - повторить автоматизм с большей силой Energy
			// Из МОЗЖУЧКА как-то отреагировать на отсуствие реакции - повторить автоматизм с большей силой Energy
			if noAutovatizmResult() { // была попытка отреагировать сильнее - в cerebellum.go
				return // чтобы не сбрасывать clinerAutomatizmRunning()
			}
		}

		//сбрасывать ожидание результата автоматизма если прошло WaitingPeriodForActionsVal пульсов
		clinerAutomatizmRunning()
		//вставить пустой кадр эпиз.памяти - прервать тему
		setInterruptionEpisosde()

		if oldBetterOrWorse < 0 { // Плохо и оператор не ответил в течение периода ожидания на важный запрос
			// Новая тема мышления
			runNewTheme(7, 2)
		}
	}
}

// отреагировать на отсуствие реакции - повторить автоматизм с большей силой Energy при EvolushnStage <= 3
func noAutovatizmResult() bool {

	if EvolushnStage > 3 {
		// осмыслить ситуацию - Активировать Дерево Понимания
		understandingSituation(1)
		clinerAutomatizmRunning()
		return true
	}

	// не опасная ситуация, можно поэкспериментировать
	if EvolushnStage == 3 && !CurrentPurposeGenetic.veryActual {
		/* в случае отсуствия автоматизма в данных условиях - послать оператору те же стимулы, чтобы посмотреть его реакцию.
		   Создание автоматизма, повторяющего действия оператора в данных условиях
		НО если уже помылался provokatorMirrorAutomatizm то больше не делать этого (бесконечный цикл)
		*/
		if oldProvokatorAutomatizm != LastAutomatizmWeiting { // не повторять, если только что был такой ответ
			provokatorMirrorAutomatizm(LastAutomatizmWeiting, &CurrentPurposeGenetic)
			clinerAutomatizmRunning()
			return true
		}
	}

	// реакция была, но оператор не обратил на нее внимания, нужно усилить силу действия мозжечковым рефлексом
	if cerebellumCoordination(LastAutomatizmWeiting, 1) {
		// и тут же снова запустить реакцию!
		if oldProvokatorAutomatizm != LastAutomatizmWeiting { // не повторять, если только что был такой ответ
			setAutomatizmRunning(LastAutomatizmWeiting, &CurrentPurposeGenetic)
			clinerAutomatizmRunning()
			return true
		}
	}
	clinerAutomatizmRunning()
	return false
}

/*
сохраняет ID автоматизма до создания зеркального
актуально для 3 стадии, когда на стимул-ответ Оператора сработал существующий автоматизм
его надо отправить на пульт, а не созданный зеркальный
если же запретить вообще создавать зеркальные для таких ситуаций
тогда зеркалить можно будет только НОВЫЕ стимулы и ответы
что обедняет механизм третьей стадии
*/
var oldAtmzAfterTreeActivatedID = 0
var IsEndWaitPeriodFunc13 = false

/*
	ПОСЛЕ ОРИЕНТИРОВОЧНОГО РЕФЛЕКСА оценивать действие запущенного автоматизма

lastBetterOrWorse НЕ ИСПОЛЬЗУЕТСЯ т.к. lastCommonDiffValue более точен и информативен
*/
func calcAutomatizmResult(lastCommonDiffValue int, wellIDarr []int) {

	lib.WritePultConsol("<span style='color:blue;background-color:#FFD0FF;'>Был ОТВЕТ ОПЕРАТОРА (func calcAutomatizmResult). Мотивационный эффект: <b>" + strconv.Itoa(lastCommonDiffValue) + "</b></span>")

	/*	полезность автоматизма в третьей стадии при отзеркаливании на втором шаге ставится в createNewMirrorAutomatizm()
		при этом нельзя менять полезность автоматизма созданного в цикле отзеркаливания на первом шаге, иначе в getAutomatizmFromNodeID()
		будет выдавать как наилучший автоматизм первого шага. То есть для отзеркаливания в 3 стадии должно быть строго из за getAutomatizmFromNodeID():
		1. автоматизм, созаднный на 1 шаге - Usefulness==0
		2. автоматизм, созаднный на 2 шаге - Usefulness==1 (устанавливается currentAutomatizmAfterTreeActivatedIDв createNewMirrorAutomatizm())
		И то же самое для отработки infoFunc13()
	*/
	if EvolushnStage != 3 && curFunc13ID == 0 {
		if lastRunVolutionAction == nil { // было НЕ произвольное действие
			automatizmCorrection(LastAutomatizmWeiting, lastCommonDiffValue, wellIDarr)
		}
	}
	if lastCommonDiffValue < 0 {
		motorActionEffect = 0 // нужно думать, искать действия в func consciousnessThinking()
	}

	// в третьей стадии по умолчанию, в более высоких по факту срабатывания func13()
	// для этого в ней используется маркер curFunc13ID >0
	if EvolushnStage == 3 || curFunc13ID > 0 {
		if GetMotorsAutomatizmListFromTreeId(detectedActiveLastNodID) != nil { // если есть список всех автоматизмов для ID узла Дерева
			// сохраняем активировавшийся на стимул автоматизм
			// для случая, когда оператор дал ответ, как отвечать - а на ответ уже есть автоматизм. В этом случае бот ответит - и его ответ пойдет в запись зеркального автоматизма
			// вместо того, что указал оператор. Чтобы этого не было, надо сохранить ответ оператора.
			oldAtmzAfterTreeActivatedID = getAutomatizmFromNodeID(detectedActiveLastNodID)
		}

		/* отзеркаливание ответа оператора не зависимо от того, стало хуже или лучше
		потому, что это был ответ оператора на действия автоматизма, значит - авторитетный ответ
		Создание автоматизма, повторяющего действия оператора в данных условиях
		*/
		createNewMirrorAutomatizm(LastAutomatizmWeiting)
		/*		if curFunc13ID > 0 {
				//clinerAutomatizmRunning() // нельзя, иначе fixEpizMemoryRules() не отработает, только fixEpizMemoryTeachRules()
				endBaseIdCycle(curFunc13ID) // подвисает
				clinerFunctionsInAllCickles(13)
				IsEndWaitPeriodFunc13 = true
				curFunc13ID = 0 // закрываем маркер цикла func13
			}*/
	}

	// >3 потому, что раньше не пишется эпизодическая память и формируются более примитивные механизмы.
	if EvolushnStage > 3 {
		if !IsSleeping { //бодрствование (при сновидении или грез, мечтаний, пассивного мышления - fixDreamsRules(lastCommonDiffValue))
			/* При каждом ответе на действия оператора - прописывать текущее правило rules
			   		и делать новый кадр эпизодической памяти
			      А так же просматривать эпизод память взад макчимум на EpisodeMemoryPause шагов или до паузы в общении > 30 шагов,
			   		фиксируя цепочку правил.
			*/

			// был запуск автоматизма или произвольной цепочки действий - записать 2 вида Правил
			if lastRunVolutionAction != nil { // было произвольное действие
				// создать автоматизм из lastRunVolutionAction
				_, azm := createAutomatizmFromNextString(lastRunVolutionAction.next, LastDetectedActiveLastNodID)
				automatizmCorrection(azm, lastCommonDiffValue, wellIDarr)
			}
			// записать 2 вида Правил для автоматизма или произвольной цепочки действий
			//		stimul, _ := CreateNewlastActionsImageID(0, curActiveActions.ActID, curActiveActions.PhraseID, curActiveActions.ToneID, curActiveActions.MoodID,true)
			// образ действий оператора Стимул:	curStimulImageID и есть Ответ: curActiveActions
			fixEpizMemoryRules(lastCommonDiffValue)

			// записать учительское Правило: как Оператор отвечает на действиЯ Beast - авторитетное Правило всегда имеет +эффект
			if !wasRunProvocationFunc { // не записывать учительское правило после провокации func infoFunc31
				// не включаем в условие curActiveActionsID > 0 т.к. ID стимула будет определен в func fixEpizMemoryTeachRules
				if ActivationTypeSensor > 1 { // только стимул оператора, а не гомео
					//если оператор ответил только во время периода ожидания
					if (LastRunAutomatizmPulsCount + WaitingPeriodForActionsVal) > PulsCount {
						fixEpizMemoryTeachRules(lastCommonDiffValue)
					}
				}
			}
			wasRunProvocationFunc = false

			if EvolushnStage > 4 && lastCommonDiffValue > 0 {
				// закрыть доминанту, котора точно подходит
				checkGestalt(lastCommonDiffValue)
			}
		} //if !IsSleeping || (IsSleeping && IsDreamsProcess){//бодрствование или процесс сновидения
	}
	return
}

/*
	для индикации период ожидания реакции оператора на действие автоматизма

//   psychicWaitingPeriodForActions()
Индикация включается после появления диалога ответа на Пульте (pult_gomeo.php: var allowShowWaightStr=0;).
*/
func WaitingPeriodForActions() (bool, int) {
	//if LastRunAutomatizmPulsCount > 0 && ActivationTypeSensor > 1 {
	if LastRunAutomatizmPulsCount > 0 {
		time := WaitingPeriodForActionsVal - (PulsCount - LastRunAutomatizmPulsCount)
		return true, time
	}
	return false, 0
}

///////////////////////////////////////////////

/*
	Определение эффекта реакции по результатам в BAD_detector.go, негативный или позитивный: возвращает значение res0.

# В  gomeostas.BetterOrWorseNow() учитывается CommonMoodAfterAction - Общее (де)мотивирующее действие с Пульта

res - стали лучше или хуже: величина измнения от -10 через 0 до 10
wellIDarr - стали лучше следующие г.параметры []int гоменостаза

wasChangingMoodCondition вызывается 1 раз: как только совершится любое действие оператора на пульте (kind==2)

Информация о том, когда вызывается (kind) раньше использовалась, но теперь она пока не используется.
*/
var CurrentMoodCondition = 0 // индикация на Пульте
func wasChangingMoodCondition() (int, []int) {
	//стало хуже или лучше теперь, возвращает величину измнения от -10 через 9 до 10
	res0, wellIDarr := gomeostas.BetterOrWorseNow()
	prevStimulsEffect = globalStimulsEffect
	globalStimulsEffect = res0 // глобализуем для func saveNewEpisodic

	// передать Боль и Радость в психику
	painValue, joyValue = gomeostas.GetCurPainJoy()
	//Скорректировать настроение по результату реакции??
	// GetCurMood() наверное здесь не нужно это делать

	// если в текущем стимуле с пульта не было действия с эффектом, нужно игнорировать эффект от предыдущего действия, иначе оно передается на действие с нулевым эффектом через CommonMoodAfterAction
	// например нажали кнопку наказать, срабатал автоматизм, потом послали вербальный стимул без кнопок - и он заблокировал предыдущий автоматизм потому, что при нулевом эффекте
	// вербального стимула эффект взялся от CommonMoodAfterAction: BetterOrWorseNow() стр. 407
	if curActions.ActID == nil {
		res0 = 0
	}

	/* это надо делать через LastAutomatipmCorrection(), потому что здесь CheckCurActions() показывает,
	   что РАНЕЕ была нажата какая то кнопка и не факт, что только что, а может и несколько шагов диалога назад.
	   В итоге может получиться ситуация:
	   "привет" : [наказать] : "Хай" - и блокируются первые 2 автоматизма. Первый понятно, что эффектом кнопки, а сама кнопка потому,
	   что эффект ее нажатия сохранился в ActionFromPult[]
	   и когда пришел пускатель "Хай", здесь прочитался отрицательный эффект от кнопки и заблокировал ее саму.
	   То есть получилось, как будто вербальный автоматизм имеет плохой эффект и потому заблокировал предыдущий автоматизм.
	   поэтому в afterTreeActivation() поставлена заглушка if !action_sensor.IsPress3or4button,
	   а обработка учительских кнопок делается через LastAutomatipmCorrection()
	*/

	if curActiveActions != nil {
		// влияние значимости первого компонента фразы + влияние тона
		tone := curActiveActions.ToneID
		k := 1
		if tone == 4 { // повышенный
			k = 2
		}

		eObj, v := getObjectsImportanceValue(curActiveActions.ID, detectedActiveLastProblemNodID)
		if eObj != nil {
			if v > 0 {
				res0 = res0 + 1*k
			}
			if v < 0 {
				res0 = res0 - 1*k
			}
		}

	} //	if curActiveActions !=nil {

	// для страницы Пульта
	if res0 > 0 {
		CurrentMoodCondition = 3
	} //лучше
	if res0 == 0 {
		CurrentMoodCondition = 2
	} // не изменилось
	if res0 < 0 {
		CurrentMoodCondition = 1
	} // Хуже

	// для лога
	lib.WritePultConsol("Изменение состояния: " + strconv.Itoa(res0))

	return res0, wellIDarr
}

//////////////////////////////////////////////////////////////

/* // текущий ID пускового стимула типов curActiveActions или curBaseStateImage
при активации дерева автоматизмов. Если тип curBaseStateImage, то ID отрицательное (ID<0)!
*/
//var currentTriggerID=0 не нужен
//var currentRulesID = 0

// для установки значимости учительских кнопок -5/+5, которые обрабатываются отдельно от прочих. Иначе установится -1/+1
//var effectPress3or4button = 0

/*  !!!Нельзя делать только для кнопок такую ф-цию! кнопки могуи активироваться и с фразой! появилось множество побочек!

только при активации действием кнопок 	в игровом режиме при воздействии кнопками Наказать, Поощрить - записывать Usefulness ранее выполненного автоматизма
Вызывается только из perception.go
Не забыть потом в afterTreeActivation() поставить заглушку для calcAutomatizmResult() - иначе будет двойной вызов calcAutomatizmResult()
wasChangingMoodCondition() не подходит для

func LastAutomatipmCorrection() {
	if LastAutomatizmWeiting == nil {
		return
	}

	// LastRunAutomatizmPulsCount - время начала периода ожидания
	if (PulsCount - LastRunAutomatizmPulsCount) > (limitOfActionsAfterStimul + 10) { // прошло более limitOfActionsAfterStimul секунд
		return
	}

	// какие акции действуют в данный момент действий с пульта - активные контексты действий с Пульта Очищаются
	ActID := action_sensor.CheckCurActionsContext()

	ActivationTypeSensor = 2 // было действие

	var effect = 0
	if lib.ExistsValInArr(ActID, 3) { // наказать
		effect = -5
	}
	if lib.ExistsValInArr(ActID, 4) { // поощрить
		effect = 5
	}
	if effect != 0 {
		var gomeoParIdSuccesArr []int
		// так как обработка действий учительских кнопок происходит ДО automatizmTreeActivation() нужно здесь задать образы активностей,
		// иначе правило не запишется. См fixEpizMemoryRules() стр. 36
		curStimulImage = curActiveActions
		curStimulImageID = curActiveActionsID
		// обработать изменение состояния фиксация ПРАВИЛА, Стимул - ОТ ОПЕРАТОРА
		calcAutomatizmResult(effect, gomeoParIdSuccesArr)
		effectPress3or4button = effect
	}
}
*/
///////////////////////////////////////////////////////////////////////////

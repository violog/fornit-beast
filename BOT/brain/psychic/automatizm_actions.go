/*  Моторные дейсвтвия автоматизма

Для каждого действия brain\reflexes\terminete_action.go задается "сила" действия в градации от 1 до 10, которая передается наПульт словами:
Максимально (сила=10), wwww (сила=8)", "Очень сильно (сила=9), ... Едва (сила=1).
При этом пропорционально расходуется энергия и могут происходить другие изменения гоместаза.
Такой результат сопоставляется с допустимым сразу при действии и корректируется установкой рефлекса мозжечка.

Две области моторного терминала уровня психики:
Область Брока VerbalFromIdArr=make(map[int]*Verbal)
отвечает за смысл распознанных слов и словосочетаний,
за конструирование собственных словосочетаний,
за моторное использование сло и словосочетаний.
За все ответственная структура - образ осмысленных слов и сочетаний.

Область моторных действий ActivityFromIdArr=make(map[int]*Activity)
отвечает за смысл распознанных действий с Пульта,
за конструирование собственных последовательностей действий,
за моторное использование действий.
За все ответственная структура - образ осмысленных действий и их сочетаний.

*/

package psychic

import (
	"BOT/brain/gomeostas"
	_ "BOT/brain/gomeostas"
	termineteAction "BOT/brain/terminete_action"
	"BOT/brain/transfer"
	word_sensor "BOT/brain/words_sensor"
	"BOT/lib"
	"strconv"
)

// блокировка действий во сне и при совершаемых действиях
var MotorTerminalBlocking = false

// ссылка на запущенный автоматизм перекрывается следующим запуском, т.е. всегда есть инфа о последнем запущенном
var lastAutomatizmRun *Automatizm
var lastAutomatizmRunPulsCount = 0

/*
НАЧАЛО ПЕРИОДА ОЖИДАНИЯ ОТВЕТА с Пульта
момент запуска автоматизма в числе пульсов -
только если LastAutomatizmWeiting был в ответ на действия Оператора!
Cбрасывать ожидание результата автоматизма если прошло WaitingPeriodForActionsVal (60) пульсов
*/
var LastRunAutomatizmPulsCount = 0 //
// период ожидания реакции оператора на действие автоматизма
const WaitingPeriodForActionsVal = 60

// ожидается результат запущенного MotAutomatizm
var LastAutomatizmWeiting *Automatizm

// предыдущий запущенный MotAutomatizm
var prevLastAutomatizmWeiting *Automatizm

// активный узел дерева в момент запуска автоматизма
var LastDetectedActiveLastNodID = 0

/*
	предыдущий момент запуска автоматизма

теперь prevLastDetectedActiveLastPulsCount==lastAutomatizmRunPulsCount НО НЕ СБРАСЫВАЕТСЯ, а есть всегда
*/
var prevLastDetectedActiveLastPulsCount = 0

/*
	запуск автоматизма на выполнение

возвращает true при успехе
*/
func RumAutomatizmID(id int) bool {

	//	a:= AutomatizmFromId[id]
	a, ok := ReadeAutomatizmFromId(id)
	if !ok {
		return false
	}
	return RumAutomatizm(a)
}

/*
	запрос из рефлексов, можно ли выполнять рефлекс if !psychic.getAllowReflexRuning(){ return }

РЕфлексы разблокируются через
*/
var notAllowReflexRuning = false

func GetAllowReflexRuning() bool {
	if notAllowReflexRuning || MotorTerminalBlocking {
		return false
	}
	return true
}

/*
	Запуск автоматизма с передачей строки на Пульт c

"10|"+message

	isTeachQuestion=true - Показать непонимание, растерянность с предложение научить

в случае отсуствия пси-реакций с вопросом о том, как нужно реагировать.

При запуске вторичного автоматизма Next не устанавливаются lastAutomatizmRun и т.п. переменные,
т.о. в периоде ожидания Эффект реакции применяется только к первому, пусковому автоматизму цепочки.
*/
var isTeachQuestion = false
var isInterruptAutmtzm = false

func RumAutomatizm(am *Automatizm) bool {
	//var isActiveLastRunAutomatizmPulsCount bool // маркер срабатывания фиксации активность мот. автоматизма
	isInterruptAutmtzm = false

	if am == nil {
		isInterruptAutmtzm = true
		return false
	}
	if MotorTerminalBlocking { //блокировка моторных терминалов во сне или произвольно
		isInterruptAutmtzm = true
		return false
	}

	//if wasRunTreeStandardAutomatizm { // уже был запущен штатный автоматизм после Стимула. ТЕПЕРЬ всегда после запуска автоматизма LastRunAutomatizmPulsCount >0
	if LastRunAutomatizmPulsCount > 0 {
		isInterruptAutmtzm = true
		return false
	}
	if wasRunPurposeActionFunc { // если ранее был запущен ментально в infoFunc17
		isInterruptAutmtzm = true
		return false
	}

	// NotAllowAnyActions ставится тогда, когда сохранение памяти должно выполняться в тишине, в бездействии
	if NotAllowAnyActions {
		isInterruptAutmtzm = true
		return false
	}
	if am.ActionsImageID == 0 {
		isInterruptAutmtzm = true
		return false
	}

	// блокировка выполнения плохого автоматизма, если только не применена "СИЛА ВОЛИ"
	if am.Usefulness < 0 && am.BranchID > 0 {
		isInterruptAutmtzm = true
		lib.WritePultConsol("Блокировка выполнения плохого (Usefulness=" + strconv.Itoa(am.Usefulness) + ") автоматизма iD=" + strconv.Itoa(am.ID))
		return false
	}

	// описание автоматизма
	res := GetAutomotizmActionsString(am, true, false) // здесь пишется "Энергичность"
	if len(res) == 0 {                                 // может не быть акции, а фраза не распознана, тогда ничего нет
		// такой автоматизм нужно удалить
		deleteAutomatizm(am)
		isInterruptAutmtzm = true
		return false
	}

	//	if am.BranchID > 0 {  любоый запуск автоматизма, даже пробный, без привязки к узлу
	notAllowReflexRuning = true // блокировка рефлексов
	// для учительского правила
	if LastAutomatizmWeiting == nil {
		prevLastAutomatizmWeiting = am
	} else {
		prevLastAutomatizmWeiting = LastAutomatizmWeiting
	}
	LastAutomatizmWeiting = am
	lastAutomatizmRun = am
	/* 	if lastAutomatizmRunPulsCount == 0 {
		prevLastDetectedActiveLastPulsCount = PulsCount
	} else {
		prevLastDetectedActiveLastPulsCount = lastAutomatizmRunPulsCount
	}		 */
	// всегда показывать время запуска автоматизма, в принципе теперь prevLastDetectedActiveLastPulsCount==lastAutomatizmRunPulsCount
	prevLastDetectedActiveLastPulsCount = PulsCount

	lastAutomatizmRunPulsCount = PulsCount
	//	}

	if NoautomatizmAfterStimul > 0 { // чтобы ставить прерывание
		// ПО-ЛЮБОМУ УБИРАЕМ МЕТКУ - ДЕТЕКТОР
		NoautomatizmAfterStimul = 0 // автоматизм выполнен, обнуляем метку
	}

	stimulCount = 0 //сколько раз был стимул от оператора после последнего запуска Ответа

	var out = "3|" // Бессознательный Автоматизм

	ta := ""
	switch levelOfRunAutomatizm {
	case 0:
		ta = "Штатный"
	case 1:
		ta = "1 уровень"
	case 2:
		ta = "2 уровень"
	case 3:
		ta = "Мышление"
	}
	out += "<div style=\"position:absolute;top:-10px;right:0;color:gray;\">" + ta + "</div>"
	//out += "<span style=\"position:relative;top:-10px;\">WWWWWWWW</span>"
	levelOfRunAutomatizm = 0

	// записать в инфо-окружение - память о происходящем сохраняется в массиве InformationEnvironmentObjects []*InformationEnvironment
	if am.BranchID > 0 {
		CurrentInformationEnvironment.AnswerImageID = lastAutomatizmRun.ActionsImageID

		//if isActiveLastRunAutomatizmPulsCount == true {
		if len(res) < 2 { // пустой и без запуска LastRunAutomatizmPulsCount не пускать
			return false
		}
		if isTeachQuestion {
			out = "10|Ответь сам на &quot;<b>" + res + "</b>&quot; чтобы показать, как лучше ответить."
			curActiveActionsPulsCount = 0
		} else {
			out += res
		}
		/*		} else {
				// не сработало isActiveLastRunAutomatizmPulsCount в процессе выполнения func13
				if isTeachQuestion {
					passivationAutomatizm(am, -1)
					return false
				}
				out += res
			}*/
	} else {
		out += res
	}

	/* на стадиях меньше 4 не активируем флаг потому, что он так и будет висеть в true и все блокировать. False ставится в consciousnessElementary(), которая активна только с 4 стадии
	if EvolushnStage > 3 {
		wasRunTreeStandardAutomatizm = true //Был запущен штатный автоматизм
	ТЕПЕРЬ всегда после запуска автоматизма LastRunAutomatizmPulsCount >0
	}*/

	lib.SentActionsForPult(out)
	isTeachQuestion = false //надо скидывать при любом раскладе - прошел попугайский автоматизм или нет

	// отслеживать последствия в automatizm_result.go при любом срабатывании автоматизма
	/* (было: начать ПЕРИОД ОЖИДАНИЯ реакции оператора - только при Стимуле Оператора, а не изменением состояния)
	Период ожидания должен начинаться ЗАНОВО после каждого ответа Beast и независимо от того, привязан ли автоматизм к ветке.
	*/
	//	if ActivationTypeSensor > 1 && am.BranchID > 0 {
	/*От Стимула curActiveActionsPulsCount до Ответа PulsCount должно быть было не более 3-х пульсов
	  		(потому как бот не может думать так долго
	  		!!!!С ЦИКЛАМИ МОЖЕТ ВЫДАТЬ ПРОИЗВОЛЬНЫЙ ОТВЕТ типа AmtzmNextString хоть через сколько времени!!!!),
	  		чтобы начался период ожидания,
	  		иначе ответ был явно не на Стимул, а, м.б. - по инициативе Beast (по ответу без стимула Правило не пишется).
	  Даже если возник в ходе решения доминанты,то он сразу не выдается на Пульт, а записывается в правило - уже более высокого порядка.
	*/
	// свежесть Стимула оператора - не позже, чем limitOfActionsAfterStimul пульса до Ответа на него
	//	if curActiveActionsID > 0 && (curActiveActionsPulsCount > (PulsCount - limitOfActionsAfterStimul)) {
	LastRunAutomatizmPulsCount = PulsCount // активность мот.автоматизма в чисде пульсов
	detectedActiveLastNodPrevID = detectedActiveLastNodID
	detectedActiveLastUnderstandingNodPrevID = detectedActiveLastUnderstandingNodID
	// может быть больше 3 пульсов - с этим надо что то делать
	// причина - циклы запускающие функции
	//isActiveLastRunAutomatizmPulsCount = true

	setAutomatizmRunning(am, curPurposeGeneticAutmtzm)
	//	}
	//	}

	/* при отзеркаливании на 2 и 4 стадиях нельзя делать попугайские с Usefulnes==-1 - иначе они могут разблокироваться и стать штатными в checkForUnbolokingAutomatizm()!!!
	кроме того, если заминусовать Usefulnes и при этом не отработал вопрос-ответ типа [out="10|Ответь сам на...] то получится заминусованный попугайский автоматизм
	который потом не даст ничего привязать к узлу, придется меняя пусковой стимул привязывать автоматизм к другому
	при создании автоматизма Usefulness и Belief ==0
	*/
	/*		if isTeachQuestion{
			isTeachQuestion = false
			am.Usefulness=0
			SetAutomatizmBelief(am,0)
		}*/
	if am.BranchID > 0 {
		//выполнить мозжечковый рефлекс сразу после выполняющегося автоматизма
		if !transfer.IsPsychicGameMode { // в игровом режиме нельзя корректировать действия - здесь только запись нового
			runCerebellumAdditionalAutomatizm(0, am.ID)
		}
		LastDetectedActiveLastNodID = detectedActiveLastNodID
		/* Блокировать выполнение рефлексов на время ожидания результата автоматизма
		вызывается из reflex_action.go рефлексов
		*/
		//isReflexesActionBloking=true // отмена в automatizm_result.go или просто isReflexesActionBloking=false

		return true
	}
	return false
}

func GetAutomotizmActionsString(am *Automatizm, writeLog bool, infoPult bool) string {
	var out = ""

	//ai:=ActionsImageArr[am.ActionsImageID]
	ai, ok := ReadeActionsImageArr(am.ActionsImageID)
	if !ok {
		lib.WritePultConsol("Нет карты ActionsImageArr для образа действий iD=" + strconv.Itoa(am.ActionsImageID))
		return ""
	}

	actAtr := ""
	praseStr := ""

	if ai.ActID != nil && am.BranchID > 0 {
		// учесть рефлекс мозжечка
		var addE int
		// для инфо-окна по щелчку таблицы автоматизмов на пульте не нужно накручивать энергию - надо просто показать текущую
		if !infoPult {
			addE = getCerebellumReflexAddEnergy(0, am.ID)
		}
		sumEnergy := am.Energy + addE
		if sumEnergy > 10 {
			sumEnergy = 10
		}
		if sumEnergy < 1 {
			sumEnergy = 1
		}
		//am.Count++
		actAtr = TerminateMotorAutomatizmActions(am.ID, ai.ActID, sumEnergy, infoPult)
	}

	if ai.PhraseID != nil && am.BranchID > 0 {
		addE := getCerebellumReflexAddEnergy(0, am.ID)
		praseStr = TerminatePraseAutomatizmActions(ai.PhraseID, am.Energy+addE)
	}
	// может не быть акции, а фраза не распознана, тогда ничего нет
	if len(actAtr) == 0 && len(praseStr) == 0 {
		return ""
	}
	out += actAtr + praseStr

	if ai.ToneID != 0 {
		out += "<br>" + getToneStrFromID(ai.ToneID) + "<br>"
	}

	if ai.MoodID != 0 {
		out += "<br>" + getMoodStrFromID(ai.MoodID) + "<br>"
	}
	if writeLog && am.BranchID > 0 {

		lib.WritePultConsol("<span style='color:blue;background-color:#FFFFA3;'>Запускается АВТОМАТИЗМ ID=" + strconv.Itoa(am.ID) + " с действием ID = " + strconv.Itoa(am.ActionsImageID) + " " + out + "</span>: ")

	}
	if am.NextID > 0 {
		out += showNextAtmtzmAction(am.ID, am.NextID, am.Energy)
	}

	// сбросить метку Не было действий более 100 пульсов
	CurrentInformationEnvironment.IsIdleness100pulse = false

	return out
}

// для функций пульта
func GetAutomotizmIDString(id int) string {

	//	am:= AutomatizmFromId[id]
	am, ok := ReadeAutomatizmFromId(id)
	if !ok {
		return "Нет автоматизма с ID = " + strconv.Itoa(id)
	}
	var out = ""
	//	ai:=ActionsImageArr[am.ActionsImageID]
	ai, ok := ReadeActionsImageArr(am.ActionsImageID)
	if ok {
		// учесть рефлекс мозжечка
		addE := getCerebellumReflexAddEnergy(0, am.ID)
		sumEnergy := am.Energy + addE
		if sumEnergy > 10 {
			sumEnergy = 10
		}
		if sumEnergy < 1 {
			sumEnergy = 1
		}
		//am.Count++
		out += TerminateMotorAutomatizmActions(am.ID, ai.ActID, sumEnergy, true)
	}

	if ai.PhraseID != nil {
		addE := getCerebellumReflexAddEnergy(0, am.ID)
		out += TerminatePraseAutomatizmActions(ai.PhraseID, am.Energy+addE)
	}
	if len(out) == 0 {
		return "ПУСТО..."
	}

	if ai.ToneID != 0 {
		out += "<br>" + getToneStrFromID(ai.ToneID) + "<br>"
	}

	if ai.MoodID != 0 {
		out += "<br>" + getMoodStrFromID(ai.MoodID) + "<br>"
	}

	return out
}

/*
	совершить МОТОРНОЕ (http://go/pages/terminal_actions.php) действие  - Dnn-часть автоматизма (не фраза)

cила действия сначала задается =5, а потот корректируется мозжечковыми рефлексами
Использование: 	TerminateMotorAutomatizmActions(actIDarr,energy)
*/
var rumAutomatizmOldID = 0 //
var rumAutomatizmOldEnergi = 0

func TerminateMotorAutomatizmActions(amID int, actIDarr []int, energy int, infoPult bool) string {
	// energy=1
	var out = ""
	var isAct = false
	for i := 0; i < len(actIDarr); i++ {
		if len(out) > 0 {
			out += ", "
		}
		// при моторном действии  меняются гомео-параметры:
		expensesGomeostatParametersAfterAction(actIDarr[i], energy)
		// выдать на Пульт:
		actName := termineteAction.TerminalActonsNameFromID[actIDarr[i]]
		// ЭНЕРГИЧНОСТЬ
		switch energy {
		case 1:
			out += "<span style=\"font-size:10px;\">" + actName + "</span>"
		case 2:
			out += "<span style=\"font-size:11px;\">" + actName + "</span>"
		case 3:
			out += "<span style=\"font-size:12px;\">" + actName + "</span>"
		case 4:
			out += "<span style=\"font-size:13px;\">" + actName + "</span>"
		case 5:
			out += "<span style=\"font-size:14px;\">" + actName + "</span>"
		case 6:
			out += "<span style=\"font-size:14px;\"><b>" + actName + "<b></span>"
		case 7:
			out += "<span style=\"font-size:17px;color:#927ACC\"><b>" + actName + "<b></span>"
		case 8:
			out += "<span style=\"font-size:19px;color:#E8A7A7\"><b>" + actName + "<b></span>"
		case 9:
			out += "<span style=\"font-size:21px;color:#E86966\"><b>" + actName + "<b></span>"
		case 10:
			out += "<span style=\"font-size:25px;color:#FF0000\"><b>" + actName + "<b></span>"
		}
		isAct = true
	}
	if isAct {
		/*Если втоматизм повторяется при одном и том же стимуле,
		то чисто "рефлекторно" повышать его силу действия с каждым разом, без мозжечкового механизма
		чтобы потом в одиночном вызове он не срабатывал.
		*/
		if rumAutomatizmOldID == amID { // повторился автоматизм
			// в игровом режиме это не учитываем, там постоянно используются одни и теже учительские кнопки при обучении
			// так же при показе инфо-окон по щелчку по строке таблицы автоматизмов на пульте просто показываем текущий уровень энергии
			if !transfer.IsPsychicGameMode && !infoPult {
				if energy+rumAutomatizmOldEnergi < len(termineteAction.EnergyDescrib) { // не превышать максимум
					rumAutomatizmOldEnergi++
				}
				energy += rumAutomatizmOldEnergi
			}
		} else {
			rumAutomatizmOldEnergi = 0
		}
		rumAutomatizmOldID = amID

		// название силы:
		var enegrName = ""
		if energy < len(termineteAction.EnergyDescrib) { // не превышать максимум
			enegrName = termineteAction.EnergyDescrib[energy]
		}
		font := getFontFromEnergy(energy)
		out = "Действие: <b>" + out + "</b><br><span style=\"" + font + "\">Энергичность: <b>" + enegrName + "</b></span><br>"

		return out
	}
	return ""
}

func getFontFromEnergy(energy int) string {
	switch energy {
	case 0:
		return "font-size:10px;color:#888888"
	case 1:
		return "font-size:11px;color:#666666"
	case 2:
		return "font-size:12px;color:#6699FF"
	case 3:
		return "font-size:13px;color:#0033FF"
	case 4:
		return "font-size:14px;color:#660099"
	case 5:
		return "font-size:15px;color:#000000"
	case 6:
		return "font-size:17px;color:#663300"
	case 7:
		return "font-size:19px;color:#CC6633;text-shadow: 0 0 1px #F74447,0 0 2px #F7B4C2,0 0 3px #F7DAE1;"
	case 8:
		return "font-size:22px;color:#FF3300;letter-spacing: 1px;text-shadow: 0 0 2px #F74447,0 0 4px #F7B4C2,0 0 8px #F7DAE1;"
	case 9:
		return "font-size:22px;color:#FF0066;letter-spacing: 2px;font-weight:bold;text-shadow: 0 0 2px #F74447,0 0 4px #F7B4C2,0 0 8px #F7DAE1;"
	case 10:
		return "font-size:22px;color:#000000"
	}
	return ""
}

/*
	совершить МОТОРНОЕ (ВЫДАТЬ ФРАЗУ) действие - Snn-часть автоматизма

cила действия сначала задается = 5, а потот корректируется мозжечковыми рефлексами
*/
func TerminatePraseAutomatizmActions(IDarr []int, energy int) string {
	// при моторном действии  меняются гомео-параметры:
	// expensesGomeostatParametersAfterAction(aI) болтать можно без устали?

	// выдать на ПУльт
	var out = ""
	if !isTeachQuestion {
		out += "Фраза Beast: "
	}

	for i := 0; i < len(IDarr); i++ {
		prase := word_sensor.GetPhraseStringsFromPhraseID(IDarr[i])
		if len(prase) > 0 {
			out += "<b>" + prase + "</b>"
		}
	}
	// название силы:
	if !isTeachQuestion {
		if energy < len(termineteAction.EnergyDescrib) {
			out += " " + termineteAction.EnergyDescrib[energy] + "</b>"
		}
	}
	return out
}

/*
	изменение гомео-параметров при действии

сила действия корректирует воздействие на параметр гомеостаза
*/
func expensesGomeostatParametersAfterAction(actID int, energy int) {
	if transfer.IsPsychicGameMode { // не воздействовать на гомео-параметры в игровом режиме
		return
	}
	se := termineteAction.TerminalActionsExpensesFromID[actID]
	if se != nil {
		for j := 0; j < len(se); j++ {
			// (2*aI.Energy/10) при силе==5 коэффициент будет 1, при силе==10 воздействие увеличиться в 2 раза
			if !gomeostas.NotAllowSetGomeostazParams {
				k := float64(2 * energy / 10)
				gomeostas.GomeostazParams[se[j].GomeoID] += se[j].Diff * k
				if gomeostas.GomeostazParams[se[j].GomeoID] > 100 {
					gomeostas.GomeostazParams[se[j].GomeoID] = 100
				}
				if gomeostas.GomeostazParams[se[j].GomeoID] < 0 {
					gomeostas.GomeostazParams[se[j].GomeoID] = 0
				}
			}
		}
	}
}

// запуск найденного автоматизма в цикле func consciousnessElementary()
func runConsciousnessAutomatizm(autmtzm *Automatizm) {
	if autmtzm == nil {
		return
	}
	RumAutomatizmID(autmtzm.ID)
	mentalInfoStruct.motorAtmzmID = 0
	mentalInfoStruct.noStaffAtmzmID = false
	motorActionEffect = autmtzm.Usefulness
}

/////////////////////////////////////////////////////////

/*
	запустить действие Игнорировать во время размышления if isRepressionStimulsNoise {

Это не автоматизм, а именно игнорирование - сообщение на пульт в виде действия
*/
func runIgnoreAction() {
	// игнорирования
	ignorImageID, _ := CreateNewlastActionsImageID(0, 0, []int{9}, nil, 0, 0, true)
	TerminateMotorAutomatizmActions(0, []int{ignorImageID}, 1, false)

	lib.WritePultConsol("<span style='color:blue;background-color:#FFFFA3;'>Игнорирование из-за глубокого размышления.</span>: ")

	// сбросить метку Не было действий более 100 пульсов
	CurrentInformationEnvironment.IsIdleness100pulse = false
}

//////////////////////////////////////////////////////////

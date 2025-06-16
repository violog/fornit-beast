/* распознаватели отклонения жизненных параметров GomeostazParams[id] от нормы

Здесь решаются две независимые задачи:
1. Определение значения  Улучшения или уходшения состояния - для оценки результата действия.
2. Оценки текщего состояния Плохо - Норма - хорошо.

1-я ЗАДАЧА
Улучшение или уходшение состояния - сигналы для оценки результата действия для automatizm_result.go,
которые передаются в психику через func BetterOrWorseNow при каждой стимуляции после действий Бота.
Для этого тслеживается ОДИН приоритетный гомео-параметр: PrioritetGomeoParID (func compareWithOld())
это соотвествует тому, что в каждый момент времени есть одна самая главная проблема (высшая значимость).
Для получения конечного сигнала ухудшения или улучшения (func BetterOrWorseNow) используется:
var GomeostazActionEffectPainV=0 // величина Боли
var GomeostazActionEffectJoyV=0 // величина радости
var GomeostazActionEffectPulsCount=0
(В psichic.mood.go: var painValue=0 var joyValue=0)
Эти значение есть даже при Норме всех жизненных параметров, если Оператор использует мотивационные Кнопки.

2-я ЗАДАЧА
Последовательность при изменении состояния:
1. Сначала распознается порции измненеия состояний по каждой шкале BadNormalWellValue в func badDetecting().
2. Изменения сохраняются в виде saveDinamicParams в func chandeDinamicPars() до следующего изменения.
3. Вычисляются суммарные отклонения по всем параметрам с учетом весов значимости параметра
curCommonBadNormalWellVal в func commonBadDetecting()
4. На пульт уходит индикация по каждому параметру и общее состояние в func GetCurGomeoStatus()
5. Экспоненциальная оценка изменения общего состояния (димит -10 и +10) commonDiffValue в func commonPerceptionNow()

При вызове func BetterOrWorseNow(kind int) возвращается информация о текущем эффекте значимости результата действия.
*/

package gomeostas

import (
	"BOT/lib"
	"math"
	"strconv"
	"strings"
)

/////////////////////////////////////////

/*
	старые значения для сравнения с PrioritetGomeoParID

Для оценки совершенного действия.
*/
var oldGomeoParID = make([]int, 10)

func initOldGomeoParID() {
	for id := 0; id < 9; id++ {
		oldGomeoParID[id] = int(GomeostazParams[id])
	}
}

// из психики: зафиксировать текущее состояние на момент срабатывания автоматизма
func BetterOrWorseInit() {
	initOldGomeoParID()
	GomeostazActionEffectPainV = 0
	GomeostazActionEffectJoyV = 0
}

// //////////////////////////////////////
// сравнить PrioritetGomeoParID со старым значением
func compareWithOld() int {
	// найти старое значение
	for id := 0; id < 9; id++ {
		if PrioritetGomeoParID == id {
			diff := 0
			if id == 1 { // Энергия наоборот
				diff = int(GomeostazParams[id]) - oldGomeoParID[id]
			} else {
				diff = oldGomeoParID[id] - int(GomeostazParams[id])
			}
			//oldGomeoParID[id] = int(GomeostazParams[id]) // сразу обновить старое состояние
			// после каждой оценки в func BetterOrWorseNow чтобы не устаревали сильно параметры.
			initOldGomeoParID()
			return diff
		}
	}
	return 0
}

////////////////////////////////////////////////////

// название Базового состояния из его ID str:=gomeostas.getBaseCondFromID(id)
func GetBaseCondFromID(id int) string {
	var out = ""
	switch id {
	case 1:
		out = "Плохo"
	case 2:
		out = "Норма"
	case 3:
		out = "Хорошо"
	}
	return out
}

// момент появления Хорошо (в тиках пульса)
//var CommonWellValueStart = 0

/* детекторы по каждому из жизненных параметров
имеют те же значения, что и в CommonBadNormalWell
Алгоритм на примере энергии.
Если энергия истощилась, то BadNormalWellValue[id] будет тем отрицательнее, чем сильнее истощилась.
Если началось восполнение энергии, то BadNormalWell[id]=3 (хорошо) и BadNormalWellValue[id] уменьшается по мере насыщения.
Но если насыщение остановилось (не меняется в течение ),
то через время BadNormalWellValue[id] снова становится ==1 (плохо), если параметр еще не восстановлен: остался еще голоден.
В природе чем сильнее голод, тем сильнее Хорошо с началом его удовлетворения и это Хорошо уменьшается с насыщением.
Но если еды было мало, то тварь довольно скоро опять почувствует голод, но уже не такой большой.
*/

// Значения динамики состояния жизненных параметров, - выходная информация детектора состояния
var BadNormalWell = make([]int, 10)      // 1 - Похо, 2 - Норма, 3 - Хорошо
var BadNormalWellValue = make([]int, 10) // насколько плохо или хорошо, значение от  -100 0 до 100 В НОРМЕ=0
//////////////////////////////////////////////////

// пороги реальной нормы
var compareNorma = make([]int, 10)

// пороги условного начала выхода параметров из нормы
var compareLimites = make([]int, 10)

func initBadDetector() {
	for id, _ := range GomeostazParams {
		BadNormalWell[id] = 2 // норма
		BadNormalWellValue[id] = 0
		dinamicValueStart[id] = 0
		//	oldBadNormalWellValue[id] = 0
	}
	/* порог выхода из нормы жиз.параметров:
	 */
	path := lib.GetMainPathExeFile()
	lines, _ := lib.ReadLines(path + "/memory_reflex/GomeostazLimits.txt")
	for i := 0; i < len(lines); i++ {
		p := strings.Split(lines[i], "|")
		id, _ := strconv.Atoi(p[0])
		/* 	В редакторе /memory_reflex/GomeostazLimits.txt порог начинается со слабого отклонения от нормы.
			Степень отклонения от нормы делится на сегменты по 1/5 всего порога,
		т.е. на критическое Плохо приходится значение порога: p[1]/5
		*/
		compareNorma[id], _ = strconv.Atoi(p[1])
		// реально плохо становится при "Значительное отклонение" - середина от ухода из нормы

		compar := compareNorma[id] / 2
		compareLimites[id] = int(compar)

		if id == 3 {
			id = 3
		}
	}
	return
}

// время в секундах удержания состояния saveDinamicParams[] для возврата в Норму
var dinamicTime = 50
var dinamicValueStart = make([]int, 10) //make(map[int]int)	// время возникновения динамического состояния Хорошо или Плохо

/*
	Определение текущего состояния по каждому базовому параметру

затем в commonBadDetecting() вычисляется общее интегральное состояние
*/
var prevGomeostazParams = make([]int, 10) // предыдыщие значения ж.параметров
var dinamicParams = make([]int, 10)       // какие парамерты изменились, в какую сторону и на сколько, (отрицательные - стало хуже)
var saveDinamicParams = make([]int, 10)   //удержание значений dinamicParams до их изменений (инд. на пульте)
var wasChangeParams = false               // были значительные изменения параметров
var prevCommonBadDetectingPulsCount = 0

func badDetecting() {
	if NotAllowSetGomeostazParams {
		return
	}

	dinamicParams = make([]int, 10) // очистить
	var difSensorPar = 2            // на какую величину изменятся параметры, чтобы это заметил детект
	isChangeDiff := false           // были значимые изменения параметров

	for id := 1; id < 9; id++ {
		curG := int(GomeostazParams[id])
		// насколько изменились параметры с прошлого сканирования, т.е. отрицательные - стало хуже
		diff := curG - prevGomeostazParams[id]
		if id > 1 { // для энергии наоборот
			diff *= -1
		}
		////////////////////////////////

		// Если нет diff (запуск или простой>50сек) то нужно определитьс с состонием статически:
		if diff == 0 {
			var isBad = false
			// только для энергии
			if id == 1 && curG < compareLimites[id] { // вышел за Сильное отклонение
				isBad = true
			}
			if id > 1 && curG > 100-compareLimites[id] { // вышел за Сильное отклонение
				isBad = true
			}
			if isBad { // когда слишком плохо
				BadNormalWell[id] = 1 // плохо
				// принимает значение параметра
				if id == 1 {
					BadNormalWellValue[id] = -(100 - curG)
				} else {
					BadNormalWellValue[id] = -curG
				}
				dinamicValueStart[id] = 0
				//continue
			}

			isWell := isNormaForGomeoParam(id) // определеие чистой нормы для параметра
			if isWell {                        // когда явная норма
				BadNormalWell[id] = 2 //
				BadNormalWellValue[id] = 0
			}
		} // КОНЕЦ if diff==0{
		///////////////////////////////////

		if PulsCount > 2 { // чтобы было уже установленные значения PulsCount!!!
			//////////// улучшения или ухудшения в диапазоне слабых отклонений от нормы
			if diff != 0 {
				if diff > difSensorPar || diff < -difSensorPar {

					if id != 7 { // не учитывать автоматическое улучшение самосохранения при малой опасности (func gomeostazUpdate())
						dinamicParams[id] = diff
						isChangeDiff = true // были значимые изменения параметров
					}

					prevGomeostazParams[id] = curG
				}

				if diff > difSensorPar { // улучшение
					// если было динамическое улучшение
					dinamicValueStart[id] = PulsCount // контроль точкой прерывания
				}
				if diff < -difSensorPar { // ухудшение
					dinamicValueStart[id] = PulsCount // контроль точкой прерывания
				}

				//if diff != 0{
			} else {
				prevGomeostazParams[id] = curG
			}

			//if PulsCount > 2 {
		} else {
			prevGomeostazParams[id] = curG
		}
	} //for id := 1; id < 9; id++ {
	/////////////////////////////

	if isChangeDiff { // были значимые изменения параметров
		chandeDinamicPars()    // сохранить до следующего изменения
		wasChangeParams = true // были изменения параметров
	} else {
		wasChangeParams = false // не было изменения параметров
	}
	// если прошли интревалы динамического Хорошо и Полохо, то погасить их
	checkDinamicTimes()

	if wasAnyChangeParams() {
		// не реагировать 2 пульса чтобы не влияли последствия изменений параметров (изменяются Стресс) Защита от дребезга
		if PulsCount-prevCommonBadDetectingPulsCount > 2 {
			prevCommonBadDetectingPulsCount = PulsCount
			commonBadDetecting()
		}
	}
}

// определеие чистой нормы для параметра
func isNormaForGomeoParam(id int) bool {
	curG := int(GomeostazParams[id])
	var isWell = false // чистая норма
	// только для энергии
	if id == 1 && curG > compareNorma[id] { // вышел в чистую норму
		isWell = true
	}
	if id > 1 && curG < compareNorma[id] { // вышел в чистую норму
		isWell = true
	}
	if isWell { // когда явная норма
		BadNormalWell[id] = 2 //
		BadNormalWellValue[id] = 0
	}
	return isWell
}

//////////////////////////////////////////////

// были значимые изменения параметров, у тех, что не изменились остается BadNormalWell[id] = 2
func chandeDinamicPars() {
	// !!! НЕТ curCommonBadNormalWellVal = 0
	saveDinamicParams = make([]int, 10)
	for id := 1; id < 9; id++ {
		saveDinamicParams[id] = dinamicParams[id]
		// чтобы BadNormalWell[] сохранились
		//BadNormalWell[id] = 2
		if dinamicParams[id] > 0 {
			BadNormalWellValue[id] = dinamicParams[id]
			BadNormalWell[id] = 3
		}
		if dinamicParams[id] < 0 {
			BadNormalWellValue[id] = dinamicParams[id]
			BadNormalWell[id] = 1
		}
	}
	return
}

// гашение Хорошо и Плохо через интервал dinamicTime в норму
func checkDinamicTimes() {
	for id := 1; id < 9; id++ {
		if dinamicValueStart[id] > 0 && (dinamicValueStart[id]+dinamicTime < PulsCount) {
			BadNormalWell[id] = 2 // норма
			BadNormalWellValue[id] = 0
			dinamicValueStart[id] = 0
		}
	}
}

////////////////////////////////__________________детектор общего базового состояния_______________

/*
	детектор общего базового состояния (индикация на Пуьте document.getElementById('common_status_id'))

1 - Похо, 2 - Норма, 3 - Хорошо (==выход из состояния Плохо)
*/
var CommonBadNormalWell = 2 // 1 - Похо, 2 - Норма, 3 - Хорошо из других пакетов: gomeostas.CommonBadNormalWell

/*
	Суммарное значение гомеопраметров с учетом весов параметров.

изменяется - в func commonBadDetecting()
в curCommonBadNormalWellVal будет нарастать изменение
curCommonBadNormalWellVal НИКОГДА НЕ Обнуляется! потому как изменения достаточно точно компенсируются
и нет критерия, когда можно было бы делать новую инициализацию.
Но и не нужно пытаться инициализировать время от времени, потому что главное - влияние текущих изменений
*/
var curCommonBadNormalWellVal = 0

// предыдущее значение  curCommonBadNormalWellVal
var OldCommonBadNormalWellVal = 0 // предыдущее curCommonBadNormalWellVal

/*
	Распознавание интегрального состояния CommonBadNormalWell

пороговый (compareLevel) сумматор значений состояний Плохо
Логика работы:
 1. Если суммарное значение Плохо выше порога - базовое состояние Плохо
 2. Если суммарное значение Плохо ниже порога и:
    2.1. если предыдущее базовое состояние было Плохо - базовое состояние Хорошо
    2.2. если предыдущее базовое состояние было Норма или Хорошо - базовое состояние Норма
 3. Если базовое состояние Хорошо держится больше dinamicTime (50 сек) - базовое состояние Норма
*/
var prevCommonBadNormalWellVal = 0 // только для func commonBadDetecting()
/*
	вариант без saveDinamicParams, непосредственное сравнение гомео-параметров при изменениях

не требует инициализации (всегда корректен).
*/
func commonBadDetecting() {
	if NotAllowSetGomeostazParams {
		return
	}

	/* Из суммы CommonBadValue исключать параметры, не влияющие на жизненные показатели:
	гон, потребность в общении, в обучении, любопытство
	*/
	prevCommonBadNormalWellVal = curCommonBadNormalWellVal
	curCommonBadNormalWellVal = 0
	isCommonWell := true // true - все важные параметры в норме
	for id := 1; id < 9; id++ {
		if id == 3 || id == 4 || id == 5 || id == 6 || id == 7 {
			continue
		}

		//если все параметры в норме, то не делать состояние плохо
		isWell := isNormaForGomeoParam(id) // определеие чистой нормы для параметра
		if !isWell {                       // когда явная норма
			isCommonWell = false
		}

		// только текущее изменение всех параметров, а не сами параметры умножаем на вес значимости параметра
		if id == 1 {
			curCommonBadNormalWellVal += (100 - int(GomeostazParams[id])) * GomeostazParamsWeight[id]
		} else {
			curCommonBadNormalWellVal += int(GomeostazParams[id]) * GomeostazParamsWeight[id]
		}
	}

	//	fmt.Print(strconv.Itoa(prevCommonBadNormalWellVal) + " - " + strconv.Itoa(curCommonBadNormalWellVal) + "\r\n")

	var compareLevel = 100  // порог начала измененияя состояния
	CommonBadNormalWell = 2 // норма

	//diff := curCommonBadNormalWellVal - prevCommonBadNormalWellVal
	if prevCommonBadNormalWellVal > curCommonBadNormalWellVal+compareLevel { // Хорошо
		CommonBadNormalWell = 3
		SetGomeoAtasEffect(3) // эффект при Критическом выходе из нормы важных жизненных параметров
	}
	//если все параметры в норме, то не делать состояние плохо
	if (curCommonBadNormalWellVal > prevCommonBadNormalWellVal+compareLevel) && !isCommonWell { // Плохо
		CommonBadNormalWell = 1 // Плохо
		SetGomeoAtasEffect(1)   // эффект при Критическом выходе из нормы важных жизненных параметров
	}
	return
}

/*
	прежний вариант, с использованием saveDinamicParams, непонятно как инициализировать время от времени
т.к. curCommonBadNormalWellVal постоянно накапливает изменения.

func commonBadDetecting() {
	if NotAllowSetGomeostazParams {
		return
	}

	// Из суммы CommonBadValue исключать параметры, не влияющие на жизненные показатели:
	//гон, потребность в общении, в обучении, любопытство
	newDiff := 0
	isCommonWell := true // true - все важные параметры в норме
	for id := 1; id < 9; id++ {
		if id == 3 || id == 4 || id == 5 || id == 6 || id == 7 {
			continue
		}

		isWell := isNormaForGomeoParam(id) // определеие чистой нормы для параметра
		if !isWell {                       // когда явная норма
			isCommonWell = false
		}

		// только текущее изменение всех параметров, а не сами параметры умножаем на вес значимости параметра
		newDiff += saveDinamicParams[id] * GomeostazParamsWeight[id]
	}

	if wasChangeParams { // были изменения параметров
		prevCommonBadNormalWellVal = curCommonBadNormalWellVal
		curCommonBadNormalWellVal += newDiff // текущие изменения модифицируют прежнее значение curCommonBadNormalWellVal
		//		oldNewDiff = newDiff
	}

	fmt.Print(strconv.Itoa(prevCommonBadNormalWellVal) + " - " + strconv.Itoa(curCommonBadNormalWellVal) + "\r\n")

	var compareLevel = 100  // порог начала измененияя состояния
	CommonBadNormalWell = 2 // норма

	//diff := curCommonBadNormalWellVal - prevCommonBadNormalWellVal
	if curCommonBadNormalWellVal > prevCommonBadNormalWellVal+compareLevel { // Хорошо
		CommonBadNormalWell = 3
		SetGomeoAtasEffect(3) // эффект при Критическом выходе из нормы важных жизненных параметров
	}
	if (prevCommonBadNormalWellVal > curCommonBadNormalWellVal+compareLevel) && !isCommonWell { // Плохо
		CommonBadNormalWell = 1 // Плохо
		SetGomeoAtasEffect(1)   // эффект при Критическом выходе из нормы важных жизненных параметров
	}
	return
}
*/
///////////////////////////////////////////////////

var prevGomeoParams = make([]int, 10) // предыдущие значения ж.параметров
func wasAnyChangeParams() bool {
	new := false
	for id := 1; id < 9; id++ {
		if id == 3 || id == 4 || id == 5 || id == 6 || id == 7 {
			continue
		}
		if prevGomeoParams[id] != saveDinamicParams[id] {
			new = true
		}
		prevGomeoParams[id] = saveDinamicParams[id]
	}
	return new
}
func initAnyChangeParams() {
	for id := 0; id < 9; id++ {
		prevGomeoParams[id] = saveDinamicParams[id]
	}
}

///////////////////////////////////////////////

/*
	Снять режимы Плохо и Хорошо при нажатии на кнопку Успокоить

По времени dinamicTime они снимаются в func checkDinamicTimes()
*/
func CliarWellBad() {

	for id := 1; id < 9; id++ {
		saveDinamicParams[id] = 0
		BadNormalWellValue[id] = 0
		dinamicValueStart[id] = 0
		BadNormalWell[id] = 2 // норма
	}
	OldCommonBadNormalWellVal = curCommonBadNormalWellVal
	CommonBadNormalWell = 2 //  норма
}

///////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////
// для Пульта индикация слайдеров и общее Плохо-Норма-Хорошоы
func GetCurGomeoStatus() string {
	var out = "0;" + strconv.Itoa(CommonBadNormalWell) + "|"
	//for id, v := range BadNormalWell {
	for id := 1; id < 9; id++ {
		out += strconv.Itoa(id) + ";" + strconv.Itoa(BadNormalWell[id]) + "|"
	}
	out += "@0|"
	for id := 1; id < 9; id++ {
		out += strconv.Itoa(saveDinamicParams[id]) + "|"
	}
	return out
}

/////////////////////////////////////////////////

/* Стало лучше или хуже?.
Вызывается из психики res:=gomeostas.BetterOrWorseNow()

BetterOrWorseNow вызывается в случаях:
1. При стимуле в период ожидания LastRunAutomatizmPulsCount (automatism_tree.go func afterTreeActivation())
2. Процесс сновидения или мечтаний (understanding_problem_tree.go func ProblemTreeActivation())

LastRunAutomatizmPulsCount устанавливается в 2-ч случаях:
1. когда было действие бота и он ждет ответ для оценки эффекта (после func RumAutomatizm )
2. Совершение произвольного действия (после func showVolutionAction)


ВОЗВРАЩАЕТ:
commonDiffValue - насколько изменилось общее состояние, значение от -10(максимально Плохо) через 0 до 10(максимально Хорошо)
GomeoParIdSuccesArr - стали лучше следующие г.параметры []int гоменостаза

Если было очень плохо, а стало не очень плохо, то commonDiffValue станет позитивным.
*/
//	var GomeoParIdBadArr []int// спиоск ж.параметров с ухудшнием
func BetterOrWorseNow() (int, []int) { // только для func BetterOrWorseNow(kind int)
	/*!!! НУЖНО ДАТЬ ВРЕМЯ НА ТО, ЧТОБЫ ПОЛУЧИТЬ НОВОЕ после Стимула ЗНАЧЕНИЕ ГОМЕОПАРАМТРОВ
	потому как действие попадает не в тот пульс, что измерение последствий, а всегда ему предшествует.
	т.е. BetterOrWorseNow нужно откладывать чуть более пульса, на 1.5 сек.
	*/
	//	time.Sleep(1500 * time.Microsecond) нормально успевают измениться параметры гомеостаза, так что не нужно

	// какие id парамектров улучшились - для Automatizm.GomeoIdSuccesArr (какие ID гомео-параметров улучшает это действие)
	var GomeoParIdSuccesArr []int // спиоск ж.параметров с улучшением
	for id := 1; id < 9; id++ {
		if saveDinamicParams[id] > 0 {
			GomeoParIdSuccesArr = append(GomeoParIdSuccesArr, id)
		}
	}
	/////////////////

	//старые значения приоритетного жизненного параметра сравниваются с новым PrioritetGomeoParID
	diff := compareWithOld() //
	diff0 := diff
	// Значение экспоненциально стремится к пределам -10 и 10
	//diff = int(10.0 - 10.0 / math.Exp(float64(lib.Abs(diff)) * 0.17))
	eval := int(11 - 11/math.Exp(float64(lib.Abs(diff))*0.1))
	if diff < 0 {
		diff = -eval
	} else {
		diff = eval
	}

	commonDiffValue := 0 // итоговый мотивационный эффект для оценки действия

	/* Боль является преимущественным мотивационным фактором, чем Радость. Детекторов боли несопоставимо больше.
	   Боль так же имеет преимущества перед улучшением жизненных показателей.
	*/
	if GomeostazActionEffectPainV > 0 { // Значение Боли
		if diff >= 0 {
			commonDiffValue = int(diff/3) - GomeostazActionEffectPainV
		} else { // diff<0 - выбрать самое плохое
			commonDiffValue = -lib.Max(-diff, GomeostazActionEffectPainV)
		}
	}

	// Если Радость
	if GomeostazActionEffectJoyV > 0 && GomeostazActionEffectPainV == 0 {
		if diff < -5 { // значительное ухудшение ж.параметра
			commonDiffValue = diff
		} else {
			commonDiffValue = lib.Max(diff, GomeostazActionEffectJoyV)
		}
	}

	lib.WritePultConsol("<span style='color:blue;background-color:#C5FFCC'>ПРАВИЛА. Эффект commonDiffValue определен в func BetterOrWorseNow(): <b>" +
		strconv.Itoa(commonDiffValue) + "</b> (гомео diff=" +
		strconv.Itoa(diff0) + ", Боль: " +
		strconv.Itoa(GomeostazActionEffectPainV) + ", радость: " +
		strconv.Itoa(GomeostazActionEffectJoyV) + ")</span>")

	// на накапливать эти состояния
	GomeostazActionEffectPainV = 0
	GomeostazActionEffectJoyV = 0

	// НЕТ curCommonBadNormalWellVal = 0

	return commonDiffValue, GomeoParIdSuccesArr
}

//////////////////////////////////////////////////////////////////

/*
	достаточно ли изменились гомео-параметры при активации дерева не оператором

чтобы проходить func consciousnessElementary()

Значения diff нужно оптимизировать при необходимости.
*/
var prevGomeoParID = make([]int, 10)

func GetGomeoParsDiff() bool {
	toAttention := false
	for id := 1; id < 9; id++ {
		if saveDinamicParams[id] > 0 {

		}
		diff := 0
		switch id {
		case 1: //Энергия
			diff = 20
		case 2: //Стресс
			diff = 40
		case 3: //Гон
			diff = 60
		case 4: //Потребность в общении
			diff = 50
		case 5: //Потребность в обучении
			diff = 50
		case 6: //Поиск
			diff = 50
		case 7: //Самосохранение
			diff = 40
		case 8: //Повреждения
			diff = 10
		}
		if lib.Abs(int(GomeostazParams[id])-prevGomeoParID[id]) > diff {
			toAttention = true
		}
		if toAttention { // достаточно сильное изменение, чтобы привлечь внимание
			for id := 1; id < 9; id++ {
				prevGomeoParID[id] = int(GomeostazParams[id])
			}
			return true
		}
	}
	return false
}

////////////////////////////////////////////////////////////////

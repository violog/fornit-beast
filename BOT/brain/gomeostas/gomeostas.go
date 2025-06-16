/*
	параметры определены в http://go/pages/gomeostaz.php

Жизненные параметры гомеостаза и Базовые стили поведения (базовые контексты рефлексов)

1 - энергия 2 - стресс 3 - гон 4-потребность в общении 5-потребность в обучении
6- Поиск, 7- Самосохранение, 8 - Повреждения
файл сохранения состояния параметров: files/GomeostazParams.txt
base_context_activnost.txt
Для параметров гомеостаза, напрямую не связанных с жизнеобеспечением (parMaxPulsCount: гон, потребность в общении,
потребность в обучении и любопытство) организована цикличность: при нарастании параметра до максимума,
он удерживается в течении 20 секунд, а потом сбрасывается.
Это позволяет создавать достаточные по времени периоды специфических контекстов реагирования.
*/
package gomeostas

// import word_sensor "BOT/brain/words_sensor"

import (
	"BOT/lib"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//////////////////////////////////////////////////////////

/*
	Приоритетно отслеживаемый ID гомео-параметр.

Только один, самый значимый гомеопараметр является объектом приоритетного внимания к изменениям.
Это параметр, который достиг критического порога "Сильное отклонение".
Из всех параметров, превысивший этот порог выбирается самый весомый.

Если нет параметров с достаточно значимым отклонением от нормы, то PrioritetGomeoParID=0
и нет необходимости обращать внимание на жизненные параметры.

PrioritetGomeoParID имеет гистерезис срабатывания, чтобы избажать дребезга.
*/
var PrioritetGomeoParID = 0

func updatePrioritetGomeoPar() { // с каждым пульсом, после определения текущих значений параметров
	var Gomeovals = make([]int, 10) // %шкал * вес
	for id := 1; id < 9; id++ {
		valP := GomeostazParams[id]
		if id == 1 { // у энергии обратная шкала
			valP = 100 - valP
		}
		Gomeovals[id] = int(valP * float64(GomeostazParamsWeight[id]))
	}
	max := 0
	maxID := 0
	for id := 1; id < 9; id++ {
		if Gomeovals[id] > max {
			max = Gomeovals[id]
			maxID = id
		}
	}
	// а теперь просто выбрать первый ID
	PrioritetGomeoParID = maxID
	/*  а пусть даже в диапазоне нормы будет - для func getCurGomeoPotencial()
	if GomeostazParams[PrioritetGomeoParID]<30{// если все в норме, то ничего не отслеживать
		PrioritetGomeoParID=0
	}*/
}

/////////////////////////////////////////////////////////////////////////////

// Текущие значения параметров гомеостаза заполняется из файла files/GomeostazParams.txt
var GomeostazParams = make([]float64, 10) //make(map[int]float64)
// старые значения, обновляемые при ухудшении ??????
var OldGomeostazParams = make([]float64, 10) //make(map[int]float64)
// скорость изменения параметров со временем
var GomeostazParamsSpeed = make([]int, 10) //make(map[int]int)
// веса значимости для GomeostazSensor От 0 до 100 %
var GomeostazParamsWeight = make([]int, 10) //make(map[int]int)
// правила активирования Базовых контекстов. Здесь могут быть значения для Боли и Радости (ID>12)
var GomeostazActivnostArr = make([][]string, 10) // make(map[int][]string)

// / НАСТРОЕНИЕ ПРИ ПОСЫЛКЕ СООБЩЕНИЯ С ПУЛЬТА
var moodeDialogEffects = make([]int, 30) // используются только с 20 по 26
// нажатия кнопок действий Пульта дают гомео-эффект:
var GomeostazActionEffectArr = make([]string, 30) //make(map[int]string)
/*
	нажатия кнопок действий Пульта дают значение эффекта Боли и Радости

На уровне рефлексов не оказывает действия - только в период ожидания результата действий Beast
- для (де)мотивации в automatizm_result.go
*/
var GomeostazActionCommonEffectArr = make([]string, 30) //make(map[int]string)

/* ДЕТЕКТОРЫ ТРАВМАТИЧЕСКИХ и УЛУЧШАЮЩИХ ПОСЛЕДСТВИЙ
Величина "Боли" или "радости" - вкавычках потому что на этом уровне уровне рефлекснв нет ощещния Боли,
а есть только детекторы травматических последствий.
Существуют рецепторы, которые могут реагировать на улучшающие воздействия или на приятные стимулы.
Они известны как рецепторы приятных ощущений или рецепторы белого вещества.
Эти рецепторы обнаружены в разных частях тела и могут реагировать на приятные тактильные,
тепловые или химические стимулы.
Например, механические рецепторы на коже способны реагировать на приятное прикосновение или массаж, половые сенсоры.
Также, терморецепторы могут реагировать на приятные тепловые стимулы.
Вкусовые рецепторы в языке могут реагировать на приятные вкусы.
Эти рецепторы передают сигналы в нервную систему, вызывая положительные реакции и улучшение настроения.
Они играют важную роль в нашем восприятии и ощущении приятности, а также могут стимулировать рефлекторные реакции
и двигательные ответы.
Однако, важно отметить, что рецепторы приятных ощущений не обязательно имеют
противоположные рецепторы травматических воздействий.
Рецепторы боли и рецепторы приятных ощущений могут функционировать независимо
друг от друга и реагировать на различные типы стимулов.
Боль и радость: fornit.ru/67646

Здесь применяются детекторы общих (де)мотивирующих действий - пока только в период ожидания результата действий Beast
- для (де)мотивации в automatizm_result.go TODO но могут использовать в безусловных рефлексах.
Максимальные значения до 10.

Боль и Радость возникают при:
1. Критическом выходе из нормы важных жизненных параметров func SetGomeoAtasEffect() - проверяется с каждым пульсом
2. При мотивации Оператором нажатием кнопок func SetGomeostazActionCommonEffectArr
3. При мотивации Оператором установкой настроения при отправке сообщения func SetMoodePraseEffectArr

На рефлекторном уровне нет ощущений Боли, а есть только мотивационное действие,
оно не влияет на рефлексы вообще никак, а только - на оценку результата действий.

Гасится в func BetterOrWorseNow при каждом Стимуле после Ответа

Боль является преимущественным мотивационным фактором, чем Радость. Детекторов боли несопоставимо больше.
Боль так же имеет преимущества перед улучшением жизненных показателей.
*/
// В psichic.mood.go: var painValue=0 var joyValue=0
var GomeostazActionEffectPainV = 0 // величина Боли
var GomeostazActionEffectJoyV = 0  // величина радости
//var GomeostazActionEffectPulsCount=0  СБРАСЫВАЮТСЯ в func BetterOrWorseNow при каждом Стимуле после Ответа

// эффект нажатия кнопок действий на пульте
func SetGomeostazActionCommonEffectArr(actionID int) {
	//GomeostazActionEffectPulsCount = PulsCount
	effect := GomeostazActionCommonEffectArr[actionID]
	kind := effect[:1]
	val, _ := strconv.Atoi(effect[1:])
	if kind == "-" {
		GomeostazActionEffectPainV = val // величина Боли
	} else {
		GomeostazActionEffectJoyV = val // величина радости
	}
	return
}

// эффект настроения при отправке сообщения
func SetMoodePraseEffectArr(moodID int) {

	effect := moodeDialogEffects[moodID]
	if effect > 0 {
		GomeostazActionEffectPainV = effect // величина Боли
	} else {
		GomeostazActionEffectJoyV = effect // величина радости
	}
}

// эффект при Критическом выходе из нормы важных жизненных параметров, проверяется с каждым пульсом
func SetGomeoAtasEffect(kind int) {
	effect := 0
	if kind == 3 { // временно удерживаемое состояние Плохo (CommonBadNormalWell = 3)
		if GomeostazParams[1] < 15 { //Энергия
			effect += 5
		}
		if GomeostazParams[2] > 85 { // Стресс
			effect += 3
		}
		if GomeostazParams[8] > 60 { //Повреждения
			effect += 6
		}
		GomeostazActionEffectPainV = effect
	}
	if kind == 1 { // временно удерживаемое состояние {Хорошо} (CommonBadNormalWell = 1)
		if GomeostazParams[1] > 85 { //Энергия
			effect += 4
		}
		if GomeostazParams[2] < 15 { // Стресс
			effect += 2
		}
		if GomeostazParams[8] < 15 { //Повреждения
			effect += 5
		}
		GomeostazActionEffectJoyV = effect
	}
}

//////////////////////////////////////////////////

var PeriodPulsCount = 20             // 20 пульсов удерживается максимальное значение некритичных для жизни 4-х параметров
var parMaxPulsCount = make([]int, 4) // 0 - гон, 1 -потребность в общении, 2 -потребность в обучении, 3 - любопытство

func init() {
	/*
		// подбор нужной крутизны экспоненты для BetterOrWorseNow() - множителем k
		var str="|"
		var val=0.0
		var k=0.17
		for i := 0; i < 70; i++ {
			if i>0 && i%10==0{str+="\r\n"}
			val=10.0 - 10.0/math.Exp(float64(i)*k)
			str+=strconv.Itoa(i)+":"+fmt.Sprintf("%10.2f", val)+"|"
		}
	*/
	path := lib.GetMainPathExeFile()
	lines, _ := lib.ReadLines(path + "/memory_reflex/GomeostazParams.txt")
	if len(lines) < 7 { // испорчен файл, восстановить
		var def = "1|10\r\n2|10\r\n3|0\r\n4|0\r\n5|0\r\n6|0\r\n7|0\r\n8|0"
		lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_reflex/GomeostazParams.txt", def)
		lines, _ = lib.ReadLines(path + "/memory_reflex/GomeostazParams.txt")
	}
	for i := 0; i < len(lines); i++ {
		p := strings.Split(lines[i], "|")
		id, _ := strconv.Atoi(p[0])
		val, _ := strconv.ParseFloat(p[1], 32)
		GomeostazParams[id] = val
		OldGomeostazParams[id] = val
	}
	lines, _ = lib.ReadLines(path + "/memory_reflex/GomeostasWeight.txt")
	for i := 0; i < len(lines); i++ {
		p := strings.Split(lines[i], "|")
		id, _ := strconv.Atoi(p[0])
		weight, _ := strconv.Atoi(p[1])
		speed, _ := strconv.Atoi(p[2])
		GomeostazParamsWeight[id] = weight
		GomeostazParamsSpeed[id] = speed
	}

	GomeostazActivnostArr = make([][]string, 10) //make(map[int][]string)
	lines, _ = lib.ReadLines(path + "/memory_reflex/base_context_activnost.txt")
	for i := 0; i < len(lines); i++ {
		p := strings.Split(lines[i], "|")
		id, _ := strconv.Atoi(p[0]) // ID параметра гомеостаза
		GomeostazActivnostArr[id] = make([]string, 7)
		GomeostazActivnostArr[id][0] = p[1] // плохо
		GomeostazActivnostArr[id][1] = p[2] // хорошо
		GomeostazActivnostArr[id][2] = p[3] //Норма
		GomeostazActivnostArr[id][3] = p[4] //Слабое отклонение
		GomeostazActivnostArr[id][4] = p[5] //Значительное отклонение
		GomeostazActivnostArr[id][5] = p[6] //Сильное отклонение
		GomeostazActivnostArr[id][6] = p[7] //Критически опасное отклонение
	}

	////////////////////////////////////////////////////////////
	GomeostazActionEffectArr = make([]string, 30)       //make(map[int]string)
	GomeostazActionCommonEffectArr = make([]string, 30) //make(map[int]string)
	lines, _ = lib.ReadLines(path + "/memory_reflex/Gomeostaz_pult_actions.txt")
	for i := 0; i < len(lines); i++ {
		p := strings.Split(lines[i], "|")
		id, _ := strconv.Atoi(p[0]) // ID параметра гомеостаза
		GomeostazActionEffectArr[id] = p[1]
		GomeostazActionCommonEffectArr[id] = p[2] // значение эффекта Боли и Радости
	}

	lines, _ = lib.ReadLines(lib.GetMainPathExeFile() + "/memory_reflex/moode_dialog_effects.txt")
	for i := 0; i < len(lines); i++ {
		p := strings.Split(lines[i], "|")
		id, _ := strconv.Atoi(p[0])
		val, _ := strconv.Atoi(p[1])
		moodeDialogEffects[id] = val
	}

	initBadDetector()
	initContextDetector()
	initOldGomeoParID()

	commonBadDetecting() // для начала
	return
}

/*
	приравнять OldGomeostazParams GomeostazParams

При каждом сравнении старого с новым приравнивать Для определения вектора изменения с каждым пульсом
*/
func copyToOldPsrams() {
	//for id, _ := range GomeostazParams { ГОНКИ чтения-запись ПОЯВЛЯЮТСЯ ТОЛЬКО ПРИ range
	for id := 1; id < 9; id++ {
		OldGomeostazParams[id] = GomeostazParams[id]
	}
}

/* ПУЛЬС ГОМЕОСТАЗА - обработка раз в секунду */
var PulsCount = 0      // передача тика Пульса из brine.go
var LifeTime = 0       // время жизни в числе пульсов
var EvolushnStage = 0  // стадия развития
var IsSlipping = false // флаг фазы сна
// коррекция текущего состояния гомеостаза и базового контекста с каждым пульсом
func GomeostazPuls(evolushnStage int, lifeTime int, puls int, isSlipping bool) {
	LifeTime = lifeTime
	EvolushnStage = evolushnStage
	PulsCount = puls // передача номера тика из более низкоуровневого пакета
	IsSlipping = isSlipping

	if EvolushnStage >= 3 { // Период подражания
		IsLevelBeginParam5 = true // уровень развития для Потребность в обучении достигнут
	}
	if EvolushnStage >= 4 { //Период преступной инициативы
		IsLevelBeginParam3 = true // уровень развития для Гона достигнут
	}
	gomeostazUpdate()

	badDetecting()
	// детектор изменения базового состояния и контекстов - проверка по каждому пульсу
	changingConditionsDetector()
	// При каждом сравнении старого с новым приравнивать Для определения вектора изменения с каждым пульсом
	copyToOldPsrams()

	updatePrioritetGomeoPar() // опредеелние приоритетного жизненного параметра
}

var NotAllowSetGomeostazParams = false // флаг процесса изменения величины базового параметра
// изменить параметр на величину
func ChangeGomeostazParametr(id int, diff float64) {
	NotAllowSetGomeostazParams = true
	OldGomeostazParams[id] = GomeostazParams[id]
	GomeostazParams[id] += diff
	if GomeostazParams[id] < 0 {
		GomeostazParams[id] = 0
	}
	if GomeostazParams[id] > 100 {
		GomeostazParams[id] = 100
	}
	NotAllowSetGomeostazParams = false
}

// выдать текущие значения жизненных параметров
func GetCurGomeoParams() string {
	if NotAllowSetGomeostazParams {
		return ""
	}
	var out = ""
	//for id, _ := range GomeostazParams { ГОНКИ чтения-запись ПОЯВЛЯЮТСЯ ТОЛЬКО ПРИ range
	for id := 1; id < 9; id++ {
		out += strconv.Itoa(int(id)) + ";" + strconv.Itoa(int(GomeostazParams[id])) + "|"
	}
	return out
}

// установка параметров гомеостаза с Пульта
func SetCurGomeoParams(parID int, parVal string) {
	NotAllowSetGomeostazParams = true
	GomeostazParams[parID], _ = strconv.ParseFloat(parVal, 64)
	SaveCurrentGomeoParams()
	NotAllowSetGomeostazParams = false
	/* определить базовые контексты при новых пераметрах гомеостаза baseContextUpdate()
	потом очистить сенсоры слов и стек слов и переактивировать Дерево понимания */
	baseContextUpdate()
	//fmt.Println("SET: ", p0Arr[0], p0Arr[1])
	return
}

func ClinerAllGomeoParams(value float64) {
	GomeostazParams[1] = value
	GomeostazParams[2] = 0.0
	GomeostazParams[3] = 0.0
	GomeostazParams[4] = 0.0
	GomeostazParams[5] = 0.0
	GomeostazParams[6] = 0.0
	GomeostazParams[7] = 0.0
	GomeostazParams[8] = 0.0
}

///////////////////////////////

// до определенной стадии развития Гон и Потребности в обучении не влияют ни на что (нет у детей гона)
var IsLevelBeginParam3 = false // true - уровень развития для Гона достигнут
var IsLevelBeginParam5 = false // true - уровень развития для Потребность в обучении достигнут

// корректировка жизненных параметров по каждому пульсу
func gomeostazUpdate() {
	if NotAllowSetGomeostazParams {
		return
	}

	// изменение со временем
	changingParVal(1) // !!! на действия затрачивается много энергии http://go/pages/terminal_actions.php
	changingParVal(2)
	// Гон начинает изменяться с некоторого уровня развития Beast
	if IsLevelBeginParam3 {
		changingParVal(3)
	}
	changingParVal(4)
	// Потребность в обучении начинает изменяться с некоторого уровня развития Beast
	if IsLevelBeginParam5 {
		changingParVal(5)
	}
	changingParVal(6)
	changingParVal(7)

	// эксклюзивные зависимости
	// повреждение
	// При энергии <5% начинает увеличивается cкорость поврежедения
	var oldValGomeostazParamsSpeed8 = GomeostazParamsSpeed[8]
	if GomeostazParams[1] < 5 {
		// при GomeostazParams[1]==0 прибавка скорости будет 0, при при GomeostazParams[1]==0 прибавка скорости станет 100
		GomeostazParamsSpeed[8] += 100 - int(GomeostazParams[1]*20.0)
	}
	// при повышенном стрессе:
	if GomeostazParams[2] > 70 {
		// при GomeostazParams[2]==70 прибавка скорости будет 0, при при GomeostazParams[2]==100 прибавка скорости станет 50
		GomeostazParamsSpeed[8] += int((GomeostazParams[2] - 70.0) * 1.7)
	}
	changingParVal(8)
	GomeostazParamsSpeed[8] = oldValGomeostazParamsSpeed8
	// При 100% повреждений  - смерть. Фон пульта менятся от повреждений, начиная с критического порога.
	// Смерть - черный фор с траурной надписью, сброс памяти.

	//самосохранение  начинает ухудшаться при превышении compareLimites[1] (Значительное отклонение)
	val1 := 0.0
	var limt = float64(compareLimites[1])
	if (limt - GomeostazParams[1]) > 0 {
		val1 = 1
	}
	val2 := 0.0
	if GomeostazParams[2] > 100-limt {
		val2 = 1
	}
	val3 := 0.0
	if GomeostazParams[8] > 100-limt {
		val3 = 1
	}
	add := val1 + val2 + val3

	if add > 1 {
		add = 1
	} // ограничить порцию изменения
	if add > 0 {
		GomeostazParams[7] += add
	} else { // если нет факторов риска, то уменьшать GomeostazParams[7]

		GomeostazParams[7] -= 5
	}

	if GomeostazParams[7] < 0 {
		GomeostazParams[7] = 0
	}
	if GomeostazParams[7] > 100 {
		GomeostazParams[7] = 100
	}

	if PulsCount > 1 && PulsCount%10 == 0 { // записать в файл текущее состояние гомеостаза раз в 10 сек
		SaveCurrentGomeoParams()
	}
	baseContextUpdate()

	/* Для параметров гомеостаза, напрямую не связанных с жизнеобеспечением (parMaxPulsCount: гон, потребность в общении,
	потребность в обучении и любопытство) организована цикличность: при нарастании параметра до максимума,
	он удерживается в течении 20 секунд, а потом сбрасывается.
	Это позволяет создавать достаточные по времени периоды специфических контекстов реагирования.
	*/
	if GomeostazParams[3] > 90 && parMaxPulsCount[0] == 0 {
		parMaxPulsCount[0] = PulsCount
	}
	if GomeostazParams[4] > 90 && parMaxPulsCount[1] == 0 {
		parMaxPulsCount[1] = PulsCount
	}
	if GomeostazParams[5] > 90 && parMaxPulsCount[2] == 0 {
		parMaxPulsCount[2] = PulsCount
	}
	if GomeostazParams[6] > 90 && parMaxPulsCount[3] == 0 {
		parMaxPulsCount[3] = PulsCount
	}
	// сброс после периода удержания
	if GomeostazParams[3] > 90 {
		if parMaxPulsCount[0]+PeriodPulsCount < PulsCount {
			parMaxPulsCount[0] = 0
			GomeostazParams[3] = 0
		}
	} else {
		parMaxPulsCount[0] = 0
	}

	if GomeostazParams[4] > 90 {
		if parMaxPulsCount[1]+PeriodPulsCount < PulsCount {
			parMaxPulsCount[1] = 0
			GomeostazParams[4] = 0
		}
	} else {
		parMaxPulsCount[1] = 0
	}

	if GomeostazParams[5] > 90 {
		if parMaxPulsCount[2]+PeriodPulsCount < PulsCount {
			parMaxPulsCount[2] = 0
			GomeostazParams[5] = 0
		}
	} else {
		parMaxPulsCount[2] = 0
	}

	if GomeostazParams[6] > 90 {
		if parMaxPulsCount[3]+PeriodPulsCount < PulsCount {
			parMaxPulsCount[3] = 0
			GomeostazParams[6] = 0
		}
	} else {
		parMaxPulsCount[3] = 0
	}
}

// Сохрнаить значения параметров гомеостаза в файле
func SaveCurrentGomeoParams() {
	if len(GomeostazParams) < 7 {
		return
	}
	var fStr = "1|" + strconv.Itoa(int(GomeostazParams[1])) + "\n" +
		"2|" + strconv.Itoa(int(GomeostazParams[2])) + "\n" +
		"3|" + strconv.Itoa(int(GomeostazParams[3])) + "\n" +
		"4|" + strconv.Itoa(int(GomeostazParams[4])) + "\n" +
		"5|" + strconv.Itoa(int(GomeostazParams[5])) + "\n" +
		"6|" + strconv.Itoa(int(GomeostazParams[6])) + "\n" +
		"7|" + strconv.Itoa(int(GomeostazParams[7])) + "\n" +
		"8|" + strconv.Itoa(int(GomeostazParams[8]))

	file, err := os.Create(lib.MainPathExeFile + "/memory_reflex/GomeostazParams.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	_, _ = file.WriteString(fStr)
}

// шаг изменения парамктра со скоростью GomeostazParamsSpeed
func changingParVal(id int) {
	if NotAllowSetGomeostazParams {
		return
	}
	step := float64(GomeostazParamsSpeed[id]) / 3600
	if id == 1 {
		ChangeGomeostazParametr(id, -step)
	} else {
		ChangeGomeostazParametr(id, step)
	}
}

var IsBeastDeath = false // true - смерть Beast при повреждении >99% gomeostas.IsBeastDeath
// true - смерть Beast при повреждении > 99%
func CheckBeastDeath() bool {
	if GomeostazParams[8] > 99.0 {
		IsBeastDeath = true
		return true
	}
	return false
}

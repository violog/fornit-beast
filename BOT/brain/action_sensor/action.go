/*
Ощущение действий с Пульта
Могут быть совершены сразу несколько действий.
Так что возникает образ действий (а далее - и общий образ действий и фраз)

*/

package action_sensor

import (
	"BOT/brain/gomeostas"
	"BOT/brain/transfer"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////
func init() {

}

/////////////////////////////////////////

var pulsCount = 0
var LifeTime = 0
var EvolushnStage = 0 // стадия развития
var IsStressing = false

////////////////////////////////////////////////////////////////////

// срабатывает с каждым пульсом
func ActionSensorPuls(evolushnStage int, lifeTime int, puls int, isSlipping bool) {
	LifeTime = lifeTime
	EvolushnStage = evolushnStage
	pulsCount = puls
	IsStressing = isSlipping
	CheckCurActions()
}

/*
	список активных действий с Пульта

Могут быть совершены сразу несколько действий, контекст которых удерживаются в течении 10 секунд.
В ActionFromPult сохраняется время активации действия в числе пульсов pulsCount

0 Нет никаких действий с Пульта
1 Непонятно
2 Понятно
3 Наказать
4 Поощрить
5 Накормить
6 Успокоить
7 Предложить поиграть
8 Предложить поучить
9 Игнорировать
10 Сделать больно
11 Сделать приятно
12 Заплакать
13 Засмеяться
14 Обрадоваться
15 Испугаться
16 Простить
17 Вылечить
*/
var ActionFromPult [18]int // время активации действия живет на время активации с пульта до завершения прохода дерева
// сохрянять текущий контекст ActionFromPultContext=ActionFromPult
var ActionFromPultContext [18]int // эти не очищаются

// название Базового контекста из его ID str:=action_sensor.GetActionNameFromID(id)
func GetActionNameFromID(id int) string {
	var out = ""
	switch id {
	case 1:
		out = "Непонятно"
	case 2:
		out = "Понятно"
	case 3:
		out = "Наказать"
	case 4:
		out = "Поощрить"
	case 5:
		out = "Накормить"
	case 6:
		out = "Успокоить"
	case 7:
		out = "Предложить поиграть"
	case 8:
		out = "Предложить поучить"
	case 9:
		out = "Игнорировать"
	case 10:
		out = "Сделать больно"
	case 11:
		out = "Сделать приятно"
	case 12:
		out = "Заплакать"
	case 13:
		out = "Засмеяться"
	case 14:
		out = "Обрадоваться"
	case 15:
		out = "Испугаться"
	case 16:
		out = "Простить"
	case 17:
		out = "Вылечить"
	}
	return out
}

//var IsPress3or4button = false /*нажата кнопка Наказать или Поощрить - пока под вопросом, но используется*/

//////////////////////////////////////////////////////////////

// учесть мотивирующий эффект Настроения сообщения Оператора -
func UpdateMoodEffectFromMessage(moodID int) {
	gomeostas.SetMoodePraseEffectArr(moodID)
}

///////////////////////////////////////////////////////////

/*
	акция с кнопок пульта

сопровождается гомеостатическими эффектами GomeostazActionEffectArr =make(map[int]string)
*/
func SetActionFromPult(actionList string, energi int) {
	// очищать прежние акции с пульта перед установкой новых т.к. сочетания передаются сразу.
	DeactivationTriggers()
	DeactivationTriggersContext()

	actArr := strings.Split(actionList, "|")
	//var listID []int
	for n := 0; n < len(actArr); n++ {
		if len(actArr[n]) == 0 {
			continue
		}
		actionID, _ := strconv.Atoi(actArr[n])
		/*
			при нажатии на кнопки включается и выключается игровой режим.
			Но он НЕ может включаться по инициативе бота
			и поэтому на пульте /pult_actions.php в function to_action(id
			жстко переключения привязаны к кнопкам для индикации игрового режима.

			//Непонятно, поучить, Сделать больно,Сделать приятно,Заплакать,Испугаться,Вылечить
			if actionID == 1 || actionID == 8 || actionID == 10 || actionID == 11  || actionID == 12 || actionID == 15 || actionID == 17{
				transfer.IsPsychicGameMode = false  // Выключить игровой режим
			}*/

		if actionID == 6 { //  Успокоить
			// Снять режимы Плохо и Хорошо
			//			gomeostas.CliarWellBad()  нажатием ESC или по крестику у плашки Плохо-Хорошо
		}

		if actionID == 7 { //  Поиграть
			transfer.IsPsychicGameMode = true // Включить игрового режима
		}

		if actionID == 8 { //  Поучить
			transfer.IsPsychicТeachingMode = true // Включить учителского режима
		}

		if actionID == 5 { //  Накормить
			switch energi {
			case 1:
				gomeostas.ChangeGomeostazParametr(1, 20.0)
			case 2:
				gomeostas.ChangeGomeostazParametr(1, 50.0)
			case 3:
				gomeostas.ChangeGomeostazParametr(1, 80.0)
			}
		} else {
			/*
				if actionID == 3 || actionID == 4 {
					IsPress3or4button = true // нажата кнопка Наказать или Поощрить
				} else {
					IsPress3or4button = false
				}*/

			if !transfer.IsPsychicGameMode { // не воздействовать на гомео-параметры в игровом режиме
				ge := gomeostas.GomeostazActionEffectArr[actionID]
				if len(ge) > 0 { // пример: 2>40,4>50,5>50,6>30,7>-20
					aeArr := strings.Split(ge, ",")
					for i := 0; i < len(aeArr); i++ {
						p := strings.Split(aeArr[i], ">")
						id, _ := strconv.Atoi(p[0])
						v, _ := strconv.ParseFloat(p[1], 64)
						gomeostas.ChangeGomeostazParametr(id, v)
					}
				}
			}
		}
		ActionFromPult[0] = 0
		ActionFromPult[actionID] = pulsCount
		ActionFromPultContext[0] = 0
		ActionFromPultContext[actionID] = pulsCount

		// эффект нажатия кнопок действий на пульте
		gomeostas.SetGomeostazActionCommonEffectArr(actionID)
	}
	// дезактивировать все акции через 10 секунд
	/* вредна, т.к. запущенный таймер может погасить нормальные акции.
	Акции должны гаситься после текущего использования в perception.go -> action_sensor.DeactivationTriggers()
	time.AfterFunc(10*time.Second, func() {
		DeactivationTriggers()
	})
	*/
	// тестирование
	// gomeostas.BetterOrWorseNow()
	return
}

// дезактивировать все пусковые стимулы при изменении условий action_sensor.DeactivationTriggers() ()
func DeactivationTriggers() {
	for i := 1; i < 18; i++ {
		ActionFromPult[i] = 0
	}
	//IsPress3or4button = false
}

// дезактивировать все контексты
func DeactivationTriggersContext() {
	for i := 1; i < 18; i++ {
		ActionFromPultContext[i] = 0
	}
}

// какие акции действуют в данный момент действий с пульта - активные контексты действий с Пульта Очищаются
func CheckCurActions() []int {
	var aArr []int
	for i := 1; i < 18; i++ {
		if ActionFromPult[i] > 0 {
			aArr = append(aArr, i)
		}
	}
	if aArr == nil {
		ActionFromPult[0] = 0
	}
	return aArr
}

// какие сохраненные акции действуют в данный момент пульса - активные контексты действий с Пульта НЕ очищаются, сохраняя контекст
func CheckCurActionsContext() []int {
	var aArr []int
	for i := 1; i < 18; i++ {
		if ActionFromPultContext[i] > 0 {
			aArr = append(aArr, i)
		}
	}
	if aArr == nil {
		ActionFromPultContext[0] = 0
	}
	return aArr
}

/////////////////////////////////////////////////////////////////////

/*
была акция с пульта - получить время акции
*/
var PultActionPulsCount = 0 // время акции
func IsPultAction() {
	PultActionPulsCount = pulsCount
}

//////////////////////////////////////////////////

/* общие глобальные дела */
package brain

import (
	"BOT/brain/action_sensor"
	"BOT/brain/gomeostas"
	"BOT/brain/psychic"
	"BOT/brain/reflexes"
	"BOT/brain/sleep"
	TerminalActions "BOT/brain/terminete_action"
)

var IsStressing = 0 // не успевает за 1 сек отсканировать
var isWorking = 0   // флаг блока обработки нейросети

/*
	Главный цикл (пульс) опроса состояния Beast и его ответов

активируется в pulsActions() каждый пульс
*/
func Activation(lifeTime int) {

	// true - смерть Beast при повреждении >99%
	if gomeostas.IsBeastDeath { // прекращение раздачи пульса
		return
	}
	if isWorking == 1 {
		IsStressing = 1
		return
	} else {
		IsStressing = 0
	}
	// начало обработки нейросети
	isWorking = 1
	isSlipping, sleepingType := sleep.GetSleepCondition()
	// текущее состояние гомеостаза и базового контекста с каждым пульсом
	gomeostas.GomeostazPuls(EvolushnStage, lifeTime, PulsCount, isSlipping)
	action_sensor.ActionSensorPuls(EvolushnStage, lifeTime, PulsCount, isSlipping)
	sleep.SleepPuls(EvolushnStage, lifeTime, PulsCount)
	reflexes.ReflexCountPuls(EvolushnStage, lifeTime, PulsCount, isSlipping)
	TerminalActions.TermineteActionCountPuls(EvolushnStage, lifeTime, PulsCount, isSlipping)
	psychic.PsychicCountPuls(EvolushnStage, lifeTime, PulsCount, sleepingType)
	psychic.SetNotAllowAnyActions(NotAllowAnyActions)
	// конец обработки нейросети
	isWorking = 0
}

////////////////////////////////////////////////////

/////////////////////////////////////////

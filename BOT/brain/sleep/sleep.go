/* Cон, его стадии и циклы

_____________________________________________________
В каждом пакете есть флаг  и там для сна выполняется
	if IsSleeping {
		sleepingProcess()
	}
в психике для этого есть psychic_sleep_process.go
_____________________________________________________

начать сон sleep.BeginSleepСondition()

!!!!!!!!
TODO нужно где-то определить функцию засыпания чтобы начать сон sleep.BeginSleepСondition()
и накапливать признаки необходимости сна нужно,
начиная с необходиомсти обработки накопившегося массива распознанных фраз чтобы во сне prepareWordArr()
!!!!!!!!
*/

package sleep

import (
	"BOT/brain/gomeostas"
	"BOT/lib"
)

// ////////////////////////////////////////
var SlipPulsCount = 0 // передача тика Пульса из brine.go
var LifeTime = 0
var EvolushnStage = 0 // стадия развития

var sleepСonditionPulsCount = 0 // начат пульс во время сна

var isStageDreams = false // это стадия сновидений

// коррекция текущего состояния с каждым пульсом
func SleepPuls(evolushnStage int, lifeTime int, puls int) {
	isStageDreams = false
	lib.BlockingAnyActions = false

	LifeTime = lifeTime
	EvolushnStage = evolushnStage
	SlipPulsCount = puls // передача номера тика из более низкоуровневого пакета

	// разбудить при включении. Сторожевая функция - в Пульсе рефлексов вызывает sleep.WakeUpping()
	if SlipPulsCount > 3 {
		WakeUpping() // разбудить при включении
	}

	if SlipPulsCount > 5 {
		sleepNecessityDetector()
	}

	// ход процесса сна
	if IsSleeping {
		sleepСonditionPulsCount++

		// просыпание из-за чего-то - в guarding_sleep_center.go

		prepareWordArr()     // обработка накопившегося массива распознанных фраз
		prepareAnyArr()      //
		isStageDreams = true // начать стадию сновидений - после первичной обработки

		/* последовательность циклов сна по пульсу
		После основной очистки начинается процесс сновидения с оптимизацией эпиз.памяти и заполнением мот.Правил.
		TODO:
		*/
		if isStageDreams {
			lib.BlockingAnyActions = true // блокировать выдачу на пульт любых действий
			//	psychic.GotoDreaming()  ЗДЕСЬ ИСПОЛЬЗОВАТЬ ФУНКЦИИ пассивного режима, а не включать сам режим!

			isStageDreams = false // закончился цикл сновидений в психике, можно прервать

		}
		// можно просыпаться, все, что нужно во сне сделано
		if sleepNecessityValue < 10 {
			WakeUpping()
			return
		}
	}
	////////////////////////////////////////////

	// можно ли заснуть
	if !IsSleeping {
		if sleepNecessityValue > 60 {
			if IsPossibleToSleep() {
				BeginSleepСondition()
			}
		}
	}

	// понижение повреждений во сне каждый пульс или в час 0.01*3600=36
	if !gomeostas.NotAllowSetGomeostazParams {
		gomeostas.GomeostazParams[8] -= 0.01
		if gomeostas.GomeostazParams[8] < 0 {
			gomeostas.GomeostazParams[8] = 0
		}
	}

}

//////////////////////////////////////////////////////

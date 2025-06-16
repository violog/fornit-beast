/* основыне данные по рефлексам

 */

package reflexes

import (
	"BOT/brain/gomeostas"
	"BOT/brain/psychic"
	"BOT/brain/transfer"
)

// ////////////////////////////////////////////////////
// отслеживане времени с последнего изменения условий с Пульта в пульсах
var lastActivnostFromPult = 0

// было изменение активности с пульта в текущем пульсе. Только одна активность допускается в течение пульса.
var activetedPulsCount = 0 // против многократных срабатываний

// переключатель игрового режима  reflexes.IsGameMode
//var IsGameMode=false

// true - была активность с Пульта в perception.go
var IsPultActionThisPuls = false

// ПУЛЬС рефлексов
var ReflexPulsCount = 0 // передача тика Пульса из brine.go
var LifeTime = 0        // время жизни
var EvolushnStage = 0   // стадия развития
var IsSlipping = false  // флаг фазы сна

// коррекция текущего состояния гомеостаза и базового контекста с каждым пульсом
func ReflexCountPuls(evolushnStage int, lifeTime int, puls int, isSlipping bool) {
	IsPultActionThisPuls = false
	LifeTime = lifeTime
	EvolushnStage = evolushnStage
	ReflexPulsCount = puls // передача номера тика из более низкоуровневого пакета
	IsSlipping = isSlipping

	if puls == 4 {
		psychic.PsychicInit() // после 3-го пульса!
	}
	if puls == 5 {
		testingRunMakeAutomatizmsFromReflexes()
	}

	if activetedPulsCount != ReflexPulsCount { // защита от повторных срабатываний
		if gomeostas.IsNewConditions { // изменились условия
			ActiveFromConditionChange()
			lastActivnostFromPult = ReflexPulsCount
			return
		}
		/* если условия не меняются более 20 сек, то пусть срабатывает простейший инстинкт
		только если Базоваое состояние Плохо или Хорошо
		*/
		if ReflexPulsCount-lastActivnostFromPult > 20 {
			bc := gomeostas.CommonBadNormalWell
			if bc != 2 {
				// найти и выполнить простейший безусловный рефлекс
				findAndExecuteSimpeReflex()
			}
			if EvolushnStage == 1 {
				transfer.IsPsychicGameMode = false //автоматический сброс игрового режима для первой стадии, автоматическая активация делается в activeGameMode()
			}
			lastActivnostFromPult = ReflexPulsCount // новый период ослеживания
		}
	}

	// обнулить причину возможного запуска рефлекса
	if oldActiveCurTriggerStimulsID > 0 && oldActiveCurTriggerStimulsPulsCount > (ReflexPulsCount+10) || transfer.IsPsychicGameMode == false {
		oldActiveCurTriggerStimulsID = 0
	}

}

////////////////////////////////////////////////
/////////////////////////////////////////////////

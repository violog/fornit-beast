/* Настроение: сила Плохо -10 ... 0 ...+10 Хорошо.
Ощущение силы Плохо и сил Хорошо и это информационно для оценок последствий действий.
Настроение в зависимости от гомеостатический параметров
а так же оченка опасноного состояния
*/

package psychic

import (
	actionSensor "BOT/brain/action_sensor"
	"BOT/brain/gomeostas"
	"BOT/lib"
)

////////////////////////////////////////
/* CurrentMood - ТЕКУЩАЯ ОЦЕНКА ИЗМЕНЕНИЯ гомео-СОСТОЯНИЯ (гомео-настроения) - сила ощущаемого настроения PsyBaseMood
-10 0 +10 изменение настроения
- основа для стремления и избегания.
*/
var CurrentMood = 0 // Так же текущее значение - в CurrentInformationEnvironment.Mood
// предыдуще значение. Так же предыдущее значение - в OldInformationEnvironment.Mood
var PreviousMood = 0

/*
	Определячет первый уровень образа Цели.

Субъективно ощущаемая оценка, текущее осознаваемое настроение, которое можно произвольно изменять.
Она стремится к нулю со временем.
Постоянное состояние Хорошо довольно быстро уходит, постоянное состояние Плохо уходит гораздо медленнее.
При резких изменениях возникает эффект "маятника настроения":
появление противоположного по знаку настроения, но меньшего значения.
Значение обновляется при значительных изменениях (CurrentMood - PreviousMood)
Предыдущее значение - в OldInformationEnvironment.PsyMood
*/
var PsyMood = 0             // плохо -10...0...10 хорошо - сила ощущаемого настроения PsyBaseMood  корректируется с каждым пульсом moodePulse()
var PsyMoodPulse = 0        // последнее изменение PsyMood в пульсах
var veryMachChanged = false // true - было сильное изменение настроения
// ///////////////////////////////////////////////////////////////////////////////////
var PsyBaseMood = 0 // -1 Плохое настроение, 0 Нормальное, 1 - хорошее настроение

/////////////////////////   БОЛЬ И РАДОСТЬ    //////////////////////////////
/* Основа этих величин - значение, получаемые в http://go/pages/gomeostaz.php
в таблицах "Активности Базовых стилей" и "Действия оператора - гомеостатический эффект".
Боль и Радость на доосознательном уровне НЕ имеют рефлекторное назначение
и используются ТОЛЬКО для психики в определении последствий реакции.

В каких животных появляется ощущение боли в ходе эволюционных усложнений?
Учитывая, что ощущение боли не возникает на уровне рефлексов, это психический феномен.
Поэтому и возникает вопрос: чтобы боль могла ощущаться, нужны какие-то рецепторы,
которые возникли задолго до появления психики.
В каких животных есть такие рецепторы, но еще нет ощущения боли?

Вопрос является сложным и до конца неисследованным.
Однако, исследования показывают, что рецепторы боли (нервные окончания,
способные реагировать на повреждения) могут существовать у разных видов животных,
даже если у них нет развитой психики и способности к осознанию боли.
Это объясняется тем, что ощущение боли является эволюционно выгодным механизмом защиты
от потенциальных угроз и повреждений.
Рецепторы боли могут помочь животному избегать опасных ситуаций и предотвращать повторные повреждения.
Некоторые исследования показывают, что рецепторы боли найдены у различных групп животных,
таких как беспозвоночные (например, раки и сепии), рыбы и даже некоторые примитивные бесхвостые хордовые
(например, амфиоксусы). Однако, понимание того,
как именно эти животные воспринимают или осознают болевые сигналы, остается предметом дальнейших исследований.

Так что на допсихическом уровне детекция Боли и Радости осуществляется в gomeostas.go в виде:
var GomeostazActionEffectPainV=0 // величина Боли
var GomeostazActionEffectJoyV=0 // величина радости
Но только безусловные рефлексы (и условные, созданные на их основе) могут использовать эти детекторы.
Однако, у же на том уровне есть детектор СтелоЛучшеИлиХуже (func BetterOrWorseNow()),
который используется в психике для оценки эффекта ответного реагирования.
На уровне рефлекснв нет Боли, а есть только детекторы травматических последствий.
Боль и радость: fornit.ru/67646
TODO Их мы могли бы использовать в безусловных рефлексов и это стоит сделать.
*/
var painValue = 0 // величина Боли
var joyValue = 0  // величина радости

///////////////////////////////////////////////////////////////////////////////////////

/*
	ПУЛЬС психики

var PulsCount=0 // передача тика Пульса из brine.go
var LifeTime=0
var EvolushnStage=0 // стадия развития
var IsSleeping=false
*/
func moodePulse() {
	updatePsyMood(false)
}

/*
	Обновляется при инициализации GetCurrentInformationEnvironment() до и после совершения действий

Это безусловно-рефлекторная оценка на освнове изменений
1) гомео параметров (улавливает их изменение >4%) и 2) основы базовых контекстов: 0- НОРМА, 1-ПЛОХО, 2-ХОРОШО
Хотя базовые контексты уже создаются на основе жизненных параметров,
но это - безусловный рефлекс на основе безусловного рефлекса активации базовых контекстов.
На этой основе происходит конкуренция Целей при их выборе
и, как следствие, произвольная оценка эмоционального состояния.
*/
func GetCurMood() int {
	danger := GetAttentionDanger()
	var mood = 0
	GomeostazParams := gomeostas.GomeostazParams
	OldGomeostazParams := gomeostas.OldGomeostazParams
	CommonBadNormalWell := gomeostas.CommonBadNormalWell
	OldCommonOldBadValue := gomeostas.OldCommonBadNormalWellVal

	for n := 1; n < 9; n++ {
		if OldGomeostazParams[n] != 0 {
			// процент изменения параметра от предыдущего значения позитивный или негативный:
			var proc = 100 * (GomeostazParams[n] - OldGomeostazParams[n]) / OldGomeostazParams[n]

			switch n {
			case 1: // энергия   +/- 10
				if proc > 4 {
					if danger {
						mood += 6
					} else {
						mood += 3
					}
				}
				if proc < -4 {
					if danger {
						mood -= 6
					} else {
						mood -= 3
					}
				}
			case 2: // стресс  +/- 20
				if proc > 6 {
					if danger {
						mood -= 4
					} else {
						mood -= 2
					}
				}
				if proc < -6 {
					if danger {
						mood += 4
					} else {
						mood += 2
					}
				}
			case 3: // гон
				if proc > 20 {
					if danger {
						mood -= 1
					} else {
						mood -= 1
					}
				}
				if proc < -20 {
					if danger {
						mood += 1
					} else {
						mood += 1
					}
				}
			case 4: // потребность в общении +/- 20
				if proc > 10 {
					if danger {
						mood -= 1
					} else {
						mood -= 1
					}
				}
				if proc < -10 {
					if danger {
						mood += 1
					} else {
						mood += 1
					}
				}
			case 5: // потребность в обучении +/- 20
				if proc > 10 {
					if danger {
						mood -= 1
					} else {
						mood -= 1
					}
				}
				if proc < -10 {
					if danger {
						mood += 1
					} else {
						mood += 1
					}
				}
			case 6: // 	Поиск любопытство  +/- 20
				if proc > 10 {
					if danger {
						mood += 1
					} else {
						mood += 1
					}
				}
				if proc < -10 {
					if danger {
						mood -= 1
					} else {
						mood -= 1
					}
				}
			case 7: // Самосохранение (Жадность, эгоизм, самозащита, страх смерти. Зависит от ситуации, может сам уменьшаться при благополучии.)
				if proc > 20 {
					if danger {
						mood -= 1
					} else {
						mood -= 1
					}
				}
				if proc < -20 {
					if danger {
						mood += 1
					} else {
						mood += 1
					}
				}
			case 8: // Повреждения
				if proc > 10 {
					if danger {
						mood -= 1
					}
				} else {
					if proc > 20 {
						if danger {
							mood -= 5
						}
					}
				}
				if proc > 40 {
					// при proc==40 mood=-6 а при proc==100 mood=-10
					mood = -int(10.0 - (100.0-proc)/15.0)
				}
			} //switch n
		}
		OldGomeostazParams[n] = GomeostazParams[n]
	}

	switch CommonBadNormalWell {
	case 2: // НОРМА
		if OldCommonOldBadValue == 1 { // было ПЛОХО
			if danger {
				mood += 5
			} else {
				mood += 2
			}
		}
		if OldCommonOldBadValue == 3 { // было ХОРОШО
			if danger {
				mood -= 5
			} else {
				mood -= 2
			}
		}
	case 1: // ПЛОХО
		if OldCommonOldBadValue == 3 { // было ХОРОШО
			if danger {
				mood -= 5
			} else {
				mood -= 2
			}
		}
		if OldCommonOldBadValue == 2 { // было НОРМА
			if danger {
				mood -= 3
			} else {
				mood -= 1
			}
		}
	case 3: // ХОРОШО
		if OldCommonOldBadValue == 2 { // было НОРМА
			if danger {
				mood += 3
			} else {
				mood += 1
			}
		}
		if OldCommonOldBadValue == 1 { // было ПЛОХО
			if danger {
				mood += 6
			} else {
				mood += 3
			}
		}
	}

	// при значениях Боли или Радости >5 они определяют настроение
	moorReduce := mood / 3
	painValue, joyValue = gomeostas.GetCurPainJoy()
	if painValue > 5 {
		mood = moorReduce - painValue
	}
	if joyValue > 5 {
		mood = moorReduce + joyValue
	}

	if mood > 10 {
		mood = 10
	}
	if mood < -10 {
		mood = -10
	}

	PreviousMood = CurrentMood
	CurrentMood = mood

	updatePsyMood(true)
	updatePsyBaseMood()

	return mood
}

///////////////////////////////////////////////////////////////////

/* оценить опасность текущей ситуации: да-нет
 */
func GetAttentionDanger() bool {
	// опасные жизненные параметры
	for k, pID := range gomeostas.BadNormalWell {
		if pID == 1 { // плохо для данного параметра гомеостаза
			if k == 1 || k == 2 || k == 7 || k == 8 {
				return true
			}
		}
	}
	// по опасному действию с Пульта

	aArr := actionSensor.CheckCurActionsContext()
	for i := 0; i < len(aArr); i++ {
		if aArr[i] == 3 || aArr[i] == 10 || aArr[i] == 12 || aArr[i] == 15 {
			return true
		}
	}
	/*
		// по опасной фразе с Пульта
		if isDangerWordFromPult(){
			return true
		}
	*/
	return false
}

/*
	PsyMood стремится к нулю со временем.

Постоянное состояние Хорошо довольно быстро уходит, постоянное состояние Плохо уходит гораздо медленнее.
При резких изменениях возникает эффект "маятника настроения":
появление противоположного по знаку настроения, но меньшего значения.
*/
func updatePsyMood(newMood bool) {
	if PreviousMood == 0 { // первое значение - просто копируем
		PsyMood = CurrentMood
		PsyMoodPulse = PulsCount
		return
	}
	// было ли очень сильное изменение - для маятника настроения
	if (lib.Abs(PsyMood-CurrentMood) > 6) || (lib.Abs(CurrentMood) >= 3 && lib.IsDiffersOfSign(PsyMood, CurrentMood)) {
		veryMachChanged = true
	}
	// PsyMood изменяется при значительных измнениях CurrentMood
	if lib.Abs(PsyMood-CurrentMood) > 3 {
		PsyMood = CurrentMood
		PsyMoodPulse = PulsCount
		return
	}
	// PsyMood изменяется при смене знака значения (одно положительное, другое отрицательное),
	// большего, чем 1 (чтобы не было постоянной смены около нуля)
	if lib.Abs(CurrentMood) >= 1 && lib.IsDiffersOfSign(PsyMood, CurrentMood) {
		PsyMood = CurrentMood
		PsyMoodPulse = PulsCount
		return
	}
	// маятник настроения - через 30 пульсов
	if veryMachChanged && (PulsCount-PsyMoodPulse > 30) {
		PsyMood = -(PsyMood - 3) // смена настроение на противоположное, но уменьшенное значение
		PsyMoodPulse = PulsCount
		veryMachChanged = false
		return
	}
	////////// режим постепенного угасания
	if PsyMood > 0 { // хорошее настроение угасает довольно быстро
		if PulsCount-PsyMoodPulse > 30 {
			if lib.Abs(CurrentMood) > 0 {
				if PsyMood > 0 {
					PsyMood--
				} else {
					PsyMood++
				}
				PsyMood = CurrentMood
				PsyMoodPulse = PulsCount
			}
		}
	}
	if PsyMood < 0 { // плохое настроение более важно, угасает медленнее
		if PulsCount-PsyMoodPulse > 60 {
			if lib.Abs(CurrentMood) > 0 {
				if PsyMood > 0 {
					PsyMood--
				} else {
					PsyMood++
				}
				PsyMood = CurrentMood
				PsyMoodPulse = PulsCount
			}
		}
	}
}

/*
	ощущаемое настроение (-1,0,1)

обновляется сразу после updatePsyMood
*/
func updatePsyBaseMood() {
	if PsyMood <= 1 {
		PsyBaseMood = -1
	}
	if PsyMood < 2 && PsyMood > -2 {
		PsyBaseMood = 0
	}
	if PsyMood >= 1 {
		PsyBaseMood = 1
	}
}

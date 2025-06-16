/* Переактивация контекстов не по гомео-событиям, а по стимулам, т.е. рефлекторная переативация.

Со стимулами действия уже связана значимость в виде боли и радости и для таких используется func ContextActiveFromStimul

На уровне психики приобретается опыт связывания значимостей с образами восприятия,
что детектируется func getExtremImportanceObject() при активации дерева автоматизмов
и если эта значимость высока, то вызывается func needPsyReactiveFromStimul

В обоих функциях есть возможность задавать стеепнь консервативности отклика
(с какого повторения сильного воздействия начнет происходить перективация).
*/

package gomeostas

import "BOT/lib"

///////////////////////////////////////////////////////////////////

/*
	Рефлекторная переактивация контекстов в зависимости от эффекта стимула, сопровождающихся болью и радостью.

effect<0 - значение боли, effect>0 - значение радости
активируется с приходом стимула-действия
Срабаывает до активации дерева рефлексов и до корректировки контекстов действием,
т.е. происходит рефлексторная переактивация.
*/

func ContextActiveFromStimul(effect int) {

	if noReactivationFromStimuls { // отключить переактивацию стимулом
		return
	}

	if !needReactiveFromStimul(effect) { // пока не реагируем
		return
	}

	reflexReactation(effect)
}

// консервативность реагирования на действие с пульта
var oldStymulEffect = 0
var oldStymulEffectCount = 0     // сколько раз повторялся тот же эффект
var conservatismStymulEffect = 2 //степень консервативности: сколько раз игнорируем
func needReactiveFromStimul(effect int) bool {
	if (lib.Abs(oldStymulEffect) < 6) || !lib.EqualSigns(oldStymulEffect, effect) { // менять контексты то0лько при повторном экстремальном стимуле
		oldStymulEffectCount = 0
		oldStymulEffect = effect
		return false
	}
	if lib.EqualSigns(oldStymulEffect, effect) {
		oldStymulEffectCount++
	}
	if oldStymulEffectCount >= conservatismStymulEffect { // нужно реагировать
		return true
	}
	oldStymulEffect = effect

	return false
}

///////////////////////////////////////////////////

//////////////////////////////////////////////////
/* Произвольная, но еще не осознанная! переактивация контекстов, например, оцененной значимостью образа.
т.е. только на уровне опыта значимости такое возможно (extremImportanceObject)
*/
func ContextActiveFromPsy(effect int) bool {
	if noReactivationFromStimuls { // отключить переактивацию
		return false
	}
	if !needPsyReactiveFromStimul(effect) { // пока не реагируем
		return false
	}

	reflexReactation(effect)
	return true
}

// консервативность реагирования на действие с пульта
var oldPsyStymulEffect = 0
var oldPsyStymulEffectCount = 0     // сколько раз повторялся тот же эффект
var conservatismPsyStymulEffect = 4 //степень консервативности: сколько раз игнорируем
func needPsyReactiveFromStimul(effect int) bool {
	if (lib.Abs(oldPsyStymulEffect) < 6) || !lib.EqualSigns(oldPsyStymulEffect, effect) { // менять контексты то0лько при повторном экстремальном стимуле
		oldPsyStymulEffectCount = 0
		oldPsyStymulEffect = effect
		return false
	}
	if lib.EqualSigns(oldPsyStymulEffect, effect) {
		oldPsyStymulEffectCount++
	}
	if oldPsyStymulEffectCount >= conservatismPsyStymulEffect { // нужно реагировать
		return true
	}
	oldPsyStymulEffect = effect

	return false
}

// ////////////////////////////////////////////////
//
//	переативировать
//
// ////////////////////////////////////////////////
// выполнение переактивации
func reflexReactation(effect int) {
	//При срабатывании от сильной боли:
	if effect < -5 {
		isStrah := false // страх
		isZlost := false // злость
		for i := 0; i < len(BaseContextActive); i++ {
			if BaseContextActive[i] && i == 8 {
				isStrah = true
			}
			if BaseContextActive[i] && i == 10 {
				isZlost = true
			}
			// сбросить все контексты
			BaseContextActive[i] = false
		}
		if !isStrah && !isZlost {
			// установить поиск,защита,страх - легальное сочетание со страницы редактора рефлексов
			BaseContextActive[2] = true
			BaseContextActive[5] = true
			BaseContextActive[8] = true
			setNewContexts() // заполнить CurStyleImage новыми значениями по порядку весов стилей
			keepingContextPulsCount = PulsCount
			return
		}
		if isZlost { // уже злится, вогнать в ступор - легальное сочетание со страницы редактора рефлексов
			BaseContextActive[7] = true
			setNewContexts() // заполнить CurStyleImage новыми значениями по порядку весов стилей
			keepingContextPulsCount = PulsCount
			return
		} else {
			// установить поиск,агрессия,злость - легальное сочетание со страницы редактора рефлексов
			BaseContextActive[2] = true
			BaseContextActive[9] = true
			BaseContextActive[10] = true
			setNewContexts() // заполнить CurStyleImage новыми значениями по порядку весов стилей
			keepingContextPulsCount = PulsCount
			return
		}
	} //if effect<-5{
	//При срабатывании от сильной радости:
	if effect > 5 {
		isDobro := false // доброта
		isLen := false   // лень
		for i := 0; i < len(BaseContextActive); i++ {
			if BaseContextActive[i] && i == 11 {
				isDobro = true
			}
			if BaseContextActive[i] && i == 6 {
				isLen = true
			}
			// сбросить все контексты
			BaseContextActive[i] = false
		}
		if !isDobro && !isLen {
			// установить Доброта - легальное сочетание со страницы редактора рефлексов
			BaseContextActive[11] = true
			setNewContexts() // заполнить CurStyleImage новыми значениями по порядку весов стилей
			keepingContextPulsCount = PulsCount
			return
		}
		if !isLen {
			// установить Лень,Доброта - легальное сочетание со страницы редактора рефлексов
			BaseContextActive[6] = true
			BaseContextActive[11] = true
			setNewContexts() // заполнить CurStyleImage новыми значениями по порядку весов стилей
			keepingContextPulsCount = PulsCount
			return
		}
	} //if effect > 5 {
}

///////////////////////////////////////////////

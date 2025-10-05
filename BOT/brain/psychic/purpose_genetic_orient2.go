/*
Для ориентировочного рефлекса типа 2
функции для определения Цели в данной ситуации - на уровне наследственных функций
исходя из текущей информационной среды CurrentInformationEnvironment:

*/

package psychic

import (
	"BOT/brain/gomeostas"
)

///////////////////////////////////////////////
// обработка автоматизма, рвущегося на выполнение, но в условиях может быть новизна news
/* Здесь - очень органиченные возможности адаптации автоматизма:
плохой - не выполнять, хороший - выполнять
при опасной ситуации выполнять тот, какой есть,
при спокойной ситуации - пробовать рефлексы мозжечка.

Из-за столь скудных возможностей и разросся функционал мыслительных автоматизмов
с их произвольностью (- перекрытием имеющихся автоматизмов новыми).
*/
var oldNodeAutomatism = 0 // прошлы раз запускался такой штатный автоматизм
func getPurposeGenetic2AndRunAutomatism(atmtzmID int) *Automatism {

	//atmzm:= AutomatismFromId[atmtzmID]
	atmzm, ok := ReadeAutomatismFromId(atmtzmID)
	if !ok {
		return nil
	}

	// Определение Цели в данной ситуации - на уровне наследственных функций
	purpose := getPurposeGenetic()
	// мозжечковые рефлексы - самый первый уровень осознания - подгонка действий под заданную Цель.

	// есть ли очень значимые новые признаки?
	newsRes := getImportantSigns()
	if newsRes { // повышенная опасность от оператора
		// срочность и важность ситуации: если очень срочно и важно - просто оставить имеющийся автоматизм
		runAutomatismFromPurpose(atmzm, purpose)
		return atmzm
	}

	if purpose.veryActual { // нужно ли вообще шевелиться?
		// высокий purpose.veryActual, 	нужно выполнить штатный автоматизм, а не придуманный

		// если в прошлый раз уже был такой автоматизм, то ничего не делать, иначе начинает повторять одно и то же
		if oldNodeAutomatism != atmzm.ID {
			runAutomatismFromPurpose(atmzm, purpose)
			oldNodeAutomatism = atmzm.ID
			return atmzm
		}

		// список всех автоматизмов для ID узла Дерева
		//aArr:=GetMotorsAutomatismListFromTreeId(detectedActiveLastNodID)

		/*aID := getAutomatismFromNodeID(detectedActiveLastNodID)
		AutomatismFromIdMapCheck()
		atmzm=AutomatismFromId[aID]
		purpose.actionID=ActionsImageArr[atmzm.ActionsImageID]
		runAutomatismFromPurpose(atmzm, purpose)*/
		return nil

		//if purpose.veryActual
	} else { // нет опасности и нет опасной новизны

		/*/ плохой автоматизм,
		if atmzm.Usefulness == -1 {	// сделать пользу ==0 и запустить с повышенным уровнем
			// была ли уже оптимизация?
			if cerebellumCoordination(atmzm, 0) {
				atmzm.Usefulness=0 // чтобы не блокировался
				runAutomatismFromPurpose(atmzm, purpose)
				return atmzm
			} else {
				if gomeostas.BaseContextActive[2] || gomeostas.BaseContextActive[3] { // активен Поиск или Игра
					// тупо метод тыка
					// Тупо поэкспериментировать для пополнения опыта (не)удачных автоматизмов
					// TODO !не проверено!
					// в отличии от createAndRunAutomatismFromPurpose(purpose) не использовать текущие рефлексы, а пробовать всякое
					// Выдавая это на стадии 3, тварь получает реакцию оператора, которую отзеркаливает
					atmzm := findAnySympleRandActions()
					return atmzm
				} else { // НЕ ИГРА  И НЕ ПОИСК, плохой автоматизм просто не выполнять
					return nil
				}
			}
		}*/
		if atmzm.Usefulness < 0 { // совсем плохой автоматизм
			if gomeostas.BaseContextActive[2] || gomeostas.BaseContextActive[3] { // активен Поиск или Игра
				atmzm := findAnySympleRandActions()
				return atmzm
			}
		}

		//все нормально, просто выполнить автоматизм и отслеживать последствия
		runAutomatismFromPurpose(atmzm, purpose)
		return atmzm
	}

	return nil
}

/////////////////////////////////////////////////////////////////

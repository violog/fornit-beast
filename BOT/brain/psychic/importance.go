/* Значимость элементов восприятия - как объекта произвольного внимания:
того из всего воспринимаемого, что имеет наибольшую значимость
т.к. именно наибольшая значимость должна осмысливаться.

Кроме того, значимости объектов - это и есть модель понимания данного объекта внимания -
его значимость в разных условиях и то, какие действия могут быть совершены при этом.

Значимость формируется в эпиз.памяти (func saveNewEpisodic) и это всегда значимость образа стимула.
Значимость стимула в PARAMS[2] зависит от count PARAMS[1] - числа подтверждений
func getOpower // сила значимости стимула в зависимости от count

Значимости в коде проекта (func getObjectImportance) обычно имеют значения величин от -10 0 до 10

Объект значимости - всегда имеет тип ActionsImage - конечный интегральный образ стимула, вне зависимости от его компонентов,
т.е. если пришла только фраза, без других компонентов, то при восприятии она имеет тип ActionsImage.
Значимость фразы с тоном или настроением или с действием - уже другая в данных условиях.


При каждом вызове consciousnessElementary определяется текущий объект наибольшой значимости в воспринимаемом -
в функции определения текущей Цели getMentalPurpose().

*/

package psychic

import "BOT/lib"

/////////////////////////////////////////////////

/*
Для определения текущих объектов восприятия и выделения одного из них - самого важного
*/
type extremImportance struct {
	objID     int //  объект значимости
	extremVal int // экстремальная значимость
}

// текущий значимый объект внимания (не сохраняется в файле):
var extremImportanceObject *extremImportance
var extremImportanceObjectOld *extremImportance // против повторения
// текущий значимый объект внимания с отрицательным эффектом extremImportance.extremVal
var problemExtremImportanceObject *extremImportance

/////////////////////////////////////////////

/*
	текущий субъект внимания (объект внимания к собственным мыслям)

используется в func getMentalEffect (Оценка суммарного ментального эффекта в период ожидания),
НУЖНО БЫ ЕГО ГДЕ-ТО выявлять
TODO м.б. в episodic_mental_memory_tree.go
*/
var extremImportanceMentalObject *extremImportance

///////////////////////////////////////////////////////

//////////////////////////////////////////////////////
/*
	найти extremImportanceObject при каждом новом стимуле curActiveActions

перекрывает прежний extremImportanceObject только если найден экстремальный объект
*/
func getExtremImportanceObject() {

	// в режиме сна мышление начинается так же с последнего curActiveActions
	if curActiveActions != nil { // curActiveActions обнуляется
		// объект объективного восприятия curActiveActions типа *ActionsImage - структура действий оператора при активации дерева автоматизмов
		eobj := getExtremObjFromID(curActiveActions.ID)
		if eobj != nil {
			if lib.Abs(eobj.extremVal) > 2 { // достаточно значимый объект
				extremImportanceObject = eobj
				CurrentInformationEnvironment.ExtremImportanceObjectID = eobj.objID
				if lib.Abs(eobj.extremVal) > 3 {
					runNewTheme(16, 2) // есть объект высокой значимости

					if eobj.extremVal < -3 { // проблемный объект с отрицательным эффектом
						// текущий значимый объект внимания с отрицательным эффектом extremImportance.extremVal
						problemExtremImportanceObject = eobj
					}
				}
			}
		}
	}
	return
}

//////////////////////////////////////////////////////

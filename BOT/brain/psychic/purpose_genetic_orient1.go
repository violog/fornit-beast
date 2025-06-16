/*
Для ориентировочного рефлекса типа 1
функции для определения Цели в данной ситуации - на уровне наследственных функций
исходя из текущей информационной среды CurrentInformationEnvironment:

*/

package psychic

import (
	"BOT/brain/gomeostas"
)

/*
	Определение Цели в данной ситуации - на уровне наследственных функций

Здесь getPurposeGenetic2AndRunAutomatizm.
*/
func getPurposeGeneticAndRunAutomatizm() *Automatizm {
	var flgBlockAtmz bool // флаг блокировки автоматизма

	purpose := getPurposeGenetic() // выбираются наиболее подходящие действия для автоматизмаы
	// мозжечковые рефлексы - самый первый уровень осознания - подгонка действий под заданную Цель.

	// нужно ли вообще шевелиться?
	// veryActualSituation: плохо для  1, 2, 7 и/или 8  параметров гомеостаза
	// на 2 стадии нужно добиться привязки к кнопкам автотизмов,  созданные по рефлексам (включая условные) - иначе на следующих стадиях будет много мусора при использвании кнопок в качестве поощрений/нказаний
	// это одна из причин, почему не стоит перескакивать стадии. Чем проще автоматизм, тем на более ранней стадии он должен быть сформирован.
	if purpose.veryActual || flgConditionReflexes == true || curActiveVerbalID == 0 {
		if purpose.actionID.ID > 0 && curActions.ActID != nil {
			/* сформировать пробный автоматизм моторного действия и сразу запустить его в действие
			Зафиксироваь время действия
			20 пульсов следить за измнением состояния и ответными действиями - считать следствием действия
			оценить результат и скорректировать силу мозжечком в записи автоматизма.
			Выбрать любое действие, т.к. послед создания автоматизма в данной ветке detectedActiveLastNodID
			он уже не вызовет orientation_1(), а будет orientation_2()
			*/
			atmzm := createAndRunAutomatizmFromPurpose(purpose)
			if atmzm.Usefulness >= 0 {
				if doWritingFile {
					SaveAutomatizm()
				}
				return atmzm
			} else {
				flgBlockAtmz = true
			}
		}
		/* нет действий (практически невозможная ситуация потому, что если нет рефлексов,
		то дейсвтвия в GetActualReflexAction() берутся из эффектов действий)
		*/
		if purpose.actionID == nil || flgBlockAtmz { // AutomatizmSuccessFromIdArr=make(map[int]*Automatizm)
			// ранее найденные удачные автоматизмы
			// выбрать из ранее удачного автоматизма, перекрыть цель новой и запустить новый автоматизм
			atmzm := chooseAutomatizmSuccessAndRun(purpose)
			if atmzm != nil {
				return atmzm
			}
			// нет действий, попробовать бессмысленно выдать фразы имеющиеся Вериниковские раз нужно что-то срочно делать
			if purpose.veryActual {
				// подобрать хоть как-то ассоциирующуюся фразу из имеющизся
				phraseID := findSuitablePhrase()
				if len(phraseID) > 0 {
					purpose.actionID.PhraseID = phraseID
					atmzm := createAutomatizm(purpose)
					// запустить автоматизм
					if RumAutomatizm(atmzm) {
						// отслеживать последствия в automatizm_result.go
						setAutomatizmRunning(atmzm, purpose)
					}
					// в automatizm_result.go после оценки результата будет осмысление с активацией Дерева Понимания
					if atmzm != nil {
						return atmzm
					}
				}
			}
		}
		//  ЗДЕСЬ активировать Дерево Понимания НЕ НУЖНО, действие уже запущено, омысление будет по результату.
	} else { // нет атаса, можно спокойно поэкспериментивроать, если есть любопытсво
		if gomeostas.BaseContextActive[2] || gomeostas.BaseContextActive[3] { // активен Поиск или Игра
			// тупо метод тыка
			// Тупо поэкспериментировать для пополнения опыта (не)удачных автоматизмов
			// TODO !не проверено!
			// в отличии от createAndRunAutomatizmFromPurpose(purpose) не использовать текущие рефлексы, а пробовать всякое
			// Выдавая это на стадии 3, тварь получает реакцию оператора, которую отзеркаливает
			atmzm := findAnySympleRandActions()
			if atmzm != nil {
				return atmzm
			}
		}
	}
	// НЕ ИГРА  И НЕ ПОИСК, нет атаса, пониженная мотивация что-то делать если нет актуальности
	if EvolushnStage == 2 {
		// нет действий, попробовать использовать AutomatizmSuccessFromIdArr.GomeoIdSuccesArr
		// выбрать из ранее удачного автоматизма, перекрыть цель новой и запустить новый автоматизм
		atmzm := chooseAutomatizmSuccessAndRun(purpose)
		return atmzm
	}
	// ЛЕНЬ
	return nil
}

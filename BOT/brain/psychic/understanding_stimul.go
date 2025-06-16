/* Понимание стимула


 */

package psychic

/////////////////////////////////////////////////////
/* распознаватель: не навязывают ли стимулом то, что не соотвествует текущей Теме и Цели?
У всех природных тварей важно строить поведение в направлении собственных потребностей, а не чужих, не попадать под воздействие чужих интересов, не выполнять чужие приказы, не учить то, что противоречит собственной адаптивности.
Уже в период преступной инициативы возникает доминирование некоего Эго – как противодействие чужому влиянию.
Например, если тварь наказывается за что-то, что отвечает ее интересам, то она должна этому противодействовать.

Меняет состояние CurrentInformationEnvironment.isStimulToForce:
	true - текущий стмул навязывает то, что не соотвествует текущей Теме и Цели

При CurrentInformationEnvironment.isStimulToForce=true игнорируется штатный автоматизм, останавливая все более низкоуровневое.
Начинается цикл подбора подходящей для Темы и Цели реакциию.

Эта функция может быть сделана множеством способов, порождая особенности реагирования стадии Инициативы.
*/
func recognizerStimulToForce() {
	CurrentInformationEnvironment.isStimulToForce = false
	//Это начинает работать с 4-й стадии (преступной инициативы).
	if EvolushnStage < 4 {
		return
	}
	if isIdleness() { // ЛЕНЬ
		return
	}
	// если ВАЖНАЯ ИЛИ ОПАСНАЯ СИТУАЦИЯ
	if CurrentInformationEnvironment.veryActualSituation || CurrentInformationEnvironment.danger {
		return
	}

	// антагонистично ли собственное состояние или нет
	var isAntagonizm = false
	/* цель определяется в infoFunc8() и заполняется структура problemTreeInfo.

		problemTreeInfo.themeID:
		1 "Негативный эффект моторного автоматизма"
		2 "Негативный эффект ментального автоматизма"
		3 "Состояние Плохо"
		4 "Стимул с Пульта"
		5 "Поисковый интерес"
		6 "Обучение с учителем"
		7 "Игнорирование оператором"
		8 "Игра"
		9 "Неудовлетворенность существущим"
		10 "Непонимание"
		11 "Разговор"
		12 "Сомнение в штатном автоматизме"
		13 "Защита"
		14 "Страх"
		15 "Агрессия"
		16 "Есть объект высокой значимости"
		17 "Улучшение настроения"

		problemTreeInfo.purposeID:
		Создается (в infoFunc8() createPurposeImageID()) образ цели типа PurposeImage
		type PurposeImage struct {
			ID int
			// цель - добиться target
			target int// 1- повторения, 2 - улучшения
			//добиться target значения PsyBaseMood -1 Плохое настроение, 0 Нормальное, 1 - хорошее настроение
			moodeID int // этот параметр будет всегда считаться заданным в getMentalPurposeEffect().
	// В случае Цели - улучшить эмоцию (target==2) - имеется в виду, что сумма весов позитивных эмоциональных контекстов превышает сумму вемов негатиных:	func isEmotonBetter(oldID int, curID int)
		emotonID int// добиться target такой эмоции

		// В случае Цели - улучшить ситуацию - func isSituationBetter(oldID int, curID int)
		situationID int// добиться PurposeImage.target данной ситуации SituationImage
	}
	*/
	var isActualTheme = false // есть актуальная для антагонизма Тема
	// текущая Тема:
	if problemTreeInfo.themeID > 0 {
		tID := problemTreeInfo.themeID
		// если нужен вес темы, то:
		//ThemeImageMapCheck()
		//themeType := ThemeImageFromID[problemTreeInfo.themeID].Type
		if tID == 1 || tID == 2 || tID == 9 || tID == 12 || tID == 13 || tID == 15 {
			isActualTheme = true
		}
	}
	// текущая цель всегда есть:
	if problemTreeInfo.purposeID > 0 {
		//		problem:=PurposeImageFromID[problemTreeInfo.purposeID]
		problem, ok := ReadePurposeImageFromID(problemTreeInfo.purposeID)
		if !ok {
			return
		}
		if isActualTheme && (problem.target == 2 || problem.emotonID > 0 || problem.situationID > 0) {
			isAntagonizm = true
		}

	}
	if !isAntagonizm { // нет причин для сопротивления
		return
	}

	if curActiveActions == nil {
		return
	}
	/* настроение при передаче фразы с Пульта:
	   	20-Хорошее    21-Плохое    22-Игровое    23-Учитель    24-Агрессивное   25-Защитное    26-Протест
	   ID возникает при добавлении 19 к номеру радиокнопки пульта, например, для Хорошее 1+19=20
	*/
	if curActiveActions.MoodID > 19 {
		mood := curActiveActions.MoodID
		if mood == 21 || mood == 24 || mood == 25 || mood == 26 {
			CurrentInformationEnvironment.isStimulToForce = true
			return
		}
		// м.б. еще какое-то распознавание ...
	}

	// сориентироваться по действия с пульта
	if curActiveActions.ActID != nil {
		act := curActiveActions.ActID
		for i := 0; i < len(act); i++ {
			if act[i] == 3 || act[i] == 9 || act[i] == 10 {
				CurrentInformationEnvironment.isStimulToForce = true
				return
			}
		}
		// м.б. еще какое-то распознавание ...
	}

	/* сориентироваться по фразе с пульта

	 */
	if curActiveActions.PhraseID != nil {

		// м.б. еще какое-то распознавание ...
	}

}

////////////////////////////////////////////////////////

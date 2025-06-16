/* Процессы осмысления: создание и использование ментальных автоматизмов
для Дерева понимания (или дерева ментальных автоматизмов)

*/

package psychic

import (
	"BOT/brain/gomeostas"
	"BOT/lib"
)

/*
	детекция ленивого состояния - по актуальности реагирования на услоивня.

Учитывает Доминанту, но не учитывает рвущегося на выполнение автоматизма.
*/
func isIdleness() bool {
	if CurrentInformationEnvironment.veryActualSituation || // не актуальная ситуация
		CurrentInformationEnvironment.danger || // нет опасности
		mentalInfoStruct.mentalPurposeID == 0 || // нет текущей поставленной Цели
		CurrentProblemDominanta != nil || // нет актуальной Доминанты
		(InterruptMemory == nil || len(InterruptMemory) == 0) { // нет прерываний размышлений к которым нужно вернуться
		return false
	}
	return true
}

// /////////////////////////////////////////////////////
// обработка структур в свободном состоянии, в первую очередь - эпизодической памяти
func processingFreeState(stopMentalWork bool) {

	if stopMentalWork { //- прекратить обработку
		return
	}

	// TODO переработка происходившего (эпизодическая память обрабатывается в GotoDreaming(true)
	//EpisodeMemoryLastCalcID - последний эпизод, который был осмыслен в лени или во сне
}

/*
	после периода ожидания:

Учесть последствия текущего ментального цикла,
выявить Правила и записать в эпизод.пямять.
Сразу после обработки периода ожидания запускается дерево понимания и объявный запуск consciousnessElementary(1,0)
так что никаких действий совершать в afterWaitingPeriod() или в самой afterWaitingPeriod() не следует.

effect =lastCommonDiffValue
*/
func afterWaitingPeriod(effect int) {

	/* Определить соотвествие желаемому и получаемому == определение степени достижения Цели PurposeImage,
	   в том числе для ментальных действий, а не только при объективных.
	      оценить совокупный эффект по нескольким атрибутам: объективному эффекту + достижение Цели + изменению значимостей
	*/
	effectValuation := getMentalEffect(effect) // детектор достижения Цели
	// записать в эпизод.пямять ментальный кадр. Запись saveNewEpisodic всегда предшествует: lastEpisodicMemID.
	saveNewMentalEpisodic(effectValuation) // здесь в конце обнуляется effectValuation

	// осмыслить результат, задав тему ThemeImage.Type=2
	if effectValuation < 0 { // негативный эффект
		// -effectValuation обычно большой и уже не перекрывается другой темой, так что задаем фиксированный 2
		runNewTheme(2, 2) //2 "Негативный эффект ментального автоматизма"
	}

}

///////////////////////////////////////////////////////

/*
Оценка суммарного ментального эффекта в период ожидания
*/
func getMentalEffect(effect0 int) int {
	/* улучшилось ли положение с учетом текущего PurposeImage 4-го узла ветки понимания?
	currentUnderstandingActivedNodes[]*UnderstandingNode // начиная с конечного к первому
	*/
	effectPurpose := getMentalPurposeEffect()

	// улучшилась ли значимость объекта внимания extremImportanceObject
	effectA := 0
	if extremImportanceObject != nil {
		oldVal := extremImportanceObject.extremVal
		eObj, newVal := getObjectsImportanceValue(extremImportanceObject.objID, detectedActiveLastProblemNodID)
		if eObj != nil {
			if newVal > oldVal { // улучшилось
				effectA = 1
			}
			if newVal < oldVal { // ухудшилось
				effectA = -1
			}
		}
	}

	// улучшилась ли значимость субъекта внимания extremImportanceMentalObject
	effectMA := 0
	if extremImportanceMentalObject != nil {
		oldVal := extremImportanceMentalObject.extremVal
		eObj, newVal := getObjectsImportanceValue(extremImportanceMentalObject.objID, detectedActiveLastProblemNodID)
		if eObj != nil { // если не нашел объект, это не значит что стало лучше/хуже
			if newVal > oldVal { // улучшилось
				effectMA = 1
			}
			if newVal < oldVal { // ухудшилось
				effectMA = -1
			}
		}
	}

	// Коэффициенты эффектов разного вида должны сильно влиять не ментальность твари: что для нее важнее.
	effectValuation := 0
	if EvolushnStage < 5 {
		effectValuation = effect0*3 + effectPurpose*1 + effectA*1 + effectMA*0
	} else { // повышенная роль заданной цели effect4 и объекта внимания
		effectValuation = effect0*2 + effectPurpose*3 + effectA*3 + effectMA*2
	}
	if effectValuation != 0 { // иначе при эффекте 0 выдаст -1 и потом заблокирует автоматизм
		if effectValuation > 1 {
			effectValuation = 1
		}
		if effectValuation < 1 {
			effectValuation = -1
		}
	}

	return effectValuation
}

/////////////////////////////////////////////////

/*
	Оценить эффект в период ожидания по ментальной цели ОБЪЕКТИВНО: после активации деревьев

Эффект от 0 до 10.
Эффект - в том, насколько достигнуто целевое состояние.
Цель считается достигнутой (эффект совершенных ментальных действий) если
заданные (НЕНУЛЕВЫЕ) параметры достигнуты.
При нескольких заданных параметрах цель может быть частично достигнута.
Эффект прикидывается в зависимости от полноты достижения цели:
при полном недостижении эффект ==0 при полном достижении +10
*/
var PsyBaseMoodOld = -10
var EmotionReceptionOld = 0
var SituationIDOld = 0

func getMentalPurposeEffect() int {
	if mentalInfoStruct.mentalPurposeID == 0 { //текущая цель не была задана
		return 0
	}
	var effect = 0
	var allPurposeHit = true // все цели достигнуты
	//	mp:=PurposeImageFromID[mentalInfoStruct.mentalPurposeID]
	mp, ok := ReadePurposeImageFromID(mentalInfoStruct.mentalPurposeID)
	if !ok {
		return 0
	}

	// настроение
	// mp.moodeID всегда считаться заданным в getMentalPurposeEffect().
	if mp.target == 1 { // добиться повторения
		if mp.moodeID == PsyBaseMood {
			effect += 3
		} else {
			allPurposeHit = false
		}
	}
	if mp.target == 2 { // добиться улучшения
		if PsyBaseMoodOld > PsyBaseMood {
			effect += 3
		} else {
			allPurposeHit = false
		}
	}
	PsyBaseMoodOld = PsyBaseMood
	////////////////////////////
	if mp.target == 1 { // добиться повторения
		if mp.emotonID > 0 {
			if mp.emotonID == CurrentEmotionReception.ID {
				effect += 3
			} else {
				allPurposeHit = false
			}
		}
	}
	///////////
	if mp.target == 2 { // добиться улучшения
		if mp.situationID > 0 {
			if isEmotonBetter(EmotionReceptionOld, CurrentEmotionReception.ID) {
				effect += 3
			} else {
				allPurposeHit = false
			}
		}
	}
	EmotionReceptionOld = CurrentEmotionReception.ID
	//////////////////////////////////
	//	var curSituationID=UnderstandingNodeFromID[detectedActiveLastUnderstandingNodID].SituationID
	unode, ok := ReadeUnderstandingNodeFromID(detectedActiveLastUnderstandingNodID)
	if !ok {
		return 0
	}
	var curSituationID = unode.SituationID
	if mp.target == 1 { // добиться повторения
		if mp.emotonID > 0 {
			if mp.situationID == curSituationID {
				effect += 3
			} else {
				allPurposeHit = false
			}
		}
	}
	///////////
	if mp.target == 2 { // добиться улучшения
		if mp.situationID > 0 {
			if isSituationBetter(SituationIDOld, curSituationID) {
				effect += 3
			} else {
				allPurposeHit = false
			}
		}
	}
	SituationIDOld = mentalInfoStruct.SituationID

	if effect > 10 {
		effect = 10
	}
	if allPurposeHit {
		return 10
	}
	return effect
}

////////////////////////////////////
/* Цель - улучшить эмоцию - имеется в виду, что
сумма весов позитивных эмоциональных контекстов превышает сумму весов негатиных.
Выбором весов и нагативаности определяется характер твари.
*/
func isEmotonBetter(oldID int, curID int) bool {
	if oldID == 0 || curID == 0 {
		return false
	}
	emOld, ok1 := EmotionFromIdArr[oldID]
	emCur, ok2 := EmotionFromIdArr[curID]
	if !ok1 || !ok2 {
		return false
	}
	// массивы активных контекстов
	aOld := emOld.BaseIDarr
	aCur := emCur.BaseIDarr
	var wArr = gomeostas.BaseContextWeight

	// старая сумма весов контекстов
	var wOld = 0 //позитивных и нешативных
	for i := 0; i < len(aOld); i++ {
		switch aOld[i] {
		case 1: //Пищевой	- Пищевое поведение, восполнение энергии, на что тратится время и тормозятся антагонистические стили поведения.
			wOld -= wArr[1]
		case 2: //	Поиск	- Поисковое поведение, любопытство. Обследование объекта внимания, поиск новых возможностей.
			wOld += wArr[1]
		case 3: //	Игра	- Игровое поведение - отработка опыта в облегченных ситуациях или при обучении.
			wOld += wArr[1]
		case 4: //	Гон	- Половое поведение. Тормозятся антагонистические стили
			wOld += wArr[1]
		case 5: //	Защита	- Оборонительные поведение для явных признаков угрозы или плохом состоянии.
			wOld -= wArr[1]
		case 6: //	Лень	- Апатия в благополучном или безысходном состоянии.
			wOld -= wArr[1]
		case 7: //	Ступор	- Оцепенелость при непреодолимой опастbase_context_ativnostности или когда нет мотивации при благополучии или отсуствии любых возможностей для активного поведения.
			wOld -= wArr[1]
		case 8: //	Страх	- Осторожность при признаках опасной ситуации.
			wOld -= wArr[1]
		case 9: //	Агрессия	- Агрессивное поведение для признаков легкой добычи или защиты (иногда - при плохом состоянии).
			wOld += wArr[1]
		case 10: //	Злость	- Безжалостность в случае низкой оценки .
			wOld -= wArr[1]
		case 11: //	Доброта	- Альтруистическое поведение.
			wOld += wArr[1]
		case 12: //	Сон - Состояние сна. Освобождение стрессового состояния. Реконструкция необработанной информации.
		}
	}

	var wCur = 0 //позитивных и негативных
	for i := 0; i < len(aCur); i++ {
		switch aOld[i] {
		case 1: //Пищевой	- Пищевое поведение, восполнение энергии, на что тратится время и тормозятся антагонистические стили поведения.
			wCur -= wArr[1]
		case 2: //	Поиск	- Поисковое поведение, любопытство. Обследование объекта внимания, поиск новых возможностей.
			wCur += wArr[1]
		case 3: //	Игра	- Игровое поведение - отработка опыта в облегченных ситуациях или при обучении.
			wCur += wArr[1]
		case 4: //	Гон	- Половое поведение. Тормозятся антагонистические стили
			wCur += wArr[1]
		case 5: //	Защита	- Оборонительные поведение для явных признаков угрозы или плохом состоянии.
			wCur -= wArr[1]
		case 6: //	Лень	- Апатия в благополучном или безысходном состоянии.
			wCur -= wArr[1]
		case 7: //	Ступор	- Оцепенелость при непреодолимой опастbase_context_ativnostности или когда нет мотивации при благополучии или отсуствии любых возможностей для активного поведения.
			wCur -= wArr[1]
		case 8: //	Страх	- Осторожность при признаках опасной ситуации.
			wCur -= wArr[1]
		case 9: //	Агрессия	- Агрессивное поведение для признаков легкой добычи или защиты (иногда - при плохом состоянии).
			wCur += wArr[1]
		case 10: //	Злость	- Безжалостность в случае низкой оценки .
			wCur -= wArr[1]
		case 11: //	Доброта	- Альтруистическое поведение.
			wCur += wArr[1]
		case 12: //	Сон - Состояние сна. Освобождение стрессового состояния. Реконструкция необработанной информации.
		}
	}
	// сравнение эмоций на позитивность
	if wCur > wOld {
		return true
	}
	return false
}

////////////////////////////////////

/*
сравнение старой и новой ситуации
*/
func isSituationBetter(oldID int, curID int) bool {
	if oldID == 0 || curID == 0 {
		return false
	}
	//sOld,ok1:=SituationImageFromIdArr[oldID]
	sOld, ok := ReadeSituationImageFromIdArr(oldID)
	if !ok {
		return false
	}

	//sCur,ok2:=SituationImageFromIdArr[curID]
	sCur, ok2 := ReadeSituationImageFromIdArr(curID)
	if !ok2 {
		return false
	}
	/////////////////////////////////
	var wOld = 0 //позитивных и негативных
	switch sOld.SituationType {
	case 4:
		wOld += 5 //все спокойно, можно экспериментивароть
	case 5:
		wOld += -5 //оператор не прореагировал на действия в течение периода ожидания - игнорирует
		// оператор выбрал настроение
	case 11:
		wOld += 10 //Хорошее
	case 12:
		wOld += -8 //Плохое
	case 13:
		wOld += 6 //Игровое
	case 14:
		wOld += 3 //Поучить
	case 15:
		wOld += 2 //Агрессивное
	case 16:
		wOld += -5 //Защитное
	case 17:
		wOld += -6 //Протест
		//оператор нажал кнопку
	case 21:
		wOld += -3 //Непонятно
	case 22:
		wOld += 3 //Понятно
	case 23:
		wOld += -8 //Наказать
	case 24:
		wOld += 8 //Поощрить
	case 25:
		wOld += 6 //Накормить
	case 26:
		wOld += 4 //Успокоить
	case 27:
		wOld += 5 //Поиграть
	case 28:
		wOld += 5 //Предложить поучить
	case 29:
		wOld += -4 //Игнорировать
	case 30:
		wOld += -4 //Сделать больно
	case 31:
		wOld += 5 //Сделать приятно
	case 32:
		wOld += -2 //Заплакать
	case 33:
		wOld += 2 //Засмеяться
	case 34:
		wOld += 3 //Обрадоваться
	case 35:
		wOld += -1 //Испугаться
	case 36:
		wOld += 3 //Простить
	case 37:
		wOld += 7 //Вылечить
	}
	//////////////////////////////////
	var wCur = 0 //позитивных и негативных
	switch sCur.SituationType {
	case 4:
		wCur += 5 //все спокойно, можно экспериментивароть
	case 5:
		wCur += -5 //оператор не прореагировал на действия в течение периода ожидания - игнорирует
		// оператор выбрал настроение
	case 11:
		wCur += 10 //Хорошее
	case 12:
		wCur += -8 //Плохое
	case 13:
		wCur += 6 //Игровое
	case 14:
		wCur += 3 //Поучить
	case 15:
		wCur += 2 //Агрессивное
	case 16:
		wCur += -5 //Защитное
	case 17:
		wCur += -6 //Протест
		//оператор нажал кнопку
	case 21:
		wCur += -3 //Непонятно
	case 22:
		wCur += 3 //Понятно
	case 23:
		wCur += -8 //Наказать
	case 24:
		wCur += 8 //Поощрить
	case 25:
		wCur += 6 //Накормить
	case 26:
		wCur += 4 //Успокоить
	case 27:
		wCur += 5 //Поиграть
	case 28:
		wCur += 5 //Предложить поучить
	case 29:
		wCur += -4 //Игнорировать
	case 30:
		wCur += -4 //Сделать больно
	case 31:
		wCur += 5 //Сделать приятно
	case 32:
		wCur += -2 //Заплакать
	case 33:
		wCur += 2 //Засмеяться
	case 34:
		wCur += 3 //Обрадоваться
	case 35:
		wCur += -1 //Испугаться
	case 36:
		wCur += 3 //Простить
	case 37:
		wCur += 7 //Вылечить
	}

	/////////////////////////////////
	// сравнение ситуаций на позитивность
	if wCur > wOld {
		return true
	}
	return false
}

//////////////////////////////////////////////////////////////////

// есть ли данные Базовые контексты в состоаве эмоции: true - да, есть.
func existsBaseContextFromEmotionID(contextsArr []int, emotionId int) bool {
	for _, v := range EmotionFromIdArr[emotionId].BaseIDarr {
		if lib.ExistsValInArr(contextsArr, v) {
			return true
		}
	}
	return false
}

//////////////////////////////////////////////

/*
Подавление мешающих стимулов при серьезном поиске решений проблемы.
Это обеспечивает "тишину мыслей" (fornit.ru/17954)
Нужно приглушать 1 и 2-й уровни осмысления чтобы не мешали,
а вместо автоматизмов выдавать действие "игнорировать" и тем самым блокировать рефлексы.
Это состояние погруженности нужно снимать при серьезных стимулах высокой значяимости?
а так же при конце атаса, наступления лени и дремы - просто установкой флага:
isRepressionStimulsNoise=false

isRepressionStimulsNoise действует при при активации дерева автоматизмов, до активации ветки
и блокирует весь процесс активации, так же нет прерываний на более важный стимул
и нет решения доминанты по аналогии.
*/
var isRepressionStimulsNoise = false

func repressionStimulsNoises() {
	if EvolushnStage < 5 {
		return
	}
	// есть ли проблемный объект с отрицательным эффектом - основа мук творчества
	if problemExtremImportanceObject != nil {
		isRepressionStimulsNoise = true
	}
}

// контроль удержания режима isRepressionStimulsNoise
func checkRepressionStimulsNoises() {
	if EvolushnStage < 5 {
		return
	}
	if !isRepressionStimulsNoise {
		return
	}
	if IsSleeping || isIdleness() {
		isRepressionStimulsNoise = false
	}

}

// сторожевая функция определения критического состояния гомеостаза
func watchdogFunctionGomeo() {
	if !allowWatchdogFunction() {
		return
	}
	CurrentInformationEnvironment.danger = GetAttentionDanger()
	if CurrentInformationEnvironment.danger {
		isRepressionStimulsNoise = false
		return
	}
	CurrentInformationEnvironment.veryActualSituation, _ = gomeostas.FindTargetGomeostazID()
	if CurrentInformationEnvironment.veryActualSituation {
		isRepressionStimulsNoise = false
		return
	}
	painValue, _ = gomeostas.GetCurPainJoy()
	if painValue > 5 {
		isRepressionStimulsNoise = false
		return
	}
}

// сторожевая функция определения очень важного объекта внимания
func watchdogFunctionStimul(curActiveActions *ActionsImage) {
	if !allowWatchdogFunction() {
		return
	}
	eobj := getExtremObjFromID(curActiveActions.ID)
	if eobj == nil {
		return
	}
	if eobj.extremVal < -5 {

		/* TODO при появлении атасного стимула записать Правило в эпиз.память ??
		 наверное не стоит, т.к. тогда всегда будет негатив при задумчивости...
		НО если оператор наказал, то для этих условий правило уместно
		*/

		isRepressionStimulsNoise = false
		return
	}
}

func allowWatchdogFunction() bool {
	if EvolushnStage < 5 {
		return false
	}
	if !isRepressionStimulsNoise {
		return false
	}
	return true
}

///////////////////////////////////////////////////

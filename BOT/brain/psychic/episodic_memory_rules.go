/*  фиксация Правил в эпизодической памяти

Обычно Правила имеют Стимул (от оператора), ответное действие Beast и эффект
но если действие Beast было не в ответ на Стимул, а самостоятельным (после провокации func infoFunc31)
то ID Стимула в Правиле будет равен нулю.

*/

package psychic

import (
	"BOT/brain/action_sensor"
	wordSensor "BOT/brain/words_sensor"
	"BOT/lib"
)

///////////////////////////////////////////////////////////

// ЗАПИСЬ ЭПИЗОДОВ С ПРАВИЛАМИ
/*
 образ действий оператора Стимул:	curStimulImageID и есть Ответ: curActiveActions
*/
func fixEpizMemoryRules(lastCommonDiffValue int) {
	if LastAutomatizmWeiting == nil {
		return
	}

	if wasRunProvocationFunc { //была  провокация оператора на действие func infoFunc31
		curStimulImageID = 0 // без стимула от оператора
	} else {
		// curStimulImageID - образ действий оператора перед Ответом для записи Правила
		if curStimulImageID == 0 {
			return
		}
		if curStimulImage == nil { // где-то сбрасывается!!??
			//		curStimulImage=ActionsImageArr[curStimulImageID]
			aImage, ok := ReadeActionsImageArr(curStimulImageID)
			if ok {
				curStimulImage = aImage
			}
		}
		if curStimulImage.ActID == nil && curStimulImage.PhraseID == nil { // не писать Правила с пустым Стимулом
			return
		}
	}

	answerID := 0 // ответ Beast
	if LastAutomatizmWeiting != nil {
		answerID = LastAutomatizmWeiting.ActionsImageID // ответный образ действий Beast
		_, ok := ReadeActionsImageArr(answerID)
		if !ok {
			return
		}

		if !wasRunProvocationFunc {
			if curStimulImageID == answerID { //попугайство просьбы показать как нужно
				return
			}
		}
	}
	if lastRunVolutionAction != nil {
		answerID = lastRunVolutionAction.ID // ответный образ действий Beast
		lib.MapFree(MapGwardAutomatizmNextStringFromID)
		if AutomatizmNextStringFromID[answerID] == nil {
			return
		}
	}

	if answerID == 0 {
		return
	}

	if LastAutomatizmWeiting.NextID > 0 { // автоматизм с довеском акций типа AmtzmNextString
		// создать AmtzmNextString из всех действий автоматизма
		asID, _ := createNextStringFromAutomatizm(LastAutomatizmWeiting)
		answerID = prefixActionIdValue + asID // увеличить ID на метку AmtzmNextString
	}
	if lastRunVolutionAction != nil { // было произвольное действие
		answerID = prefixActionIdValue + answerID // увеличить ID на метку AmtzmNextString
	}

	if !wasRunProvocationFunc {
		if detectedActiveLastNodPrevID == 0 || detectedActiveLastUnderstandingNodPrevID == 0 {
			return
		}
	}
	// Запись Правила (именно detectedActiveLastNodPrevID а не detectedActiveLastNodID)

	/*	Зачем писать ПРЯМОЕ одиночное правило, если есть автоматизм на тот же стимул? Учительское правило как пробник - другое дело
		Имхо: одиночное правило - только учительское, чтобы по нему потом пробник сделать. Анализ по правилам делается только по групповым - с целью посмотреть, чем может закончится диалог, позитивом или негативом
		Групповое правило - это динамический стимул, реакция на цепочку стимул-ответов типа:
		Ты кто? - бот, уверен? - да, [поощрить]
		Кто ты? - Бог, уверен? - да, [наказать]
		На один и тот же статический стимул  "уверен" выдается действие "да" - и получает подзатыльник или пряник. Чтобы решить проблему как реагировать, нужно смотреть всю цепочку диалога, представив ее как динамичекйи стимул
		Он то как раз разный - потому и реакции разные. Можно сказать, что бота наказывают не заего ответ на "уверен", а за ответ на "ты кто?" и "кто ты?" - или за любое другое отличие в цепочке звеньев стимул-ответов динамического стимула
	*/
	//currentRulesID, _ = createNewRules(0, detectedActiveLastNodPrevID,detectedActiveLastUnderstandingNodPrevID,[]int{TriggerAndAction},true)
	//if currentRulesID == 0{return 0}

	//lib.WritePultConsol("<span style='color:green'>Записано <b>ПРАВИЛО № "+strconv.Itoa(currentRulesID)+"</b></span>") // уже есть сообщение в createNewRules()

	//////////// ЗАПИСАТЬ В ДЕРЕВО НОВЫЙ ЭПИЗОД
	// curStimulImageID - предыдущий стимул (текущий - curActiveActionsID)
	//usedOldCondition = true    // испольховать предыдущее значение условий (до стимула от оператора)
	if wasRunProvocationFunc { //после провокации func infoFunc31 нет образа стимула т.к. beast провоцирует при отсуствии стимула
		saveNewEpisodic(0, answerID, lastCommonDiffValue, 0)
	} else {
		// prevStimulsEffect - т.к. первый стимул уже перекрыт ответом оператора на действия beast
		saveNewEpisodic(curStimulImageID, answerID, lastCommonDiffValue, prevStimulsEffect)
	}

	return
}

///////////////////////////////////////////////////////////////////////

/*
	запоминает авторитерный ответ Оператора на совершенное действие - учительское правило.

- одиночное Правило как Оператор отвечает на действиЯ Beast
>> авторитетное Правило всегда имеет эффект +1
так нельзя, иначе негативное действие оператора будет расценено как положительное и в дальнейшем будет выбираться по этому признаку
то есть, получив негативный ответ от оператора, бот все равно будет его давать, потому как он записан у него как положительный

Такое Правило используется в случае отсуствия решения как отвечать,
т.к. не пришется групповое Правило дял точного бездумного реагирования
(хотя на уровне эпиз.памяти и можно вычленять такие групповые Правила,
выделяя Стимул следующего Правила как ответ на действия Beast).
*/
func fixEpizMemoryTeachRules(lastCommonDiffValue int) {
	if curActiveActions == nil { // так было когда использовалась func LastAutomatipmCorrection (теперь ее нет)
		// Пусть на всякий случай будет: вытащить стимул из базовой рецепции:
		if action_sensor.PultActionPulsCount == PulsCount { // только если оператор только что ответил, а не берется какая-то прежняя акция
			curActiveActionsID, curActiveActions = CreateNewlastActionsImageID(0, 0, action_sensor.CheckCurActions(), wordSensor.CurrentPhrasesIDarr, wordSensor.DetectedTone, wordSensor.CurPultMood, true)
			wasChangingMoodCondition() // раз не успевает...
		}
		if curActiveActions == nil {
			return
		}
	}

	// curActiveActionsID - образ действий оператора после Ответа для записи Правила
	if curActiveActionsID == 0 {
		return
	}

	if curActiveActions.ActID == nil && curActiveActions.PhraseID == nil { // не писать Правила с пустым Стимулом
		return
	}

	/* Если есть ли автоматизм с действием оператора curActiveActionsID, и если у него atmtzm.Usefulness<0 - постепенно снять блокировку
	потому как это - новое авторитарное подтвержение полезности.
	Но если была запущена func13 то попугайский автоматизм-переспрос разблокируется на следующем шаге. Поэтому разблокировка
	только в случае, когда нет запуска func13.
	Так же нельзя разблокировать, если только что было применено наказание кнопками действий, иначе оно нивелируется
	*/
	answerID := 0 // // образ действий Beast перел ответом Оператора
	if prevLastAutomatizmWeiting != nil {
		// предыдущий момент запуска автоматизма был задолго от последующего действия оператора
		if (PulsCount - prevLastDetectedActiveLastPulsCount) > 25 {
			return
		}
		if curFunc13ID == 0 && lastCommonDiffValue >= 0 {
			checkForUnbolokingAutomatizm(curActiveActionsID)
		}
		answerID = prevLastAutomatizmWeiting.ActionsImageID // ответный образ действий Beast
		_, ok := ReadeActionsImageArr(answerID)
		if !ok {
			return
		}
		if prevLastAutomatizmWeiting.NextID > 0 {
			asID, _ := createNextStringFromAutomatizm(prevLastAutomatizmWeiting)
			answerID = prefixActionIdValue + asID
		}
	}

	if prevLastRunVolutionAction != nil {
		answerID = prevLastRunVolutionAction.ID // ответный образ действий Beast
		lib.MapFree(MapGwardAutomatizmNextStringFromID)
		if AutomatizmNextStringFromID[answerID] == nil {
			return
		}
	}

	if lastRunVolutionAction != nil { // было произвольное действие
		answerID = prefixActionIdValue + answerID // увеличить ID на метку AmtzmNextString
	}

	if answerID == 0 {
		return
	}

	if action_sensor.PultActionPulsCount == PulsCount { // только если оператор только что ответил, а не берется какая-то прежняя акция

		// уже есть стимул от оператора curActiveActionsID незачем получать новый из базовых сенсоров
		//curAct, _ := CreateNewlastActionsImageID(0, 0, action_sensor.CheckCurActions(), wordSensor.CurrentPhrasesIDarr, wordSensor.DetectedTone, wordSensor.CurPultMood, true)

		// в стадиях до 4 обнуляется в automatizmTreeActivation()
		// очистить фразу после использования, чтобы не влияла на следующую активность
		wordSensor.CurrentPhrasesIDarr = nil // остается еще wordSensor.CurretWordsIDarr

		//  ЗАПИСАТЬ В ДЕРЕВО НОВЫЙ ЭПИЗОД  эффект =100 - метка, что это - учительское правило.
		usedOldCondition = true // испольховать предыдущее значение условий (до стимула от оператора)
		saveNewEpisodic(answerID, curActiveActionsID, 100, globalStimulsEffect)
	}

	return
}

/////////////////////////////////////////////////////////////

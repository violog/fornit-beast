/* Инфо-функция поиска Цели - как объекта произвольного внимания:
того из всего воспринимаемого, что имеет наибольшую значимость
т.к. именно наибольшая значимость должна осмысливаться.

targetID [] можно пока вообще не трогать, так же как и veryActual -
оба просто констатируют что нужно бы улучшить.

Тут главное - определить PurposeImage.actionID -
выбранный образ действия бота для достижения данной цели (ActionsImage)

т.е. если не определно произвольное PurposeImage.actionID или вообще не задана Цель,
то целью автоматически становится улучшение состояния, см. func valuationPurpose().

На пятой ступени развития целью может стать решение Доминанты нерешенной проблемы, и только
если нет Доминант, то работает неосознаваемая Цель mentalInfoStruct.mentalPurposeID.
*/

package psychic

import (
	"strconv"
)

//////////////////////////////////////////

//////////////////////////////////////////////////////////
/* Ментальное определение ближайшей Цели в данной ситуации по теме в getMentalPurpose()
Для цикла мышления с новой темой mentalInfoStruct.ThemeImageType
Постановка цели для текущего цикла размышления, чтобы оценить эффект для мент.Правила.
*/
var newCycklePulsCount = 0 //время запуска infoFunc8()
/*
// переменные только дял этой функции:
var oldThemeID8=0
var oldPurposeID8=0
func initInfoFunc8pars(){// с каждым ориентировочным рефлексом функция циклов осмысления func consciousnessElementary()

		oldThemeID8=0
		oldPurposeID8=0
	}
*/
func infoFunc8(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	//  цель может модифицировать в цикле, но не подряд!
	//!!!! может быть подряд! if c.lastFuncID == 8 { // не вызывать, если только что было
	//	return false
	//}
	setCurIfoFuncID(c, 8)

	//if newCycklePulsCount==PulsCount{// от случайных повторых запусков МОЖЕТ В ОДНОМ ЦИКЛЕ СМЕНИТЬСЯ ТЕМА И ЦЕЛЬ!
	/* проверяется использование infoFunc8 в течение жизни цикла, чтобы не молотил зря:
	   если не менялась тема и цель
	   Даже если не было других инфо-фукнций после 8-й, но тема или цель менялись, то нужно отработать infoFunc8 снова.

	   	if problemTreeInfo.themeID>0 && mentalInfoStruct.mentalPurposeID>0 &&
	   	oldThemeID8==problemTreeInfo.themeID && oldPurposeID8==mentalInfoStruct.mentalPurposeID {
	   		// уже был вызов при данных теме и цели
	   // TODO м.б. в таком случае привенить что-то радикальное в случае атасной ситуации...
	   		return false
	   	}
	   //}
	   	newCycklePulsCount=PulsCount
	   	oldThemeID8=problemTreeInfo.themeID
	   	oldPurposeID8=mentalInfoStruct.mentalPurposeID
	*/

	// для каждой новой цели должен начинаться свой кадр ментальной эпиз.памяти
	// !!!! clinerFuncSequence() не сбрасывать wasRunPurposeActionFunc !!!
	infoFuncSequence = nil // т.к. с каждым стимулом обновляется главный цикл resetMineCycleAndBeginAsNew()

	// если нет темы, после просыпания и т.п.
	if problemTreeInfo.themeID == 0 {
		//"Улучшение настроения" - 17 - базовая тема когда нет ничего другого
		themeID, theme := createThemeImageID(0, 2, 17, LifeTime, true)
		problemTreeInfo.themeID = themeID
		themeType := theme.Type
		mentalInfoStruct.ThemeImageType = themeType
		oldThemeID = themeID
	} else {
		// не сменилась тема
		//		themeType := ThemeImageFromID[problemTreeInfo.themeID].Type
		node, ok := ReadeThemeImageFromID(problemTreeInfo.themeID)
		if ok {
			themeType := node.Type

			if oldThemeImageType > 0 && oldThemeImageType == themeType { //  && mentalInfoStruct.mentalPurposeID >0
				// если блокировать запуск дерево проблем при одинаковой теме, то значимости от разных стимулов будут привязываться
				// к одному узлу дерева проблем, что явно не правильно. Тема может и одна, но situationTreeID и purposeID могут быть разные
				if GetMotorsAutomatizmListFromTreeId(detectedActiveLastNodID) == nil {
					// блокировать активацию при одинаковой теме только если нет вообще никакого автоматизма - чтобы сработало func13
					return false
				}
			}

			oldThemeImageType = themeType
		}
	}

	if getMentalPurpose() { // активация дерева понимания, если готовы переменные узлов дерева
		if problemTreeInfo.situationTreeID > 0 && problemTreeInfo.themeID > 0 && problemTreeInfo.purposeID > 0 {
			//if !c.idle || show_all_logs {
			c.log += "В infoFunc8() найдена Цель <b> <span style='cursor:pointer;color:blue' onClick='get_purpose(" + strconv.Itoa(mentalInfoStruct.mentalPurposeID) + ")'>" + strconv.Itoa(mentalInfoStruct.mentalPurposeID) + "</span>" + "</b><br>"
			//}

			//При вызове  infoFunc8() происходит активация дерева проблем и НЕ ДОЛЖНА БЫТЬ новая об.активация осмысления!
			if !wasRunPurposeActionFunc { // true - уже была активация ProblemTreeActivation() по func infoFunc14
				noProblemTreeActivation = true
				ProblemTreeActivation() //Активация дерева понимания проблемы
				noProblemTreeActivation = false
			}
			// если вернет true, когда надо в consciousnessElementary отработать ситуации с негативной значимостью и плохим эффектом группового правила,
			// то в usualThinkProcess не дойдет до infoFunc2, а в самой infoFunc2 не дойдет до func13
			if mentalInfoStruct.motorAtmzmID > 0 {
				return true
			}
		}
	}
	return false
}

//////////////////////////////////////
/* определить mentalInfoStruct.mentalPurposeID
 вызов из if IsFirstActivation{

Субъективная цель. Это - не сознательная цель, а мотивирующая потребность, дающая направленность мышлению.
Текущая тема должна давать контекст размышлениям вот каким образом.
Поддержание актуальной темы-контекста общения - внимание к актуальному объекту.
Для данного объекта выбирается целевая цепочка Правил,
которая начинает выполняться по шагам, всякий раз снова корректируя цепочку правил.
Так строится разговор.
*/
func getMentalPurpose() bool { // обязательно находит Цель
	//	mentalInfoStruct.mentalPurposeID=0  не нужно! т.к. обязательно находит Цель и перекрывает

	//	theme:=mentalInfoStruct.ThemeImageType //не очищать mentalInfoStruct.ThemeImageType
	//	clinerMentalInfo()// начинается новый цикл мышления
	//	mentalInfoStruct.ThemeImageType=theme

	if mentalInfoStruct.ThemeImageType == 0 || mentalInfoStruct.ThemeImageType == 17 { // не определена тема - тупо постараться улучшить настроение
		mentalInfoStruct.ThemeImageType = 17 //Улучшение настроения - базовая тема когда нет ничего другого
		id, _ := createPurposeImageID(0, 1, 3, 0, 0, true)
		mentalInfoStruct.mentalPurposeID = id
		problemTreeInfo.purposeID = id //инфа для активации дерева проблем
		return true
	}
	///////////////////////////////////
	// есть рвущийся автоматизм, или есть тема "Сомнение в штатном автоматизме"
	if mentalInfoStruct.motorAtmzmID > 0 || mentalInfoStruct.ThemeImageType == 12 {
		unode, ok := ReadeUnderstandingNodeFromID(detectedActiveLastUnderstandingNodID)
		if ok {
			id, _ := createPurposeImageID(0, 2, unode.Mood,
				unode.EmotionID,
				unode.SituationID, true)
			mentalInfoStruct.mentalPurposeID = id
			problemTreeInfo.purposeID = id //инфа для активации дерева проблем
			// далее TODO НУЖНО ДУМАТЬ ПРО ЭТО
			//infoFunc6()//ПОДВЕРГНУТЬ СОМНЕНИЮ автоматизм, если нет опасности (не нужно реагировать аффектно) и ситуация важна
			return true
		}
	}
	///////////////////////////////////

	// Есть объект высокой негативной значимости
	if mentalInfoStruct.ThemeImageType == 16 {
		/* extremImportanceObject - объект наибольшой негативной значимости в воспринимаемом
		extremImportanceObject.extremVal - НАГАТИВНАЯ ЗНАЧИМОСТЬ объекта
		 определиться с Целью по этому объекту:
		 TODO	найти situacTreeID дерева ситуации в Правилах с этим объектом, где улучшается эффект
			situacTreeID:= ....
			и создать Цель:
			id,_:=createPurposeImageID(0,2,UnderstandingNodeFromID[situacTreeID].Mood,
							UnderstandingNodeFromID[situacTreeID].EmotionID,
							UnderstandingNodeFromID[situacTreeID].SituationID,true)
						mentalInfoStruct.mentalPurposeID=id
						return
		*/
	}
	///////////////////////////////////

	if mentalInfoStruct.ThemeImageType == 3 { //Состояние Плохо - достичь Хорошо т.е. после Плохи всегда бывает улучшение в Хорошо
		id, _ := createPurposeImageID(0, 1, 3, 0, 0, true)
		mentalInfoStruct.mentalPurposeID = id
		problemTreeInfo.purposeID = id //инфа для активации дерева проблем
		return true
	}

	/* Цель - улучшение данной ситуации - для Тем с ID:
	1 "Негативный эффект моторного автоматизма"
	2 "Негативный эффект ментального автоматизма"
	4 "Стимул с Пульта"
	7 "Игнорирование оператором"
	8 "Игра"
	10 "Непонимание"
	13 "Защита"
	14 "Страх"
	15 "Агрессия"
	16"Есть объект высокой значимости"
	*/

	switch mentalInfoStruct.ThemeImageType {
	case 1, 2, 4, 7, 10, 13, 14, 15, 16:
		unode, ok := ReadeUnderstandingNodeFromID(detectedActiveLastUnderstandingNodID)
		if ok {
			id, _ := createPurposeImageID(0, 2, unode.Mood,
				unode.EmotionID,
				unode.SituationID, true)
			mentalInfoStruct.mentalPurposeID = id
			problemTreeInfo.purposeID = id //инфа для активации дерева проблем
			return true
		}
	}

	/* Цель - повторнеие (достижение) данной ситуации - для Тем с ID:
	1 "Негативный эффект моторного автоматизма"
	2 "Негативный эффект ментального автоматизма"
	3 "Состояние Плохо"
	4 "Стимул с Пульта"
	5 "Поисковый интерес"
	6 "Обучение с учителем"
	8 "Игра"
	9 "Неудовлетворенность существущим"
	*/
	switch mentalInfoStruct.ThemeImageType {
	case 5, 6, 8, 9:
		unode, ok := ReadeUnderstandingNodeFromID(detectedActiveLastUnderstandingNodID)
		if ok {
			id, _ := createPurposeImageID(0, 1, unode.Mood,
				unode.EmotionID,
				unode.SituationID, true)
			mentalInfoStruct.mentalPurposeID = id
			problemTreeInfo.purposeID = id //инфа для активации дерева проблем
			return true
		}
	}
	///////////////////////////////////

	if mentalInfoStruct.ThemeImageType == 11 { //Действие оператора - тут TODO лучше выбрать объект по Правилам!
		/* есть доступная инфа о действии оператора:
		   	curStimulImage=curActiveActions
		      	curStimulImageID=curActiveActionsID
		       curActiveVerbalID = verbID
		      	curActions.PhraseID = PhraseID
		      	curActions.ToneID = ToneID
		      	curActions.MoodID = MoodID

		   		TODO лучше выбрать объект по Правилам!
		*/
		unode, ok := ReadeUnderstandingNodeFromID(detectedActiveLastUnderstandingNodID)
		if ok {
			id, _ := createPurposeImageID(0, 2, unode.Mood,
				unode.EmotionID,
				unode.SituationID, true)
			mentalInfoStruct.mentalPurposeID = id
			problemTreeInfo.purposeID = id //инфа для активации дерева проблем
			return true
		}
	}

	////////////////////// Цели - поиск в Правилах TODO !!!!!
	/*
		// Цели - поиск в Правилах

		// Текущая тема должна давать контекст размышлениям вот каким образом.
		Поддержание актуальной темы-контекста общения - внимание к актуальному объекту.
		Для данного объекта выбирается целевая цепочка Правил, которая начинает выполняться по шагам,
		всякий раз снова корректируя цепочку правил. Так строится разговор.*/
	/*

			//ищем PurposeImage.actionID в контексте активных деревьев

			найти - с учетом Правил!!!!
			??На стадии 4 - провоцировать оператора на ответы (почему, зачем, что такое?)
		Нужно искать не только в контексте эмоции, а активных веток деревьев detectedActiveLastNodID и detectedActiveLastUnderstandingNodID
			Алгоритм:
		Ищем в ОБЪЕКТИВНЫХ Правилах подходящее действие ,
			смотрим в последних кадрах эпизод.памяти такое Правило и в его продолжении есть позитивный эффект,
			то берем оттуда действие оператора.*/
	/*
	   	if EpisodeMemoryObjects!=nil {// еще нет эпиз.памяти, так что и цели нет...
	   		return
	   	}
	   // тупо - для 4-й ступени (и 5-й, если нет Доминанты), когда набирается значимость и Правила
	   		indexE := 0 // индекс в массиве эпизод.памяти, чтобы по нему смотреть, что было дальне
	   		rID, exact := getRulesArrFromTrigger(curActiveActionsID,true)
	   		if exact > 3 { // слишком сомнительно

	   			rulexID := 0
	   			maxSteps := 1000
	   			for limit := 5; limit > 1; limit-- {
	   				rulexID, indexE = getRulesFromEpisodicsSlice(limit, maxSteps)
	   				if rulexID > 0 {
	   					break
	   				}
	   				maxSteps = maxSteps / 2
	   			}
	   		} else { // вполне уверенно найдено Правило, ищем его индекс в эпиз.памчяти
	   			if EpisodeMemoryObjects!=nil {
	   				for i := len(EpisodeMemoryObjects); i >= 0; i-- {
	   					ep := EpisodeMemoryObjects[i]
	   					if ep.Type == 0 && ep.TriggerAndActionID == rID {
	   						indexE = i
	   						break
	   					}
	   				}
	   			}
	   		}
	   		//////////////////////////
	   		if indexE == 0 {
	   			return
	   		}

	   		// есть ли последующий кадр эпизод.памяти
	   		lastEM := EpisodeMemoryObjects[indexE+1]
	   		if lastEM == nil {
	   			return  // придется обойтись без PurposeImage.actionID
	   		}
	   		// хороший ли эфект
	   		rArr := rulesArr[lastEM.TriggerAndActionID]
	   		if rArr == nil {
	   			return  // придется обойтись без PurposeImage.actionID
	   		}
	   		// выдать конечное праило, если оно с хорошим эффектом
	   		var rеserve = 0 // резервные Правила, если не найдено точно в контексте
	   		ta := TriggerAndActionArr[lastEM.TriggerAndActionID]
	   		if ta == nil {
	   			return // придется обойтись без PurposeImage.actionID
	   		}
	   		if ta.Effect > 0 { // с хорошим эффектом
	   			rеserve = ta.Action
	   			if lastEM.NodeAID == detectedActiveLastNodID {
	   				rеserve = ta.Action
	   				if lastEM.NodeSID == detectedActiveLastUnderstandingNodID {
	   					rеserve = ta.Action
	   				}
	   			}
	   		}
	   		if rеserve > 0 { // нашли действие
	   			if savePurposeGenetic == nil {
	   				getPurposeGenetic()
	   			}
	   			// создать Цель
	   // !!!! purposeID, _ := createPurposeImageID(0, 2,savePurposeGenetic.veryActual, savePurposeGenetic.targetID, rеserve, true)
	   			// передать результат
	   			mentalInfoStruct.mentalPurposeID = 0 //purposeID
	   	}
	      ////////////////////// Цели - поиск в Правилах TODO !!!!!
	*/

	/*   оставляю как пример расшифровки эмоции
	// определить эмоциональный контекст по newEmotionID или прямо по CurrentEmotionReception
	contextArr:=CurrentEmotionReception.BaseIDarr
	for i := 0; i < len(contextArr); i++ {
		switch contextArr[i]{
		case 1:	//Пищевой	- Пищевое поведение, восполнение энергии, на что тратится время и тормозятся антагонистические стили поведения.

		case 2:	//Поиск	- Поисковое поведение, любопытство. Обследование объекта внимания, поиск новых возможностей.

		case 3:	//Игра	- Игровое поведение - отработка опыта в облегченных ситуациях или при обучении.

		case 4:	//Гон	- Половое поведение. Тормозятся антагонистические стили

		case 5:	//Защита	- Оборонительные поведение для явных признаков угрозы или плохом состоянии.

		case 6:	//Лень	- Апатия в благополучном или безысходном состоянии.

		case 7:	//Ступор	- Оцепенелость при непреодолимой опастbase_context_activnostности или когда нет мотивации при благополучии или отсуствии любых возможностей для активного поведения.

		case 8:	//Страх	- Осторожность при признаках опасной ситуации.

		case 9:	//Агрессия	- Агрессивное поведение для признаков легкой добычи или защиты (иногда - при плохом состоянии).

		case 10: //Злость	- Безжалостность в случае низкой оценки .

		case 11: //Доброта	- Альтруистическое поведение.

		case 12: //Сон - Состояние сна. Освобождение стрессового состояния. Реконструкция необработанной информации.

		}
	}
	*/

	// ? TODO мозжечковые рефлексы - самый первый уровень осознания - подгонка действий под заданную Цель?

	/* ДОМИНАНТА - осознаваемая цель - решение доминирующей осознанно сформированной проблемы
	Может по весу значимости перекрывать текущую Цель mentalInfoStruct.mentalPurposeID
	*/
	if EvolushnStage > 4 { // главное - Доминанта нерешенной пробелмы
		//dominfntaID := getMainDominanta(CurrentEmotionReception.ID)
		//if dominfntaID > 0 {
		getCurrentDominant()
		if CurrentProblemDominanta != nil {
			// TODO: использовать Цель Доминанты и определить по ней mentalInfoStruct.mentalPurposeID= !!!!
			//TODO:  ?sкакой должна быть цель Доминанты, чтобы по ней можно было определить mentalInfoStruct.mentalPurposeID

			return true
		}
	}

	// ?????????????????????????????
	//	createAndRunPurposeAutomatizm()

	return false
}

// //////////////////////////////////
func createAndRunPurposeAutomatizm() {
	if mentalInfoStruct.mentalPurposeID == 0 {
		return
	}
	// создать мент.автоматизм приоизвольной активацции ментальной цели
	/*	actImgID,_:=CreateNewlastMentalActionsImagesID(0,3,mentalInfoStruct.mentalPurposeID,true)

		id, matmzm := createMentalAutomatizmID(0, actImgID, 1)
		if id >0 {
			// запустить мент.автоматизм
			RunMentalAutomatizm(matmzm)

		}*/
}

//////////////////////////////////////////////////////////

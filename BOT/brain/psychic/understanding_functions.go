/*  функции для  consciousnessThinking(c *cycleInfo)

вынесены коды из блоков consciousnessThinking в отдельные функции чтобы было все видно

Функции обслуживают процесс в consciousnessThinking(c *cycleInfo)
они возвращают true если consciousnessThinking нужно прервать в этом месте
или false если нужно, чтобы consciousnessThinking продолжила далее процесс.

////////////////////////////////////////////////////
func xxxxxProcess(cycle *cycleInfo)bool{

	return false
}
/////////////////////////////////////////////////////
*/

package psychic

import (
	wordSensor "BOT/brain/words_sensor"
	"BOT/lib"
	"strconv"
)

//////////////////////////////////////////////////////////////////

// ////////////////////////////////// ВАЖНЫЙ ОБЪЕКТ ВНИМАНИЯ
func extremImportanceObjectProcess(cycle *cycleInfo) bool {
	cycle.lastProcessID = "extremImportanceObjectProcess"
	/////////////////  запись в ЛОГ: вытащить *importance по extremImportanceObject.extremObjID
	//oi,ok:=importanceFromID[extremImportanceObject.extremObjID]
	oi := getExtremObjFromID(extremImportanceObject.objID)
	if oi == nil {
		cycle.log += "НЕПРАВИЛЬНЫЙ ОБЪЕКТ ВНИМАНИЯ " + strconv.Itoa(extremImportanceObject.objID) + "...."
	}
	////////////////////

	if oi != nil && oi.extremVal < 3 { // проблемный объект с отрицательным эффектом - ОСНОВА ТВОРЧЕСКОГО ПОИСКА
		// текущий значимый объект внимания с отрицательным эффектом extremImportance.extremVal
		problemExtremImportanceObject = extremImportanceObject
	}
	// не опасно
	if !CurrentInformationEnvironment.veryActualSituation && !CurrentInformationEnvironment.danger {
		if !cycle.dreaming { // нет режима пассивного мышления
			if infoFunc2(cycle) {
				return true
			}
		}

	} else { // ОПАСНО
		//Срочно найти какое-то подходящее действие, получить mentalInfoStruct.runMotorAtmzmID и запустить его
		if infoFunc21(cycle) {
			return true
		} else {
			/* раз не решено в infoFunc21
			Выбрать о чем подумать (НУЖНО ПРИДУМАТЬ ЧТО-ТО НОВОЕ) в infoFunc2(), с учетом предыдущих шагов цикла.
			запустить базовую функцию infoFunc2() поиска по имеющейся информации
			пока не будет запущен мот.автоматизм infoFunc17()
			*/
			if infoFunc2(cycle) {
				return true
			}
		}
	}

	return false
}

/////////////////////////////////////////////////////

////////////////////////////////////////////////////
/* ОБЫЧНАЯ СИТУАЦИЯ, О ЧЕМ ДУМАТЬ
    Выбрать о чем подумать (НУЖНО ПРИДУМАТЬ ЧТО-ТО НОВОЕ) в infoFunc2(), с учетом предыдущих шагов цикла.
    запустить базовую функцию infoFunc2() поиска по имеющейся информации
    пока не будет запущен мот.автоматизм infoFunc17()
   		infoFunc2(cycle)

func usualThinkProcess(cycle *cycleInfo) bool {
	cycle.lastProcessID = "usualThinkProcess"

	// если сменилась тема или нет цели
	//Ментальное определение ближайшей Цели в данной ситуации по текущей теме

	///////////////////////////////////////////
	//детекция осмысленного ленивого состояния - все еще нет цели и нет доминанты
	if isIdleness() {
		if !cycle.dreaming {
			idlenessType = 2
			cycle.dreaming = true
			cycle.log += "Детекция осмысленного ленивого состояния - все еще нет цели и нет доминанты<br>"
			processingFreeState(false) // как во сне - обработка структур в свободном состоянии может быть долгой
			//чтобы остановить: processingFreeState(true)
			endDereamsCycles() // погасить все циклы дремы
		} //if cycle.dreaming!=1{
		return true
	}

	//if stimulCount > 1{//не было моторного ответа на прошлый стимул, а уже последовавл новый
	if isConfusion { // определяется в func consciousnessElementary()
		cycle.log += "Обработка конфуза: не было моторного ответа на прошлый стимул, а уже последовавл новый<br>"
		isConfusion = false
		interruptMentalWork(cycle.ID) // запомнить чтобы вернуться
		//Новая тема: Непонимание с небольшим весом
		runNewTheme(10, 1)
		cycle.log += "Новая Тема: Непонимание с небольшим весом<br>"
	}

	///////////////////////////////////

	//if currentInfoStructId!=5 { //если уже идет цикл мышления с функцией 5, то не прерывать это мышление
	//Выбрать о чем подумать (НУЖНО ПРИДУМАТЬ ЧТО-ТО НОВОЕ) в infoFunc2(), с учетом предыдущих шагов цикла.
	//запустить базовую функцию infoFunc2() поиска по имеющейся информации
	//пока не будет запущен мот.автоматизм infoFunc17()
	if infoFunc2(cycle) {
		return true
	}

	// по информации или продолжить цикл мышления или остановить после запуска моторного автоматизма
	if mentalInfoStruct.motorAtmzmID > 0 && mentalInfoStruct.motorAtmzmID == wasMentalRunMotorAtmzmID { // был запущен моторный автоматизм
		mentalInfoStruct.motorAtmzmID = 0
		wasMentalRunMotorAtmzmID = 0
		// должен прерываться цикл т.к. был запущен моторный автоматизм
		return true
	} else { // продолжить цикл мышления
		//МЕНТАЛЬНЫЙ АВТОМАТИЗМ - выдать цепочку инфо-фукнций из опыта мент.эпиз.памяти
		//которая потом будет использовать для последовательного запуска инфо-функций в findInfoIdFromExperience()
		mentAtmzmActualFuncs := getFavoritInfoFunc()
		if mentAtmzmActualFuncs != nil && len(mentAtmzmActualFuncs) > 0 {
			if mentAtmzmProcess(cycle) {
				return true
			}
		}
	}

	return false
}
*/
/////////////////////////////////////////////////////

// ////////////// состояние лени
func lenessProcess(cycle *cycleInfo) bool {
	cycle.lastProcessID = "lenessProcess"
	idlenessType = 1 // гомео-лень, т.е. непроизвольная
	//если есть проблемный, рвущещейся на выполнение автоматизма
	if mentalInfoStruct.motorAtmzmID > 0 {
		return false // даже если лень, то нужно что-то делать, поэто - смотреть далее
	}

	//если нет проблемного, рвущегося на выполнение автоматизма
	if mentalInfoStruct.motorAtmzmID == 0 {
		cycle.dreaming = false
		endAllCycles() //закончить ВСЕ циклы мышления
		cycle.log += "Состояние осоловелости<br>"
		// пофиг все, можно лениво обрабатывать накопившиеся структуры, реагирование - на уровне - до EvolushnStage < 4

		processingFreeState(false) // как во сне - обработка структур в свободном состоянии может быть долгой
		//чтобы остановить: processingFreeState(true)

		return true
	} //if mentalInfoStruct.motorAtmzmID==0{

	return false
}

//////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////

// //////////////////////////////////////////////////
var blockingNewProblemTryCount = false // от быстрого увеличения счетчика проблем в цикле, сбрасывается в func consciousnessElementary()

// только если ВАЖНАЯ ИЛИ ОПАСНАЯ СИТУАЦИЯ  не в дреме Блок повышенной акутальности
func dangerActualProcess(cycle *cycleInfo) bool {
	cycle.lastProcessID = "dangerActualProcess"
	// некогда думать, нужно что-то срочно решать

	/* newEpisodeMemory(0,0)// для данной detectedActiveLastProblemNodID пишем кадр без Правила */
	// написать нулевое Правило:
	saveNewEpisodic(curActions.ID, 0, 0, 0)

	/* увеличить счетчик нерешенной проблемы
	только для главного цикла и только в момент Стимула с пульта
	чтобы не молотило с каждым циклом.
	*/
	if cycle.isMainCycle && !blockingNewProblemTryCount {
		blockingNewProblemTryCount = true // хорош
		cycle.log += "Увеличен счетчик нерешенной проблемы<br>"
		lib.MapCheckWrite(MapGwardProblemTryCount)
		problemTryCount[detectedActiveLastProblemNodID]++
		lib.MapFree(MapGwardProblemTryCount)
		if EvolushnStage > 4 { //5-я сталия развития -
			// TODO доминанта может быть открыта произвольно сразу - тут нужно происследовать...
			lib.MapCheck(MapGwardProblemTryCount)
			if problemTryCount[detectedActiveLastProblemNodID] > 3 && CurrentInformationEnvironment.veryActualSituation {
				// создать доминанту
				createDominanta(cycle)
				return true
			}
		}
	}

	infoFunc25(cycle) //Посмотреть условия чтобы проявить инициативу

	// нет лени и нет штатного автоматизма
	if idlenessType == 0 && !mentalInfoStruct.noStaffAtmzmID {
		// Срочно найти какое-то подходящее действие, получить mentalInfoStruct.runMotorAtmzmID и запустить его
		if infoFunc21(cycle) { // если найдено, то останов цикла
			return true
		}
	}

	/*МЕНТАЛЬНЫЙ АВТОМАТИЗМ - выдать цепочку инфо-фукнций из опыта мент.эпиз.памяти
	которая потом будет использовать для последовательного запуска инфо-функций в findInfoIdFromExperience()
	*/
	mentAtmzmActualFuncs := getFavoritInfoFunc()
	if mentAtmzmActualFuncs != nil && len(mentAtmzmActualFuncs) > 0 {
		if mentAtmzmProcess(cycle) {
			return true
		}
	}

	return false
}

/////////////////////////////////////////////////////

// //////////////////////////////////////////////////
// творчество, мышление о Доминанте, фантазия - как функция интенсивного подбора ассоциаций
func dominantsProcess(cycle *cycleInfo) bool {
	cycle.lastProcessID = "dominantsProcess"
	// cycle.log+=" Отработка четвертого уровня осмысления с ID циикла "+strconv.Itoa(cycle.ID)+" мышления  "+onClickStr(cycle.ID,"show_cyckle","")+"<br>"
	// уже может быть запущена доминанта нерешенной проблемы
	// и тут - Стек отложенных дел.

	/* !!! при isRepressionStimulsNoise = true
	нет решения доминанты по аналогии т.к. заблокированы стимулы,
	Это - режим творческого решения по ментальным правилам или попытки новых действий.
	*/
	if isRepressionStimulsNoise {

	}

	if !CurrentInformationEnvironment.danger { // цикл решения доминирующей проблемы
		//выбор подходящей для detectedActiveLastProblemNodID доминанты ()
		//CurrentProblemDominanta=
		// TODO для решения по ID дерева проблем учитывать: getCurSituation(), getCurTheme() и getCurPurpose()
		infoFunc28(cycle)
	}
	return false
}

/////////////////////////////////////////////////////

////////////////////////////////////////////////////
/*  использовать ментальный автоматизм mentAtmzmActualFuncs:=getFavoritInfoFunc(), привязанный к текущей ветке дерева проблем

 */
func mentAtmzmProcess(cycle *cycleInfo) bool {
	cycle.lastProcessID = "mentAtmzmProcess"
	// есть ли полезный мент.автоматизм в узле дерева проблем, если есть сравнить полезность
	if mentAtmzmActualFuncs == nil || len(mentAtmzmActualFuncs) == 0 {
		return false
	}
	maStr := ""
	for i := 0; i < len(mentAtmzmActualFuncs); i++ {
		if i > 0 {
			maStr += ", "
		}
		maStr += strconv.Itoa(mentAtmzmActualFuncs[i])
	}
	cycle.log += "Отработка мент.автоматизма " + maStr + ", прикрепленного к ветве дерева проблем.<br>"

	/* с учетом модели понимания UnderstandingModel=make(map[int][]int) с индексом == ID объекта внимания extremImportance
	   а так же значимостей данного объета extremImportance
	   текущей Темы и Цели	(сужающий контекст)	т.е. учет правил только ID ветки проблем,
	   но при неудаче - смотреть и дургие ветки проблем.

	      (учти, что цикл мышления в автоматизме - то же самое, что ментальное правило!)
	      	чтобы не совершать снова работу по поиску подходящего ментального правила - использовать цикл из автоматизма.

	   Фактически мент.автоматизм выполняет роль поддержки Доминанты + пополняет массив tryActionArr.
	   Сначала просто пытается оптимизировать решение по инфе своего цикла,
	   потом открывается Доминанта и все временные решения заливаются для уже доимнанитного решения.
	*/

	/* TODO постепенно вбрасывать функции, а не хором
	for i:=0;i<len(mentAtmzmActualFuncs);i++ {
		funcID := mentAtmzmActualFuncs[i]
		if funcID > 0 {
			if runMentalFunctionID(cycle, funcID) {
				return true
			}
		}
	}*/

	// если найден новый вариант мышления, то УСТАНОВИТЬ НУЖНУЮ ИНФУ и запусть цикл с предполагаемым продолжением
	// в этом случае здесь -  return false

	return false
}

/////////////////////////////////////////////////////

///////////////////////////////////
/* по правилу найти или создать (в случае AmtzmNextString) автоматизм и запустить его
Если совершено действие - возвращает true
*/
func makeActionFromRooles(rules Rule) bool {
	// по правилу найти или создать (в случае AmtzmNextString) автоматизм и запустить его
	var atmzm *Automatizm

	conscienceStatus += "Найдено Правило по которому можно совершить действие.<br>"

	// выбираем Ответное действие из Правила чтобы повторить его
	//ai := ActionsImageArr[ta.Action]
	ai, ok := ReadeActionsImageArr(rules.Action)
	if ok {
		/////////////////////////////////
		if ai.Kind == 1 { // субъективный образ
			// Правило может быть неадекватным реальности
			//запускать автоматизм в опредеелнных условиях:
			if !CurrentInformationEnvironment.veryActualSituation &&
				!CurrentInformationEnvironment.danger &&
				//есть ли данные Базовые контексты 2 Поиск, 3 Игра, 4 Гон, 5 Защита - в составе текущей эмоции
				existsBaseContextFromEmotionID([]int{2, 3, 4, 5}, CurrentInformationEnvironment.PsyEmotionId) {
				// можно использовать субъективный образ
				purpose := getPurposeGenetic()
				purpose.actionID = ai
				atmzm = createAndRunAutomatizmFromPurpose(purpose)
			}
		} //if ai.Answer==1 { // субъективный образ
		/////////////////////////////////
	}

	if atmzm != nil {
		automatizmCorrection(atmzm, rules.Effect, nil)

		// вытащить образ действий успешного автоматизма и попробовать решить подходящую по аналогии Домимнату
		checkRelevantAction(curActions.ID, atmzm.ActionsImageID, atmzm.Usefulness)
		mentalInfoStruct.motorAtmzmID = 0 // сброс рассматривания автоматизма
		mentalInfoStruct.noStaffAtmzmID = false
		automatizmStatus = 0
		// если было размышление, то оно не прерывается в своем рекурсивном проходе
		conscienceStatus += "Запущен альтернативный штатному автоматизм, найденный в Правилах.<br>"
		atmtzmActualTreeNode = atmzm
		motorActionEffect = 1 // чтобы не запустилась infoFunc13
		return true           // блокировать рвущийся автоматизм т.к. запущен другой вместо него
	}

	return false
}

//////////////////////////////////////////////////

/*
	Для текущей фразы сначала смотрим есть ли для данных условий (BaseID + EmotionID) автоматзмы для известных слов фразы.

Если нет, то по Правилам выбираются подходяшие дейстивия.
Если найдены действия, то создается автоматизм, который прикрепляется к узлу дерева автоматизмов branchID.
В случае создания автоматизма возвращает autmtzm.ID иначе - 0.
Если есть автоматизмы на "привет" и на "как дела?" то на "привет как дела?" должен сформироваться автоматизм из их действий.
Если есть автоматизмы на "привет", но нет автоматизма на "как дела?" то выдаст автоматизм на "привет"

Работает с произвольным branchID, а не только при активации дерева,
так что функцию можно испольавтоматизмы на "привет"зовать для мыслительной произвольности
*/
func tryCreateAnswerForPhrese(limitActions int, branchID int) (int, *Automatizm) {
	wIDarr := wordSensor.CurretWordsIDarr // // массив wordID слов, вместо нераспознанных -1
	if wIDarr == nil {
		return 0, nil
	}
	//выбрать только известные слова
	var wrID []int
	for i := 0; i < len(wIDarr); i++ {
		if i > 10 { // смотрим не более, чем по 10 слов для данного уровня
			break
		}
		if wIDarr[i] > 0 {
			wrID = append(wrID, wIDarr[i])
		}
	}
	wIDarr = wrID
	// сформировать последовательность известных фраз из этих слов
	var phraseArr []int // знакомые фразы для поиска в дереве автоматизмов
	var tmpArr []int    // набор слов для распознавания фразы
	for n := 0; n < len(wIDarr); n++ {
		tmpArr = nil
		// сдвигаем на следующее слово
		for i := n; i < len(wIDarr); i++ {
			tmpArr = append(tmpArr, wIDarr[i])
			phraseID := wordSensor.GetExistsPraseIDFromWordArr(tmpArr) // распознаватель фразы без записи нового
			if phraseID > 0 {
				phraseArr = append(phraseArr, phraseID)
			}
		}
	}
	if phraseArr == nil {
		return 0, nil
	}

	////////////////  далее - последовательность сопособов получить autmtzm.ID

	/*1 способ: выбрать для данных условий (BaseID + EmotionID) автоматзмы для известных фраз
	 */
	emotionNode := getNodeFromLevel(2, branchID) // вернуться назад до узла эмоций и оттуда смотреть все
	if emotionNode != nil {
		/* начать поиск с данного узла по распознаннм фразам поиском по дереву автоматизмов
		Если в проходе есть автоматизм, то он учитывается.
		*/
		var aArr []int // для каждого известного слова выбрать ID действий
		var atmtzmFirst *Automatizm
		for i := 0; i < len(phraseArr); i++ {
			atmtzm := getAtmtzmFromNodesFrase(emotionNode, phraseArr[i])
			if atmtzm != nil {
				if !lib.ExistsValInArr(aArr, atmtzm.ActionsImageID) { // убираем дублеры
					aArr = append(aArr, atmtzm.ActionsImageID)
				}
				if atmtzmFirst == nil {
					atmtzmFirst = atmtzm
				}
			}
		}

		if aArr != nil && atmtzmFirst != nil {
			// для случая, когда во фразе нашелся фрагмент, на который есть автоматизм
			if len(aArr) == 1 {
				if isUnrecognizedPhraseFromAtmtzmTreeActivation { // просто выдаем автоматизм потому, что AutomatizmTreeFromID[branchID].PhraseID==0 - создаст кривой автоматизм без действий и фразы
					return atmtzmFirst.ID, atmtzmFirst
				}
				aID, atzm := createNewAutomatizmID(0, branchID, atmtzmFirst.ActionsImageID, true)
				if atzm != nil {
					return aID, atzm
				}
			} else {
				//создать автоматизм из цепочки действий, если действий >1, то - c atzm.NextID
				aID, atzm := createAutomatizmFromNextString(aArr, branchID)
				if atzm != nil {
					return aID, atzm
				}
			}
		}
	}
	//.......................

	// phraseArr=[]int{132,1229} // проверка
	//2 способ: выбрать действия из позитивных правил по каждой фразе
	var rwArr []int // для каждого известной фразы выбрать ID действий

	for i := 0; i < len(phraseArr); i++ {
		// позитивные правила для данного ID слова
		rArr := getWellRoolesFromPhraseId(phraseArr[i], true) // для текущих условий
		if rArr == nil {
			rArr = getWellRoolesFromPhraseId(phraseArr[i], false) // без учета условий
		}
		if rArr == nil {
			continue
		}
		// выбрать самые эффективные
		max := 0 // наяиная с эффективности 1
		actID := 0
		for n := 0; n < len(rArr); n++ {
			if rArr[n].Effect > max {
				max = rArr[n].Effect
				actID = rArr[n].Action
			}
		}
		// убираем дублеры
		if !lib.ExistsValInArr(rwArr, actID) {
			rwArr = append(rwArr, actID)
		}
	}

	// последовательность действий для автомтаизма
	var aArr []int
	for i := 0; i < len(rwArr); i++ {
		if rwArr[i] == 0 {
			continue
		}
		if i > limitActions {
			break
		}
		aArr = append(aArr, rwArr[i])
	}
	// создать цепочку действий
	if aArr == nil {
		return 0, nil
	}
	//id,_:=createAmtzmNextStringID(0,aArr,true)// создать образ цепочки типа AmtzmNextString

	//создать автоматизм из цепочки действий
	aID, atzm := createAutomatizmFromNextString(aArr, branchID)
	if atzm != nil {
		return aID, atzm
	}

	return 0, nil
}

/////////////////////////////////////////////////////

/*
	прогноз по действию автоматизма из последнего кадра эпиз.памяти

Возвращает эффект последнего позитивного правила в группе исторических кадров.
Пример:
Оператор: «ты кто?»
Beast: «бот»
Оператор: «уверен?»
Beast: «да»
Позитив группового правила

Оператор: «кто ты?»
Beast: «Бог»
Оператор: «уверен?»
Beast: «да»
Негатив группового правила

Казалось бы при проверке группового правила учитывать предыдущий диалог
чтобы при "уверен?" проверялось только то правило, где есть предыдущий диалог.
Но при реальной проверки диалог, который начинается на "уверен?" (даже повторенный 2 раза подряд)
всегда дает эффект == 0 и выполняется автоматизм "да".
Так что пока нет необходимости собирать все групповые правила с заданным диалогом и выбирать наиболее подходящее.
Но это еще не точно, в каких-то случаях такое может приводить к неточности, это нужно проверить TODO
*/
func getPrognoze(atmtzm *Automatizm) (int, int) {
	if atmtzm == nil {
		lib.TodoPanic("Нулевое значение параметра в func getPrognoze")
	}
	if curActiveActionsID == 0 { // стимул от гомеостаза, а не от оператора
		return 0, 0
	}
	actionID := atmtzm.ActionsImageID
	//если trigger == actionID ЗНАЧИТ БОТ ЗЕРКАЛИТ стимул
	trigger := curActiveActionsID // активация дерева оператором

	// прямое предсказание последствий выполнения действия автоматизма, т.е. эффект trigger, actionID
	accuracy, effect := getPrognoseFromAutmtzmAction(trigger, actionID)
	if accuracy == 1 { //точное предсказание для действия
		if effect < 0 { // негативное предсказание
			// но потом будет позитив?
			accuracy0, effect0 := positiveFromActionAfterStimul(curActiveActionsID, actionID)
			// если конечный эффект превышает прямой негативный
			if accuracy0 > 0 && effect0 > 0 && effect0 > -effect { //есть превышающее позитивное предсказание
				return accuracy0, effect0
			}
		}
		if effect > 0 { // позтивное предсказание
			// но потом будет негатив?
			accuracy0, effect0 := negativeFromActionAfterStimul(curActiveActionsID, actionID)
			// если конечный эффект превышает прямой негативный
			if accuracy0 > 0 && effect0 < 0 && -effect0 > effect { //есть превышающее негативное предсказание
				return accuracy0, effect0
			}
		}
		return accuracy, effect
	}
	if accuracy == 2 { //менее точное предсказание для действия
		//менее точное предсказание для действия, совершенное после Стимула curStimulImageID в данных условиях
		// нет смысла смотреть, чем кончится цепочка кадров
		return accuracy, effect
	}
	if accuracy == 3 {
		//неточное предсказание для действия при любом стимуле и любых условиях
		// дает ли действие негативный эффект ()
		accuracy, effect = negativeFromActionAfterStimul(trigger, actionID)
		if accuracy > 0 {
			return accuracy, effect
		}
	}

	return 0, 0
}

/* СТАРАЯ ВЕРСИЯ
func getPrognoze(atmtzm *Automatizm) (int, int) {
	if atmtzm == nil {
		lib.TodoPanic("Нулевое значение параметра в func getPrognoze")
	}
	actionID := atmtzm.ActionsImageID
	//	accuracy,effect:=positiveFromActionAfterStimul(curActiveActionsID,actionID)
	//Алгоритм:
	 // Поиск цепочек, заканчивающихся большим негативом (превышающим позитив промежуточных)
	 // Если не найдены цепочки, то поиск негатива в правилах с учетом эмоций и Стимула Trigger
	 // Если нет таких, то поиск правил с участием Action
	 // т.е. в первую очередь смотрим нет ли негатива и если нет, то автоматизм может запускаться.
accuracy, effect := getPrognoseFromAutmtzmAction(curActiveActionsID, actionID)
if accuracy == 1 {
//точное предсказание для действия, которое может привести к негативу, но потом будет позитив
}
if accuracy == 2 {
//точное предсказание для действия, совершенное после Стимула curStimulImageID в данных условиях
}
if accuracy == 3 {
//статистически самое позитивное предсказание для действия при любом стимуле и любых условиях
}
if accuracy > 0 {
return accuracy, effect
}
/////////// дает ди действие негативный эффект
accuracy, effect = negativeFromActionAfterStimul(curActiveActionsID, actionID)
if accuracy > 0 {
return accuracy, effect
}

return 0, 0
}
*/

///////////////////////////////////////////////////////////////////////////////

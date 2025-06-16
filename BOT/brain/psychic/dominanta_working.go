/*
Работа с Доминантой


Неудовлетворенность существующим реализуется как в виде активности темы ID=9,
так и использованием уже закрытых и открытых доминант для продолжения поиска ассоциаций.
*/

package psychic

import (
	"BOT/lib"
	"strconv"
)

////////////////////////////////////////////
/* создать доминанту
Без extremImportanceObject не создается
*/
func createDominanta(c *cycleInfo) {
	if c.ID == 0 {
		return
	}
	if EvolushnStage < 5 {
		return
	}
	if extremImportanceObject == nil {
		return
	}
	weight := 0 // насколько важна проблема - чисто оценочно, Доминанты не конкурируют по весу значимости
	weight = extremImportanceObject.extremVal

	st := getCurSituation()
	if st != nil {
		switch st.SituationType {
		case 5:
			weight += 1 //оператор не прореагировал на действия в течение периода ожидания - игнорирует
		case 12:
			weight += 1 //Плохое
		case 15:
			weight += 2 //Агрессивное
		case 16:
			weight += 1 //Защитное
			//оператор нажал кнопку
		case 21:
			weight += 1 //Непонятно
		case 23:
			weight += 3 //Наказать
		case 29:
			weight += 1 //Игнорировать
		case 30:
			weight += 3 //Сделать больно
		}
	}
	tm := getCurTheme()
	if tm != nil {
		switch tm.Type {
		case 1:
			weight += 2 //"Негативный эффект моторного автоматизма"
		case 2:
			weight += 1 //"Негативный эффект ментального автоматизма"
		case 3:
			weight += 2 //"Состояние Плохо"
		case 7:
			weight += 1 //"Игнорирование оператором"
		case 9:
			weight += 2 //"Неудовлетворенность существущим"
		case 10:
			weight += 1 //"Непонимание"
		case 14:
			weight += 2 //"Страх"
		}
	}

	birthTime := int(LifeTime / (3600 * 24)) // время рождения
	endTime := 0                             // TODO нужно определить, но время актуальности м.б. не опредеелно ==0

	// Определить целевое действие, которые желательно совершить с объектом objectID:
	//Answer int // 1 - субъективное предположение
	actArr := []int{}    //  TODO нужно определить
	зhraseArr := []int{} //  TODO нужно определить
	toneID := 0          //  TODO нужно определить
	moodID := 0          //  TODO нужно определить
	targetActionID, _ := CreateNewlastActionsImageID(0, 1, actArr, зhraseArr, toneID, moodID, true)

	typeTargetAction := 3 // TODO нужно определить

	needExecuteTargetAction := 1 //TODO: совершить действие targetActionID, 0- избегать действия targetActionID

	id, _ := createNewDominanta(0, detectedActiveLastProblemNodID,
		weight,
		endTime,
		extremImportanceObject.objID,
		targetActionID,
		typeTargetAction,
		needExecuteTargetAction,
		birthTime,
		-1,
		0, true)

	c.log += "Активирована Доминанта id=<b><span style='cursor:pointer;color:blue' onClick='show_dominant(" + strconv.Itoa(id) + ")'>" + strconv.Itoa(id) + "</span><br>"

	//перезапуск осмысления
	//reloadConsciousness(c,0)
	return
}

////////////////////////////////////////////

// выбрать подходящую Доминанту, с присвоением
func getCurrentDominant() {
	if EvolushnStage < 5 || extremImportanceObject == nil {
		/* Если нет экстремального объекта, то все равно выбрать доминанту - уже решенную
		   		для реализации неудовлетворенности существующим
		   использованием уже закрытых или открытых доминант для продолжения поиска ассоциаций
		*/
		var permanentDissatisfaction = false // true - особь с постоянным зудом неудовлетворенности - TODO может изменяться от обстоятельств
		if permanentDissatisfaction {
			getClosedCurrentDominant()
			return
		}
		tm := getCurTheme()
		if tm.Type == 9 { //"Неудовлетворенность существущим"
			getClosedCurrentDominant()
		}
		return
	}

	for _, v := range DominantaProblem {
		if v == nil {
			continue
		}
		if v.isSuccess == 3 || v.isSuccess == 4 { // не смотреть закрытые проблемы
			continue
		}
		objID := extremImportanceObject.objID

		if v.problemTreeID == detectedActiveLastProblemNodID && v.objectID == objID {
			CurrentProblemDominanta = v
			//extremImportanceObject.extremObjID=v.objectID  уже итак есть т.к. доминанта ищется по уже выявленному extremImportanceObject
			CurrentInformationEnvironment.DominantaID = v.ID

			return
		}
	}

	CurrentProblemDominanta = nil
}

// подобрать любую наиболее подходящую к условиям доминанту
func getClosedCurrentDominant() {
	for _, v := range DominantaProblem {
		if v == nil {
			continue
		}
		if v.problemTreeID == detectedActiveLastProblemNodID {
			CurrentProblemDominanta = v
			CurrentInformationEnvironment.DominantaID = v.ID
			return
		}
	}

	return
}

/////////////////////////////////////////////

/*
	Проверка, насколько данный actionsImageID может быть решением для CurrentProblemDominanta.

Если actionsImageID привел к успеху, то он может считатья решением Доминанты.

А если условия != detectedActiveLastProblemNodID , то он может быть принят как текущий curTryActionsID
и т.о. предположительное решение может быть открыто (эавристика)
по аналогии с любым наблюдаемым или совершаемым actionsImageID
и возникает "озарение" с последующей проверкой и запуском пробного действия.

Аргумент successAImgID - только успешные или авторитарно уверенные actionsImageID.
stimulID - это ID объект типа ActionsImage - действие с Пульта

Возвращает:
0 - нет подходящего,
1 - доминанта решена и завершена, можно использовать такое решение хоть сразу, хоть потом
>1 до 3 включительно - найдено пробное решение для проверок, возвращаемое число увеличивается со степенью непохожести условий.

Выбирает первую подходящую как аналог доминанту.
Из-за перебора доминант для сравненеия м.б. тяжелым процессом...
*/
func checkRelevantAction(stimulID int, successAImgID int, effect int) int {
	if EvolushnStage < 5 {
		return 0
	}
	// найти нерешенную диминанту, достаточно подходящую для данного successAImgID
	dominanta := getRelevantDominanta(stimulID)
	if dominanta == nil {
		return 0
	}

	objID := dominanta.objectID

	// прям полностью совпадают условия.
	if detectedActiveLastProblemNodID == dominanta.problemTreeID &&
		dominanta.objectID == objID {

		// является стимул stimulID решающим действием

		/* TODO переделать все с учетом
		Dominanta.objectID
		Dominanta.targetActionID
		Dominanta.typeTargetAction
		Dominanta.needExecuteTargetAction
		*/

		// доминанта решена
		writeDomTryActionArr(successAImgID, effect, objID, dominanta)
		// Доминанта решена, завершить ее и эвристически озариться.

		//Доминанта закрывается только когда после выполнения действий будет получен позитивный эффект
		dominanta.isSuccess = 2
		toConsciousHeuristics(1, successAImgID, dominanta)
		return 1
	}
	////////////////////////////////////////////

	//pT:=ProblemTreeNodeFromID[dominanta.problemTreeID]
	pT, ok := ReadeProblemTreeNodeFromID(dominanta.problemTreeID)
	if !ok {
		return 0
	}

	// совпадают detectedActiveLastNodID и detectedActiveLastUnderstandingNodID
	if detectedActiveLastNodID == pT.autTreeID && detectedActiveLastUnderstandingNodID == pT.situationTreeID &&
		dominanta.objectID == objID { // принять по ассоциации
		writeDomTryActionArr(successAImgID, effect, objID, dominanta)
		// Доминанта решена, завершить ее и эвристически озариться.
		dominanta.isSuccess = 1
		toConsciousHeuristics(2, successAImgID, dominanta)
		return 2
	}
	///////////////////////////////////////////

	// совпадают detectedActiveLastNodID
	if detectedActiveLastNodID == pT.autTreeID &&
		dominanta.objectID == objID { // принять по ассоциации
		writeDomTryActionArr(successAImgID, effect, objID, dominanta)
		// Доминанта решена, завершить ее и эвристически озариться.
		dominanta.isSuccess = 1
		toConsciousHeuristics(3, successAImgID, dominanta)
		return 3
	}
	///////////////////////////////////////////

	return 0
}

// ////////////////
func getRelevantDominanta(stimulID int) *Dominanta {

	//aiObj:=ActionsImageArr[stimulID]
	aiObj, ok := ReadeActionsImageArr(stimulID)
	if ok {
		return nil
	}
	for _, v := range DominantaProblem {
		if v == nil {
			continue
		}
		if v.isSuccess == 3 || v.isSuccess == 4 { // смотреть только НЕ закрытые доминанты
			continue
		}
		//if v.problemTreeID == detectedActiveLastProblemNodID{ НЕТ, ИНАЧЕ БУДЕТ НЕ АНАЛОГИЯ, А ПОЛНОЕ СОВПАДЕНИЕ
		if isAnalogAction(v, stimulID, aiObj) != nil {
			return v
		}
		//		}
	}

	return nil
}

// если подходит акция по аналогии, то вернуть такую доминанту
func isAnalogAction(v *Dominanta, stimulID int, aiObj *ActionsImage) *Dominanta {
	extrObj := getExtremObjFromID(v.objectID)
	if extrObj == nil {
		return nil
	}

	if stimulID == v.objectID {
		return v
	}
	return nil
}

// ///////////////////////
func writeDomTryActionArr(successAImgID int, effect int, objID int, dominanta *Dominanta) {
	dAiIDarr := getDomTryActionIDArr(dominanta.ID) //массив *tryAction для ID доминанты
	ok, tryActionsKey := lib.IndexValInArr(dAiIDarr, objID)
	if !ok { //  дополнить
		tryActionsKey, _ = addNewTryAction(0, dominanta.ID, successAImgID, effect, true)
	}
	// Доминанта предварительно решена, нужно эвристически озариться.
	dominanta.tryActionsKey = tryActionsKey
}

//////////////////////////////////////////////

//////////////////////////////////////////////
/* Эвристически озариться найденным решением
success - 1 - доминанта решена, 2 - не точно, 3 - еще менее точно
*/
func toConsciousHeuristics(success int, successAImgID int, dominanta *Dominanta) {
	if EvolushnStage < 5 {
		return
	}
	CurrentProblemDominanta = dominanta
	mentalInfoStruct.DominantaID = dominanta.ID
	mentalInfoStruct.DominantSuccessAImgID = successAImgID
	mentalInfoStruct.DominantSuccessValue = success
	cycle := createNewCycleIteration() // начать новый цикл
	// т.к. уже не первый шаг, получить текущую Цель РАЗ ОЗАРИЛО, ТО _ ЦЕЛЬ БЫЛА И РЕШЕНА!
	//	infoFunc8(cycle) // Ментальное определение ближайшей Цели mentalInfoStruct.ThemeImageType и актиуивровать дерево проблем

	//	infoFunc19(cycle)//
	//	setAsMaimCycle(cycle.ID)//	сделать цикл главным
	insight(cycle)
	// и начать осмысление

}

////////////////////////////////////////

/*
	запускать решение закрытой доминанты (т.е. взять опыт из ранее решенных проблм) в подходящих условиях.

# ПРи вызове уже учтена стадия развития и отсуствие автоматизма

exactly==false - смотреть с isSuccess>1, exactly==true - только закрытые, точные доминанты

если подходящий автоматизм найден в успешной Доминанте то он будет запущен
*/
func runDominantaAction(c *cycleInfo, exactly bool) bool {
	if extremImportanceObject == nil {
		return false
	}
	for _, v := range DominantaProblem {
		if v == nil {
			continue
		}
		if exactly {
			if v.isSuccess != 3 { // смотреть только успешно закрытые доминанты
				continue
			}
		} else {
			if v.isSuccess < 2 { // смотреть только успешно закрытые доминанты
				continue
			}
		}

		objID := extremImportanceObject.objID

		if v.problemTreeID == detectedActiveLastProblemNodID && v.objectID == objID {
			dominanta := v
			// сделать мот.автоматизм и запустить его
			var actionsImage *ActionsImage // шаблон образа действия

			extrObj := getExtremObjFromID(dominanta.objectID)
			if extrObj == nil {
				return false
			}

			//	ai:=ActionsImageArr[extrObj.extremObjID]
			ai, ok := ReadeActionsImageArr(extrObj.objID)
			if ok {
				return false
			}
			_, actionsImage = CreateNewlastActionsImageID(0, 1, ai.ActID, ai.PhraseID, 0, 0, true)

			if actionsImage != nil {
				mentalInfoStruct.motorAtmzmID = actionsImage.ID
				c.log += "Подходящий автоматизм найден в успешной Доминанте и запущен.<br>"
				infoFunc17(c) // запустить автоматизм и завершить цикл осмысления
				return true
			}
		}
	}

	return false
}

//////////////////////////////////////////

// ///////////////////////////////////////
// закрыть доминанту, котора точно подходит
func checkGestalt(effect int) {
	// curStimulImageID - образ действий оператора перед Ответом
	if curStimulImageID == 0 {
		return
	}
	if curStimulImage.ActID == nil && curStimulImage.PhraseID == nil { // не учитывать с пустым Стимулом
		return
	}
	answerID := LastAutomatizmWeiting.ActionsImageID // ответный образ действий Beast
	if answerID == 0 {
		return
	}
	_, ok := ReadeActionsImageArr(answerID)
	if ok {
		return
	}
	if curStimulImageID == answerID {
		return
	}
	// поиск доминант
	for _, v := range DominantaProblem {
		if v == nil {
			continue
		}
		if v.isSuccess == 3 || v.isSuccess == 4 { // смотреть только НЕ закрытые доминанты
			continue
		}
		if v.problemTreeID != detectedActiveLastProblemNodID {
			continue
		}
		lib.MapCheck(MapGwardTryActionArr)
		aArr := tryActionArr[v.ID]
		// точное соотвествие доминанты:
		if v.objectID == curStimulImage.ID &&
			aArr[v.tryActionsKey].ID == answerID &&
			v.isSuccess == 2 {
			v.isSuccess = 3 // закрыть доминанту
			if v.weight < effect {
				v.weight = effect
			}
			return
		}
	}

}

/////////////////////////////////////////

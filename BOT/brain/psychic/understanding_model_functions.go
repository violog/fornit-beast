/*
Модели понимания - функции
*/
package psychic

////////////////////////////////////////////////////

/*
	поиск решения по Моделям понимания

Использовать модели объектов для нахождения действий ПО АНАЛОГИИ с теми действиями,
что были в эпиз.памяти для данного объекта.
Леонардо да Винчи смотрел на штукатурку чтобы в голову пришла идея.

При удачном поиске будет создан и запущен автоматизм.
*/
func tryModelsLevel(c *cycleInfo) bool {

	actBest, _ := bestActionFromExtremImportanceObject()
	if actBest != nil { // найдено лучшее
		//получить пользу из эффекта лучшего моторного действия: будет найден и запущен лучший мот.автоматизм aBest.ID
		return getBenefitFromEpizosMemory(c, actBest)
	}
	return false
}

//////////////////////////////////////////////////

/*
найти действие в данных условиях (bestActionFromExtremImportanceObject()).
которое лучше, чем текущий extremImportanceObject
*/
func bestActionFromExtremImportanceObject() (*ActionsImage, int) {
	if extremImportanceObject == nil {
		return nil, 0
	}
	getRulesModelExactly(extremImportanceObject.objID)
	if UnderstandingRulesModel == nil {
		getRulesModelApproximately(extremImportanceObject.objID)
	}
	if UnderstandingRulesModel == nil {
		return nil, 0
	}
	///////////
	var effectMax = 0
	var actBest *ActionsImage

	for i := 0; i < len(UnderstandingRulesModel); i++ {
		rm := UnderstandingRulesModel[i]
		eff := getWpower(rm.Effect, rm.Count)
		if eff > effectMax {
			//actBest=ActionsImageArr[rm.Action] // ответ с хорошим эффектом
			actN, ok := ReadeActionsImageArr(rm.Answer)
			if !ok {
				continue
			}
			effectMax = eff
			actBest = actN // ответ с хорошим эффектом
		}
	}
	return actBest, effectMax
}

///////////////////////////////////////////////////////////

/*
выбрать все позитивные образы действий из extremImportanceObject
*/
type bestActionArr struct {
	action *ActionsImage
	effect int
}

func bestActionArrFromExtremImportanceObject() []bestActionArr {
	if extremImportanceObject == nil {
		return nil
	}
	getRulesModelExactly(extremImportanceObject.objID)
	if UnderstandingRulesModel == nil {
		getRulesModelApproximately(extremImportanceObject.objID)
	}
	if UnderstandingRulesModel == nil {
		return nil
	}
	///////////
	var out []bestActionArr

	for i := 0; i < len(UnderstandingRulesModel); i++ {
		rm := UnderstandingRulesModel[i]
		if rm.Effect > 0 {
			//actBest=ActionsImageArr[rm.Action] // ответ с хорошим эффектом
			actN, ok := ReadeActionsImageArr(rm.Answer)
			if !ok {
				continue
			}
			out = append(out, bestActionArr{actN, rm.Effect}) // ответ с хорошим эффектом
		}
	}
	return out
}

///////////////////////////////////////////////////////////

/*
попробовать для случая негативного obj *extremImportance или emExtrem *EpisodeMemory
Из нескольких акций
собрать один образ действий типа ActionsImage - из составляющих элементов с позитивными эффектами цепочки Правил
*/
func createSintezActiveActionsFromExtremImportanceObject() (ActionsImage, bool) {
	var action ActionsImage
	aArr := bestActionArrFromExtremImportanceObject()
	if aArr == nil {
		return action, false
	}
	aCount := len(aArr)
	if aCount == 1 {
		return *aArr[0].action, true
	}
	/* собрать образ действий типа ActionsImage.ID из составляющих элементов,
	вытаскивая из всех образов составляющие и выбирая из лучших один сложный образ действия.
	*/
	beast := 0
	toneID := 0
	moodID := 0

	for i := 0; i < aCount; i++ {
		for j := 0; j < len(aArr[i].action.ActID); j++ {
			action.ActID = append(action.ActID, aArr[i].action.ActID[j])
		}
		for j := 0; j < len(aArr[i].action.ActID); j++ {
			if aArr[i].action.PhraseID != nil {
				action.PhraseID = append(action.PhraseID, aArr[i].action.PhraseID[j])
			}
		}
		if aArr[i].effect > beast {
			beast = aArr[i].effect
			toneID = aArr[i].action.ToneID
			moodID = aArr[i].action.MoodID
		}
	}
	action.ToneID = toneID
	action.MoodID = moodID

	return action, true
}

////////////////////////////////////////////////////////////////////

/*
	посмотреть максимальный эффект в Модели понимания

для данного образа действия типа ActionsImage
в данных условиях дерева проблем.
*/
func getExtremEffectFromModel(actImgID int) int {
	if actImgID == 0 || detectedActiveLastProblemNodID == 0 {
		return 0
	}

	getRulesModelFromActionExactly(actImgID)
	if UnderstandingRulesModel == nil {
		return 0
	}
	///////////
	var effectMax = 0

	for i := 0; i < len(UnderstandingRulesModel); i++ {
		rm := UnderstandingRulesModel[i]
		eff := getWpower(rm.Effect, rm.Count)
		if eff > effectMax {
			effectMax = eff
		}
	}
	if effectMax != 0 { // найдено лучшее
		return effectMax
	}
	return 0
}

/////////////////////////////////////////////////////////////////

/*
	плохой ли выбор actionsImageID для мовершения действий

true - не подходящий для данного detectedActiveLastProblemNodID

Проверка только для tryModelsLevel_1() с точным результатом
*/
func isBadActionsImagefromTryActionArr(actImgID int) bool {
	getRulesModelFromActionExactly(actImgID)
	if UnderstandingRulesModel == nil {
		return false // если нет еще опыта, будем считать, что пойдет
	}
	///////////
	for i := 0; i < len(UnderstandingRulesModel); i++ {
		rm := UnderstandingRulesModel[i]
		if rm.Effect < 0 {
			return true
		}
	}
	// если нет плохого эффекта, будем считать, что пойдет
	return false
}

///////////////////////////////////////////

/*
	посмотреть, насколько опасны последствия в Модели понимания

для данного образа действия типа ActionsImage
в данных условиях дерева проблем.
*/
func isDangerousModelAutomatism(actImgID int) int {
	if detectedActiveLastProblemNodID == 0 {
		return 0
	}

	getRulesModelFromActionExactly(actImgID)
	if UnderstandingRulesModel == nil {
		return 0
	}
	///////////
	var effectMin = 0

	for i := 0; i < len(UnderstandingRulesModel); i++ {
		rm := UnderstandingRulesModel[i]
		eff := getWpower(rm.Effect, rm.Count)
		if eff < effectMin {
			effectMin = eff
		}
	}
	/*так как эффект отрицательный потому, что негативный, переводим его в значение по модулю,
	  чтобы потом сравнивать не плохой с хорошим, а определять позитивный эфеект, ПРЕВЫШАЮЩИЙ негативный
	  только в этом случае стоит бездумно отдавать ему предпочтение, но только в прогнозе на 1 шаг:
	  разрешить сделать шаг назад, если потом СРАЗУ ЖЕ будет 2 шага вперед
	  более сложный анализ предполагается в более продвинутых функциях (хотя может быть сложнее и нет ничего)
	*/
	effectMin = -1 * effectMin
	return effectMin
}

/////////////////////////////////////////////////////////////////

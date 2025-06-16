/*  Корректировка автоматизма по результатам реакции
Блокирование, разблокирование автоматизмов

Блокирование - atmtzm.Usefulness<0 (такой не выполняется) или убирание из штатных.

Принципы.
1. изменение параметра atmtzm.Count в зависимоти от стадии развития: func EvolushnStageAtmzCount
В итоге, чтобы заблокировать автоматизм в 4 стадии, если он был закреплен во 2 стадии, нужно было несколько раз оценить его негативно.

2. в 4 версии надо видимо добавить новшество: изменять оценку автоматизма в зависимости от достижения/не достижения цели getMentalEffect,
причем с учетом значимости цели.
То есть здесь начинает влиять произвольная оценка значимости цели: чем она важнее, тем сильнее будет переоценка автоматизма.

3. блокировка должна быть ОЧЕНЬ КОНСЕРВАТИВНОЙ! потому как после негатива может возникнуть позитив (цепочка правил с конечным позитивом).
Животное очень не просто заставить что-то не делать.
Только очень сильный негатив (lastCommonDiffValue < -5) может блокировать автоматизм.

4. Разблокировка и переход в штатные привяхонного к ветке возможны только на уровне сознания,
если нет автоматизмов, а опыт позитива есть или ситуация очень скверная и нужно хоть что-то делать.
Для атасной ситуации применяется функция findAlntrnativeAtmtzm()
*/

package psychic

//////////////////////////////////////////////////////////////////

/*
зависимость шага надежности от стадии развития
чем ниже стадия - тем выше шаг. Это позволит закрепить от изменений
базовые автоматизмы, формируемые в начальных стадиях
*/
func EvolushnStageAtmzCount(stepType bool) int {
	var step = 1

	switch EvolushnStage {
	case 2:
		step = 5
	case 3:
		step = 2
	}
	if !stepType && step != 1 {
		step = step / 2
	}
	return step
}

////////////////////////////////////////////////////////////////////

/*
корректируется успешность автоматизма
lastCommonDiffValue - реакция на результат (эффект от -10 до +10)
wellIDarr - ID улучшившихся гомеопараметров

Здесь могут быть только штатные и с atmtzm.UsefulnessЮ=0, незаблокированные автоматизмы раз было совершено действие.
*/
func automatizmCorrection(atmtzm *Automatizm, lastCommonDiffValue int, wellIDarr []int) {
	if atmtzm == nil {
		return
	}

	var isCommon = false

	if atmtzm.BranchID > 1000000 { // это общий, привязанный не к ветке, а к действиям или словам
		isCommon = true
	}
	//коррекция эффекта в зависимости от достижения/не достижения цели
	if EvolushnStage > 3 {
		effectValuation := getMentalEffect(lastCommonDiffValue) // детектор достижения Цели
		wMentEffrt := 2                                         // вес эффекта достижения цели
		lastCommonDiffValue += wMentEffrt * effectValuation
		if lastCommonDiffValue != 0 { // иначе выдаст -10
			if lastCommonDiffValue > 10 {
				lastCommonDiffValue = 10
			}
			if lastCommonDiffValue < -10 {
				lastCommonDiffValue = -10
			}
		}
	}

	////////////// стало хуже
	if lastCommonDiffValue < 0 {
		PsyBaseMood = -1

		// это - игнорирующий автоматизм для предовтащения действия общего автоматизма в данной ветке?
		isIgnore := isIgnoreAutomatizm(atmtzm)
		if !isCommon && !isIgnore { // не для общего автоматизма и не для игнорирующего

			passivationAutomatizm(atmtzm, lastCommonDiffValue)

			if atmtzm.Usefulness < 0 {
				// очистить списки улучшения
				atmtzm.GomeoIdSuccesArr = nil
				if AutomatizmSuccessFromIdArr[atmtzm.ID] != nil {
					AutomatizmSuccessFromIdArr[atmtzm.ID] = nil
				}

				// убрать из штатного
				SetAutomatizmBelief(atmtzm, 0)
				if EvolushnStage > 3 {
					//в коллекцию неудач
					addNewTryAction(0, -detectedActiveLastProblemNodID, atmtzm.ActionsImageID, atmtzm.Usefulness, true)
					if atmtzm.NextID > 0 {
						// действия с цепочкой при плохом эффекте
						badEffectChainAtmtzm(atmtzm)
					}
				}
			}

		} else { // для общего автоматизма
			/*ОБЩИЕ АВТОМАТИЗМЫ НЕ КОРРЕКТИРУЮТСЯ негативом!
			Концепция общих автоматизмов. Они сформированы на основе общего шаблона рефлексов и, как и рефлексы,
			перекрываются автоматизмами конца активной ветки, т.е. имеющими образ Стимула.
			Это – первичная реакция на текущее сочетание контекстов (в данном случае – на эмоцию).
			У узла эмоции может быть бесконечное число образов действия
			и блокировка общего автоматизма лишает первичной реакции у всех их.
			Но у.рефлекс может быть заморожен, а общий автоматизм сейчас – нет.
			Это значит, то в случае блокирующих действий для данной ветки, необходимо запускать автоматизм бездействия,
			останавливающий все более низкоуровневое.
			При блокировке такого автоматизма опять НЕ БЛОКРУЕТСЯ.
			Наличие игнор.автоматизма в конце ветки для func consciousnessElementary() равноценно отсуствю автоматизма.
			*/
			getIgnoreAtmtzmToBrench(atmtzm.BranchID)
		}
	}
	//////////// стало лучше
	if lastCommonDiffValue > 0 {
		PsyBaseMood = 1
		// список гомео параметро, которые улучшило это действие
		if wellIDarr != nil {
			atmtzm.GomeoIdSuccesArr = wellIDarr // м.б. nil !!!! если нет таких явных действий
		}
		// пополняется список полезных автоматизмов
		AutomatizmSuccessFromIdArr[atmtzm.ID] = atmtzm

		// задать тип автоматизма, 2 - проверенный
		SetAutomatizmBelief(atmtzm, 2) // сделать автоматизм штатным

		if !isCommon { // не для общего автоматизма
			// повысить ранее плохой штатный автоматизм
			if atmtzm.Usefulness < lastCommonDiffValue {
				atmtzm.Usefulness = lastCommonDiffValue // позитивная надежность
			}
			atmtzm.Count += EvolushnStageAtmzCount(true) // добавляем счетчик позитивной надежности
		} else { // для общего автоматизма
			// привязать общий автоматизм к активной ветке
			linkCoomonAtmtzmToBrench(atmtzm)
		}
	}
	return
}

////////////////////////////

// ухудшение автоматизма негативом
func passivationAutomatizm(atmtzm *Automatizm, negativeValue int) {

	if atmtzm.Usefulness >= 0 {
		//зависимость шага надежности от стадии развития
		atmtzm.Count -= EvolushnStageAtmzCount(false)

		// Только очень сильный негатив или очень слабый автоматизм могут блокировать сразу. Или во второй стадии
		if negativeValue < -5 || (atmtzm.Usefulness == 0 && atmtzm.Count == 0) || (EvolushnStage == 2 && atmtzm.Count < 0) {
			atmtzm.Count = 1 // счетчик уверенности уже негативного автоматизма
			atmtzm.Usefulness = -1
		} else { // не очень сильный эффект
			// консервативное ухудшение рейтинга автоматизма, чтобы не сразу блокировался
			if atmtzm.Usefulness > 0 { // не понижать меньше, чем до 0
				atmtzm.Usefulness -= 1 //lastCommonDiffValue
				atmtzm.Count--
			}
			if atmtzm.Count < 0 {
				atmtzm.Count = 0
			}
		}
	} else { // уже негативный автоматизм
		atmtzm.Usefulness--
		atmtzm.Count++ // счетчик уверенности уже негативного автоматизма
	}
}

//////////////////////////////////////////////////////////////////////////

/*
срочно, не думая, найти аварийный автоматизм, хоть что-то из привязанных к ветке или новый
привязать, сделать штатным, даже с небольшим отрицательным Usefulness
*/
func findAlternativeAtmtzm() int {

	// список всех автоматизмов для ID узла Дерева
	aArr := GetMotorsAutomatizmListFromTreeId(detectedActiveLastNodID)
	if aArr != nil {
		var usefulCount = -100 // полезность, выбрать наилучшую
		var autmtzmID = 0
		var autmtzm *Automatizm
		if aArr != nil {
			for i := 0; i < len(aArr); i++ {
				if aArr[i].Belief == 2 { // есть штатный, проверенный автоматизм
					return aArr[i].ID
				}
				if aArr[i].Usefulness > -5 && aArr[i].Usefulness*aArr[i].Count > usefulCount {
					usefulCount = aArr[i].Usefulness * aArr[i].Count
					autmtzmID = aArr[i].ID
					autmtzm = aArr[i]
				}
			}
		}
		if autmtzm != nil {
			// выправить Usefulness и вернуть в штатные
			if autmtzm.Usefulness < 0 {
				autmtzm.Usefulness = 0
				autmtzm.Count = 1
			}
			SetAutomatizmBelief(autmtzm, 2) // сделать автоматизм штатным
			return autmtzmID
		}
	}
	// не найден, пробуем все правила, включая учительские, что есть в ответ на Стимул curStimulImageID
	ePsrsIdArr = nil
	getIdArr(3, &EpisodicTree, curStimulImageID, 0)
	if ePsrsIdArr != nil {
		// собрать в targetEpisodicStrIdArr массив позитивных правил типа Rule из массива кадров эпиз.памяти
		ruleFromEpisodeIdArr(ePsrsIdArr, 1)
		// Выбрать одно лучшее Правило
		_, rule := findBestRule(targetEpisodicStrIdArr)
		if rule.Effect > 0 {
			// создать автоматизм, если такого нет и привязать к ветке
			oldID, am := checkUnicumMotorsAutomatizm(detectedActiveLastNodID, rule.Action)
			if oldID > 0 { // уже есть такой на ветке, хотя мы же выше проверяли, но на всякий, не создавать новый
				return oldID
			}
			if am != nil {
				// нужно создать дубликат автоматизма, чтобы старый не отвязывался от своей ветки
				id, _ := createDuplicateAutomatizm(detectedActiveLastNodID, am)
				if id > 0 {
					return id
				}
			}
		}
	}
	return 0
}

///////////////////////////////////////////////////////////////////////////

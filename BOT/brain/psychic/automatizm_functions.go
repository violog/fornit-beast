/* функции подержки автоматизмов

Концепция общих автоматизмов. Они сформированы на основе общего шаблона рефлексов и, как и рефлексы,
перекрываются автоматизмами конца активной ветки, т.е. имеющими образ Стимула.
Это – первичная реакция на текущее сочетание контекстов (в данном случае – на эмоцию).
У узла эмоции может быть бесконечное число образов действия
и блокировка общего автоматизма лишает первичной реакции у всех их.
Но у.рефлекс может быть заморожен, а общий автоматизм сейчас – нет.
Это значит, то в случае блокирующих действий для данной ветки, необходимо запускать автоматизм бездействия,
останавливающий все более низкоуровневое.lib.MapCheckBlock(MapGwardAutomatizmFromId)
При блокировке такого автоматизма опять НЕ БЛОКРУЕТСЯ.
Наличие игнор.автоматизма в конце ветки для func consciousnessElementary() равноценно отсуствю автоматизма.
Функции:
// это - игнорирующий штатный автоматизм - расценивается как отсуствие реакции.
func isIgnoreAutomatizmID(atmtzmID int)bool{
func isIgnoreAutomatizm(atmtzm *Automatizm)bool{
// автоматизм игнорирования для остановки общего автоматизма, штатно привязанный к активной ветке detectedActiveLastNodID
func getIgnoreAtmtzmToBrench()(int,*Automatizm)
*/

package psychic

import (
	word_sensor "BOT/brain/words_sensor"
	"BOT/lib"
	"sort"
)

//////////////////////////////////////////

/*
	Не раздумывая, а рефлексторно используя имеющуюся информацию,

ВЫБРАТЬ ЛУЧШИЙ АВТОМАТИЗМ для узла nodeID то более ранних, если нет у поздних.
а если нет, то учитывать общие автомтизмы, привязанные к действиям (виртуальная ветка ID от 1000000) и словам (>2000000)
*/
func getAutomatizmFromNodeID(nodeID int) int {
	if nodeID == 0 {
		return 0
	}

	//_, am := tryCreateAnswerForPhrese(2, detectedActiveLastNodID)
	//if am != nil {	}

	if isUnrecognizedPhraseFromAtmtzmTreeActivation { //при активации была нераспознанная фраза
		// это значит, что не может быть автоматизма при нераспознанной фразе и его нужно будет искать по infoFunc300
		return 0
	}

	// список всех автоматизмов для ID узла Дерева
	aArr := GetMotorsAutomatizmListFromTreeId(nodeID)
	var usefulness = -10 // полезность, выбрать наилучшую
	var autmtzm *Automatizm
	if aArr != nil {
		for i := 0; i < len(aArr); i++ {
			var allowRun = false

			if aArr[i].Usefulness < 0 {
				/* Не блокировать сразу, а посмотреть в Правила,
				м.б. после плохого эффекта последует следующее Правило с хорошим эффектом
				и тогда можно допустить Usefulness<0 в расчете на последующий успех.
				Не в качестве волевого усилия, а чисто автоматически использовать такую информацию?
				выбирая все хорошее (с эффектом, большим, чем  -aArr[i].Usefulness>0)
				Для сравнения используем положительное значение aArr[i].Usefulness
				*/

				//isWellEffect := isNextWellEffectFromActonRules(-aArr[i].Usefulness,aArr[i].ActionsImageID)
				// сравнивать эффект Правил с Usefulness автоматизма не корректно.... поэтому просто ищем для действия автоматизма.
				isWellEffect := isNextWellEffectFromActonRules(3, curActiveActionsID, aArr[i].ActionsImageID)
				if isWellEffect {
					allowRun = true // позволить выполниться автоматизму
				}
			}

			if aArr[i].Belief == 2 && allowRun { // есть штатный, проверенный автоматизм
				return aArr[i].ID
			}
			// сделаем приоритет выбора между одинаковыми usefulness за самым свежим
			// в 3 стадии это уберет косяк, когда создается новый вариант ответа на стимул, а на пульт отправляется старый
			// хотя новый автоматизм создается нормально
			// на стадиях больше 3 предпочтение отдаем штатному, который определяется при блокировке предыдущего штатного в SetAutomatizmBelief()/findBestNonStaff()
			// если не учитывать EvolushnStage > 3 то выбранный штатным в findBestNonStaff() автоматизм тут переназначится на "самый свежий", который не факт что лучший
			if aArr[i].Usefulness >= usefulness {
				if EvolushnStage > 3 && aArr[i].Usefulness >= 0 && aArr[i].Belief == 2 { // хотя Belief==2 и Usefulness >0 подразумевается по умолчанию, но подстрахуемся
					return aArr[i].ID
				} else {
					usefulness = aArr[i].Usefulness
					autmtzm = aArr[i]
				}
			}
		}
	}
	if usefulness >= 0 && autmtzm != nil { // выбран самый полезный из всех
		/*формирование не привязанных к узлу автоматизмов при активации дерева
		- для всех фраз - и для всех действий на основе привязанного автоматизма,
		чтобы другие ветки могли пользоваться при разных условиях.
		*/
		createNodeUnattachedAutomatizm(nodeID, autmtzm.ID)
		SetAutomatizmBelief(autmtzm, 2)
		return autmtzm.ID
	}
	//
	/* нет привязанных к данному узлу
	Это ситуация: фраза, для которой в узле нет автоматизма или нераспознанная фраза.
	*/
	if (curActions.PhraseID != nil && currentAutomatizmAfterTreeActivatedID == 0) || isUnrecognizedPhraseFromAtmtzmTreeActivation {
		/* Для текущей фразы сначала смотрим есть ли для данных условий (BaseID + EmotionID) автоматзмы для известных слов фразы.
		Если нет, то по Правилам выбираются подходяшие дейстивия.
		Если найдены действия, то создается автоматизм, который прикрепляется к узлу дерева автоматизмов branchID.
		В случае создания автоматизма возвращает autmtzm.ID иначе - 0.
		Если есть автоматизмы на "привет" и на "как дела?" то на "привет как дела?" должен сформироваться автоматизм из их действий.
		*/
		aID, atzm := tryCreateAnswerForPhrese(2, detectedActiveLastNodID) // 2 - максимальное число действия для автоматизма
		if aID > 0 && atzm.Usefulness >= 0 {                              // смотрим, чтобы не открыло заблокирванный автоматизм
			SetAutomatizmBelief(atzm, 2)
			return aID
		}
	}

	// на второй стадии и для функции 13 дальше не ищем, иначе начнет выкидывать общие автоматизмы
	if EvolushnStage == 2 || curFunc13ID > 0 {
		return 0
	}
	// в данном узле нет привязанного к нему автоматизма, либо автоматизм заблокирован
	// если это - узел действий или узел фразы, смотрим, если привязанные к таким объектам автоматизм

	//	node:=AutomatizmTreeFromID[nodeID] // должен быть обязательно, но...
	node, ok := ReadeAutomatizmTreeFromID(nodeID)
	if !ok {
		return 0
	}
	if node.PhraseID > 0 { // это узел фразы
		atmzS := GetAutomatizmBeliefFromPhraseId(node.PhraseID)
		if atmzS != nil {
			if autmtzm != nil {
				// если в найденном общем автоматизме такая акция, которая заблокирована в стандартном - не пускаем общий, ведь именно акцию и заблокировали
				if !isBranchUsefullnesAct(aArr, autmtzm.ActionsImageID) {
					return atmzS.ID //это - штатный автоматизм
				}
			} else {
				// если стандартных автоматизмов вообще нет, то запускаем общий
				return atmzS.ID //это - штатный автоматизм
			}
		}
	}
	/////////////
	if node.ActivityID > 0 && node.ToneMoodID == 90 { // это узел действий - конечный в активной ветке. 90 - ToneMoodID по умолчанию
		atmzA := GetAutomatizmBeliefFromActionId(node.ActivityID)
		if atmzA != nil {
			if autmtzm != nil {
				// если в найденном общем автоматизме такая акция, которая заблокирована в стандартном - не пускаем общий, ведь именно акцию и заблокировали
				if !isBranchUsefullnesAct(aArr, autmtzm.ActionsImageID) {
					return atmzA.ID //это - штатный автоматизм
				}
			} else {
				// для третьей стадии больше не ищем потому, что в случае со стимулом фраза + действие будет выдавать последний запущенный, а надо отзеркалить стимул
				if EvolushnStage == 3 {
					return 0
				}
				// если стандартных автоматизмов вообще нет, то запускаем общий
				return atmzA.ID //это - штатный автоматизм
			}
		}
	}

	//////////// нет штатных автоматизмов, выбрать любой нештатный на пробу
	/* такого быть не должно, т.к. штатный должен быть всегда
	if node.PhraseID>0 { // это узел фразы
		aArr = AutomatizmIdFromPhraseId[node.PhraseID]
		if aArr != nil {
			return aArr[0].ID // первый попавшийся не штатный, раз уже не нашелся штатный
		}
	}
	if node.ActivityID>0 && node.ToneMoodID==0 {
		aArr = AutomatizmIdFromActionId[node.PhraseID]
		if aArr != nil {
			return aArr[0].ID // первый попавшийся не штатный
		}
	}
	*/
	/////////// нет никаких автоматизмов хоть как-то относящихся к данному узлу
	// найти у предыдущих узел действий
	for i := len(ActiveBranchNodeArr) - 1; i > 2; i-- {

		//		node=AutomatizmTreeFromID[ActiveBranchNodeArr[i]]
		node, ok := ReadeAutomatizmTreeFromID(ActiveBranchNodeArr[i])
		if ok {
			atmzA := GetAutomatizmBeliefFromActionId(node.ActivityID)
			if atmzA != nil {
				if autmtzm != nil {
					// если в найденном общем автоматизме такая акция, которая заблокирована в стандартном - не пускаем общий, ведь именно акцию и заблокировали
					if !isBranchUsefullnesAct(aArr, autmtzm.ActionsImageID) {
						return atmzA.ID //это - штатный автоматизм
					}
				} else {
					// если стандартных автоматизмов вообще нет, то запускаем общий
					return atmzA.ID //это - штатный автоматизм
				}
				// не штатные автоматизмы для данного образа действий не будем смотреть
			}
		}
	}

	return 0
}

// проверка по действию и ветке дерева автоматизма - не заблокирован ли
func isBranchUsefullnesAct(aArr []*Automatizm, act int) bool {
	for _, v := range aArr {
		if v.ActionsImageID == act && v.Usefulness < 0 {
			return true
		}
	}
	return false
}

// выбрать самый успешный автоматизм ветки
func getBeastAutomatizmFromNodeID(node *AutomatizmNode) *Automatizm {
	if node == nil {
		return nil
	}
	// список всех автоматизмов для ID узла Дерева
	var aArr []*Automatizm
	for _, a := range AutomatizmFromId {
		if a == nil {
			continue
		}
		if a.BranchID == node.ID {
			aArr = append(aArr, a)
		}
	}

	var usefulness = -10 // полезность, выбрать наилучшую
	var autmtzm *Automatizm
	if aArr != nil {
		for i := 0; i < len(aArr); i++ {
			if aArr[i].Usefulness > usefulness {
				usefulness = aArr[i].Usefulness
				autmtzm = aArr[i]
			}
		}
		if usefulness >= 0 && autmtzm != nil {
			return autmtzm
		}
	}
	return nil
}

/////////////////////////////////////////////////
/*задать тип автоматизма Belief.
Только один из автоматизмов, прикрепленных к ветке или образу, может иметь Belief=2 - проверенное собственное знание
Если задается Belief=2, остальные Belief=2 становится Belief=0.
ТАК ПРОСТО НЕЛЬЗЯ ЗАДАВАТЬ Belief=2: LastAutomatizmWeiting.Belief=2

Для блокировки автоматизма нужно использовать SetAutomatizmBelief(atmzm,0) - для реализации func findBestNonStaff()
*/
func SetAutomatizmBelief(atmzm *Automatizm, belief int) {
	if atmzm == nil || atmzm.BranchID == 0 {
		return
	}
	if belief == 2 {
		atmzm.Belief = 2
		removeAllBeliefFormBranchID(atmzm) // все другие автоматизмы ветки вынести из штатных
		/* нельзя задавая Belief==2 оставлять Usefulness<0, он как минимум должен быть =0.
		 */
		if atmzm.Usefulness < 0 {
			atmzm.Usefulness = 0
			atmzm.Count = 1
		}
	}
	// при SetAutomatizmBelief(atmzm,0) - для выноса автоматизма из штатного
	if belief == 0 {
		atmzm.Belief = 0
		removeAllBeliefFormBranchID(atmzm) // все другие автоматизмы ветки вынести из штатных
		// посмотреть,  нет ли нештатные автоматизмы в ветке с Usefulness>0 т.е. заблокированных оператором под горячую руку
		aBest := findBestNonStaff(atmzm)
		if aBest != nil {
			aBest.Usefulness = 0 // разблокировать, сделав пробным штатным
			aBest.Belief = 2
		}
	}

	atmzm.Belief = belief
}

// //////////////////////////
// все автоматизмы ветки вынести из штатных
func removeAllBeliefFormBranchID(atmzm *Automatizm) {
	if atmzm == nil || atmzm.BranchID == 0 {
		return
	}
	// привязанные к ID узла дерева
	if atmzm.BranchID < 1000000 { // обнулить Belief у всех привязанных к узлу

		aArr := GetMotorsAutomatizmListFromTreeId(atmzm.BranchID)

		if len(aArr) > 1 {
			for i := 0; i < len(aArr); i++ {
				if aArr[i] != atmzm && aArr[i].Belief == 2 {
					aArr[i].Belief = 0
					AutomatizmBelief2FromTreeNodeId[aArr[i].BranchID] = nil
				}
			}
		}
		AutomatizmBelief2FromTreeNodeId[atmzm.BranchID] = atmzm
	}
	// привязанные к ID образа действий с пульта ActivityID
	if atmzm.BranchID > 1000000 && atmzm.BranchID < 2000000 { // обнулить Belief у всех привязанных к ActivityID
		imgID := atmzm.BranchID - 1000000
		for _, v := range AutomatizmIdFromActionId[imgID] {
			v.Belief = 0
		}
	}
	if atmzm.BranchID > 2000000 { // обнулить Belief у всех привязанных к PhraseID
		imgID := atmzm.BranchID - 2000000
		for _, v := range AutomatizmIdFromPhraseId[imgID] {
			v.Belief = 0
		}
	}
}

/*
найти нештатные автоматизмы в ветке с Usefulness>0 т.е. заблокированные оператором под горячую руку
кроме данного: atmzm *Automatizm
Только для нормальных (не групповых) автоматизмов.

Восстановление блокированных автомаптизмов ветки.
*/
func findBestNonStaff(atmzm *Automatizm) *Automatizm {
	if atmzm == nil || atmzm.BranchID == 0 {
		return nil
	}
	beat := 0
	var aBest *Automatizm

	aArr := GetMotorsAutomatizmListFromTreeId(atmzm.BranchID)

	if len(aArr) > 1 {
		for i := 0; i < len(aArr); i++ {
			if aArr[i].ID == atmzm.ID {
				continue
			}
			if aArr[i].Usefulness > 0 && aArr[i].Usefulness > beat {
				beat = aArr[i].Usefulness
				aBest = aArr[i]
			}
		}
	}
	if aBest != nil {
		return aBest
	}
	return nil
}

/////////////////////////////////////////////////////

// список всех автоматизмов для ID узла Дерева
var lastAutomatizmArrFromNodeID []*Automatizm // уже полученный массив для lastAutomatizmsNodeID чтобы не повторяться
var lastAutomatizmsNodeID int

func GetMotorsAutomatizmListFromTreeId(nodeID int) []*Automatizm {
	if nodeID == 0 {
		return nil
	}

	if AutomatizmFromId == nil {
		return nil
	}
	if lastAutomatizmsNodeID == nodeID { // уже есть массив для nodeID
		return lastAutomatizmArrFromNodeID
	}
	lastAutomatizmsNodeID = nodeID

	lastAutomatizmArrFromNodeID = nil

	for _, a := range AutomatizmFromId {
		if a == nil {
			continue
		}
		if a.BranchID < 1000000 && a.BranchID == nodeID {
			lastAutomatizmArrFromNodeID = append(lastAutomatizmArrFromNodeID, a)
		}
	}

	// нужно для работы в 3 стадии
	sort.SliceStable(lastAutomatizmArrFromNodeID, func(i, j int) bool {
		return lastAutomatizmArrFromNodeID[i].ID < lastAutomatizmArrFromNodeID[j].ID
	})
	return lastAutomatizmArrFromNodeID
}

// штатный, невредный автоматизм, привязанный к ветке
func GetBelief2AutomatizmListFromTreeId(nodeID int) *Automatizm {
	if nodeID == 0 {
		return nil
	}
	aArr := AutomatizmBelief2FromTreeNodeId[nodeID]

	if aArr == nil {
		return nil
	}
	if aArr.Usefulness >= 0 { // есть штатный, невредный
		return aArr
	}
	return nil
}

//////////////////////////////////////////////////

// есть ли штатный автоматизм (с Belief==2), привязанные к узлу дерева
func ExistsAutomatizmForThisNodeID(nodeID int) bool {
	aArr := AutomatizmBelief2FromTreeNodeId[nodeID]
	if aArr != nil {
		return true
	}
	return false
}

///////////////////////////////////////

/*
	если для прикрепленных к узлу дерева есть карта штатных AutomatizmBelief2FromTreeNodeId,

то для прикрепленных к образам нужны ФУНКЦИИ ПОЛУЧЕНИЯ ШТАТНОГО ДЛЯ ДАННОГО ОБРАЗА
*/
func GetAutomatizmBeliefFromActionId(activityID int) *Automatizm {
	if AutomatizmIdFromActionId[activityID] == nil {
		return nil
	}

	for _, v := range AutomatizmIdFromActionId[activityID] {
		if v == nil {
			continue
		}
		//ai,ok:=ActionsImageArr[v.ActionsImageID]
		ai, ok := ReadeActionsImageArr(v.ActionsImageID)
		if !ok {
			continue
		}
		if ai.ActID != nil && v.Belief == 2 {
			return v
		}
	}
	return nil
}

// /////////////////////////////////////////////////
func GetAutomatizmBeliefFromPhraseId(verbalID int) *Automatizm {
	if AutomatizmIdFromPhraseId[verbalID] == nil {
		return nil
	}
	for _, v := range AutomatizmIdFromPhraseId[verbalID] {
		if v.Belief == 2 {
			return v
		}
	}
	return nil
}

/*
формирование не привязанных к узлу автоматизмов при активации дерева
- для всех фраз - и для всех действий на основе привязанного автоматизма,
чтобы другие ветки могли пользоваться при разных условиях.
*/
func createNodeUnattachedAutomatizm(nodeID int, aID int) {

	//	node:=AutomatizmTreeFromID[nodeID] // должен быть обязательно, но...
	node, ok := ReadeAutomatizmTreeFromID(nodeID)
	if !ok {
		return
	}
	//autmzm0:= AutomatizmFromId[aID] // должен быть обязательно, но...
	autmzm0, ok := ReadeAutomatizmFromId(aID)
	if !ok {
		return
	}

	if node.PhraseID > 0 { // это узел фразы
		_, autmzm := CreateAtutomatizmNoSaveFile(2000000+node.PhraseID, autmzm0.ActionsImageID)
		if autmzm != nil && autmzm.Usefulness >= 0 { // не даем открывать заблокированные
			autmzm.Usefulness = 0          // пока предположительно
			SetAutomatizmBelief(autmzm, 2) // сделать автоматизм штатным
		}
	}
	/*	получается некий аналог у-рефлекса, когда на разные стимулы выдается один и тот же ответ
		в данном случае на полный стимул "вербальный + действие" и частичные "вербальный", "действие" создаются обшие автоматизмы с одним ответом autmzm0.ActionsImageID
	*/if node.ActivityID > 0 && node.ToneMoodID == 90 { // это узел действий - конечный в активной ветке. 90 - ToneMoodID по умолчанию
		_, autmzm := CreateAtutomatizmNoSaveFile(1000000+node.ActivityID, autmzm0.ActionsImageID)
		if autmzm != nil && autmzm.Usefulness >= 0 { // не даем открывать заблокированные
			autmzm.Usefulness = 0          // пока предположительно
			SetAutomatizmBelief(autmzm, 2) // сделать автоматизм штатным
		}
		// можно выделить из ответа только действие и создать общий автоматизм с таким ответом, но тогда получится просто клон рефлекса
		// что и так происходит в стадиях начиная со 2. Поэтому пока под сомнением - надо ли так делать, может правильнее вариант выше
		/*		if actImg, ok:=ActionsImageArr[autmzm0.ActionsImageID]; ok{
				actimgID,_:=CreateNewlastActionsImageID(0, 0, actImg.ActID, nil, 0, 0, true)
				_,autmzm:= CreateAtutomatizmNoSaveFile(1000000+node.ActivityID, actimgID)
				if autmzm!=nil{
					autmzm.Usefulness=0 // пока предположительно
					SetAutomatizmBelief(autmzm, 2) // сделать автоматизм штатным
				}
			}*/
	}
}

// разблоикровака автоматизма для http://go/pages/automatizm_table.php
func UnblockAutomatizmID(atmtzmID int) string {

	//	atmtzm:= AutomatizmFromId[atmtzmID]
	atmtzm, ok := ReadeAutomatizmFromId(atmtzmID)
	if !ok {
		return "0"
	}
	atmtzm.Usefulness = 1
	atmtzm.Count = 1
	return "1"
}

func UnblockingAllAtmtzms() {
	for _, v := range AutomatizmFromId {
		if v == nil {
			continue
		}
		if v.Usefulness < 0 {
			v.Usefulness = 0
			v.Count = 1
		}
	}

}

/////////////////////////////////////////////////////////////////////

// привязать общий автоматизм к активной ветке detectedActiveLastNodID
func linkCoomonAtmtzmToBrench(commonAutomatizm *Automatizm) {
	if LastAutomatizmWeiting.BranchID < 1000000 { // это НЕ общий - не должно такого быть
		return
	}
	atmtzm := GetBelief2AutomatizmListFromTreeId(detectedActiveLastNodID)
	if atmtzm != nil && atmtzm.ID == commonAutomatizm.ID {
		return
	}
	CreateAtutomatizmNoSaveFile(detectedActiveLastNodID, commonAutomatizm.ActionsImageID)
}

/////////////////////////////////////////////////////////////////////

// автоматизм игнорирования для остановки общиего автоматизма, штатно привязанный к активной ветке
func getIgnoreAtmtzmToBrench(BranchID int) (int, *Automatizm) {
	// игнорирующее действие:
	aID, _ := CreateNewlastActionsImageID(0, 0, []int{9}, nil, 0, 0, true)
	if aID <= 0 {
		return 0, nil
	}
	id, atmtzm := createNewAutomatizmID(0, BranchID, aID, true)

	if atmtzm != nil {
		return 0, nil
	}
	SetAutomatizmBelief(atmtzm, 2)

	return id, atmtzm
}

// определение, что это - игнорирующий ID автоматизма
func isIgnoreAutomatizmID(atmtzmID int) bool {
	//	atmtzm,ok:=AutomatizmFromId[atmtzmID]
	atmtzm, ok := ReadeAutomatizmFromId(atmtzmID)
	if !ok {
		return false
	}

	return isIgnoreAutomatizm(atmtzm)
}

// пределение, что это - игнорирующий автоматизм
func isIgnoreAutomatizm(atmtzm *Automatizm) bool {
	if atmtzm == nil {
		return false
	}
	iaID, _ := CreateNewlastActionsImageID(0, 0, []int{9}, nil, 0, 0, true)
	if atmtzm.ActionsImageID == iaID {
		return true
	}

	return false
}

/////////////////////////////////////////////////////////////////////

/*
есть ли автоматизм с действием curStimulImageID, и если у него atmtzm.Usefulness<0 - ПОСТЕПЕННО снять блокировку
потому как это - новое авторитарное подтвержение полезности.
Для текущей ветки дерева автоматизмов.
*/
func checkForUnbolokingAutomatizm(actID int) {
	for _, v := range AutomatizmFromId {
		if v == nil {
			continue
		}
		// накручиваем только Usefulness==-1, Usefulness==0 - это попугайские автоматизмы, их не трогаем!
		// при отзеркаливании на 2 и 4 стадиях нельзя делать попугайские с Usefulnes==-1 - иначе они тут же разблокируются и станут штатными
		if v.BranchID == detectedActiveLastNodID && v.ActionsImageID == actID && v.Usefulness < 0 {
			v.Usefulness = v.Usefulness + 1 // ++ не срабатывает
			if v.Usefulness == 0 {
				//v.Usefulness=0  пока предположительно
				SetAutomatizmBelief(v, 2) // сделать автоматизм штатным, полезность 1 установится там же автоматически
			}
		}
	}

}

///////////////////////////////////////////////////////

/*
	найти узел ветки данного уроня, начиная с конечного узла branchID

Первый уровень - BaseID, 6-й уровень - PhraseID
*/
func getNodeFromLevel(level int, branchID int) *AutomatizmNode {

	//ln:=AutomatizmTreeFromID[branchID]
	ln, ok := ReadeAutomatizmTreeFromID(branchID)
	if !ok {
		return nil
	}
	// сначала получить спиоск всех узлов ветки
	var nArr []*AutomatizmNode
	for ln != nil {
		nArr = append(nArr, ln)
		ln = ln.ParentNode
	}

	// берем второй уровень с эмоцией, удалив последние 2
	eNode := nArr[len(nArr)-3]
	return eNode
}

/////////////////////////////////////////////////////////////////

/*
	начать поиск с данного узла по словам поиском по дереву - лучший автоматизм из найденных

пройти 2 урояня: ActivityID и ToneMoodID до SimbolID
*/
func getAtmtzmFromNodesFrase(node *AutomatizmNode, phraseID int) *Automatizm {
	//FirstSimbolID:=word_sensor.GetFirstSymbolFromWordID(wordIDarr[0])
	FirstSimbolID := word_sensor.GetFirstSymbolFromPraseID([]int{phraseID})

	var atmzmArr []*Automatizm

	for k := 0; k < len(node.Children); k++ { // тут ActivityID
		nextlev1 := &node.Children[k]
		for l := 0; l < len(nextlev1.Children); l++ { // тут ToneMoodID
			nextlev2 := &nextlev1.Children[l]
			for m := 0; m < len(nextlev2.Children); m++ { // тут SimbolID
				if FirstSimbolID != nextlev2.Children[m].SimbolID {
					continue
				}
				nextlev3 := &nextlev2.Children[m]

				//if nextlev3.PhraseID == phraseID {
				if existsPraseIDinVerbID(nextlev3.PhraseID, phraseID) {
					a := getBeastAutomatizmFromNodeID(nextlev3)
					if a != nil {
						atmzmArr = append(atmzmArr, a)
					}
				}
				for n := 0; n < len(nextlev3.Children); n++ { // тут PhraseID
					nextlev4 := &nextlev3.Children[n]
					/*
						AutomatizmTreeMapCheck()
						verbal,ok:=VerbalFromIdArr[nextlev.PhraseID]
						if ok {
							wArr:=word_sensor.GetWordsIDarrFromPraseNodeID(verbal.PhraseID[0])
							if lib.EqualArrs(wArr, phraseArr){
								a:=getBeastAutomatizmFromNodeID(nextlev)
								atmzmArr=append(atmzmArr,a)
							}
						}*/
					//if nextlev4.PhraseID == phraseID {
					if existsPraseIDinVerbID(nextlev4.PhraseID, phraseID) { //// есть ли в данном Verb фраза praseID int
						a := getBeastAutomatizmFromNodeID(nextlev4)
						if a != nil {
							atmzmArr = append(atmzmArr, a)
						}
					}
				}
			}
		}
	}
	if atmzmArr != nil { // выбор лучшего
		max := -1
		var curA *Automatizm
		for n := 0; n < len(atmzmArr); n++ {
			if atmzmArr[n].Usefulness > max {
				max = atmzmArr[n].Usefulness
				curA = atmzmArr[n]
			}
		}
		if curA != nil {
			return curA
		}
	}

	return nil
}

///////////////////////////////////////////////////////////////////

/*
	Есть ли в узде дерева автоматизмов автоматизм с образом действий ActionsImage?

Когда для данного узла дерева автоматизмов находится новый образ действий,
нужно проверять, а не заблокирован ли такой ранее.
Если да, то незачем создавать такой же снова.

Если находит, то возвращает ID автоматизма.
*/
func existsThistActionIdInAtreeNode(nodeID int, ActionsImage int) (bool, int) {
	// это может быть вызов при создании нового автоматизма по учительскому правилу, поэтому nodeID - не обязательный параметр
	if nodeID == 0 {
		return false, 0
	}
	if ActionsImage == 0 {
		lib.TodoPanic("В func existsThistActionIdInAtreeNode неверный параметр!")
		return false, 0
	}
	// список всех автоматизмов для ID узла Дерева
	aArr := GetMotorsAutomatizmListFromTreeId(nodeID)
	if aArr == nil {
		return false, 0
	}
	for i := 0; i < len(aArr); i++ {
		if aArr[i].ActionsImageID == ActionsImage {
			return true, aArr[i].ActionsImageID
		}
	}

	return false, 0
}

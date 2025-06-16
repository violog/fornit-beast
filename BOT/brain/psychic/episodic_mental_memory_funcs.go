/* функции для работы в сентальными эпизодами памяти

В отличие от моторной эпиз.памяти здесь нет цепочек правил, разделенных пустыми кадрами.
Здесь в каждом кадре есть цепочка инфо-функций, которые применялись в данных условиях.
Тут принцип GPT - в том, что после запуска инфо-фукнции прогнозируется запуск следующей с определенной уверенностью.

В принципе здесь и не нужна оказывается историческая последовательность кадров, но она хранит время.
*/

package psychic

import "strconv"

///////////////////////////////////////////////////
/* буфер результатов поиска ID кадров
 */
var ePMentSrsIdArr []int

// Не выдавать по запросу id вспомогательных функций.
func isNoWorkInfoID(id int) bool {
	if id == 1 || id == 2 || id == 5 || id == 8 || id == 14 || id == 17 || id == 26 {
		return true
	}
	return false
}

///////////////////////////////////////////////////

/*
	Найти все ID кадров для которых выполняется условие с заданным эффектом.

NodeAID int, ThemeID int,PurposeID int могут быть ==0 НО только если есть ненулевой предыдущий,
т.е. нельзя getEpisodesMentalArrFromConditions(0, 12,0)
Тогда поиск будет максимально быстрым.
Все сразу могут быть нулевыми и тогда выдаст все кадры дерева.

Для поиска по любому из 3-х условий func getEpisodesMentalArrFromAnyConditions - более долгий перебор всех узлов дерева.

typeEffect: 0- любой, 1-позитивный, 2-негативный
*/
func getEpisodesMentalArrFromConditions(NodePID int, ThemeID int, PurposeID int, typeEffect int) []int {

	ePMentSrsIdArr = nil

	var cond []int
	if NodePID > 0 {
		cond = append(cond, NodePID)
	}
	if ThemeID > 0 {
		cond = append(cond, ThemeID)
	}
	if PurposeID > 0 {
		cond = append(cond, PurposeID)
	}

	id, _ := findEpisodicMentalBrange(0, cond, &EpisodicMentalTree)
	if id == 0 {
		return nil
	}

	//		node, ok := EpisodicTreeNodeFromID[id]
	node, ok := ReadeEpisodicMentalTreeNodeFromID(id)
	if !ok {
		return nil
	}
	getMentalIdArr(node, typeEffect)
	return ePMentSrsIdArr

	return nil
}

// вытащить все ID узлов, в которых есть условие, начиная с node
func getMentalIdArr(node *EpisodicMentalTreeNode, typeEffect int) {
	for _, child := range node.Children {

		if child.PARAMS != nil { // узел с прописанным PARAMS и там есть InfoArr
			if typeEffect == 1 && child.PARAMS[1] < 0 {
				continue
			}
			if typeEffect == 2 && child.PARAMS[1] >= 0 {
				continue
			}
			if isNoWorkInfoID(child.ID) {
				continue
			}
			ePMentSrsIdArr = append(ePMentSrsIdArr, child.ID)
		}
		// продолжим рекурсивно искать в нем далее
		getMentalIdArr(&child, typeEffect)
	}

	return
}

////////////////////////////////////////////

/*
Для поиска по любому из 3-х условий func getEpisodesMentalArrFromAnyConditions - более долгий перебор всех узлов дерева.
Любой из NodeAID int, ThemeID int,PurposeID int м.б. нулем и даже все сразу
typeEffect: 0- любой, 1-позитивный, 2-негативный
*/
func getEpisodesMentalArrFromAnyConditions(NodePID int, ThemeID int, PurposeID int, typeEffect int) []int {

	ePMentSrsIdArr = nil
	getMentalIdConditionArr(&EpisodicMentalTree, NodePID, ThemeID, PurposeID, typeEffect)
	return ePMentSrsIdArr
}

// вытащить все ID узлов, в которых есть условие, начиная с node
func getMentalIdConditionArr(node *EpisodicMentalTreeNode, NodePID int, ThemeID int, PurposeID int, typeEffect int) {
	for _, child := range node.Children {

		if child.PARAMS != nil { // узел с прописанным PARAMS и там есть InfoArr
			if NodePID > 0 && child.NodePID != NodePID {
				continue
			}
			if ThemeID > 0 && child.ThemeID != ThemeID {
				continue
			}
			if PurposeID > 0 && child.PurposeID != PurposeID {
				continue
			}
			if typeEffect == 1 && child.PARAMS[1] < 0 {
				continue
			}
			if typeEffect == 2 && child.PARAMS[1] >= 0 {
				continue
			}
			if isNoWorkInfoID(child.ID) {
				continue
			}
			ePMentSrsIdArr = append(ePMentSrsIdArr, child.ID)
		}
		// продолжим рекурсивно искать в нем далее
		getMentalIdConditionArr(&child, NodePID, ThemeID, PurposeID, typeEffect)
	}

	return
}

////////////////////////////////////////////////////////////

/*
найти ID успешной инфо-функции со все большим отклонением от текущих условий
*/
func findSuitableMentalFunc() []int {
	NodePID := detectedActiveLastProblemNodID // ID проблемы
	ThemeID := problemTreeInfo.themeID
	PurposeID := mentalInfoStruct.mentalPurposeID

	getEpisodesMentalArrFromConditions(NodePID, ThemeID, PurposeID, 1)
	if ePMentSrsIdArr == nil {
		getEpisodesMentalArrFromConditions(NodePID, ThemeID, 0, 1)
	}
	if ePMentSrsIdArr == nil {
		getEpisodesMentalArrFromConditions(NodePID, 0, 0, 1)
	}
	if ePMentSrsIdArr == nil { // в поиске по всем деревьям выбираем Цель
		getEpisodesMentalArrFromAnyConditions(0, 0, PurposeID, 1)
	}
	// выдавать все кадры нет смысла, так что на этом заканчиваем

	return ePMentSrsIdArr
}

////////////////////////////////////////////////////

/*
	Выбрать следующую финфо-функцию для запуска из цепочки наиболее подходящих для данных условий

Если еще не начато заполнение буфера исполненных функций mentAtmzmActualFuncs, то выдать первую из цепочки,
если уже выполнялись функции, дать следующую подходящую после mentAtmzmActualFuncs
GRP метод.
*/
func getNextGPTstepMental() int {
	NodePID := detectedActiveLastProblemNodID // ID проблемы
	ThemeID := problemTreeInfo.themeID
	PurposeID := mentalInfoStruct.mentalPurposeID

	getEpisodesMentalArrFromConditions(NodePID, ThemeID, PurposeID, 1)
	if ePMentSrsIdArr == nil {
		return 0 // все на этом т.к. этот метод только для точного соотвествия условиям
	}
	// найти продолжение для mentAtmzmActualFuncs
	if mentAtmzmActualFuncs == nil { // еще не было запущено функций
		return ePMentSrsIdArr[0]
	}
	lelFound := len(ePMentSrsIdArr)
	lenExists := len(mentAtmzmActualFuncs)
	if lelFound < lenExists {
		return 0
	}
	isNotEquivalentFrame := false
	for i := 0; i < lelFound && i < lenExists; i++ {
		if isNoWorkInfoID(ePMentSrsIdArr[i]) {
			continue
		}
		if ePMentSrsIdArr[i] != mentAtmzmActualFuncs[i] {
			isNotEquivalentFrame = true
		}
	}
	if isNotEquivalentFrame { // нет совпадений цепочек
		return 0
	}
	lastArr := ePMentSrsIdArr[lenExists:]
	return lastArr[0]
}

///////////////////////////////////////////////////////////////////

/*
	выбрать инфо-функцию для запуска в данных услових.

exactly - только точное совпадение условий
*/
func findInfoIdFromExperience(exactly bool) int {
	// по GPT
	fID := getNextGPTstepMental()
	if fID > 0 {
		return fID
	}
	if exactly {
		return 0
	}
	// тогда - наиболее привычную
	fArr := getFavoritInfoFunc()
	if fArr != nil {
		// не выдавать 8-ю
		for i := 0; i < len(fArr); i++ {
			if !isNoWorkInfoID(fArr[i]) {
				return fArr[i]
			}

		}
		return 0
	}

	return 0
}

/////////////////////////////////////////////////////////////

/*
МЕНТАЛЬНЫЙ АВТОМАТИЗМ - выдать номера инфо-фукнций, привычных (Count>1) для данных условий
*/
func getFavoritInfoFunc() []int {
	findSuitableMentalFunc()
	if ePMentSrsIdArr != nil {
		for i := 0; i < len(ePMentSrsIdArr); i++ {
			frame, ok := ReadeEpisodicMentalTreeNodeFromID(ePMentSrsIdArr[i])
			if ok && frame.PARAMS[1] > 1 { // Count>1  МОЖЕТ ОПЬТМИЗИРОВАТЬСЯ от опыта
				return frame.InfoArr
			}
		}
	}

	return nil
}

////////////////////////////////////////////////////////

// для пульта выдать номера инфо-фукнций, привычных для данных условий
func GetMentalAutomatizmForPult() string {
	mentAtmzmActualFuncs := getFavoritInfoFunc()
	if mentAtmzmActualFuncs != nil && len(mentAtmzmActualFuncs) > 0 {
		maStr := ""
		for i := 0; i < len(mentAtmzmActualFuncs); i++ {
			if i > 0 {
				maStr += ", "
			}
			maStr += strconv.Itoa(mentAtmzmActualFuncs[i])
		}
		return maStr
	}
	return "Нет привычной цепочки инфо-функций для данных условий."
}

////////////////////////////////////////////////////

/* Выборки из эпизодической памяти Моделей понимания
ВЫБОРКИ КАДРОВ С ПРАВИЛАМИ, в которых Trigger есть ПОЛНЫЙ образ восприятия

Типичная апись - формирует элемнты Правила модели для объекта:
saveNewEpisodic(maxE.extremObjID,maxE.kind,maxE.extremVal,1) - элемент модели понимания в условиях его выявления


Обновление массива var UnderstandingModel []ObjModel
*/

package psychic

import (
	word_sensor "BOT/brain/words_sensor"
	"BOT/lib"
)

//////////////////////////////////////////////////////////////////////

/*
найти лучший образ действий для стимула extremImportanceObject ТОЧНО
Выборка эпизодов Модели понимания, содержащими Stimul типа ActionsImage в качестве Стимула

ТОЧНО ДЛЯ ТЕКУЩИХ УСЛОВИЙ NodeAID,NodeSID,NodePID
Найти все эпизоды с Stimul типа ActionsImage в качестве Стимула,
т.е. с Trigger == Stimul и node.PARAMS[0][2]>0
Результат - в виде массива Правил UnderstandingRulesModel []RulesModel

Быстрая.
*/
func getRulesModelExactly(objID int) {
	if objID == 0 {
		lib.TodoPanic("НУЛЕВОЙ ПАРАМЕТР Stimul в func getRulesModelExactly!")
		return
	}
	UnderstandingRulesModel = nil

	cond := []int{detectedActiveLastNodID,
		detectedActiveLastUnderstandingNodID,
		detectedActiveLastProblemNodID,
		objID}
	id, _ := findEpisodicBrange(0, cond, &EpisodicTree)
	if id == 0 {
		return
	}
	node, ok := ReadeEpisodicTreeNodeFromID(id)
	if !ok {
		return
	}

	scanEpizRulesChildrens(node)
}

// вытащить все ID узлов, в которых есть условие, начиная с node
func scanEpizRulesChildrens(node *EpisodicTreeNode) {
	for _, child := range node.Children {

		if len(child.PARAMS) > 0 { // узел с прописанным PARAMS
			om := RulesModel{child.ID, child.Trigger, child.Action, child.PARAMS[0], child.PARAMS[1]}
			UnderstandingRulesModel = append(UnderstandingRulesModel, om)
		}
		// продолжим рекурсивно искать в нем далее
		scanEpizRulesChildrens(&child)
	}
	return
}

/////////////////////////////////////////////////////////////////////////////////

/*
найти лучший образ действий для стимула extremImportanceObject ПРИБЛИЗИТЕЛЬНО (только для условия detectedActiveLastNodID)
Выборка эпизодов Модели понимания с Правилами, содержащими Stimul типа ActionsImage

ПРИБЛИЗИТЕЛЬНО ДЛЯ ТЕКУЩИХ УСЛОВИЙ NodeAID
Найти все эпизоды с Stimul типа ActionsImage в качестве Стимула,
т.е. с Trigger == Stimul и node.PARAMS[0][2]>0
Результат - в виде массива Правил UnderstandingRulesModel []RulesModel

Быстрая.
*/
func getRulesModelApproximately(objID int) {
	if objID == 0 {
		lib.TodoPanic("НУЛЕВОЙ ПАРАМЕТР Stimul в func getRulesModelApproximately!")
		return
	}
	UnderstandingRulesModel = nil

	cond := []int{detectedActiveLastNodID}
	id, _ := findEpisodicBrange(0, cond, &EpisodicTree)
	if id == 0 {
		return
	}
	node, ok := ReadeEpisodicTreeNodeFromID(id)
	if !ok {
		return
	}

	scanEpizRulesChildrens2(node, objID)
}

// вытащить все ID узлов, в которых есть условие, начиная с node
func scanEpizRulesChildrens2(node *EpisodicTreeNode, stimul int) {
	for _, child := range node.Children {

		if len(child.PARAMS) > 0 { // узел с прописанным PARAMS
			if child.Trigger != stimul {
				continue
			}
			om := RulesModel{child.ID, child.Trigger, child.Action, child.PARAMS[0], child.PARAMS[1]}
			UnderstandingRulesModel = append(UnderstandingRulesModel, om)
		}
		// продолжим рекурсивно искать в нем далее
		scanEpizRulesChildrens(&child)
	}
	return
}

/////////////////////////////////////////////////////////////////////////////////

/*
	Выборка эпизодов Модели понимания , содержащими объект типа actImgID для ответного действия

ТОЧНО ДЛЯ ТЕКУЩИХ УСЛОВИЙ NodeAID,NodeSID,NodePID
Найти все эпизоды с объект типа actImgID,
т.е. с Trigger == actImgID и node.PARAMS[0][2]>0
Результат - в виде массива Правил UnderstandingRulesModel []RulesModel

Быстрая.
*/
func getRulesModelFromActionExactly(actImgID int) {
	if actImgID == 0 {
		lib.TodoPanic("НУЛЕВОЙ ПАРАМЕТР actImgID в func getRulesModelFromActionExactly!")
		return
	}
	UnderstandingRulesModel = nil

	cond := []int{detectedActiveLastNodID,
		detectedActiveLastUnderstandingNodID,
		detectedActiveLastProblemNodID}
	id, _ := findEpisodicBrange(0, cond, &EpisodicTree)
	if id == 0 {
		return
	}
	node, ok := ReadeEpisodicTreeNodeFromID(id)
	if !ok {
		return
	}

	scanEpizRulesFromActionChildrens(node, actImgID)
}

// вытащить все ID узлов, в которых есть условие, начиная с node
func scanEpizRulesFromActionChildrens(node *EpisodicTreeNode, actImgID int) {
	for _, child := range node.Children {

		if node.PARAMS != nil && len(node.PARAMS) > 0 { // узел с прописанным PARAMS
			if node.Action != actImgID {
				continue
			}
			om := RulesModel{child.ID, child.Trigger, child.Action, child.PARAMS[0], child.PARAMS[1]}
			UnderstandingRulesModel = append(UnderstandingRulesModel, om)
		}
		// продолжим рекурсивно искать в нем далее
		scanEpizRulesFromActionChildrens(&child, actImgID)
	}
	return
}

/////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////////////

/*
	Выборка эпизодов Модели понимания, содержащими СЛОВО в Стимуле (Trigger)

ТОЧНО ДЛЯ ТЕКУЩИХ УСЛОВИЙ
Найти все эпизоды с Stimul типа ActionsImage в качестве Стимула,
т.е. с Trigger содержащие wordID
Результат - в виде массива Правил UnderstandingRulesModel []RulesModel
*/
type WordModel struct {
	maxEffect    int // наилучший эффект слова
	minEffect    int // наихудший эффект слова
	beastFrameID int // ID наилучшего по эффекту кадра
	worstFrameID int // ID наихудчшего по эффекту кадра
	beastCount   int // уверенность наилучшего по эффекту кадра
	worstCount   int // уверенность наихудчшего по эффекту кадра
}

var maxEffect = 0
var minEffect = 0
var beastFrameID = 0
var worstFrameID = 0
var beastCount = 0 // уверенность наилучшего по эффекту кадра
var worstCount = 0 // уверенность наихудчшего по эффекту кадра

func getWordModelExactly(wordID int) WordModel {
	out := WordModel{0, 0, 0, 0, 0, 0}
	if wordID == 0 {
		lib.TodoPanic("НУЛЕВОЙ ПАРАМЕТР wordID в func getWordModelExactly!")
		return out
	}
	maxEffect = 0
	minEffect = 0
	beastFrameID = 0
	worstFrameID = 0
	beastCount = 0
	worstCount = 0

	cond := []int{detectedActiveLastNodID,
		detectedActiveLastUnderstandingNodID,
		detectedActiveLastProblemNodID}
	id, _ := findEpisodicBrange(0, cond, &EpisodicTree)
	if id == 0 {
		return out
	}
	node, ok := ReadeEpisodicTreeNodeFromID(id)
	if !ok {
		return out
	}

	scanEpizRulesWords(node, wordID)
	out.maxEffect = maxEffect
	out.minEffect = minEffect
	out.beastFrameID = beastFrameID
	out.worstFrameID = worstFrameID
	out.beastCount = beastCount
	out.worstCount = worstCount
	return out
}

// вытащить все ID узлов, в которых есть условие, начиная с node
func scanEpizRulesWords(node *EpisodicTreeNode, wordID int) {
	for _, child := range node.Children {

		if len(child.PARAMS) > 0 { // узел с прописанным PARAMS

			var existsWord = false
			trigger, ok := ReadeActionsImageArr(child.Trigger)
			if ok {
				if existsWordInPraseArr(wordID, trigger.PhraseID) {
					existsWord = true
				}
			}
			if trigger.PhraseID != nil { //Verbal - при активации дерева автоматизмов
				varb, ok := ReadeVerbalFromIdArr(trigger.PhraseID[0])
				if ok {
					if existsWordInPraseArr(wordID, varb.PhraseID) {
						existsWord = true
					}
				}
			}

			if !existsWord {
				continue
			}
			eff := getWpower(child.PARAMS[0], child.PARAMS[1])
			if maxEffect < eff {
				beastFrameID = child.ID
				beastCount = child.PARAMS[1]
				maxEffect = eff
			}
			if minEffect > eff {
				worstFrameID = child.ID
				worstCount = child.PARAMS[1]
				minEffect = eff
			}
		}
		// продолжим рекурсивно искать в нем далее
		scanEpizRulesWords(&child, wordID)
	}
	return
}
func existsWordInPraseArr(wordID int, phraseArr []int) bool {
	for i := 0; i < len(phraseArr); i++ {
		if phraseArr[i] == -1 { // фраза не распознанна
			continue
		}
		if existsWordInPrase(wordID, phraseArr[i]) {
			return true
		}
	}
	return false
}
func existsWordInPrase(wordID int, phrase int) bool {
	wArr := word_sensor.GetWordsArrFromPraseID(phrase)
	if wArr == nil {
		return false
	}
	for i := 0; i < len(wArr); i++ {
		if wordID == wArr[i] {
			return true
		}
	}
	return false
}

// ///////////////////////////////
// найти наилучший и наихудший случай из всех слов wArr []int
type WordArrModel struct {
	maxEffect    int // наилучший эффект слова
	minEffect    int // наихудший эффект слова
	beastFrameID int // ID наилучшего по эффекту кадра
	worstFrameID int // ID наихудчшего по эффекту кадра
	beastWordID  int // ID наилучшего по эффекту кадра
	worstWordID  int // ID наихудчшего по эффекту кадра
	beastCountID int // уверенность наилучшего по эффекту кадра
	worstCountID int // уверенность наихудчшего по эффекту кадра
}

var amaxEffect = 0
var aminEffect = 0
var abeastFrameID = 0
var aworstFrameID = 0
var abeastCount = 0 // уверенность наилучшего по эффекту кадра
var aworstCount = 0 // уверенность наихудчшего по эффекту кадра
var abeastWordID = 0
var aworstWordID = 0

func getWordArrModelExactly(wArr []int) WordArrModel {
	amaxEffect = 0
	aminEffect = 0
	abeastFrameID = 0
	aworstFrameID = 0
	abeastWordID = 0
	aworstWordID = 0
	beastCount = 0
	aworstCount = 0
	for i := 0; i < len(wArr); i++ {
		wm := getWordModelExactly(wArr[i])
		if wm.beastFrameID > 0 || wm.worstFrameID > 0 { // есть значение
			if amaxEffect < wm.maxEffect {
				amaxEffect = wm.maxEffect
				beastCount = wm.beastCount
				abeastFrameID = wm.beastFrameID
				abeastWordID = wArr[i]
			}
			if aminEffect > wm.minEffect {
				aminEffect = wm.minEffect
				aworstCount = wm.worstCount
				aworstFrameID = wm.worstFrameID
				aworstWordID = wArr[i]
			}
		}
	}
	wmArr := WordArrModel{
		amaxEffect,
		aminEffect,
		abeastFrameID,
		aworstFrameID,
		abeastWordID,
		aworstWordID,
		beastCount,
		aworstCount}
	return wmArr
}

/////////////////////////////////////////////////////////////////////////////////

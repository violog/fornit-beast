/* распознавание фразы
вербальная иерархия распознавателей

Распознавание фраз начинается в main.go с word_sensor.VerbalDetection(text_dlg, is_input_rejim, moodID)
Память о воспринятых фразах в текущем активном контексте (Vernike_detector.go): var MemoryDetectedArr []MemoryDetected
*/

package word_sensor

import (
	_ "strconv"
	_ "strings"
)

func iniPraseRecognising() {

}

// текущий уникальный ID последней активной ветки дерева - результат детекции фразы - для дальнейшего использования
var DetectedUnicumPhraseID = 0

// нераспознанный остаток
var CurrentVerbalPhraseEnd []int

// текущий найденный ID последней активной ветки дерева слов
var DetectedCurrentPhraseID = 0

// текущий номер слова распознаваемой фразы
var currentStepPhraseCount = 0

/* переносим в дерево слов достаточно повторяющиеся из tempArr для trees_former.go
Распознать и вставить новое слово-фразу в дерево:
найти подходящий узел и если еще нет - вставить новый.
*/

/*
	проход одной фразы - распознавание ID слов фразы

noCreate true - не создавать новых узлов при распознавании в рефлексах и т.п.
*/
func PhraseDetection(words []int, noCreate bool) string {
	if len(words) == 0 {
		return ""
	}
	if len(words) == 1 && words[0] == 0 {
		return ""
	} // пустые строки не писать
	CurrentVerbalPhraseEnd = nil
	DetectedUnicumPhraseID = 0
	// var pultOut=""
	DetectedCurrentPhraseID = 0
	currentStepPhraseCount = len(words)
	r := words
	// основа дерева
	cnt := len(VernikePhraseTree.Children)
	for n := 0; n < cnt; n++ {
		phraseNode := VernikePhraseTree.Children[n]
		rt := phraseNode.WordID
		if r[0] == rt {
			cldrn := VernikePhraseTree.Children[n] //.Children
			getPhraseTreeNode(r, &cldrn)
		}
		if currentStepPhraseCount == 0 {
			break
		} // распознанно точно, не смотреть другие
	}
	// результат распознавания
	if DetectedCurrentPhraseID > 0 {
		if currentStepPhraseCount == 0 { // полностью распознан
			DetectedUnicumPhraseID = DetectedCurrentPhraseID
		} else {
			var nr = len(r) - currentStepPhraseCount
			CurrentVerbalPhraseEnd = r[nr:]
		}
	}
	var needSave = false
	if DetectedUnicumPhraseID == 0 {
		// нераспознанный остаток
		if len(CurrentVerbalPhraseEnd) > 0 {
			r := CurrentVerbalPhraseEnd
			var tree *PhraseTree
			if DetectedCurrentPhraseID > 0 {
				//				tree = PhraseTreeFromID[DetectedCurrentPhraseID]
				node, ok := ReadePhraseTreeFromID(DetectedCurrentPhraseID)
				if ok {
					tree = node
				}
			} else {
				tree = &VernikePhraseTree
			}
			// просто добавить новую ветку - из диалога это стоит делать за 1 раз т.к. слова уже известны
			if !noCreate {
				node := createNewNodePhraseTree(tree, 0, r[0])
				tree = node
				pt, ok := ReadePhraseTreeFromID(tree.ID)
				if ok {
					id := createPhraseTreeNodes(r, pt)
					DetectedUnicumPhraseID = id
					needSave = true
				}
			}
		}
	}
	// нет вообще такого, добавить все слово
	if DetectedUnicumPhraseID == 0 {

		//tree := PhraseTreeFromID[0]
		tree, ok := ReadePhraseTreeFromID(0)
		if ok {
			// сразу создать первый узел
			if len(r) > 0 && !noCreate {
				node := createNewNodePhraseTree(tree, 0, r[0])
				tree = node
				if tree != nil {
					tn, ok := ReadePhraseTreeFromID(tree.ID)
					if ok {
						id := createPhraseTreeNodes(r, tn)
						DetectedUnicumPhraseID = id
						needSave = true
					}
				}
			}
		}
	}
	if needSave {
		SavePhraseTree()
	}
	out := GetPhraseStringsFromPhraseID(DetectedUnicumPhraseID)
	// заполнить PhraseTreeFromWordID
	finishScanAllTree()

	return out //pultOut+"{"+strconv.Itoa(DetectedUnicumPhraseID)+")"
}

// получить ID фразы - конечный узел дерева фраз
func getPhraseTreeNode(words []int, wt *PhraseTree) {
	if len(words) == 0 {
		return
	}
	ost := words[1:]
	if words[0] != wt.WordID {
		return
	} // пошло не туда
	DetectedCurrentPhraseID = wt.ID
	currentStepPhraseCount = len(ost)

	for n := 0; n < len(wt.Children); n++ {
		getPhraseTreeNode(ost, &wt.Children[n])
	}

	return
}

// получить число узлов в ветке
func getNodeCountFromLastID(lastID int) int {
	if lastID == 0 {
		return 0
	}
	var count = 0
	for {
		//		node := PhraseTreeFromID[lastID]
		node, ok := ReadePhraseTreeFromID(lastID)
		if !ok || node.WordID == 0 {
			break
		}
		count++
		lastID = node.ParentID
	}
	return count
}

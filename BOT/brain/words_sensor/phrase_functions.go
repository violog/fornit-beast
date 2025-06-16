/*  функции для фраз */

package word_sensor

import (
	"BOT/lib"
	"regexp"
	"strconv"
	"strings"
)

// слово из ID узла дерева фраз
func GetWordFromPraseNodeID(nodeID int) string {
	if nodeID == 0 {
		return ""
	}
	//	ph := PhraseTreeFromID[nodeID]
	ph, ok := ReadePhraseTreeFromID(nodeID)
	if !ok {
		return ""
	}
	word := GetWordFromWordID(ph.WordID)

	return word
}

// []ID слов из ID узла дерева фраз
func GetWordsIDarrFromPraseNodeID(nodeID int) []int {
	if nodeID == 0 {
		return nil
	}
	//	ph := PhraseTreeFromID[nodeID]
	ph, ok := ReadePhraseTreeFromID(nodeID)
	if !ok {
		return nil
	}
	var reversID []int
	reversID = append(reversID, ph.WordID)
	for ph.ParentNode != nil {
		reversID = append(reversID, ph.ParentNode.WordID)
		if ph.ParentNode.ParentNode == nil {
			break
		}
		ph.ParentNode = ph.ParentNode.ParentNode
	}
	var out []int
	for i := len(reversID) - 1; i >= 0; i-- {
		if reversID[i] > 0 {
			out = append(out, reversID[i])
		}
	}

	return out
}

////////////////////////////////////////////////

// []ID слов из ID фразы
func GetWordsArrFromPraseID(praseID int) []int {
	var wArr []int
	for {
		//node := PhraseTreeFromID[lastID]
		node, ok := ReadePhraseTreeFromID(praseID)
		if !ok {
			break
		}
		wArr = append(wArr, node.WordID)
		praseID = node.ParentID
		if praseID == 0 {
			break
		}
	}

	return wArr
}

/////////////////////////////////////////////////////

// строка из ID фразы дерева фраз
func GetPhraseStringsFromPhraseID(lastID int) string {
	var idArr []string

	for {
		//node := PhraseTreeFromID[lastID]
		node, ok := ReadePhraseTreeFromID(lastID)
		if !ok {
			break
		}
		w := GetWordFromWordID(node.WordID)
		idArr = append(idArr, w)
		lastID = node.ParentID
		if lastID == 0 {
			break
		}
	}

	var str = ""
	for i := len(idArr) - 1; i >= 0; i-- {
		if len(str) > 0 {
			str += " "
		}
		str += idArr[i]
	}

	return str
}

/*
	выдать строку из массива wordsArr[]int

используется в update_genom.go
*/
func GetStrFromArrID(wArr []int) string {
	var out = ""

	for i := 0; i < len(wArr); i++ {
		out += GetWordFromWordID(wArr[i]) + " "
	}

	return out
}

// очистить фразу от неалфавитных символов
func ClinerNotAlphavit(prase string) string {
	var out = ""

	reg := regexp.MustCompile(`[а-я ]`)
	res := reg.FindAllString(prase, -1)
	for i := 0; i < len(res); i++ {
		out += res[i]
	}

	return out
}

// если есть такая фраза в Дереве, то выдать ее ID - ТОЛЬКО РАСПОЗНАВАНИЕ, А НЕ СОЗДАНИЕ НОВОЙ
// var WordsRecognisedIdArr []int //массив слов, не короче 4 символов
func GetExistsPraseID(text string) int {
	var id = 0
	//	WordsRecognisedIdArr = nil

	// чистим лишние пробелы
	rp := regexp.MustCompile("s+")
	text = rp.ReplaceAllString(text, " ")
	text = strings.TrimSpace(text)

	wordsArr := GetWordIDfromPhrase(text, true)
	//	WordsRecognisedIdArr = GetWordIDfromPhrase4limit(text, 4) //СЛОВА НЕ КОРОЧЕ 4-x символов
	wArr := strings.Split(text, " ")
	// проверяем, число слов wArr должно быть равно числу распознанных
	if len(wordsArr) != len(wArr) {
		return 0 // фраза не распознана
	}
	id = GetExistsPraseIDFromWordArr(wordsArr)

	return id
}

////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////

// если есть такая фраза в Дереве, то выдать ее ID - ТОЛЬКО РАСПОЗНАВАНИЕ, А НЕ СОЗДАНИЕ НОВОЙ
func GetExistsPraseIDFromWordArr(wordsArr []int) int {
	var id = 0
	if wordsArr == nil || len(wordsArr) == 0 {
		return 0 // фразы нет
	}
	str := PhraseDetection(wordsArr, true) // распознаватель фразы
	if len(str) > 0 {
		id = DetectedUnicumPhraseID
	}
	return id
}

/////////////////////////////////////////////////////////////////////

// удалить слово во всех упоминаниях в Дереве фраз
func deleteWordFromPhrase(wordID int) {
	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_reflex/phrase_tree.txt")
	var out = ""
	var parentNewID = 0
	var parentOdID = 0

	for n := 0; n < len(strArr); n++ {
		p := strings.Split(strArr[n], "|")
		id, _ := strconv.Atoi(p[0])
		pID, _ := strconv.Atoi(p[1])
		//		node := PhraseTreeFromID[id]
		node, ok := ReadePhraseTreeFromID(id)
		if !ok {
			return
		}
		p = strings.Split(strArr[n], "|#|")
		wID, _ := strconv.Atoi(p[1])
		if wID == wordID {
			if len(node.Children) > 0 { // всем дочкам переписать родителей - node.ParentID
				parentNewID = node.ParentID
				parentOdID = node.ID
			} // если нет родителя, то можно просто удалить
			continue // не писать удаляемую строку
		}

		if pID > 0 && pID == parentOdID { // заменить родителя
			out += strconv.Itoa(id) + "|" + strconv.Itoa(parentNewID) + "|#|" + strconv.Itoa(wID) + "\r\n"
		} else {
			out += strArr[n] + "\r\n"
		}
	}
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_reflex/phrase_tree.txt", out)
}

// вытащить первый символ из фразы
func GetFirstSymbolFromPraseID(PhraseID []int) int {
	if len(PhraseID) == 0 {
		return 0
	}
	// аналогично
	// GetPhraseStringsFromPhraseID(PhraseID[0])
	lastID := PhraseID[0]
	word := ""
	for {
		//		node := PhraseTreeFromID[lastID]
		node, ok := ReadePhraseTreeFromID(lastID)
		if !ok {
			break
		}

		word = GetWordFromWordID(node.WordID)
		lastID = node.ParentID
		if lastID == 0 {
			break
		}
	}
	if len(word) == 0 {
		return 0
	}
	r := []rune(word)
	first := GetSymbolIDfromRune(r[0])
	// проверка
	// w:=GetWordFromWordID(wID); if len(w)>0{}
	// s := GetSynbolFromID(first); if len(s) > 0 { }

	return first
}

/////////////////////////////////////////////

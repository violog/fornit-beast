/*
Распознаватель слов по символьно для заполнения дерева слов word_tree.go
и распознавания слов при вводе с пульта
*/

package word_sensor

import (
	"BOT/lib"
	"fmt"
	"strings"
)

func wordRecognizerInit() {
	//	WordDetection("играть") // распознавание слова
}

// текущий уникальный ID последней активной ветки дерева - результат детекции фразы - для дальнейшего использования
var DetectedUnicumID = 0
var FirstSimbolID = 0       // при активации дерева слов
var CurrentVerbalEnd []rune // нераспознанный остаток
var detectedCurrentID = 0   // текущий найденный ID последней активной ветки дерева слов
var currentStepCount = 0    // текущеее число нераспознанных символов

/*
	var lastFoundID=0 // последний ID символа при проходе дерева

проход одного слова - распознавание слова
возвращает найденное ID слова или похожей альтернативы

noCreate true - не создавать новых узлов при распознавании в рефлексах и т.п.
*/
func WordDetection(word string, noCreate bool) int {
	word = strings.TrimSpace(word)
	if len(word) == 0 {
		return 0
	}

	CurrentVerbalEnd = []rune("")
	DetectedUnicumID = 0
	// var pultOut=""
	detectedCurrentID = 0
	currentStepCount = 0

	// попробовать найти подходящее слово по готовому массиву ID слов
	DetectedUnicumID = tryWordRecognize(word)
	if DetectedUnicumID > 0 {
		return DetectedUnicumID
	}

	/* Дальше пойдет только если на пульте включен "режим форсированной обработки".
	   Иначе слова заходят в дерево при накоплении в words_temp_arr.txt
	   	if !NoCheckWordCount {
	   //		updateWordTreeFromTempArr(2,4)
	   		return 0
	   	}
	*/
	r := []rune(word)
	// основа дерева
	cnt := len(VernikeWordTree.Children)
	var curFirstLevelID = 0
	for n := 0; n < cnt; n++ {
		smblNode := VernikeWordTree.Children[n]
		rt := []rune(smblNode.Symbol)
		if r[0] == rt[0] {
			FirstSimbolID = GetSymbolIDfromRune(r[0])
			curFirstLevelID = smblNode.ID
			if len(r) == 1 { // это - символ, присвоить слову ID символа
				DetectedUnicumID = VernikeWordTree.Children[n].ID
				CurrentVerbalEnd = []rune("")
				return DetectedUnicumID
			}

			cldrn := VernikeWordTree.Children[n].Children
			cnt := len(cldrn)
			for n := 0; n < cnt; n++ {
				getWordTreeNode(r[1:], &cldrn[n])
			}
		}
	}
	// результат распознавания
	if detectedCurrentID > 0 {
		if currentStepCount == 0 { // полностью распознан
			DetectedUnicumID = detectedCurrentID
		} else {
			var nr = len(r) - currentStepCount
			CurrentVerbalEnd = r[nr:]
		}
	}

	var needSave = false
	if DetectedUnicumID == 0 {
		/* лучше это делать до прохода дерева потому как сравнивается только со старыми словами, переместил
		// если обычный режим диалога (на ПУльте не стоит галка "набивка рабочих фраз без отсеивания мусорных слов ")
		if !NoCheckWordCount {
			//отсеивать мусорных (редких - менее 4 повторов в tempArr) слов
			repeet := getWordTemparrCount(text)
			if repeet < 4 {
				// попробовать найти подходящее слово
				DetectedUnicumID = getAlternative(text)
				if DetectedUnicumID != 0 {
					return DetectedUnicumID
				}
			}
		}
		*/

		// нераспознанный остаток
		if len(CurrentVerbalEnd) > 0 {
			r := CurrentVerbalEnd
			var tree *WordTree
			if detectedCurrentID > 0 {
				tree = WordTreeFromID[detectedCurrentID]
			} else {
				tree = &VernikeWordTree
			}
			// просто добавить новую ветку - из диалога это стоит делать за 1 раз т.к. слова уже известны
			if !noCreate {
				id := createWordTreeNodes(r, tree)
				DetectedUnicumID = id
				WordIdFormWord[word] = id // добавить в список слов
				// SaveWordTree()
				needSave = true
			}
		}
	}

	// нет вообще такого, добавить все слово, пройдя первый символьный уровень дерева
	if DetectedUnicumID == 0 {
		tree := WordTreeFromID[curFirstLevelID] //пройдя первый символьный уровень дерева
		ost := r[1:]
		if len(ost) > 0 && !noCreate {
			// добавить в дерево остальные символы, начиная с узла node
			id := createWordTreeNodes(ost, tree) //WordTreeFromID[tree.ID]
			DetectedUnicumID = id
			WordIdFormWord[word] = id // добавить в список слов
			//SaveWordTree()
			needSave = true
		}
	}
	if needSave {
		SaveWordTree()
	}

	return DetectedUnicumID //pultOut+"{"+strconv.Itoa(DetectedUnicumID)+")"
}

// cканирование следует строго по нужной ветке
func getWordTreeNode(word []rune, wt *WordTree) {
	if len(word) == 0 {
		return
	}

	ost := word[1:]
	if string(word[0]) != wt.Symbol {
		return
	} // пошло не туда

	detectedCurrentID = wt.ID
	currentStepCount = len(ost)
	for n := 0; n < len(wt.Children); n++ {
		getWordTreeNode(ost, &wt.Children[n])
	}
}

/*
	найти слово в tempArr и выдать его повторяемость, если слова нет в tempArr - добавить его.
	слово еще раньше добавляется в tempArr если его там нет

так что оно уже обязательно там будет
*/
func getWordTemparrCount(word string) int {
	lib.MapCheckBlock(MapGwardWordTempArr)
	for k, v := range tempArr {
		if k == word {
			lib.MapFree(MapGwardWordTempArr)
			return v[0]
		}
	}
	lib.MapFree(MapGwardWordTempArr)
	return 0 // на всякий случай
}

/*
	ранее - не найдено в WordIdFormWord[word]

попробовать найти подходящее слово с альтрнативным ID
Первые буквы должны совпадать, а остальные, кроме последней (разные окончания),
быть перемешаны, но в наличии >80%.
Сканирует дерево с начальной буквы строго по числу чимволов слова.
Это имитирует свойство персептронного распознавателя.
*/
func getAlternative(word string) int {
	defer func() { // ловим панику
		if err := recover(); err != nil {
			fmt.Println(err) // просто выдать сообщение чтобы выловить с breakpointe здесь и пройти по стеку
		}
	}()
	rw := []rune(word)
	var rwLen = len(rw)
	if rwLen < 4 {
		return 0
	}
	// выбрать известные слова с первой и последней буквой как у word
	var wArr = make(map[int][]rune)
	for w, id := range WordIdFormWord {
		r := []rune(w)
		rLen := len(r)
		if rLen < 3 || r[0] != rw[0] || rLen != rwLen {
			continue
		}
		if rw[rwLen-1] == r[rLen-1] {
			r0 := r[1:]
			r0 = r0[:(rLen - 1)]
			wArr[id] = r0
		}
	}
	if len(wArr) == 0 {
		return 0
	}
	// проверять все ли внутренние буквы совпадают
	rw0 := rw[1:]
	rw0 = rw0[:(rwLen - 1)]
	for id, r := range wArr {
		// сопоставляются только внутренние части слов (без первой и последней букв)
		if isEquivalented(r, rw0) {
			return id
		}
	}

	return 0
}

// все внутренние буквы должны присуствовать
func isEquivalented(r1 []rune, r2 []rune) bool {
	for n := 0; n < len(r1); n++ {
		var isAbsent = 1
		for m := 0; m < len(r2); m++ {
			if r1[n] == r2[m] {
				isAbsent = 0
				break
			}
		}
		if isAbsent == 1 {
			return false
		}
	}
	return true
}

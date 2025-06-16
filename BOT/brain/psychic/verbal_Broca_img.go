/* Словестные образы (область Брока)  для 4, 5 и 6-го уровня дерева автоматизмов.
Смысл (осознанную значимость) образ приобретеает только в контексте Дерева Понимания (дерева мент.автоматизмов)

Детекторы зоны Вернике распознают слова и словосочетания,
а область Брока отвечает за смысл распознанных слов и словосочетений,
за конструирование собственных словосочетаний,
за моторное использование слов и словосочетаний.
За все ответственная структура - образ осмысленных слов и сочетаний.

! Нужно иметь в виду, что в Vernike_detector.go есть массив памяти фраз, накапливается в течении дня
var MemoryDetectedArr []MemoryDetected - структур фразы с контекстным окружением
и Verbal.PhraseID[] - можно найти в этом массиве для ориентировки что бы ло раньше и позже.
MemoryDetectedArr - как бы оперативная память фраз для сопоставлений.
*/

package psychic

import (
	wordSensor "BOT/brain/words_sensor"
	"BOT/lib"
	"strconv"
	"strings"
)

/////////////////////////////////////////////

func verbalInit() {
	loadVerbalFromIdArr()

	/*
	   var tm=922// "Обычный, Хорошее"
	   	str:=getToneMoodStrFromID(tm)
	   	if len(str)>0{	}
	*/

}

////////////////////////////////////////////

/*
	для оптимизации поиска по дереву перед узлом Verbal идет узел первого символа : var symbolsArr из word_tree.go

Смысл (осознанную значимость) образ приобретеает только в контексте Дерева Понимания (дерева мент.автоматизмов)
Если член PhraseID[] == -1 - фраза есть, но нераспознана.

Используется в purpose_genetic.go и rules_functions.go
Создается при активации дерева автоматизмов, НО ТАМ ВЕТКА PhraseID int//!!!! на самом деле пишется lev6 = verb.PhraseID[0] !!!
*/
type Verbal struct {
	ID int
	// для текущего сообщения с Пусльта:
	SimbolID int   // id первого символа первой фразы PhraseID: var symbolsArr из word_tree.go
	PhraseID []int // массив фразID (DetectedUnicumPhraseID) слова каждой фразы вытаскиваются wordSensor.GetWordArrFromPhraseID(PhraseID[0])
	//0 - обычный, 1 - восклицательный, 2- вопросительный, 3- вялый, 4 - Повышенный
	ToneID int // тон сообщения с Пульта
	//0 - обычный, 1 (из 20)-Хорошее    2 (21)-Плохое    3(22)-Игровое    4(23)-Учитель
	//5(24)-Агрессивное   6(25)-Защитное    7(26)-Протест
	MoodID int // настроение оператора
}

// var VerbalFromIdArr=make(map[int]*Verbal)
var VerbalFromIdArr []*Verbal // сам массив
// var AFromID = make([]*aNode, len(strArr))//задать сразу имеющиеся в файле число при загрузке
// запись члена
func WriteVerbalFromIdArr(index int, value *Verbal) {
	addVerbalFromIdArr(index)
	VerbalFromIdArr[index] = value
}
func addVerbalFromIdArr(index int) {
	if index >= len(VerbalFromIdArr) {
		newSlice := make([]*Verbal, index+1)
		copy(newSlice, VerbalFromIdArr)
		VerbalFromIdArr = newSlice
	}
}

// считывание члена
func ReadeVerbalFromIdArr(index int) (*Verbal, bool) {
	if index >= len(VerbalFromIdArr) || VerbalFromIdArr[index] == nil {
		return nil, false
	}
	return VerbalFromIdArr[index], true
}

//////////////////////////////////////////
// для поиска по ID фразы
//var VerbalFromPhraseIdArr=make(map[int]*Verbal)
///////////////////////////////////////////

// создать образ сочетаний пусковых стимулов
// В случае отсуствия пусковых стимулов создается ID такого отсутсвия, пример такой записи: 2|||0|0| - ID=2
var lastVerbalID = 0

func createNewlastVerbalID(id int, SimbolID int, PhraseID []int, ToneID int, MoodID int, CheckUnicum bool) (int, *Verbal) {
	if MoodID > 19 {
		MoodID = MoodID - 19
	} else {
		MoodID = 0
	}
	if CheckUnicum {
		oldID, oldVal := checkUnicumVerbal(PhraseID, ToneID, MoodID)
		if oldVal != nil {
			return oldID, oldVal
		}
	}

	if id == 0 {
		lastVerbalID++
		id = lastVerbalID
	} else {
		//		newW.ID=id
		if lastVerbalID < id {
			lastVerbalID = id
		}
	}

	var node Verbal
	node.ID = id
	node.SimbolID = SimbolID
	node.PhraseID = PhraseID
	node.ToneID = ToneID
	node.MoodID = MoodID

	//VerbalFromIdArr[id]=&node
	WriteVerbalFromIdArr(id, &node)

	return id, &node
}
func checkUnicumVerbal(PhraseID []int, ToneID int, MoodID int) (int, *Verbal) {
	for id, v := range VerbalFromIdArr {
		if v == nil {
			continue
		}
		if !lib.EqualArrs(PhraseID, v.PhraseID) {
			continue
		}
		if ToneID != v.ToneID || MoodID != v.MoodID {
			continue
		}
		return id, v
	}

	return 0, nil
}

// ///////////////////////////////////////
// создать новый вербальный образ, если такого еще нет
func CreateVerbalImage(FirstSimbolID int, PhraseID []int, ToneID int, MoodID int) (int, *Verbal) {
	if PhraseID == nil || isUnrecognizedPhraseFromAtmtzmTreeActivation {
		return 0, nil
	}
	// достаем первый символ первой фразы
	// получить последователньость wordID из уникального идентификатора первой фразы
	/*
	   	wordIDarr:=wordSensor.GetWordArrFromPhraseID(PhraseID[0])
	   	// первое слово в виде строки
	   	if wordIDarr==nil || len(wordIDarr)==0{
	   		return 0,nil
	   	}
	   	word:=wordSensor.GetWordFromWordID(wordIDarr[0])
	   	//rw:=[]rune(word)
	   	//SimbolID:=wordSensor.GetSymbolIDfromRune(rw[0])
	   //	word:=wordSensor.GetPhraseStringsFromPhraseID(PhraseID[0])
	   	//SimbolID:=wordSensor.GetSymbolIDfromString(rw[0])
	*/
	id, verb := createNewlastVerbalID(0, FirstSimbolID, PhraseID, ToneID, MoodID, true)

	if doWritingFile {
		SaveVerbalFromIdArr()
	}

	return id, verb
}

/////////////////////////////////////////

// ////////////////// сохранить образы сочетаний пусковых стимулов
// В случае отсуствия пусковых стимулов создается ID такого отсутсвия, пример такой записи: 2|||0|0|
func SaveVerbalFromIdArr() {
	var out = ""
	for k, v := range VerbalFromIdArr {
		if v == nil {
			continue
		}
		out += strconv.Itoa(k) + "|"
		out += strconv.Itoa(v.SimbolID) + "|"
		for i := 0; i < len(v.PhraseID); i++ {
			out += strconv.Itoa(v.PhraseID[i]) + ","
		}
		out += "|"
		out += strconv.Itoa(v.ToneID) + "|"
		out += strconv.Itoa(v.MoodID) + "|"
		out += "\r\n"
	}
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/verbal_images.txt", out)

}

// //////////////////  загрузить образы сочетаний пусковых стимулов
func loadVerbalFromIdArr() {

	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/verbal_images.txt")
	cunt := len(strArr)
	//VerbalFromIdArr=make(map[int]*Verbal)
	VerbalFromIdArr = make([]*Verbal, cunt) //задать сразу имеющиеся в файле число
	for n := 0; n < cunt; n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		p := strings.Split(strArr[n], "|")
		id, _ := strconv.Atoi(p[0])
		SimbolID, _ := strconv.Atoi(p[1])
		s := strings.Split(p[2], ",")
		var PhraseID []int
		for i := 0; i < len(s); i++ {
			if len(s[i]) == 0 {
				continue
			}
			si, _ := strconv.Atoi(s[i])
			PhraseID = append(PhraseID, si)
		}
		ToneID, _ := strconv.Atoi(p[3])
		MoodID, _ := strconv.Atoi(p[4])
		var saveDoWritingFile = doWritingFile
		doWritingFile = false
		createNewlastVerbalID(id, SimbolID, PhraseID, ToneID, MoodID, false)
		doWritingFile = saveDoWritingFile
	}
	return

}

//////////////////////////////

/*
	получить уникальное сочетание в виде int из двух компонентов int

На входе int2 -  виде ID настроения (jn 20 до 26) преобразуется в диапазон от 1 до 7
простым вычитанием int2-19
*/
func GetToneMoodID(int1 int, int2 int) int {
	// вмето первой 0 (для "обычный") ставим 9 !!!
	if int1 == 0 {
		int1 = 9
	}
	s := strconv.Itoa(int1)
	if int2 >= 19 {
		int2 -= 19
	}
	s += strconv.Itoa((int2))
	ToneMoodID, _ := strconv.Atoi(s)
	return ToneMoodID
}

// ////////////////////////////////
// получить тон и настроение из уникального сочетания
func GetToneMoodFromImg(img int) (int, int) {
	tonmoode := strconv.Itoa(img)
	var t = 0
	ton := tonmoode[:1]
	if ton == "9" {
		t = 0
	} else {
		t, _ = strconv.Atoi(ton)
	}
	m, _ := strconv.Atoi(tonmoode[1:])
	return t, m
}

// расшифровка в виде строки
func GetToneMoodStrFromID(img int) string {
	t, m := GetToneMoodFromImg(img)
	out := "<b>" + getToneStrFromID(t) + " - "
	out += getMoodStrFromID(m) + "</b>"

	return out
}

//////////////////////////////////////////////////

// /////////////////////////////
func getToneStrFromID(id int) string {
	var ret = ""
	//0 - обычный, 1 - восклицательный, 2- вопросительный, 3- вялый, 4 - Повышенный
	switch id {
	case 0:
		ret = "обычный"
	case 1:
		ret = "восклицательный"
	case 2:
		ret = "вопросительный"
	case 3:
		ret = "вялый"
	case 4:
		ret = "Повышенный"
	default:
		ret = "ощибка (тон = " + strconv.Itoa(id) + ")"
	}
	return ret
}

// //////////////////////////////
func getMoodStrFromID(id int) string {
	var ret = ""
	// из 20-Хорошее    21-Плохое    22-Игровое    23-Учитель    24-Агрессивное   25-Защитное    26-Протест
	// отнимаем 19
	switch id {
	case 0:
		ret = "Нормальное"
	case 1:
		ret = "Хорошее"
	case 2:
		ret = "Плохое"
	case 3:
		ret = "Игровое"
	case 4:
		ret = "Учитель"
	case 5:
		ret = "Агрессивное"
	case 6:
		ret = "Защитное"
	case 7:
		ret = "Протест"

	}
	return ret
}

////////////////////////////////

// выдать инфу про PhraseID
func GetStringsFromVerbalID(vID int) string {

	//	v:=VerbalFromIdArr[vID]
	v, ok := ReadeVerbalFromIdArr(vID)
	if !ok {
		lib.WritePultConsol("Нет карты VerbalFromIdArr для iD=" + strconv.Itoa(vID))
		return "Ошибка получения информации."
	}
	out := "Вербальный образ ID=" + strconv.Itoa(vID)

	out += "<br> Фраза: "
	for i := 0; i < len(v.PhraseID); i++ {
		if i > 0 {
			out += " | "
		}
		out += wordSensor.GetPhraseStringsFromPhraseID(v.PhraseID[i]) + "<br>"
	}
	out += "Тон: " + getToneStrFromID(v.ToneID) + "<br>"
	out += "Настроение: " + getMoodStrFromID(v.MoodID) + "<br>"
	return out
}

// ///////////////////////////////////////////////////
// выдать инфу про PhraseID
func GetPraseStringsFromVerbalID(vID int) string {

	//	v:=VerbalFromIdArr[vID]
	v, ok := ReadeVerbalFromIdArr(vID)
	if !ok {
		lib.WritePultConsol("Нет карты VerbalFromIdArr для iD=" + strconv.Itoa(vID))
		return "Ошибка получения информации."
	}
	out := ""
	for i := 0; i < len(v.PhraseID); i++ {
		if i > 0 {
			out += " | "
		}
		out += wordSensor.GetPhraseStringsFromPhraseID(v.PhraseID[i]) + "<br>"
	}

	return out[:len(out)-4]
}

/////////////////////////////////////////////////////

// ПРОИЗВОЛЬНОЕ ТВОРЕНИЕ ФРАЗ

/*
	сделать фразу PhraseID из wordID []

Возвращает PhraseID
*/
func AddwordIDToPhraseTree(wordID []int) []int {
	// засунуть фразу в дерево слов и дерево фраз
	// проход одной фразы - распознавание ID слов фразы
	wordSensor.PhraseDetection(wordID, false)
	PhraseID := wordSensor.CurrentPhrasesIDarr

	// первый символ ответной фразы
	FirstSimbolID := wordSensor.GetFirstSymbolFromPraseID(PhraseID)
	// создать образ Брока
	CreateVerbalImage(FirstSimbolID, PhraseID, 0, 0)

	return PhraseID
}

////////////////////////////////////////////

/*
	сделать фразу PhraseID из string

Возвращает PhraseID
*/
func AddStringToPhraseTree(str string) []int {
	// засунуть фразу в дерево слов и дерево фраз
	wordSensor.VerbalDetection(str, 0, 0, 0)
	PhraseID := wordSensor.CurrentPhrasesIDarr

	// первый символ ответной фразы
	FirstSimbolID := wordSensor.GetFirstSymbolFromPraseID(PhraseID)
	// создать образ Брока
	CreateVerbalImage(FirstSimbolID, PhraseID, 0, 0)

	return PhraseID
}

////////////////////////////////////////////

// есть ли в данном Verb фраза praseID int
func existsPraseIDinVerbID(verbID int, praseID int) bool {
	verb, ok := ReadeVerbalFromIdArr(verbID)
	if !ok {
		return false
	}
	for n := 0; n < len(verb.PhraseID); n++ {
		if verb.PhraseID[n] == praseID {
			return true
		}
	}
	return false
}

////////////////////////////////////////////////////

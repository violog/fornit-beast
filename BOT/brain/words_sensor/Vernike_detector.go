/*  распознаватель слов и фраз по типу зоны Вернике в мозге.
Память о воспринятых фразах в текущем активном контексте (Vernike_detector.go): var MemoryDetectedArr []MemoryDetected

Распознавание фраз начинается в main.go с word_sensor.VerbalDetection(text_dlg, is_input_rejim, moodID)
С ПУльта приходит текст, который в VerbalDetection() разбирается на фрзацы (\r\n):

абзацы в PhraseSeparator() разбираются на фразы по разделителям (знаки препинания)
фразы в WordDetection() разбиваются на слова.
Распознанные (и нераспознанные) последовательности сохраняются в оперативной памяти Beast MemoryDetectedArr.
где распознанный текст представлен в виде уникального laslID фразы

ОПИСКИ при вводе слова. Если слово не распознается и оно имеет более 3-х символов,
то делается предположение об описке внутренних символов
(в природном распознавателе слово узнается если точно совпали первая и последняя буквы,
а внутренние буквы могут быть как угодно перемешаны)
Если слово распознается, то подставляется ID слова.

Нераспознанной фразы НЕ БЫВАЕТ т.к. она тут же создается

Тон фразы можно задать 1) с помощью знаков ! и ? в конце фразы
или задать преимущественно - выбрав Тон под окном ввода фразы.
*/

package word_sensor

import (
	_ "BOT/lib"
	"regexp"
	"strconv"
	"strings"
)

// запрет показа карт WordTreeFromID и PhraseTreeFromID во время распознавания и записи
// против паники типа "одновременная запись и считывание карты"
var notAllowScanInThisTime = false

// индикация, что дерево загружено, можно вводить тексты
var isReadyWordSensorLevel = 0

func IsReadyWordSensorLevel() bool {
	if isReadyWordSensorLevel > 0 {
		return true
	} // связь с корнями проекта
	return false
}

// word_sensor.VerbalDetectin("активностный")
// для использования в SetNewWordTreeNode

/*
	в CurrentPhrasesIDarr набирается массив lastID распознанных фраз для дерева автоматизмов

При активации дерева автоматизмов:
if len(wordSensor.CurrentPhrasesIDarr) > 0 { ...

Если CurrentPhrasesIDarr[n]==-1 - фраза есть, но она нераспознана.
*/
var CurrentPhrasesIDarr []int
var CurretWordsIDarr []int // массив словID и вместо нераспознанного -1

// тон сообщения с Пульта при передаче фразы
var CurPultTone = 0

// настроение с Пульта при передаче фразы
var CurPultMood = 0

// текущий тон фразы: 0-обычный, 1-восклицательный, 2-вопросительный
var DetectedTone = 0

/* Память о воспринятых фразах в текущем активном контексте:
В принципе это не нужно, т.к. имеется эпизодическая память
(правда только для Правил, но это как раз и всязывает фразы с их полезностью)

type MemoryDetected struct {
	//распознанный текст в виде lastID выделенных фраз
	PhrasesID []int // массив структур распознанных фраз
	Tone      int   // Тон: 0-обычный, 1-восклицательный, 2-вопросительный, 3-вялый, 4-Повышенный
	Mood      int   // настроение при передаче фразы с Пульта: 20-Хорошее 21-Плохое 22-Игровое 23-Учитель 24-Агрессивное 25-Защитное 26-Протест
	// индекс==ID активного базового контекста, значение - вес этого контекста
	ActveContextWeight map[int]int
}

// добавить строку в массив памяти о воспринятых фразах в текущем активном контексте
func addNewMemoryDetected() {
	var newM MemoryDetected
	newM.PhrasesID = CurrentPhrasesIDarr
	// тон может указываться 1) в виде ! или ? во фразе - DetectedTone 	И/ИЛИ 2) в виде радиокнопки Тон с Пульта - CurPultTone
	var tone = 0
	if DetectedTone > 0 { // преимущество - у задатчика тона 2)
		tone = DetectedTone
	} else { // есть ! или ? во фразе
		tone = CurPultTone
	}
	newM.Tone = tone
	newM.Mood = CurPultMood // настроение с Пульта: повышенный нормальный вялый Хорошее Плохое Игровое Учитель Агрессивное Защитное Протест
	newM.ActveContextWeight = gomeostas.GetActiveContextInfo()
	MemoryDetectedArr = append(MemoryDetectedArr, newM)
}

// массив памяти накапливается в течении дня, обрабатывается и очищается во сне
var MemoryDetectedArr []MemoryDetected
*/

/*  распознавание фразы с Пульта - бывает только в нижнем регистре
 */
var wlev = 0

// // если обычный режим диалога (на ПУльте не стоит галка "набивка рабочих фраз без отсеивания мусорных слов ")
var NoCheckWordCount = false // is_input_rejim - набивка рабочих фраз с отсеиванием мусорных слов

/*
	вызывается фразой с Пульта

text="привет|!|#|хочу| |научить| |тебя| |простым| |вещам|.|#|для| |начала| |давай| |усвоим| |длинные| |тексты|.|#|в| |них| |могут| |попадаться| |знакомые| |значащие| |слова| |вроде| |улыбнуться|,| |давай| |учить|,| |бестолочь|.|#|ты| |должен| |найти| |их|,| |выделить| |правила| |и| |обобщить| |слова| |в| |ответное| |общее| |действие|.|"

func UpdateWordTreeFromTempArr //переносим в дерево слов достаточно повторяющиеся из tempArr
NoCheckWordCount = true //игнорировать getWordTemparrCount и всегда распознавать слова
*/
var NeedCheckTempList = false // был ввод фразы с пульта - обработать в pula.go func pulsActions() в свободное время.
func VerbalDetection(text string, noFilterInput int, toneID int, moodID int) string {
	NeedCheckTempList = true
	notAllowScanInThisTime = true // запрет показа карты при обновлении
	NoCheckWordCount = false
	CurrentPhrasesIDarr = nil
	CurretWordsIDarr = nil
	if noFilterInput == 0 { // 0 - это набивка рабочих фраз без отсеивания мусорных слов
		// игнорировать getWordTemparrCount и всегда распознавать слова
		NoCheckWordCount = true
	}
	CurPultTone = toneID
	CurPultMood = moodID

	pultOut := ""

	// стандартно разделить текст на короткие фразы, отправить на накопление в /memory_reflex/words_temp_arr.txt
	// разделяем на фразы, получая strArr: "привет| |о| |тупой"
	strArr := strings.Split(text, "|#") // а не |#| - чтобы оставлять разделитель "|"
	for i := 0; i < len(strArr); i++ {
		//if i > 0 { pultOut += "<br>" }

		// добавление в файл /memory_reflex/words_temp_arr.txt накопления новых слов и фраз
		if !NoCheckWordCount { // накапливать статистику если не задан режим распознавания сразу
			praseID, pultOut := addNewtempArr(strArr[i])
			if praseID == 0 { // фраза нераспознана
				/*	м.б. в таком случае подставлять наиболее похожую фразу, как это сделал бы персептрон при распознавании
					но смысл изменяется буквально он любого слова и знака, так что не нужно, просто - сигнал НЕ ПОНЯЛ
				*/
				CurrentPhrasesIDarr = append(CurrentPhrasesIDarr, -1)
			}
			// в func UpdateWordTreeFromTempArr() будут учтены новые слова и фразы, встратившиеся достаточное число раз
			// ответ на Пульт:
			notAllowScanInThisTime = false
			return pultOut
		} else { // распознавать за 1 раз - просто проход фразы с распознаванием

			text0 := strings.Replace(strArr[i], "|", "", -1)
			pultOut += PhraseSeparator(text0)
			if DetectedUnicumPhraseID > 0 { // распознанная фраза
				CurrentPhrasesIDarr = append(CurrentPhrasesIDarr, DetectedUnicumPhraseID)
				pultOut += "<br>"
				//return pultOut
			}
			// ответ на Пульт:
			// но нераспознанной фразы НЕ БЫВАЕТ т.к. она тут же создается
			//return "Фраза не распознана."
		}
		//pultOut += "<br>"
	} //for i := 0; i < len(strArr); i++ {
	// добавить в стек памяти распознанных
	//addNewMemoryDetected()

	// reflexes.ActiveFromPhrase() // активировать дерево рефлексов фразой - только для условных рефлексов
	// ответ на Пульт:
	//	notAllowScanInThisTime = false
	return pultOut
}

// проход одной фразы (т.е. по разделителям в предложении, а не по \r\n)
func PhraseSeparator(text string) string {
	var pultOut = ""
	// чистим лишние пробелы
	rp := regexp.MustCompile("s+")
	text = rp.ReplaceAllString(text, " ")
	text = strings.TrimSpace(text)
	wordsArr := GetWordIDfromPhrase(text, false) // распознаватель слов
	//	time.Sleep(1000 * time.Microsecond)
	str := PhraseDetection(wordsArr, false) // распознаватель фразы
	if DetectedUnicumPhraseID > 0 {
		pultOut += str + "(" + strconv.Itoa(DetectedUnicumPhraseID) + ")"
	}

	// тон сообщения
	DetectedTone = 0
	if strings.Contains(text, "!") {
		DetectedTone = 1
	}
	if strings.Contains(text, "?") {
		DetectedTone = 2
	}

	return pultOut
}

/*
	получить последователньость wordID из уникального идентификатора фразы CurrentPhrasesIDarr[i]

начиная с любого узла дерева Фраз (не обязательно конечного!) - к первому узлу ветки
*/
func GetWordArrFromPhraseID(PhraseNodeID int) []int {
	var wArr []int
	// пройти фразу от последнего слова до первого
	//	w:=PhraseTreeFromID[PhraseNodeID]
	w, ok := ReadePhraseTreeFromID(PhraseNodeID)
	if !ok {
		return nil
	}
	wArr = append(wArr, w.ID)
	for w.ParentID > 0 {
		pht, ok := ReadePhraseTreeFromID(w.ParentID)
		if ok {
			wArr = append(wArr, pht.ID)
			w = pht
		}
	}

	// нельзя так делать, уходит в бесконечный цикл: обращение к w.ID возвращает предыдущее значение
	/*	for w.ParentID > 0 {
		//w=PhraseTreeFromID[w.ParentID]
		w, ok := ReadePhraseTreeFromID(w.ParentID)
		if ok {
			wArr = append(wArr, w.ID)
		}
	}*/
	// восстановить порядок слов
	var wordIDarr []int
	for i := len(wArr) - 1; i >= 0; i-- {
		wordIDarr = append(wordIDarr, PhraseTreeFromID[wArr[i]].WordID)
	}
	return wordIDarr
}

////////////////////////////////////////

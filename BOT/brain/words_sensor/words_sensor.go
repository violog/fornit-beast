/* сенсор символов, слов и фраз */
package word_sensor

import (
	"BOT/lib"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
	вспомогательный шаблон воспринятых слов с числом повторений восприятия

Индекс - фраза, [0] = число повторнеий [1] - сколько раз уже сохраняется в списке

Фомат words_temp_arr.txt: Число повторнеий|#|Фраза|#|число сохранений (старость)

Число сохранений - в func UpdateWordTreeFromTempArr при пересохранении страых слов, фраз списков
*/
var tempArr = make(map[string][]int)
var MapGwardWordTempArr = lib.RegNewMapGuard()

///////////////////////////////////////

func init() {
	// залить сохранненный tempArr
	//!!!loadTempArr()
	afterLoadTempArr()
	//GetExistsPraseID("дурак")
}

// сохранить шаблон воспринятых слов
func SaveTempArr() {
	var str []string
	/* Вы не поверите, но при переборе tempArr вдруг очищается tempArr!?
	for k, v := range tempArr {
		str = append(str, k + "|#|" + strconv.Itoa(v[0]) + "|#|" + strconv.Itoa(v[1]))
	}
	А вот если разделить действи как ниже, то очищается, но нормально записывает в str. Жуть...
	*/
	lib.MapCheckBlock(MapGwardWordTempArr)
	for s, v := range tempArr {
		val0 := v[0]
		val1 := v[1]
		fileS := s + "|#|" + strconv.Itoa(val0) + "|#|" + strconv.Itoa(val1)
		str = append(str, fileS)
	}
	lib.MapFree(MapGwardWordTempArr)

	sort.Strings(str)
	var out = ""
	for i := 0; i < len(str); i++ {
		p := strings.Split(str[i], "|#|")
		out += p[1]
		out += "|#|" + p[0]
		if len(p) > 2 { // новый член формата
			out += "|#|" + p[2]
		} else {
			out += "|#|0"
		}
		//		out+= strconv.Itoa(int(gomeostas.LifeTime/100))+ "|#|"
		out += "\r\n"
	}

	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_reflex/words_temp_arr.txt", out)
}

// загрузить шаблон воспринятых слов
func loadTempArr() {
	lib.MapCheckWrite(MapGwardWordTempArr)
	tempArr = make(map[string][]int)
	lib.MapFree(MapGwardWordTempArr)
	wArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_reflex/words_temp_arr.txt")
	for n := 0; n < len(wArr); n++ {
		if len(wArr[n]) < 4 {
			continue
		}
		p := strings.Split(wArr[n], "|#|")
		v, _ := strconv.Atoi(p[0])
		c := 0
		if len(p) > 2 { // новый член формата
			c, _ = strconv.Atoi(p[2])
		}
		lib.MapCheckWrite(MapGwardWordTempArr)
		tempArr[p[1]] = []int{v, c}
		lib.MapFree(MapGwardWordTempArr)
	}
}

// Новая порция текста для формирования дерева слов
// Эта функция работает при накачке текстов из main.go
func SetNewTextBlock(txt string) string {
	/*
		txt, err := url.QueryUnescape(txt)
		if err != nil {
			log.Fatal(err)
			return "ОШИБКА раскодировки"
		}
	*/
	txt = strings.Replace(txt, "{#1}", "%", -1)
	// txt= strings.Replace(txt, "{#2}", "", -1)// кавычки просто очищены (пусть будет афазия :)
	var res = ""
	// разделяем на фразы
	strArr := strings.Split(txt, "|#") // а не |#| - чтобы оставлять разделитель "|"
	for i := 0; i < len(strArr); i++ {
		addNewtempArr(strArr[i] + "|")
	}
	UpdateWordTreeFromTempArr(4, 6)

	return res
}

/*
	добавляются как целиком фраза, так и все слова во фразе

Тут же дозаполняетс дерево слов уже многократно провторяющимися элементами.

Вернуть разпознанное ID фразы и саму фразу (режим диалога с Пультом, а не режим накачки текстами)
Если фраза не распознана, то возвращается 0 и фраза с нераспознанными словами.

При превышении числа накопленных слов в func UpdateWordTreeFromTempArr
происходит их распознавание и удаление из /memory_reflex/words_temp_arr.txt

Здесь str - с разделителями '|' (привет| |о| |тупой)
*/
func addNewtempArr(str string) (int, string) {
	loadTempArr() // вский раз получать обновленный массив накоплений слов

	//не добавлять в words_temp_arr.txt, если такая фраза (str) есть
	// убрать разделители '|'
	str0 := strings.Replace(str, "|", "", -1)
	// если есть такая фраза в Дереве, то выдать ее ID - ТОЛЬКО РАСПОЗНАВАНИЕ, А НЕ СОЗДАНИЕ НОВОЙ
	phrID := GetExistsPraseID(str0)
	if phrID > 0 {
		CurrentPhrasesIDarr = append(CurrentPhrasesIDarr, DetectedUnicumPhraseID)
		// если вся фраза известа, то ничего больше не делать
		//out_str = append(out_str, out)

		CurretWordsIDarr = GetWordsIDarrFromPraseNodeID(phrID)

		return phrID, str0
		//noAddToList=true
	}
	// из-за непонятного глюка (внезапная очистка) с непосредственным добавлением ввожу промежуточный буфер:
	var tArr []string // новые фразы и слова, которые нужно добавить в tempArr (words_temp_arr.txt)

	if len(str) < 2 {
		return phrID, str
	}
	var out_str string // массив распознанных слов во фразе - для вывода на пульт
	// разделить слова
	// разбить по разделителям
	sArr := strings.Split(str, "|")
	var wordCount = 0     // число только слов, с пробелами
	var onlyWordCount = 0 // число только слов, без пробелов
	// фраза без разделителей и отдельные слова (более 1 символа) - в шаблон:
	var out = ""
	//var isBegin = true
	// подсчет числа слов
	cnt := len(sArr)
	for i := 0; i < cnt; i++ {
		if len(sArr[i]) == 0 {
			continue
		}
		if sArr[i] != " " {
			onlyWordCount++
		}
		wordCount++
		// во фразе оставляем и слова и символы, т.е. восстанавливаем фразу как она была
		out += sArr[i]
	}

	if onlyWordCount > 1 && onlyWordCount < 6 { // записать всю фразу, но не длинее 6 слов
		tArr = append(tArr, out)
	}
	///////////////   конец для всей фразы

	// Отдельные слова.
	zp := regexp.MustCompile(`[ ,:.+\(\)]`)
	sPath := zp.Split(out, -1) // Выделить слова по пробелам и знакам препинания
	//sArr = strings.Split(sPath, " ")
	for i := 0; i < len(sPath); i++ {
		wID, ok := WordIdFormWord[sPath[i]]
		if ok { // слово уже известно
			out_str += " " + sPath[i]
			CurretWordsIDarr = append(CurretWordsIDarr, wID)
			continue
		}
		r := []rune(sPath[i])
		if len(r) < 2 {
			continue
		}
		tArr = append(tArr, sPath[i])
		CurretWordsIDarr = append(CurretWordsIDarr, -1)
		out_str += " <span style='color:red' title='НЕ РАСПОЗНАННО'>" + sPath[i] + "</span> "
	}

	// добавляем новое
	for i := 0; i < len(tArr); i++ {
		nc := 0
		lib.MapCheck(MapGwardWordTempArr)
		_, ok := tempArr[tArr[i]] // var tempArr=make(map[string][]int)
		if ok {
			lib.MapCheck(MapGwardWordTempArr)
			nc = tempArr[tArr[i]][0]
		}
		lib.MapCheckWrite(MapGwardWordTempArr)
		tempArr[tArr[i]] = []int{nc + 1, 0}
		lib.MapFree(MapGwardWordTempArr)
	}

	SaveTempArr()

	return phrID, out_str
}

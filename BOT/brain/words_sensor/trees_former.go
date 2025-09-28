/*  Формирователь Дерева слов и Дерева фраз при наборе текстов
из редактора http://go/pages/words.php
и при общении с Beast с Пульта http://go/pult.php
*/

package word_sensor

import (
	"BOT/lib"
	"sort"
	"strconv"
	"strings"
)

/* переносим в дерево слов достаточно повторяющиеся из tempArr
 limitWord - число повторений отдельных слов, с которых начинается пеернос.
limitFraze - число повторений отдельных фраз (с несколькими словами), с которых начинается пеернос.
Для автозаливки - (4,6),
для сообщения с Пульта - (2,4)
Фомат words_temp_arr.txt: Число повторений|#|Фраза|#|число сохранений (старость)
*/
//var curSizeWordsTempArrTxt=0
func UpdateWordTreeFromTempArr(limitWord int, limitFraze int) {
	// return

	loadTempArr() // вский раз получать обновленный массив накоплений слов

	// Удалять использованные строки накопительного массива
	var newTempFileStr = "" // это будет записано после процесса в /memory_reflex/words_temp_arr.txt
	var str []string        // для выбранных повторяющихся слов - чтобы потом отсортирвать
	// чтобы фразы могли использовать слова (минимизация размера дерева)
	// слово или фраза?
	var existsWords = false
	var existsPhrase = false

	// сначала прходим только слова, потом - только фразы,
	lib.MapCheckBlock(MapGwardWordTempArr)
	for k, v := range tempArr {
		sps := strings.Split(k, " ")
		if len(sps) == 1 { // это слово
			if v[0] >= limitWord {
				str = append(str, k)

			} else {
				if v[1] > 1000 { // старое забывать, - очистка мусора и давно не воспринимаемого
					continue
				}
				//оставляем строку со словом, но добавляем ее старость: strconv.Itoa(v[1]+1)
				newTempFileStr += strconv.Itoa(v[0]) + "|#|" + k + "|#|" + strconv.Itoa(v[1]+1) + "\r\n"
			}
		}
	}
	lib.MapFree(MapGwardWordTempArr)

	/* str - список слов, повторившихся в memory_reflex/words_temp_arr.txt более limitWord раз
	Распознать их в сенсоре, дополнив базу слов.
	*/
	//oldNCh:=NoCheckWordCount
	//NoCheckWordCount=true
	sort.Strings(str) // по алфавиту, чтобы максимально облегчить последовательное разделение слов
	for i := 0; i < len(str); i++ {
		cur := str[i]
		SetNewWordTreeNode(cur, false)
		SetNewPhraseTreeNode(cur)
		existsWords = true
	}
	//NoCheckWordCount=oldNCh
	if existsWords {
		SaveWordTree()
		//	SavePhraseTree()// для формирования WordIdFormWord[]
		existsPhrase = true
	}

	// проход для фраз
	str = nil
	lib.MapCheckBlock(MapGwardWordTempArr)
	for k, v := range tempArr {
		sps := strings.Split(k, " ")
		if len(sps) > 1 { // это фраза
			if v[0] >= limitFraze {
				str = append(str, k)
			} else {
				if v[1] > 1000 { // старое забывать, - очистка мусора и давно не воспринимаемого
					continue
				}
				//оставляем строку с фразой, но добавляем ее старость: strconv.Itoa(v[1]+1)
				newTempFileStr += strconv.Itoa(v[0]) + "|#|" + k + "|#|" + strconv.Itoa(v[1]+1) + "\r\n"
			}
		}
	}
	lib.MapFree(MapGwardWordTempArr)

	// на распознавание фраз
	sort.Strings(str) // по алфавиту, чтобы максимально облегчить последовательное разделение слов
	for i := 0; i < len(str); i++ {
		cur := str[i]
		SetNewPhraseTreeNode(cur)
		existsPhrase = true
		// SaveWordTree() // для пошагового контроля
		// if(i>1){}
	}
	if existsPhrase {
		SavePhraseTree()
	}

	// Удаление использованных строк накопительного массива: запись только незатронутых
	lib.WriteFileContentExactly(lib.GetMainPathExeFile()+"/memory_reflex/words_temp_arr.txt", newTempFileStr)
}

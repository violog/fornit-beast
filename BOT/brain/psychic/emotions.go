/* Эмоции.
Распознавание активности текущих Базовых контекстов в виде структур
- осмысленной значимости сочетаний активных базовых контекстов.
Произвольно возможна активация имеющегося образа, создание образа новых сочетаний.
*/

package psychic

import (
	"BOT/brain/gomeostas"
	"BOT/lib"
	"strconv"
	"strings"
)

///////////////////////////////////

// ////////////////////////////
func emotionsInit() {
	loadEmotionArr()
}

////////////////////////////////

/*
Образ сочетания базовых контекстов
*/
type Emotion struct {
	ID        int   // идентификатор данного сочетания контекстов
	BaseIDarr []int // сочетание базовых контекстов
}

////////////////////////////////

var CurrentCommonBadNormalWell = 0 // 1 - Похо, 2 - Норма, 3 - Хорошо
var oldCommonBadNormalWell = 0

// последняя определенная эмоция
var CurrentEmotionReception *Emotion
var oldCurrentEmotionReception *Emotion

var EmotionFromIdArr = make(map[int]*Emotion)

/*  создать новую эмоцию, если такой еще нет
 */
var lastEmotionID = 0

func createNewBaseStyle(id int, BaseIDarr []int, CheckUnicum bool) (int, *Emotion) {
	if CheckUnicum {
		oldID, oldVal := checkUnicumEmotion(BaseIDarr)
		if oldVal != nil {
			//CurrentEmotionReception=oldVal
			return oldID, oldVal
		}
	}

	if id == 0 {
		lastEmotionID++
		id = lastEmotionID
	} else {
		//		newW.ID=id
		if lastEmotionID < id {
			lastEmotionID = id
		}
	}

	var node Emotion
	node.ID = id
	node.BaseIDarr = BaseIDarr

	EmotionFromIdArr[id] = &node

	if doWritingFile {
		SaveEmotionArr()
	}

	//CurrentEmotionReception=&node  - только при активации дерева актоматизмов!!!

	return id, &node
}
func checkUnicumEmotion(bArr []int) (int, *Emotion) {
	for id, v := range EmotionFromIdArr {
		if v == nil {
			continue
		}
		if lib.EqualArrs(bArr, v.BaseIDarr) {
			return id, v
		}
	}

	return 0, nil
}

////////////////////////////////////////

// ///////////////  сохранить образы сочетаний базовых стилей
func SaveEmotionArr() {
	var out = ""
	for k, v := range EmotionFromIdArr {
		if v == nil {
			continue
		}
		out += strconv.Itoa(k) + "|"
		for i := 0; i < len(v.BaseIDarr); i++ {
			out += strconv.Itoa(v.BaseIDarr[i]) + ","
		}
		out += "\r\n"
	}
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/emotion_images.txt", out)
}

// ////////////////  загрузить образы сочетаний базовых стилей
func loadEmotionArr() {
	EmotionFromIdArr = make(map[int]*Emotion)
	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/emotion_images.txt")
	cunt := len(strArr)
	for n := 0; n < cunt; n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		p := strings.Split(strArr[n], "|")
		id, _ := strconv.Atoi(p[0])
		s := strings.Split(p[1], ",")
		var BaseIDarr []int
		for i := 0; i < len(s); i++ {
			if len(s[i]) == 0 {
				continue
			}
			bc, _ := strconv.Atoi(s[i])
			BaseIDarr = append(BaseIDarr, bc)
		}
		var saveDoWritingFile = doWritingFile
		doWritingFile = false
		createNewBaseStyle(id, BaseIDarr, false)
		doWritingFile = saveDoWritingFile
	}
	return
}

/////////////////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////////////////
// Описать словами текущую эмоцию
func getEmotonsComponentStr(em *Emotion, strID bool) string {
	var out = " "

	if strID {
		out = "(ID=" + strconv.Itoa(em.ID) + ") "
	}
	if em == nil {
		return "НЕТ"
	}
	for i := 0; i < len(em.BaseIDarr); i++ {
		if i > 0 {
			out += ", "
		}
		out += gomeostas.GetBaseContextCondFromID(em.BaseIDarr[i])
	}

	return out
}

//////////////////////////////////////////

// последняя испытанная эмоция в виде строки
func GetCurrentEmotionReception() string {
	if CurrentEmotionReception == nil {
		return "Эмоция еще не определена."
	}
	return getEmotonsComponentStr(CurrentEmotionReception, true)
}

///////////////////////////////////////////

// есть ли данный компонент в текущей эмоции
func existsBaseContext(baseContID int) bool {
	if CurrentEmotionReception == nil {
		return false
	}
	for i := 0; i < len(CurrentEmotionReception.BaseIDarr); i++ {
		if CurrentEmotionReception.BaseIDarr[i] == baseContID {
			return true
		}
	}
	return false
}

//////////////////////////////////////////////

// для пульта
/* Уже есть в automatism_tree_pult_show.go func GetStrnameFromBaseImageID(id int)(string)
func GetEmotionContexts(emID int)string{
	em,ok:=EmotionFromIdArr[emID]
if !ok{
	return "Нет эмоции с ID="+strconv.Itoa(emID)
}
	out:=getEmotonsComponentStr(em, false)

	return out
}*/
//////////////////////////////////////////

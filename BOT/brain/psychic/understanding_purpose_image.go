/* Это - осознаваемая, но не осмысливаемая цель (не поставленная осмысленно), а мотивирующая потребность, дающая направленность мышлению.
Формируется на основе актуальной Темы размышления, возникающей при некоторых событиях.
Но может быть выбрана и произвольно.

Для осмысливаемых целей - Доминанта и там надо будет более детально прописывать желаемое.
*/

package psychic

import (
	"BOT/lib"
	"strconv"
	"strings"
)

///////////////////////////////

/*
	образ цели действия.

Осознаваемая Цель - проблема Доминанты.
Отражает состояние 3-х уровней дерева ситуации:
1-MoodID, 2-EmotionID, 3-situationImageID (здесь определена текущая цель problemTreeInfo.purposeID)
Если есть доминанта, то в ней такая цель зафиксирована в виде problemTreeID (из текущей problemTreeInfo.purposeID)
Так что улучшение ситуации без доминанты подразумевает решение текущей проблемы,
а с доминантой - с той проблемой, которая бала зафиксирована в виде problemTreeID.

Дополнительно - целями мошут быть:
добиться возможности совершить данное действие и
добиться достижения эффекта моторного Правила с rules.ID

mentalInfoStruct.mentalPurposeID - поставлена текущая цель

Цель считается достигнутой (эффект совершенных ментальных действий) если
заданные (НЕНУЛЕВЫЕ) параметры достигнуты.
При нескольких заданных параметрах цель может быть частично достигнута.
Эффект прикидывается в зависимости от полноты достижения цели:
при полном недостижении эффект ==0 при полном достижении +10
*/
type PurposeImage struct {
	ID int
	// цель - добиться target
	target int // 1- повторения, 2 - улучшения
	//добиться target значения PsyBaseMood -1 Плохое настроение, 0 Нормальное, 1 - хорошее настроение
	moodeID int // этот параметр будет всегда считаться заданным в getMentalPurposeEffect().

	/* В случае Цели - улучшить эмоцию (target==2) - имеется в виду, что
	сумма весов позитивных эмоциональных контекстов превышает сумму вемов негатиных:
	func isEmotonBetter(oldID int, curID int)
	*/
	emotonID int // добиться target такой эмоции

	/* В случае Цели - улучшить ситуацию - func isSituationBetter(oldID int, curID int)

	 */
	situationID int // добиться PurposeImage.target данной ситуации SituationImage

}

// var PurposeImageFromID=make(map[int]*PurposeImage)
var PurposeImageFromID []*PurposeImage // сам массив
// var AFromID = make([]*aNode, len(strArr))//задать сразу имеющиеся в файле число при загрузке
// запись члена
func WritePurposeImageFromID(index int, value *PurposeImage) {
	addPurposeImageFromID(index)
	PurposeImageFromID[index] = value
}
func addPurposeImageFromID(index int) {
	if index >= len(PurposeImageFromID) {
		newSlice := make([]*PurposeImage, index+1)
		copy(newSlice, PurposeImageFromID)
		PurposeImageFromID = newSlice
	}
}

// считывание члена
func ReadePurposeImageFromID(index int) (*PurposeImage, bool) {
	if index >= len(PurposeImageFromID) || PurposeImageFromID[index] == nil {
		return nil, false
	}
	return PurposeImageFromID[index], true
}

///////////////////////////////////////////////////////////////////

// создать новый образ желаемой цели, если такого еще нет
var lastPurposeImagePurposeID = 0

func createPurposeImageID(id int, target int, moodeID int, emotonID int, situationID int, CheckUnicum bool) (int, *PurposeImage) {
	if CheckUnicum {
		oldID, oldVal := checkUnicumPurposeImage(target, moodeID, emotonID, situationID)
		if oldVal != nil {
			return oldID, oldVal
		}
	}

	if id == 0 {
		lastPurposeImagePurposeID++
		id = lastPurposeImagePurposeID
	} else {
		//		newW.ID=id
		if lastPurposeImagePurposeID < id {
			lastPurposeImagePurposeID = id
		}
	}

	var node PurposeImage
	node.ID = id
	node.target = target
	node.moodeID = moodeID
	node.emotonID = emotonID
	node.situationID = situationID

	//	PurposeImageFromID[id]=&node
	WritePurposeImageFromID(id, &node)

	if doWritingFile {
		SavePurposeImageFromIdArr()
	}

	return id, &node
}
func checkUnicumPurposeImage(target int, moodeID int, emotonID int, situationID int) (int, *PurposeImage) {
	for id, v := range PurposeImageFromID {
		if v == nil {
			continue
		}
		if target != v.target || moodeID != v.moodeID || emotonID != v.emotonID || situationID != v.situationID {
			continue
		}
		return id, v
	}
	return 0, nil
}

/////////////////////////////////////////

// ///////////////////////////////////////
// сохранить образы
func SavePurposeImageFromIdArr() {
	var out = ""
	for k, v := range PurposeImageFromID {
		if v == nil {
			continue
		}
		out += strconv.Itoa(k) + "|"
		out += strconv.Itoa(v.target) + "|"
		out += strconv.Itoa(v.moodeID) + "|"
		out += strconv.Itoa(v.emotonID) + "|"
		out += strconv.Itoa(v.situationID)
		out += "\r\n"
	}
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/purpose_images.txt", out)

}

// //////////////////  загрузить образы
func loadPurposeImageFromIdArr() {

	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/purpose_images.txt")
	cunt := len(strArr)
	//PurposeImageFromID=make(map[int]*PurposeImage)
	PurposeImageFromID = make([]*PurposeImage, cunt) //задать сразу имеющиеся в файле число
	for n := 0; n < cunt; n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		p := strings.Split(strArr[n], "|")
		id, _ := strconv.Atoi(p[0])
		target, _ := strconv.Atoi(p[1])
		moodeID, _ := strconv.Atoi(p[2])
		emotonID, _ := strconv.Atoi(p[3])
		situationID, _ := strconv.Atoi(p[4])

		var saveDoWritingFile = doWritingFile
		doWritingFile = false
		createPurposeImageID(id, target, moodeID, emotonID, situationID, false)
		doWritingFile = saveDoWritingFile
	}
	return

}

//////////////////////////////

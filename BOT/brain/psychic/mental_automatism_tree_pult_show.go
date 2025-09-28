/*  Выдать на Пульт дерево ментальных автоматизмов
id|Usefulness|ActionsImageID|motAutmtzmID
*/

package psychic

import (
	"strconv"
)

////////////////////////////////////////////

/*запрет показа карты на пульте (func GetAutomatizmTreeForPult()) при обновлении
против паники типа "одновременная запись и считывание карты"
Использовать для всех операций записи узлов дерева
*/

// образ дерева автоматизмов для вывода
var automatizmMentalTreeModel = ""

// ///////////////////////////////////////////////
func GetMentalAutomatizmTreeForPult(limit int) string {
	// против паники типа "одновременная запись и считывание карты"
	if notAllowScanInTreeThisTime {
		return "!Временно запрещена работа func GetAutomatizmTreeForPult() т.к. идет параллельная обработка."
	}
	if len(UnderstandingTree.Children) == 0 { // еще нет никаких веток
		return "Еще нет Дерева ситуации"
	}

	//посмотреть число имеющихся узлов дерева
	//strArr,_:=lib.ReadLines(lib.GetMainPathExeFile()+"/memory_psy/automatizm_tree.txt")
	automatizmMentalTreeModel = "" // иначе дублирует блоки дерева при каждом выводе страницы
	scanMentalAutomatizmNodes(-1, &UnderstandingTree)

	if len(automatizmMentalTreeModel) < 10 {
		return "Еще нет информационных веток дерева"
	}

	return automatizmMentalTreeModel

}

// ////////////////////
// ID|ParentNode|Mood|EmotionID|SituationID|PurposeID
func scanMentalAutomatizmNodes(level int, node *UnderstandingNode) {

	if node.ID > 0 {
		automatizmMentalTreeModel += "<span style='color:#666666;'>" + strconv.Itoa(node.ID) + ":</span> "

		automatizmMentalTreeModel += setOutShift(level)

		switch level {
		case 0: // Mood
			automatizmMentalTreeModel += getMoodStr(node)
		case 1: // EmotionID
			automatizmMentalTreeModel += getEmotionStr(node)
		case 2: // SituationID
			automatizmMentalTreeModel += getSituationStr(node)
			//case 3: // PurposeID
			//	automatizmMentalTreeModel +=getPurposeStr(node)
		}
		automatizmMentalTreeModel += "<br>\n"
	}
	level++
	for n := 0; n < len(node.Children); n++ {
		scanMentalAutomatizmNodes(level, &node.Children[n])
	}
}

// отступ
func setOutShift(level int) string {
	var sh = ""
	for n := 0; n < level; n++ {
		sh += "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"
	}
	return sh
}

func GetMentalSituationsForNodeInfo(sID int) string {
	//	st:=SituationImageFromIdArr[sID]
	st, ok := ReadeSituationImageFromIdArr(sID)
	if !ok {
		return "Нет образа ситуации с ID=" + strconv.Itoa(sID)
	}

	switch st.SituationType {
	case 1:
		return "Было ответное действие"
	case 2:
		return "Был запуск автоматизма ветки"
	case 3:
		return "Ничего не делали, но нужно осмысление"
	case 4:
		return "Все спокойно, можно экспериментировать"
	case 5:
		return "Оператор не прореагировал на действия в течение периода ожидания"

	case 11:
		return "Оператор выбрал настроение Хорошее"
	case 12:
		return "Оператор выбрал настроение Плохое"
	case 13:
		return "Оператор выбрал настроение Игровое"
	case 14:
		return "Оператор выбрал настроение Учитель"
	case 15:
		return "Оператор выбрал настроение Агрессивное"
	case 16:
		return "Оператор выбрал настроение Защитное"
	case 17:
		return "Оператор выбрал настроение Протест"

	//case 20: return "Оператор нажал кнопку "
	case 21:
		return "Оператор нажал кнопку Непонятно"
	case 22:
		return "Оператор нажал кнопку Понятно"
	case 23:
		return "Оператор нажал кнопку Наказать"
	case 24:
		return "Оператор нажал кнопку Поощрить"
	case 25:
		return "Оператор нажал кнопку Накормить"
	case 26:
		return "Оператор нажал кнопку Успокоить"
	case 27:
		return "Оператор нажал кнопку Предложить поиграть"
	case 28:
		return "Оператор нажал кнопку Предложить поучить"
	case 29:
		return "Оператор нажал кнопку Игнорировать"
	case 30:
		return "Оператор нажал кнопку Сделать больно"
	case 31:
		return "Оператор нажал кнопку Сделать приятно"
	case 32:
		return "Оператор нажал кнопку Заплакать"
	case 33:
		return "Оператор нажал кнопку Засмеяться"
	case 34:
		return "Оператор нажал кнопку Обрадоваться"
	case 35:
		return "Оператор нажал кнопку Испугаться"
	case 36:
		return "Оператор нажал кнопку Простить"
	case 37:
		return "Оператор нажал кнопку Вылечить"
	}
	return ""
}

func GetMentalPurposeForNodeInfo(pID int) string {
	out := "Цель - "
	outMode := ""
	outEmo := ""
	outSit := ""
	//	pp:=PurposeImageFromID[pID]
	pp, ok := ReadePurposeImageFromID(pID)
	if !ok {
		return "Нет образа цели для pID=" + strconv.Itoa(pID) + " - в func GetMentalPurposeForNodeInfo"
	}

	switch {
	case pp.moodeID < 0:
		outMode = "плохого настроения<br>"
	case pp.moodeID == 0:
		outMode = "нормального настроения<br>"
	case pp.moodeID > 0:
		outMode = "хорошего настроения<br>"
	}
	if pp.emotonID > 0 {
		em := EmotionFromIdArr[pp.emotonID]
		outEmo = "ID эмоции =" + strconv.Itoa(pp.emotonID) + ": " + getEmotonsComponentStr(em, false) + "<br>"
	}
	if pp.situationID > 0 {
		outSit = "ID ситуации =" + strconv.Itoa(pp.situationID) + ": " + getSituationString(pp.situationID) + "<br>"
	}

	if pp.target == 1 {
		out += "повторение:<br>"
	} else {
		out += "улучшение:<br>"
	}

	out += outMode
	if outEmo != "" {
		out += outEmo
	}
	if outSit != "" {
		out += outSit
	}

	return out
}

func GetMentalThemeForNodeInfo(tID int) string {
	out := ""

	//	tImg:=ThemeImageFromID[tID]
	tImg, ok := ReadeThemeImageFromID(tID)
	if ok {
		out += "Тип темы: <b>" + GetThemeImageName(tImg.Type) + "</b><br>"
		out += "Вес значимости темы: <b>" + strconv.Itoa(tImg.Weight) + "</b><br>"
		diff := LifeTime - tImg.PulsCount
		out += "Время жизни темы: <b>" + strconv.Itoa(diff) + " секунд</b><br>"
		return out
	}
	return ""
}

// Mood
func getMoodStr(node *UnderstandingNode) string {
	return getMoodString(node.Mood)
}
func getMoodString(id int) string {
	moodS := ""
	switch id {
	case -1:
		moodS = "Плохое"
	case 0:
		moodS = "Нормальное"
	case 1:
		moodS = "Хорошее"
	}
	out := "Настроение: <b> " + moodS + "</b>"
	return out
}

// EmotionID
func getEmotionStr(node *UnderstandingNode) string {
	return getEmotionString(node.EmotionID)
}
func getEmotionString(id int) string {
	em := EmotionFromIdArr[id]
	out := "Эмоция (" + strconv.Itoa(em.ID) + "):<b> " + getEmotonsComponentStr(em, false) + "</b>"
	return out
}

// SituationID
func getSituationStr(node *UnderstandingNode) string {
	return getSituationString(node.SituationID)
}
func getSituationString(id int) string {

	//pp := SituationImageFromIdArr[id]
	out := ""
	_, ok := ReadeSituationImageFromIdArr(id)
	if !ok {
		out += "<span style='color:red'>Нет образа Ситуации с ID=" + strconv.Itoa(id) + "</span>"
	} else {
		out += "Ситуация: <b> <span style='cursor:pointer;color:blue' onClick='get_situation(" + strconv.Itoa(id) + ")'>" + strconv.Itoa(id) + "</span>" + "</b>"
	}
	return out
}
func getSituationDetaileString(id int) string {
	//	pp := SituationImageFromIdArr[id]
	out := ""
	_, ok := ReadeSituationImageFromIdArr(id)
	if !ok {
		out += "<span style='color:red'>Нет образа Ситуации с ID=" + strconv.Itoa(id) + "</span>"
	} else {
		out += "Ситуация: <b> " + GetMentalSituationsForNodeInfo(id) + "</b>"
	}
	return out
}

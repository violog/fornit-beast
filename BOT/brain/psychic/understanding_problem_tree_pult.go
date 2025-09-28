/* Вывод дерева проблем на пульт
 */

package psychic

import (
	"strconv"
)

// образ дерева автоматизмов для вывода
var problemTreeStr = ""

func GetMentalPriblemTreeForPult(limit int) string {
	problemTreeStr = ""
	// против паники типа "одновременная запись и считывание карты"
	if notAllowScanInTreeThisTime {
		return "!Временно запрещена работа func GetAutomatizmTreeForPult() т.к. идет параллельная обработка."
	}
	if len(ProblemTree.Children) == 0 { // еще нет никаких веток
		return "Еще нет Дерева проблем"
	}

	scanMentalPriblemNodes(-1, &ProblemTree)

	if len(problemTreeStr) < 10 {
		return "Еще нет информационных веток дерева проблем"
	}
	return problemTreeStr
}

// ID|ParentNode|autTreeID|situationTreeID|themeID|purposeID
func scanMentalPriblemNodes(level int, node *ProblemTreeNode) {
	if node.ID > 0 {
		out := "<span style='color:#666666;'>" + strconv.Itoa(node.ID) + ":</span> "

		out += setOutShift(level)

		switch level {
		case 0: // ID Дерева автоматизмов
			out += getAutomTreeNodeString(node.autTreeID)
		case 1: // ID ситуации
			if node.situationTreeID == 0 {
				return
			}
			out += getSituationString(node.situationTreeID)
		case 2: // ID Темы
			if node.themeID == 0 {
				return
			}
			out += getThemeStr(node.themeID)
		case 3: // ID Цели
			if node.purposeID == 0 {
				return
			}
			out += getPurposeString(node.purposeID)
			/* // привязанный мент.автоматизм
			ma:=ExistsStaffMentAutomatizmForThisNodeID(node.ID)
			if ma!=nil {
				out += " Мент.автоматизм c ID=<b> <span style='cursor:pointer;color:blue' onClick='get_ment_automatizm(" + strconv.Itoa(ma.ID) + ")'></b>" + strconv.Itoa(ma.ID) + "</span>" + "\n"
			}else{out += " Мент.авт-м не привязан."}
			*/
		}
		out += "<br>\n"
		problemTreeStr += out
	}
	level++
	for n := 0; n < len(node.Children); n++ {

		scanMentalPriblemNodes(level, &node.Children[n])
	}
}

func getAutomTreeNodeString(id int) string {

	//	atmzm:=AutomatizmTreeFromID[id]
	atmzm, ok := ReadeAutomatizmTreeFromID(id)
	if !ok {
		return ""
	}
	out := ""
	if atmzm == nil {
		out += "<span style='color:red'>Нет узла дерева автоматизма с ID=" + strconv.Itoa(id) + "</span>"
	} else {
		out += "ID дерева автоматизмов: <b> <span style='cursor:pointer;color:blue' onClick='show_node_automatizms(" + strconv.Itoa(id) + ")'>" + strconv.Itoa(id) + "</span>" + "</b>"
	}
	return out

}

func getThemeStr(id int) string {

	//	theme:=ThemeImageFromID[id]
	out := ""
	theme, ok := ReadeThemeImageFromID(id)
	if !ok {
		out += "<span style='color:red'>Нет Темы с ID=" + strconv.Itoa(id) + "</span>"
	} else {
		out += "Тема: <b> <span style='cursor:pointer;color:blue' onClick='get_theme(" + strconv.Itoa(theme.ID) + ")'>" + strconv.Itoa(id) + "</span>" + "</b>"
	}
	return out
}

// PurposeID
func getPurposeString(id int) string {
	//	pp:=PurposeImageFromID[id]
	out := ""
	_, ok := ReadePurposeImageFromID(id)
	if !ok {
		out += "<span style='color:red'>Нет образа цели с func getPurposeString id=" + strconv.Itoa(id) + "</span>"
	} else {
		out += "Цель: <b> <span style='cursor:pointer;color:blue' onClick='get_purpose(" + strconv.Itoa(id) + ")'>" + strconv.Itoa(id) + "</span>" + "</b>"
	}
	return out
}

func getPurposeDetaileString(id int) string {
	if id == 0 {
		return "<span style='color:red'>Нет образа цели.</span>"
	}
	//	pp:=PurposeImageFromID[id]
	out := ""
	_, ok := ReadePurposeImageFromID(id)
	if ok {
		out += "Цель: <b> " + GetMentalPurposeForNodeInfo(id) + "</b>"
	} else {
		out += "<span style='color:red'>Нет образа цели с ID=" + strconv.Itoa(id) + "</span>"
	}
	return out
}

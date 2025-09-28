/*
для пульта показ дерева ментальной эпиз.памяти
*/
package psychic

import "strconv"

// ////////////////////////////////////////////////
// 100 последних кадров, записанных в кадрах эпизодической памяти
// var idEpArr = make(map[int]param)// массивы params для данного ID кадра памяти
func GetMentallastRules() string {

	hStr := "<b>Историческая последовательность кадров:</b><br>"

	// последний индекс исторической памяти EpisodicMentalHistoryArr[]
	lastI := len(EpisodicMentalHistoryArr) - 1
	var node *EpisodicMentalTreeNode
	out := "<hr><br>До 100 последних Правил, записанных в кадрах эпизодической памяти." +
		"<table cellpadding=0 cellspacing=0 border=1 class='main_table'><tr>" +
		"<th class='table_header' title='Номер кадра эпизодической памяти - индекс ячейки исторической памяти'>№ кадра</th>" +
		"<th class='table_header' title='ID кадра эпизодической памяти'>ID</th>" +
		"<th class='table_header' title='ID дерева проблем'>NodePID</th>" +
		"<th class='table_header' title='Тема мышления'>ThemeID</th>" +
		"<th class='table_header' title='Текущая цель'>PurposeID</th>" +
		"<th class='table_header' title='последовательность ID инфо функций, которая применялось после активации NodePID'>InfoArr</th>" +
		"<th class='table_header' title='Эффет от цепочки инфо-функций'>Effect</th>" +
		"<th class='table_header' title='Уверенность'>Уверенность</th></tr>"
	rCount := 0
	style := "style='font-size:19px;font-weight:bold;cursor:pointer'"
	for i := lastI; i >= 0 || rCount > 100; i-- {
		curID := EpisodicMentalHistoryArr[i].ID
		if curID == -1 {
			hStr += frameMentStr(-1)
			continue
		}

		node0, ok := ReadeEpisodicMentalTreeNodeFromID(curID)
		if !ok {
			continue
		}
		node = node0

		infoStr := ""
		for n := 0; n < len(node.InfoArr); n++ {
			if n > 0 {
				infoStr += ", "
			}
			infoStr += "<span onClick='show_mental_func(" + strconv.Itoa(node.InfoArr[n]) + ")'>" + strconv.Itoa(node.InfoArr[n]) + "</span>"
		}

		pars := node.PARAMS
		if len(pars) == 0 {
			continue
		}

		bg := "#ffffff" // фон строки учительского правила - bg=="#FFEBE4"
		// Показываем  node.PARAMS
		effect := pars[0]

		out += "<tr style='background-color:" + bg + ";' title='Каддр памяти'>" +
			"<td >" + strconv.Itoa(i) + "</td>" +
			"<td >" + strconv.Itoa(curID) + "</td>" +
			"<td " + style + " onClick='show_problem_tree(" + strconv.Itoa(node.NodePID) + ")'>" + strconv.Itoa(node.NodePID) + "</td>" +
			"<td " + style + " onClick='show_theme_strings(1," + strconv.Itoa(node.ThemeID) + ")'>" + strconv.Itoa(node.ThemeID) + "</td>" +
			"<td " + style + " onClick='get_purpose(1," + strconv.Itoa(node.PurposeID) + ")'>" + strconv.Itoa(node.PurposeID) + "</td>" +
			"<td " + style + ">" + infoStr + "</td>" +
			"<td " + style + ">" + strconv.Itoa(effect) + "</td>" +
			"<td " + style + ">" + strconv.Itoa(pars[1]) + "</td></tr>"
		rCount++
		hStr += frameMentStr(curID)
	}
	if rCount == 0 {
		out += "<tr><td colspan=7>Не найдены кадры в эпизодической памяти.</td></tr>"
	}
	out += "</table>"

	out = hStr + out

	return out
}

func frameMentStr(id int) string {

	if id == -1 {
		return "<div class='frameEm' title='пустой кадр - конец темы'>&nbsp;</div>"
	}
	return "<div class='frameEP'>" + strconv.Itoa(id) + "</div>"
}

///////////////////////////////////////////////////////////

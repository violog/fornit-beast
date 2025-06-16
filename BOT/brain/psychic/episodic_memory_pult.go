/* Показ эпиз.памяти на Пульте

 */

package psychic

import (
	"BOT/lib"
	"sort"
	"strconv"
)

/*
	100 последних Правил, записанных в кадрах эпизодической памяти, включая "учительские"
	var idEpArr = make(map[int]param)// массивы params для данного ID кадра памяти

Если objID !=0 то выдать не 100, а все и только где есть объект objID
*/
func GetCur100lastRules(objID int) string {

	hStr := "<b>Историческая последовательность 100 последних кадров эпизодической памяти:</b><br>"
	if objID > 0 {
		hStr = "<b>Модель понимания: значимости кадров эпизодической памяти для объекта ID=" + strconv.Itoa(objID) + ":</b><br>"
	}

	// последний индекс исторической памяти EpisodicHistoryArr[]
	lastI := len(EpisodicHistoryArr) - 1
	var node *EpisodicTreeNode
	out := "<hr><br>параметры Правил, записанных в кадрах эпизодической памяти, включая учительские." +
		"<table cellpadding=0 cellspacing=0 border=1 class='main_table'><tr>" +
		"<th class='table_header' title='ID кадра эпизодической памяти'>ID</th>" +
		"<th class='table_header' title='1 - Похо, 2 - Норма, 3 - Хорошо'>BaseID</th>" +
		"<th class='table_header' title='ID эмоции'>EmotionID</th>" +
		"<th class='table_header' title='ID дерева проблем'>NodePID</th>" +
		"<th class='table_header' title='Стимул с Пульта'>Trigger</th>" +
		"<th class='table_header' title='Ответное действие'>Action</th>" +
		"<th class='table_header' title='Эффет от действия'>Effect</th>" +
		"<th class='table_header' title='Число подтверждений при записи такого кадра'>Уверенность</th>" +
		"<th class='table_header' title='Значимость образа в модели понимания'>Значимость</th></tr>"
	rCount := 0
	curID := 0
	var curIdArr []int
	style := "style='font-size:19px;font-weight:bold;cursor:pointer'"
	limit := 100
	if objID > 0 {
		limit = 100000000
	}
	for i := lastI; i >= 0 && rCount < limit; i-- {
		curID = EpisodicHistoryArr[i].ID
		if curID == -1 {
			hStr += frameStr(-1, 0)
			continue
		}

		node0, ok := ReadeEpisodicTreeNodeFromID(curID)
		if !ok {
			continue
		}
		if objID > 0 {
			if node0.Trigger != objID && node0.Action != objID {
				continue
			}
		}
		node = node0

		pars := node.PARAMS
		if len(pars) == 0 {
			continue
		}

		if !lib.ExistsValInArr(curIdArr, curID) { // это таблица-расшифровка параметров кадров памяти, тут не нужна история
			curIdArr = append(curIdArr, curID)
		}
		rCount++
		hStr += frameStr(curID, 0)
	}

	if rCount == 0 {
		out += "<tr><td colspan=7>Не найдены Правила в эпизодической памяти.</td></tr>"
	} else {
		sort.Ints(curIdArr)
		for i := 0; i < len(curIdArr); i++ {
			node0, ok := ReadeEpisodicTreeNodeFromID(curIdArr[i])
			if !ok {
				continue
			}
			node = node0

			pars := node.PARAMS
			bg := "#ffffff" // фон строки учительского правила - bg=="#FFEBE4"
			ttr := ""
			// Показываем  node.PARAMS
			effect := pars[0]
			if effect == 100 {
				bg = "#FFEBE4"
				ttr = "Учительское правило"
			}
			out += "<tr style='background-color:" + bg + ";' title='" + ttr + "'>" +
				"<td >" + strconv.Itoa(curIdArr[i]) + "</td>" +
				"<td " + style + " title='1 - Похо, 2 - Норма, 3 - Хорошо' >" + strconv.Itoa(node.BaseID) + "</td>" +
				"<td " + style + " onClick='show_emotion(" + strconv.Itoa(node.EmotionID) + ")'>" + strconv.Itoa(node.EmotionID) + "</td>" +
				"<td " + style + " onClick='show_problem_tree(" + strconv.Itoa(node.NodePID) + ")'>" + strconv.Itoa(node.NodePID) + "</td>" +
				"<td " + style + " onClick='show_object(1," + strconv.Itoa(node.Trigger) + ")'>" + strconv.Itoa(node.Trigger) + "</td>" +
				"<td " + style + " onClick='show_object(1," + strconv.Itoa(node.Action) + ")'>" + strconv.Itoa(node.Action) + "</td>" +
				"<td " + style + ">" + strconv.Itoa(effect) + "</td>" +
				"<td " + style + ">" + strconv.Itoa(pars[1]) + "</td>"
			if effect == 100 { // для учительского правила
				out += "<td style='font-size:19px;font-weight:bold;cursor:pointer;color:blue' onClick='get_undestand_model(" + strconv.Itoa(node.Action) + ")' title='Значимость node.Action'>" + strconv.Itoa(pars[2]) + "</td>"
			} else {
				out += "<td style='font-size:19px;font-weight:bold;cursor:pointer;color:blue' onClick='get_undestand_model(" + strconv.Itoa(node.Trigger) + ")'  title='Значимость node.Trigger'>" + strconv.Itoa(pars[2]) + "</td>"
			}
			out += "</tr>"
		}
	}
	out += "</table>"
	out = hStr + out

	return out
}

func frameStr(id int, kind int) string {

	if id == -1 {
		return "<div class='frameEm' title='пустой кадр - конец темы'>&nbsp;</div>"
	}
	if kind == 1 {
		return "<div class='frameEv' title='Кадр объекта Модели понимания'>" + strconv.Itoa(id) + "</div>"
	}
	return "<div class='frameEP'>" + strconv.Itoa(id) + "</div>"
}

///////////////////////////////////////////////////////////
/* модель понимания объекта с ID
Найти все вхождения объекта в эпиз.память и выдать таблицу как для func GetCur100lastRules()
но только для данного объекта
*/
func Get_undestand_model_from_object(objID int) string {

	return GetCur100lastRules(objID)
}

//////////////////////////////////////////////////////////////

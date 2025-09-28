/*   показ доминант ДЛЯ ПУЛЬТА

 */

package psychic

import (
	"BOT/lib"
	"strconv"
)

func GetDominantsListToPult() string {
	var out = ""

	for k, v := range DominantaProblem {
		if v == nil {
			continue
		}
		out += "ID=<b><span style='cursor:pointer;color:blue' onClick='show_dominant(" + strconv.Itoa(k) + ")'>" + strconv.Itoa(k) + "</span>|"
		out += " <span title='вес значимости проблемы'>Вес: " + strconv.Itoa(v.weight) + "</span>|"
		out += " <span title='степень решения проблемы'>Стадия решения: " + strconv.Itoa(v.isSuccess) + "</span>|"
		out += " <span title='ID rонечного узла активной ветки дерева проблем'>ID дерева проблем: <b><span style='cursor:pointer;color:blue' onClick='get_problem_tree(" + strconv.Itoa(v.problemTreeID) + ")'>" + strconv.Itoa(v.problemTreeID) + "</span>" + "</span>|"
		out += " <span title='объект размышления типа extremImportance'>Объект проблемы: <b><span style='cursor:pointer;color:blue' onClick='show_object(" + strconv.Itoa(v.objectID) + ")'>" + strconv.Itoa(v.objectID) + "</span>" + "</span>|"
		lib.MapCheck(MapGwardTryActionArr)
		aArr := tryActionArr[v.tryActionsKey]
		out += " Попыток решения <b>" + strconv.Itoa(len(aArr)) + "<br><br>\n"
	}
	return out
}

////////////////////////////////////////////////////////////

func DominantaInfoStr(id int) string {
	var out = ""
	//	dm:=DominantaProblem[id]
	dm, ok := ReadeDominantaProblem(id)
	if !ok {
		return "Нет доминанты с ID=" + strconv.Itoa(id)
	}

	out += "Доминанта ID=" + strconv.Itoa(id) + ":<br>"

	out += "вес значимости проблемы: " + strconv.Itoa(dm.weight) + "<br>"
	out += "степень решения проблемы: " + strconv.Itoa(dm.isSuccess) + "<br>"

	out += "Конечный узел активной ветки дерева проблем<b><span style='cursor:pointer;color:blue' onClick='get_problem_tree(" + strconv.Itoa(dm.problemTreeID) + ")'>" + strconv.Itoa(dm.problemTreeID) + "</span>" + "<br>\n"
	out += "объект размышления типа extremImportance <b><span style='cursor:pointer;color:blue' onClick='show_object(" + strconv.Itoa(dm.objectID) + ")'>" + strconv.Itoa(dm.objectID) + "</span>" + "<br>\n"
	lib.MapCheck(MapGwardTryActionArr)
	aArr := tryActionArr[dm.tryActionsKey]
	out += "число попыток решения <b>" + strconv.Itoa(len(aArr)) + " <br>\n"

	return out
}

/////////////////////////////////////////

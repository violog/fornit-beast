/* Вывод моделей понимания на Пульт

 */

package psychic

import (
	"strconv"
)

/*
func GetMentalUndestandingModelsForPult()string{
var out=""
	getAllExtrОbjIDarr()
if UnderstandingModelArr==nil{
	return "Нет Моделей понимания в эпиз.памяти."
}

	// этот вывод далается при пульсации, и чтобы не прыгало случайным образом на пульте нужно делать сортировку
	keys := make([]int, 0, len(UnderstandingModelArr))
	for id,v:= range UnderstandingModelArr {
		if v==nil{continue}
		keys = append(keys, id)
	}

	sort.Ints(keys)
	for _, id := range keys{
		arr,ok:=ReadeUnderstandingModelArr(id)
		if !ok {
			continue
		}
		out += "<b>Образ ID=<b> <span style='cursor:pointer;color:blue' onClick='get_ment_model_index(" + strconv.Itoa(id) + ")'>" + strconv.Itoa(id) + "</span>" + ":</b> "
		out += "<b>Кадры: <b>"
		for i:=0; i < len(arr); i++ {
			if i > 0 {
				out += ", "
			}
			out += "<span style='cursor:pointer;color:blue' onClick='get_epiz_memory_info(" + strconv.Itoa(arr[i]) + ")'>" + strconv.Itoa(arr[i]) + "</span>" + "</b>"
		}
		//out += ","
		out += "<br>\r\n"
	}
	return out
}
*/

func GetModelExtremImportanceInfo(id int) string {
	/*curImportanceObjectArr не сохраняется, а возникает только в данной сессиии
	Сохраняется extremImportance.extremObjID в importance.txt
	Поэтому id - это importance
	*/
	var out = ""
	//etremObj:= getExtremObjFromID(id)
	//	obj,ok:=importanceFromID[id]
	obj := getExtremObjFromID(id)
	if obj == nil {
		out += "ID объекта типа importance " + strconv.Itoa(id) + " не существует.<br>"
	} else {
		out += "ID объекта типа importance " + strconv.Itoa(id) + ":<br>"
		out += GetImportanceObjectInfo(obj.objID)
	}

	return out
}

// узел дерева проблем
func GetProblemTreeForNodeInfo(id int) string {
	var out = ""
	//	node:= ProblemTreeNodeFromID[id]
	node, ok := ReadeProblemTreeNodeFromID(id)
	if !ok {
		out += "Узел дерева проблем ID=" + strconv.Itoa(id) + " не существует.<br>"
	} else {
		out += "Узел дерева проблем: ID=" + strconv.Itoa(id) + ":<br>"

		out += "Конечный узел активной ветки дерева моторных автоматизмов: <span style='cursor:pointer;color:blue' onClick='show_atmzm_tree(" + strconv.Itoa(node.autTreeID) + ")'>" + strconv.Itoa(node.autTreeID) + "</span>" + "<br>\n"
		out += "Конечный узел активной ветки дерева ситуации: <span style='cursor:pointer;color:blue' onClick='show_unde_tree(" + strconv.Itoa(node.situationTreeID) + ")'>" + strconv.Itoa(node.situationTreeID) + "</span>" + "<br>\n"
		out += "Тема: <span style='cursor:pointer;color:blue' onClick='get_theme(" + strconv.Itoa(node.themeID) + ")'>" + strconv.Itoa(node.themeID) + "</span>" + "<br>\n"
		out += "Цель: <span style='cursor:pointer;color:blue' onClick='get_purpose(" + strconv.Itoa(node.purposeID) + ")'>" + strconv.Itoa(node.purposeID) + "</span>" + "<br>\n"
	}

	return out
}

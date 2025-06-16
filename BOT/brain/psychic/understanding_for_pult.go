/* показывать состояние психики на Пульте

 */

package psychic

import (
	"BOT/lib"
	"sort"
	"strconv"
)

/////////////////////////////////////////////

// выдать текущую инфу для http://go/pages/mental_cicles.php
/*
func GetCicklesToPult()(string){


	style:="style='font-size:19px;font-weight:bold;cursor:pointer'"
	//	Текущая ситуация:
	out:="<b>Текущая ситуация:</b><br>"
		out+="<table cellpadding=0 cellspacing=0 border=1 class='main_table'>"
		out+="<tr>"
		out+="<th class='table_header'>ID дерева понимания ситуации</th>"
		out+="<th class='table_header'>ID дерева автоматизмов</th>"
		out+="</tr>"

		out+="<tr>"
		out+="<td class='table_cell'><span "+style+" onClick='show_unde_tree("+strconv.Itoa(detectedActiveLastUnderstandingNodID)+")'>"+strconv.Itoa(detectedActiveLastUnderstandingNodID)+
		"</span></td><td class='table_cell'><span "+style+" onClick='show_atmzm_tree("+strconv.Itoa(detectedActiveLastNodID)+")'>"+strconv.Itoa(detectedActiveLastNodID)+
		"</span></td></tr>"
		out+="</table>"
/////////////  прерывания ращмышлений
if InterruptMemory!=nil && len(InterruptMemory)>0{
	out="<b>Прерывания размышлений:</b><br>"
	out+="<table cellpadding=0 cellspacing=0 border=1 class='main_table'>"
	out+="<tr>"
	out+="<th class='table_header'>ID итерации</th>"
	out+="<th class='table_header'>Тема размышления</th>"
	out+="<th class='table_header'>ID поставленной цели</th>"
	out+="</tr>"
	for i := 0; i < len(InterruptMemory); i++ {
		im := InterruptMemory[i]
		theme:=ThemeImageFromID[mentalInfoStruct.ThemeImageType]
	out+= "<tr><td class='table_cell'><span " + style + " )'>" + strconv.Itoa(im.goNextID) +
	"</span></td><td class='table_cell'><span " + style + " >" + ThemeTypeStr[theme.Type] +
	"</span></td><td class='table_cell'><span " + style + " onClick='get_purpose(" + strconv.Itoa(im.PurposeImageID) + ")'>" + strconv.Itoa(im.PurposeImageID) +
	"</span></td></tr>"
	}
	out+="</table><br>"
}

//////////////   циклы осмысления
	cycle:=getCyckleGoNextIdArr(activedCyckleID)
	if cycle==nil || len(cycle)==0{
		out+="Нет текущего цикла осмысления.<hr><br><br>"
	}else {// выдать таблицу всх составляющих цикла

cLen:=len(cycle)
isLimit:=0
if cLen >20{// не больше 20 показывать
	cycle=cycle[cLen:]
	isLimit=cLen-20
}
		out+=showGoNextCykles(cycle,isLimit)
	}

	if idlenessType>0{
		detStr:="Нет осознанной цели и проблемы"
		if idlenessType==1{
			detStr="Благополучное состояние (осоловелость)"
		}
		out+="<br><span style='style='font-size:19px;'>Нет мышления  - <b>ЛЕНИВОЕ СОСТОЯНИЕ:</b> "+
			detStr+"</span><br>"
	}


	// последнгие 20 кадров кратковременной памяти
	//	termMemory=[]shortTermMemory{{4800,5,1},{7188,4,2},{4800,3,3}} // тестирование
	if termMemory == nil || len(termMemory)==0{
		out+="Еще нет кратковременной памяти."
	}else {
	var termMemoryFrag []shortTermMemory
	if len(termMemory)>20{
		termMemoryFrag = termMemory[:20]
	}else{
		termMemoryFrag=termMemory
	}
	out+="<br><b>Кратковременная память (последние 20 кадров):</b><br>"

		out += "<hr><table cellpadding=0 cellspacing=0 border=1 class='main_table'>"
		out += "<tr><th class='table_header'>goNext ID</th>"
		out += "<th class='table_header'>ID дерева понимания</th>"
		out += "<th class='table_header'>ID дерева автоматизмов</th>"
		out += "</tr>"
		for i := len(termMemoryFrag) - 1; i >= 0; i-- {
			sm := termMemory[i]
			if sm.GoNextID == 0 {
				//return "Нулевой образ звена цикла в GetCicklesToPult()."
				out += "<tr><td colspan=10>Нулевой образ звена цикла в GetCicklesToPult() с ID = " + strconv.Itoa(sm.GoNextID) + "</td></tr>"
			}
			goNextFromIDMapCheck()
			gn := goNextFromID[sm.GoNextID]
			if gn == nil {
				out += "<tr><td colspan=10>Нет образа звена цикла с ID = " + strconv.Itoa(sm.GoNextID) + "</td></tr>"
			}else {
				style := "style='font-size:19px;font-weight:bold;cursor:pointer'"
				out += "<tr><td class='table_cell'><span style='color:#666666'>" + strconv.Itoa(gn.ID) +
					"</span></td><td class='table_cell'><span " + style + " onClick='show_unde_tree(" + strconv.Itoa(sm.uTreeNodID) + ")'>" + strconv.Itoa(sm.uTreeNodID) +
					"</span></td><td class='table_cell'><span " + style + " onClick='show_atmzm_tree(" + strconv.Itoa(gn.MotorBranchID) + ")'>" + strconv.Itoa(gn.MotorBranchID) +
					"</span></td></tr>"
			}
		}
		out += "</table>"
	}
	out:="Сделать показ логов циклов мышления по сохраненным в файлах !!!!!"
	return out
}*/
//////////////////////////////////////////
func showGoNextCykles(goNext []int, isLimit int) string {
	//style:="style='font-size:19px;font-weight:bold;cursor:pointer'"
	out := "<b>Последовательность инфо-функций:</b>"
	out += "<br>"

	for i := 0; i < len(goNext); i++ {
		gn := goNext[i]
		if gn != 0 {
			out += strconv.Itoa(gn) + ": " + getMentalFunctionString(gn)
		}
		if isLimit > 0 {
			out += "...Еще есть " + strconv.Itoa(isLimit) + " функций...<br>"
		}
		out += "<br>"
	}
	return out
}

//для http://go/pages/mental_rules.php инфа о ID goNext
/*
func GetGoNextInfo(id int)(string){
	if id==0{
		return "Нулевой образ звена цикла в GetGoNextInfo."
	}

	goNextFromIDMapCheck()
		gn:= goNextFromID[id]
		if gn==nil{
			return "Нет образа звена цикла с ID = "+strconv.Itoa(id)
		}
		style:="style='font-size:19px;font-weight:bold;cursor:pointer'"
		out:="<tr><td class='table_cell'><span style='color:#666666'>"+strconv.Itoa(id)+
			"</span></td><td class='table_cell'><span "+style+" onClick='show_atmzm_tree("+strconv.Itoa(gn.MotorBranchID)+")'>"+strconv.Itoa(gn.MotorBranchID)+"</span> "+
			"</span></td><td class='table_cell'><span "+style+" onClick='show_ment_act_img("+strconv.Itoa(gn.MentalActionsImagesID)+")'>"+strconv.Itoa(gn.MentalActionsImagesID)+
			"</span></td></tr>"

	return out
}*/
//////////////////////////////////////////////////////////////////
// для GetCicklesToPult() инфа о ветке дерева автоматизмов
func GetAtmzmTreeInfo(id int) string {
	out := ""

	//node:=AutomatizmTreeFromID[id]
	node, ok := ReadeAutomatizmTreeFromID(id)
	if !ok {
		return "Нет такого узла дерева автоматизмов."
	}
	out += getStrFromCond(0, node.BaseID) + "<br>"
	out += getStrFromCond(1, node.EmotionID) + "<br>"
	out += getStrFromCond(2, node.ActivityID) + "<br>"
	out += getStrFromCond(3, node.ToneMoodID) + "<br>"
	//if node.PhraseID>0 { нафиг масикировать лажи! первый символ не должен быть в ветке, если нет фразы!
	out += getStrFromCond(4, node.SimbolID) + "<br>"
	//}
	out += getStrFromCond(5, node.PhraseID) + "<br>"
	return out
}

// для GetCicklesToPult() инфа о ветке дерева понимания
func GetUndstgTreeInfo(id int) string {
	out := ""
	//	node:=UnderstandingNodeFromID[id]
	node, ok := ReadeUnderstandingNodeFromID(id)
	if !ok {
		return "Нет такого узла дерева понимания."
	}
	out += getMoodStr(node) + "<br>"
	out += getEmotionStr(node) + "<br>"
	out += getSituationStr(node) + "<br>"
	//out += getPurposeString(node.ID) + "<br>" // это дерево ситуации, а не дерево проблем
	return out
}

// для GetCicklesToPult() инфа о ментальном автоматизме
func GetMentAtmzmInfo(id int) string {
	/*out:=""
	//	atmz:=MentalAutomatizmsFromID[id]
		atmz,ok:=ReadeAFromID(id)
		if !ok{
			return "Нет такого узла дерева понимания."
		}
		if atmz.IsStaff==1{
			out += "Это - штатный автоматизм<br>"
		}else{
			out += "Это - НЕ штатный автоматизм<br>"
		}
		out += "Число шагов в цикле goNext: "+strconv.Itoa(len(atmz.funcArr))+"<br>"
		out += "(Бес)Полезность: "+strconv.Itoa(atmz.Usefulness)+"<br>"
		out += "ID запущенного моторного автоматизма: <span onClick='show_automatizms("+strconv.Itoa(atmz.motAutmtzmID)+")' style='cursor:pointer;color:blue'>"+strconv.Itoa(atmz.motAutmtzmID)+"</span><br>"
	*/
	return "Мент.Автоматизм - наиболее привычная цепь инфо-фукций - специфична к кокретным условиям."
}

//////////////////////////////////////////////////////////////////

// выдать текущую инфу Oбъекты значимости для http://go/pages/mental_importance.php
func GetImportanceToPult() string {
	/*
		//	saveFromNextIDcurretCicle=[]int{1,2,3} // тестирвоание
		out := ""
		iArr := importanceFromID
		if iArr == nil || len(iArr) == 0 {
			out += "Еще нет объектов значимости.<hr><br><br>"
		} else {
			out += "<table cellpadding=0 cellspacing=0 border=1 class='main_table'>"
			out += "<tr><th class='table_header'>ID образа<br>значимости</th>"
			out += "<th class='table_header'>ID<br>объекта</th>"
			out += "<th class='table_header'>ID<br>дерева<br>проблем</th></th>"
			out += "<th class='table_header' width=350>Тип объекта</th>"
			out += "<th class='table_header'>Значимость</th></tr>"

			keys := make([]int, 0, len(iArr))

			for k, v := range iArr {
				if v == nil {
					continue
				}
				keys = append(keys, k)
			}

			sort.Ints(keys)

			for _, k := range keys {
				//			oi,ok:=importanceFromID[k]
				oi, ok := ReadeImportanceFromID(k)
				if !ok {
					out += "<tr><td class='table_cell'>Нет объекта с ID=" + strconv.Itoa(k) + "</span>"
					continue
				}

				style := "style='font-size:19px;font-weight:bold;cursor:pointer'"
				out += "<tr><td class='table_cell'><span style='color:#666666'>" + strconv.Itoa(oi.ID) + "</span>" +
					"<td class='table_cell'>" + strconv.Itoa(oi.ObjectID) + "</td>"

				if oi.Type == 0 || oi.Type > 8 {
					out += "<tr><td class='table_cell'>Нет объекта с importanceTypeName =" + strconv.Itoa(oi.Type) + "</span>"
					continue
				}

				out += "<td class='table_cell' " + style + " onClick='show_problem_tree(" + strconv.Itoa(oi.ProblemID) + ")'>" + strconv.Itoa(oi.ProblemID)

				out += "</td><td class='table_cell' onClick='show_object(" + strconv.Itoa(oi.Type) + "," + strconv.Itoa(oi.ObjectID) + ")'>type=" + strconv.Itoa(oi.Type) + ": " +
					"</td><td class='table_cell'>" + strconv.Itoa(oi.Value) +
					"</td></tr>"
			}
		}
		out += "</table>"

		return out
	*/
	return ""
}

// информация об объекте значимости
func GetImportanceObjectInfo(objID int) string {
	out := ""

	out = GetActionsString(objID) +
		"<br><br><span style='cursor:pointer;color:blue' onClick='get_undestand_model(" + strconv.Itoa(objID) + ")'>Модель понимания объекта ID=" + strconv.Itoa(objID) + "</span>"

	return out
}

/////////////////////////////////////////////////

/*
func show_atmzm_tree(actID int)string{
	out:=""
	MentalActionsImagesMapCheck()
	act:=MentalActionsImagesArr[actID]
	switch act.typeID{
	case 1: out="активация настроения Mood в дереве понимания: "+strconv.Itoa(act.typeID)
	case 2: out="активация эмоции "+strconv.Itoa(act.typeID)+" в дереве понимания"
	case 3: out="активация PurposeImage "+strconv.Itoa(act.typeID)+"в дереве понимания"
	case 4: out="запуск инфо-функции infoFunc"+strconv.Itoa(act.typeID)+"()"
	case 5: out="запуск моторного автоматизма ID="+strconv.Itoa(act.typeID)
	case 6: out="запуск Доминанты ID="+strconv.Itoa(act.typeID)
	case 7: out="создание новой Доминанты"
	default: out="<span style=''>НЕТ ТАКОГО МЕНТАЛЬНОГО ДЕЙСТВИЯ</span> c ID="+strconv.Itoa(act.typeID)
	}

	return out
}*/
/////////////////////////////////////////////////

/* вставка строки с вызовом onClick
 */
func onClickStr(id int, onclick string, title string) string {
	out := "<b><span style='cursor:pointer;color:blue' title='" + title + "'"
	out += "onClick='" + onclick + "(" + strconv.Itoa(id) + ")'>" + strconv.Itoa(id) + "</span>" + "</b>"
	return out
}
func onClick2Str(par1 int, par2 int, id int, onclick string, title string) string {
	out := "<b><span style='cursor:pointer;color:blue' title='" + title + "'"
	out += "onClick='" + onclick + "(" + strconv.Itoa(par1) + "," + strconv.Itoa(par2) + ")'>" + strconv.Itoa(id) + "</span>" + "</b>"
	return out
}

///////////////////////////////////////////////////

// Страница Сознание
// var conscienceInfoCount=0
func GetConscienceInfo() string {
	var out = "" //strconv.Itoa(conscienceInfoCount)+"#|#"
	//out+=GetSelfPerceptionInfo()

	var cyclesListStr = ""
	/* сортировать по возрастанию order, можно - по ID циклов всегда возрастает с новым циклом.
	   Но главный - всегда последний (главным может стать любой предшествующий).
	*/
	var mainID = 0
	keys := make([]int, 0, len(cyclesArr))

	for id, v := range cyclesArr {
		if v == nil {
			continue
		}
		if v.isMainCycle {
			mainID = id
		} else {
			keys = append(keys, id)
		}
	}

	sort.Ints(keys)
	var n = 0
	for _, id := range keys {
		if n > 100 {
			ost := len(cyclesArr) - 100
			out += "&nbsp;&nbsp;&nbsp;&nbsp;Еще " + strconv.Itoa(ost) + " циклов..."
			break
		}
		n++
		cyclesListStr += setCicleStr(id)
	}
	if mainID > 0 {
		cyclesListStr += setCicleStr(mainID)
	}

	out += cyclesListStr + "#|#" // для кнопок текущих циклов

	// общая текущая инфа:
	out += "<br>Число параллельных цклов: <b>" + strconv.Itoa(len(cyclesArr)) + "</b><br>"
	out += "<br>Количество Прерываний осознания: <b>" + strconv.Itoa(len(InterruptMemory)) + "</b><br>"
	cStr := ""
	for i := 0; i < len(infoFuncSequence); i++ {
		if i > 0 {
			cStr += ", "
		}
		cStr += strconv.Itoa(infoFuncSequence[i])
	}
	if len(cStr) > 0 {
		out += "последовательность ID выполненных инфо-функций всех циклов: " + cStr + "< br > "
	}

	// текущий актуальный, главный, последний цикл при создании
	if activedCyckleID > 0 {
		out += "<br>Главный цикл мышления  с начальным ID=: " + onClickStr(activedCyckleID, "show_cyckle", "") + "<br>"
		out += GetCycleLocInfo0(activedCyckleID)
	} else {
		out += "<br>Нет главного цикла мышления.<br>"
	}
	return out
}

// //////////////////
func setCicleStr(cID int) string {
	out := ""
	//	cycle,ok:=cyclesArr[cID]
	cycle, ok := ReadecyclesArr(cID)
	if ok {
		if cycle.isMainCycle {
			//mainCycle=&v
			out += "m," + strconv.Itoa(cycle.ID) + "," + strconv.Itoa(cycle.order) + "|"
		} else {
			out += "0," + strconv.Itoa(cycle.ID) + "," + strconv.Itoa(cycle.order) + "|"
		}
	}
	return out
}

// /////////////////////////////////////
// инфа для пульта по get_short_info(
func GetCycleLocInfo0(cID int) string {
	//	c,ok:= cyclesArr[cID]
	c, ok := ReadecyclesArr(cID)
	if ok {
		return c.log
	}
	return "Нет цикла с ID=" + strconv.Itoa(cID)
}

////////////////////////////////////////////////////////////////

/*
	запись инфы о циклах в файл папки пульта /cycle_logs/

по каждому циклу - файл cycle.order+".txt"
С каждым пробуждением папка очищается чтобы начать новый лог.
*/
func updateCycleLogsFiles(c *cycleInfo) {
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/cycle_logs/"+strconv.Itoa(c.order)+".txt", c.log)
}

// //////////////////////////////////////////////////////////////
func clinerCycleLogsFiles() {
	dir := lib.GetMainPathExeFile() + "/cycle_logs/"
	lib.ClinerAllFromDir(dir)
}

////////////////////////////////////////////////////////////////

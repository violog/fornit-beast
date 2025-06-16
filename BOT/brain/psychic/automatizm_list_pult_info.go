/* список моторных автоматизмов для Пульта
для http://go/pages/automatizm_table.php
*/

package psychic

import (
	"BOT/brain/gomeostas"
	termineteAction "BOT/brain/terminete_action"
	"BOT/lib"
	"sort"
	"strconv"
	"strings"
)

//////////////////////////////////////////

// //////////////////////////////////////////
func GetAutomatizmInfo(limitBasicID int) string {
	var out = ""
	// сколько рефлексов есть
	uAutomatizmCount := len(AutomatizmFromId)
	// переключатель диапазона вывода
	out += "<br>Показывать: "

	out += "<span style='cursor:pointer;"
	if limitBasicID == 0 { // показать 100 последних с обратной сортировкой
		out += "background-color:#FFFF9D;font-weight:bold;"
	}
	out += "' onClick='show_level(0)'>100последних</span> "

	out += "<span style='cursor:pointer;"
	if limitBasicID == 1 {
		out += "background-color:#FFFF9D;font-weight:bold;"
	}
	out += "' onClick='show_level(1)'>Плохо</span> "

	out += "<span style='cursor:pointer;"
	if limitBasicID == 2 {
		out += "background-color:#FFFF9D;font-weight:bold;"
	}
	out += "' onClick='show_level(2)'>Норма</span> "

	out += "<span style='cursor:pointer;"
	if limitBasicID == 3 {
		out += "background-color:#FFFF9D;font-weight:bold;"
	}
	out += "' onClick='show_level(3)'>Хорошо</span> "

	out += "<span style='cursor:pointer;"
	if limitBasicID == 4 {
		out += "background-color:#FFFF9D;font-weight:bold;"
	}
	out += "' onClick='show_level(4)' title='Автоматизмы, не привязанные к определенным условиям состояния Beast.'>Свободные</span> "

	out += "<span style='cursor:pointer;"
	if limitBasicID == 5 {
		out += "background-color:#FFFF9D;font-weight:bold;"
	}
	out += "' onClick='show_level(5)' title='Автоматизмы с плохой Полезностью, которые блокируются при запуске.'>Заблокированные</span> "
	/////////////////////////////////

	header := "<tr><th width=70 class='table_header'>ID ав-ма</th>" +
		"<th width=70 class='table_header' style='background-color:#CCC5FF;' title='Если <1000000, то - узел ветки дерева\nЕсли <2000000 - ID образа действий\nостальное - ID фразы'>ID <br>объекта<br>Привязки</th>" +
		"<th width=70 class='table_header' style='background-color:#CCC5FF;'  title='ID базового состояния'>BaseID Дерева</th>" +
		"<th width='70' class='table_header' style='background-color:#CCC5FF;'  title='ID образа сочетания эмоций'>Эмоции Дерева</th>" +
		"<th width='10%' class='table_header' style='background-color:#CCC5FF;'  title='ID образа пускового стимула'>ДействияID-ФразаID-НастроениеID</th>" +
		"<th class='table_header'  title='Действия автоматизма, их может быть много видов.'>Образ действия</th>" +
		"<th width='30' class='table_header' title='ID следующего автоматизма в цепочке'>NextID</th>" +
		"<th width='30' class='table_header' title='Энергичность'>Energy</th>" +
		"<th width='30' class='table_header' title='(БЕС)ПОЛЕЗНОСТЬ: -10 вред 0 +10 +n польза'>Польза</th>" +
		"<th width='30' class='table_header' title='Уверенность: 0 - пробник, 1 - чужие сведения, 2 - уверенный (штатный)'>Belief</th>" +
		"<th width='30' class='table_header' title='Число успешных использований'>Надежность</th></tr>"

	out += "<table class='main_table'  cellpadding=0 cellspacing=0 border=1 width='1000px' style='font-size:14px;'>" +
		header

	if len(AutomatizmFromId) == 0 {
		return out + "</table><br>Подождите пока не активируется психика (не более 4 секунд) и нажмите <a href='/pages/automatizm_table.php'>Обновить</a>"
	}

	keys := make([]int, 0, uAutomatizmCount)

	for k, v := range AutomatizmFromId {
		if v != nil {
			keys = append(keys, k)
		}
	}

	if limitBasicID == 0 { // обратная сортировка
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] > keys[j]
		})
	} else {
		sort.Ints(keys)
	}
	///////////////////////////////////////////////////

	var count = 0
	for _, k := range keys {
		//v:= AutomatizmFromId[k]
		v, ok := ReadeAutomatizmFromId(k)
		if !ok {
			count++
			continue
		}
		if limitBasicID == 5 && v.Usefulness >= 0 {
			count++
			continue
		}
		if limitBasicID < 4 && v.BranchID >= 1000000 {
			count++
			continue
		}

		id := strconv.Itoa(k)
		nodeID := ""
		baseID := ""
		emotionID := ""
		emotionTitle := ""
		trStr := ""
		onclickID := ""
		//activityTitle:=""
		sStyle := "style='cursor:pointer;color:blue'"
		emClick := ""
		actClick := ""

		// узел дерева, к которому прикреплен автоматизм   && limitBasicID<4
		if v.BranchID < 1000000 {

			//	nodeA:=AutomatizmTreeFromID[v.BranchID]
			nodeA, ok := ReadeAutomatizmTreeFromID(v.BranchID)
			if ok {
				nodeID = strconv.Itoa(nodeA.ID)
				if limitBasicID > 0 && limitBasicID < 5 {
					if nodeA.BaseID != limitBasicID {
						count++
						continue
					}
				}
				baseID = strconv.Itoa(nodeA.BaseID) + " " + gomeostas.GetBaseCondFromID(nodeA.BaseID)

				/// эмоции
				emotionID = strconv.Itoa(nodeA.EmotionID)
				emo := EmotionFromIdArr[nodeA.EmotionID]
				if emo == nil { // так не должно быть! эмоция д.б. всегда
					//lib.WritePultConsol("В emo:=EmotionFromIdArr[nodeA.EmotionID] НЕТ ЭМОЦИИ!!!! Какой-то образ неверен, нужен программист!")
					// лучше бы выкинуть панику: panic("НЕТ ЭМОЦИИ!!!!")
					emotionTitle += "<span style='color:red'>НЕТ ЭМОЦИИ! нужно разобраться с EmotionFromIdArr[]</span>"
					count++
					continue
				}
				for i := 0; i < len(emo.BaseIDarr); i++ {
					if i > 0 {
						emotionTitle += ", "
					}
					emotionTitle += gomeostas.GetBaseContextCondFromID(emo.BaseIDarr[i])
				}
				/////////////////////////////
				if id == "19" {
					id = "19"
				}

				// пусковой стимул

				trStr = strconv.Itoa(nodeA.ActivityID) + "-" + strconv.Itoa(nodeA.PhraseID) + "-" + strconv.Itoa(nodeA.ToneMoodID)

				emClick = "onClick='show_emotion(" + emotionID + ")'"
				actClick = "onClick='show_trigger(" + nodeID + ")'"

			} //if nodeA!=nil{

			//if v.BranchID<1000000{
		} else {

			nodeID = strconv.Itoa(v.BranchID)
			onclickID = "onClick='show_object(" + strconv.Itoa(v.BranchID) + ")' style='cursor:pointer;color:blue'"
			baseID = "не"
			emotionID = "привязан"
			//	emotionTitle="привязан"
			trStr = "к дереву"
			sStyle = ""

		}

		//////////////
		out += "<tr >"
		out += "<td class='table_cell' >" + id + "</td>"
		out += "<td class='table_cell' " + onclickID + ">" + nodeID + "</td>"
		out += "<td class='table_cell' ><nobr>" + baseID + "</nobr></td>"
		out += "<td class='table_cell' title='" + emotionTitle + "' " + emClick + " " + sStyle + "><nobr>" + emotionID + "</nobr></td>"
		out += "<td class='table_cell' title='Информация по клику' " + actClick + " " + sStyle + ">" + trStr + "</td>"
		out += "<td class='table_cell' title='Информация по клику'  onClick='show_actions(" + id + ")' style='cursor:pointer;color:blue'><nobr>" + strconv.Itoa(v.ActionsImageID) + "</nobr></td>"
		out += "<td class='table_cell' title='Информация по клику'  onClick='show_next_actions(" + strconv.Itoa(v.NextID) + ")' style='cursor:pointer;color:blue'><nobr>" + strconv.Itoa(v.NextID) + "</nobr></td>"
		//>"+strconv.Itoa(v.NextID)+"</td>";
		out += "<td class='table_cell' >" + strconv.Itoa(v.Energy) + "</td>"
		var usefulness = strconv.Itoa(v.Usefulness)
		//if limitBasicID==5{
		if v.Usefulness < 0 {
			usefulness = "<span " + sStyle + " onClick='cliner_block(" + strconv.Itoa(v.ID) + ")' title='разблокировать (установить Полезность в 1)'><b>" + usefulness + "</b></span>"
		}
		out += "<td class='table_cell' >" + usefulness + "</td>"
		out += "<td class='table_cell' >" + strconv.Itoa(v.Belief) + "</td>"
		out += "<td class='table_cell' >" + strconv.Itoa(v.Count) + "</td>"
		out += "</tr>"
		count++
		if limitBasicID == 0 && count > 100 {
			break
		}
	}
	out += "</table>"
	return out
}

////////////////////////////////////

// показать пусковой стимул, к которому привязан автоматизм
func GetStrnameFromobjectID(objectID int) string {
	if objectID > 1000000 && objectID < 2000000 { // это - действия кнопок с пульта
		imgID := objectID - 1000000
		act := ActivityFromIdArr[imgID]
		if act == nil {
			return ""
		}
		out := ""
		for i := 0; i < len(act.ActID); i++ {
			if i > 0 {
				out += ", "
			}
			out += termineteAction.TerminalActonsNameFromID[act.ActID[i]]
		}
		return "Пусковые действия:<br><br><b>" + out + "</b>"
	}
	if objectID > 2000000 { // это - фраза
		imgID := objectID - 2000000
		return "Пусковая фраза:<br><br><b>" + GetPraseStringsFromVerbalID(imgID) + "</b>"
	}
	return ""
}

/////////////////////////////////////

/*
	Список фраз, для которых есть автоматизм Beast в этих условиях.

Для иконки http://go/img/words.png нал полем ввода.
*/
func GetAutomatizmPraseList(basicID int, contexts string) string {
	out := ""
	nCol := 0
	// образ эмоции
	var contextsArr []int
	cArr := strings.Split(contexts, ",")
	for i := 0; i < len(cArr); i++ {
		cID, _ := strconv.Atoi(cArr[i])
		contextsArr = append(contextsArr, cID)
	}
	emID, _ := createNewBaseStyle(0, contextsArr, true)
	// ШТАТНЫЕ автоматизмы, прикрепленные к ID узла Дерева
	var outArr []string
	for _, v := range AutomatizmBelief2FromTreeNodeId {
		//		brnch:=AutomatizmTreeFromID[v.BranchID]
		brnch, ok := ReadeAutomatizmTreeFromID(v.BranchID)
		if !ok {
			continue
		}
		if brnch.BaseID != basicID && brnch.EmotionID != emID {
			continue
		}
		if brnch.PhraseID != 0 {
			prase := GetPraseStringsFromVerbalID(brnch.PhraseID)
			//			prase:=word_sensor.GetPhraseStringsFromPhraseID(brnch.PhraseID)
			if len(prase) > 1 {
				if !lib.ExistsValInStringArr(outArr, prase) {
					outArr = append(outArr, prase)
				}
				//	}
			}
		}
	}
	if len(outArr) > 0 {
		sort.Strings(outArr)
		nCol = 0
		out = "<b>Штатные автоматизмы:</b> <table border=0 style='width:800px;font-size:14px;'><tr>"
		for i := 0; i < len(outArr); i++ {
			if !lib.ExistsValInStringArr(outArr, outArr[i]) {
				outArr = append(outArr, outArr[i])
			}
			if nCol == 6 {
				out += "</tr><tr>"
				nCol = 0
			}
			out += "<td align='left' style='cursor:pointer;' onClick='insert_pult_word(`" + outArr[i] + "`)'><nobr>" + outArr[i] + "</nobr></td>"
			nCol++
		}
		out += "</tr></table><hr>"
	}

	//автоматизмы, привязанные к ID фразе PhraseID и тогда их branchID начинается с 2000000
	outArr = nil
	out2 := ""
	for praaseID, _ := range AutomatizmIdFromPhraseId {
		prase := GetPraseStringsFromVerbalID(praaseID)
		if len(prase) > 1 {
			if !lib.ExistsValInStringArr(outArr, prase) {
				outArr = append(outArr, prase)
			}
		}
	}
	if len(outArr) > 0 {
		sort.Strings(outArr)
		nCol = 0
		out2 = "<span title='Эти автоматизмы срабаотывают по фразам, отвечая на них.'><b>Общие автоматизмы:</b></span> <table border=0 style='width:800px;font-size:14px;'><tr>"
		for i := 0; i < len(outArr); i++ {
			if !lib.ExistsValInStringArr(outArr, outArr[i]) {
				outArr = append(outArr, outArr[i])
			}
			if nCol == 6 {
				out2 += "</tr><tr>"
				nCol = 0
			}
			out2 += "<td align='left' style='cursor:pointer;' onClick='insert_pult_word(`" + outArr[i] + "`)'><nobr>" + outArr[i] + "</nobr></td>"
			nCol++
		}
		out2 += "</tr></table>"
		out += out2
	}

	return out
}

//////////////////////////////////////////////////

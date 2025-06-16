/* Для Пульта выдать текущее состояние Самоощещения.

 */

package psychic

import (
	"BOT/brain/gomeostas"
	"strconv"
)

///////////////////////////////////////

// информация самоощущения для Пульта
func GetSelfPerceptionInfo() string {

	//!!! refreshCurrentInformationEnvironment() инфа просто перекрывается новой
	// опасность
	danger := GetAttentionDanger()
	// актуальность ситуации
	veryActualSituation, _ := gomeostas.FindTargetGomeostazID()

	// против паники типа "одновременная запись и считывание карты"
	//	if notAllowScanInTreeThisTime{
	//		return "!!!"
	//	}
	ie := CurrentInformationEnvironment
	// нужно задать постоянную высоту блока чтобы он не мигал по высоте <div style='min-height:230px;'>
	var out = "<div style='min-height:230px;'><br>Общее состояние жизненных параметров: <b>" //background-color:#FFEBFC;
	if gomeostas.CommonBadNormalWell == 1 {
		out += "ПЛОХО"
	}
	if gomeostas.CommonBadNormalWell == 2 {
		out += "НОРМА"
	}
	if gomeostas.CommonBadNormalWell == 3 {
		out += "ХОРОШО"
	}
	out += "</b>"
	////////////////////////////
	out += "<br>Ощущаемое настроение: <b>"
	if ie.PsyMood == -1 {
		out += "Плохое (" + strconv.Itoa(ie.Mood) + ")"
	}
	if ie.PsyMood == 0 {
		out += "Нормальное (" + strconv.Itoa(ie.Mood) + ")"
	}
	if ie.PsyMood == 1 {
		out += "Хорошее (" + strconv.Itoa(ie.Mood) + ")"
	}
	out += "</b>"
	////////////////////////////////
	out += "<br>Текущая эмоция: " + GetCurrentEmotionReception() //getEmotonsComponentStr(ie.PsyEmotionImg)
	///////////////////////////////
	out += "<br>Опасность состояния: <b>"
	if danger {
		out += "Опасное"
	} else {
		out += "Неопасное"
	}
	out += "</b>"
	out += "<br>Важность состояния: <b>"
	if veryActualSituation {
		out += "Очень важное состояние"
	} else {
		out += "Спокойное состояние."
	}
	out += "</b>"

	switch idlenessType {
	case 0:
		out += "<br><span style='color:#888888'>Лень: <b>отсуствует</b></span>"
	case 1:
		out += "<br>Лень: <b>гомеостатическая</b>"
	case 2:
		out += "<br>Лень: <b>осознанная лень</b>"
	}

	////////////////////////////////
	if ie.ActionsImageID > 0 {
		str := GetActionsString(ie.ActionsImageID)
		out += "<br>Текущий образ сочетания действий с Пульта: <b>" + str + "</b>"
	} else {
		out += "<br>текущий образ сочетания действий с Пульта: отсуствует"
	}
	////////////////////////////////

	out += "<br>Субъективно ощущаемая оценка, текущее осознаваемое настроение: <b>" + strconv.Itoa(ie.PsyMood) + "</b><br>"

	if ie.AnswerImageID > 0 {
		str := GetActionsString(ie.AnswerImageID)
		out += "<br>Текущий образ сочетания ОТВЕТНОГО действия мот.автоматизма: <b>" + str + "</b>"
	} else {
		out += "<br>текущий образ сочетания ОТВЕТНОГО действия мот.автоматизма: отсуствует"
	}

	if ie.ExtremImportanceObjectID > 0 {
		//oi, ok := importanceFromID[ie.ExtremImportanceObjectID]
		oi := getExtremObjFromID(ie.ExtremImportanceObjectID)
		if oi != nil {
			out += "<br>Нет объекта importanceTypeName </span>"
		} else {
			out += "<br>Наиболее важный образ типа extremImportance: " + onClickStr(1, "show_object", "")
		}
	}
	out += "<br>Текущая Доминанта нерешенной проблемы: " + onClickStr(ie.DominantaID, "show_dominant", "")

	// TODO остальное

	out += "</div>"
	return out
}

///////////

/* образы действий, оператора или действий Beast.
Образ действий оператора - это Стимул (образ восприятия)
Образ действий Beast - это Акция (образ действия)

При каждой стимуляции с Пульта Дерева автоматизмов возникает образ восприятия curActiveActionsID, curActiveActions
Еще есть функция "вытащить образ Стимула ActiveActions из ID узла дерева автоматизмов"
func getActiveActionsFromAutomatizmTreeNode(automatizmTreeNodeID int)(int,*ActionsImage)

Фактически структура повторяет TriggerStimuls из рефлексов и позволяет сохранять
как образы действий в автоматизмах, так и образы действий оператора, отражаемые в дереве мот.автомтаизмов.
Используется для формирования пар стимул (действия оператора) - действия (ответ beast)
для эпизодической памяти и структуры rules - Правил примитивного опыта.

Обоснование:
Обощенные образы восприятия, возникающие в теменной ассоциативной коре полностью соотвествуют воспринимаемому
и не могут меняться.
Но в лобной коре есть отражение этих образов,
с возможностью произвольно создавать любые новые из известных элементов старого.
Поэтому для области рефлексов используется TriggerStimuls,
а для области психики - ActionsImage (с меткой Answer int // 0 - объектиное действие, 1 - субъективное предположение).
Эти два вида структур локализуются по обе стороны двигательных программ
(https://scorcher.ru/optional_class/MVAP_lections.pdf)
образуя основу "зеркальной" системы подражания
(https://scorcher.ru/axiomatics/axiom_show.php?id=522).

Формат: ID|RSarr через ,|PhraseID через ,|ToneID|MoodID|Answer
*/

package psychic

import (
	termineteAction "BOT/brain/terminete_action"
	word_sensor "BOT/brain/words_sensor"
	"BOT/lib"
	"strconv"
	"strings"
)

///////////////////////////////////////////

/*
  - это метка AmtzmNextString ID (десять миллионов) -

ID акции является цепочкой действий AmtzmNextString, а не одним действием
используется для automatizm_next.go
*/
var prefixActionIdValue = 10000000 // если ID лействия больше prefixActionIdValue, то AmtzmNextString.ID=ID-prefixActionIdValue
//////////////////////////////////////////

// универсальный образ восритятия Стимула и образ действия имеют одну и ту же структуру.
type ActionsImage struct {
	ID int // идентификатор данного сочетания пусковых стимулов
	/* метка о том, что действия не является объективным Стимулом (реально воспринятым из Пульта)
	   или реально выполненным действием (предположение, Правило из сновидения)
	*/
	Kind int // 0 - объектиное действие, 1 - субъективное предположение

	ActID []int // массив действий с Пульта или Ответного действия
	// для текущего сообщения с Пусльта:
	PhraseID []int // массив фразID (DetectedUnicumPhraseID) слова каждой фразы вытаскиваются wordSensor.GetWordArrFromPhraseID(PhraseID[0])
	ToneID   int   // тон сообщения с Пульта или Ответного действия
	/* настроение при передаче фразы с Пульта:
		20-Хорошее    21-Плохое    22-Игровое    23-Учитель    24-Агрессивное   25-Защитное    26-Протест
	ID возникает при добавлении 19 к номеру радиокнопки пульта, например, для Хорошее 1+19=20
	*/
	MoodID int // настроение оператора или Ответного действия
}

// var ActionsImageArr=make(map[int]*ActionsImage)
var ActionsImageArr []*ActionsImage // сам массив
// var AFromID = make([]*aNode, 20000)//задать сразу имеющиеся в файле число при загрузке
// запись члена
func WriteActionsImageArr(index int, value *ActionsImage) {
	if index >= len(ActionsImageArr) {
		newSlice := make([]*ActionsImage, index+1)
		copy(newSlice, ActionsImageArr)
		ActionsImageArr = newSlice
	}
	ActionsImageArr[index] = value
}

// считывание члена
func ReadeActionsImageArr(index int) (*ActionsImage, bool) {
	if index >= len(ActionsImageArr) || ActionsImageArr[index] == nil {
		return nil, false
	}
	return ActionsImageArr[index], true
}

//////////////////////////////////////////1

// вызывается из psychic.go
func ActionsImageInit() {
	loadActionsImageArr()
}

// //////////////////////////////////////////////
// создать новое сочетание ответных действий если такого еще нет
var lastActionsImageID = 0

func CreateNewlastActionsImageID(id int, Kind int, ActID []int, PhraseID []int, ToneID int, MoodID int, CheckUnicum bool) (int, *ActionsImage) {
	// не создавать образ с пустым действием и вербальным сенсором - такое может быть при новом слове и отключенной галке Форсажа
	if ActID == nil && (PhraseID == nil || isUnrecognizedPhraseFromAtmtzmTreeActivation) {
		return 0, nil
	}

	if CheckUnicum {
		oldID, oldVal := checkUnicumActionsImage(Kind, ActID, PhraseID, ToneID, MoodID)
		if oldVal != nil {
			return oldID, oldVal
		}
	}

	if id == 0 {
		lastActionsImageID++
		id = lastActionsImageID
	} else {
		if lastActionsImageID < id {
			lastActionsImageID = id
		}
	}

	var node ActionsImage
	node.ID = id
	node.Kind = Kind
	node.ActID = ActID
	node.PhraseID = PhraseID
	node.ToneID = ToneID
	node.MoodID = MoodID

	//ActionsImageArr[id]=&node
	WriteActionsImageArr(id, &node)

	if doWritingFile {
		SaveActionsImageArr()
	}

	return id, &node
}
func checkUnicumActionsImage(Kind int, ActID []int, PhraseID []int, ToneID int, MoodID int) (int, *ActionsImage) {

	for id, v := range ActionsImageArr {
		if v == nil || Kind != v.Kind {
			continue
		}
		if !lib.EqualArrs(ActID, v.ActID) {
			continue
		}
		if !lib.EqualArrs(PhraseID, v.PhraseID) {
			continue
		}
		if ToneID != v.ToneID || MoodID != v.MoodID {
			continue
		}

		return id, v
	}

	return 0, nil
}

/////////////////////////////////////////

// ////////////////// сохранить образы сочетаний ответных действий
// В случае отсуствия ответных действий создается ID такого отсутсвия, пример такой записи: 2|||0|0|
// ID|ActID через ,|PhraseID через ,|ToneID|MoodID|Answer
func SaveActionsImageArr() {
	var out = ""
	for k, v := range ActionsImageArr {
		if v == nil {
			continue
		}
		out += strconv.Itoa(k) + "|"
		for i := 0; i < len(v.ActID); i++ {
			out += strconv.Itoa(v.ActID[i]) + ","
		}
		out += "|"
		for i := 0; i < len(v.PhraseID); i++ {
			out += strconv.Itoa(v.PhraseID[i]) + ","
		}
		out += "|"
		out += strconv.Itoa(v.ToneID) + "|"
		out += strconv.Itoa(v.MoodID) + "|"
		out += strconv.Itoa(v.Kind)
		out += "\r\n"
	}

	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/action_images.txt", out)

}

// //////////////////  загрузить образы сочетаний ответных действий
// ID|ActID через ,|PhraseID через ,|ToneID|MoodID|Answer
func loadActionsImageArr() {
	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/action_images.txt")
	cunt := len(strArr)
	ActionsImageArr = make([]*ActionsImage, cunt) //задать сразу имеющиеся в файле число
	for n := 0; n < cunt; n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		p := strings.Split(strArr[n], "|")
		id, _ := strconv.Atoi(p[0])

		s := strings.Split(p[1], ",")
		var ActID []int
		for i := 0; i < len(s); i++ {
			if len(s[i]) == 0 {
				continue
			}
			si, _ := strconv.Atoi(s[i])
			ActID = append(ActID, si)
		}

		s = strings.Split(p[2], ",")
		var PhraseID []int
		for i := 0; i < len(s); i++ {
			if len(s[i]) == 0 {
				continue
			}
			si, _ := strconv.Atoi(s[i])
			PhraseID = append(PhraseID, si)
		}
		ToneID, _ := strconv.Atoi(p[3])
		MoodID, _ := strconv.Atoi(p[4])
		Kind := 0
		if len(p[5]) > 0 {
			Kind, _ = strconv.Atoi(p[5])
		}

		var saveDoWritingFile = doWritingFile
		doWritingFile = false
		CreateNewlastActionsImageID(id, Kind, ActID, PhraseID, ToneID, MoodID, false)
		doWritingFile = saveDoWritingFile
	}
	return

}

///////////////////////////////////////////

// выдать строку из одиночной или цепочечной (AmtzmNextString) акции
func GetActionsString(act int) string {
	var out = ""
	if act < prefixActionIdValue {
		return getSingleAtcStr(act)
	} else { // это последовательность действий типа AmtzmNextString
		lib.MapFree(MapGwardAutomatizmNextStringFromID)
		anArr, ok := AutomatizmNextStringFromID[act-prefixActionIdValue]
		if !ok {
			return ""
		}
		out = "<b>Цепочка действий:</b> "
		for i := 0; i < len(anArr.next); i++ {
			if i > 0 {
				out += " | "
			}
			out += getSingleAtcStr(anArr.next[i])
		}
		return out
	}

	return ""
}

// строку из одиночной акции
func getSingleAtcStr(act int) string {
	if ActionsImageArr == nil || len(ActionsImageArr) == 0 {
		return ""
	}

	//ai, ok := ActionsImageArr[act]
	ai, ok := ReadeActionsImageArr(act)
	if !ok {
		return ""
	}
	out := ""
	var kind = ""
	if ai.Kind == 0 {
		kind = "Объективнй образ"
	}
	if ai.Kind == 1 {
		kind = "Субъективнй образ"
	}
	if ai.ActID != nil {
		out += "Действие (" + kind + "): "
		for i := 0; i < len(ai.ActID); i++ {
			if i > 0 {
				out += ", "
			}
			actName := termineteAction.TerminalActonsNameFromID[ai.ActID[i]]
			out += "<b>" + actName + "</b>"
		}
		out += " "
	}

	if ai.PhraseID != nil {
		out += "Фраза: "
		for i := 0; i < len(ai.PhraseID); i++ {
			if i > 0 {
				out += " "
			}
			prase := word_sensor.GetPhraseStringsFromPhraseID(ai.PhraseID[i])
			out += "<b>\"" + prase + "\"</b>"
		}
		out += " "
	}

	if ai.ToneID != 0 {
		out += " Тон: " + getToneStrFromID(ai.ToneID) + " "
	}

	// так просто нельзя менять ActionsImageArr !!! Могту появляться дубли после изменений. Нужно обязательтно проверять на уникальность
	/*	if ai.MoodID < 0 {
				ai.MoodID *= -1
			}
			if ai.MoodID >= 19 { последний ID равен 26 а не 19!!!! т.е. if ai.MoodID >= 27 {
				ai.MoodID -= 19
			}
		Пусть остается закомментированным, если будут вылезать лажи нужно будет смотреть где и почему
			сбивается диапазон значений настроения. Просто маскировать это неправильно.
	*/
	if ai.MoodID != 0 {
		out += " Настрой: " + getMoodStrFromID(ai.MoodID) + "<br>"
	}
	return out
}

////////////////////////////////////////

/* Рефлексы мозжечка для корректировки автоматизмов

Допонение автоматизма другими корректирующими действиеями
или корректировка самого автоматизма.

 Для размышления:
24	 	Как оператору оценивать результат, когда бот выдает автоматизм с большей энергичностью, если это касается вербального общения? Типа он громче/злее отвечает?
25	 	Если бы это был механический автоматизм типа пары антагонистических движений, в этом был бы смысл как в попытке найти оптимальное корректирующее действие.
26	 	А сейчас это просто бестолковое накручивание энергичности которое никак не оценивается оператором потому, что не понятно, как это оценивать.
27	 	Может вообще побрить эти мозжечковые коррекции как бессмысленные для нашего вербального проекта?
28	 	Или может привязать вместо энергичности изменение тона и контекста сообщения, что будет более логично.
29	 	Тем более, что до сих пор мы никак не использовали их в плане адаптации.
30
31	 	Если вместо увеличения энергичности ставить в порядке возрастания (см. флажки на пульте):
32	 	тон повышенный + контекст нормальный
33	 	тон повышенный + контекст агрессивный
34	 	тон повышенный + контекст защитный
35	 	тон повышенный + контекст протест
36
37	 	И в порядке убывания:
38	 	тон вялый + контекст нормальный
39	 	тон вялый + контекст плохой
40
41	 	Это уже не будет +-10 единиц, но по сути это будет активация дополнительных элементов воздействия на оператора, которые могли в теории образоваться
42	 	как рефлекторная реакция на изменение энергичности автоматизма. То есть изменение энергичности выполнения автоматизма во многих случаях это добавление
43	 	каких то составляющих к нему потому, что не все автоматизмы это мышечные сокращения.
*/

package psychic

import (
	"BOT/brain/transfer"
	"BOT/lib"
	"strconv"
	"strings"
)

/*
	по результатам выполнения автоматизма выбираются дополнительные действия

или изменяется сила действия автоматизма.
Это - средство не переписывать автоматизм, а оптимизировать его.
В качестве дополнительных действий используются имеющиеся автоматизмы на основе которых создаются
мозжечковые рефлексы
*/
type cerebellumReflex struct {
	id                 int
	typeAut            int // тип корректруемого автоматизма: 0 - это ID моторного, 1 - ID ментального
	sourceAutomatismID int // корректируемый моторный или ментальный автоматизм
	addEnergy          int // добавление (-убавление) силы действия Energy. Может быть отрицательное число, чтобы уменьшить энергию автоматизма
	// кроме корректирования самого автоматизма по силе действия, могут быть запущены дополнительные автоматизмы:
	additionalAutomatismID []int // массив ID дополнительных моторных автоматизмов
	additionalMentalAutID  []int // массив ID дополнительных ментальных автоматизмов
}

// общий массив рефлексов мозжечка
var cerebellumReflexFromID = make(map[int]*cerebellumReflex)

// рефлексы мозжечка по моторному ID
var cerebellumReflexFromMotorsID = make(map[int]*cerebellumReflex)

// рефлексы мозжечка по ментальному ID
var cerebellumReflexFromMentalsID = make(map[int]*cerebellumReflex)

func cerebellumReflexInit() {
	loadCerebellumReflex()
}

// последний ID рефлексов мохжечка
var lastCRid = 0

// создать новый автоматизм
// В случае отсуствия автоматизма создается ID такого отсутсвия, пример такой записи: 2|||0|0| - ID=2
func createNewCerebellumReflex(id int, typeAut int, sourceAutomatismID int, CheckUnicum bool) (int, *cerebellumReflex) {
	if CheckUnicum {
		oldID, oldVal := checkUnicumCerebellumReflex(typeAut, sourceAutomatismID)
		if oldVal != nil {
			return oldID, oldVal
		}
	}

	if id == 0 {
		lastCRid++
		id = lastCRid
	} else {
		if lastCRid < id {
			lastCRid = id
		}
	}

	var node cerebellumReflex
	node.id = id
	node.typeAut = typeAut
	node.sourceAutomatismID = sourceAutomatismID
	node.addEnergy = 5 // сразу придаст максимум, сложившись с энергией автоматизма :)
	// node.additionalAutomatismID = additionalAutomatismID
	cerebellumReflexFromID[id] = &node
	if typeAut == 0 {
		cerebellumReflexFromMotorsID[sourceAutomatismID] = &node
	} else {
		cerebellumReflexFromMentalsID[sourceAutomatismID] = &node
	}
	if doWritingFile {
		SaveCerebellumReflex()
	}
	return id, &node
}

// поиск рефлекса мозжечка
func checkUnicumCerebellumReflex(typeAut int, sourceAutomatismID int) (int, *cerebellumReflex) {
	for id, v := range cerebellumReflexFromID {
		if v == nil {
			continue
		}
		if typeAut == v.typeAut && sourceAutomatismID == v.sourceAutomatismID {
			return id, v
		}
	}
	return 0, nil
}

// Сохранить рефлексы мозжечка
// структура записи: id|typeAut|sourceAutomatismID||addEnergy|additionalAutomatismID|additionalMentalAutID
// В случае отсуствия пусковых стимулов создается ID такого отсутсвия, пример такой записи: 2|||0|0|
func SaveCerebellumReflex() {
	var out = ""
	for k, v := range cerebellumReflexFromID {
		if v == nil {
			continue
		}
		out += strconv.Itoa(k) + "|"
		out += strconv.Itoa(v.typeAut) + "|"
		out += strconv.Itoa(v.sourceAutomatismID) + "|"
		out += strconv.Itoa(v.addEnergy) + "|"
		for i := 0; i < len(v.additionalAutomatismID); i++ {
			out += strconv.Itoa(v.additionalAutomatismID[i]) + ","
		}
		out += "|"
		for i := 0; i < len(v.additionalMentalAutID); i++ {
			out += strconv.Itoa(v.additionalMentalAutID[i]) + ","
		}
		out += "\r\n"
	}
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/cerebellum_reflex.txt", out)
}

// Загрузить рефлексы мозжечка
// структура записи: id|typeAut|sourceAutomatismID||addEnergy|additionalAutomatismID|additionalMentalAutID
func loadCerebellumReflex() {
	cerebellumReflexFromID = make(map[int]*cerebellumReflex)
	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/cerebellum_reflex.txt")
	for n := 0; n < len(strArr); n++ {
		p := strings.Split(strArr[n], "|")
		if len(p) < 5 {
			return
		}
		id, _ := strconv.Atoi(p[0])
		typeAut, _ := strconv.Atoi(p[1])
		sourceAutomatismID, _ := strconv.Atoi(p[2])
		addEnergy, _ := strconv.Atoi(p[3])
		a := strings.Split(p[4], "|")
		var additionalAutomatismID []int
		for i := 0; i < len(a); i++ {
			aid, _ := strconv.Atoi(a[i])
			additionalAutomatismID = append(additionalAutomatismID, aid)
		}
		a = strings.Split(p[5], "|")
		var additionalMentalAutID []int
		for i := 0; i < len(a); i++ {
			aid, _ := strconv.Atoi(a[i])
			additionalMentalAutID = append(additionalMentalAutID, aid)
		}
		var saveDoWritingFile = doWritingFile
		doWritingFile = false
		_, ca := createNewCerebellumReflex(id, typeAut, sourceAutomatismID, false)
		doWritingFile = saveDoWritingFile
		ca.addEnergy = addEnergy
		ca.additionalAutomatismID = additionalAutomatismID
		ca.additionalMentalAutID = additionalMentalAutID
	}
	return
}

// вернуть скорректированную силу действия
func getCerebellumReflexAddEnergy(kind int, automatismID int) int {
	var e *cerebellumReflex

	if transfer.IsPsychicGameMode {
		return 0
	} // в игровом режиме не должно быть коррекций, иначе обучающие действия кнопок начнут неадекватно прогрессировать
	if kind == 0 {
		e = cerebellumReflexFromMotorsID[automatismID]
	} else {
		e = cerebellumReflexFromMentalsID[automatismID]
	}

	if e == nil {
		return 0
	}
	return e.addEnergy
}

// выполнить дополнительные мозжечковые автоматизмы сразу после выполняющегося автоматизма
var wasRunAutmzmID = 0 // защелка от бесконечного цикла RumAutomatism() - runCerebellumAdditionalAutomatism() - RumAutomatismID() - RumAutomatism()
var wasRunMentalAutmzmID = 0

func runCerebellumAdditionalAutomatism(kind int, automatismID int) {
	if wasRunAutmzmID > 0 { // только один вызов runCerebellumAdditionalAutomatism для данного автоматизма
		wasRunAutmzmID = 0
		return
	}
	if wasRunMentalAutmzmID > 0 { // только один вызов runCerebellumAdditionalAutomatism для данного автоматизма
		wasRunMentalAutmzmID = 0
		return
	}
	var cr *cerebellumReflex

	if kind == 0 {
		cr = cerebellumReflexFromMotorsID[automatismID]
	} else {
		cr = cerebellumReflexFromMentalsID[automatismID]
	}
	if cr == nil {
		return
	}
	if kind == 0 {
		aArr := cr.additionalAutomatismID
		for i := 0; i < len(aArr); i++ {
			wasRunAutmzmID = aArr[i]
			RumAutomatismID(aArr[i])
		}
	} else {
		aArr := cr.additionalMentalAutID
		for i := 0; i < len(aArr); i++ {
			// TODO это старый вариант мент.автоматизмов, который нужно просто заменить на моторные автоматизмы, копровождаюзие мозжечковый рефлекс.
			wasRunMentalAutmzmID = aArr[i] //???????
			//RunMentalAutomatismsID(aArr[i])
		}
	}
}

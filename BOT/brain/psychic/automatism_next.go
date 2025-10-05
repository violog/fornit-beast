/* Automatism.NextID - поддержка вторичных автоматизмов цепочки, запускаемой от первого автоматизма.
Это - аналог природных программ действия, запускаетмых по пусковому стимулу.
Next - следующий автоматизм в цепочке исполнения.
Цепь запускается автоматически при запуске первого автоматизма из func GetAutomotizmActionsString()
Цепь может быть пройдена прогностически (чтобы узнать каким действием заканчивается), ментально, без выполнения ее автоматизмов,
для этого не вызывается моторное выполнение, а просто - проход цепочки с просмотром ее звеньев
или цепь может быть прервана осознанно.

В реальности есть цепочки последовательности действий, но строго по модальности,
например, последовательности мышечных сокращений. Это дает возможность моментального прогноза,
чем закончилось все действие и возможность осознанно корректировать любую фазу действий.
Вот что-то такое и нужно реализовать.
Хотя это - цепочка действий, но у в структуре действий нет NextID
и автоматизм полностью обеспечивает последовательность действий.

Так что AmtzmNextString - это просто продолжение действия в цепочке.
Для AmtzmNextString Energy наследуется от родителя-автоматизма
	и поэтому к AmtzmNextString ПОКА не применяется мозжечковый рефлекс.

Для общих автоматизмов (.BranchID>1000000) AmtzmNextString применять не следует как и для if EvolushnStage < 4 {

Чаще всего цепочка AmtzmNextString должно появляться в состоянии обучения при отзеркаливании чудого действия
с момента его предудыщего действия.
Но может быть добавлено в ходе эксперимента в состоянии любопытства.
В обоих случаях нужно использовать infoFunc22().
*/

package psychic

import (
	"BOT/lib"
	"strconv"
	"strings"
)

///////////////////////////////////////////////

type AmtzmNextString struct {
	ID   int
	next []int // цепочка ID ActionsImage
}

var AutomatismNextStringFromID = make(map[int]*AmtzmNextString)
var MapGwardAutomatismNextStringFromID = lib.RegNewMapGuard()

///////////////////////////////////////

/* создать новый AmtzmNextString
 */
var lastAmtzmNextStringID = 0 // ID последнего созданного AmtzmNextString
func createAmtzmNextStringID(id int, ActionsImageIDArr []int, CheckUnicum bool) (int, *AmtzmNextString) {
	// не разрешать создавать цепочки из нулевых образов действий
	if lib.ExistsValInArr(ActionsImageIDArr, 0) {
		return 0, nil
	}
	// убрать дублеры в массиве
	ActionsImageIDArr = lib.UniqueArr(ActionsImageIDArr)

	if CheckUnicum {
		oldID, oldVal := checkUnicumAmtzmNextString(ActionsImageIDArr)
		if oldVal != nil {
			return oldID, oldVal
		}
	}

	if id == 0 {
		lastAmtzmNextStringID++
		id = lastAmtzmNextStringID
	} else {
		if lastAmtzmNextStringID < id {
			lastAmtzmNextStringID = id
		}
	}

	var node AmtzmNextString
	node.ID = id
	node.next = ActionsImageIDArr

	lib.MapCheckWrite(MapGwardAutomatismNextStringFromID)
	AutomatismNextStringFromID[id] = &node
	lib.MapFree(MapGwardAutomatismNextStringFromID)

	return id, &node
}

//////////////////////////////////////////
/* Функцию можно использовать для выборки ActionsImageIDArr по ActionsImageID
Вернет оригинал из всех найденных клонов, т.к. у клона всегда ID больше оригинала.
*/
func checkUnicumAmtzmNextString(ActionsImageIDArr []int) (int, *AmtzmNextString) {
	if AutomatismFromId == nil {
		return 0, nil
	}
	minID := 1000000000
	var nS *AmtzmNextString
	lib.MapCheckBlock(MapGwardAutomatismNextStringFromID)
	for _, v := range AutomatismNextStringFromID {
		if v == nil {
			continue
		}
		if !lib.EqualArrs(ActionsImageIDArr, v.next) {
			continue
		}
		//return v.ID,v
		if v.ID < minID {
			minID = v.ID
			nS = v
		}
	}
	lib.MapFree(MapGwardAutomatismNextStringFromID)
	if nS != nil {
		return nS.ID, nS
	}
	return 0, nil
}

////////////////////////////////////////////

/*
	создать новую цепочку на основе существующей - клонировать

клоны создаются для удлинения и становятся не похожи на оригинал,
но некоторое время они тождестенны по next
*/
func createCloneNextStringFromID(id int) (int, *AmtzmNextString) {
	lib.MapCheck(MapGwardAutomatismNextStringFromID)
	nS, ok := AutomatismNextStringFromID[id]
	if !ok {
		return 0, nil
	}
	return createCloneNextString(nS)
}

// создать новую цепочку на основе существующей - клонировать
func createCloneNextString(nstr *AmtzmNextString) (int, *AmtzmNextString) {
	lastAmtzmNextStringID++
	id := lastAmtzmNextStringID
	var node AmtzmNextString
	node.ID = id
	node.next = nstr.next

	lib.MapCheckWrite(MapGwardAutomatismNextStringFromID)
	AutomatismNextStringFromID[id] = &node
	lib.MapFree(MapGwardAutomatismNextStringFromID)

	return id, &node
}

///////////////////////////////////////////

// найти цепочку AmtzmNextString с таким действием, если нет - создать, если есть создать новый клон
func createClonForActins(act []int) int {
	id, nS := checkUnicumAmtzmNextString(act)
	if id > 0 {
		id, _ = createCloneNextString(nS)
	} else {
		id, _ = createAmtzmNextStringID(0, act, true)
	}
	return id
}

////////////////////////////////////////////////

// СОХРАНИТЬ структура записи: id|ActionsImage[]
func SaveAmtzmNextString() {
	var out = ""
	lib.MapCheckBlock(MapGwardAutomatismNextStringFromID)
	for k, v := range AutomatismNextStringFromID {
		out += strconv.Itoa(k) + "|"
		for i := 0; i < len(v.next); i++ {
			out += strconv.Itoa(v.next[i]) + ","
		}
		out += "\r\n"
	}
	lib.MapFree(MapGwardAutomatismNextStringFromID)
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/automatism_next.txt", out)
}

// ЗАГРУЗИТЬ структура записи: id|ActionsImage[]
func loadAmtzmNextString() {
	AutomatismNextStringFromID = make(map[int]*AmtzmNextString)

	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/automatism_next.txt")
	if strArr == nil {
		return
	}
	for n := 0; n < len(strArr); n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		p := strings.Split(strArr[n], "|")
		id, _ := strconv.Atoi(p[0])
		s := strings.Split(p[1], ",")
		var next []int
		for i := 0; i < len(s); i++ {
			if len(s[i]) == 0 {
				continue
			}
			ai, _ := strconv.Atoi(s[i])
			next = append(next, ai)
		}
		createAmtzmNextStringID(id, next, false) // без проверки на уникальность
	}
	return
}

/////////////////////////////////////////////////////////////

/* Создать новую цепочку AmtzmNextString с привязкой к Automatism.NextID
 */
func createNextAutomatism(NextID int, ActionsImageID []int, CheckUnicum bool) (int, *AmtzmNextString) {

	//parent,ok:=AutomatismFromId[NextID]
	parent, ok := ReadeAutomatismFromId(NextID)
	if !ok {
		lib.TodoPanic("Родитель для создания вторичного автоматизма с ID=" + strconv.Itoa(NextID) + " не существует!")
	}
	if parent.BranchID > 1000000 || EvolushnStage < 4 {
		return 0, nil
	}

	idN, am := createAmtzmNextStringID(0, ActionsImageID, true)

	parent.NextID = idN // дорастили цепочку
	wasCreateNextAtmtztID = parent.ID

	return idN, am
}

////////////////////////////////////////////////////////////

/* Добавить к цепочке AmtzmNextString автоматизма Automatism.NextID новое звено образа действия
 */
func addNextAutomatism(atmztmID int, ActionsImageID []int, CheckUnicum bool) (int, *AmtzmNextString) {

	//parent,ok:=AutomatismFromId[atmztmID]
	parent, ok := ReadeAutomatismFromId(atmztmID)
	if !ok {
		lib.TodoPanic("Родитель для создания вторичного автоматизма с ID=" + strconv.Itoa(atmztmID) + " не существует!")
	}
	if parent.BranchID > 1000000 || EvolushnStage < 4 {
		return 0, nil
	}

	idN, am := createAmtzmNextStringID(0, ActionsImageID, true)

	parent.NextID = idN // дорастили цепочку
	wasCreateNextAtmtztID = parent.ID

	return idN, am
}

////////////////////////////////////////////////////////////

/*
	Добавить к цепочке AmtzmNextStringID новое звено образа действия

если еще нет AmtzmNextStringID, то сначала создать.
*/
func addAmtzmNextString(AmtzmNextStringID int, ActionsImageID []int) (int, *AmtzmNextString) {
	if AmtzmNextStringID == 0 {
		sN, am := createAmtzmNextStringID(0, ActionsImageID, true)
		return sN, am
	}
	lib.MapCheck(MapGwardAutomatismNextStringFromID)
	sN, ok := AutomatismNextStringFromID[AmtzmNextStringID]
	if !ok {
		sN, am := createAmtzmNextStringID(0, ActionsImageID, true)
		return sN, am
	}
	for i := 0; i < len(ActionsImageID); i++ {
		sN.next = append(sN.next, ActionsImageID[i])
	}
	// убрать дублеры в массиве
	sN.next = lib.UniqueArr(sN.next)

	return sN.ID, sN
}

////////////////////////////////////////////////////////////

/*
	Вставить в цепочку Automatism.NextID новой образ действия, - в любую часть цепочки.

При pos==-1 вставляется в конец. При pos==0 - в начало.
*/
func insertImageActionToAmtzmNextID(atmztmID int, pos int, imgActID int) {

	//	parent,ok:=AutomatismFromId[atmztmID]
	parent, ok := ReadeAutomatismFromId(atmztmID)
	if !ok {
		lib.TodoPanic("Родитель для создания вторичного автоматизма с ID=" + strconv.Itoa(atmztmID) + " не существует!")
	}
	if parent.BranchID > 1000000 || EvolushnStage < 4 {
		return
	}
	if parent.NextID == 0 { // еще нет, создать next и првязать есдинственное звено (любая позиция pos)
		createNextAutomatism(parent.ID, []int{imgActID}, true)
		return
	}

	lib.MapCheck(MapGwardAutomatismNextStringFromID)
	nArr, ok := AutomatismNextStringFromID[parent.NextID]
	if !ok {
		return
	}
	insertImageActionToAmtzmNext(nArr, pos, imgActID)
}
func insertImageActionToAmtzmNext(next *AmtzmNextString, pos int, imgActID int) {
	if pos == -1 || pos >= len(next.next) {
		next.next = append(next.next, imgActID)
		return
	}
	nArr := next.next
	var newArr []int
	for i := 0; i < len(nArr); i++ {
		if i == pos {
			newArr = append(newArr, imgActID)
		}
		newArr = append(newArr, nArr[i])
	}
	next.next = newArr
}

////////////////////////////////////////////////////////////

/*
	удалить образ действия с позицией в цепочке pos

При pos==-1 - последний. При pos==0 - первый.
*/
func deleteImageActionToAmtzmNextID(nextID int, pos int) {
	lib.MapCheck(MapGwardAutomatismNextStringFromID)
	nArr, ok := AutomatismNextStringFromID[nextID]
	if !ok {
		return
	}
	deleteImageActionToAmtzmNext(nArr, pos)
}
func deleteImageActionToAmtzmNext(next *AmtzmNextString, pos int) {
	if pos == -1 || pos >= (len(next.next)-1) {
		next.next = next.next[:len(next.next)-1] // удалить последний элемент
		return
	}
	if pos == 0 {
		next.next = next.next[1:] // удалить первый элемент
		return
	}
	nArr := next.next
	var newArr []int
	for i := 0; i < len(nArr); i++ {
		if i == pos {
			continue
		}
		newArr = append(newArr, nArr[i])
	}
	next.next = newArr
}

////////////////////////////////////////////////////////////////

/*
	создать автоматизм из цепочки действий

м.б., что в этой ветке уже есть автоматизм с действием, не даполненный NextID... он дополнится.
*/
func createAutomatismFromNextString(next []int, branchID int) (int, *Automatism) {
	if len(next) < 2 { // только одно действие
		return 0, nil
	}
	// создать автоматизм из первой акции
	id, atmtzm := createNewAutomatismID(0, branchID, next[0], true)
	if len(next) > 1 {
		// создать AmtzmNextString из остальной части действий
		nOstID, _ := createAmtzmNextStringID(0, next[1:], true)
		// остальная часть действий
		atmtzm.NextID = nOstID
	}

	return id, atmtzm
}

/////////////////////////////////////////////////////////////////

/* создать AmtzmNextString из всех действий автоматизма
 */
func createNextStringFromAutomatism(atmtzm *Automatism) (int, *AmtzmNextString) {
	var n []int
	n = append(n, atmtzm.ActionsImageID)
	if atmtzm.NextID > 0 {
		lib.MapCheck(MapGwardAutomatismNextStringFromID)
		nArr := AutomatismNextStringFromID[atmtzm.NextID]
		for i := 0; i < len(nArr.next); i++ {
			n = append(n, nArr.next[i])
		}
	}
	// создать AmtzmNextString из остальной части действий
	id, nObj := createAmtzmNextStringID(0, n, true)

	return id, nObj
}

/////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////
/* Запуск действия Automatism.NextID в цепочке,
для func RumAutomatism()
*/

func showNextAtmtzmAction(atmzmID int, NextID int, Energy int) string {
	out := showVolutionAction(NextID, Energy)
	lastRunVolutionAction = nil // не использовать механизм периода ожидания автоматизма
	return out
}

/////////////////////////////////////////////

////////////
/* произвольный запуск действий
приводит к периоду ожидания LastRunAutomatismPulsCount и оценке результата.
*/
//lastRunVolutionAction очищается после периода ожидания:
var lastRunVolutionAction *AmtzmNextString // последняя запущенная цепочка на исполнение.

var lastRunVolutionPulsCount = 0
var prevLastRunVolutionAction *AmtzmNextString // предпоследняя запущенная цепочка на исполнение - для ученического Правила
var prevLastRunVolutionPulsCount = 0

/*
	Совершение произвольного действия (не автоматизм, а предшествующая альтернатива).

Запуск всех действия NextID в цепочке.
Может запускаться произвольно, как свободные действия вне автоматизма (showNextAtmtzmAction(0,AmtzmNextStringID,5)).
*/
func showVolutionAction(NextID int, Energy int) string {
	// Не запускать в период ожидания ответа оператора. Такое же при запуске автоматизма.
	if LastRunAutomatismPulsCount > 0 {
		return ""
	}
	lib.MapCheck(MapGwardAutomatismNextStringFromID)
	as, ok := AutomatismNextStringFromID[NextID]
	if !ok {
		return ""
	}
	var out = ""
	for i := 0; i < len(as.next); i++ {
		if i > 0 {
			out += ""
		}

		//ai:=ActionsImageArr[as.next[i]]
		ai, ok := ReadeActionsImageArr(as.next[i])
		if !ok {
			continue
		}
		// каждый образ - в одну строку
		if ai.ActID != nil {
			out += TerminateMotorAutomatismActions(0, ai.ActID, Energy, true) + " "
		}
		if ai.PhraseID != nil {
			out += TerminatePraseAutomatismActions(ai.PhraseID, Energy) + " "
		}

		if ai.ToneID != 0 {
			out += " (" + getToneStrFromID(ai.ToneID)
		}

		if ai.MoodID != 0 {
			out += " " + getMoodStrFromID(ai.MoodID) + ")"
		}
	}

	// для периода ожидания ответа оператора
	prevLastRunVolutionAction = lastRunVolutionAction
	prevLastRunVolutionPulsCount = lastRunVolutionPulsCount
	lastRunVolutionAction = as
	lastRunVolutionPulsCount = PulsCount

	LastRunAutomatismPulsCount = PulsCount // активность произвольного действия в чисде пульсов
	LastDetectedActiveLastNodID = detectedActiveLastNodID

	return out
}

/////////////////////////////////////////////

////////////////////////////////////////////

/*
	действия с цепочкой при плохом эффекте ВРЕМЕННО-ПРЕДПОЛОЖИТЕЛЬНАЯ ВЕРСИЯ

Пока что создать дубль автоматизма из первого в цепочке, но с NextID=0 и сделать его штатным с полезностью 0.
Полезность и коунтер - начальные (==0).

Результат оценивается в func calcAutomatismResult() с записью 2 видов Правил.
*/
var wasCreateNextAtmtztID = 0

func badEffectChainAtmtzm(am *Automatism) {

	/* если стало плохо после добавления нового звена, то
	 */
	if wasCreateNextAtmtztID == am.ID {
		//сделать дубликат для штатного
		_, dAm := createDuplicateAutomatism(am.BranchID, am)
		// в дубликате привязкать цепочку Neat
		dAm.NextID = am.NextID
		// в дубликате  удалить последнее звено Next
		deleteImageActionToAmtzmNextID(dAm.NextID, -1)
		// сделать автоматизм штатным
		SetAutomatismBelief(dAm, 2)
	}

	// TODO - пока больше ничего
}

/////////////////////////////////////////////////

// /////////////////////////////////////////////////
// найти готовый или создать новый автоматизм из AmtzmNextString.ID
func createAndRunAutomatismFromAmtzmNextString(actionID int) *Automatism {
	// на всякий случай убережемся, что акция - именно типа AmtzmNextString
	if actionID < prefixActionIdValue { // НЕ последовательность образов действий AmtzmNextString.ID
		return nil
	}
	nID := actionID - prefixActionIdValue // чистый ID без префикса
	lib.MapCheck(MapGwardAutomatismNextStringFromID)
	nextString, ok := AutomatismNextStringFromID[nID]
	if !ok { // такого не должно быть, НО...
		return nil
	}
	_, atmtzm := createAutomatismFromNextString(nextString.next, detectedActiveLastNodID)
	return atmtzm
}

/////////////////////////////////////////////////////

// //////////////////////////////////////////////
// выдать на пульт инфу о цепочке действий
func GetNextActionsInfo(nextID int) string {
	if nextID == 0 {
		return "Отсутствует цепочка действий."
	}

	lib.MapCheck(MapGwardAutomatismNextStringFromID)
	as, ok := AutomatismNextStringFromID[nextID]
	if !ok {
		return ""
	}
	var out = "<b>Цепочка действий:</b> "
	for i := 0; i < len(as.next); i++ {
		if i > 0 {
			out += "<br>"
		}
		out += getSingleAtcStr(as.next[i])
	}
	return out
}

/////////////////////////////////////////

func GetNextActionsInfoList() string {
	var out = ""
	lib.MapCheckBlock(MapGwardAutomatismNextStringFromID)
	for k, v := range AutomatismNextStringFromID {
		if v == nil {
			continue
		}
		out += " <span style='cursor:pointer;color:blue' onClick='show_next_actions(" + strconv.Itoa(k) + ")'>ID=<b>" + strconv.Itoa(k) + "</b></span> Действия цепочки: "
		for i := 0; i < len(v.next); i++ {
			out += strconv.Itoa(v.next[i]) + ","
		}
		out += "<br>\r\n"
	}
	lib.MapFree(MapGwardAutomatismNextStringFromID)
	return out
}

//////////////////////////////////////////////

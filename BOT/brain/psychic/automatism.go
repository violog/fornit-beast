/* Автоматизмы, могут совершать внешние действия или внутренние произвольные действия.
К ветке дерева может быть прикреплено сколько угодно автоматизмов: GetMotorsAutomatismListFromTreeId(branchID)
но только один из автоматизмов, прикрепленных к ветке, может иметь Belief=2 - проверенное собственное знание
Автоматизмы могут быть и не привязаны к конкретной ветке дерева, а быть привязаны к отдельным значениям AutomatismNode:
- к ID образа действий с пульта ActivityID и тогда branchID начинается с 1000000,
сохраняются в карте AutomatismIdFromActionId
- к ID фразы PhraseID  и тогда branchID начинается с 2000000,
сохраняются в карте AutomatismIdFromPhraseId

Если задается Belief=2, остальные Belief=2 становится Belief=0.
!!! ПОЭТОМУ ВСЕГДА нужно задавать с помощью SetAutomatismBelief(atmzm *Automatism,belief int))

Если для прикрепленных к узлу дерева есть карта штатных AutomatismBelief2FromTreeNodeId,
то для прикрепленных к образам нужны ФУНКЦИИ ПОЛУЧЕНИЯ ШТАТНОГО ДЛЯ ДАННОГО ОБРАЗА:
func GetAutomatismBeliefFromActionId(activityID int)(*Automatism){
func GetAutomatismBeliefFromPhraseId(verbalID int)(*Automatism){

Концепция общих автоматизмов. Они сформированы на основе общего шаблона рефлексов и, как и рефлексы, перекрываются автоматизмами конца активной ветки, т.е. имеющими образ Стимула. Это – первичная реакция на текущее сочетание контекстов (в данном случае – на эмоцию). У узла эмоции может быть бесконечное число образов действия и блокировка общего автоматизма лишает первичной реакции у всех их. Но у.рефлекс может быть заморожен, а общий автоматизм сейчас – нет.
Это значит, то в случае блокирующих действий для данной ветки, необходимо запускать автоматизм бездействия, останавливающий все более низкоуровневое.
При блокировке такого автоматизма опять НЕ БЛОКРУЕТСЯ.
Наличие игнор.автоматизма в конце ветки для func consciousnessElementary() равноценно отсуствю автоматизма.

Структура записи: id|BranchID|Usefulness|ActionsImageID|NextID|Energy|Belief|motAutmtzmID
*/

package psychic

import (
	"BOT/lib"
	"strconv"
	"strings"
)

// инициализирующий блок - в порядке последовательности инициализаций
// вызывается из psychic.go
func automatismInit() {
	loadAutomatism()
	//res:=RumAutomatism(AutomatismFromId[1])
	//if res{}
}

type Automatism struct {
	ID int
	/* id объекта к кторому привязан автоматизм:
	   (он может быть  привязан к узлу дерева, к фразе (AutomatismIdFromPhraseId) или действиям (AutomatismIdFromActionId)
	   т.е. втоматизмы могут быть и не привязаны к конкретному узлу ветки дерева,
	   а быть привязаны к отдельным значениям AutomatismNode:
	      к ID образа действий с пульта ActivityID и тогда branchID начинается с 1000000,
	   сохраняются в карте AutomatismIdFromActionId
	      к ID фразы PhraseID  и тогда branchID начинается с 2000000,
	   сохраняются в карте AutomatismIdFromPhraseId

	   Если бы автоматизм не привязывался к ветке через BranchID, а чтобы была карта для ветки с привязанными автоматизмами.
	   Тогда один и тот же автоматизм мог бы быть привязан к любым веткам.
	   Но тогда он бы утратил специфику NextID, полезности и повторяемости для данной ветки.
	   Хотя можно было бы сделать структуру для каждой ветки для привязки автоматизмов и учета их специфики, но тогда автоматизм просто вырождается в действие и становится не нужным.
	*/
	BranchID int
	//TODO: достаточно -1, 0, +1. В паре с Count это даст величину градации (бес)полезности автоматизма
	Usefulness int // (БЕС)ПОЛЕЗНОСТЬ: -10 вред 0 +10 +n польза
	// образ действий типа ActionsImage
	ActionsImageID int

	/* Next - следующая цепочка действий типа AmtzmNextString:
	т.е. последовательность действий ActionsImageID, начинающаяся с AmtzmNextString.ID.
	Цепь может быть пройдена ментально, без выполнения автоматизмов,
	для этого не вызывается моторное выполнение а просто - проход цепочки с просмотром ее звеньев.
	или цепь может быть прервана осознанно.
	*/
	NextID int

	/* Энергичность действия или фразы.
	      Но т.к. автоматизм может использоваться в разных случаях,
	      лучше для этих конкретных случаев использования уточнять энергичность
	      с помощью мозжечковых рефлексов.
	   Для NextID Energy наследуется от родителя и поэтому к Next не применяется мозжечковый рефлекс.
	*/
	Energy int // от 1 до 10, по умолчанию = 5
	/* Уверенность в авторитарном автоматизме высока в период авторитарного обучения
		и сильно падает в период собственной инициативы, когда нужно на себе проверить,
		а даст ли такой автоматизм в самом деле обещанное улучшение.
		Только один из автоматизмов, прикрепленных к ветке, может иметь Belief=2 - проверенное собственное знание
		Если задается Belief=2, остальные Belief=2 становится Belief=0.
	!!! ПОЭТОМУ ВСЕГДА нужно задавать с помощью SetAutomatismBelief(atmzm *Automatism,belief int)
	*/
	Belief int // 0 - предположение, 1 - чужие сведения, 2 - проверенное собственное знание
	/* В случае, если в результате автоматизма его Usefulness изменит знак, то
	motAutmtzmID обнулится, а при таком же знаке - увеличивается на 1.
	*/
	Count int // надежность: число использований с подтверждением (бес)полезности Usefulness
	/* какие ID гомео-параметров улучшает это действие
	по аналогии и функциональности с http://go/pages/terminal_actions.php
	*/
	GomeoIdSuccesArr []int
}

// все, привязанные к узлу дерева или привязанные к id образа действия и к id фразы.
// var AutomatismFromId = make(map[int]*Automatism)
var AutomatismFromId []*Automatism

// запись члена
func WriteAutomatismFromId(index int, value *Automatism) {
	if index >= len(AutomatismFromId) {
		newSlice := make([]*Automatism, index+1)
		copy(newSlice, AutomatismFromId)
		AutomatismFromId = newSlice
	}
	AutomatismFromId[index] = value
}

// считывание члена
func ReadeAutomatismFromId(index int) (*Automatism, bool) {
	if index >= len(AutomatismFromId) || AutomatismFromId[index] == nil {
		return nil, false
	}
	return AutomatismFromId[index], true
}

///////////////////////////////////////

// ШТАТНЫЕ автоматизмы, прикрепленные к ID узла Дерева с Belief==2 т.е. ШТАТНЫЕ, выполняющиеся не раздумывая
// у узла может быть только один штатный автоматизм с Belief==2
var AutomatismBelief2FromTreeNodeId = make(map[int]*Automatism)

// ЭТОТ ПУСТЬ БУДЕТ НА КАРТЕ

// ///////////////////////////////////////////////////////////////
// привязанные к ID образа действий с пульта ActivityID и тогда их branchID начинается с 1000000
// среди привязанный к данному образуID может быть один штатный с Belief==2
var AutomatismIdFromActionId = make(map[int][]*Automatism)

// привязанные к ID фразы PhraseID и тогда их branchID начинается с 2000000
// среди привязанных к данной фразеID (неважны предыдущие условия) может быть один штатный с Belief==2
var AutomatismIdFromPhraseId = make(map[int][]*Automatism)

/*
	список удачных автоматизмов, относящихся к определенным условиям (привзяанных к определенной ветке Дерева)

В этом списке поле Usefulness >0
*/
var AutomatismSuccessFromIdArr = make(map[int]*Automatism)

// GetMotorsAutomatismListFromTreeId(nodeID int) список всех автоматизмов для ID узла Дерева
// ExistsAutomatismForThisNodeID(nodeID int) есть ли штатный автоматизм (с Belief==2), привязанные к узлу дерева
// GetBelief2AutomatismListFromTreeId(nodeID int) штатный, невредный автоматизм, привязанный к ветке

var lastAutomatismID = 0        // ID последнего созданного автоматизма
var NoWarningCreateShow = false // true - не выдавать сообщение о новом автоматизме

// удалить автоматизм
func deleteAutomatism(a *Automatism) {
	delete(AutomatismBelief2FromTreeNodeId, a.ID)

	//	delete(AutomatismFromId, a.ID)
	WriteAutomatismFromId(a.ID, nil)

	a = nil
}

////////////////////////////////////////////

/*
	создать новый автоматизм

checkLevel - глубина проверки на идентичность: 0 - нет проверки, 1 - поверхностная, 2 - полная
*/
func createNewAutomatismID(id int, BranchID int, ActionsImageID int, CheckUnicum bool) (int, *Automatism) {
	/* Автоматизмы уникальны по сочетанию BranchID и ActionsImageID.
	   	При попытке создать с таким же сочетанием возвращается уже созданный.
	    к одной вентке могут быть прикреплены неограниченное число автоматизмов
	*/

	// не создавать автоматизм с нулевым образом - такое может быть при новом слове и отключенной галке Форсажа
	if ActionsImageID == 0 {
		return 0, nil
	}
	if CheckUnicum {
		oldID, oldVal := checkUnicumMotorsAutomatism(BranchID, ActionsImageID)
		if oldVal != nil {
			return oldID, oldVal
		}
	}

	if id == 0 {
		lastAutomatismID++
		id = lastAutomatismID
	} else {
		if lastAutomatismID < id {
			lastAutomatismID = id
		}
	}

	var node Automatism
	node.ID = id
	node.BranchID = BranchID
	node.Energy = 5
	node.ActionsImageID = ActionsImageID

	//	AutomatismFromId[id] = &node
	WriteAutomatismFromId(id, &node)

	if BranchID > 1000000 && BranchID < 2000000 {
		imgID := BranchID - 1000000
		AutomatismIdFromActionId[imgID] = append(AutomatismIdFromActionId[imgID], &node)
	}
	if BranchID > 2000000 {
		imgID := BranchID - 2000000
		AutomatismIdFromPhraseId[imgID] = append(AutomatismIdFromPhraseId[imgID], &node)
	}

	if !NoWarningCreateShow {
		lib.WritePultConsol("Создан новый автоматизм.")
	}
	return id, &node
}

/*
	Автоматизмы уникальны по сочетанию BranchID и ActionsImageID.

Функцию можно использовать для выборки автоматизма с заданными BranchID и ActionsImageID
*/
func checkUnicumMotorsAutomatism(BranchID int, ActionsImageID int) (int, *Automatism) {
	if AutomatismFromId == nil {
		return 0, nil
	}

	for id, v := range AutomatismFromId {
		if v == nil || (BranchID != v.BranchID || ActionsImageID != v.ActionsImageID) {
			continue
		}

		return id, v
	}

	return 0, nil
}

// создать новый автоматизм
func CreateNewAutomatism(BranchID int, ActionsImageID int) (int, *Automatism) {
	// BranchID может быть ==0 для мозжечковых рефлексов
	if ActionsImageID == 0 {
		return 0, nil
	}

	id, verb := createNewAutomatismID(0, BranchID, ActionsImageID, true)

	if doWritingFile {
		SaveAutomatism()
	}

	return id, verb
}

/*
	создать дубликат автоматизма по образу имеющегося am *Automatism

т.е. неуникальный по привязке к ветке дерева и действию - для пополнения коллекции автоматизмов ветки.
Полезность и коунтер - начальные (==0).
*/
func createDuplicateAutomatism(BranchID int, am *Automatism) (int, *Automatism) {
	id, am := createNewAutomatismID(0, BranchID, am.ActionsImageID, false)
	return id, am
}

// создать новый автоматизм без записи в файл
func CreateAtutomatizmNoSaveFile(BranchID int, ActionsImageID int) (int, *Automatism) {
	// BranchID может быть ==0 для мозжечковых рефлексов
	if ActionsImageID == 0 {
		return 0, nil
	}

	id, verb := createNewAutomatismID(0, BranchID, ActionsImageID, true)

	return id, verb
}

// СОХРАНИТЬ структура записи: id|BranchID|Usefulness|ActionsImageID|NextID|Energy|Belief
// В случае отсуствия пусковых стимулов создается ID такого отсутсвия, пример такой записи: 2|||0|0|
func SaveAutomatism() {
	var out = ""
	for k, v := range AutomatismFromId {
		if v == nil {
			continue
		}
		out += strconv.Itoa(k) + "|"
		out += strconv.Itoa(v.BranchID) + "|"
		out += strconv.Itoa(v.Usefulness) + "|"
		out += strconv.Itoa(v.ActionsImageID) + "|"
		out += strconv.Itoa(v.NextID) + "|"
		out += strconv.Itoa(v.Energy) + "|"
		out += strconv.Itoa(v.Belief) + "|"
		out += strconv.Itoa(v.Count) + "|"
		for i := 0; i < len(v.GomeoIdSuccesArr); i++ {
			out += strconv.Itoa(v.GomeoIdSuccesArr[i]) + ","
		}
		out += "\r\n"
	}

	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/automatism_images.txt", out)
}

// ЗАГРУЗИТЬ структура записи: id|BranchID|Usefulness||ActionsImageID||NextID|Energy|Belief
func loadAutomatism() {
	NoWarningCreateShow = true

	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/automatism_images.txt")
	if strArr == nil {
		return
	}
	//AutomatismFromId = make(map[int]*Automatism)
	AutomatismFromId = make([]*Automatism, len(strArr)) //задать сразу имеющиеся в файле число
	for n := 0; n < len(strArr); n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		p := strings.Split(strArr[n], "|")
		id, _ := strconv.Atoi(p[0])
		BranchID, _ := strconv.Atoi(p[1])
		Usefulness, _ := strconv.Atoi(p[2])
		ActionsImageID, _ := strconv.Atoi(p[3])
		NextID, _ := strconv.Atoi(p[4])
		Energy, _ := strconv.Atoi(p[5])
		Belief, _ := strconv.Atoi(p[6])
		Count, _ := strconv.Atoi(p[7])
		s := strings.Split(p[4], ",")
		var GomeoIdSuccesArr []int
		for i := 0; i < len(s); i++ {
			if len(s[i]) == 0 {
				continue
			}
			sp, _ := strconv.Atoi(s[i])
			GomeoIdSuccesArr = append(GomeoIdSuccesArr, sp)
		}
		var saveDoWritingFile = doWritingFile
		doWritingFile = false
		_, a := createNewAutomatismID(id, BranchID, ActionsImageID, false) // без проверки на уникальность
		a.NextID = NextID
		a.Usefulness = Usefulness
		a.Energy = Energy
		a.Count = Count
		SetAutomatismBelief(a, Belief)
		doWritingFile = saveDoWritingFile
	}
	NoWarningCreateShow = false
	return
}

/////////////////////////////////////////////////////////////

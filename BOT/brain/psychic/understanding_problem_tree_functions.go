/* Функции для Дерева понимания проблемы
запись ID|ParentNode|autTreeID|situationTreeID|themeID|purposeID
*/

package psychic

import (
	"BOT/lib"
	"strconv"
	"strings"
)

///////////////////////////////

func getCurSituation() *SituationImage {
	//	stID:=UnderstandingNodeFromID[detectedActiveLastUnderstandingNodID]
	stID, ok := ReadeUnderstandingNodeFromID(detectedActiveLastUnderstandingNodID)
	if !ok {
		return nil
	}
	node, ok := ReadeSituationImageFromIdArr(stID.SituationID)
	if ok {
		return node
	}
	return nil
}

// /////////////////////////////////////
func getCurTheme() *ThemeImage {
	//stID:=ProblemTreeNodeFromID[detectedActiveLastProblemNodID]
	stID, ok := ReadeProblemTreeNodeFromID(detectedActiveLastProblemNodID)
	if !ok {
		return nil
	}

	node, ok := ReadeThemeImageFromID(stID.themeID)
	if ok {
		return node
	}
	return nil
}

// /////////////////////////////////////
func getCurPurpose() *PurposeImage {
	//stID:=ProblemTreeNodeFromID[detectedActiveLastProblemNodID]
	stID, ok := ReadeProblemTreeNodeFromID(detectedActiveLastProblemNodID)
	if !ok {
		return nil
	}

	node, ok := ReadePurposeImageFromID(stID.purposeID)
	if ok {
		return node
	}
	return nil
}

///////////////////////////////////////

/*
	Создать новый узел дерева понимания проблемы.

Формат записи:
ID|ParentNode|autTreeID|situationTreeID|themeID|purposeID
*/
var lastProblemTreeNodeID = 0

func createNewProblemTreeNode(parent *ProblemTreeNode, id int, autTreeID int, situationTreeID int,
	themeID int, purposeID int, CheckUnicum bool) (int, *ProblemTreeNode) {

	if parent == nil {
		return 0, nil
	}

	/*	var flgNoAtmzAdt=false
		// пока закоментим - может и не надо это отсекать
		if atmz,ok:=AutomatizmTreeFromID[autTreeID];ok{
			if atmz.PhraseID==0 && atmz.SimbolID==0{
				if atmzAct,ok:=ActivityFromIdArr[atmz.ActivityID];ok{
					if atmzAct.ActID==nil{
						flgNoAtmzAdt=true
					}
				}else{
					flgNoAtmzAdt=true
				}
			}
		}
		if flgNoAtmzAdt{ //не зачем создавать ветку проблем, если в пусковом стимуле автоматизма нет акций и вербальных действий
			return 0,nil
		}*/

	// если есть такой узел, то не создавать
	if CheckUnicum {
		idOld, nodeOld := FindProblemTreeNodeFromCondition(autTreeID, situationTreeID, themeID, purposeID)
		if idOld > 0 {
			return idOld, nodeOld
		}
	}

	if id == 0 {
		lastProblemTreeNodeID++
		id = lastProblemTreeNodeID
	} else {
		if lastProblemTreeNodeID < id {
			lastProblemTreeNodeID = id
		}
	}

	var node ProblemTreeNode
	node.ID = id
	node.ParentNode = parent
	node.ParentID = parent.ID

	node.autTreeID = autTreeID
	node.situationTreeID = situationTreeID
	node.themeID = themeID
	node.purposeID = purposeID

	parent.Children = append(parent.Children, node)

	// четко находим новый вставленный член (а не &parent.Children[count-1])
	var newN *ProblemTreeNode
	for i := 0; i < len(parent.Children); i++ {
		if parent.Children[i].ID == node.ID {
			newN = &parent.Children[i]
		}
	}

	//ProblemTreeNodeFromID[id]=&node
	WriteProblemTreeNodeFromID(id, &node)

	// т.к. append меняет длину массива, перетусовывая адреса, то нужно обновить адреса в ProblemTreeNodeFromID:
	updateProblemTreeNodeFromID(parent) // здесь потому, что при загрузке из файла нужно на лету получать адреса

	return id, newN
}

// создать первые 4 ветки базовых состояний
func createBasicProblemTree() {
	notAllowScanInTreeThisTime = true // запрет показа карты при обновлении
	ProblemTree.ID = 0
	//ProblemTreeNodeFromID[0] = &ProblemTree
	WriteProblemTreeNodeFromID(0, &ProblemTree)

	if doWritingFile {
		SaveProblemTree()
	}
	// SaveProblemTree()
	notAllowScanInTreeThisTime = false // запрет показа карты при обновлении
}

// корректируем адреса всех узлов
func updateProblemTreeNodeFromID(parent *ProblemTreeNode) {
	//updatingProblemTreeNodeFromID(&VernikePhraseTree)
	updatingProblemTreeNodeFromID(parent)
}

// проход всего дерева
func updatingProblemTreeNodeFromID(rt *ProblemTreeNode) {
	if rt.ID > 0 {
		//rt.ParentNode=ProblemTreeNodeFromID[rt.ParentID] // wr.ParentNode адрес меняется из=за corretsParent(,
		node, ok := ReadeProblemTreeNodeFromID(rt.ParentID)
		if ok {
			rt.ParentNode = node
		}
		ProblemTreeNodeFromID[rt.ID] = rt
	}
	if rt.Children == nil { // конец ветки
		return
	}
	for i := 0; i < len(rt.Children); i++ {
		updatingProblemTreeNodeFromID(&rt.Children[i])
	}
}

/*
	загрузить записанное дерево

Формат записи:
ID|ParentNode|autTreeID|situationTreeID|themeID
*/
func loadProblemTree() {

	//нулевой узел
	//[0]=&ProblemTree// все по нулям по умолчанию

	//ProblemTreeNodeFromID[0]=rt
	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/Problem_tree.txt")
	cunt := len(strArr)
	ProblemTreeNodeFromID = make([]*ProblemTreeNode, cunt) //задать сразу имеющиеся в файле число
	WriteProblemTreeNodeFromID(0, &ProblemTree)
	//просто проход по всем строкам файла подряд так что сначала идут дочки, потом - их родители
	for n := 0; n < cunt; n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		if len(strArr[n]) < 2 {
			panic("Сбой загрузки дерева проблем: [" + strconv.Itoa(n) + "] " + strArr[n])
			return
		}
		p := strings.Split(strArr[n], "|")
		id, _ := strconv.Atoi(p[0])
		parentID, _ := strconv.Atoi(p[1])
		autTreeID, _ := strconv.Atoi(p[2]) // PsyBaseautTreeID: -1 Плохое настроение, 0 Нормальное, 1 - хорошее настроение
		situationTreeID, _ := strconv.Atoi(p[3])
		themeID, _ := strconv.Atoi(p[4])
		purposeID, _ := strconv.Atoi(p[5])

		// новый узел с каждой строкой из файла
		var saveDoWritingFile = doWritingFile
		doWritingFile = false
		node, ok := ReadeProblemTreeNodeFromID(parentID)
		if !ok {
			continue
		}
		createNewProblemTreeNode(node, id, autTreeID, situationTreeID,
			themeID, purposeID, false)
		doWritingFile = saveDoWritingFile
	}
	return
}

// ID|ParentNode|autTreeID|situationTreeID|themeID
func SaveProblemTree() {
	if EvolushnStage < 4 { // только со стадии развития 4
		return
	}
	notAllowScanInTreeThisTime = true
	var out = ""
	cnt := len(ProblemTree.Children)
	for n := 0; n < cnt; n++ { // чтобы записывалось по порядку родителей
		out += getProblemTreeNode(&ProblemTree.Children[n])
	}

	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/Problem_tree.txt", out)
	notAllowScanInTreeThisTime = false
	return
}

// такой проход чтодбы дочерние узлы шли по порядку и всегда были бы родители
func getProblemTreeNode(wt *ProblemTreeNode) string {
	var out = "" //ID|ParentNode|autTreeID|situationTreeID|themeID
	//	if wt.ParentID>0 {
	out += strconv.Itoa(wt.ID) + "|"
	out += strconv.Itoa(wt.ParentID) + "|"
	out += strconv.Itoa(wt.autTreeID) + "|"
	out += strconv.Itoa(wt.situationTreeID) + "|"
	out += strconv.Itoa(wt.themeID) + "|"
	out += strconv.Itoa(wt.purposeID)
	out += "\r\n"
	//	}
	if wt.Children == nil { // конец
		return out
	}
	for n := 0; n < len(wt.Children); n++ {
		out += getProblemTreeNode(&wt.Children[n])
	}
	return out
}

// найти КОНЕЧНЫЙ узел по условиям
func FindProblemTreeNodeFromCondition(autTreeID int, situationTreeID int,
	themeID int, purposeID int) (int, *ProblemTreeNode) {
	/*
		ProblemTreeNodeMapCheck()
			for k, v := range ProblemTreeNodeFromID {
			if v==nil{continue}
		ProblemTreeNodeMapCheck()
				if v.autTreeID==autTreeID && v.situationTreeID==situationTreeID &&
					v.themeID==themeID {
					return k, v
				}
			}
	*/
	var id = 0
	var aut *ProblemTreeNode
	cnt := len(ProblemTree.Children)
	for n := 0; n < cnt; n++ {
		id, aut = checkProblemTree(&ProblemTree.Children[n], autTreeID, situationTreeID, themeID, purposeID)
		if id > 0 {
			return id, aut
		}
	}
	return 0, nil
}

func checkProblemTree(v *ProblemTreeNode, autTreeID int, situationTreeID int,
	themeID int, purposeID int) (int, *ProblemTreeNode) {
	var id = v.ID
	var aut = v

	// как только наткнется в предыдущих на такое услове - выдаст ID этой ветки
	if v.autTreeID == autTreeID && v.situationTreeID == situationTreeID && v.themeID == themeID && v.purposeID == purposeID {
		return v.ID, v
	}

	if v.Children == nil { // конец
		return 0, nil
	}
	for n := 0; n < len(v.Children); n++ {
		id, aut = checkProblemTree(&v.Children[n], autTreeID, situationTreeID, themeID, purposeID)
		if id > 0 {
			return id, aut
		}
	}
	return 0, nil //v.ID

}

// выдать массив узлов ветки по конечному ID, начиная с конечного к первому
func getcurrentProblemActivedNodes(lastID int) []*ProblemTreeNode {
	var nodws []*ProblemTreeNode
	//	node:=ProblemTreeNodeFromID[lastID]
	node, ok := ReadeProblemTreeNodeFromID(lastID)
	if ok {
		nodws = append(nodws, node)
		node = node.ParentNode
	}
	return nodws
}

/*
	создание иерархии АКТИВНЫХ образов контекстов условий и пусковых стимулов в виде ID образов в [4]int

создать последовательность уровней условий в виде массива  ID последовательности ID уровней
*/
func getProblemActiveConditionsArr(lev1 int, lev2 int, lev3 int, lev4 int) []int {
	arr := make([]int, 5)
	arr[0] = 0 //т.к. в conditionProblemFound сразу ost:=cond[1:] - пусть обрежет нулевой член
	arr[1] = lev1
	arr[2] = lev2
	arr[3] = lev3
	arr[4] = lev4
	return arr
}

// создание новой ветки, начиная с заданного узла
func addNewProblemBranchFromNodes(level int, cond []int, node *ProblemTreeNode) int {
	if node == nil {
		return 0
	}
	if level >= len(cond) {
		return node.ID
	}

	if cond[2] == 0 || cond[3] == 0 || cond[4] == 0 {
		return 0
	}

	var id = 0
	switch level {
	case 0: // нет ни одонго узла
		//		id,_=createNewProblemTreeNode(node,0,0,0,0,0,true)
	case 1:
		id, _ = createNewProblemTreeNode(node, 0, cond[1], 0, 0, 0, true)
	case 2:
		id, _ = createNewProblemTreeNode(node, 0, cond[1], cond[2], 0, 0, true)
	case 3:
		id, _ = createNewProblemTreeNode(node, 0, cond[1], cond[2], cond[3], 0, true)
	case 4:
		id, _ = createNewProblemTreeNode(node, 0, cond[1], cond[2], cond[3], cond[4], true)
	}
	level++
	node, ok := ReadeProblemTreeNodeFromID(id)
	if !ok {
		return 0
	}
	id = addNewProblemBranchFromNodes(level, cond, node)
	return id
}

// создание ветки, начиная с заданного узла fromID
func formingProblemBranch(fromID int, lastLevel int, condArr []int) int {
	lastNode := &ProblemTree // нит ни одного узла, нужно сделать всю ветку
	if fromID > 0 {
		//		lastNode = ProblemTreeNodeFromID[fromID]
		node, ok := ReadeProblemTreeNodeFromID(fromID)
		if ok {
			lastNode = node
		}
	}

	lastNodeID := addNewProblemBranchFromNodes(lastLevel, condArr, lastNode)
	if lastNodeID > 0 {
		//		 SaveProblemTree() // сохранять в общем порядке, при закрытии и по времени сохранения
	}
	return lastNodeID
}

// выдать массив узлов ветки по заданному ID узла
func getBrangeProblemTreeNodeArr(lastNodeId int) []*ProblemTreeNode {
	var nArr []*ProblemTreeNode
	//	node:=ProblemTreeNodeFromID[lastNodeId]
	node, ok := ReadeProblemTreeNodeFromID(lastNodeId)
	if !ok {
		return nil
	}
	for {
		if node == nil {
			break
		}
		nArr = append(nArr, node)
		node = node.ParentNode
	}
	return nArr
}

// выдать массив ID узлов ветки по заданному ID узла
func getBrangeProblemTreeNodeIdArr(lastNodeId int) []int {
	var nArr []int
	//node:=ProblemTreeNodeFromID[lastNodeId]
	node, ok := ReadeProblemTreeNodeFromID(lastNodeId)
	if !ok {
		return nil
	}
	for {
		if node == nil {
			break
		}
		nArr = append(nArr, node.ID)
		node = node.ParentNode
	}
	return nArr
}

// для Пульта - инфа о ветке Дерева
func ShowProblemTreeInfo(pID int) string {
	if pID == 0 {
		return "Нулевой ID..."
	}
	if len(ProblemTree.Children) == 0 { // еще нет никаких веток
		return "Еще нет Дерева проблем"
	}
	//	node:=ProblemTreeNodeFromID[pID]
	node, ok := ReadeProblemTreeNodeFromID(pID)
	if !ok {
		return "Нет ветки дерева с ID=" + strconv.Itoa(pID)
	}
	out := "<b>Ветка дерева проблем с ID=" + strconv.Itoa(pID) + "</b><br>"
	out += "ID дерева автоматизмов: <span onclick='show_atmzm_tree(" + strconv.Itoa(node.autTreeID) + ")'><b>" + strconv.Itoa(node.autTreeID) + "</b><span>"
	out += "ID дерева ситуации: <span onclick='show_unde_tree(" + strconv.Itoa(node.situationTreeID) + ")'><b>" + strconv.Itoa(node.situationTreeID) + "</b><span>"
	out += "ID Темы мышления: <span  onclick='get_situation(" + strconv.Itoa(node.themeID) + ")'><b>" + strconv.Itoa(node.themeID) + "</b><span>"
	out += "ID Цели мышления: <span onclick='get_purpose(" + strconv.Itoa(node.purposeID) + ")'><b>" + strconv.Itoa(node.purposeID) + "</b><span>"

	return out
}

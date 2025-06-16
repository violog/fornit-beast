/*  дерево фраз
вербальная иерархия распознавателей
Первый уровень дерева фраз может может заполняться любыми ID слов
Память о воспринятых фразах в текущем активном контексте (Vernike_detector.go): var MemoryDetectedArr []MemoryDetected
*/

package word_sensor

import (
	"BOT/lib"
	"regexp"
	"strconv"
	"strings"
)

// подошла очередь инициализации
func afterLoadPhraseArr() {
	loadPhraseTree()
	/*
		//%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
		SetNewPhraseTreeNode("повести и игра") //
		//%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
		SavePhraseTree()
	*/
	iniPraseRecognising()
	afetrInitPhraseTree()
	// для старых слов получить WordIdFormWord
	getWordIdFormWord()
	// deleteWordFromPhrase(11)
}

// дерево фраз, разбитых на слова, формат записи ID|ParentID|#|WordID
type PhraseTree struct {
	ID       int // id узла слова
	ParentID int // ID родителя
	WordID   int // одно  слово, м.б. пробелорм или любым символом

	Children   []PhraseTree // дочерние узлы (ветвление) НЕ АДРЕСА, А РЕАЛЬНЫЕ ОБЪЕКТЫ
	ParentNode *PhraseTree  // адрес родителя
}

// дерево фраз
var VernikePhraseTree PhraseTree

// карта поиска дерева фраз
// var PhraseTreeFromID = make(map[int]*PhraseTree)
var PhraseTreeFromID []*PhraseTree // сам массив
// var AFromID = make([]*aNode, len(strArr))//задать сразу имеющиеся в файле число при загрузке
// запись члена
func WritePhraseTreeFromID(index int, value *PhraseTree) {
	addPhraseTreeFromID(index)
	PhraseTreeFromID[index] = value
}
func addPhraseTreeFromID(index int) {
	if index >= len(PhraseTreeFromID) {
		newSlice := make([]*PhraseTree, index+1)
		copy(newSlice, PhraseTreeFromID)
		PhraseTreeFromID = newSlice
	}
}

// считывание члена
func ReadePhraseTreeFromID(index int) (*PhraseTree, bool) {
	if index < 0 {
		return nil, false
	}
	if index >= len(PhraseTreeFromID) || PhraseTreeFromID[index] == nil {
		return nil, false
	}
	return PhraseTreeFromID[index], true
}

/////////////////////////////////////////////////

// Последовательность wordID в ветке дерева нужно получать из GetWordArrFromPhraseID(PhraseID int)

///////////////////////////////////////////////////

/* ID фразы - по строке фразы - вместо GetExistsPraseID(str)
НО ВСЕ ФРАЗЫ ЗАПОМНАТЬ - КРУТОВАТО
Эта карта запоминается в отдельном файле потому как нет другой связки
(непросто получать при загрузке файла дерева фраз, точнее НУЖНО БУДЕТ ПОТРАТИТЬ ВРЕМЯ НА ВОССТАВЛЕНИЕ ФРАЗ ПО ID)

var PhraseIDFromPraseStr = make(map[string]int)
///// предотвращение ПАНИКИ КАРТ типа "concurrent map writes"
var PhraseIDFromPraseStrMapFlag=false //true- в карту идет запись
// перед чтением карты включать функцию уникального для данной карты арбитра MyMapCheck
func PhraseIDFromPraseStrMapCheck(){
	for PhraseIDFromPraseStrMapFlag{
		time.Sleep(10 * time.Microsecond)
	}
}
/////////////////////////////////
func savePhraseIDFromPraseStr() { // фраза|#|ID фразы
	var out = ""
	for str, id := range PhraseIDFromPraseStr {
		out+=str+"|#|"+id+"\r\n"

	}
	lib.WriteFileContent(lib.GetMainPathExeFile() + "/memory_reflex/phrase_tree.txt", out)
	return
}
////////////////////////
func loadPhraseIDFromPraseStr() {

}
*/

/////////////////////////////////////////////////

// Все ID фраз по wordID: в каких ID фраз содержится данное слово
// var PhraseTreeFromWordID = make(map[int][]int)
var PhraseTreeFromWordID [][]int // сам массив
// var AFromID = make([]*aNode, len(strArr))//задать сразу имеющиеся в файле число при загрузке
// запись члена
func WritePhraseTreeFromWordID(index int, value []int) {
	addPhraseTreeFromWordID(index)
	PhraseTreeFromWordID[index] = value
}
func addPhraseTreeFromWordID(index int) {
	if index >= len(PhraseTreeFromWordID) {
		newSlice := make([][]int, index+1)
		copy(newSlice, PhraseTreeFromWordID)
		PhraseTreeFromWordID = newSlice
	}
}

// считывание члена
func ReadePhraseTreeFromWordID(index int) ([]int, bool) {
	if index >= len(PhraseTreeFromWordID) || PhraseTreeFromWordID[index] == nil {
		return nil, false
	}
	return PhraseTreeFromWordID[index], true
}

/////////////////////////////////////////////////

// для обеспечения уникальности узлов:
/*  лишнее
type PhraseUnicum struct {
	ID int
	wordID int
}
var PhraseUnicumIdStr=make(map[PhraseUnicum]int)// для каждого сочетания  выдается ID узла
*/

var lastPhraseTreeID = 0 // конечный узел дерева фраз

// создать новый узел дерева фраз
func createNewNodePhraseTree(parent *PhraseTree, id int, wordID int) *PhraseTree {
	if parent == nil {
		return nil
	}
	//if wordID==0{ return nil }

	// после удаления слова - запрет на вставку новых слов до перезагрузки
	if blockingNewInsertWordAfterDeleted {
		lib.WritePultConsol("ПОСЛЕ УДАЛЕНИЯ СЛОВА - ЗАПРЕТ НА ВСТАВКУ НОВЫХ СЛОВ ДО ПЕРЕЗАГРУЗКИ")
		return nil
	}

	// notAllowScanInThisTime=true // запрет показа карты при обновлении
	if id == 0 {
		lastPhraseTreeID++
		id = lastPhraseTreeID
	} else {
		if lastPhraseTreeID < id {
			lastPhraseTreeID = id
		}
	}

	var newW PhraseTree
	newW.ID = id
	newW.ParentID = parent.ID
	newW.ParentNode = parent
	newW.WordID = wordID

	parent.Children = append(parent.Children, newW)
	// четко находим новый вставленный член (а не &parent.Children[count-1])
	var newTP *PhraseTree
	for i := 0; i < len(parent.Children); i++ {
		if parent.Children[i].ID == newW.ID {
			newTP = &parent.Children[i]
		}
	}
	//!!!! PhraseTreeFromID[new.ID]=new  т.к. append меняет длину массива, перетусовывая адреса, то нужно:
	scanAllTree(parent) // здесь потому, что при загрузке из файла нужно на лету получать адреса
	/*if newTP != nil {
		WordsArrFromPhraseID[newTP.ID] = append(WordsArrFromPhraseID[newTP.ID], newTP.WordID)
		PhraseTreeFromWordID[newTP.WordID] = append(PhraseTreeFromWordID[newTP.WordID], newTP)
	}*/
	// notAllowScanInThisTime=false

	return newTP
}

// корректируем адреса всех узлов
func scanAllTree(parent *PhraseTree) {
	updatingPhraseTreeFromID(parent)
}
func updatingPhraseTreeFromID(wt *PhraseTree) {
	if wt.ID > 0 {
		//		wt.ParentNode = PhraseTreeFromID[wt.ParentID] // wt.ParentNode адрес меняется из=за corretsParent(,
		node, ok := ReadePhraseTreeFromID(wt.ParentID)
		if ok {
			wt.ParentNode = node
			//		PhraseTreeFromID[wt.ID] = wt
			WritePhraseTreeFromID(wt.ID, wt)
		}
	}
	if wt.Children == nil {
		return
	} // конец ветки
	for i := 0; i < len(wt.Children); i++ {
		updatingPhraseTreeFromID(&wt.Children[i])
	}
}

// Загрузка дерева фраз
func loadPhraseTree() {
	initPhraseTree(&VernikePhraseTree)
	// initPhraseTree(&VernikePhraseTree)
	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_reflex/phrase_tree.txt")
	cunt := len(strArr)
	// просто проход по всем строкам файла подряд так что сначала идут дочки, потом - их родители
	for n := 0; n < cunt; n++ {
		if len(strArr[n]) < 2 {
			panic("Сбой загрузки дерева фраз: [" + strconv.Itoa(n) + "] " + strArr[n])
			return
		}
		p := strings.Split(strArr[n], "|#|")
		id, _ := strconv.Atoi(p[1])
		wordID := id
		if WordTreeFromID[wordID] == nil {
			continue
		} // нет такого узла дерева слов
		idP := strings.Split(p[0], "|")
		id, _ = strconv.Atoi(idP[0])
		parentID, _ := strconv.Atoi(idP[1])
		// новый узел с каждой строкой из файла
		node, ok := ReadePhraseTreeFromID(parentID)
		if ok {
			createNewNodePhraseTree(node, id, wordID)
		}
	}

	// заполнить PhraseTreeFromWordID
	finishScanAllTree()

	return
}

var curBrangeArr []int

func finishScanAllTree() {

	//	PhraseTreeFromWordID = make(map[int][]int)
	PhraseTreeFromWordID = nil

	curBrangeArr = nil
	curScanAllTree(&VernikePhraseTree)
}
func curScanAllTree(wt *PhraseTree) {
	if wt.ID > 0 {
		//		wt.ParentNode = PhraseTreeFromID[wt.ParentID] // wt.ParentNode адрес меняется из=за corretsParent(,
		node, ok := ReadePhraseTreeFromID(wt.ParentID)
		if ok {
			wt.ParentNode = node
			curBrangeArr = append(curBrangeArr, wt.WordID)
		}
	}
	if wt.Children == nil { // конец ветки
		/*if проверка не заблокирован ли уже {
			WordsArrFromPhraseIDmu.Lock()
			defer WordsArrFromPhraseIDmu.Unlock()
		}*/

		// перебрать все слова                       wordID - ЭТО НЕТ nodePhraseID !!!!!!!!!
		for i := 0; i < len(curBrangeArr); i++ {
			addPhraseTreeFromWordID(curBrangeArr[i])
			PhraseTreeFromWordID[curBrangeArr[i]] = append(PhraseTreeFromWordID[curBrangeArr[i]], wt.ID)
			/*
				val, ok := ReadePhraseTreeFromWordID(curBrangeArr[i])
					if (ok) {
						val = append(val, wt.ID)
					}*/
		}
		if len(curBrangeArr) > 1 {
			curBrangeArr = nil
		}

		/* ЗАЧЕМ ТУТ условие есои по любому внизу curBrangeArr=nil???
		if len(PhraseTreeFromWordID[wt.ID])>1{
			curBrangeArr=nil
		}
		*/
		curBrangeArr = nil
		return
	}
	for i := 0; i < len(wt.Children); i++ {
		//curBrangeArr=nil
		curScanAllTree(&wt.Children[i])
	}
}

////////////////////////////////////////////////////////////

// ///////////////////////////////////////////////////////
// создать первый, нулевой уровень дерева
func initPhraseTree(vt *PhraseTree) {
	// createNewNodePhraseTree(vt,0,0)
	vt.ID = 0
	vt.WordID = 0

	//	PhraseTreeFromID[vt.ID] = vt
	WritePhraseTreeFromID(vt.ID, vt)

	//updateWordTreeFromID()
	return
}

// Сохранить дерево фраз
// ID|ParentID|#|WordID
func SavePhraseTree() {
	var out = ""
	cnt := len(VernikePhraseTree.Children)
	for n := 0; n < cnt; n++ {
		out += getPtreeNode(&VernikePhraseTree.Children[n])
	}
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_reflex/phrase_tree.txt", out)
	return
}

// получить ветку дерева фраз в виде строки
func getPtreeNode(wt *PhraseTree) string {
	var out = ""

	out += strconv.Itoa(wt.ID) + "|"
	out += strconv.Itoa(wt.ParentID) + "|#|"
	out += strconv.Itoa(wt.WordID) + "\r\n"

	if wt.Children == nil {
		return out
	} // конец
	for n := 0; n < len(wt.Children); n++ {
		out += getPtreeNode(&wt.Children[n])
	}
	return out
}

/*
	вставка новой фразы со вставкой новых слов фразы,

так что фраза будет распознанна всегда.
*/
func SetNewPhraseTreeNode(phrace string) *WordTree {
	// чистим лишние пробелы
	rp := regexp.MustCompile("s+")
	phrace = rp.ReplaceAllString(phrace, " ")
	phrace = strings.TrimSpace(phrace)

	var wordsIDstr []int // строка (не)распознанных слов

	/* сначала добавляем слова в дерево слов, потом - всю фразу в дерево фраз
	   Делим фразу на слова (в строке нет других разделительных символов,
	т.к. они уже сработали при разделении на фразы).
	*/
	wArr := strings.Split(phrace, " ")
	for n := 0; n < len(wArr); n++ { // перебор отдельных слов
		curWord := strings.TrimSpace(wArr[n])
		if len(curWord) == 0 {
			return nil
		}

		id := SetNewWordTreeNode(curWord, false)
		// распознавание будет ВСЕГДА т.к. в случае новго слова оно вставляется в дерево слов тут же
		wordsIDstr = append(wordsIDstr, id)
	} //for n := 0; n < len(wArr); n++ { закончен проход отдельных слов
	// updateWordTreeFromID()// обновляем массив адресов узлов после всех append() родителей, меняющих адреса

	//  проход фразы
	// var needSave=false
	if len(wordsIDstr) > 0 {
		PhraseDetection(wordsIDstr, false)
		if DetectedUnicumPhraseID > 0 { // распознанная фраза
			CurrentPhrasesIDarr = append(CurrentPhrasesIDarr, DetectedUnicumPhraseID)

			//PhraseIDFromPraseStrMapFlag = true
			//PhraseIDFromPraseStr[phrace] = DetectedUnicumPhraseID
			//PhraseIDFromPraseStrMapFlag = false
		}
		// Запись недостающего в дерево фраз происходит в PhraseDetection(wordsIDstr)
	}
	// if needSave {
	//	savePhraseTree()
	// }
	return nil
}

// создание ветки фраз, начиная с заданного узла
func createPhraseTreeNodes(word []int, wt *PhraseTree) int {
	ost := word[1:]
	if len(ost) == 0 {
		return wt.ID
	}

	pn, ok := ReadePhraseTreeFromID(wt.ID)
	if ok {
		node := createNewNodePhraseTree(pn, 0, ost[0])
		pn2, ok2 := ReadePhraseTreeFromID(node.ID)
		if ok2 {
			id := createPhraseTreeNodes(ost, pn2)
			return id
		}
	}
	return 0
}

/*  Дерево эпизодической памяти.
для записи моторных правил для данных условий и фиксации значимостей образов стимулов.
Стало ясно, что моторная эпиз.память формируется в районе гиппокампа, а ментальная - где-то в лобных долях (там тоже есть циклы).
К дереву прилагается массив для фиксации исторической последовательности эпизодов.
В эпиз.памяти есть все для извлечения информации моделей понимания (в understanding_model.go) - системы значимостей образа в разных условиях.

Дерево легко использовать для
1) предказаний, что будет, если после данного стимула применить такой-то ответ
2) выбора ответов с позитивным эффектом

В дереве 5 узлов, до узла Action включительно:
NodeAID|NodeSID|NodePID|Trigger|Action|PARAMS
т.е. после Action идут его дочки типа PARAMS (числом от ни одной до сколько угодно)

значения PARAMS[] - это свойства эпизода: Effect|Count(уверенность)|stimulsEffect(значимость стимула)
Если при записи нового эпизода в дерево уже есть такая ветка, то эффект правила усредняется (func averageEffect)
Сила эффекта правил в зависимости от count отпределеяется func getWpower
Сила значимости стимула в зависимости от count отпределеяется func getОpower

В файл пишется:
ID|ParentID|BaseID|EmotionID|NodePID|Trigger|Action#Effect|Count|stimulsEffect
Нет смысла выносить PARAMS[] в отдельныю структуру со связью по ID
т.к. память это не сэкономит, но приведет к лишнему геммору.

Дерево не активируется по каким-то стимулам, в него просто делается запись Правил
и используется для получения информации.

ПРИНЦИП ПОИСКА НАИБОЛЕЕ СТАТИСТИЧЕСКИ ВЕРНОГО ДЕЙСТВИЯ (GPT):
Eсть историческая лента прошлых кадров. По дереву ищутся кадры, максимально похожие на текущую ситуацию.
Начинается от них просмотр далее чтобы найти действие с позитивным результатом.
Каждая цепоска проходится или до позитива (он может быть или сразу или через несколько негативов)
или до минутного перерыва (считаешь, что опят на этом кончился неудачно).
Все найденные позитивные цепочки сравниваются и выбирается с самым большим позитивом.
Начинается действие первого шага, пусть негативного, но скоро же будет позитив.
Так реализуется принцип GPT.

ПРИМЕРЫ ИСПОЛЬЗОВАНИЯ:
1. saveNewEpisodic(actionsImageID, autmtzm.ActionsImageID, effect, averageEffect) - новый кадр эпизодической памяти, сохраняющий Правил
2. saveNewEpisodic(answerID, curAct, 100, averageEffect) - учительское правило
3. saveNewEpisodic(curActions.ID, 0, 0, 0) - записать пустое Правило (нет Action)

Т.к. в дереве эп.памяти не может быть дубликатов, то каждый узел уникален, а время записи памяти фиксируется в истории EpisodicHistoryArr

*/
/////////////////////////////////////////////////////////////////////

package psychic

import (
	"BOT/lib"
	"strconv"
	"strings"
)

///////////////////////////////////////////

type EpisodicTreeNode struct { // узел дерева Эпизодической памяти
	ID int
	// Узлы дерева:
	BaseID    int // 1 - Похо, 2 - Норма, 3 - Хорошо
	EmotionID int // эмоция
	/* субъективный узел активной ветки дерева проблем (последний активировавшийся)
	По нему можно найти:
	node,ok:=ReadeProblemTreeNodeFromID()
	if ok{
	NodeAID:=node.autTreeID //конечный узел активной ветки дерева моторных автоматизмов
	NodeSID:=node.situationTreeID //конечный узел активной ветки дерева ситуации
	}*/
	NodePID int

	// Cтимул (полный образ Стимула с Пульта) = ActionsImage.ID
	Trigger int

	// Образ ответных действий - ActionsImage.ID
	Action int
	// Свойства данного узла - массив PARAMS:
	PARAMS []int
	/* [3] иммет три значения:
	Effect int // эффект от действий: (меньше 0) или 0 или (больше 0) Effect учительского Правила ==100.
	Count int // всегда>0 - число усреднений Effect  (используется Count) - уверенность в Правиле и значимости стимула.
	stimulsEffect int // значимость стимула усредняется с каждой новой записью в дерево (используется Count).
	*/

	//Связь узлов дерева:
	Children   []EpisodicTreeNode // дочерние узлы (ветвление) НЕ АДРЕСА, А РЕАЛЬНЫЕ ОБЪЕКТЫ
	ParentID   int                // ID родителя
	ParentNode *EpisodicTreeNode  // адрес родителя
}

// ////////////////////////////////////////////////////
var EpisodicTree EpisodicTreeNode

// var EpisodicTreeNodeFromID=make(map[int]*EpisodicTreeNode)
var EpisodicTreeNodeFromID []*EpisodicTreeNode // сам массив
// запись члена
func WriteEpisodicTreeNodeFromID(index int, value *EpisodicTreeNode) {
	addEpisodicTreeNodeFromID(index)
	EpisodicTreeNodeFromID[index] = value
}
func addEpisodicTreeNodeFromID(index int) {
	if index >= len(EpisodicTreeNodeFromID) {
		newSlice := make([]*EpisodicTreeNode, index+1)
		copy(newSlice, EpisodicTreeNodeFromID)
		EpisodicTreeNodeFromID = newSlice
	}
}

// считывание члена
func ReadeEpisodicTreeNodeFromID(index int) (*EpisodicTreeNode, bool) {
	if index < 0 {
		return nil, false
	}
	if index >= len(EpisodicTreeNodeFromID) || EpisodicTreeNodeFromID[index] == nil {
		return nil, false
	}
	return EpisodicTreeNodeFromID[index], true
}

//////////////////////////////////////////////////////////////////////

/*
Связь узлов историческая - ID конечных узлов веток EpisodicTree подряд по порядку возникновения новых.
Записывается в файл "/memory_psy/episodic_history.txt" - сплошная строка EpisodicTreeNode.ID с разделителем "|"
В случе добавления newPARAMS в существующий EpisodicTreeNode пишется его EpisodicTreeNode.ID
так что в EpisodicHistoryArr могут появляться несколько EpisodicTreeNode.ID.

Пустые кадры (с ID == -1) разделяют цепочки памяти одной темы.

Поиск в EpisodicHistoryArr нужного EpisodicTreeNode.ID
При оптимизации во сне порядок сохраняется, но могут вырезаться отдельные составляющие.
При вырезании из EpisodicHistoryArr пишется в файл episodic_history.txt новая строка
и т.о. EpisodicHistoryArr []int остается с непрерывной цепочкой индексов.

Найти в массиве EpisodicHistoryArr[] индексы всех членов по значению:
iArr:=lib.FindIndexes(EpisodicHistoryArr,ID)

Запись в файле: ID1,LifeTime1|ID2,LifeTime2 ...
*/
type History struct {
	ID       int // ID EpisodicTreeNode
	LifeTime int // время появления узла в пульсах от рождения
}

var EpisodicHistoryArr []History

///////////////////////////////////////////////////////

/*
	Элемент отдельного Правила

для вывода целевой цепочки Правил и т.п.
*/
type Rule struct { //
	Trigger    int // Cтимул (действий оператора)
	Action     int // образ ответных действий - всегда ActionsImage
	Effect     int // эффект от действий: -1 или 0 или 1
	Count      int // всегда>0 - число усреднений Effect - уверенность в Правиле
	Importence int // значимость стимула (Trigger если Effect<100 или Action если Effect==100)
}

///////////////////////////////////////////////////////

// ////////// ЗАПИСАТЬ В ДЕРЕВО НОВЫЙ ЭПИЗОД
var lastEpisodicMemID = 0

/*
	если нужно испольховать предыдущее значение условий (до стимула от оператора) - для прямого правила func fixEpizMemoryRules

т.к. послед действия оператора может сильно измениться базовый контекст и эмоция.
*/
var usedOldCondition = false

func saveNewEpisodic(Trigger int, Action int, Effect int, stimulsEffect int) {

	//	NodeAID=65;NodeSID=38;NodePID=8;Trigger=1;Action=1;Effect=4;stimulsEffect=1;
	//	NodeAID=65;NodeSID=38;NodePID=8;Trigger=1;Action=2;Effect=4;stimulsEffect=1;

	var BaseID = 0
	var EmotionID = 0
	var NodePID = 0                                                // ID проблемы
	if usedOldCondition && oldDetectedActiveLastProblemNodID > 0 { // только если уже есть предыдущие значения!
		BaseID = oldCommonBadNormalWell
		EmotionID = oldCurrentEmotionReception.ID
		NodePID = oldDetectedActiveLastProblemNodID // ID проблемы
	} else {
		BaseID = CurrentCommonBadNormalWell
		EmotionID = CurrentEmotionReception.ID
		NodePID = detectedActiveLastProblemNodID // ID проблемы
	}

	usedOldCondition = false

	//	getEpisodesArrFromConditions(0)

	// !!!! после провокации func infoFunc31 нет Стимула!
	//if Action == 0 { // может быть пустое правило!
	//	return
	//}
	params := []int{Effect, 1, stimulsEffect}
	var condArr = []int{BaseID, EmotionID, NodePID, Trigger, Action}

	//проверяются дубликаты
	idOld, nodeOld := checkEpisodicBranchFromCondition(BaseID, EmotionID, NodePID, Trigger, Action)
	if idOld > 0 {
		EpisodicHistoryArr = append(EpisodicHistoryArr, History{idOld, LifeTime})
		// усреднить эффект
		averageEffect(nodeOld, Effect, stimulsEffect)
		params = nodeOld.PARAMS
	}

	if Effect != 100 && params != nil {
		if params[0] > 10 {
			params[0] = 10
		}
		if params[0] < -10 {
			params[0] = -10
		}
	}

	lastEpisodicMemID = addEpisodicFromNodeIDsToBrange(0, 0, condArr, params)
	//lastEpisodicMemID - ID последнего кадра

	if idOld == 0 && lastEpisodicMemID >= 0 { // не давать записывать кадры прерывания цепочки: lastEpisodicMemID ==-1
		EpisodicHistoryArr = append(EpisodicHistoryArr, History{lastEpisodicMemID, LifeTime})
	}

	return
}

// ////////////////////////////////////////
// усреднить эффект в PARAMS узла
func averageEffect(node *EpisodicTreeNode, effect int, stimulsEffect int) {
	if node.PARAMS == nil {
		return
	}
	// вес нового эффекта
	count := node.PARAMS[1] + 1
	if count == 0 {
		lib.TodoPanic("func averageEffect НУЛЕВОЕ ЗНАЧЕНИЕ count")
		return
	}
	if effect != 100 {
		// усреднить эффект правила
		w := int(((node.PARAMS[0] * (count - 1)) + effect) / count)
		if w > 10 {
			w = 10
		}
		if w < -10 {
			w = -10
		}
		node.PARAMS[0] = w
	}
	// усреднить значимость стимула
	w := int(((node.PARAMS[2] * (count - 1)) + stimulsEffect) / count)
	if w > 10 {
		w = 10
	}
	if w < -10 {
		w = -10
	}
	node.PARAMS[2] = w
}

///////////////////////////////////////////////////////////////////////////////////

/*
	Для создания нового узла ветки дерева

Возвращает ID узла и указатель на него - в любом случае, создан ли новый узел или найден такой же уже имеющийся.
*/
var lastIDNodeTree = 0 // счетчик ID узлов, хранит ID последнего созданного узла
func createEpisodicNodeTree(parent *EpisodicTreeNode, id int, BaseID int, EmotionID int, NodePID int, Trigger int, Action int, PARAMS []int) (int, *EpisodicTreeNode) {
	if parent == nil {
		return 0, nil
	}

	if id == 0 {
		lastIDNodeTree++
		id = lastIDNodeTree
	} else {
		if lastIDNodeTree < id {
			lastIDNodeTree = id
		}
	}

	var node EpisodicTreeNode
	node.ID = id
	node.ParentNode = parent
	node.ParentID = parent.ID
	node.BaseID = BaseID
	node.EmotionID = EmotionID
	node.NodePID = NodePID
	node.Trigger = Trigger
	node.Action = Action
	node.PARAMS = PARAMS

	parent.Children = append(parent.Children, node)
	// четко находим новый вставленный член (а не &parent.Children[count-1])
	var newN *EpisodicTreeNode
	for i := 0; i < len(parent.Children); i++ {
		if parent.Children[i].ID == node.ID {
			newN = &parent.Children[i]
			//		newN.Children=append(newN.Children, node)
		}
	}

	//EpisodicTreeNodeFromID[node.ID]=&node
	WriteEpisodicTreeNodeFromID(node.ID, &node)

	//WriteEpisodicTreeNodeFromID(parent.ID, parent)

	// т.к. append меняет длину массива, перетусовывая адреса, то нужно обновить адреса в EpisodicTreeNodeFromID:
	// ДЕЛАТЬ ТОЛЬКО ДЛЯ ДЕРЕВЬЕВ!
	updatingEpisodicTreeNodeFromID(parent) // здесь потому, что при загрузке из файла нужно на лету получать адреса
	return id, newN
}

//////////////////////////////////////////////////////////

// ////////////////////////////////
func checkEpisodicBranchFromCondition(BaseID int, EmotionID int, NodePID int, Trigger int, Action int) (int, *EpisodicTreeNode) {
	cond := []int{BaseID, EmotionID, NodePID, Trigger, Action}
	//maxLev:=5
	id, level := findEpisodicBrange(0, cond, &EpisodicTree)
	if id > 0 {
		// найти уровень lev0 ненулевых параметров
		lev0 := getTrueLevel(cond)
		if lev0 > level {
			return 0, nil
		}
		//if maxLevel==maxLev{// есть такая ветка, просто возвращаем ее значения

		//node,ok:=EpisodicTreeNodeFromID[id]
		node, ok := ReadeEpisodicTreeNodeFromID(id)
		if ok {
			return id, node
		} else { // такого не должно быть
			lib.TodoPanic("В func FindEpisodicTreeNodeFromCondition должно быть значение карты EpisodicTreeNodeFromID.") //вызвать панику
		}

		//}else{// нужно дорастить ветку
		//	return 0,nil
		//}
	} else { // ID==0 вообще нет совпадений, нужно наращивать с основы
		return 0, nil
	}

	return 0, nil
}

// ////////
// найти уровень lev0 ненулевых параметров из массива условий
func getTrueLevel(cond []int) int {
	lev0 := 0
	for i := 0; i < len(cond); i++ {
		if cond[i] == 0 {
			break
		}
		lev0++
	}
	return lev0
}

// //////////
// рекурсивно корректируем адреса всех узлов
// ДЕЛАТЬ ТОЛЬКО ДЛЯ ДЕРЕВЬЕВ!
func updatingEpisodicTreeNodeFromID(rt *EpisodicTreeNode) {
	if rt.ID > 0 {
		rt.ParentNode = EpisodicTreeNodeFromID[rt.ParentID] // wr.ParentNode адрес меняется из=за corretsParent(,
		_, ok := ReadeEpisodicTreeNodeFromID(rt.ParentID)
		if ok {
			WriteEpisodicTreeNodeFromID(rt.ID, rt)
		}
	}
	if rt.Children == nil { // конец ветки
		return
	}
	for i := 0; i < len(rt.Children); i++ {
		updatingEpisodicTreeNodeFromID(&rt.Children[i])
	}
}

///////////////////////////////////////////////////////////

// //////////////////////////////////////////////
// получить массив PARAMS[3] последнего узла дерева
// Только узел c ActionID имеет PARAMS[][]
func getEpisodicPARAMS(nodeID int) []int {

	//var actionNode,ok=EpisodicTreeNodeFromID[nodeID]
	actionNode, ok := ReadeEpisodicTreeNodeFromID(nodeID)
	if !ok {
		return nil
	}

	if actionNode.Action == 0 {
		return nil
	}

	return actionNode.PARAMS //[]int{}
}

////

/////////////////////////////////////////////////////////////////////////

/*
	Доращивание ветки, начиная с заданного узла fromID по массиву всех значений для ветки

newPARAMS - добавить к ветке, если она существует новый массив Effect|LifeTime
Возвращает ID конечного узла ветки.
*/
func addEpisodicFromNodeIDsToBrange(fromID int, lastLevel int, condArr []int, newPARAMS []int) int {

	var bNode *EpisodicTreeNode
	if fromID > 0 {
		//node, ok := EpisodicTreeNodeFromID[fromID]
		node, ok := ReadeEpisodicTreeNodeFromID(fromID)
		if !ok {
			return 0
		}
		bNode = node
	} else { // не определен начальный узел, хотя он может и быть
		bNode = &EpisodicTree
	}

	lastNodeID, _ := addEpisodicNodesToBrange(bNode, lastLevel, condArr, newPARAMS)

	return lastNodeID
}

/*
рекурсивно создать все недостающие уровни ветки по значениям condArr []int
*/
func addEpisodicNodesToBrange(fromNode *EpisodicTreeNode, level int, condArr []int, newPARAMS []int) (int, *EpisodicTreeNode) {
	if fromNode == nil {
		return 0, nil
	}
	if level >= len(condArr) {
		return fromNode.ID, fromNode
	}
	//vArr := make([]int, len(condArr))// все vArr[n] имеют нулевые значения
	vArr := []int{0, 0, 0, 0, 0}
	for i := 0; i <= level; i++ { // заполнить vArr до текущего уровня включительно, остальные - оставить нулями
		vArr[i] = condArr[i]
	}
	// с каждым уровнем vArr[n] добавляется новое значение для создания узла ветки следующего уровня.
	// если такой узел есть (checkUnicum==true), то просто вернет его node
	var node *EpisodicTreeNode
	var pars []int

	//проверяются дубликаты потому, что новый узел нужно делать только если такого еще нет!!!
	idOld, nodeOld := checkEpisodicBranchFromCondition(vArr[0], vArr[1], vArr[2], vArr[3], vArr[4])

	//node.PARAMS=append(node.PARAMS,pars)
	if idOld > 0 {
		node = nodeOld
		if level == 4 && newPARAMS != nil {
			// усреднить эффект
			averageEffect(node, newPARAMS[0], newPARAMS[2])
		}
	} else { // НОВЫЙ
		if level == 4 {
			pars = newPARAMS
		}

		if pars != nil && pars[0] != 100 {
			if pars[0] > 10 {
				pars[0] = 10
			}
			if pars[0] < -10 {
				pars[0] = -10
			}
		}

		_, node = createEpisodicNodeTree(fromNode, 0, vArr[0], vArr[1], vArr[2], vArr[3], vArr[4], pars)
	}

	level++
	if level >= len(condArr) {
		return node.ID, node
	}
	id, node := addEpisodicNodesToBrange(node, level, condArr, newPARAMS)
	return id, node
}

///////////////////////////////////////

/*
	поиск конечного узла ветки (lastBrangeID) дерева (root - ID начального узла == 0) по массиву ID узлов ветки (cond []int)

Поиск начинается с первого узла (по значению cond[0]).
Если первый узел найден, то findSampleBrange вызывается рекурсивно для поиска следующего узла
и так далее, пока не будет найден удел последнего члена cond.
Если для текущего cond[level] узел в редеве не существует, то в дереве доращивается ветка по значениям cond []int.

Возвращает ID последней найденной ветки и ее уровень, если конечный узел не найден (если найден, то len(cond)-1).
Если вернула уровень меньше, чем максимальное число уровней ветки,
то ветка доращивается func addNodesToBrange

Ищет до узла Action включительно, а к этому узлу могут быть прикреплены несколько PARAMS
*/
func findEpisodicBrange(level int, cond []int, root *EpisodicTreeNode) (int, int) {

	// Обработка случая когда мы достигли конца дерева
	if cond == nil || len(cond) <= 0 || level >= len(cond) {
		return root.ID, len(cond)
	}

	// Поиск узла с ID из списка cond в дочерних узлах текущего узла
	for _, child := range root.Children {
		if isEquivalentCondition(level, &child, cond) {
			// Если узел найден, продолжим рекурсивно искать в нем далее
			id, lev := findEpisodicBrange(level+1, cond, &child)
			if id == 0 {
				/*Eсли условие cond []int содержит значение 0 для какого-то уровня, то на этом уровне поиск заканчивается.
				if cond[lev]==0{// считается нормальный поиск по неполному условию

				}*/
				return child.ID, lev // поиск закончен на child.ID
				//return 0,lev // не найдено совпадение на данном уровне
			}

			return id, lev
		}
		//  пусть перебирает дочки! return 0,level // не найдено совпадение на данном уровне
	}

	// не найден узел на данном уровне
	return 0, level
}

// //////////////////////////////////////////////////
func isEquivalentCondition(level int, node *EpisodicTreeNode, cond []int) bool {
	// массив значений по уровням - в зависимости от числа уровней в ветке, прописывается вручную
	nArr := []int{node.BaseID, node.EmotionID, node.NodePID, node.Trigger, node.Action}

	for i := 0; i < len(cond) && i <= level; i++ {
		if cond[i] != nArr[i] {
			return false
		}
	}
	return true
}

/////////////////////////////////////////////////////////

/* загрузить записанное дерево
 */
func loadEpisodicTree() {
	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/episodic_tree.txt")
	initEpisodicTree(strArr)

	str := lib.ReadFileContent(lib.GetMainPathExeFile() + "/memory_psy/episodic_history.txt")

	sArr := strings.Split(str, "|")
	for n := 0; n < len(sArr); n++ {
		if len(sArr[n]) == 0 {
			break
		}
		il := strings.Split(sArr[n], ",")
		id, _ := strconv.Atoi(il[0])
		lifeTime, _ := strconv.Atoi(il[1])
		EpisodicHistoryArr = append(EpisodicHistoryArr, History{id, lifeTime})
	}

	// ОТЛАДКА, ТЕСТИРОВАНИЕ
	// saveNewEpisodic(67,38,107,1,3,3)
	//	saveNewEpisodic(200,39,100,1,3,3)

	// getEpisodesArrFromConditions(0,curStimulImageID,0)
	// getEpisodesArrFromConditions(0,1,0)
	// getEpisodesArrFromConditions(1,1,0)

	// тестирование getTargetEpisodicStrIdArr()
	testing := false
	//	testing=true
	if testing { // здесь еще не активированные деревья
		detectedActiveLastNodID = 67
		detectedActiveLastUnderstandingNodID = 38
		detectedActiveLastProblemNodID = 8
		curActiveActionsID = 1
		getTargetEpisodicStrIdArr(curActiveActionsID, getLimitCountEM())
		if targetEpisodicStrIdArr != nil {

		}
		typeRule := 1 //typeRule - 1-искать только прямые Правила, 2-искать только учительские Правила, 3- искать все виды Правил
		rule := getSingleBestRule(typeRule, curActiveActionsID)
		if rule.Action > 0 {

		}
	}
	setInterruptionEpisosde() // обязательно в конце втыкаем пустой кадр!!! Иначе после просыпания начнет дописывать новые кадры к старым цепочкам
	return
}
func initEpisodicTree(strArr []string) {

	cunt := len(strArr)
	EpisodicTreeNodeFromID = make([]*EpisodicTreeNode, cunt) // задать сразу имеющиеся в файле число при загрузке из файла
	lastIDNodeTree = 1                                       // нужно, чтобы дерево памяти началось с ID==2 (lastIDNodeTree++ ==2), иначе оно не загрузится. см ReadeEpisodicTreeNodeFromID()
	// определить нулевой узел
	//EpisodicTreeNodeFromID[0]=&EpisodicTree// все по нулям по умолчанию
	WriteEpisodicTreeNodeFromID(0, &EpisodicTree)

	//просто проход по всем строкам файла подряд так что сначала идут дочки, потом - их родители
	for n := 0; n < cunt; n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		if len(strArr[n]) < 2 {
			panic("Сбой загрузки дерева: [" + strconv.Itoa(n) + "] " + strArr[n])
			return
		}

		par := strings.Split(strArr[n], "#")
		p := strings.Split(par[0], "|")
		var parArr []string
		if len(par) > 1 {
			parArr = strings.Split(par[1], "|")
		}

		id, _ := strconv.Atoi(p[0])
		parentID, _ := strconv.Atoi(p[1])
		BaseID, _ := strconv.Atoi(p[2])
		EmotionID, _ := strconv.Atoi(p[3])
		NodePID, _ := strconv.Atoi(p[4])
		Trigger, _ := strconv.Atoi(p[5])
		Action, _ := strconv.Atoi(p[6])

		// считывание PARAMS
		var params []int
		if Action > 0 {
			idP, _ := strconv.Atoi(parArr[0])
			count, _ := strconv.Atoi(parArr[1])
			kind, _ := strconv.Atoi(parArr[2])
			params = []int{idP, count, kind}
		}

		parentNode, ok := ReadeEpisodicTreeNodeFromID(parentID)
		if ok {

			// Создание нового узла дерева
			node := EpisodicTreeNode{
				ID:         id,
				BaseID:     BaseID,
				EmotionID:  EmotionID,
				NodePID:    NodePID,
				Trigger:    Trigger,
				Action:     Action,
				Children:   []EpisodicTreeNode{},
				ParentID:   parentID,
				ParentNode: parentNode,
			}
			node.PARAMS = params

			node.ParentID = parentNode.ID
			node.ParentNode = parentNode
			parentNode.Children = append(parentNode.Children, node)

			// Добавление узла в массив по ID
			WriteEpisodicTreeNodeFromID(id, &node)
			lastIDNodeTree++

			// т.к. append меняет длину массива, перетусовывая адреса, то нужно обновить адреса в EpisodicTreeNodeFromID:
			// ДЕЛАТЬ ТОЛЬКО ДЛЯ ДЕРЕВЬЕВ!
			updatingEpisodicTreeNodeFromID(parentNode) // здесь потому, что при загрузке из файла нужно на лету получать адреса
		}
		/*
			// новый узел с каждой строкой из файла
			//parent,ok:=EpisodicTreeNodeFromID[parentID]
			parent,ok:=ReadeEpisodicTreeNodeFromID(parentID)
			if ok {
				// false - не проверять наличие ветки т.к. еще нет веток, они только глузятся
				id,newNodw:=createEpisodicNodeTree(parent, id, BaseID, EmotionID, NodePID, Trigger, Action, params)
				if id>0 && newNodw!=nil{

					if id>0{
						id=id
					}
				}
			}*/
	}
	return
}

/////////////////////////////////////

func SaveEpisodicTree() {
	var out = ""
	cnt := len(EpisodicTree.Children)
	for n := 0; n < cnt; n++ { // чтобы записывалось по порядку родителей
		out += getEpisodicTreeNode(&EpisodicTree.Children[n])
	}
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/episodic_tree.txt", out)

	// запись стории эпизодов EpisodicHistoryArr []int
	out = ""
	for n := 0; n < len(EpisodicHistoryArr); n++ {
		out += strconv.Itoa(EpisodicHistoryArr[n].ID) + "," + strconv.Itoa(EpisodicHistoryArr[n].LifeTime) + "|"
	}
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/episodic_history.txt", out)

	return
}
func getEpisodicTreeNode(wt *EpisodicTreeNode) string {
	var out = ""
	//	if wt.ParentID>0 {
	out += strconv.Itoa(wt.ID) + "|"
	out += strconv.Itoa(wt.ParentID) + "|"
	out += strconv.Itoa(wt.BaseID) + "|"
	out += strconv.Itoa(wt.EmotionID) + "|"
	out += strconv.Itoa(wt.NodePID) + "|"
	out += strconv.Itoa(wt.Trigger) + "|"
	out += strconv.Itoa(wt.Action) //+ "#"

	// записать массивы PARAMS
	parsArr := getEpisodicPARAMS(wt.ID)
	if parsArr != nil {
		out += "#"
		out += strconv.Itoa(parsArr[0]) + "|"
		out += strconv.Itoa(parsArr[1]) + "|"
		out += strconv.Itoa(parsArr[2]) + "|"
	}

	out += "\r\n"
	//	}
	if wt.Children == nil { // конец
		return out
	}
	for n := 0; n < len(wt.Children); n++ {
		out += getEpisodicTreeNode(&wt.Children[n])
	}
	return out
}

/////////////////////////////////////

// Очистить эпизодическую память
func ClianEpisodicMemory() {
	lib.WriteFileContentExactly(lib.GetMainPathExeFile()+"/memory_psy/episodic_tree.txt", "")
	lib.WriteFileContentExactly(lib.GetMainPathExeFile()+"/memory_psy/episodic_history.txt", "")

	EpisodicHistoryArr = nil
	lastIDNodeTree = 1 // нельзя 0, не сработает ReadeEpisodicTreeNodeFromID()!!!
	WriteEpisodicTreeNodeFromID(0, &EpisodicTree)
}

/////////////////////////////////////

/*
	вставить кадр прерывания цепочки (темы)

Означает, что цепочка кадров памяти данной темы закончиена (fornit.ru/67675).
Вставляется по щелчку по плащке ожидания ответа (func StopWaitingWeriodFromOperator())
TODO вставлять при смене темы и произвольно.
*/
func setInterruptionEpisosde() {
	// первым не включать
	if len(EpisodicHistoryArr) == 0 {
		return
	}

	// если только что уже был пустой кадр, то не вставлять лишний
	lastI := len(EpisodicHistoryArr) - 1
	if EpisodicHistoryArr[lastI].ID == -1 {
		return
	}
	// пустой кадр с ID=-1
	EpisodicHistoryArr = append(EpisodicHistoryArr, History{-1, LifeTime})
}

///////////////////////////////////////////////////////////

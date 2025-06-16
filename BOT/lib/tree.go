/* Пример кода для пддержки деревьев
типа с фиксированным числом уровней - наиболее физиологичный тип.

!!! Для деревьев с переменной длиной веток, как WordTree или PhraseTree нужно адаптировать функции
т.к. там в структурах узлов - значение только одного текущего уровня. Зато там максимально компактная запись в файл.

Для использования кода просто "sample" везде менять на название вида дерева.

Формат записи примера дерева sampleTree в файл:
ID|ParentNode|levVal1|levVal2|levVal3|levVal4

Для тестирования раскомментируй tryExperiment() в func init() !
*/

package lib

import (
	"strconv"
	"strings"
	//"BOT/lib"
)

///////////////////////////////////////////////

func init() {

	// tryExperiment() //  проверка работы дерева в разных ситуациях

	return
}

///////////////////////////////////////////

type sampleNode struct { // узел дерева
	ID int
	/* в ветке узелы получают все значения предыдущих, кроме пока еще не достигших даноого уровня.
	   т.е. если в каждом узле сохраняется инфа о всех значениях узлов предыдущих уровней вентки.
	*/
	levVal1 int // значение узла первого уровня (всегда ID образа)
	levVal2 int // значение узла второго уровня (всегда ID образа)
	levVal3 int // значение узла третьего уровня (всегда ID образа)
	levVal4 int // значение узла четвертого уровня (всегда ID образа)

	Children   []sampleNode // дочерние узлы (ветвление) НЕ АДРЕСА, А РЕАЛЬНЫЕ ОБЪЕКТЫ
	ParentID   int          // ID родителя
	ParentNode *sampleNode  // адрес родителя
}

var sampleTree sampleNode

var sampleTreeFromID = make(map[int]*sampleNode)

// /// предотвращение ПАНИКИ КАРТ типа "concurrent map writes"
var sampleTreeMapFlag = false //true- в карту идет запись
var MapGwardSampleTreeFromID = RegNewMapGuard()

///////////////////////////////////////

/////////////////////////////////////////// функции поддержки
/* Для создания нового узла ветки дерева
Возвращает ID узла и указатель на него - в любом случае, создан ли новый узел или найден такой же уже имеющийся.
*/
var lastIDNodeTree = 0 // счетчик ID узлов, хранит ID хранит ID последнего созданного узла
func createNodeTree(parent *sampleNode, id int, levVal1 int, levVal2 int, levVal3 int, levVal4 int, checkUnicum bool) (int, *sampleNode) {
	if parent == nil {
		return 0, nil
	}

	// если есть такой узел, то не создавать новый
	if checkUnicum {
		idOld, nodeOld := checkSampleBranchFromCondition(levVal1, levVal2, levVal3, levVal4)
		if idOld > 0 {
			return idOld, nodeOld
		}
	}

	if id == 0 {
		lastIDNodeTree++
		id = lastIDNodeTree
	} else {
		//		newW.ID=id
		if lastIDNodeTree < id {
			lastIDNodeTree = id
		}
	}

	var node sampleNode
	node.ID = id
	node.ParentNode = parent
	node.ParentID = parent.ID
	node.levVal1 = levVal1
	node.levVal2 = levVal2
	node.levVal3 = levVal3
	node.levVal4 = levVal4

	parent.Children = append(parent.Children, node)
	// четко находим новый вставленный член (а не &parent.Children[count-1])
	var newN *sampleNode
	for i := 0; i < len(parent.Children); i++ {
		if parent.Children[i].ID == node.ID {
			newN = &parent.Children[i]
		}
	}

	sampleTreeMapFlag = true
	sampleTreeFromID[node.ID] = &node
	sampleTreeMapFlag = false

	// т.к. append меняет длину массива, перетусовывая адреса, то нужно обновить адреса в AutomatizmTreeFromID:
	// ЭТО НУЖНО ДЕЛАТЬ ТОЛЬКО ДЛЯ ДЕРЕВЬЕВ!
	updatingSampleTreeFromID(parent) // здесь потому, что при загрузке из файла нужно на лету получать адреса

	return id, newN
}

// ////////////////////////////////
func checkSampleBranchFromCondition(levVal1 int, levVal2 int, levVal3 int, levVal4 int) (int, *sampleNode) {
	cond := []int{levVal1, levVal2, levVal3, levVal4}
	maxLev := 4
	/*
		maxLev:=0// максимальный уровень в условии cond
		for i := 0; i < len(cond); i++ {
			if cond[i]==0{
				break
			}
			maxLev++
		}*/
	id, maxLevel := findSampleBrange(0, cond, &sampleTree)
	if id > 0 {
		if maxLevel == maxLev { // есть такая ветка, просто возвращаем ее значения
			MapCheckBlock(MapGwardSampleTreeFromID)
			node, ok := sampleTreeFromID[id]
			if ok {
				return id, node
			} else { // такого не должно быть
				TodoPanic("В func FindSampleTreeNodeFromCondition должно быть значение карты sampleTreeFromID.") //вызвать панику
			}
			MapFree(MapGwardSampleTreeFromID)
		} else { // нужно дорастить ветку
			return 0, nil
		}
	} else { // ID==0 вообще нет совпадений, нужно наращивать с основы
		return 0, nil
	}

	return 0, nil
}

// ////////
// рекурсивно корректируем адреса всех узлов
func updatingSampleTreeFromID(rt *sampleNode) {
	if rt.ID > 0 {
		MapCheckBlock(MapGwardSampleTreeFromID)
		rt.ParentNode = sampleTreeFromID[rt.ParentID] // wr.ParentNode адрес меняется из=за corretsParent(,
		MapFree(MapGwardSampleTreeFromID)
		sampleTreeFromID[rt.ID] = rt
	}
	if rt.Children == nil { // конец ветки
		return
	}
	for i := 0; i < len(rt.Children); i++ {
		updatingSampleTreeFromID(&rt.Children[i])
	}
}

///////////////////////////////////////////////////////////

// проверка работы дерева в разных ситуациях
func tryExperiment() {

	// ветки дерева для тестиварования
	fileTree := []string{
		"1|0|1|0|0|0",
		"2|0|2|0|0|0",
		"3|0|3|0|0|0",
		// второй уровень
		"4|1|1|10|0|0",
		"5|1|1|11|0|0",
		// третий уровень
		"6|5|1|11|2|0",
		// четвертый уровень
		"7|6|1|11|2|3",
	}
	// загрузить дерево из данных вверху
	initSampleTree(fileTree)

	//cond:= []int{9,2,3,4, } // вообще нет такой ветки ни на одном уровне  (maxLevel==0)
	//cond:= []int{1,4,3,4, } // есть совпадение только для первого уровня (maxLevel==1)
	cond := []int{1, 11, 5, 8} // не хватает 3 и 4-го уровня (maxLevel==2)
	//cond:= []int{1,11,2,8, } // не хватает 4-го уровня  (maxLevel==3)
	//cond:= []int{1,11,2,3, } // существующая ветка дерева (maxLevel==4)

	//cond:= []int{1,11,0,0, } // условие поиска не до конца ветки - нулевые ID образов 3 и 4 уровней НЕ ДОЛЖНО БЫТЬ ТАКОГО

	ID, maxLevel := findSampleBrange(0, cond, &sampleTree)
	if ID > 0 {
		if maxLevel == len(cond) { // найден конечный узел ветки

			// что-то сделать при нормальном распознавании
			return
		} else { // нужно дорастить ветку
			//Eсли условие cond []int содержит значение 0 для какого-то уровня, то на этом уровне поиск заканчивается.
			// считается нормальный поиск по неполному условию, доращивать нет смысла
			if cond[maxLevel] > 0 {
				ID = addFromNodeIDsToBrange(ID, maxLevel, cond)
				return
			}
		}
	} else { // ID==0 вообще нет совпадений, нужно наращивать с основы
		ID = addFromNodeIDsToBrange(sampleTree.ID, maxLevel, cond)
		return
	}

	return
}

//////////////////////////////////////////////////////////////

/*
	поиск конечного узла ветки (lastBrangeID) дерева (root - ID начального узла == 0) по массиву ID узлов ветки (cond []int)

Поиск начинается с первого узла (по значению cond[0]).
Если первый узел найден, то findSampleBrange вызывается рекурсивно для поиска следующего узла
и так далее, пока не будет найден удел последнего члена cond.
Если для текущего cond[level] узел в редеве не существует, то в дереве доращивается ветка по значениям cond []int.

Возвращает ID последней найденной ветки и ее уровень, если конечный узел не найден (если найден, то len(cond)-1).
Если вернула уровень меньше, чем максимальное число уровней ветки,
то ветка доращивается func addNodesToBrange
*/
func findSampleBrange(level int, cond []int, root *sampleNode) (int, int) {

	// Обработка случая когда мы достигли конца дерева
	if cond == nil || len(cond) <= 0 || level >= len(cond) {
		return root.ID, len(cond)
	}

	// Поиск узла с ID из списка cond в дочерних узлах текущего узла
	for _, child := range root.Children {
		if isEquivalentCondition(level, &child, cond) {
			// Если узел найден, продолжим рекурсивно искать в нем
			id, lev := findSampleBrange(level+1, cond, &child)
			if id == 0 {
				/*Eсли условие cond []int содержит значение 0 для какого-то уровня, то на этом уровне поиск заканчивается.
				if cond[lev]==0{// считается нормальный поиск по неполному условию

				}*/
				//return child.ID,lev
				return 0, lev // не найдено совпадение на данном уровне
			}
			// сюда доходит только после полного прохода всех итераций, так что просто возвращаем достигнутое
			return id, lev
		}
	}

	// не найден узел на данном уровне
	return 0, level
}

// //
func isEquivalentCondition(level int, node *sampleNode, cond []int) bool {
	// массив значений по уровням - в зависимости от числа уровней в ветке, прописывается вручную
	nArr := []int{node.levVal1, node.levVal2, node.levVal3, node.levVal4}

	for i := 0; i < len(cond) && i <= level; i++ {
		if cond[i] != nArr[i] {
			return false
		}
	}
	return true
}

/////////////////////////////////////////////////////////////////////////

/*
	Доращивание ветки, начиная с заданного узла fromID по массиву всех значений для ветки

Возвращает ID конечного узла ветки.
*/
func addFromNodeIDsToBrange(fromID int, lastLevel int, condArr []int) int {

	MapCheckBlock(MapGwardSampleTreeFromID)
	node, ok := sampleTreeFromID[fromID]
	MapFree(MapGwardSampleTreeFromID)
	if !ok {
		return 0
	}

	lastNodeID, _ := addNodesToBrange(node, lastLevel, condArr)

	return lastNodeID
}

/*
	рекурсивно создать все недостающие уровни ветки по значениям condArr []int

func addNodesToBrange запускается ТОЛЬКО ЕСЛИ точно известно, что узла с condArr [0] у fromNode нет !!!
Проверки на существование здесь нет: createNodeTree(fromNode, ..., false)
*/
func addNodesToBrange(fromNode *sampleNode, level int, condArr []int) (int, *sampleNode) {
	if fromNode == nil {
		return 0, nil
	}
	if level >= len(condArr) {
		return fromNode.ID, fromNode
	}
	//vArr := make([]int, len(condArr))// все vArr[n] имеют нулевые значения
	vArr := []int{0, 0, 0, 0}
	for i := 0; i <= level; i++ { // заполнить vArr до текущего уровня включительно, остальные - оставить нулями
		vArr[i] = condArr[i]
	}
	// с каждым уровнем vArr[n] добавляется новое значение для создания узла ветки следующего уровня.
	//true потому, что новый узел нужно делать только если такого еще нет!!!
	_, node := createNodeTree(fromNode, 0, vArr[0], vArr[1], vArr[2], vArr[3], true)
	level++
	id, node := addNodesToBrange(node, level, condArr)
	return id, node
}

///////////////////////////////////////

/* загрузить записанное дерево
 */
func loadSampleTree() {
	//	strArr,_:=lib.ReadLines(lib.GetMainPathExeFile()+"/memory_psy/automatizm_tree.txt")
	//	initSampleTree(strArr)
}
func initSampleTree(strArr []string) {
	sampleTreeMapFlag = true
	sampleTreeFromID = make(map[int]*sampleNode)
	sampleTreeMapFlag = false
	// определить нулевой узел
	sampleTreeFromID[0] = &sampleTree // все по нулям по умолчанию

	cunt := len(strArr)
	//просто проход по всем строкам файла подряд так что сначала идут дочки, потом - их родители
	for n := 0; n < cunt; n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		if len(strArr[n]) < 2 {
			panic("Сбой загрузки дерева: [" + strconv.Itoa(n) + "] " + strArr[n])
			return
		}
		p := strings.Split(strArr[n], "|")
		id, _ := strconv.Atoi(p[0])
		parentID, _ := strconv.Atoi(p[1])

		levVal1, _ := strconv.Atoi(p[2])
		levVal2, _ := strconv.Atoi(p[3])
		levVal3, _ := strconv.Atoi(p[4])
		levVal4, _ := strconv.Atoi(p[5])

		// новый узел с каждой строкой из файла
		MapCheckBlock(MapGwardSampleTreeFromID)
		parent, ok := sampleTreeFromID[parentID]
		MapFree(MapGwardSampleTreeFromID)
		if ok {
			// false - не проверять наличие ветки т.к. еще нет веток, они только глузятся
			createNodeTree(parent, id, levVal1, levVal2, levVal3, levVal4, false)
		}
	}
	return
}

// ///////////////////////////////////
func SaveSampleTree() {
	var out = ""
	cnt := len(sampleTree.Children)
	for n := 0; n < cnt; n++ { // чтобы записывалось по порядку родителей
		out += getSampleNode(&sampleTree.Children[n])
	}
	//lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/automatizm_tree.txt",out)
	return
}
func getSampleNode(wt *sampleNode) string {
	var out = ""
	//	if wt.ParentID>0 {
	out += strconv.Itoa(wt.ID) + "|"
	out += strconv.Itoa(wt.levVal1) + "|"
	out += strconv.Itoa(wt.levVal2) + "|"
	out += strconv.Itoa(wt.levVal3) + "|"
	out += strconv.Itoa(wt.levVal4)
	out += "\r\n"
	//	}
	if wt.Children == nil { // конец
		return out
	}
	for n := 0; n < len(wt.Children); n++ {
		out += getSampleNode(&wt.Children[n])
	}
	return out
}

/////////////////////////////////////

///////    попытка челез интерфесы и reflect оазывается слишком накрученной...
/*  из-за статической типизации в Go, создание новых узлов динамически в общем случае будет очень сложно реализовать,
особенно если структура sampleNode имеет разное количество уровней nodeN int.

type nodeI interface {
	CreateBranch(cond []int) nodeI
	// Остальные методы Node, которые можно определить
}
func findSampleBrange(level int, cond []int, root interface{}) int {

	if root == nil || level >= len(cond) {
		return 0
	}

	val := reflect.ValueOf(root)
	typeOfS := val.Type()


	// Достаем ID узла из структуры и проверяем соответствие с числом из cond
	for i := 0; i < val.NumField(); i++ {
		if typeOfS.Field(i).Name == "ID" && val.Field(i).Interface().(int) == cond[level] {
			// Увеличиваем уровень
			level++
			if level >= len(cond) {
				return val.Field(i).Interface().(int)
				// Когда достигаете конца текущей ветви:
				//newNode := nodeI.CreateBranch(cond[level:])
				//return findSampleBrange(level+1, cond, newNode)
			}
		}
	}

	// Если ветка поиска по cond закончилась, ищем далее по объекту root
	for i := 0; i < val.NumField(); i++ {
		if typeOfS.Field(i).Name == "Children" {
			children := val.Field(i)
			for j := 0; j < children.Len(); j++ {
				result := findSampleBrange(level, cond, children.Index(j).Interface())
				if result > 0 {
					return result
				}
			}
		}
	}

	return 0
}
*/
//////////////////////////////////////////////////

/*
Счетчик проблем для открытия доминаниы.
Массив пробных действий доминанты

*/

package psychic

import (
	"BOT/lib"
	"strconv"
	"strings"
)

/////////////////////////////////////

/*
Счетчик проблем: если для данного ID дерева проблем problemTryCount[id] >3 -
на пятой стпени разивтия открывать доминанту. Счетчик продолжает накапливаться.
Сохраняется в файле problem_counts.txt
*/
var problemTryCount = make(map[int]int)
var MapGwardProblemTryCount = lib.RegNewMapGuard()

///////////////////////////////////////

func saveProblemTryCount() {
	var out = ""
	lib.MapCheckBlock(MapGwardProblemTryCount)
	for k, v := range problemTryCount {
		out += strconv.Itoa(k) + "|"
		out += strconv.Itoa(v)
		out += "\r\n"
	}
	lib.MapFree(MapGwardProblemTryCount)
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/dominanta_try_count.txt", out)
}

func loadProblemTryCount() {
	lib.MapCheckWrite(MapGwardProblemTryCount)
	problemTryCount = make(map[int]int)
	lib.MapFree(MapGwardProblemTryCount)

	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/dominanta_try_count.txt")
	cunt := len(strArr)
	for n := 0; n < cunt; n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		p := strings.Split(strArr[n], "|")
		ID, _ := strconv.Atoi(p[0])
		count, _ := strconv.Atoi(p[1])

		//problemTryCountMu.Lock()
		//defer problemTryCountMu.Unlock()

		lib.MapCheckWrite(MapGwardProblemTryCount)
		problemTryCount[ID] = count
		lib.MapFree(MapGwardProblemTryCount)
	}
}

///////////////////////////////////////

/////////////////////////////////////
/* действия, которые были опробованы, не только для доминанты.
Сохраняются в файле dominanta_try_ctions.txt
в формате dominantaID|[]tryAction 1|[]tryAction 2|....

Очищаются с удалением доминанты с помощью removeTryActionArr(dominantaID int)
в том числе и для временных массивов (отрицательные dominantaID)
*/
type tryAction struct {
	ID             int
	actionsImageID int
	effect         int // отрицательный, нулевой или положительный
}

/*
	массив испробованных действий для данных условий объекта внимания importance

index == Dominanta.ID

Могут быть пробные действия и без доминанты,
тогда индексом в tryActionArr[-id] отрицательное значение ID дерева проблем.
Такие действия не записываются в файл и сохраняются только до перезагрузки (как кратковременная память).
*/
var tryActionArr = make(map[int][]*tryAction)
var MapGwardTryActionArr = lib.RegNewMapGuard()

///////////////////////////////////////

// удаление массива пробных действий
func removeTryActionArr(dominantaID int) {
	lib.MapCheckBlock(MapGwardTryActionArr)
	ad := tryActionArr[dominantaID]
	for k, _ := range ad {
		delete(tryActionArr, k)
	}
	lib.MapFree(MapGwardTryActionArr)
}

/////////////////////////////////

// удаление временных массивов пробных действий (отрицательные значяения индеков) - во сне
func RemoveTemporeryTryActionArr() {
	lib.MapCheckBlock(MapGwardTryActionArr)
	for k, _ := range tryActionArr {
		if k > 0 {
			continue
		}
		delete(tryActionArr, k)
	}
	lib.MapFree(MapGwardTryActionArr)
}

/////////////////////////////////

// // сохраниение tryActionArr в формате dominantaID|[]tryAction 1|[]tryAction 2|....
// полный вариант: dominantaID|id,actionsImageID,effect|id2,actionsImageID2,effect2|....
func saveTryActionArr() {
	var out = ""
	lib.MapCheckBlock(MapGwardTryActionArr)
	for k, v := range tryActionArr {
		if k < 0 { // такие действия хранятся временно
			continue
		}
		out += strconv.Itoa(k) + "|"
		for n := 0; n < len(v); n++ {
			out += strconv.Itoa(v[n].ID) + ","
			out += strconv.Itoa(v[n].actionsImageID) + ","
			out += strconv.Itoa(v[n].effect) // в конце нельзя ставить / иначе в loadTryActionArr будет определять пустую строку и вылетать
		}
		out += "\r\n"
	}
	lib.MapFree(MapGwardTryActionArr)
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/dominanta_try_actions.txt", out)
}

func loadTryActionArr() {

	tryActionArr = make(map[int][]*tryAction)

	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/dominanta_try_actions.txt")
	cunt := len(strArr)
	for n := 0; n < cunt; n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		p := strings.Split(strArr[n], "|")
		dominantaID, _ := strconv.Atoi(p[0])
		// все записанные tryAction:
		for m := 1; m < len(p); m++ {
			a := strings.Split(p[m], ",")
			nodeID, _ := strconv.Atoi(a[0])
			nodeActionsImageID, _ := strconv.Atoi(a[1])
			nodeEffect, _ := strconv.Atoi(a[2])
			addNewTryAction(nodeID, dominantaID, nodeActionsImageID, nodeEffect, false)
		}
	}
}

//////////////////////////////////////////////////////

/*
	добавить новое действие actionsImageID для доминанты или временное в массив действий, которые были опробованы

Для одной dominantaID м.б. массив []*tryAction
Возвращает ID новой tryAction и сам *tryAction
*/
var tryActionLastID = 0

func addNewTryAction(id int, dominantaID int, actionsImageID int, effect int, CheckUnicum bool) (int, *tryAction) {
	if CheckUnicum {
		oldID, oldVal := checkTryAction(dominantaID, actionsImageID)
		if oldVal != nil {
			return oldID, oldVal
		}
	}

	if id == 0 {
		tryActionLastID++
		id = tryActionLastID
	} else {
		//		newW.ID=id
		if tryActionLastID < id {
			tryActionLastID = id
		}
	}
	var node tryAction
	node.ID = id
	node.actionsImageID = actionsImageID
	node.effect = effect

	lib.MapCheckWrite(MapGwardTryActionArr)
	tryActionArr[dominantaID] = append(tryActionArr[dominantaID], &node)
	lib.MapFree(MapGwardTryActionArr)

	return id, &node
}

////////////////////////

// выдать *tryAction по actionsImageID
func checkTryAction(dominantaID int, actionsImageID int) (int, *tryAction) {
	lib.MapCheckBlock(MapGwardTryActionArr)
	for id, v := range tryActionArr {
		for n := 0; n < len(v); n++ {
			if v[n].actionsImageID == actionsImageID {
				lib.MapFree(MapGwardTryActionArr)
				return id, v[n]
			}
		}
	}
	lib.MapFree(MapGwardTryActionArr)
	return 0, nil
}

/////////////////////////////////////////

// выдать массив *tryAction для ID доминанты
func getDomTryActionArr(dominantaID int) []*tryAction {
	// т.к. у tryActionArr=make(map[int][]*tryAction)  index == Dominanta.ID
	return tryActionArr[dominantaID]
}

// //////////////////////////////////////
// выдать массив []int c tryAction.ID для ID доминанты
func getDomTryActionIDArr(dominantaID int) []int {
	var arr []int
	lib.MapCheckBlock(MapGwardTryActionArr)
	for id, v := range tryActionArr {
		if id == dominantaID {
			for n := 0; n < len(v); n++ {
				arr = append(arr, v[n].ID)
			}
		}
	}
	lib.MapFree(MapGwardTryActionArr)
	return arr
}

////////////////////////////////////////

/*
	выявления завершения актуальности доминанты при сканировании массива доминант

Это - функция, обратная открытию доминанты - закрытие уже не актуальной доминанты.
В первую очередь учитывается endTime
*/
func lookForUnactualDominant() {
	idleness := isIdleness()
	for dID, v := range DominantaProblem {
		if v == nil {
			continue
		}

		if v.isSuccess == 3 || v.isSuccess == 4 { // не смотреть закрытые проблемы
			continue
		}
		if v.endTime != 0 {
			diffT := LifeTime - v.birthTime
			if v.endTime > diffT { // уже не актуальна
				// закрыть доминанту
				v.isSuccess = 4
				clinerProblemDominenta(dID)
			}
		}

		// TODO другие критерии потери актуальности

		// необходимо время от времени возобновлять в памяти, напоминая зачем они были поставлены.
		if idleness { //Если процесс в лени, а не во сне
			// TODO выйти из лени и начать обрабатывать доминанту - ее экстремальный объект
		}

	}

}

//////////////////////////////////////////////////////////////////////////

/*
При появлении экстремального стимула на 3-м уровне осмысления определить, нет ли доминанты с Dominanta.objectID.
Если есть, то попробовать решить доминанту реализацией ее действия Dominanta.typeTargetAction.
Если идентифицирован Dominanta.objectID возвращает true
*/
func checkDominantsJbject(objectID *extremImportance) bool {
	if objectID == nil {
		return false
	}
	// при опасной ситуцации не смотреть
	if CurrentInformationEnvironment.veryActualSituation || CurrentInformationEnvironment.danger {
		return false
	}

	// TODO

	return false
}

////////////////////////////////////////////////////////////////////////////

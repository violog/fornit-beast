/* Доминанта нерешенной проблемы
Это - атрибут 5-й стадии развития - творчества

ДОМИНАНТА - нерешенная проблема,
т.е. не найденное пока что ответное действие (или цепочка ответных действий) или цепочка мыслительных операций.
Цель - решение проблемы problemTreeID где ставится цель:
в func infoFunc8 ищестя текущая цель problemTreeInfo.purposeID устанавливается в getMentalPurpose()

!!! Активная Доминанта, если она есть, ВСЕ ВРЕМЯ сопровождает появление новой инфы восприятия так,
чтобы была возможность найти решение по аналогии. Модели понимания дают аналогии действий.
Только этим она и отличается от простой мыслительной Темы и Цели.

В структуре доминанты есть самая общая цель типа ProblemTreeNode.purposeID,
а также тема мышления ProblemTreeNode.themeID
которые сохраняются как образ текущей ветки дерева проблем detectedActiveLastProblemNodID

Главным идентификатором Доминанты является detectedActiveLastProblemNodID, НО
под одним detectedActiveLastProblemNodID могут быть создано несколько доминант для разных objectID (которого может и не быть).
Поэтому при выборе Доминанты следует находить ту, что соотвествует текущему значению extremImportanceObject.

Решением доминанты (закрытие гештальта) является нахождение удачного действия типа tryAction

Сохраняется в файле /memory_psy/dominanta.txt
в формате:
ID|problemTreeID|weight|birthTime|endTime|objectID|targetActionID|typeTargetAction|needExecuteTargetAction|tryActionsKey|isSuccess

Доминанты не удаляются и составляют базу личного опыта. TODO - как-то использовать.
Если доминанта решена, то можно использовать такое решение хоть сразу, хоть потом,
но озарение возникает сразу в func checkRelevantAction(

В творчестве доминанта используется в трех ипостасях:
1. нахождение действия по случайной аналогии (это сделано)
и то, что использует стек обобщения saveBaseLinksCicleSynthesis (который не имеет пока поддержки в mental_memories.go):
2. ТВОРЧЕСКОЕ СОПОСТАВЛЕНИЕ И ОБОБЩЕНИЕ
3. ТВОРЧЕСКАЯ ФАНТАЗИЯ
*/

package psychic

import (
	"BOT/lib"
	"strconv"
	"strings"
)

// /////////////////////////////////////
type Dominanta struct {
	ID            int
	problemTreeID int // ID дерева проблем - общие ситуация, тема и цель
	// вес значимости проблемы - чисто оценочно, Доминанты не конкурируют по весу значимости
	weight int
	// время рождения проблемы в днях жизни (как в условных рефлексах)
	birthTime int
	// срок актуальности цели, в днях, может быть очень приблизительным или не ограниченным (endTime==0)
	endTime int
	/*объект размышления типа extremImportance Тут может быть новая фраза и т.п. новые дела
	- объект, с которым возможно совершить какое-то действие.
	Это должен быть Стимул, для которого ищется Ответ типа extremImportanceObject
	т.е. с каждым стимулом нужно проверять, не это ли данный объект, с которым возможно попробовать совершить действие
	*/
	objectID int
	/* целевое действие, которые желательно совершить с объектом objectID
	TODO в принципе, тут может быть массив действий или же цепочка действий
	*/
	targetActionID int
	/* тип образа действия:
	0 - изменение состояния жизненных параметров - уточняющихся в образе objectID
	1 - изменение эмоционального состояния
	2 - достижение (или избегание в случае needExecuteTargetAction==0) образа ситуации
	3 - моторное действие типа ActionsImage
	4 - образ понимания сути объекта objectID - реализуется в ходе решения проблемы доминанты
	с получением определенного вида опыта (модели понимания, правила, категории, обобщения и т.п.)
	Решение об удовлетворенности реализацией типа 4 принимается осознанно т.е. через информационную картину, с закрытием проблемы.

	*/
	typeTargetAction int
	//метка о том, нужно ли совершить такое действие (1) или, наоборот, нужно его избегать (0).
	needExecuteTargetAction int // только 0 или 1

	/* действия, которые были опробованы при попытке решения (не только для доминанты).
	индекс для tryActionsID[] действия, которое принято как пробное решение, вначале =-1
	выдать массив *tryAction для ID доминанты
	dAIarr:=getDomTryActionArr(dominantaID int)
	или просто aArr:=tryActionArr[v.ID] т.к. у tryActionArr=make(map[int][]*tryAction)  index == Dominanta.ID
	При открытии доминанты tryActionsKey может быть ==-1 но заполняться ID образов действий в ходе решения доминанты.
	TODO с вводом targetActionID вроде бы нет смысла поддерживать tryActionsKey но в структуре tryAction есть ээффект от попытки,
	тут нужно оптмизировать структуру...
	*/
	tryActionsKey int // для каждой доминанты является ключом tryActionArr=make(map[int][]*tryAction)

	/* // степень решения проблемы:
	0 - ничего нет,
	1 - начато решение, - уже есть найденные tryActionsKey но не точные
	2 - частичное решение, - есть точно подходящий к условиям tryActionsKey
	3 - успешное решение - эффект проверен, гельштат закрыт, можно не думать, а пользолваться опытом
	4 - потеря актуальности доминанты, более не рассматривается
	*/
	isSuccess int
}

// var DominantaProblem=make(map[int]*Dominanta)
var DominantaProblem []*Dominanta // сам массив
// var AFromID = make([]*aNode, len(strArr))//задать сразу имеющиеся в файле число при загрузке
// запись члена
func WriteDominantaProblem(index int, value *Dominanta) {
	if index >= len(DominantaProblem) {
		newSlice := make([]*Dominanta, index+1)
		copy(newSlice, DominantaProblem)
		DominantaProblem = newSlice
	}
	DominantaProblem[index] = value
}

// считывание члена
func ReadeDominantaProblem(index int) (*Dominanta, bool) {
	if index >= len(DominantaProblem) || DominantaProblem[index] == nil {
		return nil, false
	}
	return DominantaProblem[index], true
}

///////////////////////////////////////

var CurrentProblemDominanta *Dominanta

////////////////////////////////////////////

////////////////////////////////////////////
/* создать новую доминанту проблемы для detectedActiveLastUnderstandingNodID

 */
var lastDominantaID = 0

func createNewDominanta(id int, problemTreeID int, weight int, endTime int, objectID int, targetActionID int, typeTargetAction int, needExecuteTargetAction int, birthTime int, tryActionsKey int, isSuccess int, CheckUnicum bool) (int, *Dominanta) {
	if CheckUnicum {
		oldID, oldVal := checkUnicumDominanta(problemTreeID, objectID, targetActionID, typeTargetAction, needExecuteTargetAction)
		if oldVal != nil {
			return oldID, oldVal
		}
	}
	if id == 0 {
		lastDominantaID++
		id = lastDominantaID
	} else {
		//		newW.ID=id
		if lastDominantaID < id {
			lastDominantaID = id
		}
	}

	var node Dominanta
	node.ID = id
	node.problemTreeID = problemTreeID
	node.weight = weight
	node.objectID = objectID
	node.targetActionID = targetActionID
	node.typeTargetAction = typeTargetAction
	node.endTime = endTime
	node.needExecuteTargetAction = needExecuteTargetAction
	node.birthTime = birthTime
	node.tryActionsKey = tryActionsKey
	node.isSuccess = isSuccess

	//DominantaProblem[id]=&node
	WriteDominantaProblem(id, &node)

	return id, &node
}

// /////////////
func checkUnicumDominanta(problemTreeID int, objectID int, targetActionID int, typeTargetAction int, needExecuteTargetAction int) (int, *Dominanta) {
	for id, v := range DominantaProblem {
		if v == nil {
			continue
		}
		if problemTreeID != v.problemTreeID || objectID != v.objectID ||
			targetActionID != v.targetActionID ||
			typeTargetAction != v.typeTargetAction ||
			needExecuteTargetAction != v.needExecuteTargetAction {
			continue
		}
		return id, v
	}
	return 0, nil
}

//////////////////////////////////////////////////////

// ///////////////////////////////  if doWritingFile{SaveProblemDominenta() }
func SaveProblemDominenta() {
	saveProblemTryCount()
	saveTryActionArr()

	var out = ""
	for k, v := range DominantaProblem {
		if v == nil {
			continue
		}
		out += strconv.Itoa(k) + "|"
		out += strconv.Itoa(v.problemTreeID) + "|"
		out += strconv.Itoa(v.weight) + "|"
		out += strconv.Itoa(v.birthTime) + "|"
		out += strconv.Itoa(v.endTime) + "|"
		out += strconv.Itoa(v.objectID) + "|"
		out += strconv.Itoa(v.targetActionID) + "|"
		out += strconv.Itoa(v.typeTargetAction) + "|"
		out += strconv.Itoa(v.needExecuteTargetAction) + "|"

		out += strconv.Itoa(v.tryActionsKey) + "|"
		out += strconv.Itoa(v.isSuccess)
		out += "\r\n"
	}
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/dominanta.txt", out)
}

/////////////////////////////////////////

func loadProblemDominenta() {
	loadProblemTryCount()
	loadTryActionArr()

	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/dominanta.txt")
	cunt := len(strArr)
	//	DominantaProblem=make(map[int]*Dominanta)
	DominantaProblem = make([]*Dominanta, cunt) //задать сразу имеющиеся в файле число
	for n := 0; n < cunt; n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		p := strings.Split(strArr[n], "|")
		id, _ := strconv.Atoi(p[0])
		problemTreeID, _ := strconv.Atoi(p[1])
		weight, _ := strconv.Atoi(p[2])
		birthTime, _ := strconv.Atoi(p[3])
		endTime, _ := strconv.Atoi(p[4])
		objectID, _ := strconv.Atoi(p[5])
		targetActionID, _ := strconv.Atoi(p[6])
		typeTargetAction, _ := strconv.Atoi(p[7])
		needExecuteTargetAction, _ := strconv.Atoi(p[8])
		tryActionsKey, _ := strconv.Atoi(p[9])
		isSuccess, _ := strconv.Atoi(p[10])

		var saveDoWritingFile = doWritingFile
		doWritingFile = false

		createNewDominanta(id, problemTreeID, weight, endTime, objectID, targetActionID, typeTargetAction, needExecuteTargetAction, birthTime, tryActionsKey, isSuccess, false)

		doWritingFile = saveDoWritingFile
	}
	return
}

//////////////////////////////////

// степень решенной Доминанты (зактыть гештальт isSuccess=3 )
func solutionDominanta(dID int, isSuccess int) {
	//	DominantaProblem[dID].isSuccess=isSuccess
	node, ok := ReadeDominantaProblem(dID)
	if !ok {
		return
	}
	node.isSuccess = isSuccess
	clinerProblemDominenta(dID) // удаление массива пробных действий
}

// ///////////////////////////////////////////
// удаление доминанты, проставив isSuccess=true
func clinerProblemDominenta(dID int) {

	// удаление массива пробных действий
	removeTryActionArr(dID)
}

//////////////////////////////////

///////////////////////////////////
/*  Это не нужно!
// наиболее важная доминанта в заданном эмоциональном контексте
func getMainDominanta(emotionID int) int {
	var dominantaID=0
	// TODO

		switch emotionID {
		case 1: //Пищевой	- Пищевое поведение, восполнение энергии, на что тратится время и тормозятся антагонистические стили поведения.

		case 2: //Поиск	- Поисковое поведение, любопытство. Обследование объекта внимания, поиск новых возможностей.

		case 3: //Игра	- Игровое поведение - отработка опыта в облегченных ситуациях или при обучении.

		case 4: //Гон	- Половое поведение. Тормозятся антагонистические стили

		case 5: //Защита	- Оборонительные поведение для явных признаков угрозы или плохом состоянии.

		case 6: //Лень	- Апатия в благополучном или безысходном состоянии.

		case 7: //Ступор	- Оцепенелость при непреодолимой опастbase_context_activnostности или когда нет мотивации при благополучии или отсуствии любых возможностей для активного поведения.

		case 8: //Страх	- Осторожность при признаках опасной ситуации.

		case 9: //Агрессия	- Агрессивное поведение для признаков легкой добычи или защиты (иногда - при плохом состоянии).

		case 10: //Злость	- Безжалостность в случае низкой оценки .

		case 11: //Доброта	- Альтруистическое поведение.

		case 12: //Сон - Состояние сна. Освобождение стрессового состояния. Реконструкция необработанной информации.

		}

return dominantaID
}*/
////////////////////////////////////////////////////

// для пульта
func GetDominantaIDString(id int) string {

	return "Еще не сделан вывод Доминант"
}

////////////////////////////////////////////

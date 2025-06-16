/* Создание, останов, перезапуск ЦИКЛОВ МЫШЛЕНИЯ

 */

package psychic

import (
	"strconv"
)

///////////////////////////////////////////////////////////
/* ID текущего главного циклаfunctionsInAllCickles
НЕ ОБНУЛЯЕТСЯ, а только перекрывается, сохраняя инфу о последнем.
*/
var activedCyckleID = 0 //

// ID запускаемых во всех циклах инфо-функций
var functionsInAllCickles []int

//////////////////////////////////////////////

/////////////////////// ИНФОРМАЦИЯ О ЦИКЛАХ
/* Каждый вновь запускаемый цикл становится главный осознаваемым, а другие уходят в подсознание.

Поддержка многоцикленности мышления:
*/
type cycleInfo struct {
	ID              int    // ID первого звена цикла. Из любого ID шага получается - id:=getBaseIdCyckle(stepID)
	order           int    // порядок нумерации циклов при возникновении нового в createNewCycleIteration
	isMainCycle     bool   // это - главный, осознаваемый цикл мышления
	count           int    // число проделанных шагов
	isWaitingPeriod bool   // это - период ожидания ответа с пульта на действие ТОЛЬКО ДЛЯ ГЛАВНОГО ЦИКЛА
	weight          int    // значимость цикла в конкуренции параллельных
	pulsCount       int    // старые циклы и взаимная конкурентность циклов - уменьшает их число
	themeType       int    // тема мышления в момент создания цикла, ThemeImage.Type
	problevID       int    // проблема в момент создания цикла
	impObjID        int    // последний extremImportanceObject.extremObjID (перекрывается новым)
	lastFuncID      int    // какая последняя инфо-функция вызывалась в этом цикле - против дребезга
	lastProcessID   string //название последней из функций процесса в understanding_functions.go вызывалась
	func0Arr        []int  // последовательность вызываемых функций этого цикла в течении одной итерации  func consciousnessThinking()
	funcArr         []int  // последовательность вызываемых функций этого цикла
	idle            bool   // пустой цикл (крутит вхолостую) - запускаются в 5 раз реже: PulsCount%5 !=0
	isStupor        bool   // этот цикл, пришедший в ступор, больше не перезепускать
	/* Фоновые циклы могут продолжать обрабатывать инфу в режиме dreaming, но если это происходит в главном цикле,
	   то это значит общий режим пассивного размышления, который тормозит восприятие стимулов (isRepressionStimulsNoise).
	*/
	dreaming bool // пассивный режим размышления (сновидение или мечтания)
	/*	ID кадра истории эпиз.памяти в сновидении или мечтании при процессе текущего прохода,
		после чего isDreamInterrupt устанавливается в значение для продолжения прохода.
		При isDreamInterrupt==true процесс приостанавливается и потом продолжается с текущего dreamingEpisodeHistoryID
	*/
	dreamingEpisodeHistoryID int // текущий ID исторической памяти EpisodicHistoryArr[] при процессе текущего прохода

	log string // последовательность шагов для пульта
}

// инфа о работающих параллельно ментальных циклах - глобальный детектор занятости
// var cyclesArr =make(map[int]*cycleInfo)
var cyclesArr []*cycleInfo // сам массив
// var AFromID = make([]*aNode, len(strArr))//задать сразу имеющиеся в файле число при загрузке
// запись члена
func WritecyclesArr(index int, value *cycleInfo) {
	addcyclesArr(index)
	cyclesArr[index] = value
}
func addcyclesArr(index int) {
	if index >= len(cyclesArr) {
		newSlice := make([]*cycleInfo, index+1)
		copy(newSlice, cyclesArr)
		cyclesArr = newSlice
	}
}

// считывание члена
func ReadecyclesArr(index int) (*cycleInfo, bool) {
	if index >= len(cyclesArr) || cyclesArr[index] == nil {
		return nil, false
	}
	return cyclesArr[index], true
}

////////////////////////////////////

// усли превысит 1000, то в func sleepNecessityDetector() - потребность во сне
func GetCycleCount() int {
	return len(cyclesArr)
}

///////////////////////////////////

/* создание нового cycleInfo
 */
var countbaseID = 0

func addCounterGoNext() *cycleInfo {
	var new cycleInfo
	countbaseID++
	new.ID = countbaseID
	cycle := &new
	cycle.pulsCount = PulsCount
	cycle.isMainCycle = true // новый - всегда главный
	activedCyckleID = cycle.ID

	//	cyclesArr[countbaseID] = cycle
	WritecyclesArr(countbaseID, cycle)

	return cycle
}

/////////////////////////////////////

// /  установить цикл как главный
func setAsMaimCycle(baseID int) {

	for id, v := range cyclesArr {
		if v == nil {
			continue
		}
		if id == baseID {
			v.isMainCycle = true
			activedCyckleID = id
		} else {
			v.isMainCycle = false
			// 			v.dreaming=1 //все не главные циклы делать дремными
			//			НЕТ отличие только в озарении для неглавных
		}
	}

	return
}

///////////////////////////////////

// сбросить все в цикле чтобы начал как новый
func resetCycleAndBeginAsNew(cycle *cycleInfo) {
	cycle.count = 0
	cycle.isWaitingPeriod = false
	cycle.dreaming = false
	cycle.pulsCount = 0
	cycle.themeType = 0
	cycle.problevID = 0
	cycle.impObjID = 0
	cycle.lastFuncID = 0
	cycle.lastProcessID = ""
	cycle.func0Arr = nil
	cycle.funcArr = nil
	cycle.idle = false
	cycle.isStupor = false
}

// ////////////////////// возвращает главный цикл
func resetMineCycleAndBeginAsNew() *cycleInfo {
	if len(cyclesArr) == 0 { // еще нет циклов
		beginMentalCycle() // начать первый, главный цикл мышлления с вызовом func infoFunc8()
	}
	var mainС *cycleInfo
	for _, v := range cyclesArr {
		if v == nil {
			continue
		}
		if v.isMainCycle {
			mainС = v
			resetCycleAndBeginAsNew(v) // сбросить все в цикле чтобы начал как новый
			break
		}
	}
	return mainС
}

/////////////////////////////////////////////////

// найти новый главный цикл по значимости, кроме удаляемого
func foundMainCycle(delID int) {
	// удаленный не является главным?
	//c,ok:=ReadecyclesArr(delID)
	c, ok := ReadecyclesArr(delID)
	if ok && !c.isMainCycle {
		return // значит есть уже главный
	}
	var max = 0
	var idMax = 0

	for id, v := range cyclesArr {
		if v == nil {
			continue
		}
		if v.weight > max {
			max = v.weight
			idMax = id
		}
	}

	if idMax > 0 {
		setAsMaimCycle(idMax)
	}
}

///////////////////////////////////////////

func deleteCycle(baseID int) {
	foundMainCycle(baseID) //найти новый главный цикл по значимости, кроме удаляемого

	//	delete(cyclesArr, baseID)
	WritecyclesArr(baseID, nil)
}

/*
	закончить ВСЕ циклы мышления

приведет к прерыванию циклов в func consciousnessElementary: //если нет такого цикла, то завершить эту итерацию
Используется только в func endAllCycles()
*/
func deleteAllCounterGoNext() {
	//cyclesArr =make(map[int]*cycleInfo) // сброс всех счетчиков
	cyclesArr = nil
}

//////////////////////////////////////////

// погасить все циклы кроме дремы
func endNoDereamsCycles() {

	for id, v := range cyclesArr {
		if v == nil {
			continue
		}
		if !v.dreaming {
			endBaseIdCycle(id)
		}
	}
}

// погасить все циклы дремы  mine - только главные
func endDereamsCycles(mine bool) {
	for id, v := range cyclesArr {
		if v == nil {
			continue
		}
		if mine && !v.isMainCycle {
			continue
		}
		if v.dreaming {
			endBaseIdCycle(id)
		}
	}
}

// все циклы - дремы
func setDereamsForAllCycles() {
	for _, v := range cyclesArr {
		if v == nil {
			continue
		}
		v.dreaming = true
	}
}

// главный цикл вывести из дремы
func setWakuoForMainCycles() {
	for _, v := range cyclesArr {
		if v == nil {
			continue
		}
		if v.isMainCycle {
			v.dreaming = false
		}
	}
}

////////////////////////////////////////////////

/*
	Начать новый цикл мышления. Или начать новый с прерванного fromID.

Создать новое Базовое звено цепи для данной активности деревьев
и пройти цепочку до конца, чтобы продолжить цикл от него.

Создать уникальный ID нового цикла - cycleID (cycleID,goNext:=createNewCycleIteration()) делать НЕ СТОИТ
т.к. тогда придется вызывать func consciousnessElementary с этим параметром

типично:
createNewCycleIteration()// начать новый цикл
*/
var curCycleOrder = 0 // порядок создания цикла, всегда УНИКАЛЬНЫЕ ЗНАЧЕНИЯ даже при удалении промежуточных циклов
func createNewCycleIteration() *cycleInfo {
	shotrMemepMemArr = nil
	extremImportanceObject = nil
	problemExtremImportanceObject = nil
	extremImportanceMentalObject = nil

	// НАЧАТЬ НОВЫЙ ЦИКЛ
	c := addCounterGoNext() // новый цикл - всегда главный
	activedCyckleID = c.ID
	curCycleOrder++
	c.order = curCycleOrder

	node, ok := ReadeThemeImageFromID(oldThemeID)
	if ok {
		c.themeType = node.Type // - от последнего func runNewTheme и там перекроется, если цикл создан в ней.
	}

	c.problevID = detectedActiveLastUnderstandingNodID
	c.weight = getCycleWeight()
	c.log += GetSelfPerceptionInfo() //getCurInformEnviroment()
	c.log += conscienceStatus        // сохранить conscienceStatus объективного прохода

	setAsMaimCycle(c.ID) // сделать главным

	reduseUpdating() //старые циклы и взаимная конкурентность циклов - уменьшает их число

	conscienceStatus = ""
	c.log += "<span style='color:#006300'><b>Начат новый цикл мышления c ID=" + strconv.Itoa(c.ID) + ".</b></span><br>"

	return c
}

//////////////////////////////////////////////////

/*
	прерывание осмысления с сохранением предыдущего и началом нового цикла - итак происходит - новый цикл - главный

возврат к прерванному мышлению rememberInterruptImage()
*/
func interruptMentalWork(cycleID int) {
	if len(cyclesArr) == 0 {
		return
	}
	// прерывание только если важная ситуация
	if !CurrentInformationEnvironment.veryActualSituation && !CurrentInformationEnvironment.danger {
		return
	}
	addInterruptMemory(cycleID)
	idlenessType = 0

	//	wasInterruptedMentCickle = true // начало прохода ментального цикла
	//	functionsInAllCickles = nil     //ID запускаемых в текущем цикле инфо-функций

}

//////////////////////////////////////////////////////
/* выборка прерванного размышления из стека и начало размышления об этом

 */
func rememberInterruptImage(c *cycleInfo) *InterruptImage {
	// вспомнить последнее прерванное размышление
	lastImg := InterruptMemory[len(InterruptMemory)-1]
	mentalInfoStruct.ThemeImageType = lastImg.ThemeImageType
	mentalInfoStruct.mentalPurposeID = lastImg.PurposeImageID

	mentalInfoStruct.ExtremObjID = lastImg.ExtremObjID

	// удалить последний элемент массива
	if len(InterruptMemory) > 0 {
		InterruptMemory = InterruptMemory[:(len(InterruptMemory) - 1)]
	}

	// начать размышление с прерванного шага, не с самого начала
	//запустить прерванное размышление
	//createNewCycleIteration(lastID)// тут обнулялось extremImportanceObject
	setAsMaimCycle(c.ID) ///  просто установить цикл как главный

	if lastImg.ExtremObjID > 0 {
		extremImportanceObject = getExtremObjFromID(lastImg.ExtremObjID)
	}

	return lastImg
}

//////////////////////////////////////////////////////////////////////

// ///////////////////////////////////////
// закончить цикл с базовым ID
func endBaseIdCycle(baseID int) {
	foundMainCycle(baseID) //найти новый главный цикл по значимости, кроме удаляемого

	//Вызывает постоянную блокировку т.к. во внешнем цикле уже есть неснятая блокировка
	//!!!	lib.MapCheck(MapGwardCyclesArr)

	deleteCycle(baseID)
}

/////////////////////////////////////////
/* закончить ВСЕ циклы мышления
приведет к прерыванию циклов в func consciousnessElementary: //если нет такого цикла, то завершить эту итерацию
*/
func endAllCycles() {
	deleteAllCounterGoNext() // удаление всех счетчиков всех циклов -
	stopCyckles()            //  закончить обслуживание циклов мышления
}

// ///////////////////////////////////////
//
//	закончить обслуживание всех циклов мышления
func stopCyckles() {
	functionsInAllCickles = nil //ID запускаемых в текущем цикле инфо-функций

	conscienceStatus += "Остановка всех циклов мышления.<br>"
}

////////////////////////////////////////////

// старые циклы и взаимная конкурентность циклов - уменьшает их число
func reduseUpdating() {
	var delList = "" // список удаленных
	// удалить очень старые
	for id, v := range cyclesArr {
		if v == nil {
			continue
		}
		if (PulsCount-v.pulsCount) > 3600 && !v.isMainCycle { // старше часа
			delList += strconv.Itoa(v.order) + ";"

			//delete(cyclesArr, id)
			WritecyclesArr(id, nil)

		}
	}

	// взаимное торможение в каждой группе goNextFromUnderstandingNodeIDArr, оставлять только один самый весомый или главныый цикл из всего пакета
	//разбить по группам
	var themArr = make(map[int][]*cycleInfo)
	/*  не нужно по проблемам, раз есть по темам ?
	var problemArr []int // какие ШВ проблем в циклах
	// сгруппировать по проблемам TODO
		for id, v := range problemArr {
		if v==nil{continue}
			cyclesArrMapCheck()
			base,ok := cyclesArr[v[0].ID]
			if ok{
				themArr[id]=append(themArr[id],base)
			}
		}
		// конкуренция по каждой группе
		for _, v := range themArr {
		if v==nil{continue}
			delList+=cycleLateralBraking(v)
		}
	*/
	/////////////////////////////////////////////// удаление циклов
	// взаимное торможение в каждой группе одинаковых Тем, оставлять только один самый весомый или главныый цикл из всего пакета
	//разбить по темам
	themArr = make(map[int][]*cycleInfo)

	for _, v := range cyclesArr {
		if v == nil {
			continue
		}
		themArr[v.themeType] = append(themArr[v.themeType], v)
	}

	// конкуренция по каждой теме
	for _, v := range themArr {
		if v == nil {
			continue
		}
		delList += cycleLateralBraking(v)
	}

	if len(delList) > 0 {
		conscienceStatus += "<span style='color:red'>" + delList + " - циклы удалены в ходе взаимной конкуренции и устаревания.</span><br>"
	}
}

// / конкуренция в каждом массиве циклов
func cycleLateralBraking(cArr []*cycleInfo) string { // возвращает число удаленных
	if len(cArr) < 2 { // раз один, то нет конкуренции
		return ""
	}
	var delList = ""
	max := 0
	cID := 0
	age := 0 // возраст
	for i := 0; i < len(cArr); i++ {
		v := cArr[i]

		//base,ok := cyclesArr[v.ID]
		base, ok := ReadecyclesArr(v.ID)
		if !ok {
			continue
		}
		if base.isMainCycle {
			cID = v.ID
			break
		}
		if base.weight > max {
			max = base.weight
			age = base.pulsCount
			cID = v.ID
		}
		// если равны по значимости, то преимущество в более молодом возрасте
		if base.weight == max && base.pulsCount < age {
			max = base.weight
			age = base.pulsCount
			cID = v.ID
		}
	}
	if cID > 0 { // удалить все кроме cID
		for i := 0; i < len(cArr); i++ {
			if cArr[i].ID == cID {
				continue
			}
			if !cArr[i].isMainCycle {
				delList += strconv.Itoa(cArr[i].order) + ";"

				//				delete(cyclesArr, cArr[i].ID)
				WritecyclesArr(cArr[i].ID, nil)

			}
		}
	}
	return delList
}

////////////////////////////////////////////////////////////////////////////

/*
	если данный цикл не выполняет никакого действия,

а тупо крутится вхолостую, то замедлить его цркуляцию.
Вызывается в самом конце функции активации осознания, значит, прошло без пользы.
*/
func endIdleCycle(cycle *cycleInfo) bool {
	// такое возможно ли?
	/*
		 if functionsInAllCickles == nil{//ID запускаемых в текущем цикле инфо-функций нет
			 conscienceStatus+="<span style='color:red'>Шаг цикла ID="+ strconv.Itoa(cycle.ID)+" прошел без вызова инфо-функций. Цикл закрывается</span><br>"
			 //cycle.needFinish=true
			 endBaseIdCycle(cycle.ID)
			 return true
		 }*/
	// просто тормозим проход активации на 0.5 секунды чтобы с такой скоростью следил за событиями.
	//	time.Sleep(500 * time.Millisecond)//БЕСПОЛЕЗНО ПРИМЕНЯТЬ SLEEP т.к. он блокирует не только данную горутину, но и другие!!!!
	return true
	/*
		 // Уменьшенное числа холостых циклов по сравнению с limitMaxCycleStepCount:
			if cycle.count > limitMaxCycleStepCount/10  {
				conscienceStatus+="<span style='color:red'>Шаг цикла ID="+ strconv.Itoa(cycle.ID)+" прошел в холостую "+ strconv.Itoa(limitMaxCycleStepCount/10)+" раз. Цикл закрывается</span><br>"
		//		cycle.needFinish=true
				endBaseIdCycle(cycle.ID)
				return true
			}
			return false */
}

////////////////////////////////////////////////

/*
	  В момент создания нового цикла (func addCounterGoNext)
			определить его первоначальную значимость.

В ходе цикла значимость может меняться.
*/
func getCycleWeight() int {
	var w = 0
	if CurrentInformationEnvironment.veryActualSituation {
		w += 1
	}
	if CurrentInformationEnvironment.danger {
		w += 1
	}
	if problemExtremImportanceObject != nil {
		w += 1
	}

	return w
}

// ///////////////////////////////////////////////
// инфа для пульта по get_short_info(
func GetCycleLocInfo(cID int) string {
	//	c,ok:= cyclesArr[cID]
	c, ok := ReadecyclesArr(cID)
	if ok {
		return c.log
	}
	return "Нет цикла с ID=" + strconv.Itoa(cID)
}

// ///////////////////////////////////////
// инфа для пульта по function show_cyckle(id)
func GetCycleInfo(id int) string {
	out := ""
	cInfo := ""
	// точно ли есть такой живой цикл
	//	c,ok:= cyclesArr[id]
	c, ok := ReadecyclesArr(id)
	if ok {
		if c.isMainCycle {
			cInfo += "Это - главный, осознаваемый цикл мышления.<br>"
		}
		if c.dreaming {
			cInfo += "Режим пассивного размышления.<br>"
		}

		cInfo += "Значимость цикла: " + strconv.Itoa(c.weight) + "<br>"
		cInfo += "Тема мышления: " + strconv.Itoa(c.themeType) + "<br>"
		cInfo += "<span style='color:blue;cursor:pointer;' onClick='get_short_info(" + strconv.Itoa(c.ID) + ")'>Последовательность шагов</span><br>"
	} else {
		return "Нет такого Цикла мышления."
	}
	out = "Цикл мышления " + strconv.Itoa(c.count) + " шагов<br>"
	out += cInfo

	return out
}

////////////////////////////////////////////

// инфа о цикле
func GetCycleStrInfo(cID int) string {
	out := "Нет цикла с ID = " + strconv.Itoa(cID) + "<br>"
	//	cycle,ok:= cyclesArr[cID]
	cycle, ok := ReadecyclesArr(cID)
	if ok {
		out = "Цикл ID = " + strconv.Itoa(cID) + "<br>"
		if cycle.isMainCycle {
			out += "Это - главный цикл.<br>"
		}
		out += "Число звеньев цепочки: " + strconv.Itoa(cycle.count) + "<br>"
		out += "Значимость цикла: " + strconv.Itoa(cycle.weight) + "<br>"
		if cycle.dreaming {
			out += "Это цикл пассивного размышления.<br>"
		}
	}
	return out
}

//////////////////////////////////////////////////////

// в главном цикле - режим пассивного размышления
func IsDreamMainProcess() *cycleInfo {
	for _, v := range cyclesArr {
		if v == nil {
			continue
		}
		if v.dreaming {
			return v
		}
	}
	return nil
}

///////////////////////////////////////////////////////////////////

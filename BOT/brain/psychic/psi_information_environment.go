/* Информационная среда - основа текущего самоощущения

с каждым обновление инф.среды предыдущее запоминается в массиве InformationEnvironmentObjects
и еть OldInformationEnvironment
*/

package psychic

import (
	"BOT/brain/gomeostas"
	"BOT/lib"
	"strconv"
)

/////////////////////////////////////////////////////
/* Среда условий текущего состояния - интегративная информационная среда
Образ интегративной информационной среды Формируется временно и не сохранятся в файле.
Объекты InformationEnvironment накапливаются в оперативке, указалеи входят в  эпизодическую память и удаляются во сне.

Структуру информационной среды можно дополнять по мере необходимости, т.к. она не сохраняется в файле.

Заполняться начинает на уровне рефлексов.
*/
type InformationEnvironment struct {
	LifeTime   int  // момент создания кадра инф.окружения
	IsPsyLevel bool // true - данные записаны на стороне психики
	IsSleep    bool // true - организм спит (во сне контекст задает тоже InformationEnvironment) if IsSleeping {

	//устанавливается в func PsychicCountPuls, сбрасывается в func GetAutomotizmActionsString и в infoFunc25()
	IsIdleness100pulse bool // Не было действий Beast более 100 пульсов

	// общая оценка гомео-настроения
	Mood int //сила Плохо -10 ... 0 ...+10 Хорошо
	//ID парамктров гомеостаза как цели для улучшения в данных условиях
	curTargetArrID []int
	// текущая эмоция Emotion, может быть произвольно изменена
	PsyEmotionId int
	// опасность состояния
	danger              bool // получить из GetAttentionDanger
	veryActualSituation bool // оценка важности ситуации, необходиомсть срочных действий

	PsyMood int //Субъективно ощущаемая оценка, текущее осознаваемое настроение, которое можно произвольно изменять.

	// текущий образ сочетания действий с Пульта
	ActionsImageID          int // акция типа ActionsImage
	IsUnknownActionsImageID int //Образ, не имеющий модели понимания, вызывает беспокойство и активирует пассивный режим

	// текущий образ сочетания ОТВЕТНОГО действия мот.автоматизма
	AnswerImageID int // акция типа ActionsImage
	// текущий образ ментального автоматизма
	//MentalAutomatizmID int    //  типа MentalAutomatizm

	/*Это - период ожидания ответа с пульта на действие
	чтобы инфо-функции знали, работают ли они на Стимул или произвольно.
	В случае произвольности размышления о наиболее значащем могут давать действия.
	*/
	IsWaitingPeriod bool // это - период ожидания ответа с пульта на действие.

	/*наиболее важный образ типа extremImportance - хранится в обычном массиве curImportanceObjectArr []extremImportance
	Можно по extremImportance найти importance: getImportanceFromExtremImportance(Stimul int, kind int)
	*/
	ExtremImportanceObjectID int
	// актуальнsq по эффекту Правило, выделенное из ExtremImportanceObjectID в ходе перебора infoFunc5()
	ActualEpisodicMemoryID int
	// текущая Доминанта нерешенной проблемы
	DominantaID int

	//Нужно подумать о проблеме автоматизма или проявить инициативу, в общем, запустить func infoFunc25()
	needThinkingAboutAutomatizm bool

	// true - текущий стмул навязывает то, что не соотвествует текущей Теме и Цели
	isStimulToForce bool
}

/*
	с каждым обновлением инф.среды предыдущее запоминается в массиве

InformationEnvironmentObjects и еще еть OldInformationEnvironment
Кратковременная память кадров ИЕ. В файл не записывается, освобождается во сне.
TODO: Позволяет удобно вспомнить, что было недавно, использовать для выбора сновидения, темы общения и для мечтаний.
В инфо-функции infoFunc10() позволяет вспомнить самое значимое, что было для начала размышения (кроме возврата на прерванное).
*/
var InformationEnvironmentObjects []*InformationEnvironment

func GetInformationEnvironmentObjectsLength() int { // для сна
	return len(InformationEnvironmentObjects)
}

var CurrentInformationEnvironment InformationEnvironment

func initCurrentInformationEnvironment() {

	CurrentInformationEnvironment.ActionsImageID = 0
	CurrentInformationEnvironment.IsUnknownActionsImageID = 0
	CurrentInformationEnvironment.LifeTime = LifeTime

	CurrentInformationEnvironment.ActualEpisodicMemoryID = 0
	CurrentInformationEnvironment.ExtremImportanceObjectID = 0
}

var OldInformationEnvironment InformationEnvironment

//////////////////////////////////////////////////////////////////////
/* создать новй кадр IE из OldInformationEnvironment
и записать его адрес в InformationEnvironmentObjects
*/
func saveOldIE() {
	var ie InformationEnvironment // создан новый реальный объект инфосреды, от которых нужно освобождать память
	ie = OldInformationEnvironment
	InformationEnvironmentObjects = append(InformationEnvironmentObjects, &ie)
	OldInformationEnvironment = CurrentInformationEnvironment
	CurrentInformationEnvironment = ie
}

///////////////////////////////////////

///////////////////////////////////////////////////////
/*  получение текущего состояния информационной среды: отражение Базового состояния и Активных Базовых контекстов
только при ориентировчном рефлексе и осмыслении результатов - обновление самоощущения!
*/
func GetCurrentInformationEnvironment() {
	saveOldIE()

	// !!! 	initCurrentInformationEnvironment() НИЧЕГО НЕ СБРАСЫВАТЬ, ТОЛЬКО ПЕРЕКРЫВАТЬ НОВЫМИ ЗНАЧЕНИЯМИ!

	CurrentInformationEnvironment.LifeTime = LifeTime // момент обновления

	CurrentInformationEnvironment.IsSleep = IsSleeping

	// определение текущего сочетания ID Базовых контекстов - оно есть всегда, даже если ничего не сделано на Пульте - нулевое сочетание.
	bsIDarr := gomeostas.GetCurContextActiveIDarr()
	// текущая эмоция Emotion, может быть произвольно изменена
	CurrentInformationEnvironment.PsyEmotionId, _ = createNewBaseStyle(0, bsIDarr, true)

	CurrentInformationEnvironment.veryActualSituation, CurrentInformationEnvironment.curTargetArrID = gomeostas.FindTargetGomeostazID()

	//ActID:=action_sensor.CheckCurActionsContext()  модификация - только в automatism_tree.go при акции с пульта!
	//CurrentInformationEnvironment.ActionsImageID,_=createNewlastActivityID(0,ActID,true)// текущий образ сочетания действий с Пульта Activity

	if lastAutomatizmRun != nil {
		CurrentInformationEnvironment.AnswerImageID = lastAutomatizmRun.ActionsImageID
	}

	CurrentInformationEnvironment.danger = GetAttentionDanger()
	CurrentInformationEnvironment.Mood = GetCurMood()
	CurrentInformationEnvironment.PsyMood = PsyMood

	writeInformationEnvironmentMarker()
	return
}

///////////////////////////////////////////////////////

// /////////////////////////////
// обновление состояния информационной среды
func refreshCurrentInformationEnvironment() {
	//return  // инфа просто перекрывается новой
	///////// Информационная среда осознания ситуации
	// Нужно собрать всю информацию, которая может повлиять на решение.
	//  получение текущего состояния информационной среды: отражение Базового состояния и Активных Базовых контекстов
	GetCurrentInformationEnvironment()

	// оценка опасности ситуации, необходиомсть срочных действий
	veryActualSituation = CurrentInformationEnvironment.veryActualSituation
	// выявить ID парамктров гомеостаза как цели для улучшения в данных условиях
	curTargetArrID = CurrentInformationEnvironment.curTargetArrID

	/* Еще информация:
	жизненный опыт  psy_Experience.go
	доминанта psy_problem_dominanta.go
	субъектиная оценка ситуации для применения произвольности
	*/

	// актуальной инфой являются узлы активной ветки дерева понимания, особенно контекст SituationID
}

///////////////////////////////////////////////

// ////////////////////////////////////////////////////
// записать метку изменения information_environment при каждом обновлении
func writeInformationEnvironmentMarker() {
	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/self_perception_count.txt")
	var old = 0
	if strArr != nil {
		old, _ = strconv.Atoi(strArr[0])
	}
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/self_perception_count.txt", strconv.Itoa(old+1))
}

//////////////////////////////////////////////////////

/* Информационные функции, вызываемые как действия ментального автоматизма по их ID функции.
Так же они могут вызываться непосредсвтенно из процессов функции осознания.

Инфо-функции - разные методы получения инфы, систематизации, поиска и т.п.
с целью найти верное действие для моторного автоматизма, а если нет,
то создания нового ментального автоматизма для продолжения итеации поиска.

У инфо-функций не должно быть аргумента, иначе невозможно будет их вызывать из runMentalFunctionID(id int)
Поэтому в инфо-функции могут вызываться вспомогательные функции с аргументами, полученными в инфофункци
которые вызываются только если есть нужная инфа, например, сохраненная в mentalInfoStruct

Результат работы инфо-функции записывается в mentalInfoStruct

________________________________ПРИМЕР
В цикле мышления нужно запустить имеющийся мот.автоматизм с aID.
Где-то можно просто: infoFunc17()//запустить моторный автоматизм mentalInfoStruct.motorAtmzmID
или
//С использованием образа запуска моторного автоматизма:
maID,ma:=CreateNewlastMentalActionsImagesID(0,6,aID,true)
и тогда infoFunc2()//Подобрать ActionsImage для последующего звена цепочки
видя, что до него есть такой образ, может решить запустить его:
infoFunc7()//запустить моторный автоматизм mentalInfoStruct.motorAtmzmID

Для новых мот.автоматизмов:
infoFunc7()//создать новый моторный автоматизм по действию ActionsImageID
затем
infoFunc7()//запустить моторный автоматизм mentalInfoStruct.motorAtmzmID

TODO Предствоит добавить "размышление" о значимых объектах -
проход по Правилам чтобы посмотреть, чем это может кончиться
и при этом записать новые превила (переработка информации в свободное время).
То, что должно происходить и в сновидениях, но не ограничивая уровни осознания.
Это уже создаст немало шагов мышления в цикле.

Возвращаемое значение нужно для определения cycle.idle

ПРО ОТЗЕРКАЛИВАНИЕ или Как формируются новые автоматизмы с неизвестными ранее действиями?
Отзеркаливание - единственный способ получения автоматизмов с новыми фразами и действиями.
Могут быть следующие случаи.
1. Каждый стимул с Пульта, пришедший в ответ на совершенное действие, создает учительское Правило в эпиз.памяти. Отсюда сразу возникает то, чем нужно отвечать в аналогичных условиях.
2. Стимул с Пульта пришел не связанный с действиями Beast (более 60 пульсов от последнего действия).
В таком случае может быть тут же опробовано "попугайство" (func infoFunc13 В отличии от молчаливогоо отзеркаливания в infoFunc11() для 4-й стадии), что покажет Оператору, что Beast хочет оценить возможный эффект.
3. На 4-й стадии нужно отзеркаливать только то, что в действиях оператора достигает поставленной цели (func infoFunc11 ).
Отзеркаленную фразу можно поначалу испоьзовать как есть, без изменений, но если они в новых условиях дает нежелательный эффект, то она может быть "исправлена" так, чтобы содержать только слова, которые раньше приводили к позитивному эффекту в последовательности кадров эпиз.памяти.
*/

package psychic

import (
	"BOT/brain/gomeostas"
	"BOT/brain/transfer"
	wordSensor "BOT/brain/words_sensor"
	"BOT/lib"
	"strconv"
)

// /////////////////////////////////////////////////////////////////////////
/*
	Дополнительное инфо-окружение. Оперативная память цикла осмысления.

Общая для всех информационных функций структура (типа информацонного окружения)
для сохранения найденной информации.
*/
type mentalInfo struct {
	ActionsImageID int //ID ActionsImage моторного действия (стимула или ответа)
	/* AmtzmNextStringID - наращивание цепочки последовательности действий
	   	для произвольного выполнения или прикрепления к автоматизму.
	   Можно произвольно набирать и корректировать AmtzmNextStringID.
	   Можно произвольно выполнять отдельные ActionsImageID и смотреть результат
	   	чтобы составлять цепочку AmtzmNextStringID.

	   !AmtzmNextStringID - ID без добавленного prefixActionIdValue!
	*/
	AmtzmNextStringID int // последовательность образов действий AmtzmNextString.ID

	toAutmtzmActionsImageID int  //в infoFunc2() создать автоматизм infoFunc7() по ID ActionsImage моторного действия (стимула или ответа)
	noStaffAtmzmID          bool // true - нет шаштного моторного автоматизма
	motorAtmzmID            int  // ID моторного автоматизма
	motorAtmzmBlockedID     int  // ID заблокированного как опасного моторного автоматизма
	runMotorAtmzmID         int  // ID моторного автоматизма который нужно запустить в infoFunc2()
	//var runningMotAutmtzmID=0 - запущенный в infoFunc17() моторный автоматизм
	//mentalAtmzmID int // ID ментального автоматизма
	ThemeImageType   int  // ТИП актуальной темы размышления
	mentalPurposeID  int  // ID ментальной цели, альтернативной текущей  PurposeImage
	notOldAutomatizm bool // true - НЕ позволить запустить рвущийся на выполнение старый автоматизм
	runInfoFuncID    int  // запуск инфо-функции
	epizodFrameIndex int  // ID успешной инфо-функции со все большим отклонением от условий - по ментальным правилам
	volutionReload   int  // для функ 10: 0 не найдены параметры перезапуска дерева, 1 - найдены
	// для произвольной переактивации дерева ситуации infoFunc14():
	moodeID     int
	emotonID    int
	SituationID int
	fromNextID  int // передача инфы о номере шага итерации
	ExtremObjID int // extremImportanceObject.extremObjID - при возврате к прерванному объекту мышления

	DominantaID           int // ID доминанты для коотрой переданы науденные действия
	DominantSuccessAImgID int // найдено успешное действие
	DominantSuccessValue  int // оцененная успешность: 1 - доминанта решена, 2 - не точно, 3 - еще менее точно
	prognoseEffect        int // <0 - плохой прогноз цепочки правил, >0 - хороший прогно(з начиная с действия рвущегося автоматизма)

	noOperatorStimul bool // не было стимула от оператора > 10 сек при значительном изменении гомео-параметров - режим провокации: бот пытается вызвать реакцию оператора

}

var mentalInfoStruct mentalInfo

func clinerMentalInfo() {
	mentalInfoStruct.ActionsImageID = 0
	mentalInfoStruct.AmtzmNextStringID = 0
	mentalInfoStruct.toAutmtzmActionsImageID = 0
	mentalInfoStruct.runMotorAtmzmID = 0
	mentalInfoStruct.noStaffAtmzmID = false
	mentalInfoStruct.motorAtmzmBlockedID = 0
	mentalInfoStruct.motorAtmzmID = 0
	//mentalInfoStruct.mentalAtmzmID=0
	// никогда не очищать Цель! только перекрывать новой	mentalInfoStruct.mentalPurposeID=0 Иначе просто зацикливается
	mentalInfoStruct.notOldAutomatizm = false
	mentalInfoStruct.runInfoFuncID = 0 // запуск инфо-функции
	mentalInfoStruct.epizodFrameIndex = 0
	// для произвольной переактивации дерева ситуации infoFunc14():
	mentalInfoStruct.moodeID = 0
	mentalInfoStruct.emotonID = 0
	mentalInfoStruct.SituationID = 0
	mentalInfoStruct.volutionReload = 0
	mentalInfoStruct.ExtremObjID = 0

	mentalInfoStruct.DominantaID = 0
	mentalInfoStruct.DominantSuccessAImgID = 0
	mentalInfoStruct.DominantSuccessValue = 0
	mentalInfoStruct.prognoseEffect = 0

	mentalInfoStruct.noOperatorStimul = false

	// ВСТАВЛЯТЬ ДРУГИЕ ЧЛЕНЫ ПО МЕРЕ ПОЯВЛЕНИЯ !!!!

}

/*
	произвольно активированные параметры, определяются при замуске ментального автоматизма.

Держатся на время, пока не изменятся генетически определенные соотвествующие параметры или

	если активация была в данном пульсе
*/
var mentalMoodVolitionID = 0        // произвольно активированное настроение
var mentalMoodVolitionPulsCount = 0 // произвольно активированное настроение

var mentalEmotionVolitionID = 0        // произвольно активированная эмоция
var mentalEmotionVolitionPulsCount = 0 // произвольно активированная эмоция

var mentalSituationVolitionID = 0        // произвольно активированная ситуация
var mentalSituationVolitionPulsCount = 0 // произвольно активированная ситуация

//var mentalPurposeImageID=0// призвольно активированная цель
//var mentalPurposeImagePulsCount=0// призвольно активированная цель

var runningMotAutmtzmID = 0 // запущенный в infoFunc17() моторный автоматизм - для MentalAutomatizm.motAutmtzmID

func GetgetInfoFuncInfoStr(fID int) string {
	return getMentalFunctionString(fID)
}

func getMentalFunctionString(id int) string {
	switch id {
	case 1:
		return "Подготовить продолжение цикла осмысления"
	case 2:
		return "Выбрать о чем и как подумать, с учетом предыдущих шагов цикла"
	case 3:
		return "Найти подходящую инфо-функцию по опыту ментальных Правил"
	case 4:
		return "Находит и запускает мот.автоматизм в плохой ситуации"
	case 5:
		return "Идет цикл мышления об экстремально важном объекте extremImportanceObject"
	case 6:
		return "Подвергнуть сомнению автоматизм, если нет опасности (не нужно реагировать аффектно) и ситуация важна"
	case 7:
		return "Создать новый моторный автоматизм по действию ActionsImageID"
	case 8:
		return "Ментальное определение ближайшей Цели в данной ситуации"
	case 9:
		return "Найти способ улучшения значимости объекта внимания extremImportanceObject"
	case 10:
		return "Вспомнить самое значимое, что было для начала размышения (кроме возврата на прерванное) с произвольной переактивацией дерева"
	case 11:
		return "Ментальное отзеркаливание действия оператора"
	case 12:
		return "Cинтез новой фразы из компонентов, имеющих плюсы в Правилах"
	/*case 122:
	return "Выбрать ID действия имеющего плюсы в Правилах"
	*/
	case 13:
		return "Тупое повторение Стимула оператора. Попугайство."
	case 14:
		return "Ментально переактиваровать дерево понимания с заданными узлами"
	case 15:
		return "Для условия дерева автоматизмов (NodeAID) в одиночных Правилах выбираем наилучшее"
	case 16:
		return "Случайно выдать любую известную фразу и затем infoFunc7() + infoFunc17()"
	case 17:
		return "Запустить моторный автоматизм mentalInfoStruct.motorAtmzmID"
	case 18:
		return "Ощущение наиболее значимого из параметров CurrentInformationEnvironment"
	case 19:
		return "Было эвристическое озарение и найдено mentalInfoStruct.ActionsImageID для пробы действия"
	case 20:
		return "Запустить процесс осмысления НАЗАД эпиз.памяти."
	case 202:
		return "Запустить процесс осмысления ВПЕРЕД эпиз.памяти."
	case 21:
		return "Срочно найти какое-то подходящее действие и запустить его"
	case 22:
		return "Добавление нового действия mentalInfoStruct.ActionsImageID в цепочку действий автоматизма Automatizm.NextID"
	case 23:
		return "Создание новой цепочки произвольных действий mentalInfoStruct.AmtzmNextStringID\nс добавлением mentalInfoStruct.motorAtmzmID"
	case 24:
		return "Добавление нового звена mentalInfoStruct.ActionsImageID в цепочку произвольных действий mentalInfoStruct.AmtzmNextStringID"
	case 25:
		return "Посмотреть условия чтобы проявить инициативу"
	case 26:
		return "Запуск на выполнение на пульте имеющейся цепочки произвольных действий mentalInfoStruct.AmtzmNextStringID"
	case 27:
		return "Создание автоматизма на основе имеющейся цепочки произвольных действий mentalInfoStruct.AmtzmNextStringID"
	case 28:
		return "фантазия - как функция интенсивного подбора ассоциаций для решения доминанты"
	case 29:
		return "Более внимательно рассмотреть ситуацию с Правилами: ментальными и связанными с ними моторными"
	case 30:
		return "Анализировать длинную или недоопредеелнную фразу и найти обобщающие данные"
	case 300:
		return "найти в неопределенной фразе что-то ранее понятное и выполнить действия с понятным стимулом"
	case 31:
		return "Провокационные действия, когда нет стимула от оператора, но нужно что-то делать"

		//	case NN: retrun "Решение по опыту закрытых доминант"

		/* !!! кроме этого есть специализированные инфо-функции для поддержки пассивного режима мышления,
		определенные в dreaming.go:
		func GotoDreaming
		func gotoPassiveMaind
		func epizodicGPTprognoze

		*/
	}
	return "Нет функции с ID = " + strconv.Itoa(id)
}

//////////////////////////////////////////////////////////

/*
	Функция вызова пронумерованной функции

Нужна для вызова случайной функции и т.п.
*/
func runMentalFunctionID(c *cycleInfo, id int) bool {
	switch id {
	case 1:
		infoFunc1(c)
		return true //Подготовить продолжение цикла осмысления
	case 2:
		infoFunc2(c)
		return true //выбрать о чем подумать, с учетом предыдущих шагов цикла
	case 3:
		infoFunc3(c)
		return true //найти подходящее мент.действие по опыту ментальных Правил
	case 4:
		infoFunc4(c)
		return true //Находит и запускает мот.автоматизм в плохой ситуации
	case 5:
		infoFunc5(c)
		return true //Идет цикл мышления об экстремально важном объекте extremImportanceObject
	case 6:
		infoFunc6(c)
		return true //ПОДВЕРГНУТЬ СОМНЕНИЮ моторный автоматизм, если нет опасности (не нужно реагировать аффектно) и ситуация важна
	case 7:
		infoFunc7(c)
		return true //создать новый моторный автоматизм по действию ActionsImageID
	case 8:
		infoFunc8(c)
		return true //Ментальное определение ближайшей Цели в данной ситуации
	case 9:
		infoFunc9(c)
		return true //найти способ улучшения значимости объекта внимания extremImportanceObject
	case 10:
		infoFunc10(c)
		return true //вспомнить самое значимое, что было для начала размышения (кроме возврата на прерванное) по InformationEnvironmentObjects []*InformationEnvironment
	case 11:
		infoFunc11(c)
		return true //Ментальное отзеркаливание действия оператора
	case 12:
		infoFunc12(c)
		return true //Cинтез новой фразы из компонентов, имеющих плюсы в Правилах
	/*case 122:
	infoFunc122(c)
	return true //выбрать ID действия имеющего плюсы в Правилах
	*/
	case 13:
		infoFunc13(c)
		return true //Тупое повторение Стимула оператора. Попугайство.
	case 14:
		infoFunc14(c)
		return true //ментально переактиваровать дерево понимания с заданными узлами
	case 15:
		infoFunc15(c)
		return true //Для условия дерева автоматизмов (NodeAID) в одиночных Правилах выбираем наилучшее
	case 16:
		infoFunc16(c)
		return true //Случайно выдать любую известную фразу или действие и затем infoFunc7() + infoFunc17()
	case 17:
		infoFunc17(c)
		return true //запустить моторный автоматизм mentalInfoStruct.motorAtmzmID
	case 18:
		infoFunc18(c)
		return true //Ощущение наиболее значимого из параметров CurrentInformationEnvironment
	case 19:
		infoFunc19(c)
		return true // было эвристическое озарение и найдено mentalInfoStruct.ActionsImageID для пробы действия
	case 20:
		infoFunc20(c)
		return true // выбрать кадр эпиз.памяти dreamingEpisodeHistoryID и проходить назад
	case 202:
		infoFunc202(c)
		return true // выбрать кадр эпиз.памяти dreamingEpisodeHistoryID и проходить вперед
	case 21:
		infoFunc21(c)
		return true // Срочно найти какое-то подходящее действие и запустить его
	case 22:
		infoFunc22(c)
		return true // Добавление нового действия в цепочку действий автоматизма Automatizm.NextID
	case 23:
		infoFunc23(c)
		return true //Создание новой цепочки произвольных действий mentalInfoStruct.AmtzmNextStringID с добавлением mentalInfoStruct.motorAtmzmID
	case 24:
		infoFunc24(c)
		return true //добавление нового звена mentalInfoStruct.motorAtmzmID в цепочку произвольных действий mentalInfoStruct.AmtzmNextStringID
	case 25:
		infoFunc25(c)
		return true //Посмотреть условия чтобы проявить инициативу
	case 26:
		infoFunc26(c)
		return true //Запуск на выполнение на пульте имеющейся цепочки произвольных действий mentalInfoStruct.AmtzmNextStringID
	case 27:
		infoFunc27(c)
		return true //Создание автоматизма на основе имеющейся цепочки произвольных действий mentalInfoStruct.AmtzmNextStringID
	case 28:
		infoFunc28(c)
		return true //фантазия - как функция интенсивного подбора ассоциаций для решения доминанты
	case 29:
		infoFunc29(c)
		return true //Более внимательно рассмотреть ситуацию с Правилами: ментальными и связанными с ними моторными
	case 30:
		infoFunc30(c)
		return true //Анализировать длинную или недоопредеелнную фразу и найти обобщающие данные
	case 300:
		infoFunc300(c)
		return true //найти в неопределенной фразе что-то ранее понятное и выполнить действия с понятным стимулом
	case 31:
		infoFunc31(c)
		return true //"Провокационные действия, когда нет стимула от оператора, но нужно что-то делать"
	default:
		return false
	}

	return false
}

//////////////////////////////////////////////////

//////////////////////////////////////////////////////////
/* далее идут ПРОНУМЕРОВАННЫЕ ИНФОРМАЦИОННЫЕ ФУНКЦИИ,
для которых в mental_automatizm_INFO_structs.go определяются ИНФОРМАЦИОННЫЕ ГЛОБАЛЬНЫЕ СТРУКТУРЫ - для
передачи в них полученной информации.
Так же для передачи информации в инфо-функции (если это нужно, например, что найти) применяюися входне структуры.
*/
//////////////////////////////////////////////////////////

/*
	НЕ НУЖНО БОЛЬШЕ

1. Подготовить ПРОДОЛЖЕНИЕ (а не новый) цикла осмсысления.
Вызывается только из consciousnessElementary

Новый цикл делается в fromNextID := createNewCycleIteration()
*/
func infoFunc1(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	setCurIfoFuncID(c, 1)
	/* может вызываться подряд при пустом цикле
	if c.lastFuncID == 1{// не вызывать, если только что было
		return
	}*/
	c.lastFuncID = 1
	//////////////////////

	//	runningMotAutmtzmID=0
	//	dreamingEpisodeHistoryID=0// прервать размышление, сновидение
	//	curActiveActions=nil // новый шаг мышления не начинать с объекта восприятия curActiveActions

	// продолжить циркуляцию через 1 сек
	//if true {// true - бесконечные параллельные циклы вызывают паники по разным случаям, нужно вылавливать...
	//	sleepCycle(c.ID) // задержать перезапуск на 1 сек
	//}

	//	reloadConsciousness(c,fromNextID)
	return true
}

/////////////////////////////////////////////////////////
/*
func sleepCycle(runID int) {
	cChannelbig := make(chan int,1)

	go func() {
		cChannelbig <- runID

		timer := time.NewTimer(1000 * time.Millisecond)
		select {
		case <-timer.C:
//			consciousnessElementary(2, runID)
		}

		close(cChannelbig) // закрываем канал
	}()
}*/
////////////////////////////////////////////////////////

////////////////////////////////////////////////////
/* №2 - выбрать о чем подумать, с учетом предыдущих шагов цикла.

Уже не были выбраны ментальные Правила, НУЖНО ПРИДУМАТЬ ЧТО-ТО НОВОЕ,
в том числе новый вариант продолжения мент-х Правил при предшествующей коррекции цикла мент.автоматизма.

Здесь выбор той или иной функции делается в контексте имеющейся инфо-среды,
а если этого не удается, то выбирается случайно одна из инфо-функций.

Если удалось найти решение в виде подходящео действия, то тут же оно и запускается.

Не обязательно - ответ на Стимул, а может быть собственная ИНИЦИАТИВА при дреме
(начинается со случайных действий infoFindRundomMentalFunction()).

Возвращаемое значение нужно для определения cycle.idle
*/
func infoFunc2(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 2 { // не вызывать, если только что было
		return false
	}
	if motorActionEffect > 0 && !c.dreaming { // не нужно искать решение, раз уже был нормальный ответ
		return false
	}
	setCurIfoFuncID(c, 2)
	c.log += "Запущена Инфо-функция № 2: <i>\"" + getMentalFunctionString(2) + "\"</i><br>"
	return mentalSimpleReflexSolution(c)
}
func mentalSimpleReflexSolution(c *cycleInfo) bool {

	if c.isWaitingPeriod { // не совершать действий, а ждать ответа
		//TODO думать о сделанном только что, если нет полной уверенности, и успеть еще раз ответить до ответа с пульта?
		// хотя это должен быть личный опыт... лучше не нагромождать.
		return false
	}
	// нельзя прерывать главный цикл, если для него установлено ожидание
	val, ok := ReadecyclesArr(activedCyckleID)
	if !ok && val != nil {
		if val.isWaitingPeriod == true {
			return false
		}
	}

	// infoFunc31(c) // тестирование

	/* не было стимула от оператора > 10 сек при значительном изменении гомео-параметров
	и при этом ситуация атасная или нужно найти действие

	TODO что-то не срабатывает infoFunc31(c) - нет сочетания условий поэтому сделано еще в func PsychicCountPuls
	*/
	if mentalInfoStruct.noOperatorStimul {
		if isNeedForCommunication() { // нужно провоцировать оператора
			if infoFunc31(c) {
				return true
			}
		}
	} else { // есть стимул, но нет автоматизма

		if mentalInfoStruct.noStaffAtmzmID {
			// в спокойной ситуации
			if !CurrentInformationEnvironment.veryActualSituation && !CurrentInformationEnvironment.danger {

				if isUnrecognizedPhraseFromAtmtzmTreeActivation { //при активации была нераспознанная фраза
					/*Если фраза не понята, то в порядке слов смотреть, на что похожа цепочки слов.
					При выборке сделать массив слов фразы и в каждом кадре - тоже массив слов.
					Это проще, чем распознавание по мере появления новых слов.

					"привет тебе!"
					*/
					if EvolushnStage <= 4 {
						return infoFunc300(c)
					} // для EvolushnStage > 4 есть infoFunc30(c) - Анализировать длинную или недоопределенную фразу, найти обобщающие данные и что-то сделать

				}

			}

		}
	}

	//Нужно подумать о проблеме автоматизма или проявить инициативу, в общем, запустить func infoFunc25()
	if CurrentInformationEnvironment.needThinkingAboutAutomatizm {
		//Более внимательно рассмотреть ситуацию с Правилами: ментальными и связанными с ними моторными
		if !infoFunc29(c) {
			infoFunc25(c)
		}
		return true
	}

	//если есть действие - создать автоматизм infoFunc7() по ID ActionsImage моторного действия (стимула или ответа)
	if mentalInfoStruct.toAutmtzmActionsImageID > 0 {
		//		c.log+="в infoFunc2() ID образа действий, который нужно запустить в автоматизме: "+strconv.Itoa(mentalInfoStruct.toAutmtzmActionsImageID)+"<br>"
		if c.isMainCycle {
			runAutomatizmAfterCheck(c)
		} else {
			insight(c)
		}
		return true
	}
	// ID моторного автоматизма который нужно запустить в infoFunc2()
	if mentalInfoStruct.runMotorAtmzmID > 0 {
		//	c.log+="в infoFunc2() ID моторного автоматизма который нужно запустить: "+strconv.Itoa(mentalInfoStruct.runMotorAtmzmID)+"<br>"
		// подготовить инфу для infoFunc6() из инфы от infoFunc7():
		mentalInfoStruct.motorAtmzmID = mentalInfoStruct.runMotorAtmzmID
		infoFunc6(c) // проверить и если норм. - запустить, сделав штатным
		// на этом и закончить данный цикл
		//endCurrentCyckle()
		return true
	}

	if (curActions.PhraseID != nil && len(curActions.PhraseID) > 1) || isUnrecognizedPhraseFromAtmtzmTreeActivation {
		if EvolushnStage > 4 { //5-я сталия развития. Творчество.
			//Анализировать длинную или недоопределенную фразу, найти обобщающие данные и что-то сделать
			return infoFunc30(c)
		}

	}

	if mentalInfoStruct.epizodFrameIndex > 0 { // было найдено сделующее звено мышления в infoFunc3()
		c.log += "в infoFunc2() было найдено сделующее звено мышления в infoFunc3()<br>"
		// mentalInfoStruct.epizodFrameIndex это - номер инфо-функции
		// нужно перезапустить осмысление в новым шагом
		funcID := mentalInfoStruct.epizodFrameIndex
		mentalInfoStruct.epizodFrameIndex = 0
		mentalInfoStruct.epizodFrameIndex = 0
		if runMentalFunctionID(c, funcID) {
			return true
		}
	}

	/* Если есть mentalInfoStruct.motorAtmzmID > 0 но автоматизм сразу не выполнен,
	   то нужно ПОДВЕРГНУТЬ СОМНЕНИЮ автоматизм и если нормально, то выполнить
	*/
	if mentalInfoStruct.motorAtmzmID > 0 { //
		/*Если заблоикрован mentalInfoStruct.motorAtmzmID и найден новый, то он СРАЗУ ЗАПУСКАТЕСЯ.
		  Если mentalInfoStruct.motorAtmzmID нормальный, то он так же сразу запускается.
		*/
		infoFunc6(c)

		//	c.log+="в infoFunc2() сразу запускается mentalInfoStruct.motorAtmzmID="+strconv.Itoa(mentalInfoStruct.motorAtmzmID)+" <br>"

		mentalInfoStruct.motorAtmzmID = 0
		// на этом закончить данный цикл
		//endCurrentCyckle()
		return true
	}
	/////////////////////////////////

	//!!!! clinerMentalInfo() не терять Тему и Цель!!!!!

	// здесь ищем Какое действие нужно совершить, не только инфо-функции

	if atmtzmActualTreeNode == nil || // нет вообще автоматизма
		atmtzmActualTreeNode.Usefulness < 0 { // запущенный автоматизм заблокирован (м.б. не штатный)
		///////////////////////  поиск решения в мент.Правилах и Моделях понимания
		/* при кажой неудаче предыдушего уровня перезапуск осмыления и более глубокий поиск.
		 */

		/*  ПОИСК РЕШЕНИЯ В МЕНТАЛЬНЫХ ПРАВИЛАХ tryRulesLevel(c,refinementLevel) для данного detectedActiveLastProblemNodID
		Мент.Правила, в конечном счете - последовательность infoFuncSequence []int
				- последовательность ID выполненных инфо-функций одной активации consciousnessElementary()
		Тут нет конкретного образа действия, но есть ID инфо-функций, которые приводили к достижению Цели,
				но С последней ф-цией может быть связан и MentalAutomatizm.motAutmtzmID
				По действию - аналогично func infoFunc3
		*/
		/*funcID := findSuitableMentalFunc()
		if funcID!=nil {
			// выбираем первую функцию TODO подумать
			if runMentalFunctionID(c,funcID[0]){
				return true
			}
		}*/
		fID := findInfoIdFromExperience(true)
		if fID > 0 {
			if runMentalFunctionID(c, fID) {
				return true
			}
		}

		// поиск в Модели понимания
		/*Использовать модели объектов для нахождения действий ПО АНАЛОГИИ с теми действиями,
					что были в эпиз.памяти для данного объекта.
		Леонардо да Винчи смотрел на штукатурку чтобы в голову пришла идея.
		*/
		//
		resM := tryModelsLevel(c)
		if resM { // найдено решение
			c.log += "Удачный поиск в Модели понимания в infoFunc6() будет конец осмысления из-за запуска мот.авто-ма<br>"
			return true
		}
		//}
	}
	//		} //if refinementLevel<6{
	//	}

	// Правила и Модели не сработали, ИДЕМ ДАЛЬШЕ.
	// tryRandExperimtntalAction() дальше неплохо прописано вроде

	/*Далее - быстрое нахождение решения, без сложных размышлений. Размышлять будет когда есть время
	в processingFreeState(stopMentalWork) // как во сне - обработка структур в свободном
	} */
	////////////////////////////////////////

	//	infoFunc13();return true // тестивароние попугайского отзеркаливания

	/* ГЛАВНОЕ - ПРИВЛЕЧЕНИЕ ОСОЗНАННОГО ВНИМАНИЯ к наиболее значимому,
	 есть ли актуальный объект внимания с отрицательной значимостью.
	Информационная функция "Понимание объекта восприятия" - выборка данного образа восприятия в дереве с прослеживанием,
	что оно означало в разных условиях. т.к. образ включает в себя все составляющие объекта восприятия, то он - обобщение,
	а его понимание - Вид всех последствий в разных условиях.
	Для образа фразы type Verbal struct - составляющие отдельные слова, которые могут быть в разных образах фраз,
	так что можно сделать функцию Вида - выборки данного слова в разных условиях с последствиями.
	*/
	if extremImportanceObject != nil {
		// найти способ улучшения значимости объекта extremImportanceObject - запуск моторного автоматизма

		if infoFindAttentionObjImprovement(c) { //  infoFunc9() сделано, но еще не тестировалось
			c.log += "в infoFunc2() найден способ улучшения значимости объекта extremImportanceObject - запуск моторного автоматизма<br>"
			/*при наличии mentalInfoStruct.ActionsImageID > 0
			создать автоматизм (если такого еще нет),
			проверить его и запустить.
			*/
			if c.isMainCycle {
				runAutomatizmAfterCheck(c)
			} else {
				insight(c)
			}
			return true // больше не искать, уже создан мент.автоматизм объективного действия
		}
	}

	if EvolushnStage > 4 { //5-я стадия развития. Творчество.
		/* В тяжелом ментальном состоянии сделать попытку улучшить настроение, эмоцию и ситуацию произвольной переактивацией
		чтобы выйти на решения в лучших условиях. Но нет опасности и можно подумать.
		Произвольная переактивация через func infoFunc14
		*/
		if CurrentMood < 5 && !CurrentInformationEnvironment.danger {
			// прикинуть как переактивировать по состоянию гомеопараметров
			if findReactivadeParams(c) {
				return true // больше не искать
			}
		}

		/*  текущий субъект внимания (наиболее нзачимое в мыслях)
			TODO ЕЩЕ НЕТ ПОДДЕРЖКИ объектов собственных мыслей по аналогии с importance.go И НЕ НУЖНО
			var extremImportanceMentalObject extremImportance
		Для processingFreeState(stopMentalWork) // как во сне - обработка структур в свободном ?
			Здесь должно находиться решение (выбор действия) на основе использования информации
			о наиболее важном в мыслях, скорее всего связанное с Доминантами.

			if infoFindAttentionObjMentalImprovement(){
				return // больше не искать, уже создан мент.автоматизм объективного действия
			}*/
		///////////////////////////////////////

		/*если для данного сочетания Стимул-Ответ есть только один вид эффекта,
		то это - уже Информация, и чем больше опыт (количество обобщенных правил), тем такая информация полезнее.
		infoFunc15()
		*/
		//getDominantEffect(triggerID int, actionID int)

		/*Если оператор нажал кнопку Учитель, это - стимул привлечения внимания для наблюдения за ним,
		за достигаемой им целью и как он ее достигает. Это - начиная с 4-й стадии развития.
		Для понимания отзеркаленных действий оператора цель, которую ставил тварь перед своими действиями PurposeGenetic -
		фксируется в узлах дерева понимания PurposeImage -> SituationImage .
		Кнопки действий теперь активируют смысл (контекст) ситуации SituationImage
		*/
		// TODO

		/*Посмотреть, есть ли другие ветки у данного узла и использовать их начальную инфу.
		В ходе цикла осмысления могут быть переключения на другие условия goNext.MotorBranchID деревьев (перезапуск осмысления)
		при том же контексте Цели (сначала PurposeImage -> SituationImage). Попробовать решения других веток.
		М.б. применить goNextFromUnderstandingNodeIDArr
		В общем - анализ состояния веток циклов осмысления для данной Цели, т.к. кнопки действий активируют смысл (контекст) ситуации SituationImage.
		В том числе, в контексте, если оператор нажал кнопку Учитель, это - стимул привлечения внимания для наблюдения за ним,
		за достигаемой им целью и как он ее достигает.
		*/
		// TODO

	} //if EvolushnStage > 4 { //5-я стадия развития. Творчество.

	/////////////////////////////////////////////////////////
	// СЛУЧАЙНЫЙ ВЫБОР
	needRunMenyalFunc := false // будет запуск инфо-функции,
	funcID := 0
	if EvolushnStage == 4 {
		if CurrentInformationEnvironment.veryActualSituation || CurrentInformationEnvironment.danger {
			c.log += "в infoFunc2() спросить, как правильно ответить, цикл осмысления прекращается<br>"
			funcID = 13 // спросить, как правильно ответить, цикл осмысления прекращается
		} else { // рискованные эксперименты
			// случайный выбор инфо-функции - поддержка продолжения цикла осмысления
			funcID = infoFindRundomMentalFunction()
		}
		needRunMenyalFunc = true
	}
	if EvolushnStage > 4 {
		//определить, нет ли доминанты с таким объектом, если нет опасности (определяется в checkDominantsJbject)
		if !checkDominantsJbject(extremImportanceObject) {
			if isIdleness() { // ЛЕНЬ
				funcID = 13 // спросить, как правильно ответить, цикл осмысления прекращается
			} else { // рискованные эксперименты
				// случайный выбор инфо-функции - поддержка продолжения цикла осмысления
				c.log += "в infoFunc2() случайный выбор инфо-функции<br>"
				funcID = infoFindRundomMentalFunction()
			}
			needRunMenyalFunc = true
		}
	}

	// завершение шага цикла с записью инфы goNext
	if needRunMenyalFunc {
		if runMentalFunctionID(c, funcID) { // запустить инфо-функцию
			return true
		}
	}

	return false
}

//////////////////////////////////////////////////////////

/*
	№3 найти подходящую infoFunc по опыту ментальных Правил

для данного detectedActiveLastProblemNodID

Тут нет конкретного образа действия, но есть ID инфо-функций, которые приводили к достижению Цели, но
С последней ф-цией может быть связан и MentalAutomatizm.motAutmtzmIDrulesm

Допускается не точное совпадение условий ментального правила.
*/
func infoFunc3(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	/*
		if c.lastFuncID == 3{// не вызывать, если только что было
			return
		}*/
	mentalInfoStruct.epizodFrameIndex = 0
	setCurIfoFuncID(c, 3)
	//	clinerMentalInfo()
	return infoFindRightMentalRules(c)

}
func infoFindRightMentalRules(c *cycleInfo) bool {
	mentalInfoStruct.epizodFrameIndex = 0
	/*
		funcID := findSuitableMentalFunc()
		if funcID!=nil {
		// выбираем первую функцию
			mentalInfoStruct.epizodFrameIndex = funcID[0]
			return true
		}*/
	fID := findInfoIdFromExperience(false) // false - можно не точное совпадение условий правила
	if fID > 0 {
		mentalInfoStruct.epizodFrameIndex = fID
		return true
	}
	return false
}

//////////////////////////////////////////////////////////

/*
	анализ инфо стркутуры mentalInfoStruct и др. информации  и выдача решения.

Нужна таблица, какие инфо-функции вызывать при данной ситуации.
Находит и запускает мот.автоматизм в плохой ситуации
*/
func infoFunc4(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 4 { // раз в прошлый раз уже вызывалась функция infoFunc4() то  не зацикливать
		return false
	}
	setCurIfoFuncID(c, 4)
	//	clinerMentalInfo()
	return analisAndSintez(c.ID, c)

}

/*
	анализ инфо стркутуры и др. информации и выдача решения

Нужна таблица, какие инфо-функции вызывать при данной ситуации??
*/
func analisAndSintez(fromNextID int, c *cycleInfo) bool {
	//sitID:=UnderstandingNodeFromID[detectedActiveLastUnderstandingNodID].SituationID
	unode, oku := ReadeUnderstandingNodeFromID(detectedActiveLastUnderstandingNodID)
	if !oku {
		return false
	}
	sitID := unode.SituationID

	//sit:=SituationImageFromIdArr[sitID].SituationType // текущая ситуация SituationImage
	node, ok := ReadeSituationImageFromIdArr(sitID)
	if !ok {
		return false
	}
	sit := node.SituationType // текущая ситуация SituationImage

	//	theme:=mentalInfoStruct.ThemeImageType // текущая тема Нет тем в Правилах, так что не приплетать

	var stimulsArr = []int{15, 23, 30, 32}

	// псих оценка ситуации: плохо, нужно решать проблему
	var dangAct = 0
	for _, v := range stimulsArr {
		if sit == v {
			dangAct = v
			break
		}
	}
	// псих оценка ситуации: хорошо, можно ничего не делать, даже если действия были опасными (может и не быть такого)
	if sit == 24 || sit == 27 || sit == 31 {
		dangAct = 0
	}
	if dangAct > 0 {
		// выбрать лучшее действие из обозримой эпиз.памяти с хорошим эффектом после стимулов с перечисленными проблемами
		actBestID := 0 // getBestAnswerFromEpisodeStimulsArr(dangAct)
		if actBestID > 0 {
			//			actBest:=ActionsImageArr[actBestID]
			actBest, ok := ReadeActionsImageArr(actBestID)
			if ok {
				/*получить пользу из эффекта лучшего моторного действия
				будет найден и запущен лучший мот.автоматизм aBest.ID
				mentalInfoStruct.motorAtmzmID=aBest.ID
				infoFunc17()// запустить автоматизм и завершить цикл осмысления
				*/
				getBenefitFromEpizosMemory(c, actBest)
				return true
			}
		}
	}

	return false
}

//////////////////////////////////////////////////////////

/*
	№5 Идет цикл мышления об экстремально важном объекте extremImportanceObject

Сновидение или мечтание или пассивное обдумывание или беседа (с кем-то или сам с собою)

// НАПРАВЛЕНИЕ МЫШЛЕНИЯ
Раскрутка сюжета по моторным Правилам со сменой сюжета в направлении наивысшей значимости.

	Теоретически здесь возможно произвольно выбирать глубину детализации объекта внимания ExtremImportanceObjectID
	"символьное" или "образное мышление". А пока что это делается автоматически.

	При выявлении ценного правила возникает озарение,
	прерывание эвристического цикла и начало нового с обновлением в инфо-картине члена ExtremImportanceObjectID.

	Найденное Правило дает новый extremImportanceObject, который там есть и следующий шаг цикла идет по нему.
	Каждый раз в CurrentInformationEnvironment.ExtremImportanceObjectID появляется текущий объект с его значимостью.

	Цель - эвристический синтез новых важных Правил. При этом возникает озарение в виде завершения такого цикла,
	новое правило становится объектом осмысления, extremImportanceObject обнуляется так, что ждет появления нового объекта.
*/
var shotrMemepMemArr []bestActionArr // память на время одного цикла мышления
func infoFunc5(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 5 { // не вызывать, если только что было
		return false
	}

	if c.isWaitingPeriod { // не совершать действий, а ждать ответа
		return false
	}

	if extremImportanceObject == nil { // нет экстремально важного объекта
		return false
	}

	if extremImportanceObjectOld == extremImportanceObject { // уже обрабатывался такой объект
		return false
	}
	extremImportanceObjectOld = extremImportanceObject

	setCurIfoFuncID(c, 5)
	c.log += "Запущена Инфо-функция № 5: Идет цикл мышления об экстремально важном объекте extremImportanceObject ID=" + strconv.Itoa(extremImportanceObject.objID) + "<br>"

	return infoMentalScaning(c)
}
func infoMentalScaning(c *cycleInfo) bool {
	if extremImportanceObject == nil {
		return false
	}

	shotrMemepMemArr = bestActionArrFromExtremImportanceObject()

	// В начале цикла все увиденные по аналогии с extremImportanceObject объекты Модели понимания типа extremImportance
	if shotrMemepMemArr == nil { // еще нет памяти о значащих объектах Модели понимания
		// первичная модель понимания, из топового объекта которой будут возникать другие модели
		// массив ID эпиз памяти для экстр-го объекта
		ActionsImage, effect := bestActionFromExtremImportanceObject()
		if ActionsImage == nil {
			// !!!! extremImportanceObject = nil // все на этом
			return false
		}
		bai := bestActionArr{ActionsImage, effect}
		shotrMemepMemArr = []bestActionArr{bai}
		// с каждым шагом цикла проходим по shotrMemepMemArr, начиная с самых топовых, назнаяая их как extremImportanceObject

		//Выбираем топовый по эффекту элемент shotrMemepMemArr
		eobj := getTopShotrMemepMemArr()
		if eobj != nil {
			extremImportanceObject = eobj
			return true
		}

	} else { // смотрим текущий элемент shotrMemepMemArr ставший extremImportanceObject
		/* Обработка информации об extremImportanceObject может идти по самым разным вариантам, их бесконечное множество.
		Это может стать поистине бесконечным процессом, если после нахождения экстремально важного кадра эпиз.памяти
		пройти по цепочке Правил, создавая по ним новые extremImportanceObject и помещая в стек shotrMemepMemArr[]
		Это бы походило на ветвление природного механизма сновидений у человека
		(но у других видов животных могут совсем другие наследственные алгоритмы).
		Далее реализован упрошенный вариант.
		*/

		// собрать образ действий типа *ActionsImage из элементов с позитивными эффектами цепочки
		sintezAction, ok := createSintezActiveActionsFromExtremImportanceObject()
		if !ok {
			return false
		}
		// добавить новую акцию в опыт
		actionsImageID, _ := CreateNewlastActionsImageID(0, 0, sintezAction.ActID, sintezAction.PhraseID, sintezAction.ToneID, sintezAction.MoodID, true)
		if actionsImageID > prefixActionIdValue { // AmtzmNextString.ID
			// это ID цепочки действия 	AmtzmNextString.ID
			mentalInfoStruct.AmtzmNextStringID = actionsImageID
			//выдать его на пульт
			infoFunc26(c) // запуск с периодом ожидания
			mentalInfoStruct.AmtzmNextStringID = 0
			return true
		} else { // одиночное действие
			if mentalInfoStruct.ActionsImageID > 0 { // создать, проверить и запустить сразу, не в цикле.
				//раз прошли, то просто вырезаем следующий топовый элемент shotrMemepMemArr для продолжения мышления о нем
				//eobj := cutLastExtremObj()
				eobj, _ := getObjectsImportanceValue(mentalInfoStruct.ActionsImageID, detectedActiveLastProblemNodID)
				if eobj != nil {
					extremImportanceObject = eobj
				}
				if c.isMainCycle {
					runAutomatizmAfterCheck(c)
				} else {
					insight(c)
				}
				return true
			}
		}
	}

	//reloadConsciousness(c,mentalInfoStruct.fromNextID)
	return false
}

//////////////////////////////
/* объект extremImportance с максимальным эффектом из стека shotrMemepMemArr[]
Выбираются объекты,
*/
func getTopShotrMemepMemArr() *extremImportance {
	if shotrMemepMemArr == nil {
		return nil
	}

	max := 0
	var maxE *extremImportance
	// вытащить экстрем из массива моделей понимания shotrMemepMemArr
	for i := 0; i < len(shotrMemepMemArr); i++ {

		eobj := getExtremObjFromID(shotrMemepMemArr[i].action.ID)
		if eobj != nil {
			if lib.Abs(max) < lib.Abs(shotrMemepMemArr[i].effect) {
				max = shotrMemepMemArr[i].effect
				maxE = eobj
			}
		}
	}
	if maxE != nil {
		return maxE
	}

	return nil
}

/////////////////////////////////////////////
/*попробовать для случая негативного obj *extremImportance или emExtrem *EpisodeMemory найти прошлый опыт как вывернуться
Из нескольких акций собрать один образ действий типа ActionsImage - из составляющих элементов с позитивными эффектами цепочки Правил
*/
func tryCreteNewActiveActionsFromRulesStr() int {
	sintezAction, ok := createSintezActiveActionsFromExtremImportanceObject()
	if !ok {
		return 0
	}
	// добавить новую акцию в опыт
	aID, _ := CreateNewlastActionsImageID(0, 0, sintezAction.ActID, sintezAction.PhraseID, sintezAction.ToneID, sintezAction.MoodID, true)
	return aID
}

//////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
/* №6 нужно ПОДВЕРГНУТЬ СОМНЕНИЮ автоматизм mentalInfoStruct.motorAtmzmID,
если нет опасности (не нужно реагировать аффектно) и ситуация важна.
Если заблоикрован mentalInfoStruct.motorAtmzmID и найден новый, то он СРАЗУ ЗАПУСКАТЕСЯ.
Если mentalInfoStruct.motorAtmzmID нормальный, то он так же сразу запускается тут.

Найденная по Значимстям и Моделям опасность при запуске будет НАМЕРТВО БЛОКИРОВАТЬ АВТОМАТИЗМ
поэтому в isNextWellEffectFromActonRules() нужно давать возможность преодолевать эту блокирующую значимость
если это приводит в Правилах к победе.

infoFunc6() дает результат сразу, не требуя следующего шага осмысления.
*/
func infoFunc6(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 6 { // не вызывать, если только что было
		return false
	}
	setCurIfoFuncID(c, 6)
	mentalInfoStruct.notOldAutomatizm = false // true - запрет запуска штатного автоматизма mentalInfoStruct.motorAtmzmID
	if (EvolushnStage == 4 || CurrentInformationEnvironment.veryActualSituation) &&
		!CurrentInformationEnvironment.danger {
		if mentalInfoStruct.motorAtmzmID > 0 {
			return infoCreateAndRunNewActionMentAtmzmFromAction(c)
		}
	}
	return false
}
func infoCreateAndRunNewActionMentAtmzmFromAction(c *cycleInfo) bool {

	if mentalInfoStruct.motorAtmzmID == 0 { // ID действия мот.автоматизма, рвущегося на выполнение
		return false
	}
	motorAtmzmID := mentalInfoStruct.motorAtmzmID

	//atmzm,ok := AutomatizmFromId[motorAtmzmID]
	atmzm, ok := ReadeAutomatizmFromId(motorAtmzmID)
	if !ok {
		return false
	}
	actImgID := atmzm.ActionsImageID
	// посмотреть, грозит ли опасностной значимостью запускаемый автоматизм
	res1 := isDangerousImportansAutomatism(actImgID)
	// посмотреть, грозит ли опасностными последствиями в Модели понимания запускаемый автоматизм
	res2 := isDangerousModelAutomatism(actImgID)
	// величины опасной значимости:
	if res1 > 0 || res2 > 0 {
		mentalInfoStruct.motorAtmzmBlockedID = mentalInfoStruct.motorAtmzmID
		/* 	нужно давать возможность произвольно преодолевать эту блокирующую значимость
		если это привоит в Правилах к победе. Сделать ПОИСК В ПРАВИЛАХ !!!!
		*/
		harm := res1
		if res2 > res1 {
			harm = res2
		}
		/* есть ли положительный эффект у Правила, превышающий величину вреда (harm), следующего за действием автоматизма,
		чтобы если он хороший, посчитать такое действие приемлемым и запустить автоматизм?
		*/
		if !isNextWellEffectFromActonRules(3, harm, actImgID) { // нет последующего позитивного эффекта
			mentalInfoStruct.notOldAutomatizm = true // не запускать такой автоматизм
			c.log += "Не запускать старый штатный автоматизм ID=" + strconv.Itoa(actImgID) + "<br>"
			// попробовать найти альтернативу

		}

		/* записать неудачный автоматизм mentalInfoStruct.motorAtmzmBlockedID в массив попыток действия
		и убрать из штатных.
		*/
		if mentalInfoStruct.motorAtmzmBlockedID > 0 {

			//autmzm := AutomatizmFromId[mentalInfoStruct.motorAtmzmBlockedID]
			autmzm, ok := ReadeAutomatizmFromId(mentalInfoStruct.motorAtmzmBlockedID)
			if !ok {
				return false
			}
			SetAutomatizmBelief(autmzm, 0) // убрать из штатных
			addNewTryAction(0, -detectedActiveLastProblemNodID, autmzm.ActionsImageID, -harm, true)
			mentalInfoStruct.motorAtmzmID = 0 // забыть про него
			return false
		}
		mentalInfoStruct.motorAtmzmBlockedID = 0
	} else { // mentalInfoStruct.motorAtmzmID не опасен и сразу тут запускается:
		c.log += "Запустить автоматизм ID=<b><span style='cursor:pointer;color:blue' onClick='show_automatizms(" + strconv.Itoa(mentalInfoStruct.motorAtmzmID) + ")'>" + strconv.Itoa(mentalInfoStruct.motorAtmzmID) + "</span></b><br>"
		// вытащить образ действий успешного автоматизма и попробовать решить подходящую по аналогии Домимнату
		checkRelevantAction(curActions.ID, atmzm.ActionsImageID, atmzm.Usefulness)

		infoFunc17(c) // запустить автоматизм, сделав его штатным и завершить цикл осмысления
		mentalInfoStruct.motorAtmzmID = 0
		return true
	}
	return false
}

/*
	сопоставление, грозит ли опасностной значимостью запускаемый автоматизм

Для данных условий
возвращает экстремальную негативную значимость объекта, или 0 если все в порядке
*/
func isDangerousImportansAutomatism(actImgID int) int {

	actImg, ok := ReadeActionsImageArr(actImgID)
	if !ok {
		return -10 // это странная ситуация, что нет данных про actImgID
	}
	objImportance := getExtremObjFromID(actImg.ID)
	if objImportance != nil && objImportance.extremVal < -3 {
		// нет ли потом в правилах этого фрагмента превышающий позитив - нет смысла при большом негативе
		return objImportance.extremVal // вернуть негатив
	}
	return 0
}

//////////////////////////////////////////////////////////
/*  ПОДВЕРГНУТЬ СОМНЕНИЮ автоматизм
используется при быстрой, БЕЗДУМНОЙ проверки штатного автоматизма на первом уровне осмысления.
В основном повторяет проверку в func infoFunc6
Если нет атаса выбирать среди них наиболее подходящий, даже если он не штатный и запускать его (не меняя штатность).
*/
func checkAutomatizm(atmzm *Automatizm) *Automatizm {

	if !CurrentInformationEnvironment.veryActualSituation && !CurrentInformationEnvironment.danger {
		//нет атаса, можно выбрать не штатный автоматизм из привязанных к ветке
		// список всех автоматизмов для ID узла Дерева

		aArr := GetMotorsAutomatizmListFromTreeId(detectedActiveLastNodID)

		harmMax := 1000
		var maxA *Automatizm
		for i := 0; i < len(aArr); i++ {
			actImgID := aArr[i].ActionsImageID
			// посмотреть, грозит ли опасностной значимостью запускаемый автоматизм
			res1 := isDangerousImportansAutomatism(actImgID)
			// посмотреть, грозит ли опасностными последствиями в Модели понимания запускаемый автоматизм
			res2 := isDangerousModelAutomatism(actImgID)
			if res1 > 0 || res2 > 0 {
				harm := res1
				if res2 > res1 {
					harm = res2
				}
				if harmMax > harm {
					harmMax = harm
					maxA = aArr[i]
				}
			}
		}
		if harmMax < 1000 { // выбран самый безопасный
			atmzm = maxA
		}
	}

	if atmzm == nil {
		return nil
	}

	actImgID := atmzm.ActionsImageID
	// посмотреть, грозит ли опасностной значимостью запускаемый автоматизм
	res1 := isDangerousImportansAutomatism(actImgID)
	// посмотреть, грозит ли опасностными последствиями в Модели понимания запускаемый автоматизм
	res2 := isDangerousModelAutomatism(actImgID)
	// величины опасной значимости:
	if res1 > 0 || res2 > 0 {
		harm := res1
		if res2 > res1 {
			harm = res2
		}
		// есть ли положительный эффект у Правила, превышающий величину вреда (harm), следующего за действием автоматизма,
		//чтобы если он хороший, посчитать такое действие приемлемым и запустить автоматизм?

		if !isNextWellEffectFromActonRules(3, harm, actImgID) { // нет последующего позитивного эффекта
			// не запускать такой автоматизм
			// записать неудачный автоматизм mentalInfoStruct.motorAtmzmBlockedID в массив попыток действия
			//и убрать из штатных.
			if mentalInfoStruct.motorAtmzmBlockedID > 0 {

				//	autmzm := AutomatizmFromId[mentalInfoStruct.motorAtmzmBlockedID]
				autmzm, ok := ReadeAutomatizmFromId(mentalInfoStruct.motorAtmzmBlockedID)
				if !ok {
					return nil
				}
				SetAutomatizmBelief(autmzm, 0) // убрать из штатных
				addNewTryAction(0, -detectedActiveLastProblemNodID, autmzm.ActionsImageID, -harm, true)
				// забыть про него
			}
			return nil
		}
		return atmzm
	} else {
		// mentalInfoStruct.motorAtmzmID не опасен и сразу тут запускается:
		// вытащить образ действий успешного автоматизма и попробовать решить подходящую по аналогии Домимнату
		checkRelevantAction(curActions.ID, atmzm.ActionsImageID, atmzm.Usefulness)
		return atmzm
	}
	return nil
}

///////////////////////////////////////////////////////////////////////

/*
	№7 создать новый моторный автоматизм по действию ActionsImageID -

ВСЕГДА ПОСЛЕ ПОЛУЧЕНИЯ ОБРАЗА ДЕЙСТВИЯ mentalInfoStruct.ActionsImageID
Создается моторный автоматизм (если такого еще нет), штатно привязанный к ветке текущей активности дерева автоматизмов.
И по mentalInfoStruct.runMotorAtmzmID=motorID запускается в infoFunc2() -> infoFunc17()
*/
var prevMotorAtmzmID = 0

func infoFunc7(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 7 { // не вызывать, если только что было
		return false
	}
	if mentalInfoStruct.ActionsImageID == 0 { // infoFunc7 не может работать без mentalInfoStruct.ActionsImageID >0
		return false
	}
	setCurIfoFuncID(c, 7)
	if mentalInfoStruct.ActionsImageID > 0 {
		infoCreateAndRunMentMotorAtmzmFromAction(mentalInfoStruct.ActionsImageID, c)
		return true
	}
	return false
}
func infoCreateAndRunMentMotorAtmzmFromAction(ActionsImageID int, c *cycleInfo) bool {
	if ActionsImageID == 0 {
		return false
	}
	motorID, motorAtmzm := createNewAutomatizmID(0, detectedActiveLastNodID, ActionsImageID, true)
	if motorID == 0 {
		mentalInfoStruct.ActionsImageID = 0
		mentalInfoStruct.runMotorAtmzmID = 0
		return false
	}
	if prevMotorAtmzmID == motorID { // раньше был запущен ментально такой мот.автоматизм
		// применить мозжечковый рефлекс
		cerebellumCoordination(motorAtmzm, 1) // 1 - усилить действие
	}
	prevMotorAtmzmID = motorID
	//	clinerMentalInfo()
	//mentalInfoStruct.motorAtmzmID=motorID
	/*if motorID==0{
		mentalInfoStruct.ActionsImageID=0
		mentalInfoStruct.runMotorAtmzmID=0
		return false
	}*/
	// ID моторного автоматизма который нужно запустить в infoFunc2() -> infoFunc17()
	mentalInfoStruct.runMotorAtmzmID = motorID
	return true
}

//////////////////////////////////////////////////////////

/*
	попытаться улучшить значимость текущего объекта extremImportanceObject по моделям понимания (understanding_model.go)

Задача - найти действие в данных условиях (bestActionFromExtremImportanceObject()).
которое лучше, чем текущий extremImportanceObject
и запустить по нему автоматизм, чтобы выработать правило с таким эффектом,
дающее при активации данной ветки автоматизмов более позитивный extremImportanceObject
*/
func infoFunc9(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 9 { // не вызывать, если только что было
		return false
	}
	setCurIfoFuncID(c, 9)
	//	clinerMentalInfo()
	return infoFindAttentionObjImprovement(c)

}

// улучшение значимости объекта внимания
func infoFindAttentionObjImprovement(c *cycleInfo) bool {
	if extremImportanceObject == nil {
		return false
	}

	act, _ := bestActionFromExtremImportanceObject()
	if act != nil {
		//infoFunc2() -> infoFunc7() -> infoFunc17()
		if act.ID > prefixActionIdValue { // последовательность образов действий AmtzmNextString.ID
			mentalInfoStruct.AmtzmNextStringID = act.ID - prefixActionIdValue
			return true
		}
		mentalInfoStruct.ActionsImageID = act.ID //в infoFunc2() создать автоматизм infoFunc7() по ID ActionsImage моторного действия
		return true
	}
	return false
}

////////////////////////////////////////////////////////////////////////

/*
	вспомнить самое значимое, что было - для начала размышения

(кроме возврата на прерванное) по InformationEnvironmentObjects []*InformationEnvironment

Это задает тему для нового цикла осмысления.

Если не найдено подходящей темы, то переактивации не будет.
Если тема найдена - после переактивации м.б. начнется размышление об этом.

TODO Функция может быть значительно усложнена, пока что сделано только самое поверхностное.
*/
func infoFunc10(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 10 { // не вызывать, если только что было
		return false
	}
	setCurIfoFuncID(c, 10)
	return infoFindAttentionObjMentalImprovement(c)
}

// улучшение объекта внимания с произвольной переактивацией дерева
func infoFindAttentionObjMentalImprovement(c *cycleInfo) bool {
	if InformationEnvironmentObjects == nil {
		mentalInfoStruct.volutionReload = 0
		return false
	}
	var needRecalingConsciousness = false
	// посмотреть инфо-окружения назад до первого очень важного объекта
	max := len(InformationEnvironmentObjects)
	for i := max - 1; i >= 0; i-- {
		eo := InformationEnvironmentObjects[i]
		eObj := eo.ExtremImportanceObjectID
		//		obj:=importanceFromID[eObj]
		obj, _ := getObjectsImportanceValue(eObj, detectedActiveLastUnderstandingNodID)
		if obj == nil {
			mentalInfoStruct.volutionReload = 0
			return false
		}
		if obj.extremVal > 3 { // вспомнить хорошее
			if eo.PsyMood > 0 { // хорошее настроение
				createThemeImageID(0, 2, 3, LifeTime, true) //Поисковый интерес
				needRecalingConsciousness = true
			}
			if eo.PsyMood < 0 { // плохое настроение
				createThemeImageID(0, 2, 5, LifeTime, true)
				needRecalingConsciousness = true
			}
		}

		if obj.extremVal < -3 { // вспомнить плохое
			if eo.PsyMood > 0 { // хорошее настроение
				createThemeImageID(0, 3, 17, LifeTime, true) //Состояние Плохо
				needRecalingConsciousness = true
			}
			if eo.PsyMood < 0 { // плохое настроение
				createThemeImageID(0, 3, 3, LifeTime, true) //Улучшение настроения
				needRecalingConsciousness = true
			}
		}

		if needRecalingConsciousness {
			clinerMentalInfo()
			mentalInfoStruct.volutionReload = 1
			// переактивировать с eo.PsyMood и eo.PsyEmotionId
			mentalMoodVolitionID = eo.PsyMood
			mentalEmotionVolitionID = eo.PsyEmotionId
			mentalMoodVolitionPulsCount = PulsCount
			mentalEmotionVolitionPulsCount = PulsCount
			// произвольлная переактивация
			understandingSituation(2)
			return true // не смотреть дальше
		}
	}
	return false
}

//////////////////////////////////////////

/*
	Ментальное отзеркаливание действия оператора

На 4-й стадии нужно отзеркаливать только то, что в действиях оператора достигает поставленной цели.
Щенок наблюдает как спускается по ступенькам и видит, что цель достигнута, начинает сам пробовать так же действовать.

if EvolushnStage == 4{// повышенная вероятность выбора в func infoFindRundomMentalFunction() для 11 и 13 - отзеркаливание
*/
func infoFunc11(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 11 { // не вызывать, если только что было
		return false
	}
	setCurIfoFuncID(c, 11)
	//	clinerMentalInfo()
	return infoMentalMirriring(c)
}
func infoMentalMirriring(c *cycleInfo) bool {
	//есть ли фраза в действиях оператора
	if curActiveActions == nil || curActiveActions.PhraseID == nil {
		return false
	}
	/* алгоритм:
	1. Найти методом GPT такую фразу в Ответах Beast, в Правилах: rulesID
	2. Посмотреть какое последовало действие оператора на это - в эпиз.памяти после rulesID: answer
	3. Создать автоматизм на такое действие и выдать на Пульт.
	*/
	getTargetEpisodicStrIdArr(curActiveActionsID, getLimitCountEM())
	//Результат получаем в глобальном targetEpisodicStrIdArr []Rule
	//- окончательно выбранная целевая цепочка Правил откуда берем следующее Действие.
	if targetEpisodicStrIdArr != nil {
		// есть цепочка с конечным плюсовым эффектом, значит можно так действвоать
		//rules := targetEpisodicStrIdArr[len(targetEpisodicStrIdArr)-1]
		acting, ok := ReadeActionsImageArr(targetEpisodicStrIdArr[0].Action)
		if ok {
			mentalInfoStruct.ActionsImageID = acting.ID //в infoFunc2() создать автоматизм infoFunc7() по ID ActionsImage моторного действия
			return true
		}
	}

	mentalInfoStruct.motorAtmzmID = 0
	return false
}

////////////////////////////////////////////////////////////////////////

/*
Cинтез новой фразы из компонентов, имеющих плюсы в Правилах
Возможен при хорошем словарном запасе и хорошем опыте значимости (importance) фраз.

При генерации нового словосочетания в infoFunc12 должна быть запись в дерево фраз?
НЕТ т.к. появляется автоматизм с цепочкой действий-фраз. Не нужен сенсорный детектор такой фразы.
*/
func infoFunc12(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 12 { // не вызывать, если только что было
		return false
	}
	setCurIfoFuncID(c, 12)
	//	clinerMentalInfo()
	return infoSynthesisOwnPrase(c)
}
func infoSynthesisOwnPrase(c *cycleInfo) bool {
	/*сделать фразу, состояющую из 2-3-х известных фраз, найденных в Правилах при данных условиях и выдать ее на Пульт
	  Чтобы фраза не была бессмысленным попугайством, нужно проверять ее смысл по importanceFromID
	  	importance.Type=5//фраза
	    importance.ObjectID=praseID
	    importance.Value>0
	  	для текущих условий
	    importance.NodeAID
	    importance.NodeSID
	*/
	getTargetEpisodicStrIdArr(curActiveActionsID, getLimitCountEM())
	//Результат получаем в глобальном targetEpisodicStrIdArr []Rule
	//- окончательно выбранная целевая цепочка Правил откуда берем следующее Действие.
	if targetEpisodicStrIdArr != nil {
		// есть цепочка с конечным плюсовым эффектом, значит можно так действвоать
		//rules := targetEpisodicStrIdArr[len(targetEpisodicStrIdArr)-1]
		var sumPrase []int
		for i := 0; i < len(targetEpisodicStrIdArr); i++ {
			// собираем фразу
			acting, ok := ReadeActionsImageArr(targetEpisodicStrIdArr[i].Action)
			if ok {
				if acting.PhraseID != nil {
					for j := 0; j < len(acting.PhraseID); j++ {
						sumPrase = append(sumPrase, acting.PhraseID[j])
					}
				}
			}
		} // for
		if len(sumPrase) > 0 { // выдать нормальным тоном с настроением wordSensor.CurPultMood
			actID, _ := CreateNewlastActionsImageID(0, 0, nil, sumPrase, 0, wordSensor.CurPultMood, true)
			if actID > 0 {
				//infoFunc2() -> infoFunc7() -> infoFunc17()
				mentalInfoStruct.ActionsImageID = actID //в infoFunc2() создать автоматизм infoFunc7() по ID ActionsImage моторного действия
				return true
			}
		}
	}

	return false
}

/////////////////////////////////////////////////////////////////////////

//TODO: надо переделать логику работы таких функций, работающих в виде цепочек диалога ожидающих ответа на каждом шаге. Это ИНСТИНКТЫ.
// Сделать в виде отдельного потока действий с ожиданием, перехватывающего на себя все внимание. Пока выполняется такой поток-цепочка, никакие функции и циклы не должны встревать, все ждут окончания цепочки либо превышения времения ожидания ответа при шаге цепочки
// в этом случае она прекращается и уступает дорогу другим. Иначе замахаемся отлавливать ситуации, когда работает цикл с func13, и параддельно лезут другие такие же циклы func13
// и разрывают цепочку действий.

/*
	Тупое повторение Стимула оператора. Попугайство.

Показать непонимания, растерянность с вопросом о том, как нужно реагировать.
В отличии от молчаливогоо отзеркаливания в infoFunc11()
создает и использует автоматизм с фразой:
"Ответь сам на "+спопугайничать curActiveActionsID оператора+" чтобы показать, как лучше ответить."
В последствии такой автоматизм может входить в Правила.

Отзеркалить последний Стимул от оператора и совершить такое же действие.
Нужно учесть, что отзеркаливание происходит и в orientation_reflexes.go func orientation_1()!!! м.б. стоит оттуда это убрать?
Отзеркаливает авторитерный ответ Оператора на совершенное действие с помощью func fixNewTeachRules()
- запись авторитарного Правила.

if EvolushnStage == 4{// повышенная вероятность выбора в func infoFindRundomMentalFunction() для 11 и 13 - отзеркаливание
*/
func infoFunc13(c *cycleInfo) bool {
	//	return false // пока блокируем, чтобы не мешалась. Возможно вообще надо убрать
	if c == nil {
		return false
	}
	if isUnrecognizedPhraseFromAtmtzmTreeActivation { // не делать автоматизм из нераспознанной фразы
		return false
	}
	if c.lastFuncID == 13 { // не вызывать, если только что было
		return false
	}
	if curFunc13ID > 0 { // не вызывать, если не отработан второй шаг цикла вопрос-ответ
		return false
	}
	// вызывает лишнюю паузу, два раза бота спрашивать приходится - пока комментим
	//if WasOperatorActiveted==false {// не вызывать, если не было стимула от Оператора
	//	return false
	//}

	if motorActionEffect > 0 {
		return false
	}
	// для случая, когда нажали учительскую кнопку, давшую отрицательный эффект - не нужно запускать в этом случае func 13
	// иначе просто создастся автоматизм с действием кнопки и все
	if GetBelief2AutomatizmListFromTreeId(detectedActiveLastNodID) != nil && mentalInfoStruct.motorAtmzmID > 0 {
		return false
	}

	setCurIfoFuncID(c, 13)
	//	clinerMentalInfo()
	return infoMirroringStimul(c)
}

// маркер для отработки второго шага отзеркаливания после ответа оператора
// в calcAutomatizmResult(): EvolushnStage == 3 || curFunc13ID > 0
var curFunc13ID int

func infoMirroringStimul(c *cycleInfo) bool {
	// свежесть ответа оператора - не позже, чем limitOfActionsAfterStimul пульсов назад
	// число ожидания пульсов должно совпадать с аналогичным числом в RumAutomatizm(), иначе 2-шаговый алгоритм infoFunc13() будет срабатывать не стабильно
	// периодически прерываясь на первом шаге и закрепляя в виде шататного попугайский автотизм, а не авторитарный, который должен помочь создать Оператор
	if curActiveActionsID > 0 {
		if curActiveActionsPulsCount > (PulsCount - limitOfActionsAfterStimul) {

			//actionsImageID,_:=CreateNewlastActionsImageID(0,0,[]int{1,21},nil,0,0,true)
			/* не создавать автоматизм с просьбой показать что-то, а прямо спопугайничть: сделать то же самое, что был 	curActiveActionsID
			в расчете получить реакцию и оценить такой прием как полражание.
			*/
			isTeachQuestion = true //Показать непонимание, растерянность с предложением научить
			motorID, atmz := createNewAutomatizmID(0, detectedActiveLastNodID, curActiveActionsID, true)
			oldUseful := atmz.Usefulness
			wasRunPurposeActionFunc = false // иначе может не пропустить в RumAutomatizmID()
			if atmz.Usefulness < 0 {
				atmz.Usefulness = 0 //если попугайский был заблокирован и получен его код - открываем его, иначе не получится задать вопрос
				atmz.Count = 1
			}
			lib.NoReflexWithAutomatizm = true // не показывать акции рефлексов с автоматизмами в одной плашке действия Beast на пульте
			if RumAutomatizmID(motorID) {
				c.log += "Инфо-функция 13: запуск моторного автоматизма <b> <span style='cursor:pointer;color:blue' onClick='show_automatizms(" + strconv.Itoa(mentalInfoStruct.motorAtmzmID) + ")'>" + strconv.Itoa(mentalInfoStruct.motorAtmzmID) + "</span>" + "</b>.<br>"
				runningMotAutmtzmID = motorID
				// при запуске прекратить думать про extremImportanceObject
				//!!! extremImportanceObject=nil
				mentalInfoStruct.motorAtmzmID = 0
				mentalInfoStruct.noStaffAtmzmID = false
				//			wasRunPurposeActionFunc =true  еще не было запуска!
				motorActionEffect = 1

				//setAsMaimCycle(c.ID)

				c.isWaitingPeriod = true    // блокируем повторную активацию этого цикла
				curFunc13ID = c.ID          // создаем маркер активации цикла, закрываем его в calcAutomatizmResult()
				atmz.Usefulness = oldUseful // возвращаем успешность обратно
				return true
			}
			atmz.Usefulness = oldUseful // возвращаем успешность обратно
			isTeachQuestion = false     // подстрахуемся, надо скидывать при любом раскладе - прошел автоматизм или нет
		}
	}
	return false
}

/////////////////////////////////////////////////////////////////////

/*
	Ментально переактиваровать дерево понимания с заданными параметрами узлов,

# НЕ МЕНЯЯ БАЗОВЫЕ КОНТЕКСТЫ

func14 запускается в infoFindRundomMentalFunction() только до 4-й стадии
ИЛИ при if mentalInfoStruct.ThemeImageType == 5 { //Поисковый интерес
т.к. происходит случайный выбор параметров переактивации для набора ментального опыта.
Если не задано mentalInfoStruct.moodeID - то moodeID будет выбрано случайно
Если не задано mentalInfoStruct.emotonID - то выбирается случайно
Если не задано mentalInfoStruct.SituationID - то ситуация выбирается случайно

Функция может использоваться для определенной целевой переактивации TODO,
например, для выхода из плохого ментального состояния (в func infoFunc2).

с mentalInfoStruct.moodeID и mentalInfoStruct.emotonID и mentalInfoStruct.SituationID

При переактивации изменяется Тема и Цель мышления
в func understandingSituation будет перезапущен главный цикл чтобы снова определить тему и цель
*/
func infoFunc14(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 14 { // не вызывать, если только что было
		return false
	}
	setCurIfoFuncID(c, 14)
	// clinerMentalInfo() чтобы можно было задавать mentalInfoStruct.moodeID и mentalInfoStruct.emotonID
	return reactivateEmotionUnderstandingЕree(c)
}
func reactivateEmotionUnderstandingЕree(c *cycleInfo) bool {
	if mentalInfoStruct.moodeID == 0 { // выбрать случайно - для infoFindRundomMentalFunction()
		var infoArr = []int{-1, 0, 1}
		mentalMoodVolitionID = lib.RandChooseIntArr(infoArr)
	} else {
		mentalMoodVolitionID = mentalInfoStruct.moodeID
	}
	if mentalInfoStruct.emotonID == 0 { // выбрать случайно - для infoFindRundomMentalFunction()
		var infoArr []int
		for k, v := range EmotionFromIdArr {
			if v == nil {
				continue
			}
			infoArr = append(infoArr, k)
		}
		mentalEmotionVolitionID = lib.RandChooseIntArr(infoArr)
	} else {
		mentalEmotionVolitionID = mentalInfoStruct.emotonID
	}

	if mentalInfoStruct.SituationID == 0 { // выбрать случайно - для infoFindRundomMentalFunction()
		var infoArr = getSituationTypeArrID()
		mentalSituationVolitionID = lib.RandChooseIntArr(infoArr)
	} else {
		mentalSituationVolitionID = mentalInfoStruct.SituationID
	}
	// это - уже конец главного цикла т.к. ьудет перезапуск understandingSituation(2)
	//wasRunPurposeActionFunc = true   // будет влиять на 1 и 2 уровнях т.к. сбрасывается на 3-м уровне if cycle.count == 0 {
	infoFuncSequence = append(infoFuncSequence, 14) // т.к. более не будет наполняться infoFuncSequence

	mentalMoodVolitionPulsCount = PulsCount
	mentalEmotionVolitionPulsCount = PulsCount
	mentalSituationVolitionPulsCount = PulsCount
	understandingSituation(2)
	clinerMentalInfo()
	return true
}

// переактивировать через func infoFunc14 в лучшее рабочее настроение
func findReactivadeParams(c *cycleInfo) bool {
	mentalInfoStruct.moodeID = 1 // сделать небольшой позитивчик
	// создать подходящую эмоцию
	gPars := gomeostas.GomeostazParams
	var bsIDarr []int
	bsIDarr = append(bsIDarr, 2) // поиск
	if gPars[1] < 30 {
		bsIDarr = append(bsIDarr, 1) // недостаток энергии
	}
	if gPars[8] > 50 { //большие повреждения
		bsIDarr = append(bsIDarr, 5) // защита
	}
	if gPars[4] > 70 && len(bsIDarr) < 3 { // потребность в общении
		bsIDarr = append(bsIDarr, 3) //
	}
	mentalInfoStruct.emotonID, _ = createNewBaseStyle(0, bsIDarr, true)
	// выбрать ситуацию
	mentalInfoStruct.SituationID = 4 //все спокойно, можно экспериментивароть
	if infoFunc14(c) {
		return true // больше не искать
	}
	return false
}

/////////////////////////////////////////////////////////////////////

/*
	для условия дерева автоматизмов (NodeAID) в одиночных Правилах выбираем наилучшее

Eсли для данного сочетания Стимул-Ответ есть только один вид эффекта,

	то это - уже Информация, и чем больше опыт (количество обобщенных правил), тем такая информация полезнее.
*/
func infoFunc15(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 16 { // не вызывать, если только что было
		return false
	}
	setCurIfoFuncID(c, 15)
	//	clinerMentalInfo()
	return beastIDRulesFromCondA(c)
}
func beastIDRulesFromCondA(c *cycleInfo) bool {
	/* В getTargetEpisodicStrIdArr получаем цепочку (или одно правило), заканчивающееся позитивом.
	В цепочке только последний член кончается удачей, остальные отрицательные,
	но смысл цепочки в том, что они приводит к удаче и поеэтому в данных конкретных условиях
	следует выполнить первую акций из этой цепочки.
	*/
	getTargetEpisodicStrIdArr(curActiveActionsID, getLimitCountEM())
	//Результат получаем в глобальном targetEpisodicStrIdArr []Rule
	//- окончательно выбранная целевая цепочка Правил откуда берем следующее Действие.
	if targetEpisodicStrIdArr != nil {
		// есть цепочка с конечным плюсовым эффектом, значит можно так действвоать
		rules := targetEpisodicStrIdArr[len(targetEpisodicStrIdArr)-1]
		if rules.Action > 0 {
			//res := makeActionFromRooles(rules) //	!!!!!ПЕРЕДЕЛАТЬ для type Rule struct {
			if rules.Action > prefixActionIdValue { // последовательность образов действий AmtzmNextString.ID
				mentalInfoStruct.AmtzmNextStringID = rules.Action - prefixActionIdValue
				return true
			}
			mentalInfoStruct.ActionsImageID = rules.Action //в infoFunc2() создать автоматизм infoFunc7() по ID ActionsImage моторного действия
			return true
		}
	}

	return false
}

///////////////////////////////////////////////////

/*
	По лучшей ЗНАЧИМОСТИ в данных условиях выдать действие и затем infoFunc7()

и сразу запустить (Писать мент.Правило на случайным действиям бесполезно)
*/
func infoFunc16(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 16 { // не вызывать, если только что было
		return false
	}
	setCurIfoFuncID(c, 16)
	//	clinerMentalInfo()
	return randomAction(c)
}

var oldRandActionsIDarr []int // выполненные действияID, чтобы не повторяться. Сбрасывается при просыпании if IsFirstActivation{
var wasRandPhrase = false     // последний раз была не фраза
func randomAction(c *cycleInfo) bool {
	var rActID = 0 //значимый образ ActionsImage

	if curActiveActions != nil {

	} //if curActiveActions != nil && curActiveActions.PhraseID!= nil{

	// попробовать найти по значимостям
	getBestRuleFromImpotrents()
	if bestRule.Count > 0 { // найден
		if bestRule.Effect == 100 {
			rActID = bestRule.Action
		} else {
			rActID = bestRule.Trigger
		}
	}

	if rActID == 0 {
		// выдать случайное действие из уже известных ActionsImageArr
		var actualArr []int
		for k, v := range ActionsImageArr {
			if v == nil {
				continue
			}
			actualArr = append(actualArr, k)
		}
		rActID = lib.RandChooseIntArr(actualArr)
	}

	if rActID > 0 {
		// не повторяться
		if !lib.ExistsValInArr(oldRandActionsIDarr, rActID) {
			return false
		}

		oldRandActionsIDarr = append(oldRandActionsIDarr, rActID)
		//infoFunc2() -> infoFunc7() -> infoFunc17()
		mentalInfoStruct.ActionsImageID = rActID //в infoFunc2() создать автоматизм infoFunc7() по ID ActionsImage моторного действия
		return true
	}

	return false
}

// /////////////////////////////////////////////////
/* инфа, сопровозжающая ментальный запуск мот.автоматизма НЕ ИСПОЛЬЗУЕТСЯ
var wasMentalRunMotorAtmzmID = 0 // ID моторного автомтаизма, запущенного ментально в infoFunc17()
// при ментальном запуске автоматизма таково состояние extremImportance
var wasExtremImportanceObject *extremImportance
*/
/////////////////////////////////////////////////
/*
	запустить моторный автоматизм mentalInfoStruct.motorAtmzmID

Записать этот автоматизм как штатный для узла дерева автоматизмов.
зафиксировать текущий goNex
и завершить цикл осмысления
*/
func infoFunc17(c *cycleInfo) bool { //запустить моторный автоматизм mentalInfoStruct.motorAtmzmID
	if c == nil {
		return false
	}
	if c.lastFuncID == 17 { // не вызывать, если только что было
		return false
	}
	setCurIfoFuncID(c, 17)
	if mentalInfoStruct.motorAtmzmID == 0 {
		return false
	}

	//autmzm:=AutomatizmFromId[mentalInfoStruct.motorAtmzmID]
	autmzm, ok := ReadeAutomatizmFromId(mentalInfoStruct.motorAtmzmID)
	if !ok {
		return false
	}
	// нет смысла запускать заблокированный автоматизм, его остановят в RumAutomatizmID()
	// и тем более нет смысла делать его штатным
	if autmzm.Usefulness >= 0 {
		SetAutomatizmBelief(autmzm, 2) // сделать автоматизм штатным
	}
	// инфа, сопровозжающая ментальный запуск мот.автоматизма
	//wasMentalRunMotorAtmzmID = mentalInfoStruct.motorAtmzmID
	//wasExtremImportanceObject = extremImportanceObject

	if RumAutomatizmID(mentalInfoStruct.motorAtmzmID) {
		c.log += "Инфо-функция 17: запуск моторного автоматизма <b> <span style='cursor:pointer;color:blue' onClick='show_automatizms(" + strconv.Itoa(mentalInfoStruct.motorAtmzmID) + ")'>" + strconv.Itoa(mentalInfoStruct.motorAtmzmID) + "</span>" + "</b>.<br>"
		runningMotAutmtzmID = mentalInfoStruct.motorAtmzmID
		// при запуске прекратить думать про extremImportanceObject
		//!!! extremImportanceObject=nil
		mentalInfoStruct.motorAtmzmID = 0
		mentalInfoStruct.noStaffAtmzmID = false
		lib.NoReflexWithAutomatizm = true // не показывать акции рефлексов с автоматизмами в одной плашке действия Beast на пульте
		wasRunPurposeActionFunc = true    // уже был запущен автоматизм
		// итак успевает
		//infoFuncSequence = append(infoFuncSequence, 17)// т.к. после не будет наполняться infoFuncSequence
		motorActionEffect = autmzm.Usefulness
		levelOfRunAutomatizm = 3 // для передачи в пульт при ответе бота - на каком уровне осмысления был дан ответ
		return true
	}

	//по любому остановить цикл т.к. запущен или старый или новый автоматизм
	//endBaseIdCycle(c.ID)
	//	reloadConsciousness(c,0)
	return false
}

/////////////////////////////////////////////////////////////

/*
	Ощущение CurrentInformationEnvironment - по наиболее значимому параметру.

Чем именно грозит наиболее значимое из параметров CurrentInformationEnvironment
1. выбрать  наиболее значимое изменение в параметрах CurrentInformationEnvironment - Здесь есть конкурентность с имеющимся extremImportanceObject (если он есть)
2. найти для него extremImportanceObject
После infoFunc18() всегда перезапуск осмысления
*/
func infoFunc18(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 18 { // не вызывать, если только что было
		return false
	}
	setCurIfoFuncID(c, 18)
	return locForInformationEnvironment(c)
}
func locForInformationEnvironment(c *cycleInfo) bool {

	// конкурентная значимость уже имебщегося объекта внимания extremImportanceObject для сравнения со значимостями пааметров ИнфоОкружения

	var extremVal = 0 // значимость extremImportanceObject если он есть
	var eObj *extremImportance
	// доминанта имеет доминирующий приоритет
	CurrentInformationEnvironment.DominantaID = 0
	if EvolushnStage > 4 {
		if CurrentProblemDominanta != nil {
			dm := CurrentProblemDominanta
			CurrentInformationEnvironment.DominantaID = dm.ID
			extrem0 := dm.weight
			if lib.Abs(extremVal) < lib.Abs(extrem0) {
				extremVal = extrem0
				if dm.objectID > 0 {
					eObj0 := getExtremObjFromID(dm.objectID)
					if eObj0 != nil {
						eObj = eObj0
					}
				}
			}
		}
	}
	//////////////////////////////////////////////////////
	// другие вды значимого
	if CurrentInformationEnvironment.DominantaID == 0 {
		if CurrentInformationEnvironment.ExtremImportanceObjectID > 0 {
			extremObj := getExtremObjFromID(CurrentInformationEnvironment.ExtremImportanceObjectID)
			if extremObj != nil {
				extremVal = extremObj.extremVal
				eObj = extremObj
			}
		}
		// выявить, что изменилось после OldInformationEnvironment
		if CurrentInformationEnvironment.ActionsImageID > 0 && OldInformationEnvironment.ActionsImageID != CurrentInformationEnvironment.ActionsImageID {
			// текущий образ сочетания действий с Пульта
			extrem0 := getExtremEffectFromModel(CurrentInformationEnvironment.ActionsImageID)
			if lib.Abs(extremVal) < lib.Abs(extrem0) {
				extremVal = extrem0
				eObj = &extremImportance{CurrentInformationEnvironment.ActionsImageID, extrem0}
			}
		}
		if OldInformationEnvironment.AnswerImageID != CurrentInformationEnvironment.AnswerImageID {
			// текущий образ сочетания действий с Пульта
			extrem0 := getExtremEffectFromModel(CurrentInformationEnvironment.AnswerImageID)
			if lib.Abs(extremVal) < lib.Abs(extrem0) {
				extremVal = extrem0
				eObj = &extremImportance{CurrentInformationEnvironment.ActionsImageID, extrem0}
			}
		}
	}

	/* Тут могут быть проверки и всех других параметров CurrentInformationEnvironment
	   Их нужно придумать, но УЖЕ ПРИНЦИП осмысления по текущему ИнфоОкружени. ПОКАЗАН.
	   TODO - добавлять сканирования других OldInformationEnvironment.nnn != CurrentInformationEnvironment.nnn по мере необъодимости
	*/

	////////////////////////  наивысший по значимости объект внимания
	if eObj != nil {
		//новый extremImportance
		//addNewcurImportanceObjectArr(*eObj)
		extremImportanceObject = eObj
		c.log += "Выбрана Цель в виде актуального объекта мышления<br> "
	}
	//////////////////////////////////
	/*
		if extremImportanceObject != nil {
			GotoDreaming() // начать сновидение или мечтание(пассивное размышление, не в ответ на Стимул)
			return true
		}
	*/

	return false
}

/////////////////////////////////////////////////////////////////////

/*
	было эвристическое озарение и найдено:

mentalInfoStruct.DominantaID //ID доминанты для коотрой переданы науденные действия
mentalInfoStruct.DominantSuccessAImgID int // найдено успешное действие
mentalInfoStruct.DominantSuccessValue int // оцененная успешность: 1 - доминанта решена, 2 - не точно, 3 - еще менее точно
*/
func infoFunc19(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 19 { // не вызывать, если только что было
		return false
	}
	setCurIfoFuncID(c, 19)
	return thinkingAboutHeuristics(c)
}
func thinkingAboutHeuristics(c *cycleInfo) bool {
	if CurrentInformationEnvironment.DominantaID == mentalInfoStruct.DominantaID {
		// сразу запустить
		if mentalInfoStruct.DominantSuccessValue == 1 {
			mentalInfoStruct.ActionsImageID = mentalInfoStruct.DominantSuccessAImgID
			runAutomatizmAfterCheck(c)
			return true
		}
		if mentalInfoStruct.DominantSuccessValue > 1 {
			// запустить если спосокйно

			return true
		}
	} else {
		/* если нет моторного автоматизма в func consciousnessElementary
		запускать runDominantaAction() (сделано)
		*/
		// при озарении из неглавного цикла (подсознание)
		if mentalInfoStruct.ActionsImageID > 0 {
			runAutomatizmAfterCheck(c)
			return true
		}
	}
	return false
}

/////////////////////////////////////////////////////////////////////////////////////////////

// Воспоминание назад (начиная с dreamingEpisodeHistoryID и уходя в прошлое истории).
/* выбрать последний пустой (т.е. без Правил) кадр эпиз.памяти dreamingEpisodeHistoryID и запустить процесс осмысления infoFunc5()
Чтобы прервать процесс размышления о прошлом:	dreamingEpisodeHistoryID=0
*/
func infoFunc20(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 20 { // не вызывать, если только что было
		return false
	}
	setCurIfoFuncID(c, 20)
	return beginPrepareLastEmptyEpisode(c)
}
func beginPrepareLastEmptyEpisode(c *cycleInfo) bool {

	// последний индекс исторической памяти EpisodicHistoryArr[] но не далее 1000 эпизодов
	if c.dreamingEpisodeHistoryID == 0 {
		c.dreamingEpisodeHistoryID = len(EpisodicHistoryArr) - 1
	}
	lastI := c.dreamingEpisodeHistoryID
	for i := lastI; i >= 0 && i > (lastI-1000); i-- {

		node, ok := ReadeEpisodicTreeNodeFromID(EpisodicHistoryArr[i].ID)
		if !ok {
			continue
		}
		if node.PARAMS == nil { // на всякий случай, хотя такого не должно быть
			continue
		}
		if node.Action == 0 {
			// вытащить образ Стимула
			//_,ai:=getActiveActionsFromAutomatizmTreeNode(node.NodeAID)
			ai, ok := ReadeActionsImageArr(node.Action)
			// выделить наиболее значимое в восприятии в массив типа extremImportance
			if ok {
				eobj := getExtremObjFromID(ai.ID) // для текущего detectedActiveLastProblemNodID
				if eobj != nil {
					extremImportanceObject = eobj
					if extremImportanceObject.extremVal < -5 { // большой негатив
						//dreamAboutExtremImportanceObject(c) // размышление об extremImportanceObject
						if isDreamInterrupt {
							continue
						}
						if infoFunc5(c) { // при отсуствии extremImportanceObject не пойдет
							continue
						}
					}

					CurrentInformationEnvironment.ExtremImportanceObjectID = extremImportanceObject.objID
					// Можно по extremImportance найти все []*importance: getImportanceFromExtremImportance(Stimul int, kind int)
					// extremImportanceObject.extremVal - НАГАТИВНАЯ ЗНАЧИМОСТЬ объекта
					runNewTheme(16, 2) // при этом создается новый цикл мышления ???

					// i - последний отработанный индекс массива EpisodicHistoryArr []int
					c.dreamingEpisodeHistoryID = i // начать процесс размышления или сновидения
					return true
				}
			}
			break
		}
	}
	c.dreamingEpisodeHistoryID = 0 // прошли до самого начала истории эпиз.памяти
	return false
}

/////////////////////////////////////////////////////////////////////////////////////////////

// Воспоминание вперед (начиная с dreamingEpisodeHistoryID и уходя вперед по истории).
/* выбрать пустой (т.е. без Правил) кадр эпиз.памяти dreamingEpisodeHistoryID
Чтобы прервать процесс размышления:	dreamingEpisodeHistoryID=0
*/
func infoFunc202(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 202 { // не вызывать, если только что было
		return false
	}
	setCurIfoFuncID(c, 202)
	return beginPrepareNextEmptyEpisode(c)
}
func beginPrepareNextEmptyEpisode(c *cycleInfo) bool {
	//	if c.dreamingEpisodeHistoryID == 0 { НАЧИНАТЬ С НАЧАЛА ИСТОРИИ

	lastI := c.dreamingEpisodeHistoryID
	for i := lastI; i < len(EpisodicHistoryArr); i++ {

		node, ok := ReadeEpisodicTreeNodeFromID(EpisodicHistoryArr[i].ID)
		if !ok {
			continue
		}
		if node.Action > 0 { // пройти назад, чтобы начать рассматривать последовательность кадров с пустого
			i = returnEmptyHistoryID(i)
			n, ok := ReadeEpisodicTreeNodeFromID(EpisodicHistoryArr[i].ID)
			if !ok {
				return false // непонятно что случилось...
			}
			node = n
		}
		if node.PARAMS == nil { // на всякий случай, хотя такого не должно быть
			continue
		}
		if node.Action == 0 {
			// вытащить образ Стимула
			//_,ai:=getActiveActionsFromAutomatizmTreeNode(node.NodeAID)
			ai, ok := ReadeActionsImageArr(node.Action)
			// выделить наиболее значимое в восприятии в массив типа extremImportance
			if ok {
				eobj := getExtremObjFromID(ai.ID) //для текущего detectedActiveLastProblemNodID
				if eobj != nil {
					extremImportanceObject = eobj
					/* это - при новом проходе func GotoDreaming
					if extremImportanceObject.extremVal < -5 { // большой негатив
						dreamAboutExtremImportanceObject(c) // размышление об extremImportanceObject в infoFunc5()
					}
					*/
					CurrentInformationEnvironment.ExtremImportanceObjectID = extremImportanceObject.objID
					// Можно по extremImportance найти все []*importance: getImportanceFromExtremImportance(Stimul int, kind int)
					// extremImportanceObject.extremVal - НАГАТИВНАЯ ЗНАЧИМОСТЬ объекта
					runNewTheme(16, 2) // 16: "Есть объект высокой значимости" при этом создается новый цикл мышления

					// i - последний отработанный индекс массива EpisodicHistoryArr []int
					c.dreamingEpisodeHistoryID = i // чтобы начать следующий проход или сновидения
					return true
				}
			}
			break
		}
	}
	c.dreamingEpisodeHistoryID = 0 // прошли до самого конца истории эпиз.памяти
	return false
}

/////////////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////////////////////////
/* Срочно найти какое-то подходящее действие, получить mentalInfoStruct.runMotorAtmzmID
и тут же запустить его
*/
func infoFunc21(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if c.lastFuncID == 21 { // не вызывать, если только что было
		return false
	}
	if motorActionEffect > 0 && !c.dreaming { // не нужно искать решение, раз уже был нормальный ответ
		return false
	}
	setCurIfoFuncID(c, 21)
	return findRunMotorAtmzmID(c)
}
func findRunMotorAtmzmID(c *cycleInfo) bool {

	actionsImageID := findSuitableMotorAction(c)
	if actionsImageID > prefixActionIdValue { // AmtzmNextString.ID
		// это ID цепочки действия 	AmtzmNextString.ID
		mentalInfoStruct.AmtzmNextStringID = actionsImageID
	} else { // одиночное действие
		mentalInfoStruct.ActionsImageID = actionsImageID
	}

	//////////////////  что-нибудь попробовать против опасности
	if mentalInfoStruct.ActionsImageID == 0 {
		if CurrentInformationEnvironment.veryActualSituation || CurrentInformationEnvironment.danger {

			//  ПАНИКА
			if mentalInfoStruct.ActionsImageID == 0 {
				//создать автоматизм просьбы о помощи из действия ID: 1 (не понимает), 21 (паника)
				//		mentalInfoStruct.ActionsImageID,_=CreateNewlastActionsImageID(0,0,[]int{1,21},nil,0,0,true)
				// вызвать состояние ступора:этот ступорный цикл больше не перезапускать, для ожидания Стимула.
				if infoFunc31(c) {
					return true
				} // сначала попробовать найти решение через провокацию
				c.isStupor = true

				// как бы невербально выражаемая растерянность, не автоматизм, а типа без.рефлекса, но на уровне психики
				lib.SentConfusion("Не могу понять, что делать.<br>СТУПОР!<br>Просьба о помощи.") // непосредственно сообщить о панике
				c.log += "Ступорный цикл остановить. СТУПОР! Просьба о помощи.<br>"
				return true
			}
		}
	} else { // нет опасности
		// TODO что-нибудь попробовать в МЕНТ-Х ПРАВИЛАХ
		if infoFunc3(c) { //найти подходящую infoFunc по опыту ментальных Правил
			runMentalFunctionID(c, mentalInfoStruct.epizodFrameIndex)
			mentalInfoStruct.ActionsImageID = 0
			return true
		}
	}
	// выполнить, если есть что
	if mentalInfoStruct.ActionsImageID > 0 { // создать, проверить и запустить сразу, не в цикле.
		runAutomatizmAfterCheck(c)
		mentalInfoStruct.ActionsImageID = 0
		return true
	}
	if mentalInfoStruct.AmtzmNextStringID > 0 {
		//выдать его на пульт
		infoFunc26(c) // запуск с периодом ожидания
		mentalInfoStruct.AmtzmNextStringID = 0
		return false
	}
	return false
}

/////////////////////////////////
/* срочно найти подходящий образ действия mentalInfoStruct.ActionsImageID
по экстремальному объекту внимания
*/
func findSuitableMotorAction(c *cycleInfo) int {
	actBest, _ := bestActionFromExtremImportanceObject()
	if actBest != nil { // найдено лучшее
		CurrentInformationEnvironment.ActionsImageID = actBest.ID // в текущую инфо-картину
		return actBest.ID
	}

	return 0
}

/////////////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////
/* Добавление нового действия mentalInfoStruct.ActionsImageID в цепочку действий автоматизма Automatizm.NextID.

Можно в цикле мышления начать формировать произвльную цепочку действий в infoFunc22()
на основе имеющегося автоматизма,
добавляя новое звено, возможно запуская ее по мере добавления нового звена
чтобы оценивать достижение цели с этого звена и с этим звеном.
*/
func infoFunc22(c *cycleInfo) {
	if c == nil {
		return
	}
	// инициативная произвольность возможна только в главном цикле
	if !c.isMainCycle {
		return
	}
	if c.lastFuncID == 22 { // не вызывать, если только что было
		return
	}
	if mentalInfoStruct.ActionsImageID == 0 {
		return
	}
	setCurIfoFuncID(c, 22)
	addNextAtmzmID(c)
}
func addNextAtmzmID(c *cycleInfo) {
	// добавить в конец
	id, _ := addAmtzmNextString(mentalInfoStruct.AmtzmNextStringID, []int{mentalInfoStruct.ActionsImageID})
	if id > 0 {
		mentalInfoStruct.AmtzmNextStringID = id
	}
	//insertImageActionToAmtzmNextID(mentalInfoStruct.ActionsImageID,-1,mentalInfoStruct.motorAtmzmID)
	// запустить mentalInfoStruct.motorAtmzmID с довеском Next
	// НЕТ можно доавбять в цикле, не запуская infoFunc17(c)
}

///////////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////
/* Создание новой цепочки произвольных действий mentalInfoStruct.AmtzmNextStringID
с добавлением mentalInfoStruct.ActionsImageID
Просто начать новую произвольную цепочку: mentalInfoStruct.AmtzmNextStringID=0

Можно в цикле мышления начать формировать произвольную цепочку действий mentalInfoStruct.AmtzmNextStringID
не привязанную пока к автоматизму, но которую можно выполнить showNextAtmtzmAction(0,mentalInfoStruct.AmtzmNextStringID,5)
чтобы оценивать достижение цели с этого звена и с этим звеном (непонятно как),
добавляя infoFunc24() новое звено, возможно запуская ее по мере добавления нового звена.
*/
func infoFunc23(c *cycleInfo) {
	if c == nil {
		return
	}
	// инициативная произвольность возможна только в главном цикле
	if !c.isMainCycle {
		return
	}
	if c.lastFuncID == 23 { // не вызывать, если только что было
		return
	}
	if mentalInfoStruct.ActionsImageID == 0 {
		return
	}
	setCurIfoFuncID(c, 23)

	if mentalInfoStruct.ActionsImageID > 0 {
		// добавить в конец
		id, _ := addAmtzmNextString(mentalInfoStruct.AmtzmNextStringID, []int{mentalInfoStruct.ActionsImageID})
		if id > 0 {
			mentalInfoStruct.AmtzmNextStringID = id
		}

		//mentalInfoStruct.AmtzmNextStringID,_=createAmtzmNextStringID(0,[]int{mentalInfoStruct.ActionsImageID},true)
	}

}

///////////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////
/* добавление нового звена mentalInfoStruct.motorAtmzmID
в цепочку произвольных действий mentalInfoStruct.AmtzmNextStringID,
возможно запуская ее по мере добавления нового звена infoFunc26()
*/
func infoFunc24(c *cycleInfo) {
	if c == nil {
		return
	}
	// инициативная произвольность возможна только в главном цикле
	if !c.isMainCycle {
		return
	}
	if mentalInfoStruct.ActionsImageID == 0 {
		return
	}
	setCurIfoFuncID(c, 24)
	addVolutionString(c)
}
func addVolutionString(c *cycleInfo) {
	lib.MapFree(MapGwardAutomatizmNextStringFromID)
	nArr, ok := AutomatizmNextStringFromID[mentalInfoStruct.AmtzmNextStringID]
	if !ok {
		return
	}
	insertImageActionToAmtzmNext(nArr, -1, mentalInfoStruct.motorAtmzmID)
}

///////////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////
/* Посмотреть условия чтобы проявить инициативу.
СЛУЧАИ:

1. Не было действий Beast  более 100 пульсов, но есть необходимость попробовать что-то сделать (нет другой возможности повлиять на состояние).
Так что ЕСЛИ ФИГОВО И НЕТ АВТОМАТИМА ДЛЯ ДАННОГО СОСТОЯНИЯ (при просыпании проявить инициативу и что-то сделать-сказать).

2. При стимуле с Пульта нет автоматизма или пытается найти альтренативу, в сомнении.
НО Нет походящего Правила, Значимости и Модели чтобы выработать реакцию.
При ИССЛЕДОВАТЕЛЬСКОМ ИЛИ УЧЕНИЧЕСКОМ КОНТЕКСТЕ нужно экспериментировать (иначе - просто промолчать).

В этих случаях запускается функция для 1) найти новое infoFunc23(c)
и запустить как произвольность showNextAtmtzmAction, 2) дополнить сомнительный автоматизм infoFunc22(c),
в том числе если уже есть довесок Automatizm.NextID то добавить еще infoFunc24(c).
*/
func infoFunc25(c *cycleInfo) {
	if c == nil {
		return
	}
	if EvolushnStage < 4 {
		return
	}

	setCurIfoFuncID(c, 25)

	lookForInitiativeAction(c)

	//Нужно подумать о проблеме автоматизма или проявить инициативу, в общем, запустить func infoFunc25()
	CurrentInformationEnvironment.needThinkingAboutAutomatizm = false
}
func lookForInitiativeAction(c *cycleInfo) {

	// для отладки:
	//	CurrentInformationEnvironment.veryActualSituation = true
	//	CurrentInformationEnvironment.IsIdleness100pulse = true

	mentalInfoStruct.AmtzmNextStringID = 0 // если будут найдены действия, то будет создана цепочка с таким ID
	// попробовать найти уже имеющуюся подходящую цепочку для текущей проблемы detectedActiveLastProblemNodID
	//чтобы нарастить ее клон (сохранив оригинал)
	if extremImportanceObject != nil {
		act, _ := bestActionFromExtremImportanceObject()
		if act != nil {
			mentalInfoStruct.AmtzmNextStringID = createClonForActins([]int{act.ID})
			//return
		}
	}

	////////////////////////////////////////s

	//Не было действий Beast  более 100 пульсов, возможно есть условия проявить инициативу
	if CurrentInformationEnvironment.IsIdleness100pulse {
		CurrentInformationEnvironment.IsIdleness100pulse = false // чтобы больше не активировало

		// нужно же что-то сделать, раз нет другой возможности повлиять на состояние
		if CurrentInformationEnvironment.veryActualSituation || CurrentInformationEnvironment.danger {

			// найти подходящий образ действия по экстремальному объекту внимания
			actionsImageID := findSuitableMotorAction(c)
			if actionsImageID > 0 {
				if actionsImageID > prefixActionIdValue { // AmtzmNextString.ID
					// это ID цепочки действия 	AmtzmNextString.ID
					mentalInfoStruct.AmtzmNextStringID = actionsImageID - prefixActionIdValue
					//выдать его на пульт
					infoFunc26(c) // запуск с периодом ожидания
					mentalInfoStruct.AmtzmNextStringID = 0
					return
				} else { // одиночное действие
					if mentalInfoStruct.ActionsImageID > 0 {
						//Создание новой цепочки произвольных действий с добавлением mentalInfoStruct.motorAtmzmID
						//infoFunc23(c)
						// добавить найденное действие actID в конец существующей цепочки в mentalInfoStruct.AmtzmNextStringID
						addAmtzmNextString(mentalInfoStruct.ActionsImageID, []int{mentalInfoStruct.ActionsImageID})
						mentalInfoStruct.ActionsImageID = 0
						// в опасной ситуации достаточно одного действия, чтобы выдать его на пульт
						infoFunc26(c) // запуск с периодом ожидания
						mentalInfoStruct.AmtzmNextStringID = 0
						return
					}
				}
			}

			// нет действия по экстремальному объекту, ищем в правилах
			/*Найти методом GPT последнее известное Правило по цепочке последних limit кадров эпиз.памяти
			  Найти самое свежее Правило, имеющее схожесть с последними кадрами эпизодической памяти
			  для текущей ветки дерева проблем detectedActiveLastUnderstandingNodID, а значит, для текущей совокупности условий.
			  В эпиз.памяти записываются Правила с учетом последовательности образов действий типа AmtzmNextString.ID
			  				так что функция getRulesFromEpisodicsSlice должна работать корректо.

			  Последний Стимул при активации Дерева автоматизмов: curActiveActionsID, найти для него наиболее преспективную цепочку Правил
			  				и по нему - рекомендуемое действие.
			*/
			getTargetEpisodicStrIdArr(curActiveActionsID, getLimitCountEM())
			//Результат получаем в глобальном targetEpisodicStrIdArr []Rule
			//- окончательно выбранная целевая цепочка Правил откуда берем следующее Действие.
			if targetEpisodicStrIdArr != nil {
				// есть цепочка с конечным плюсовым эффектом, значит можно так действвоать
				rules := targetEpisodicStrIdArr[len(targetEpisodicStrIdArr)-1]
				if rules.Action > 0 {
					// найти цепочку AmtzmNextString с таким действием, если нет - создать, если есть - создать клон
					mentalInfoStruct.AmtzmNextStringID = createClonForActins([]int{rules.Action})
				}
			}

		} else { // Не было действий Beast  более 100 пульсов, но неопасная ситуация
			// контекст 2 "Поиск"
			if existsBaseContext(2) {
				var rules Rule // найти Правило
				/*Найти методом GPT последнее известное Правило по цепочке последних limit кадров эпиз.памяти
				  Найти самое свежее Правило, имеющее схожесть с последними кадрами эпизодической памяти
				  для текущей ветки дерева проблем detectedActiveLastUnderstandingNodID, а значит, для текущей совокупности условий.
				  В эпиз.памяти записываются Правила с учетом последовательности образов действий типа AmtzmNextString.ID
				  				так что функция getRulesFromEpisodicsSlice должна работать корректо.

				  Последний Стимул при активации Дерева автоматизмов: curActiveActionsID, найти для него наиболее преспективную цепочку Правил
				  				и по нему - рекомендуемое действие.
				*/
				getTargetEpisodicStrIdArr(curActiveActionsID, getLimitCountEM())
				//Результат получаем в глобальном targetEpisodicStrIdArr []Rule
				//- окончательно выбранная целевая цепочка Правил откуда берем следующее Действие.
				if targetEpisodicStrIdArr != nil {
					// есть цепочка с конечным плюсовым эффектом, значит можно так действвоать
					rules = targetEpisodicStrIdArr[len(targetEpisodicStrIdArr)-1]
				}
				if rules.Action > 0 {
					//По правилу найти или создать (в случае AmtzmNextString) автоматизм и запустить его. Если совершено действие - возвращает true
					res := makeActionFromRooles(rules)
					if res {
						return
					}
				}

				// найти подходящий образ действия  в цепочках в спокойной исследовательской ситуации
				findResearchMotorAction(c) // наращена цепочка mentalInfoStruct.AmtzmNextStringID
				if mentalInfoStruct.AmtzmNextStringID > 0 {
					infoFunc26(c) // запуск с периодом ожидания
					mentalInfoStruct.AmtzmNextStringID = 0
					return
				}
			}
		}
	}
	//if CurrentInformationEnvironment.IsIdleness100pulse{

	// набор цепочики mentalInfoStruct.AmtzmNextStringID разными способами
	if mentalInfoStruct.AmtzmNextStringID == 0 {

		// стоит насущная проблема (вызов из func dangerActualProcess)
		lib.MapCheck(MapGwardProblemTryCount)
		if problemTryCount[detectedActiveLastProblemNodID] > 1 { // счетчик нерешенной проблемы для текущего detectedActiveLastProblemNodID
			// найти подходящий образ действия в цепочках в спокойной исследовательской ситуации
			findResearchMotorAction(c) // наращена цепочка mentalInfoStruct.AmtzmNextStringID
		}
		if mentalInfoStruct.AmtzmNextStringID == 0 {
			act := tryCreteNewActiveActionsFromRulesStr()
			// найти цепочку AmtzmNextString с таким действием, если нет - создать, если есть - создать клон
			mentalInfoStruct.AmtzmNextStringID = createClonForActins([]int{act})
		}
	}
	///////////////////////////////////////////

	///////////////////////////////////////////
	if mentalInfoStruct.AmtzmNextStringID > 0 {
		infoFunc26(c) // запуск с периодом ожидания
		mentalInfoStruct.AmtzmNextStringID = 0
		return
	}
	mentalInfoStruct.AmtzmNextStringID = 0

}

//////////////////////////////////////////////////////////////////////

////////////////////////
/* спокойно найти подходящий образ действия mentalInfoStruct.motorAtmzmID в исследовательской ситуации
Выбрать все, что подходит под ситуацию, набирая цепочку
из Правил и из карты AutomatizmNextStringFromID[] (прогнозы по Правилам).

Алгоритм: начиная с текущего CurrentInformationEnvironment.ExtremImportanceObjectID
найти действие, добавить в цепочку, при этом выявляя следующий экстремальный объект
и начинать поиск от него, пока не наберется достаточная (по прогнозам?) цепочка.
*/
func findResearchMotorAction(c *cycleInfo) {

	// по экстремальному образу - м.б. совокупным (сделанным из нескольких позитивных)
	// чтобы найти самый лучший одинчный образ: actBest:=bestActionFromExtremImportanceObject()
	actID := tryCreteNewActiveActionsFromRulesStr()
	if actID != 0 {
		// добавить найденные действия в mentalInfoStruct.AmtzmNextStringID
		addAmtzmNextString(mentalInfoStruct.ActionsImageID, []int{actID})
	}

	// по правилам для текущей ситуации
	getTargetEpisodicStrIdArr(curActiveActionsID, getLimitCountEM())
	//Результат получаем в глобальном targetEpisodicStrIdArr []Rule
	//- окончательно выбранная целевая цепочка Правил откуда берем следующее Действие.
	if targetEpisodicStrIdArr != nil {
		// есть цепочка с конечным плюсовым эффектом, значит можно так действвоать
		rules := targetEpisodicStrIdArr[len(targetEpisodicStrIdArr)-1]
		if rules.Action > 0 {
			// добавить найденные действия в mentalInfoStruct.AmtzmNextStringID
			addAmtzmNextString(mentalInfoStruct.ActionsImageID, []int{rules.Action})
		}
	}
	return
}

///////////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////
/* Запуск на выполнение на пульте имеющейся цепочки произвольных действий mentalInfoStruct.AmtzmNextStringID
без создания автоматизма
с периодом ожидания после которого автоматизм создается по результату гомео-эффекта.
*/
func infoFunc26(c *cycleInfo) {
	if c == nil {
		return
	}
	// инициативная произвольность возможна только в главном цикле
	if !c.isMainCycle {
		return
	}
	if mentalInfoStruct.AmtzmNextStringID == 0 {
		return
	}
	setCurIfoFuncID(c, 26)
	if LastRunAutomatizmPulsCount > 0 { // не запускать в период ожидания
		return
	}
	runVolutionString(c)
}
func runVolutionString(c *cycleInfo) {
	levelOfRunAutomatizm = 3 // для передачи в пульт при ответе бота - на каком уровне осмысления был дан ответ

	infoFuncSequence = append(infoFuncSequence, 26) // т.к. после не будет наполняться infoFuncSequence

	// запуск с энергией == 5
	showNextAtmtzmAction(0, mentalInfoStruct.AmtzmNextStringID, 5)

	lib.NoReflexWithAutomatizm = true // не показывать акции рефлексов с автоматизмами в одной плашке действия Beast на пульте
	wasRunPurposeActionFunc = true    // уже была запущена цепочка действий
}

///////////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////
/* Создание автоматизма на основе имеющейся цепочки произвольных действий mentalInfoStruct.AmtzmNextStringID
Получение значения mentalInfoStruct.motorAtmzmID для запуска.

Эта функция - на всякий случай, т.к.
АВТОМАТИЗМ ВСЕГДА СОЗДАЕТСЯ ПРИ ЗАПУСКЕ mentalInfoStruct.AmtzmNextStringID на выполнение
	по результату периода ожидания - в func calcAutomatizmResult()
*/
func infoFunc27(c *cycleInfo) {
	if c == nil {
		return
	}
	// инициативная произвольность возможна только в главном цикле
	if !c.isMainCycle {
		return
	}
	if mentalInfoStruct.AmtzmNextStringID == 0 {
		return
	}
	setCurIfoFuncID(c, 27)
	createAnvtzmFromVolutionString(c)
}
func createAnvtzmFromVolutionString(c *cycleInfo) {
	automatizm := createAndRunAutomatizmFromAmtzmNextString(mentalInfoStruct.AmtzmNextStringID)
	mentalInfoStruct.motorAtmzmID = automatizm.ID
}

///////////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////
/* фантазия - как функция интенсивного подбора ассоциаций для решения доминанты

Это - гипотетичкая реализация, - намеки на будущую, задел.

В результате - // НЕТ получение цепочки произвольных действий mentalInfoStruct.AmtzmNextStringID ????
- получение значения actionsImageID.

Эта функция - на всякий случай, т.к.
АВТОМАТИЗМ ВСЕГДА СОЗДАЕТСЯ ПРИ ЗАПУСКЕ mentalInfoStruct.AmtzmNextStringID на выполнение
	по результату периода ожидания - в func calcAutomatizmResult()

При каждом проходе генерировать 1 ассоциацию и примерять ее к актуальной доминанте.
*/
func infoFunc28(c *cycleInfo) {
	if c == nil {
		return
	}
	// инициативная произвольность возможна только в главном цикле
	if !c.isMainCycle {
		return
	}
	if mentalInfoStruct.AmtzmNextStringID == 0 {
		return
	}
	setCurIfoFuncID(c, 28)
	createAnvtzmFromVolutionString(c)
}
func fantasing(c *cycleInfo) {

	var mainDominanta *Dominanta
	// выбрать наиболее актуальную для detectedActiveLastProblemNodID доминанту
	maxWeight := 0 // пока выбор по весу, но нужно по detectedActiveLastProblemNodID
	for _, v := range DominantaProblem {
		if v == nil {
			continue
		}
		if v.weight > maxWeight {
			maxWeight = v.weight
			mainDominanta = v
		}
	}

	if maxWeight == 0 {
		return
	}
	CurrentProblemDominanta = mainDominanta

	effect := 1          // предположительный эффект
	actionsImageID := 0  // найденное действие для ассоциации с доминантами
	actionSucxessID := 0 // найденное действие успешного автоматизма, имеющего actionsImageID
	// в принципе actionsImageID == actionSucxessID

	/* генерация ассоцияций
	   ??? нужно бы учитывать: getCurSituation(), getCurTheme() и getCurPurpose()
	*/
	limit := 20
	n := 0
	for {
		actionsImageID = associationGeneration()
		if actionsImageID > 0 {
			// проверка, насколько данная ассоциаяция подходит для целей CurrentProblemDominanta
			if chechDominantPopose() {
				break
			}
		}
		if n > limit {
			return
		}
		n++
		continue
	}

	//попробовать решить подходящую по аналогии Домимнату с предположительным эффектом==effect
	res := checkRelevantAction(actionsImageID, curActions.ID, effect)
	if res > 1 {
		/* Эвристически озариться найденным решением
		success - 1 - доминанта решена, 2 - не точно, 3 - еще менее точно
		*/

		toConsciousHeuristics(actionsImageID, actionSucxessID, CurrentProblemDominanta)
	}
}

/*
	сгенерировать ассоциацию в виде ID действия

Правила - только один из механизмов поиска ассоциаций (и тут может применять прогноз по групповым правилам).
Причем использование своих правил - одно, а учительских - даст другое.
Другим способом нахождения ассоциаций м.б. перебор моделей значимых объектов,
чтобы отзеркалить действия, совпадающие с целью доминанты.
Ну и сама выборка правил м.б. на основе значимостей объектов, далеких от текущей ветки дерева,
но отвечающих цели.
*/
var oldAssociationsArr []*ActionsImage // не генерировать уже проверенные ассоциации
func associationGeneration() int {

	// TODO

	return 0
}

// проверка, насколько данная ассоциаяция подходит для целей CurrentProblemDominanta
func chechDominantPopose() bool {

	// TODO
	return true
}

///////////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////
/* Более внимательно рассмотреть ситуацию с Правилами: ментальными и связанными с ними моторными.
В результате найти полезное действие и
*/
func infoFunc29(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if !c.isMainCycle {
		return false
	}
	setCurIfoFuncID(c, 29)

	return lookForRulesMentalAndMotor(c)
}
func lookForRulesMentalAndMotor(c *cycleInfo) bool {
	actBestID := 0
	// найти кадры ментальной эпиз.памяти для данных условий с позитивным эффектом
	NodePID := detectedActiveLastProblemNodID // ID проблемы
	ThemeID := problemTreeInfo.themeID
	PurposeID := mentalInfoStruct.mentalPurposeID
	getEpisodesMentalArrFromConditions(NodePID, ThemeID, PurposeID, 1)
	if ePMentSrsIdArr == nil {
		return false
	}
	/* найти кадры моторной эпиз.памяти, связанные с ментальными по времени создания
	и сопоставить их
	*/
	for i := 0; i < len(ePMentSrsIdArr); i++ {
		motorID := EpisodicMentalHistoryArr[ePMentSrsIdArr[i]].lastEpisodicMemID
		if motorID > 0 {
			motorEP, ok := ReadeEpisodicTreeNodeFromID(motorID)
			if ok && motorEP != nil {

				// TODO придумать как сопоставить и найти лучший вариант actBestID
			}
		}
	}
	////////////////////////////////////////////

	if actBestID > 0 {
		actBest, ok := ReadeActionsImageArr(actBestID)
		if ok {
			/*получить пользу из эффекта лучшего моторного действия
			будет найден и запущен лучший мот.автоматизм aBest.ID
			mentalInfoStruct.motorAtmzmID=aBest.ID
			infoFunc17()// запустить автоматизм и завершить цикл осмысления
			*/
			getBenefitFromEpizosMemory(c, actBest)
			return true
		}
	}
	return false
}

/////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////
/* найти в неопределенной фразе что-то ранее понятное и выполнить действия
 */
func infoFunc300(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if !c.isMainCycle {
		return false
	}
	setCurIfoFuncID(c, 300)

	actsID := 0 // ищем эффективную акцию

	// CurretWordsIDarr - массив словID и вместо нераспознанного -1
	if len(wordSensor.CurretWordsIDarr) == 0 {
		return false
	}
	wArr := wordSensor.CurretWordsIDarr

	/* найти слово в стимуле правил для текущих условий и более общих, но не во всей эп.памяти
	оценивая значимость стимула с таким словом и игнорируя малую значимость.
	Вернуть акцию наиболее успешного правила
	*/
	actsID = getBestActionFromWordID(wArr)

	if actsID == 0 {
		wordsModel := getWordArrModelExactly(wArr)
		if wordsModel.beastFrameID > 0 { // ID наилучшего по эффекту кадра памяти
			actionNode, ok := ReadeEpisodicTreeNodeFromID(wordsModel.beastFrameID)
			if ok {
				actsID = actionNode.Action
			}
		}
	}

	/*  если бы был набор фраз, то можно было бы так:
	if isUnrecognizedPhraseFromAtmtzmTreeActivation && curActions.PhraseID != nil && len(curActions.PhraseID) > 1 {
			actsID := prepareUnrecognizedPhrase(curActions.PhraseID) // найти в неопределенной фразе что-то ранее понятное и выполнить действия

		}
	*/
	if actsID > 0 { // ID действия с наиболее позитивным Эффектом
		mentalInfoStruct.ActionsImageID = actsID
		runAutomatizmAfterCheck(c)
		isUnrecognizedPhraseFromAtmtzmTreeActivation = false
		return true
	}
	return false
}

/////////////////////////////////////////////////////////////////////
/* Анализировать длинную или недоопредеелнную фразу и найти обобщающие данные
 */
func infoFunc30(c *cycleInfo) bool {
	if c == nil {
		return false
	}
	if !c.isMainCycle {
		return false
	}
	setCurIfoFuncID(c, 30)

	if curActions.PhraseID != nil && len(curActions.PhraseID) > 1 {
		return preparelongPhrase(curActions.PhraseID) // TODO действия по найденным данным
	}
	// нараспознанная фраза
	if isUnrecognizedPhraseFromAtmtzmTreeActivation {
		return preparelongPhrase(curActions.PhraseID) // TODO действия по найденным данным
	}

	return false
}

// ///////////////////////////////////////////////////////////
/* Провокационные действия, когда нет стимула от оператора, но нужно что-то делать
не было стимула от оператора > 10 сек при значительном изменении гомео-параметров
*/
func infoFunc31(c *cycleInfo) bool {
	// !!! infoFunc31pulsCount = 0
	if c == nil {
		return false
	}
	if !c.isMainCycle {
		return false
	}
	idleness := isIdleness() // лень?
	if IsSleeping || !c.dreaming || idleness {
		return false
	}
	setCurIfoFuncID(c, 31)

	//if wasRunTreeStandardAutomatizm { // уже был запущен штатный автоматизм после Стимула.
	if LastRunAutomatizmPulsCount > 0 { // уже был запущен штатный автоматизм после Стимула - период ожидания
		return false
	}
	if wasRunPurposeActionFunc { // если ранее был запущен ментально в infoFunc17
		return false
	}

	// wasRunProvocationFunc используется как флаг "провокация func infoFunc31"
	wasRunProvocationFunc = true //сработала провокация оператора на действие, очистка - в clinerAutomatizmRunning()
	/*Выбрать лучшее Правило, как можно с более точным учетом условий, искать все виды Правил
	т.к. поиск идет при Стимуле==0, то может находится много правил, из них выбирается самое эффективное.
	При ответе оператора на провокацию будет записано Правила со Стимулом равным 0
	что можно использовать как выбор или избегание такого действия, но с учетом условия - действие без стимула опрератора.
	*/
	rule := getProvocationBestRule()
	if rule.Action > 0 {
		//по правилу найти или создать (в случае AmtzmNextString) автоматизм и запустить его
		//res := makeActionFromRooles(rule)
		ai, ok := ReadeActionsImageArr(rule.Action)
		if ok {
			purpose := getPurposeGenetic()
			purpose.actionID = ai
			levelOfRunAutomatizm = 3 // для передачи в пульт при ответе бота - на каком уровне осмысления был дан ответ
			createAndRunAutomatizmFromPurpose(purpose)
			if !isInterruptAutmtzm { // не выполнился, например если уже был запущен штатный автоматизм после Стимула и т.п.
				wasRunPurposeActionFunc = false //иначе не пропускает на исполнение
				// судя по всему, происходит наложение активации lib.SentActionsForPult(out) от рефлексов и автоматизмов, и выдает действия обоих.
				lib.NoReflexWithAutomatizm = true // не показывать акции рефлексов с автоматизмами в одной плашке действия Beast на пульте
				//action_sensor.DeactivationTriggers()
				//notAllowReflexRuning = true
				return true
			} else {
				wasRunProvocationFunc = false // раз автоматизм не запустился
			}
		}
	}

	return false
}

/////////////////////////////////////////////////////
//
//
//
//    КОНЕЦ ТЕЛ ИНФО_ФУНКЦИЙ
////////////////////////////////////////////////////////

/////////////// ПОДДЕРЖКА ИНФО_ФУНКЦИЙ ////////////////////////////////////////////////

// озарение при найденном решении в неглавном цикле (подсознании)
func insight(c *cycleInfo) {

	infoFunc19(c) //- было эвристическое озарение

	setAsMaimCycle(c.ID) //	сделать цикл главным
}

////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////
/* при наличии mentalInfoStruct.ActionsImageID > 0
создать автоматизм (если такого еще нет),
проверить его и запустить.
*/
func runAutomatizmAfterCheck(c *cycleInfo) {
	if mentalInfoStruct.ActionsImageID > 0 { // создать, проверить и запустить сразу, не в цикле.
		c.log += "По образу действий создать автоматизм, проверить его и запустить.<br>"
		if infoFunc7(c) { //создать новый моторный автоматизм по действию ActionsImageID
			// подготовить инфу для infoFunc6() из инфы от infoFunc7():
			mentalInfoStruct.motorAtmzmID = mentalInfoStruct.runMotorAtmzmID
			infoFunc6(c) // проверить и если норм. - запустить, сделав штатным
		}
	}
}

/////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////
/* случайный выбор ментальной функции, из тех, что еще не использовались в данном цикле (нет в infoFuncSequence)

В infoFuncSequence сохраняются активированные инфоID так что можно не вызывать то, что уже прошло в цикле.

TODO СЛУЧАЙНОЕ ДЕЙСТВИЕ НУЖНО ОГРАНИЧИВАТЬ ТЕМОЙ И ЦЕЛЬЮ!
*/
func infoFindRundomMentalFunction() int {
	//	return 13 // тестирование

	if transfer.IsPsychicТeachingMode {
		return 13
	}

	switch mentalInfoStruct.ThemeImageType {
	case 3: // Состояние Плохо
		return 122
	case 6: // Обучение с учителем
		return 13
	}

	clinerMentalInfo() // чтобы было случайное clinerMentalInfo() для 14-й функции

	// 11 и 13	отзеркаливание
	// весь набор допустимых функций повышенная вероятность для 13 и пониженная для 14
	var infoArr = []int{4, 9, 11, 11, 11, 12, 13, 13, 13, 14, 15, 16} // повышенная вероятность для 11 и 13  - отзеркаливание
	if EvolushnStage > 4 {                                            // повышенная вероятность для 11 - отзеркаливание
		infoArr = []int{4, 9, 11, 11, 11, 12, 15, 16}
	}
	if mentalInfoStruct.ThemeImageType == 5 { //Поисковый интерес
		infoArr = []int{11, 12, 14}
	}

	/* var actualArr []int
	чтобы func13 прошла нужно удалить ее код из functionsInAllCickles. Блокировка вызова 2 экземпляров делается через curFunc13ID
	иначе в functionsInAllCickles добавится 13 от какого то цикла и больше не даст запустить отзеркаливание не явным образом
	clinerFunctionsInAllCickles(13)
		for i := 0; i < len(infoArr); i++ {
			if !lib.ExistsValInArr(functionsInAllCickles, infoArr[i]) {
				actualArr = append(actualArr, infoArr[i])
			}
		}*/
	actualArr := listUnusedInfoId(infoArr) // только еще не использованные
	if actualArr != nil && len(actualArr) > 0 {
		funcID := lib.RandChooseIntArr(actualArr)
		/*не вызывалась ли такая случайная функция funcID в цикле? чтобы не вызывать подряд
		exists := lib.ExistsValInArr(functionsInAllCickles, funcID)
		var tryN = 20
		for exists && tryN > 0 {
			funcID := lib.RandChooseIntArr(actualArr)
			exists = lib.ExistsValInArr(functionsInAllCickles, funcID)
			tryN--
		}
		if !exists {*/
		if funcID > 0 {
			return funcID
		}
	}
	return 0
}

// список функций, котоорые еще не были использованы
func listUnusedInfoId(list []int) []int {
	var out []int
	for i := 0; i < len(list); i++ {
		if !lib.ExistsValInArr(infoFuncSequence, list[i]) {
			out = append(out, list[i])
		}
	}
	return out
}

////////////////////////////////////////////////

/* чтобы после отработки func13 infoFindRundomMentalFunction позволяло выдать новый код 13
нужно удалить из functionsInAllCickles коды прошлых вызовов 13
func clinerFunctionsInAllCickles(val int) {
	for id, n := range functionsInAllCickles {
		if n == val {
			functionsInAllCickles = lib.RemoveArrIndex(functionsInAllCickles, id)
			break
		}
	}
}*/

/*
*************** Группировка зеркальных автоматизмов *********************
Были зафиксированы две цепочки диалога по 2 шага каждая (1 шаг - 1 пара вопрос + ответ):
а) Привет – привет, как дела – нормально
б) Приветствую – привет, Ты как – отлично
Из них сформировались 2 зеркальных автоматизма: привет - как дела, привет - ты как. Их можно сгруппировать и вывести в отдельном
массиве варианты ответов на один пускатель: привет - как дела, ты как.
При поиске ответа нужно искать в этом массиве и выбирать варианты например по счетчику успешности. Это будет намного быстрее,
чем перебирать весь массив автоматизмов.
Для такой группировки нужно при создании нового зеркального автоматизма дописывать
в этот массив новый вариант в нужной строке: находить в массиве пусковое слово и добавлять к нему вариант ответа.
По сути это групповой запрос, только выделенный в динамическую таблицу. Так БД-шники часто делают, если приходится ворочать объемные
данные под миллионы записей. Вместо тяжелых тормозных запросов строятся буферные таблицы и забиваются через хранимки при совершении
операций с данными.
*/

//////////////////////////////////////////////////

///////////////////////////////////////////////////////

// при запуске любой инфо-функций - общие установки
func setCurIfoFuncID(c *cycleInfo, infofID int) {
	if c == nil {
		return
	}
	if infofID == 0 {
		return
	}
	// набор инфоID для func saveNewMentalEpisodic
	if !wasRunPurposeActionFunc { //После запуска автоматизма прекратить набор кадров и ждать ответа
		infoFuncSequence = append(infoFuncSequence, infofID)
	}

	c.lastFuncID = infofID
	c.func0Arr = append(c.func0Arr, infofID)
	c.funcArr = append(c.funcArr, infofID)

	if !lib.ExistsValInArr(functionsInAllCickles, infofID) {
		functionsInAllCickles = append(functionsInAllCickles, infofID)
	}
	// В лог - если это не функции из списка:
	if infofID != 1 && infofID != 2 && infofID != 5 && infofID != 8 && infofID != 17 {
		fName := getMentalFunctionString(infofID)
		if !c.idle && !c.isMainCycle || show_all_logs {
			c.log += "Запущена Инфо-функция № " + strconv.Itoa(infofID) + ": <i>\"" + fName + "\"</i><br>"
		}
	}

	if infofID == 17 {
		// вывод в функции
	}
	if infofID == 1 {
		//c.log+="Завершен цикл мышления.<br>"
	}

}

//////////////////////

/*
	Прдотвращние постоянного повторения ряда инфо-функций.

Если в цикле funcArr начинает повторяться последовательность выовов infofID,
то нужно пропустить вызов infofID начала такоей последовательности.
Последний фрагмент cycle.funcArr всегда == cycle.func0Arr
*/
func correctRepeatCycleFuncStr(cycle *cycleInfo) bool {
	if cycle.func0Arr == nil {
		return false
	}
	a0Arr := cycle.func0Arr
	aArr := cycle.funcArr
	len0F := len(a0Arr)
	lenF := len(aArr)
	// есть ли в предшествующем фрагменте cycle.funcArr повторение cycle.func0Arr
	if lenF > 2*len0F {
		// предшествующй фрагмент cycle.funcArr
		fArr := aArr[:(lenF - len0F)]
		prevF := fArr[(len(fArr) - len0F):]
		if lib.EqualArrs(prevF, a0Arr) {
			return true
		}
	}
	return false
}

// ////////////////////////////////////////////////////////////////////////
var infoFunc31pulsCount = 0
var waitingTimeBeforeProvocation = 50 // сколько ждет пульсов до следующей попытки провокации
// нужно провоцировать оператора
func isNeedForCommunication() bool {

	if infoFunc31pulsCount > 0 && (PulsCount-infoFunc31pulsCount) < 30 {
		// уже была провокация
		return false
	}
	if (PulsCount - curActiveActionsPulsCount) > waitingTimeBeforeProvocation { // прошло > 10 пульсов со времени последнего стимула от оператора
		if CurrentInformationEnvironment.veryActualSituation ||
			CurrentInformationEnvironment.danger ||
			CurrentInformationEnvironment.needThinkingAboutAutomatizm ||
			gomeostas.IsNeedForCommunication() {
			infoFunc31pulsCount = PulsCount
			return true
		}
	}
	return false
}

///////////////////////////////////////////////////////////////////////////

/*
получить пользу из эффекта лучшего моторного действия
будет найден и запущен лучший мот.автоматизм aBest.ID
mentalInfoStruct.motorAtmzmID=aBest.ID
infoFunc17()// запустить автоматизм и завершить цикл осмысления
*/
func getBenefitFromEpizosMemory(c *cycleInfo, actBest *ActionsImage) bool {
	if detectedActiveLastNodID == 0 { // так не должно быть..
		return false
	}
	// может быть несколько автоматизмов с одним и тем же действием, но с разными BranchID веками дерева автоматизмов или вообще не привязанные
	var eMax = 0
	var aBest *Automatizm

	//штатный автоматизм активного узла дерева
	sA := AutomatizmBelief2FromTreeNodeId[detectedActiveLastNodID]
	if sA != nil {
		if sA.ActionsImageID == actBest.ID {
			// как раз штатный автоматизм и делает это
			mentalInfoStruct.motorAtmzmID = sA.ID
			infoFunc17(c) // запустить автоматизм и завершить цикл осмысления
			return true
		}
	}

	//найти мот.автоматизмы с таким действием и сравинить со штатным автоматизмом активного узла detectedActiveLastNodID

	for _, v := range AutomatizmFromId {
		if v == nil {
			continue
		}
		if v.ActionsImageID != actBest.ID || v.Usefulness <= 0 {
			continue
		}
		if eMax < v.Usefulness {
			eMax = v.Usefulness
			aBest = v
		}
	}

	if aBest != nil {
		mentalInfoStruct.motorAtmzmID = aBest.ID
		infoFunc17(c) // запустить автоматизм и завершить цикл осмысления
		return true
	} else { // нет мот.автоматизм в конечном звене цикла, создать такой по действию actBest.ID
		//в infoFunc2() -> infoFunc7() -> infoFunc17()
		mentalInfoStruct.toAutmtzmActionsImageID = actBest.ID
	}

	return false
}

////////////////////////////////////////////

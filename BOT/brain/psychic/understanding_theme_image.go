/* Тема мышления - для выбора Цели
АКТУАЛЬНАЯ ТЕМА (mentalInfoStruct.ThemeImageType) изменятеся по событиям ThemeTypeStr
с перезапуском размышления runNewTheme(Weight,Type)
после чего в infoFunc8() выбирается цель mentalInfoStruct.mentalPurposeID

Тема определяется по многим разным объективным событиям (с помощью runNewTheme)
и в такой момент она запускает новый цикл размышления,
если только вес предыдущей темы не больше. В начале цикла по теме определяется цель (мтивация, потребность).
И циклы направляются в сторону достижения этой цели.
События изменения темы могут быть как при Стимуде (с Пульта), так и вне его,
так что сами по себе запускают новый цикл размышления.

Тема удаляется из ThemeImageFromID при решении проблемы или остается нерешенной и
тогда становится Доминантой.

ID темы используется в ментальных Правилах (как и ID Цели): вместо ID дерева автоматизмов+ситуации как в моторных Правилах, используются ID темы+Цели
*/

package psychic

import (
	"BOT/lib"
	"strconv"
	"strings"
)

/*
	какие бывают темы для ThemeImage.Type (многие из type SituationImage struct {)

фактически - перечисление проблем,
которые могут стать доминантой если ThemeImage.PulsCount старый и ThemeImage.Weight >3

Поиск по проекту детектора данной темы: "runNewTheme(N,"

1 "Негативный эффект моторного автоматизма"
2 "Негативный эффект ментального автоматизма"
3 "Состояние Плохо"
4 "Стимул с Пульта"
5 "Поисковый интерес"
6 "Обучение с учителем"
7 "Игнорирование оператором"
8 "Игра"
9 "Неудовлетворенность существущим"
10 "Непонимание"
11 "Разговор"
12 "Сомнение в штатном автоматизме"
13 "Защита"
14 "Страх"
15 "Агрессия"
16 "Есть объект высокой значимости"
17 "Улучшение настроения"
*/
var ThemeTypeStr = []string{
	"Нет темы", //0
	"Негативный эффект моторного автоматизма",   //1  -ЕСТЬ ДЕТЕКТОР runNewTheme(1,
	"Негативный эффект ментального автоматизма", //2 -ЕСТЬ ДЕТЕКТОР runNewTheme(2,
	"Состояние Плохо",          //3 Очень плохо -ЕСТЬ ДЕТЕКТОР runNewTheme(3,
	"Стимул с Пульта",          //4 -ЕСТЬ ДЕТЕКТОР runNewTheme(4,
	"Поисковый интерес",        //5 - любопытство, экспериментирование, набор опыта  -ЕСТЬ ДЕТЕКТОР в simeThemeDetection()
	"Обучение с учителем",      //6 -ЕСТЬ ДЕТЕКТОР в simeThemeDetection()
	"Игнорирование оператором", //7 оператор не ответил в течение периода ожидания на важный запрос  -ЕСТЬ ДЕТЕКТОР
	"Игра", //8  -ЕСТЬ ДЕТЕКТОР runNewTheme(8, -ЕСТЬ ДЕТЕКТОР в simeThemeDetection()
	"Неудовлетворенность существущим", //9  -ЕСТЬ ДЕТЕКТОР в simeThemeDetection()
	"Непонимание",                    //10  -ЕСТЬ один вид непонимания: runNewTheme(10,1)
	"Действие оператора",             //11 текстовое общение с оператором -ЕСТЬ ДЕТЕКТОР runNewTheme(11,
	"Сомнение в штатном автоматизме", //12  проблема mentalInfoStruct.motorAtmzmID > 0 -ЕСТЬ ДЕТЕКТОР runNewTheme(12,
	"Защита",   //13 -ЕСТЬ ДЕТЕКТОР в simeThemeDetection()
	"Страх",    //14 -ЕСТЬ ДЕТЕКТОР в simeThemeDetection()
	"Агрессия", //15 -ЕСТЬ ДЕТЕКТОР в simeThemeDetection()
	"Есть объект высокой значимости", // 16 -ЕСТЬ ДЕТЕКТОР runNewTheme(16,
	"Улучшение настроения",           //17 - базовая тема когда нет ничего другого, после просыпания и т.п.
}

// название типа темы
func GetThemeImageName(id int) string {
	return ThemeTypeStr[id]
}

//////////////////////////////////////////////

/* образ темы
 */
type ThemeImage struct {
	ID int
	// вес значимости для конкурентности 1 - 10 Используется только в runNewTheme( для конкурентности.
	Weight    int // типичное значение ==2 если не нужно чтобы эта тема доминировала
	Type      int //ID типов тем одно значение из ThemeTypeID
	PulsCount int // время актуализации темы
}

// var ThemeImageFromID=make(map[int]*ThemeImage)
var ThemeImageFromID []*ThemeImage // сам массив
// var AFromID = make([]*aNode, len(strArr))//задать сразу имеющиеся в файле число при загрузке
// запись члена
func WriteThemeImageFromID(index int, value *ThemeImage) {
	addThemeImageFromID(index)
	ThemeImageFromID[index] = value
}
func addThemeImageFromID(index int) {
	if index >= len(ThemeImageFromID) {
		newSlice := make([]*ThemeImage, index+1)
		copy(newSlice, ThemeImageFromID)
		ThemeImageFromID = newSlice
	}
}

// считывание члена
func ReadeThemeImageFromID(index int) (*ThemeImage, bool) {
	if index >= len(ThemeImageFromID) || ThemeImageFromID[index] == nil {
		return nil, false
	}
	return ThemeImageFromID[index], true
}

///////////////////////////////////////////////////////////////////

/*
	создать новый образ темы, если такой еще нет

createThemeImageID(id int,2,3,LifeTime,true)
*/
var lastThemeImageThemeID = 0

func createThemeImageID(id int, Weight int, Type int, PulsCount int, CheckUnicum bool) (int, *ThemeImage) {
	if CheckUnicum {
		oldID, oldVal := checkUnicumThemeImage(Type)
		if oldVal != nil {
			// обновить время создания и вес
			oldVal.PulsCount = LifeTime
			oldVal.Weight = Weight
			return oldID, oldVal
		}
	}

	if id == 0 {
		lastThemeImageThemeID++
		id = lastThemeImageThemeID
	} else {
		if lastThemeImageThemeID < id {
			lastThemeImageThemeID = id
		}
	}

	var node ThemeImage
	node.ID = id
	node.Weight = Weight
	node.Type = Type
	node.PulsCount = PulsCount

	//ThemeImageFromID[id]=&node
	WriteThemeImageFromID(id, &node)

	if doWritingFile {
		SaveThemeImageFromIdArr()
	}

	return id, &node
}
func checkUnicumThemeImage(Type int) (int, *ThemeImage) {

	for id, v := range ThemeImageFromID {
		if v == nil {
			continue
		}
		if Type != v.Type {
			continue
		}
		return id, v
	}
	return 0, nil
}

// сохранить образы
func SaveThemeImageFromIdArr() {
	var out = ""
	for k, v := range ThemeImageFromID {
		if v == nil {
			continue
		}
		out += strconv.Itoa(k) + "|"
		out += strconv.Itoa(v.Weight) + "|"
		out += strconv.Itoa(v.Type) + "|"
		out += strconv.Itoa(v.PulsCount)
		out += "\r\n"
	}
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/Theme_images.txt", out)
}

// загрузить образы, обнуляя вес старых тем
func loadThemeImageFromIdArr() {

	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/Theme_images.txt")
	cunt := len(strArr)
	//ThemeImageFromID=make(map[int]*ThemeImage)
	ThemeImageFromID = make([]*ThemeImage, cunt) //задать сразу имеющиеся в файле число
	for n := 0; n < cunt; n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		p := strings.Split(strArr[n], "|")
		id, _ := strconv.Atoi(p[0])
		weight, _ := strconv.Atoi(p[1])
		kind, _ := strconv.Atoi(p[2])
		pulsCount, _ := strconv.Atoi(p[3])
		// если возраст темы > 10дней - убрать ее конкурентный вес
		if LifeTime-pulsCount > 864000 { // //нельзя удалять тему т.к. она используется в ментральных Правилах
			//continue
			weight = 0
		}

		var saveDoWritingFile = doWritingFile
		doWritingFile = false
		createThemeImageID(id, weight, kind, pulsCount, false)
		doWritingFile = saveDoWritingFile
	}
	return
}

/*  Детектор появления новой темы - по событиям из ThemeTypeStr
Если появляется новая тема, но ее вес меньше старой, то думается про старую.

Тема меняется только с помощью этой функции!

Старое размышление не прерывается, а создается новый цикл.
*/

var oldThemeID = 0 // код карты ThemeImageFromID, это не ThemeImage.Type
func runNewTheme(kind int, weight int) {
	// только со стадии развития > 3
	if EvolushnStage < 4 {
		return
	}
	themeID, _ := createThemeImageID(0, weight, kind, LifeTime, true)
	/* новая тема может ли перекрыть старую?
	   т.е. исчерпана ли старая тема, имеет ли она меньший вес, чем новая?
	*/
	node, ok := ReadeThemeImageFromID(oldThemeID)
	if !ok {
		return
	}
	if oldThemeID > 0 && node.Weight > weight {
		return // остается старая тема, размышление не прерывается на новую
	}

	if oldThemeID != themeID { // если тема сменилась
		mentalInfoStruct.ThemeImageType = kind
		oldThemeID = themeID

		//инфа для активации дерева проблем:
		problemTreeInfo.themeID = themeID

		// !!!!  interruptMentalWork()
		cycle := createNewCycleIteration()
		themrCycle = cycle
		cycle.log += "Определена тема для мышления: " + ThemeTypeStr[kind] + "<br>"
		cycle.log += "Запущен новый цикл <span style='color:#006300' onclick='get_cycle_info(" + strconv.Itoa(cycle.ID) + ")'><b>" + strconv.Itoa(cycle.ID) + "</b></span><br>"
		cycle.themeType = kind
		//infoFunc8(cycle) // определить mentalInfoStruct.ThemeImageType и актиуивровать дерево проблем
		return
	}
	return
}

// жестко начать новую тему, даже если такая уже была перед этим
func NewTeme() {
	mentalInfoStruct.ThemeImageType = 0
	oldThemeID = 0
}

/////////////////////////////////////////////////////////

/*изменение веса темы после размышления или обнуление веса неактуальной темы:
 */
func ThemeAfterThinking(themeID int, weight int) {
	if weight < 0 { // удалить неактуальную тему
		//нельзя удалять тему т.к. она используется в ментральных Правилах
		//delete(ThemeImageFromID, themeID)
		//return
		weight = 0
	}
	if weight > 10 {
		weight = 10
	}
	//	ThemeImageFromID[themeID].Weight=weight
	node, ok := ReadeThemeImageFromID(themeID)
	if ok {
		node.Weight = weight
	}
}

/*
	детектор некоторых тем в зависимости от условий

Распознает тему и запускает ее итерацию
Вызов только из одного места: перед концом объективного вызова func consciousnessElementary( в understanding.go
*/
var themrCycle *cycleInfo

func someThemeDetection() bool {
	themrCycle = nil //или только перекрывать?
	//	if mentalInfoStruct.motorAtmzmID>0{ уже прошли 1 и 2 уровни
	//		return false
	//	}
	var themType = 0   // тип темы
	var themWeight = 0 // вес темы
	// плохо и опасно
	if PsyMood < -8 && CurrentInformationEnvironment.veryActualSituation && CurrentInformationEnvironment.danger {
		runNewTheme(10, 3)
		return true
	} else { // нет атаса
		if PsyMood > 0 && !CurrentInformationEnvironment.danger && !existsBaseContext(3) {
			if existsBaseContext(2) { // текущая эмоция - поиск
				runNewTheme(5, 5)
				return true
			}
			if existsBaseContext(3) { // текущая эмоция - игра
				runNewTheme(8, 2)
				return true
			}
			if existsBaseContext(2) && existsBaseContext(3) { // текущая эмоция - поиск+игра
				runNewTheme(8, 2) //- обучение
				return true
			}

		}
		if existsBaseContext(5) { //- защита
			runNewTheme(13, 5)
			return true
		}
		if existsBaseContext(8) { //- страх
			runNewTheme(14, 5)
			return true
		}
		if existsBaseContext(9) { //- агрессия
			runNewTheme(15, 3)
			return true
		}

		// Неудовлетворенность существующим
		if existsBaseContext(6) && existsBaseContext(11) { //- лень+доброта
			runNewTheme(9, 2) // ОТ ВЕСА МНОГОЕ ЗАВИСИТ
			return true
		}
	}

	// если есть какая-то проблема:
	if themType > 0 && themWeight > 0 {
		runNewTheme(themType, themWeight)
		return true
	}
	return false
}

///////////////////////////////////////////////////////

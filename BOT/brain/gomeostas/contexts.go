/*
	БАЗОВЫЕ КОНТЕКСТЫ (основные стили поведения) - contexts.go

1	Пищевой	- Пищевое поведение, восполнение энергии, на что тратится время и тормозятся антагонистические стили поведения.
2	Поиск	- Поисковое поведение, любопытство. Обследование объекта внимания, поиск новых возможностей.
3	Игра	- Игровое поведение - отработка опыта в облегченных ситуациях или при обучении.
4	Гон	- Половое поведение. Тормозятся антагонистические стили
5	Защита	- Оборонительные поведение для явных признаков угрозы или плохом состоянии.
6	Лень	- Апатия в благополучном или безысходном состоянии.
7	Ступор	- Оцепенелость при непреодолимой опастbase_context_activnostности или когда нет мотивации при благополучии или отсуствии любых возможностей для активного поведения.
8	Страх	- Осторожность при признаках опасной ситуации.
9	Агрессия	- Агрессивное поведение для признаков легкой добычи или защиты (иногда - при плохом состоянии).
10	Злость	- Безжалостность в случае низкой оценки .
11	Доброта	- Альтруистическое поведение.
12	Сон - Состояние сна. Освобождение стрессового состояния. Реконструкция необработанной информации.

Число одновеременно активных контекстов - не более 3-х!!! Остаются только с наибольшим весом,
а лишние будут отсеиваться в порядке убывания весов контекстов.
Это имитирует распознаватель с активацией по частично-активному профилю на входе.
*/
package gomeostas

import (
	"BOT/lib"
	"sort"
	"strconv"
	"strings"
)

func init() {
	//diff:=1
	//	diff = int(11.0 - 11.0 / math.Exp(float64(lib.Abs(diff)) * 0.17))
	//diff = int(11 - 11 / math.Exp(float64(lib.Abs(diff)) * 0.1))

	return
}

////////////////////////////////////////////////////////////

/*
	CurStyleImage удерживает активные стили с гистерезиом. Только CurStyleImage опредееляет, какие стили активны.

А BaseContextActive[] лишь промежуточный массив активных рвущихся на включение стилей.

На Пульте на странице редактора гомеостаза /pages/gomeostaz.php таблица "Активности Базовых стилей"
определяет порядок переключения сочетний базовых контекстов. Переключения происходят скачкообразно:
так, если значение параметра гомеостаза находилось около границ диапазона,
то пересечения границы флуктуацией параметра может вызвать дребезг - многократное срабатывание.
ПОэтому вводится порог переключения hysteresisLimitVal.
Это повышает стабильность удержания текущего CurStyleImage
(интерпретирует латеральное торможение конкурирующих сочетаний стилей).
Если задать hysteresisLimitVal=0, то стили будут переключаться точно по таблице "Активности Базовых стилей" (по-старому).

curGomeoPotencial вычисляется в функции getCurGomeoPotencial()
*/
var hysteresisLimitVal = 10 // значение порога (в % шкал) переключения CurStyleImage
type StyleImage struct {
	Styles    [3]int // максимум три текущих стиля
	potencial int    // значение curGomeoPotencial при обновлении CurStyleImage
}

var CurStyleImage StyleImage

// заполнить CurStyleImage новыми значениями (порядок весов стилей уже есть в BaseContextActive)
func setNewContexts() {
	n := 0
	for id, _ := range BaseContextActive {
		if BaseContextActive[id] {
			CurStyleImage.Styles[n] = id
			n++
			if n > 2 {
				break
			}
		}
	}
	CurStyleImage.potencial = curGomeoPotencial
}

var curGomeoPotencial = 0

/*
	текущее значениие curGomeoPotencial Корректруется по каждому пульсу.

Вычисляется в зависимости от состояния гомео-параметров (а не стилей поведения!).
При плавном изменении значений гомеопараметро, изменяется curGomeoPotencial
и при превышении CurStyleImage.potencial на порог  hysteresisLimitVal происходит переключение на новый дежурный стиль.
*/
func getCurGomeoPotencial() {
	// отслеживается только один приоритетный параметр гомеостаза PrioritetGomeoParID
	if PrioritetGomeoParID > 0 {
		curGomeoPotencial = 100 - int(GomeostazParams[PrioritetGomeoParID])
		if PrioritetGomeoParID == 1 {
			curGomeoPotencial = 100 - curGomeoPotencial
		}
		return
	}
	curGomeoPotencial = 0
	//если PrioritetGomeoParID - значит глухая норма, CurStyleImage будет для диапазона Норма

	/* вариант отслеживания по всем шкалам параметров нверное хуже*/

}

//////////////////////////////////////////////////////////////

// название Базового контекста из его ID str:=gomeostas.GetBaseContextCondFromID(id)
func GetBaseContextCondFromID(id int) string {
	var out = ""
	switch id {
	case 1:
		out = "Пищевой"
	case 2:
		out = "Поиск"
	case 3:
		out = "Игра"
	case 4:
		out = "Гон"
	case 5:
		out = "Защита"
	case 6:
		out = "Лень"
	case 7:
		out = "Ступор"
	case 8:
		out = "Страх"
	case 9:
		out = "Агрессия"
	case 10:
		out = "Злость"
	case 11:
		out = "Доброта"
	case 12:
		out = "Сон"
	}
	return out
}

//var IsGameMode=false

/*
	масссив активностей базовых контекстов

активный - true, неактивный - false
!!!ВНИМАНИЕ! на пульте результат изменений BaseContextActive не показывается без setNewContexts()
*/
var BaseContextActive [15]bool // index 0 НЕ ИСПОЛЬЗУЕТСЯ! т.е. начинаетсмя с BaseContextActive[1]
/* массив весов (значимостей) базовых контекстов */
var BaseContextWeight [15]int // index 0 НЕ ИСПОЛЬЗУЕТСЯ!

// антагонисты контекстов
var antagonists = make(map[int][]int)

/*
	Прошивка несовместимых сочетаний контекстов

Для каждого основного контекста - антагонисты
*/
func initContextDetector() {
	antagonists = make(map[int][]int)
	path := lib.GetMainPathExeFile()
	lines, _ := lib.ReadLines(path + "/memory_reflex/base_context_antagonists.txt")
	for i := 0; i < len(lines); i++ {
		p := strings.Split(lines[i], "|")
		id, _ := strconv.Atoi(p[0]) // ID параметра гомеостаза
		a := strings.Split(p[1], ",")
		for n := 0; n < len(a); n++ {
			aID, _ := strconv.Atoi(a[n])
			antagonists[id] = append(antagonists[id], aID)
		}
	}
	lines, _ = lib.ReadLines(path + "/memory_reflex/base_context_weight.txt")
	for i := 0; i < len(lines); i++ {
		p := strings.Split(lines[i], "|")
		id, _ := strconv.Atoi(p[0])
		val, _ := strconv.Atoi(p[1])
		BaseContextWeight[id] = val
	}
	for id, _ := range BaseContextWeight {
		BaseContextActive[id] = false
	}

	/* // проверка ограничителя
	BaseContextActive[1]=true
	BaseContextActive[6]=true
	BaseContextActive[9]=true
	BaseContextActive[12]=true
	BaseContextActive[9]=true
	var activedC=make(map[int]int)
	for id, v := range BaseContextActive {
		if v{
			activedC[id]=BaseContextWeight[id]
		}
	}
	//карта только активных контекстов
	keys := make([]int, 0, len(activedC))
	for k := range activedC {
		keys = append(keys, k)
	}
	//СОРТИРОВКА ПО ЗНАЧЕНИЮ даже если значения повторяются
	sort.SliceStable(keys , func(i, j int) bool {
		return activedC[keys[i]] > activedC[keys[j]]
	})
	// ограничить только первыми тремя
	if len(keys)>3 {
		keys = keys[:3]
	}
	for id, _ := range BaseContextActive {
		BaseContextActive[id]=false
		for i := 0; i < len(keys); i++ {
			if id == keys[i]{
				BaseContextActive[id]=true
			}
		}
	}
	*/
	return
}

/*
	состояние базового контекста зависит

1) от выхода из нормы жизненных параметров
2) от безусловно прошитых признаков восприятия
Антагонисты конкурируют между собой со своими весами значимости.

!!! BaseContextActive[6]=true - Ступор регулируется как отсуствие реакций в данном контексте при опасности

!!!ВНИМАНИЕ! на пульте результат изменений не показывается т.к. нет setNewContexts()
*/
func baseContextUpdate() { // корректруется по каждому пульсу

	if waitForReactivationPeriod() {
		return
	}

	getCurGomeoPotencial()

	// сначала все контексты выключены:
	//for id, _ := range BaseContextWeight {
	for id := 1; id < 9; id++ {
		//	if id != 12 { // сон не гасить
		BaseContextActive[id] = false
		//	}
	}
	// если активен сон, все остальные неактивны
	if IsSlipping {
		BaseContextActive[12] = true
		return // больше никаких других контекстов для активных действий
	} else {
		// При бодрствовании обязательно должен быть определен базовый контекст,
		// поэтому делаем по умолчанию Лень и гасим его в нужных случаях: BaseContextActive[6]=false
		// проверка в самом конце BaseContextActive[6] = true

		/* Нельзя утанавливать только один контект, это не позволит обучать в разных условиях!
		if IsGameMode {
			BaseContextActive[3] = true
			return // больше никаких других контекстов для активных действий
		}*/
	}

	// Определяем текущее сочетания активных базовых контекстов:
	for id := 1; id < 9; id++ {
		gp := int(GomeostazParams[id])
		// в каком диапазоне находтся жизненный параметр? Выше или инже порога?
		isNorma := false                         // статическая норма
		if id == 1 && gp >= compareLimites[id] { // для энергии
			isNorma = true
		}
		if id > 1 && gp < compareLimites[id] { // для остальных
			isNorma = true
		}
		if isNorma {
			// правило для данного диапазона
			rule := GomeostazActivnostArr[id][2] // норма
			// активируем или пассивируем контексты по заданному правилу в http://go/pages/gomeostaz.php
			activeOrPassiveContext(rule)
			continue // при явной норме не устанавливаются другие контексты
		} else { // статические диапазоны отклонения от нормы
			diapazoN := getBadDiapazon(id)
			rule := GomeostazActivnostArr[id][diapazoN] // диапазон ухудшения
			// активируем или пассивируем контексты по заданному правилу в http://go/pages/gomeostaz.php
			activeOrPassiveContext(rule)
		}

		/*
			if BadNormalWell[i] == 2 { // НОРМА
				// правило для данного диапазона
				rule := GomeostazActivnostArr[i][2] // норма
				// активируем или пассивируем контексты по заданному правилу в http://go/pages/gomeostaz.php
				activeOrPassiveContext(rule)
			}else{// установившееся значение Плохо, .а не динамические, которые учитываются внизу
				diapazoN := getBadDiapazon(i)
				rule := GomeostazActivnostArr[i][diapazoN] // диапазон ухудшения
				// активируем или пассивируем контексты по заданному правилу в http://go/pages/gomeostaz.php
				activeOrPassiveContext(rule)
			}
		*/

		// ДИНАМИЧЕСКИЕ ИЗМНЕНИЯ (т.е. моменты улучшения или ухудшения) перекрывают предыдущее
		if BadNormalWell[id] == 3 { // ХОРОШО
			rule := GomeostazActivnostArr[id][1] // хорошо
			// активируем или пассивируем контексты по заданному правилу в http://go/pages/gomeostaz.php
			activeOrPassiveContext(rule)
		}
		// перекрывает предыдущее
		if BadNormalWell[id] == 1 { // ПЛОХО
			rule := GomeostazActivnostArr[id][0]
			// активируем или пассивируем контексты по заданному правилу в http://go/pages/gomeostaz.php
			activeOrPassiveContext(rule)
		}

	}

	// Конкурентность антагонистов
	keys := make([]int, 0, len(BaseContextActive))
	for id, v := range BaseContextActive {
		if v {
			keys = append(keys, id)
		}
	}
	sort.Ints(keys)
	for _, id := range keys {
		for _, ida := range antagonists[id] {
			if BaseContextActive[ida] && BaseContextWeight[ida] > BaseContextWeight[id] { // активный антагонист значимее
				BaseContextActive[id] = false // погасить текущий контекст
				break                         // больше не нужно смотреть других антагонистов
			} else {
				BaseContextActive[ida] = false // погасить антогониста
			}
		}
	}
	/* ограничение на число компонентов в образе б.контекстов: их число будет не более 3-х,
	а лишние будут отсеиваться в порядке убывания весов контекстов.
	Это неплохо имитирует распознаватель с активацией по частично-активному профилю на входе.
	*/
	var activedC = make(map[int]int)
	for id, v := range BaseContextActive {
		if v {
			activedC[id] = BaseContextWeight[id]
		}
	}

	////////////////////////////////////////
	// карта только активных контекстов
	keys = make([]int, 0, len(activedC))
	for k := range activedC {
		keys = append(keys, k)
	}
	// СОРТИРОВКА ПО ЗНАЧЕНИЮ даже если значения повторяются
	sort.SliceStable(keys, func(i, j int) bool {
		return activedC[keys[i]] > activedC[keys[j]]
	})
	// ограничить только первыми тремя
	if len(keys) > 3 {
		keys = keys[:3]
	}
	for id, _ := range BaseContextActive {
		BaseContextActive[id] = false
		for i := 0; i < len(keys); i++ {
			if id == keys[i] {
				BaseContextActive[id] = true
			}
		}
	}
}

////////////////////////
/* в каком из 5 диапазоне Плохо находится Базовый параметр, если он вне Нормы
0 - предельно плохо
      2 Норма   - getBadDiapazon не вызывается при норме
3 Слабое отклонение
4 Значительное отклонение
5 Сильное отклонение
6 Критически опасное отклонение
*/
func getBadDiapazon(pID int) int {
	gp := int(GomeostazParams[pID]) // текущее значение жизненного параметра
	limit := compareNorma[pID]      // порог начала отклонения параметров из нормы

	var bad = 0
	if pID == 1 {
		bad = limit     // остаток параметра вне критического
		gp = limit - gp // убираем критическую часть
	} else {
		//bad = limit // остаток параметра вне критического
		bad = 100 - limit // остаток параметра вне критического
		gp = gp - limit
	}
	// какой процент составляет gp
	proc := int((gp * 100) / bad)
	if proc < 20 {
		return 3
	}
	if proc < 40 {
		return 4
	}
	if proc < 60 {
		return 5
	}
	if proc < 80 {
		return 6
	}

	return 0 // предельно плохо
}

/////////////////////////////////////////////////////////////////////

// активируем или пассивируем контексты по заданному правилу в http://go/pages/gomeostaz.php
var curPainValue = 0
var curJoyValue = 0

// !!!ВНИМАНИЕ! на пульте результат изменений не показывается т.к. нет setNewContexts()
func activeOrPassiveContext(rule string) {

	if waitForReactivationPeriod() {
		return
	}

	if len(rule) == 0 {
		return
	}
	// выделяем ID контекстов
	p := strings.Split(rule, ",")
	for n := 0; n < len(p); n++ {
		p[n] = strings.TrimSpace(p[n])
		if len(p[n]) == 0 {
			return
		}
		// активируем или пассивируем контексты
		cID, _ := strconv.Atoi(p[n])
		if cID > 19 && cID < 30 {
			curPainValue = cID - 20
			curJoyValue = 0
			continue
		}
		if cID > 29 {
			curJoyValue = cID - 30
			curPainValue = 0
			continue
		}
		if cID > 0 {
			BaseContextActive[cID] = true
		} else {
			BaseContextActive[-cID] = false
		}
	}
}

///////////////////////////////////////////////////

// для определения текущего сочетания ID Безовых контекстов gomeostas.GetCurContextActiveIDarr()
func GetCurContextActiveIDarr() []int {
	var out []int
	/*
		// concurrent map iteration and map write
		for id, v := range BaseContextActive {
			if v {
				out = append(out, id)
			}
		}*/
	for i := 0; i < 3; i++ {
		out = append(out, CurStyleImage.Styles[i])
	}
	return out
}

// для Пульта
func GetCurContextActive() string {
	var out = ""
	/*
		for id, v := range BaseContextActive {
			if v {
				out += strconv.Itoa(id) + ";1|"
			} else {
				out += strconv.Itoa(id) + ";0|"
			}
		}*/
	for id, _ := range BaseContextActive {
		exists := false
		for i := 0; i < 3; i++ {
			if CurStyleImage.Styles[i] == id { // есть такой активный контекст
				out += strconv.Itoa(id) + ";1|"
				exists = true
				break
			}
		}
		if !exists {
			out += strconv.Itoa(id) + ";0|"
		}
	}
	return out
}

/*
// контекст распознавания текущей фразы с Пульта для Vernike_detector.go
func GetActiveContextInfo() map[int]int {
	var activeCW = make(map[int]int)

	keys := make([]int, 0, len(BaseContextActive))
	for id, v := range BaseContextActive {
		if v {
			keys = append(keys, id)
		}
	}
	sort.Ints(keys)
	for _, k := range keys {
		activeCW[k] = BaseContextWeight[k]
	}
	return activeCW
}*/

var IsNewConditions = false    // флаг изменения условий (пусковых стимулов)
var oldBaseCondition = 0       // предыдущее базовое состояние
var oldActiveContextstStr = "" // строка ID старого сочетания активных Базовых контекстов

/*
	переактивация контекстов в зависимости от состояния гомео-параметров

детектор изменения базового состояния и контекстов - проверка по каждому пульсу
*/

func changingConditionsDetector() {

	if waitForReactivationPeriod() {
		return
	}

	if oldBaseCondition != CommonBadNormalWell {
		oldBaseCondition = CommonBadNormalWell

		correctPainJoy() // скорректировать состояния Боль, Радость

		IsNewConditions = true
		return
	}

	var activeContextstStr = ""
	keys := make([]int, 0, len(BaseContextActive))
	for k := range BaseContextActive {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for k, v := range keys {
		if BaseContextActive[v] {
			activeContextstStr += strconv.Itoa(k) + "_" // "_" нужно разделять цифры, иначе будет ошибаться
		}
	}
	if oldActiveContextstStr != activeContextstStr {
		// можно переключать контексты если увеличение или уменьшение curGomeoPotencial превысит порог
		if lib.Abs(CurStyleImage.potencial-curGomeoPotencial) > hysteresisLimitVal {
			// заполнить CurStyleImage новыми значениями по порядку весов стилей
			setNewContexts()

			oldActiveContextstStr = activeContextstStr
			IsNewConditions = true
			return
		} /////////////////////////////////

	}
	IsNewConditions = false
}

/* Есть ли хоть какая то активность контекстов
func IsContextActive()bool {
  for _, v :=range BaseContextActive {
   if v { return true }
 }
 return false
}


// для произвольного задания текущего сочетания ID Безовых контекстов gomeostas.SetCurContextActiveIDarr()
func SetCurContextActiveIDarr(c []int) {
	// concurrent map iteration and map write
	for i := 0; i < len(c); i++ {
		BaseContextActive[c[i]]=true
	}
}*/
////////////////////////////////////////////////////////

// скорректировать состояния Боль, Радость при изменении контекcтов
func correctPainJoy() {
	GomeostazActionEffectPainV = curPainValue // величина Боли
	GomeostazActionEffectJoyV = curJoyValue   // величина радости
}
func GetCurPainJoy() (int, int) { // запрос с психики
	return GomeostazActionEffectPainV, GomeostazActionEffectJoyV
}

//////////////////////////////////////////////////////////

// //////////////////////////////////////////////////////
var noReactivationFromStimuls = false // true чтобы не работали функции переактивации
/*
вряд ли переактивация может быть продолжительной из-за changingConditionsDetector() и ContextActiveFromStimul() которые опять изменять контексты
Но на время keepingContextTime changingConditionsDetector() блокируется
а ContextActiveFromStimul() блокируется только для незначительных
*/
var keepingContextTime = 20
var keepingContextPulsCount = 0

func waitForReactivationPeriod() bool {
	if keepingContextPulsCount == 0 {
		return false
	}
	if PulsCount-keepingContextPulsCount < keepingContextTime {
		// на время keepingContextTime сохраняются контексты после переактиваций
		return true
	}
	keepingContextPulsCount = 0
	baseContextUpdate()
	changingConditionsDetector() //переактивация контекстов в зависимости от состояния гомео-параметров
	setNewContexts()             // заполнить CurStyleImage новыми значениями по порядку весов стилей
	return false
}

///////////////////////////////////////////////////////////

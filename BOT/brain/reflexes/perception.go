/*
восприятие действий и фраз с Пульта

Все рефлекторные и автоматические активности начинаются отсюда.
_____________________________
Сначала активируется Дерево рефлексов и собираются рефлексы на выполнение, но пока не выполняются,
потом активируется Дерево автоматизмов и собираются автоматизмы на выполнение, но пока не выполняются,
если возникает ориентировочный рефлекс, то
активируется Дерево понимания (ментальных автоматизмов) и решается, что делать дальше.
если нет ориентировочного рефлекса, то
потом выполняются автоматизмы, если их нет - то рефлексы.
________________________________
Создание образов различной иерархии контексfunc ActiveFromActionтов восприятия:
BaseStyleArr - образ сочетаний активных Базовых контекстов
TriggerStimulsArr - образ сочетаний пусковых стимулов
*/

package reflexes

import (
	"BOT/brain/action_sensor"
	"BOT/brain/gomeostas"
	"BOT/brain/psychic"
	"BOT/lib"
	"strconv"
)

// Создание образов сочетаний ID действий:
func loadImagesArrs() {
	// загрузить образы сочетаний базовых стилей
	loadBaseStyleArr()
	// загрузить образы сочетаний пусковых стимулов
	loadTriggerStimulsArr()
}

// АКТИВАЦИЯ ДЕРЕВА РЕФЛЕКСОВ ПО изменению условий, действиям с Пульта или фразе с Пульта

/*
	Вид активации дерева рефлексов:

1 - изменение сочетания базовых контекстов
2 - действия с Пульта
3 - фраза с Пульта
*/
var ActivationTypeSensor = 0

// текущее восприятие ID образов
// обновляющихся при каждом событии с Пульта или достаточно сильном изменении Базовых параметров
var ActiveCurBaseID = 0           // ID Базового состояния CommonBadNormalWell
var ActiveCurBaseStyleID = 0      // ID сочетания базовых контекстов BaseStyle
var ActiveCurTriggerStimulsID = 0 // ID теущего активного образа сочетаний пусковых стимулов TriggerStimuls
/*
	предыдущий образ сочетания пусковых стимулов

используется как ПРИЧИНА последующих событий - для формирования условных рефлексов
потому как ОБНУЛЯЕТСЯ при:
1) активации дерева рефлексов, если вызвало какое-то действие
2) через 10 пульсов после записи значения - типа причина устаревает
*/
var oldActiveCurTriggerStimulsID = 0

// момент записи значения в тике Пульса
var oldActiveCurTriggerStimulsPulsCount = 0

// Сохранить предыдущий образ сочетаний пусковых стимулов
func setOldActiveCurTriggerStimulsVal(val int) {
	oldActiveCurTriggerStimulsID = val
	oldActiveCurTriggerStimulsPulsCount = ReflexPulsCount
}

/* Активация дерева рефлексов при любом изменении условий с проверкой по каждому пульсу. */
func ActiveFromConditionChange() {
	if IsSlipping { // блокировака восприятия во время сна
		return
	}
	IsPultActionThisPuls = true

	// :) это не сработает при прерывании точной останова т.к. за это время ReflexPulsCount увеличится!
	if activetedPulsCount == ReflexPulsCount { // ждет следующего пульса
		return
	}
	// очищать прежние акции с пульта при смене сочетания Базовых контекстов.
	action_sensor.DeactivationTriggers()

	activetedPulsCount = ReflexPulsCount
	ActivationTypeSensor = 1

	ActiveCurBaseID = gomeostas.CommonBadNormalWell

	// определение текущего сочетания ID Базовых контекстов
	bsIDarr := gomeostas.GetCurContextActiveIDarr()

	// создаем новый образ Базовых контекстов, если такого еще нет
	ActiveCurBaseStyleID, _ = createNewBaseStyle(0, bsIDarr, true)

	// нет пусковых стимулов с Пульта, ActiveCurTriggerStimulsID обнуляем
	//чтобы при активации деревьев было ясно, что нет Стимула с Пульта.
	ActiveCurTriggerStimulsID = 0

	// активировать дерево рефлексов
	activeReflexTree()
	//  lib.SentActionsForPult("1|WWWWWWWW") // запустить рефлекс для тестирования

	// активировать дерево автоматизмов
	psychic.WasConditionsActiveted = true
	res := psychic.SensorActivation(ActivationTypeSensor)

	if res { // блокировать выполнение рефлексов
		if len(oldReflexesIdArr) > 0 || len(geneticReflexesIdArr) > 0 {
			lib.WritePultConsol("<span style='color:red'>Рефлекс <b>заблокирован</b></span>")
		}
		action_sensor.DeactivationTriggers()
		return
	}
	// запустить рефлексы
	toRunRefleses()

	// сбросить контекст акций по кнопкам Пульта
	action_sensor.DeactivationTriggers()
	psychic.WasConditionsActiveted = false
}

/*
	активировать дерево автоматизмов ТОЛЬКО действием в main.go reflexes.ActiveFromAction()

Если кроме действий будет и фраза, то сработает func ActiveFromPhrase()
*/
func ActiveFromAction() {
	if IsSlipping { // блокировака восприятия во время сна
		return
	}

	//	activeGameMode()//Активировать Игровой режим
	IsPultActionThisPuls = true
	/* Нужно сбрасывать буфер фраз после каждого действия кнопок потому, что действие означает оценку вербального автоматизма и завершение перебора вариантов
	Сюда перенесено так как тогда будет учитываться воспитательные действия любых кнопок, а не только Поощрить/Наказать
	*/
	psychic.UsedPraseIdArr = nil

	// определение текущего сочетания ID Базовых контекстов
	bsIDarr := gomeostas.GetCurContextActiveIDarr()

	//переактивация контекстов в зависимости от эффекта стимула, сопровождающихся болью и радостью
	rSarr := action_sensor.CheckCurActions()
	if rSarr != nil {
		effect := 0
		for i := 0; i < len(rSarr); i++ {
			ef := gomeostas.GomeostazActionCommonEffectArr[rSarr[i]] // значение эффекта Боли и Радости
			val, _ := strconv.Atoi(ef)
			effect += val
		}
		if lib.Abs(effect) > 5 {
			gomeostas.ContextActiveFromStimul(effect)
		}
	}

	ActivationTypeSensor = 2
	// при воздействии кнопками Наказать, Поощрить - записывать Usefulness ранее выполненного автоматизма
	// !!! psychic.LastAutomatipmCorrection() эти кнопки могут сопровождать и фразу! нефиг такие эксклюзивы делать!

	if activetedPulsCount == ReflexPulsCount { // ждет следующего пульса
		return // :) это не сработает при прерывании точкой останова т.к. за это время ReflexPulsCount увеличится!
	}
	activetedPulsCount = ReflexPulsCount

	ActiveCurBaseID = gomeostas.CommonBadNormalWell

	// создаем новый образ Базовых контекстов, если такого еще нет
	ActiveCurBaseStyleID, _ = createNewBaseStyle(0, bsIDarr, true)

	// создаем новый образ Пусковых стимулов ActiveCurTriggerStimulsID (TriggerStimuls.ID), если такого еще нет
	CreateNewTriggerStimulsImage()

	// активировать дерево рефлексов
	activeReflexTree()

	/* Это используется для определения момента реакция оператора Пульта на действия автоматизма - для психики.
	За 20 сек г.параметры могут просто натечь и сработает ожидание реакции оператора.
	Флаг сбрасывается через пульс после запуска автоматизма.
	*/
	psychic.WasOperatorActiveted = true

	// активировать дерево автоматизмов
	res := psychic.SensorActivation(ActivationTypeSensor)
	if res { // блокировать выполнение рефлексов
		if len(oldReflexesIdArr) > 0 || len(geneticReflexesIdArr) > 0 {
			lib.WritePultConsol("<span style='color:red'>Рефлекс <b>заблокирован</b></span>")
		}
		action_sensor.DeactivationTriggers()
		return
	}

	toRunRefleses()

	// сбросить контекст акций по кнопкам Пульта
	action_sensor.DeactivationTriggers()
}

// активировать дерево фразой  reflexes.ActiveFromPhrase()
func ActiveFromPhrase() {
	if IsSlipping { // блокировака восприятия во время сна
		return
	}

	IsPultActionThisPuls = true
	//	activeGameMode()
	//если вместе с вербальным стимулом идет действие, значит оно оценивает предыдущий ответ, и надо сбрасывать UsedPraseIdArr() - см. выше, почему
	if curPultActionsArr != nil {
		psychic.UsedPraseIdArr = nil
	}

	// :) это не сработает при прерывании точной останова т.к. за это время ReflexPulsCount увеличится!
	if activetedPulsCount == ReflexPulsCount { // ждет следующего пульса
		return
	}
	activetedPulsCount = ReflexPulsCount
	ActivationTypeSensor = 3

	ActiveCurBaseID = gomeostas.CommonBadNormalWell
	// определение текущего сочетания ID Базовых контекстов
	bsIDarr := gomeostas.GetCurContextActiveIDarr()
	// создаем новый образ Базовых контекстов, если такого еще нет
	ActiveCurBaseStyleID, _ = createNewBaseStyle(0, bsIDarr, true)

	// создать новое сочетание пусковых стимулов если такого еще нет
	CreateNewTriggerStimulsImage()

	// активировать дерево рефлексов
	activeReflexTree()

	/* Это используется для определения момента реакция оператора Пульта на действия автоматизма - для психики.
	За 20 сек г.параметры могут просто натечь и сработает ожидание реакции оператора.
	Флаг сбрасывается через пульс после запуска автоматизма.
	*/
	psychic.WasOperatorActiveted = true
	// убираем активацию рефлекса бездействия, чтобы его действия не цеплялись к создаваемому автоматизму
	// иначе сутуация на 2 стадии в процессе отзеркаливания стимула и создания автоматизма:
	// посылаем "привет" - получаем ответ: "привет" + [предлагает поиграть]
	if action_sensor.CheckCurActions() == nil {
		psychic.DeleteActualRelextActon()
	}

	// активировать дерево автоматизмов
	res := psychic.SensorActivation(ActivationTypeSensor)
	if res { // блокировать выполнение рефлексов
		if len(oldReflexesIdArr) > 0 || len(geneticReflexesIdArr) > 0 {
			lib.WritePultConsol("<span style='color:red'>Рефлекс <b>заблокирован</b></span>")
		}
		action_sensor.DeactivationTriggers()
		return
	}

	toRunRefleses()

	// сбросить контекст акций по кнопкам Пульта
	action_sensor.DeactivationTriggers()
}

/* Активация Игрового режима действием или фразой
Активация Игрового режима во всех стадиях происходит автоматически при первом стимуле в контексте Поиск или Игра
сброс актвиации для 1 стадии делается в ReflexCountPuls() при превышении времени ожидания рефлекторного ответа 20 сек
сброс активации для 2,3 стадий в делается в automatizmActionsPuls() при превышении времени ожидания ответа WaitingPeriodForActionsVal
В 4 стадии это так же можно делать произвольно, при помощи IsArbitraryGameMode: True - активировать игровой режим. Этот же флаг, пока True, не дает игровому режиму погаснуть в automatizmActionsPuls()
То есть с 4 стадии просто добавляется возможность ПРОИЗВОЛЬНО активировать и дезактивировать Игровой режим в любом контексте, а его автоматическая активация по контексту так и продолжается.

Но так как SetActionFromPult() в котором происходит изменение общего состояния активируется ПЕРЕД activeGameMode(), то при первой активации игрового режиме
возможно изменение состояния, что необходимо для естественной реакции Beast.

Во всех случаях имеет место специфический механизм УДЕРЖАНИЯ ОБРАЗА ТЕКУЩЕГО СОСТОЯНИЯ - важнейшая часть функционала ОР, который необходим для корректного формирования новых у-рефлексов и автоматизмов в одном контексте,
что особенно критично в начальных стадиях - иначе запишется сумбур. А начиная с 4 стадии он нужен чтобы не писать сумбурные групповые правила, где собираются в одну группу разные по конечному эффекту цепочки.

Активировать Игровой режим в baseContextUpdate() просто по факту наличия контекстов Поиск или Игра не нужно потому, что есть еще и внутренние реакции, не имеющие никакого отношения к игре. Например рефлексы на
отсутствие реакции со стороны Оператора. Так же имеет место изменение уровней Жизненных параметров при пульсации, что изменяет контексты. Такие естественные изменения контекстов
нужно оставить, чтобы прототип был более близок к природной реализации.

func activeGameMode(){
	if psychic.IsArbitraryGameMode || gomeostas.BaseContextActive[2] || gomeostas.BaseContextActive[3]{
		if true {// отключение для тестирования
			transfer.IsPsychicGameMode = true
			psychic.LimitGroupRules++
		}
	}
}*/

/*
	создание иерархии образов контекстов условий и пусковых стимулов в виде ID образов в [3]int
	создать последовательность уровней условий в виде массива  ID последовательности ID уровней

В случае отсуствия пусковых стимулов создается ID такого отсутсвия, пример такой записи: 2|||0|0|
*/
func getConditionsArr(lev1ID int, lev2 []int, lev3 []int, PhraseID []int, ToneID int, MoodID int) []int {
	arr := make([]int, 3)
	arr[0] = lev1ID
	arr[1], _ = createNewBaseStyle(0, lev2, true)
	arr[2], _ = CreateNewlastTriggerStimulsID(0, lev3, PhraseID, ToneID, MoodID, true)
	return arr
}

// получить сохраненное (последнее активное) сочетание пусоквых стимулов-кнопок
// reflexes.GetCurPultActionsContext()
func GetCurPultActionsContext() []int {
	var ActID []int
	if ActiveCurTriggerStimulsID > 0 {
		//ActID = TriggerStimulsArr[ActiveCurTriggerStimulsID].RSarr
		node, ok := ReadeTriggerStimulsArr(ActiveCurTriggerStimulsID)
		if !ok {
			return nil
		}
		ActID = node.RSarr
	}
	return ActID
}

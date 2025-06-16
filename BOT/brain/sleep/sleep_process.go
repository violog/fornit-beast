/* процесс погружения в сон и его стадии и циклы

Период сновидений (обработки оставшейся со дня информации) начинается с активации
CurrentInformationEnvironment - текущего образа информационной среды - выявляются актуальные проблемы
и активные доминанты нерешенной проблемы.
Это становится контекстом развития экспериментирования с оставшимися со дня образами.


при повышении активности "Хочу спать"
Механизм преждевременного пробуждения.

В каждом пакете есть флаг  и там для сна выполняется
	if IsSlipping {
		sleepingProcess()
	}
в психике для этого есть psychic_sleep_process.go


Сон начинается с перевода главного цикла мышления в пассивный режим (сновидения)
Если есть инсайт, то просырание и все заново,
если нет инсайта, начинать по очереди фоновые циклы переводить в главные в пассивном режиме.
Это - на уровне не func consciousnessElementary(), а в func dispetchConsciousnessThinking()->func consciousnessThinking
*/

package sleep

import (
	"BOT/brain/psychic"
	"BOT/lib"
)

// true - это СОН  sleep.IsSleeping
var IsSleeping = false
var SleepingType = 0 // 0 - нет сна, 1 - глубокий сон, 2- сновидение !Фаза сновидений реализуется на уровне психики!

func GetSleepCondition() (bool, int) {
	return IsSleeping, SleepingType
}

// ///////////////////////////////////
var sleepNecessityValue = 0 // параметр необходимости сна
/*
	sleepNecessityValue повышается при бодрствовании от разных причин

и уменьашется во сне, когда эти причины восстанавливаются (редуцируется память и т.п.)
При значении sleepNecessityValue >50 и отсуствии значимых раздражителей наступает сон.
При значении > 100 сон наступает несмотря ни на какие раздражители.

Детектор необходимости сна работает по каждому пульсу.
*/
func sleepNecessityDetector() {

	if IsSleeping {
		return
	}

	sleepNecessityValue = 0 // определяется заново с каждым пульсом

	// число параллельных циклов мышления
	cicleCount := psychic.GetCycleCount()
	sleepNecessityValue += getSNweight(cicleCount, 1000, 0.1) //

	// массив распознанных фраз
	//	words:=len(word_sensor.MemoryDetectedArr)
	//	sleepNecessityValue+=getSNweight(words,1000,0.1)// при превышении на 1000 sleepNecessityValue достигнет 100

	// массив прошлых состояний информационного окружения
	infs := psychic.GetInformationEnvironmentObjectsLength()
	sleepNecessityValue += getSNweight(infs, 100000, 0.001) // при превышении на 100000 sleepNecessityValue достигнет 100

	// массив эпизодической памяти
	/* TODO epizM:=len(psychic.EpisodeMemoryObjects)
	sleepNecessityValue+=getSNweight(epizM,100000,0.001)// при превышении на 100000 sleepNecessityValue достигнет 100
	*/
	// масив curImportanceObjectArr объектов (а не адресов)
	//impM:=psychic.GetcurImportanceObjectArrLength()
	//sleepNecessityValue+=getSNweight(impM,100000,0.001)// при превышении на 100000 sleepNecessityValue достигнет 100

	if sleepNecessityValue > 100 {
		BeginSleepСondition()
	}
}

/*
	рассчитать добавку для sleepNecessityValue для параметра, вызывающего сон (param),

действие которого начинается со значения параметра = beginLimit
умноженного на коэффициент (k) параметра для необходимости сна.
*/
func getSNweight(param int, beginLimit int, k float64) int {
	balance := param - beginLimit
	if balance <= 0 {
		return 0
	}
	val := k * float64(balance)
	res := lib.RoundToInt(val)
	return res
}

///////////////////////////////////

/*
начать сон sleep.BeginSleepСondition()
*/
func BeginSleepСondition() {
	// можно ли заснуть
	if !IsPossibleToSleep() || !psychic.IsPossibleToSleep() { // нельзя спать
		return
	}

	IsSleeping = true
	SleepingType = 2 // сразу же главный цикл мышления перевести в пассивный режим (сновидение)
	psychic.ReadiStatus = 0
	sleepСonditionPulsCount = SlipPulsCount
}

////////////////////////////////

/* разбудить - сторожевая функция по важным объектам
 */
func WakeUpping() {
	if !IsSleeping {
		return
	}
	sleepСonditionPulsCount = 0
	psychic.ReadiStatus = 1
	//	psychic.WakeUpping() // в understanding.go
	IsSleeping = false
}

////////////////////////////////////

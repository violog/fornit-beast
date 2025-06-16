package brain

import (
	"BOT/brain/reflexes"
	word_sensor "BOT/brain/words_sensor"
	"BOT/lib"
	_ "fmt"
	"strconv"
	"time"
)

/* Пульс - сердце Beast :)
запускает  каждую секунду Аctivatin() в net.go, поддерживая возможность самостоятельных действий Beast
В общем, это обеспечивает возможность параллельной работы функций Beast, отводя им выполнение с каждым пульсом
*/

var PulsCount = 0     // счетчик пульса
var noWorkung = false // не активировать очередной цикл если true - тишина
var startWait = 0     // начало ожидания
var secWaiting = 1    // время ожидания
var LifeTime int      // время жизни в числе пульсов
var EvolushnStage int // стадия развития

func init() {
	// Puls() Запускается только в одном месте (main.go) чтобы не было двух потоков пульса!!!
}

// сохранить время жизни в файл
func saveLifeTime() {
	if LifeTime > 10 { // иногда life_time.txt обнулялся...
		lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_reflex/life_time.txt", strconv.Itoa(LifeTime))
	}
}

/*
запуск пульса по сигналам из getParams := r.FormValue("get_params")
*/
var blockinfP = 0

func SincroTic() {
	blockinfP = 1
	Puls()
}

////////////////////////////////

// ВОДИТЕЛЬ РИТМА ПУЛЬСА
var IsPultActivnost = false // начало активности с Пульта  brain.IsPultActivnost=true brain.IsPultActivnost=false
func Puls() {
	if blockinfP > 0 { // не запускать генеретор пульса в этот раз - синхронизация из Пульта
		pulsActions() // действия для этого такта пульса
		blockinfP = 0
		return
	}
	pulsActions()

	sleepPuls()
	/*
		time.AfterFunc(1 * time.Second, func() {
			// действия дял этого такта пульса
			pulsActions()
			Puls()
		})*/
}
func sleepPuls() {
	pChannel := make(chan int, 1)
	go func() {
		pChannel <- 1
		timer := time.NewTimer(1 * time.Second)
		select {
		case <-timer.C:
			Puls()
		}
		close(pChannel) // закрываем канал
	}()
}

///////////////////////////////////////////

/*
	действия, совершаемые по каждому пульсу:

буквально вся работа восприяти - действия идет последовательно,
так что pulsActions() ждет пока все не выполниться.
*/
var isBusyFromWork = false // если pulsActions() не успело выполниться, то пропускает вызов текущего пульса
func pulsActions() {

	if isBusyFromWork {
		return
	}
	isBusyFromWork = true // если pulsActions() не успело выполниться, то пропускает текущий цикл
	// сканирование состояния (пульс):
	//now := time.Now()
	//curTime := now.Second()//.Millisecond()
	//log.Println("пульс",curTime)
	if PulsCount == 2 {
		lib.WritePultConsol("Beast активируется.")
	}
	/* ЗАЧЕМ ОСТАНАВЛИВАТЬ АКТИВНОСТЬ ПРИ СОБЫТИИ С ПУЛЬТА??
		 NotAllowAnyActions=true должно использоваться ТОЛЬКО для остановки активности дял ЗАПИСИ ПАМЯТИ!
	if  IsPultActivnost{// остановка любой активности Beast
		NotAllowAnyActions=true
	}else{
		NotAllowAnyActions=false
	}
	*/
	if !noWorkung {
		Activation(LifeTime) // Главный цикл (пульс) опроса состояния Beast и его ответов в brain.go
	}
	if startWait > 0 && !IsPultActivnost { // ждем следующего цикла и выполняем inaction
		if PulsCount > (startWait + 1) {
			inaction()
			noWorkung = false // конец ожидания
		}
	}
	LifeTime++
	//if (PulsCount + 1) % 10 == 0 {
	//	saveLifeTime()
	//}
	//ac := time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC)
	//lastPulsTime=curTime
	PulsCount++

	isBusyFromWork = false // если pulsActions() не успело выполниться, то пропускает текущий цикл

	// не было активности с пульта и нет перегруза (раз дошло до сюда)
	if !reflexes.IsPultActionThisPuls && word_sensor.NeedCheckTempList {
		word_sensor.NeedCheckTempList = false // только 1 раз
		word_sensor.UpdateWordTreeFromTempArr(3, 5)
	}
}

///////////////////////////////////////////////////

var NotAllowAnyActions = false // - запрет любой активности  brain.NotAllowAnyActions
// то, что должно выполнять в тишине, в бездействии
func inaction() {
	/* Память сохраняется при корректном выключении, попытки автоматизировать - очень геморные и ненадежные
	// сохранять все каждые 100 сек если нет действий
	if (PulsCount+1)%100==0{
		NotAllowAnyActions=true
		SaveAll()
		NotAllowAnyActions=false
	}
	*/
}

/*
	если нужно остановить все на время waitSec для выполнения какой-то функции, то

Пример для функции созранения всей памяти:

	func WaitSaveAllPryMemory(){
		OwnFuncForRun=SaveAllPryMemory
		StopAll() // выполнить  saveAllPryMemory() в тишине, когда ничто не работает
	}
*/
func StopAll() {
	noWorkung = true // останавливает puls()
	startWait = PulsCount
}

func StopRunAll(stop bool) {
	if stop { // останавливает puls()
		noWorkung = true
	} else {
		noWorkung = false
	}
}

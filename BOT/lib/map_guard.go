/*
Способ защиты карт от ПРОБЛЕМЫ С КАРТАМИ "concurrent map writes"
ПРИМЕР реализации предотвращения конфликтов карты
_________________________________________________________
Штатный способ через мьютексы - для любых ситуаций:

import "sync"

var mapArr = make(map[int]int)
var MutexMapArr = &sync.RWMutex{}

// при записи:
MutexMapArr.Lock()
mapArr[key] = value
MutexMapArr.Unlock()

// при чтении или перед перебором for:
MutexMapArr.RLock()
value,ok := mapArr[key]
if ok{ ... }
MutexMapArr.RUnlock()
Но концепция мьютексом очень капризна и бывают ситуации гонок без какого-либо оповещения,
просто функции перестают работать из-за гонок: дедлоком (взаимной блокировкой) между RLock() и Lock().
Примекр такой ситуации описан в комменте в самом низу пакета.
Хотя есть встроенный инструмент для обнаружения состояний гонки, называемый "Race Detector" он не спасает положения.

Из-за такой непродуманности для нашего случая использования только одной главной горутины и запросов GET и POST,
которые и приводят к конкурентности, лучше использовать следующий способ.
Его доддержка остается в этом пакете всегда активной.

Самодельный способ, применимый в случае отсуствия конкурирующих горутин (одновременность только за счет POST и GET обработки)
// при записи - ждем когда отпустит и снова блокируем
lib.MapCheckBlock(mapGwardYyyArr)
yyyArr[1]=i
lib.MapFree(mapGwardYyyArr)

// при одиночном чтении - просто ждем когда отпустит
lib.MapCheck(mapGwardYyyArr)
_,ok:=yyyArr[1]

// при чтении или перед перебором - ждем когда отпустит и снова блокируем
lib.MapCheckBlock(mapGwardYyyArr)

	for k, _ := range yyyArr {
		if k > 0 {	}
	}

lib.MapFree(mapGwardYyyArr)
______________________________________________________
*/
package lib

import (
	"math/rand"
	"time"
)

//////////////////////////////////////////////////

/*
// ПРОВЕРКА РАБОТЫ ВАРИАНТА С МЬЮТЕКСАМИ:
// тестовый вызов из brain/init.go RunInitialisation(): lib.MapExperiment()
func MapExperiment(){
	fmt.Println("ПРОВЕРКА")
	yyyArr[1]=333
	go www()
	go rrr()

	time.Sleep(3 * time.Second)

	return
}
//////////////////////////////////////
func rrr(){
	mutexYyyArr.Lock()
	for i := 0; i < 10000; i++ {
		for k, _ := range yyyArr {
			if k > 0 {
			}
		}
	}
	mutexYyyArr.Unlock()

	for i := 0; i < 10000; i++ {
		_,ok:=readYyyArr(1)
		if 	ok{

		}

		//readYyyArr(1)

	}
	return
}
func www(){

	for i := 0; i < 10000; i++ {
		writeYyyArr(1,i)
	}
}


/////////////////////////////////////
var yyyArr=make(map[int]int)
var mutexYyyArr = &sync.RWMutex{}
func writeYyyArr(key, value int) {
	mutexYyyArr.Lock()
	yyyArr[key] = value
	mutexYyyArr.Unlock()
}

func readYyyArr(key int) (int,bool) {
	mutexYyyArr.RLock()
	value,ok := yyyArr[key]
	mutexYyyArr.RUnlock()
	return value,ok
}
*/

/*
//////////////////////////////////////////////////////////
// ПРОВЕРКА РАБОТЫ ВАРИАНТА С выжидательной блокировкой:
// тестовый вызов из func init()

var yyyArr=make(map[int]int)
var mapGwardYyyArr=RegNewMapGuard()

func MapExperiment(){
	fmt.Println("ПРОВЕРКА")
	yyyArr[1]=333
	go www()
	go rrr()

	time.Sleep(10 * time.Second)// подождем когда все горутины пройдут

	return
}
//////////////////////////////////////
func rrr(){
	MapCheckBlock(mapGwardYyyArr)
	for i := 0; i < 10000; i++ {
		for k, _ := range yyyArr {
			if k > 0 {
			}
		}
	}
	MapFree(mapGwardYyyArr)

	for i := 0; i < 10000; i++ {
		MapCheck(mapGwardYyyArr)
		_,ok:=yyyArr[1]
		if 	ok{

		}
	}
	return
}
func www(){
	MapCheckWrite(mapGwardYyyArr)
	for i := 0; i < 10000; i++ {
		yyyArr[1]=i
	}
	MapFree(mapGwardYyyArr)
}
*/

/////////////////// ПОДДЕРЖКА  MapFlag["map1Arr"]=
/* массив флагов блокировки для всех карт
Если MapFlag[nn]==true, то не может быть записи или считывания.
Это - коунтер блокировок.

var mapGwardActionsImageArr=lib.RegNewMapGuard()

MapCheckWrite(mapGwardActionsImageArr) - блокировка при записи
lib.MapCheckBlock(mapGwardActionsImageArr) - блокировка при чтении больших фрагментов
lib.MapFree(mapGwardActionsImageArr) - отпускаем на один шажок

lib.MapCheck(mapGwardActionsImageArr)  - просто ждем когда отпустит

lib.MapCheck(mapGwardActionsImageArr)
val,ok:=yyyArr[nn]
if ! ok{
return nil
}
///////////////////////////////////////////////////
*/

var MapFlag []int

var mapGuardCounter = 0

func RegNewMapGuard() int {
	if mapGuardCounter == 0 {
		MapFlag = append(MapFlag, 0)
	}
	mapGuardCounter++
	MapFlag = append(MapFlag, 0)
	MapFlag[mapGuardCounter] = 0
	return mapGuardCounter
}

func init() {
	rand.Seed(time.Now().UnixNano())

	// MapExperiment() // Для тестирования нужно раскомментировать "ПРОВЕРКА РАБОТЫ ВАРИАНТА С выжидательной блокировкой"
}

///////////////////////////////////

func MapCheck(index int) {
	/* В очень сложных, вложенных конструкциях возникают дублирования блокировок и код подвешивается.
	С помощью TodoPanic("В коде есть место, вызывающее постоянную блокировку!")
	можно находить такие места и разруливать. Но это не всегда возможно,
	поэтому вполне приемлемо будет просто сбрасывать такой цикл: MapFlag[index]=0;break

	if n>1000{ должно быть больше, чем время на запись в карту между MapCheckWrite() и MapFree()!!!
	*/
	n := 0
	for MapFlag[index] > 0 {
		if n > 1000 {
			// TodoPanic("В коде есть место, вызывающее постоянную блокировку!")
			MapFlag[index] = 0
			break // освобождение после каких-то лаж в коде...
		}
		n++
		time.Sleep(10 * time.Microsecond)
	}
}

// блокировка при записи с удалением одновременности
func MapCheckWrite(index int) {
	/* нарушить одновременность вызова MapCheck с данным индексом
		Вероятность одновременного вызова, превышающего скорость отработки флага, ничтожна,
		но механизмы провоцируют строгую одновременность. Поэтому такую одновременность нарушим,
		выжидая случайное число микросекунд (от 0 до 100).
		Такое время при записи является приемлемым для задач проекта Beast.
	После этого какой-то из вызовов MapCheckBlock(index) окажется первым и закроет последующий MapCheck(index).
	*/
	MapCheck(index)
	if true { // false - чтобы отключить лекарство от одновременности вызова MapCheckWrite
		randNum := rand.Intn(10000)
		//time.Sleep(time.Duration(randNum) * time.Nanosecond) очень нестабильна, зависит от текущей нагрузки
		// задержка - просто выполнением числа операций
		for randNum > 0 {
			randNum--
		}
		MapCheck(index)
	}
	MapFlag[index]++
}

// блокировка при переборе типа чтении for k, _ := range yyyArr или больших фрагментов кода
func MapCheckBlock(index int) {
	MapCheck(index)
	MapFlag[index]++
}

func MapFree(index int) {
	MapFlag[index]--
	if MapFlag[index] < 0 {
		MapFlag[index] = 0
	}
}

///////////////////////////////////////////////////////////////
/*  ИСПОЛЬЗОВАНИЕ:
// при записи - ждем когда отпустит и снова блокируем
lib.MapCheckBlock(mapGwardYyyArr)
yyyArr[1]=i
lib.MapFree(mapGwardYyyArr)

// при одиночном чтении - просто ждем когда отпустит
lib.MapCheck(mapGwardYyyArr)
_,ok:=yyyArr[1]

// при чтении или перед перебором - ждем когда отпустит и снова блокируем
lib.MapCheckBlock(mapGwardYyyArr)
for k, _ := range yyyArr {
	if k > 0 {	}
}
lib.MapFree(mapGwardYyyArr)
*/
//////////////////////////////////////

/* Как можно выловить лажу мьютекса:
в строке MutexTempArr.Lock() правок кнопкой по Lock выйти на ее реализацию в пакете mutext.go
там поставить точку прерывания на runtime_SemacquireMutex:
	if r != 0 && atomic.AddInt32(&rw.readerWait, r) != 0 {
		runtime_SemacquireMutex(&rw.writerSem, false, 0)
	}
и тестировать разные случаи пока не будет останов в точке, после чего прояти по стеку и найти откула пошел конфликт.

Пример сбоя в ГО при использовании мьютексов для карты cyclesArr:
func endDereamsCycles(){
	MutexCyclesArr.RLock()
	for id, v := range cyclesArr {
		if v.dreaming>0 {
			endBaseIdCycle(id)
		}
	}
	MutexCyclesArr.RUnlock()
}
но в функции endBaseIdCycle(id) уже используется мьютекс:
func endBaseIdCycle(baseID int){
	foundMainCycle(baseID)//найти новый главный цикл по значимости, кроме удаляемого
	MutexCyclesArr.Lock()
	delete(cyclesArr, baseID)
	MutexCyclesArr.Unlock()
}
Возникает гонка, останавливающая выполнение функции в
func (rw *RWMutex) Lock() {
...
if r != 0 && atomic.AddInt32(&rw.readerWait, r) != 0 {
		runtime_SemacquireMutex(&rw.writerSem, false, 0)
	}
}
Как правильно использовать мьютексты в таких случаях?


В данном случае проблема связана с дедлоком (взаимной блокировкой) между RLock() и Lock(). Вы пытаетесь получить мьютекс на запись в то время как у вас уже есть чтения блокировка этого же мьютекса. Итог - взаимная блокировка.

Метод RLock() не предназначен для "повышения" на блокировку записи. Если вы захватили блокировку чтения с помощью RLock(), вы не сможете получить блокировку записи с помощью Lock(), пока не освободите блокировку чтения с помощью RUnlock().

Вам нужно будет пересмотреть свою стратегию использования мьютекса. Вы можете:

Использовать блокировку записи Lock() в endDereamsCycles(), а не блокировку чтения. Это значит, что во время выполнения endDereamsCycles() другие горутины не смогут читать или писать в cyclesArr.

Разделить функцию endDereamsCycles() на две части: одна для чтения и фильтрации значение из cyclesArr, и другая для вызова endBaseIdCycle(). Это позволит вам освободить блокировку чтения перед вызовом endBaseIdCycle().

Сделать копию данных (cyclesArr), которые будут обрабатываться, защитив их с помощью блокировки чтения, а затем производить изменения вне блокированного блока.

Пожалуйста, имейте в виду, что правильное решение зависит от контекста вашего приложения и того, какие гарантии вы хотите предоставить.

Возможно, лучшим вариантом для данного случая будет подход, при котором вы создаете копию keys, которые нужно обработать. Затем обрабатываете каждый key / value в отдельности.

Код ниже показывает пример такого подхода:


import "sync"

var cyclesArr= make(map[int]*MyStruct) // предположим, что MyStruct - ваш произвольный тип
var MutexCyclesArr = &sync.RWMutex{}

func endDereamsCycles(){
    // захватываем блокировку чтения
    MutexCyclesArr.RLock()

    // создаем копию ключей
    keys := make([]int, 0, len(cyclesArr))
    for id, v := range cyclesArr {
        if v.dreaming > 0 {
            keys = append(keys, id)
        }
    }
    // освобождаем блокировку чтения
    MutexCyclesArr.RUnlock()

    // теперь обрабатываем наши элементы
    for _, id := range keys {
        endBaseIdCycle(id)
    }
}

// найти новый главный цикл по значимости, кроме удаляемого - предположительные операции
func foundMainCycle(baseID int){
  //добавьте здесь ваши операции
}

func endBaseIdCycle(baseID int) {
    foundMainCycle(baseID)
    MutexCyclesArr.Lock()
    delete(cyclesArr, baseID)
    MutexCyclesArr.Unlock()
}
В этом коде, на первом этапе функции endDereamsCycles() мы захватываем блокировку только для чтения и копирования ключей нуждающихся в обработке элементов.

Затем, после того как все ключи были скопированы, мы освобождаем блокировку чтения и переходим к обработке каждого ключа в отдельности с захватом блокировки записи в endBaseIdCycle().


Go имеет встроенный инструмент для обнаружения состояний гонки, называемый "Race Detector". Этот инструмент может автоматически обнаруживать состояния гонки во время выполнения и обычно очень полезен для обнаружения сложных ошибок конкурентной среды.

Использование детектора состояний гонок ("Race Detector") в Go легко. Вам просто нужно добавить флаг -race при запуске вашей программы или тестов.

Для запуска все своей программы с детектором состояний гонок, вы можете использовать команду go run -race myprogram.go.

Для запуска тестов с использованием детектора состояний гонок, вы можете использовать команду go test -race mypackage.

В обоих случаях детектор состояний гонок автоматически обнаруживает и сообщает о любых встреченных условиях гонки.

Однако, важно отметить, что детектор состояний гонок не может найти потенциальные состояния гонок, которые могут произойти в будущем. Он обнаруживает только те состояния гонок, которые фактически произошли во время выполнения.

Также стоит отметить, что, несмотря на то что "Race Detector" - это полезный  инструмент, его использование увеличивает время выполнения программы и требует больше памяти, поэтому его обычно не используют в продакшен-среде.

В IDE GoLand, вы можете включить "Race Detector" несколькими способами:

Настройка среды выполнения для конкретной задачи:

Откройте настройки "Run/Debug Configurations".
Выберите конкретное задание на выполнение (или создайте новое).
В разделе "Go build flags" напишите -race.
Сохраните изменения и запустите задание на выполнение.
Непосредственно перед запуском кода:

Откройте терминал GoLand.
Используйте go run -race или go test -race followed by your .go file or package.
В обоих случаях GoLand запустит ваш код с "Race detector", и если будут обнаружены состояния гонки, они будут отображены в консоли вывода.

*/

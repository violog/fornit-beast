/* Передача инфы на Пульт */

package lib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// строка вывода на пульт - для func WritePultConsol
var WritePultConsolStr = ""

/*
	вывести на консоль Пульта

Сообщения накапливаются в WritePultConsolStr и откправлются с каждым пульсом
*/
func WritePultConsol(print string) {
	WritePultConsolStr = print + "<br>" + WritePultConsolStr
	// очищать от форматирования тегами, проверка - в func RunInitialisation()
	reg := regexp.MustCompile(`<\/?[^>]+>`)
	print = reg.ReplaceAllString(print, " ")
	fmt.Println("НА ПУЛЬТ: ", print)
}

// функция вызова паники с информированием в логе Пульта
func TodoPanic(panicWarning string) {
	WritePultConsol("<span style='color:red;font-size:19px;font-weight:bold;'>ПАНИКА: </span> " + panicWarning)
	panic(panicWarning)
}

//////////////////////////////////////////////////

// В main.go действия БЛОКИРУЮТСЯ В РЕЖИМЕ СНА: !sleep.IsSleeping
var ActionsForPultStr = ""     // Строка действий для Пульта
var BlockingAnyActions = false // блокировать выдачу на пульт любых действий во время сна или dreaming
/*
	вывести на Пульт действия Бота строкой lib.SentActionsForPult("xcvxvxcv")

Каждая акция - в формате: вид действия (1 - действие рефлекса, 2 - фраза) затем строка акции,
например: "1|Предлогает поиграть" или "2|Привет!"
Можно передавать неограниченную последовательность акций, разделяя их "||"
например: "1|Предлогает поиграть||2|Привет!"

За один пульс может набраться очередь нескольких действий, например,
после действия может измениться условие и будет новое действие.
Поэтому их все нужно выводить, но не допускать дублирования.

Дубли так же контролируются в doublesRemove - проверкой строки "энергичность",
т.к. могут быть несколько одинаковых сообщений с разной энергичностью.
*/
func SentActionsForPult(act string) {
	if BlockingAnyActions {
		ActionsForPultStr = ""
		return
	}
	if len(ActionsForPultStr) > 0 { // еще не прочитана предыдущая инфа т.к. читается раз в пульс, а после действия может измениться условие и будет новое действие
		// посмотреть дубли в очереди
		sArr := strings.Split(ActionsForPultStr, "||")
		for _, v := range sArr {
			if act == v || len(v) < 4 { // уже есть такое действие, не добавлять в очередь
				return
			}
		}
		ActionsForPultStr += "||" + act // добавить новую к предыдущему
		return
	}
	ActionsForPultStr = act
}

/*
Дубли контролируются проверкой строки "энергичность",
т.к. могут быть несколько одинаковых сообщений с разной энергичностью.
Для одинаковых действий ставится максимальная Энергичность из имеющихся.

Вызывается в main.go перед отправкой на Пульт.
*/
func DoublesRemove(str string) string {
	if str == "" {
		return ""
	}

	out := ""
	var s1Arr = make(map[int]string)
	var efArr = make(map[int]int)
	var s2Arr = make(map[int]string)
	sArr := strings.Split(str, "||")
	// собираем инфу
	for i := 0; i < len(sArr); i++ {
		s1, ef, s2 := cutEnerg(sArr[i])
		s1Arr[i] = s1
		efArr[i] = ef
		s2Arr[i] = s2
	}
	// ищем дубли
	var dArr []int
	for i := 0; i < len(sArr); i++ {
		if ExistsValInArr(dArr, i) {
			continue
		} // дубли не смотреть
		var max = efArr[i]
		out += s1Arr[i] + " Энергичность:"
		for j := 0; j < len(sArr); j++ {
			if i == j {
				continue
			} // саму себя не сравнивать
			if s1Arr[i] == s1Arr[j] {
				if efArr[j] > max {
					max = efArr[j]
				}
				dArr = append(dArr, j) // не смотреть это
				continue
			}
		}

		out += strconv.Itoa(max) + "</b>" + s2Arr[i]
		out += "||"
	}

	return out
}

// рубим сообщение на части
func cutEnerg(str string) (string, int, string) {
	if str == "" {
		return "", 0, ""
	}
	s1 := strings.Split(str, "Энергичность:")
	out1 := s1[0]
	if len(s1) == 1 {
		return out1, 0, ""
	}
	s2 := strings.Split(s1[1], ")</b>")
	out2, _ := strconv.Atoi(strings.Split(s2[0], "=")[1])
	out3 := s2[1]
	return out1, out2, out3
}

//////////////////////////////////////////////////

/*
	Показать непонимания, растерянность -

в случае отсуствия пси-реакций но не Лени.
lib.SentConfusion()
*/
func SentConfusion(detaile string) {
	ActionsForPultStr = "10|" + detaile
}

/*
можно ли показывать акции рефлексов с автоматизмами в одной плашке действия Beast на пульте
ТЕПЕРЬ: НИКОГДА НЕЛЬЗЯ показывать акции рефлексов с автоматизмами в одной плашке!
РАНБШЕ БЫЛО:
Eсли noReflexWithAutomatizm==true и если в ActionsForPultStr есть автоматизм, то строки рефлексов удаляются.
При запуске автоматизма из цикла мышления (на 3-м уровне осмысления) устанавливается NoReflexWithAutomatizm =true
Применение:
ActionsForPultStr=SharedRflexWithAutomatizm(true)
*/

var NoReflexWithAutomatizm = false

func SharedReflexWithAutomatizm() string {
	if !NoReflexWithAutomatizm {
		//		return ActionsForPultStr  НИКОГДА НЕЛЬЗЯ показывать акции рефлексов с автоматизмами в одной плашке!
	}
	out := ""
	sArr := strings.Split(ActionsForPultStr, "||")
	n := 0
	for i := 0; i < len(sArr); i++ {
		id, _ := strconv.Atoi(sArr[i][:1]) // первый символ - идентификатор типа акции
		if id > 1 || sArr[i][:2] == "10" { // нужно еще учесть активацию infoFunc(13)
			if n > 0 {
				out += "||"
			}
			out += sArr[i]
			n++
		}
	}
	return out
}

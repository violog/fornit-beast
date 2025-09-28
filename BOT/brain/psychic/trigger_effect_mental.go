/*
	Образ ментального Правила.

Ментальное Правило используется как опыт использования ментальных автоматизмов для нахождения решения.

Начало фиксации Правила - ментальный запуск моторного автомтаизма (MentalAutomatizm.ActionsImageID ->  activateMotorID).
Эффект Правила отражает насколько мент.автоматизм привел к успешному решению,
т.е. оценивается в момент действий оператора в период ожидания.
После этого в эпизод память и MentalCyckleEffectArr записывается новое Правило.
Так же при этом проставляется MentalAutomatizm.Usefulness автоматизму, запустившему действие (вообще MentalAutomatizm нивелируется Правилами).

При ментальном запуске моторного автомтаизма фиксируется фрагмент Кратковременной памяти (infoFuncSequence),
начиная с последней объективной активации consciousnessElementary т.е. только цепочка для данной активности дерева актоматизмов,
но могут быть разные активности дерева понимания из-за произвольной активации
(т.к. при переактивации деревьев могут измениться базовые циклы) и эти звенья оцениваются как успешнвые или нет.
Это Правило записывается в массив Правил rulesMentalArr, и в кадр эпиз.памяти с Type=1 .
*/
package psychic

import (
	"BOT/lib"
	"strconv"
	"strings"
)

/////////////////////////
/* В отличие от объективнх Правил здесь нет пошаговой сверки с реакцией Оператора (это внутренний цикл поиска решения),
так что нужно просто смотреть наиболее подходящее Правило.
Каждое Правило - один из усвоенных алгоритмов поиска решений - циклом мышления.
Удачный цикл мышления записывается в мент.автоматизм, привязанный к ветке древа проблем.
Так что в ветке проблем фиксируется удачное ментальное Правило == цикл мышления.

funcID []int последовательность ID выполненных инфо-функций до срабатывания автоматизма по func infoFunc17,
т.е. была такая последовательность пока не сработал мот.автоматизм,ч то привело к эффекту Effect.
*/
type MentalCyckleEffect struct {
	ID     int
	funcID []int // infoFuncSequence[]int данной активации - последовательность ID выполненных инфо-функций
	Effect int   // эффект от действий накапливается при каждой новой перезаписи и используется уже суммарное значение.
}

// //////////////////////
var MentalCyckleEffectArr = make(map[int]*MentalCyckleEffect)
var MapGwardMentalCyckleEffectArr = lib.RegNewMapGuard()

//////////////////////////////////////////

// //////////////////////////////////////////////
// создать новый эффект ментального правила если такого еще нет
var lastMentalCyckleEffectID = 0

func createNewlastMentalCyckleEffectID(id int, ShortTermMemoryID []int, Effect int, CheckUnicum bool) (int, *MentalCyckleEffect) {
	if Effect < 0 {
		Effect = -1
	}
	if Effect > 0 {
		Effect = 1
	}

	if CheckUnicum {
		oldID, oldVal := checkUnicumMentalCyckleEffect(ShortTermMemoryID, Effect)
		if oldVal != nil {
			return oldID, oldVal
		}
	}

	if id == 0 {
		lastMentalCyckleEffectID++
		id = lastMentalCyckleEffectID
	} else {
		//		newW.ID=id
		if lastMentalCyckleEffectID < id {
			lastMentalCyckleEffectID = id
		}
	}

	var node MentalCyckleEffect
	node.ID = id
	node.funcID = infoFuncSequence
	node.Effect = Effect

	lib.MapCheckWrite(MapGwardMentalCyckleEffectArr)
	MentalCyckleEffectArr[id] = &node
	lib.MapFree(MapGwardMentalCyckleEffectArr)

	if doWritingFile {
		SaveMentalCyckleEffectArr()
	}

	return id, &node
}
func checkUnicumMentalCyckleEffect(ShortTermMemoryID []int, Effect int) (int, *MentalCyckleEffect) {
	lib.MapCheckBlock(MapGwardMentalCyckleEffectArr)
	for id, v := range MentalCyckleEffectArr {
		if v == nil {
			continue
		}
		if !lib.EqualArrs(ShortTermMemoryID, v.funcID) {
			continue
		}
		if Effect != v.Effect {
			continue
		}
		lib.MapFree(MapGwardMentalCyckleEffectArr)
		return id, v
	}
	lib.MapFree(MapGwardMentalCyckleEffectArr)
	return 0, nil
}

/////////////////////////////////////////

// ////////////////// сохранить Образы стимула (действий оператора) - ответа Beast
// В случае отсуствия ответных действий создается ID такого отсутсвия, пример такой записи: 2|||0|0|
// ID|ActID через ,|PhraseID через ,|ToneID|MoodID|
func SaveMentalCyckleEffectArr() {
	var out = ""
	lib.MapCheckBlock(MapGwardMentalCyckleEffectArr)
	for k, v := range MentalCyckleEffectArr {
		if v == nil {
			continue
		}
		out += strconv.Itoa(k) + "|"
		for i := 0; i < len(v.funcID); i++ {
			if i > 0 {
				out += ","
			}
			out += strconv.Itoa(v.funcID[i])
		}
		out += "|"
		out += strconv.Itoa(v.Effect)
		out += "\r\n"
	}
	lib.MapFree(MapGwardMentalCyckleEffectArr)
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/trigger_and_actions_mental.txt", out)

}

// //////////////////  загрузить образы стимула (действий оператора) - ответа Beast
// ID|ActID через ,|PhraseID через ,|ToneID|MoodID|
func loadMentalCyckleEffectArr() {
	MentalCyckleEffectArr = make(map[int]*MentalCyckleEffect)
	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/trigger_and_actions_mental.txt")
	cunt := len(strArr)
	for n := 0; n < cunt; n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		p := strings.Split(strArr[n], "|")
		id, _ := strconv.Atoi(p[0])

		s := strings.Split(p[1], ",")
		var ShortTermMemoryID []int
		for i := 0; i < len(s); i++ {
			si, _ := strconv.Atoi(s[i])
			ShortTermMemoryID = append(ShortTermMemoryID, si)
		}
		Effect, _ := strconv.Atoi(p[2])
		var saveDoWritingFile = doWritingFile
		doWritingFile = false
		createNewlastMentalCyckleEffectID(id, ShortTermMemoryID, Effect, false)
		doWritingFile = saveDoWritingFile
	}
	return

}

///////////////////////////////////////////

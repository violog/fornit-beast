/* Короткие цепочки вспомогательной ментальной памяти.
Два вида короткой памяти:
1. Стек прерывания размышления
2. Стек обобщения смыслов воспринятого
Величина этих стеков должна оптимизироваться к текущим условиям жизни.
Здесь же она задается жестко - по 10 ячеек памяти в каждом стеке.
______________________________________
и стек памяти прерванных размышлений  var InterruptMemory []*InterruptImage
где созраняется прерванный ID итерации цикла, ID темы и ID поставленной цели,
Так что при возврате к прерванному размышлению, восстанавлявается его образы.
Этот стек сохраняется в файле чтобы после перезагрузки можно было начать с прерванного.
_____________________________________
Стек памяти обобщения смыслов - буфер для сбора отдельных смыслов
при восприятии флинной фразы, которую нужно разбить на отдельные смыслы (importance)
Этот стек НЕ сохраняется в файле.
При восприятии фразы с пульта нужно смотреть, распознается ли смысл (importance) по мере прохождения, как только
удается понять текущий смысл фрагмента, если да, то фрагмент кладется в стек обобщения и туда добавляются все смыслы,
чтобы затем понять общий. Ребенок узнает слово по буквам, вставляя в стек отдельные буквы.

Пока что остается за скобками то, что использует стек обобщения:
ТВОРЧЕСКОЕ СОПОСТАВЛЕНИЕ И ОБОБЩЕНИЕ
ТВОРЧЕСКАЯ ФАНТАЗИЯ
*/

package psychic

import (
	"BOT/lib"
	"strconv"
	"strings"
)

/*
	стек для обобщений: 7 Базовых cycleID

Сохраняет 7 Базовых cycleID (звенья начала цепочки)
*/
var saveBaseLinksCicleSynthesis []*int

func addMewBaseLinksMemory(BaseLinksID int) {
	if len(saveBaseLinksCicleSynthesis) > 7 {
		// удалить первый
		saveBaseLinksCicleSynthesis = saveBaseLinksCicleSynthesis[1:]
	}

	saveBaseLinksCicleSynthesis = append(saveBaseLinksCicleSynthesis, &BaseLinksID)
}

//////////////////////////////////////////////////////////////////////////

/////////////////////   стек памяти прерванных размышлений
/* стек до 7 прерванных задач
Запоминаются в файле и актуализируются при просыпании

Здесь нет базовых условий, и поэтому при их изменении образ прерванной цели обессмысливается,
чтобы вспомнить нужно опять попасть в ситуацию с прерыванием.
Можно улучшить природу и сделать полное восстановление условий ментальным перезапуском,
но это может привести к неожиданным последствиям.
*/
// образ прерванного размышления
type InterruptImage struct {
	cycleID        int // на каком ID goNext было прерванно размышление
	ThemeImageType int //Type темы размышления
	PurposeImageID int //ID поставленной цели
	ExtremObjID    int // extremImportanceObject.extremObjID
}

var InterruptMemory []*InterruptImage

// добавить в стек прерывание
func addInterruptMemory(cycleID int) {
	// Добавить в стек до 7 прерванных решений
	if len(InterruptMemory) > 7 {
		// удалить первый
		InterruptMemory = InterruptMemory[1:]
	}
	var node InterruptImage
	node.cycleID = cycleID
	node.ThemeImageType = mentalInfoStruct.ThemeImageType
	node.PurposeImageID = mentalInfoStruct.mentalPurposeID
	if extremImportanceObject != nil {
		node.ExtremObjID = extremImportanceObject.objID
	}
	InterruptMemory = append(InterruptMemory, &node)
}

////////////////////////////////////////////////////////////////

func loadInterruptMemory() {
	InterruptMemory = nil
	strArr, _ := lib.ReadLines(lib.GetMainPathExeFile() + "/memory_psy/interrupt_memory.txt")
	cunt := len(strArr)
	for n := 0; n < cunt; n++ {
		if len(strArr[n]) == 0 {
			continue
		}
		p := strings.Split(strArr[n], "|")
		var node InterruptImage
		node.cycleID, _ = strconv.Atoi(p[0])
		node.ThemeImageType, _ = strconv.Atoi(p[1])
		node.PurposeImageID, _ = strconv.Atoi(p[2])
		InterruptMemory = append(InterruptMemory, &node)
	}
}
func saveInterruptMemory() {
	var out = ""
	for n := 0; n < len(InterruptMemory); n++ {
		out += strconv.Itoa(InterruptMemory[n].cycleID) + "|"
		out += strconv.Itoa(InterruptMemory[n].ThemeImageType) + "|"
		out += strconv.Itoa(InterruptMemory[n].PurposeImageID)
	}
	out += "\r\n"
	lib.WriteFileContent(lib.GetMainPathExeFile()+"/memory_psy/interrupt_memory.txt", out)
}

//////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////
/////////////////////////  стек обобщения смыслов
/*  Здесь хранятся распознанные  ID смыслов типа type importance struct
для элементов фразы Стимула из образа действия ActionsImage или нескольких подрях ActionsImage
т.е. смысл обобщается не только в одной длинной фразе, а и в ходе диалога.

*/
// образ прерванного размышления
type SynthesizeImportanceImage struct {
	PhraseID     []int // // массив фразID (DetectedUnicumPhraseID) слова каждой фразы вытаскиваются wordSensor.GetWordArrFromPhraseID(PhraseID[0])
	importanceID int   //ID importance (var importanceFromID=make(map[int]*importance))
}

var SynthesizeImportanceMemory []*SynthesizeImportanceImage

/* Тут должны быть функци работы с обобщением нескольких смыслов в один общий.
Это - пока не очень понятный алгоритм
НО В ДАННОЙ РЕАЛИЗАЦИИ Beast нет особого смысл в обобщении смыслов...
*/
// TODO когда-нибудь, но это  не принципиально для общей системы

///////////////////////////////////////////////////////////////////////

/*
	Был запущен штатный автоматизм после Стимула. Выясняется на 1-м и 2-м уровне осмысления.

Если был запущен автоматизм, то блокировать запуски последующих, кроме ментального запуска wasRunPurposeActionFunc.
Если вдруг нужно будет, подумав, запустить ментально автоматизм, то нужно wasRunTreeStandardAutomatizm=false.
wasRunTreeStandardAutomatizm=true устанавливается при любом запуске автоматизма в func RumAutomatizm

var wasRunTreeStandardAutomatizm = false  ТЕПЕРЬ всегда после запуска автоматизма LastRunAutomatizmPulsCount >0
*/

/*
последовательность ID выполненных инфо-функций, ментальный эизод памяти, в одной активации consciousnessElementary()
infoFuncSequence=append(infoFuncSequence,infofID) - в
для ментальных правил func afterWaitingPeriod(effect int) в understanding_process.go
Когда сработал автоматизм и после периода ожидания возник Стимул,
то происходит запись м.Правил func afterWaitingPeriod(effect int) в understanding_process.go
с детектором достижения Цели getMentalEffect(effect0 int)
После этого или после завершения периода ожидания (clinerAutomatizmRunning()) infoFuncSequence очищается.

TODO м.б. не самая последняя функция, а ряд последних дает эффект мент.Правила!...
*/
var infoFuncSequence []int

var wasRunPurposeActionFunc = false //true - сработала func infoFunc17, 14 и 26 - прекратилась запись в infoFuncSequence[]
func clinerFuncSequence() {
	infoFuncSequence = nil
	wasRunPurposeActionFunc = false
	//curFunc13ID = 0
}

var wasRunProvocationFunc = false //true - сработала провокационная функция infoFunc31

//////////////////////////////////////////////////////////////////////////

func initMentalMemories() {
	//savePorposeIDcurrentCicle=nil
	saveBaseLinksCicleSynthesis = nil
	loadInterruptMemory()
}

//////////////////////////////////////////////////////////////////////////////

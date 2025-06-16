/* Формирование зеркальных автоматизмов */

package psychic

import (
	wordSensor "BOT/brain/words_sensor"
	"BOT/lib"
	"strconv"
	"strings"
)

// Формирование зеркальных автоматизмов на основе списка ответов lib/mirror_reflexes_basic_phrases/...
// тестирование - запуск из psychic.go
func FormingMirrorAutomatizmFromList(file string) string {
	path := lib.GetMainPathExeFile()
	strArr, _ := lib.ReadLines(path + file)
	// triggPhrase|baseID|ContID_list|answerPhrase|Ton,Mood|actions1,...
	if len(strArr) < 2 {
		return "Пустой файл"
	}
	/* Эти автоматизмы привязываются к baseID|ContID_list|0|, т.е. к нулевому образу пусковых
	и к нулевому тону-настроению 90.
	Но TODO: сделать более мягкую активацию автоматизмов дерева:
		 ID|ParentNode|BaseID|EmotionID|ActivityID|ToneMoodID|SimbolID|PhraseID
		 Если нет автоматизма для данного узла > ActivityID то смотреть для других узлов, начиная с данного уровня.
	Т.е. если автоматизм привязан к ToneMoodID==90 а активировалась ветка с ToneMoodID==12 где нет автоматизма,
	то пусть бы срабатывал привязанный к ToneMoodID==90 !!!
		 НО ОСТОРОЖНО (понизить силу?)
	*/

	// первую строку пропускаем из-за #utf8 bom
	for n := 1; n < len(strArr); n++ {
		if len(strArr[n]) < 10 {
			continue
		}
		p := strings.Split(strArr[n], "|")
		// УСЛОВИЯ ДЕРЕВА
		// базовое состояние
		baseID, _ := strconv.Atoi(p[1])

		// базовые контексты
		pn := strings.Split(p[2], ",")
		var lev2 []int
		for i := 0; i < len(pn); i++ {
			b, _ := strconv.Atoi(pn[i])
			if b > 0 {
				lev2 = append(lev2, b)
			}
		}

		// образ отсуствия тона и настроения
		tm := 90

		// засунуть фразу в дерево слов и дерево фраз
		prase := p[0]
		wordSensor.VerbalDetection(prase, 0, 0, 0)
		PhraseID := wordSensor.CurrentPhrasesIDarr

		// первый символ ответной фразы
		FirstSimbolID := wordSensor.GetFirstSymbolFromPraseID(PhraseID)
		// создать образ Брока
		_, verbal := CreateVerbalImage(FirstSimbolID, PhraseID, 0, 0)

		nodeID := FindConditionsNode(baseID, lev2, nil, FirstSimbolID, tm, verbal.ID)
		/* если есть привязанный к узлу автоматизм, то он просто перестанет быть штатным,
		т.к. авторитерный (зеркальный) автоматизм важнее
		exists:=ExistsAutomatizmForThisNodeID(nodeID)
		if exists {
			continue
		}	*/
		if nodeID > 0 {
			// создать автоматизм и привязать его к nodeID
			//var sequence = "Snn:" // ответная фраза
			// засунуть фразу в дерево слов и дерево фраз
			wordSensor.VerbalDetection(p[3], 0, 0, 0)
			answerID := wordSensor.CurrentPhrasesIDarr

			//sequence += "|Тnn:" // тон и настроение
			tnArr := strings.Split(p[4], ",")
			t, _ := strconv.Atoi(tnArr[0])
			m, _ := strconv.Atoi(tnArr[1])

			var aArr []int
			aD := strings.Split(p[5], ",")
			for i := 0; i < len(aD); i++ {
				a, _ := strconv.Atoi(aD[i])
				aArr = append(aArr, a)
			}

			NoWarningCreateShow = true

			ActionsImageID, _ := CreateNewlastActionsImageID(0, 0, aArr, answerID, t, m, true)
			_, autmzm := CreateAtutomatizmNoSaveFile(nodeID, ActionsImageID)
			NoWarningCreateShow = false
			if autmzm != nil {
				autmzm.Usefulness = 1          //авторитарный
				SetAutomatizmBelief(autmzm, 2) // сделать автоматизм штатным
				// ?? autmzm.GomeoIdSuccesArr какие ID гомео-параметров улучшает это действие
			}
		}
	}
	SaveAllPsihicMemory()
	return "OK"
}

/*
	на основе общего шаблона ответов lib/mirror_basic_phrases_common.txt

Создаются автоматизмы, привязанные к пусковой фразе, а не к узлу дерева,
с BranchID > 2000000.
var AutomatizmIdFromPhraseId=make(map[int] []*Automatizm)
// тестирование - запуск из psychic.go
*/
func FormingMirrorAutomatizmFromTempList(file string) string {
	path := lib.GetMainPathExeFile()
	strArr, _ := lib.ReadLines(path + file)
	// triggPhrase|baseID|ContID_list|answerPhrase|Ton,Mood|actions1,...
	if len(strArr) == 0 {
		return "Пустой файл"
	}
	/* Эти автоматизмы привязываются к baseID|ContID_list|0|, т.е. к нулевому образу пусковых
	   и к нулевому тону-настроению 90.
	   Но TODO: сделать более мягкую активацию автоматизмов дерева:
	      ID|ParentNode|BaseID|EmotionID|ActivityID|ToneMoodID|SimbolID|PhraseID
	      Если нет автоматизма для данного узла > ActivityID то смотреть для других узлов, начиная с данного уровня.
	   Т.е. если автоматизм привязан к ToneMoodID==90 а активировалась ветка с ToneMoodID==12 где нет автоматизма,
	   то пусть бы срабатывал привязанный к ToneMoodID==90 !!!
	      НО ОСТОРОЖНО (понизить силу?)
	*/

	for n := 0; n < len(strArr); n++ {
		if len(strArr[n]) < 10 {
			continue
		}
		p := strings.Split(strArr[n], "|")
		// УСЛОВИЯ ДЕРЕВА
		// пусковая фраза
		triggerPrase := p[0]

		// ответ
		answerPrase := p[1]

		// тон, настроение
		pt := strings.Split(p[2], ",")
		t, _ := strconv.Atoi(pt[0])
		m, _ := strconv.Atoi(pt[1])
		//tm := GetToneMoodID(t, m + 19)

		// засунуть фразу в дерево слов и дерево фраз
		wordSensor.VerbalDetection(triggerPrase, 0, 0, 0)
		triggerPraseID := wordSensor.CurrentPhrasesIDarr

		wordSensor.VerbalDetection(answerPrase, 0, 0, 0)
		answerPraseID := wordSensor.CurrentPhrasesIDarr

		// создать автоматизм и привязать его к объекту
		NoWarningCreateShow = true
		// для фразы triggerPraseID создаем привязанный к ней автоматизм
		ActionsImageID, _ := CreateNewlastActionsImageID(0, 0, nil, answerPraseID, t, m, true)
		_, autmzm := CreateAtutomatizmNoSaveFile(2000000+triggerPraseID[0], ActionsImageID)
		NoWarningCreateShow = false
		if autmzm != nil {
			autmzm.Usefulness = 1          //авторитарный
			SetAutomatizmBelief(autmzm, 2) // сделать автоматизм штатным
			// ?? autmzm.GomeoIdSuccesArr какие ID гомео-параметров улучшает это действие
		}
	}
	SaveAllPsihicMemory()
	return "OK"
}

/*
	создание зеркального автоматизма, повторяющего действия оператора в данных условиях

в ответ на действия sourceAtmzm - причина ответа оператора
Только что действиями оператора была активирована ветка detectedActiveLastNodID дерева и
есть информация об этих действиях в curActiveActions
Автоматизм прикрепляется к ветке предыдущей активации дерева LastDetectedActiveLastNodID (причине) -
которая становится пусковым стимулом отзеркаливания.
*/
func createNewMirrorAutomatizm(sourceAtmzm *Automatizm) {
	if sourceAtmzm == nil || curActiveActions == nil {
		return
	}
	if curActiveActions.ActID == nil && isUnrecognizedPhraseFromAtmtzmTreeActivation {
		//sourceAtmzm.Usefulness=-1
		sourceAtmzm.Usefulness = 0
		sourceAtmzm.Count = 0
		return
	}
	/* вытащить действия исходного автоматизма чтобы найти или сделать узел дерева с таким пускателем
	   и существубшими BaseID и EmotionID
	*/

	//	curNode:=AutomatizmTreeFromID[detectedActiveLastNodID]
	curNode, ok := ReadeAutomatizmTreeFromID(detectedActiveLastNodID)
	if !ok {
		return
	}
	targetNodeID := findTreeNodeFromAutomatizmActionsImage(curNode.BaseID, curNode.EmotionID, sourceAtmzm)
	if targetNodeID == 0 {
		return
	}
	//	SaveAutomatizmTree()
	// найти узел, который может реагировать на данные действия и если нет - создать его чтобы привязать зеркальный автоматизм

	// создать автоматизм и привязать его к объекту
	// NoWarningCreateShow=true
	ActionsImageID, _ := CreateNewlastActionsImageID(0, 0, curActiveActions.ActID, curActiveActions.PhraseID, curActiveActions.ToneID, curActiveActions.MoodID, true)
	_, autmzm := CreateAtutomatizmNoSaveFile(targetNodeID, ActionsImageID)
	//	NoWarningCreateShow=false
	if autmzm != nil {
		detectedActiveLastNodID = targetNodeID
		// сделать автоматизм штатным, т.к. действия авторитарно верные
		autmzm.Usefulness = 1 //авторитарный
		SetAutomatizmBelief(autmzm, 2)
		autmzm.Count = EvolushnStageAtmzCount(true) // накручиваем счетчик успешных повторов в зависимости от стадии развития
		if doWritingFile {
			SaveAutomatizm()
		}
	}
}

/*
	вытащить действия исходного автоматизма чтобы найти или сделать узел дерева с таким пускателем

и существубшими BaseID и EmotionID
*/
func findTreeNodeFromAutomatizmActionsImage(baseID int, EmotionID int, atmz *Automatizm) int {
	lev2 := EmotionFromIdArr[EmotionID].BaseIDarr

	actImage := atmz.ActionsImageID
	//ai:=ActionsImageArr[actImage]
	ai, ok := ReadeActionsImageArr(actImage)
	if !ok {
		return 0
	}

	// первый символ ответной фразы. В случае ответной цепочки это первый символ первого слова цепочки
	simbolID := wordSensor.GetFirstSymbolFromPraseID(ai.PhraseID)

	/* настроение при передаче фразы с Пульта:
	   	20-Хорошее    21-Плохое    22-Игровое    23-Учитель    24-Агрессивное   25-Защитное    26-Протест
	   ID возникает при добавлении 19 к номеру радиокнопки пульта, например, для Хорошее 1+19=20
	*/
	tm := GetToneMoodID(ai.ToneID, ai.MoodID+19)

	prase := 0
	var actSum []int
	if ai.PhraseID != nil {
		prase = ai.PhraseID[0]
	}
	if atmz.NextID == 0 {
		actSum = append(actSum, ai.ActID...)
	} else {
		var praseSum []int
		lib.MapFree(MapGwardAutomatizmNextStringFromID)
		arr := AutomatizmNextStringFromID[atmz.NextID].next
		praseSum = append(praseSum, prase)
		actSum = append(actSum, ai.ActID...)
		for i := 0; i < len(arr); i++ {
			actImg, ok := ReadeActionsImageArr(arr[i])
			if !ok {
				continue
			}
			//if actImg, ok:=ActionsImageArr[arr[i]]; ok{
			actSum = append(actSum, actImg.ActID...)
			praseSum = append(praseSum, actImg.PhraseID...)
			//}
		}
		var wordsArr []int
		// получаем слова из всех фраз
		for i := 0; i < len(praseSum); i++ {
			wa := wordSensor.GetWordArrFromPhraseID(praseSum[i])
			if wa != nil {
				wordsArr = append(wordsArr, wa...)
			}
		}
		// создаем новую фразу из слов
		str := wordSensor.PhraseDetection(wordsArr, false)
		if len(str) > 0 {
			prase = wordSensor.DetectedUnicumPhraseID
		}
	}
	FirstSimbolID := wordSensor.GetFirstSymbolFromPraseID([]int{prase})
	_, verbal := CreateVerbalImage(FirstSimbolID, []int{prase}, ai.ToneID, ai.MoodID+19)
	nodeID := FindConditionsNode(baseID, lev2, actSum, tm, simbolID, verbal.ID)
	if nodeID > 0 {
		return nodeID
	}

	return 0
}

////////////////////////////////////////////////////////

// чтобы не повторять ответ еще раз после каждого игнорирования
var oldProvokatorAutomatizm *Automatizm

/*
	в случае отсуствия автоматизма в данных условиях - послать оператору те же стимулы, чтобы посмотреть его реакцию.

Создание автоматизма, повторяющего действия оператора в данных условиях
*/
func provokatorMirrorAutomatizm(sourceAtmzm *Automatizm, purposeGenetic *PurposeGenetic) {
	if sourceAtmzm == nil || purposeGenetic == nil {
		return
	}

	ActionsImageID, _ := CreateNewlastActionsImageID(0, 0, curActiveActions.ActID, curActiveActions.PhraseID, curActiveActions.ToneID, curActiveActions.MoodID, true)
	// NoWarningCreateShow=true
	// для фразы triggerPraseID создаем привязанный к ней автоматизм
	_, autmzm := CreateAtutomatizmNoSaveFile(detectedActiveLastNodID, ActionsImageID)
	// NoWarningCreateShow=false
	if autmzm != nil {
		oldProvokatorAutomatizm = autmzm
		//autmzm.BranchID += linkID // не привязывать к узлу
		autmzm.Usefulness = 1          // авторитарная полезность
		SetAutomatizmBelief(autmzm, 2) // сделать автоматизм штатным, т.к. действия авторитарно верные (копируем действия оператора)

	}
	// и тут же запустить реакцию с ожиданием ответа
	setAutomatizmRunning(autmzm, purposeGenetic)
}

//////////////////////////////////

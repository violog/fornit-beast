/* ПАССИВНЫЙ (бесцельный) РЕЖИМ РАЗМЫШЛЕНИЯ https://scorcher.ru/thems/29/Passivniy_rezhim_myshleniya.htm
В отличие от целевого мышления о проблеме, пассивное мышление не имеет определенной цели,
но оно всегда начинается с выбора кадра эпиз.памяти (опорного кадра) с которого начинается такое мышление.
В результате прогностического прохода по эпиз.памяти возникают новые опорные кадры для следующего цикла.
При этом правила из кадров могут изменяться так, что возникают новые цепочки действий с их прогнозами.
Приоритет выбора опорных кадров задается в func GotoDreaming.
В первую очередь это - объективный стимул экстремального значения.
Затем субъективно найденный стимул экстремального значения из прошлого. и т.д.

При любом выборе опорного кадра мышление развивается одним алгоритмом - func gotoPassiveMaind(memFr *EpisodicTreeNode)

Обработка актуальных значимостей в режиме свободном от стимулов - пассивный режим,
активный в состоянии, когда Beast не занято выполнением какой-либо задачи, связанной со стимулами вопринятия,
бездействует, грезит наяву или образатывает недоосмысленное.
При актуальных доминантах в этом режиме
фантазируются попытки решения, рассуждения об этом, поиски методов по аналогии с моделями понимания и т.п.

Так же режим активируется при пассивном наблюдении за внешними объектами
- стимулами, не требующими немедленного реагирования.
Стимул, не имеющий модели понимания, вызывает беспокойство и активирует пассивный режим
чтобы исследовать признаки образа, найти ассоциации к ним в личном опыте.
В принципе такой анализ должен проводиться при каждом стимуле, не только при IsDream и откладывать в кеш осмысление нового.

В случае удачного поиска по моделям понимания и наблюдения за объектами (камень ловко катится)
формируется правило, привязанное к данной модели понимания,
которые могут быть использованы для пробных действий и формирование правил в этиз.памяти.
Правила, привязываемые к объектам моделей понимания, формируются при отзеркаливаним действий объекта и эффекта.

При появлении нового стимула режим дорабатывает текущую запущенную функцию и завершается,
оставляя инфо-картину прерванных мыслей для последующей активации пассивного режима.
Но если проблема решается на 1 или 2-м уровне, то режим не прерывается.

Таким образом, режим возникает на 3-м уровне осознания стимула.
Если после данного стимула не находится тема для размышления, следующий может ее дать.
Если нет стимулов более 30 пульсов и нет пассивного размышления, то запустить новый поиск темы.

ПАССИВНЫЙ РЕЖИМ РАЗМЫШЛЕНИЯ имеет метку в цикле cycle.dreaming
так что отдельные циклы могут продолжать пребывать не в режиме размышления (тут сложно - нужно бы решить более универсально...)
Стимул может начать новый главный цикл, но не закрывать режим размышления в уже имеющихся циклах,
который перрывается чтобы выделить ресурсы на текущую проблему, и восстанавливается после ее решения.

Во сне работают оставшиеся циклы дремы и могут активироваться новые циклы пассивного размышления.

func IsDreamMainProcess()// в главном цикле - режим пассивного размышления

ЗАДЕЙСТВАННЫЕ ИНФО-ФУНКЦИИ:
infoFunc10(cycle) - вспомнить самое значимое, что было для начала размышения
func infoFunc18 - Чем именно грозит наиболее значимое из параметров CurrentInformationEnvironment
func infoFunc5 - цикл мышления об экстремально важном объекте extremImportanceObject
func infoFunc20 - Запустить процесс осмысления НАЗАД эпиз.памяти infoFunc5()
func infoFunc202 - Запустить процесс осмысления ВПЕРЕД эпиз.памяти infoFunc5()


В пассивном мышлении могут находится альтернативные правила, которые надо бы запоминать (иначе нафиг нужно такое мышление).
То, что память модифицируется новым опытом - доказано.

Алгоритм описан в комментах над func GotoDreaming(cycle *cycleInfo)
*/

package psychic

import (
	"BOT/lib"
	"strconv"
)

// //////////////////////////////////////////////////////////////////
// время ожидания до активации пассивного режима размышлнения
var waitingTimeForDreamRejim = 10 // может думать в период ожидания ответа

var isDreamInterrupt = false // прерывание процесса размышления с новой актуальной проблемой после стимула. Снимается после решения проблемы или перевода в доминанту.

///////////////////////// ПОддержка шагов итерации пассивного мышления
/* Изменением инфо-картины при шаге итерации Пассивного режима
происходит только в главном цикле и в сновидениях
*/
type passiveEnvironment struct {
	mood          int // сила ощущаемого настроения: плохо -10...0...10 хорошо
	emotionID     int
	problemNodID  int
	danger        bool
	stimulID      int // акция оператора типа ActionsImage
	answerID      int // акция Beast типа ActionsImage
	extremID      int // новый объект для следующей итерации может быть как stimulID, так и answerID
	effect        int // значимость: -10...0...10
	episodFrameID int // кадр эпизода, на котором закончилась итерация
}

// текущее значение после шага пассвного осмысления
var newPassiveEnvironment passiveEnvironment

// обновление инфо-картины при шаге пассивного осмысления
func updateInformationEnvironment() {
	CurrentInformationEnvironment.PsyMood = newPassiveEnvironment.mood
	CurrentInformationEnvironment.PsyEmotionId = newPassiveEnvironment.emotionID
	CurrentInformationEnvironment.danger = newPassiveEnvironment.danger
	CurrentInformationEnvironment.ActionsImageID = newPassiveEnvironment.stimulID
	CurrentInformationEnvironment.AnswerImageID = newPassiveEnvironment.answerID
	CurrentInformationEnvironment.ActualEpisodicMemoryID = newPassiveEnvironment.episodFrameID
	CurrentInformationEnvironment.ExtremImportanceObjectID = newPassiveEnvironment.extremID
}

/*
	метка исчерпывания прохода по цепочкам значимости от начального эпизода

Определяется в
*/
var endOfPassiveIteratuins = true

/////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
/*
	начать процесс ПАССИВНЫЙ РЕЖИМ РАЗМЫШЛЕНИЯ
и новые итерации после прохода предыдущих, которые запускаются текущей активной функцией.
Тут приоритетная последовательность видов нахождения объектоа для размышления.

АЛГОРИТМ
С функции func GotoDreaming(cycle *cycleInfo)
 начинается цикл пассивного режима. В конце фоункции, если при обработке было найдено продолжение, то снова вызывается func GotoDreaming. Это - цикл в цикле cycle *cycleInfo.
1. При первом проходе выявляются наиболее значимые объекты внимания и по ним находятся наиболее значащий один кадр эпиз.памяти из всех найденных.
2. По этому кадру начинается шаг обработки func gotoPassiveMaind(cycle *cycleInfo, memFr *EpisodicTreeNode).
В func getNextHistoryEpisodicFrameArr находятся следующие после заданного кадры эпизодов в исторической цепочке до прерывания (пустой кадр) БЕЗ УСЛОВИЙ кроме настроения, что и дает диапазон фантазирования.
после этого шага имеем исходный кадр frm0 памяти и полученный эфктремальный frm2.
4. каждый шаг завершается обработкой с записью новой информации в общую инфокартину (func updateInformationEnvironment).
При проходе, если исходным кадром был с неопределенной значимостью стимула, то эта значимость изменяется найденными, усредняясь.
Если это было нахождение наиболее экстремального продолжения цепочки, и найдена высокая значимость (> insiteCompare = 5), процесс пассивного мышления прерываниеся инсайтом с переходом цикла в главный и его осознанием. Если это случилось во время сновидения - пробуждение.
В обоих случаях записываются два кадра эпиз.памяти подряд (frm0 и frm2) - как результат осознания инсайта.

*/
var curCondintions []int                    // совокупность условий для поиска по эпизодам для findEpisodicBrangeFromObject
var needOptimisationFrame *EpisodicTreeNode // нужно переписать в этом кадре новую значимость для стимула
var needOptimisationFrameCount = 0

func GotoDreaming(cycle *cycleInfo) {
	if isDreamInterrupt {
		return
	}

	cycle.dreaming = true
	if IsSleepingDream {
		cycle.log += "Сновидение пассивного размышления.<br>"
	} else {
		cycle.log += "Состояние пассивного размышления.<br>"
	}

	var curActionImgID = 0                    // найти объект для начала прохода
	var curEpisodicTreeNode *EpisodicTreeNode // этот эпизод нужно оптимизировать после прохода
	if endOfPassiveIteratuins {               // еще не было проходов найти объект для начала пассивного мышления
		needOptimisationFrame = nil
		needOptimisationFrameCount = 0
		curCondintions = []int{ // сначала текущие условия, потом они могут загруляться и изменяться
			CurrentCommonBadNormalWell,
			CurrentEmotionReception.ID,
			detectedActiveLastProblemNodID,
			curActionImgID}

		if cycle.isMainCycle && !IsSleepingDream { // только для главного цикла
			isRepressionStimulsNoise = true //Подавление мешающих стимулов при серьезном поиске решений проблемы.
		}
		//В первую очередь - объективный стимул экстремального значения
		if extremImportanceObject != nil {
			curActionImgID = extremImportanceObject.objID
		}
		if curActionImgID == 0 {
			///выбрать наиболее значимое из еще не угасших циклов кроме главного
			objectsFromPhoneCyclesArr()                                                   // TODO
			if objectsForPassiveToughtArr != nil && len(objectsForPassiveToughtArr) > 0 { // мыслить об этом
				curActionImgID = objectsForPassiveToughtArr[0]
				objectsForPassiveToughtArr = objectsForPassiveToughtArr[0:] // удалить использованный

				if IsSleepingDream {
					resetMineCycleAndBeginAsNew() // погачить главный цикл и начать новый

				} else {
					if !cycle.isMainCycle {
						endBaseIdCycle(cycle.ID) // погасить этот цикл
					}
				}
			}
		}
		if curActionImgID == 0 {
			//Стимул, не имеющий модели понимания, вызывает беспокойство
			if CurrentInformationEnvironment.IsUnknownActionsImageID > 0 {
				curActionImgID = CurrentInformationEnvironment.IsUnknownActionsImageID
			}
		}
		if curActionImgID == 0 {
			/* при опасной ситуации в первую очередь думаем про ответное действие,
			по статусу состояния отсуствия автоматизма: automatizmStatus
			0 - сброс рассматривания автоматизма
			1 - если автоматизм заблокирован
			2 - нет автоматизма или есть старый автоматизм, например, при игнорировании Ответа оператором
			*/
			if CurrentInformationEnvironment.veryActualSituation || CurrentInformationEnvironment.danger { // опасно
				if automatizmStatus > 0 && mentalInfoStruct.motorAtmzmID > 0 {
					curActionImgID = mentalInfoStruct.motorAtmzmID
				}
			}

		}
		if curActionImgID == 0 {
			// думаем про доминанту
			if EvolushnStage > 4 {
				// TODO
			}
		}
		if curActionImgID == 0 {
			/*находим недооцененные по значимости стимула неучительские кадры, т.е. с PARAMS[2]==0
			первые 3 объекта с конца, они заполнятся в эпиз.памяти и не будут здесь возникать
			*/
			curEpisodicTreeNode = nil
			unknownObjectsFromMemoryArr()
			if unknownObjectsFrameArr != nil && len(unknownObjectsFrameArr) > 0 { // мыслить об этом
				curActionImgID = unknownObjectsFrameArr[0].Trigger
				curEpisodicTreeNode = unknownObjectsFrameArr[0]     // этот эпизод нужно оптимизировать после прохода
				unknownObjectsFrameArr = unknownObjectsFrameArr[0:] // удалить использованный
			}
		}
	}
	////////////////////////////////////
	if endOfPassiveIteratuins {
		if curActionImgID > 0 { // начать проход пассивного мышления
			// задать условия поиска эпизода
			curCondintions = []int{
				CurrentCommonBadNormalWell,
				CurrentEmotionReception.ID,
				detectedActiveLastProblemNodID,
				curActionImgID}
		}
	} else {
		// задать условия поиска эпизода из newPassiveEnvironment, могут изменять только в сторону упрощений!
		curCondintions = []int{
			newPassiveEnvironment.mood,
			newPassiveEnvironment.emotionID,
			newPassiveEnvironment.problemNodID,
			newPassiveEnvironment.extremID}
	}

	if curActionImgID == 0 { // ничего нет для пассивного мышления
		if cycle.isMainCycle {
			if !IsSleepingDream {
				isRepressionStimulsNoise = false // умолчательное состояние Подавления мешающих стимулов
			}
			if IsSleepingDream {
				id := cycle.ID
				endBaseIdCycle(cycle.ID) // погасить этот цикл
				// найти новый главный цикл по значимости, кроме удаляемого
				foundMainCycle(id)
			}
		} else {
			endBaseIdCycle(cycle.ID) // погасить этот фоновый цикл
		}
		//  закончить пассивное мышление в этом цикле
		return
	}
	/////////////////////////////////////

	if curEpisodicTreeNode == nil { // м.б. определен и тогда не нужно findEpisodicBrangeFromObject
		/* найти самый экстремальный опорный кадр эп.памяти со стимулом extremImportanceObject
		прямыми правилами (а не учительскими)
		*/
		curEpisodicTreeNode = findEpisodicBrangeFromObject(curCondintions)
	} else {
		needOptimisationFrame = curEpisodicTreeNode // нужно переписать в этом кадре новую значимость для стимула
	}

	///////////////////////  эпизод начала мышления выбран
	if curEpisodicTreeNode != nil { // т.к. не смотрим учительские правила, то может не быть найдено

		// с curEpisodicTreeNode кадра жпиз.памяти НАЧИНАЕМ ИТЕРАЦИЮ ПАССИВНОГО МЫШЛЕНИЯ
		endOfPassiveIteratuins = gotoPassiveMaind(cycle, curEpisodicTreeNode) // один шаг обработки

		if needOptimisationFrame != nil { // нужно переписать в этом кадре новую значимость для стимула
			if newPassiveEnvironment.effect != 0 {
				// но условия могут отличаться! нужно усреднять эффект
				needOptimisationFrameCount++
				needOptimisationFrame.PARAMS[2] = int((needOptimisationFrame.PARAMS[2]*(needOptimisationFrameCount-1) + newPassiveEnvironment.effect) / needOptimisationFrameCount)
			}
		}

		// осознание результата, дающее контекст инфо-картины
		if cycle.isMainCycle { // только для главного цикла
			updateInformationEnvironment()
		}

		// исчерпание цепочки или ИНСАЙТ
		var insiteCompare = 5
		if endOfPassiveIteratuins || lib.Abs(newPassiveEnvironment.effect) > insiteCompare {
			// иссякла цепочка образов, конец пассивного мышления высокий эффект, требующий инсайта
			// если найдено что-то важное, то - инсайт (просыпание после сновидения).
			if lib.Abs(newPassiveEnvironment.effect) > insiteCompare {
				// ЭТО ИНСАЙТ
				setAsMaimCycle(cycle.ID) //	сделать цикл главным
				if IsSleeping {
					wakingUp() // разбудить
				}
				/*зписать пару кадров подряд в эпиз.память saveBeginEpisodicTreeNode + curEpisodicTreeNode
				получилась фантастическая цепочка
				*/
				if true {
					saveNewEpisodic(saveBeginEpisodicTreeNode.Trigger, saveBeginEpisodicTreeNode.Action, saveBeginEpisodicTreeNode.PARAMS[0], saveBeginEpisodicTreeNode.PARAMS[2])
					saveNewEpisodic(curEpisodicTreeNode.Trigger, curEpisodicTreeNode.Action, curEpisodicTreeNode.PARAMS[0], curEpisodicTreeNode.PARAMS[2])
				}
			}

			if !IsSleepingDream {
				isRepressionStimulsNoise = false // умолчательное состояние Подавления мешающих стимулов
			}
			return
		}
		GotoDreaming(cycle) // новая итерация
		return
	}

	return
}

//////////////////////////////////////////

// если нет главных циклов с пассивным мышлением, снять блокировку стимулов

////////////////////////////////////////////////////
/* если было прерывание, то восстановить инфу о прерванном цикле InterruptImage

 В режиме сновидений продолжается нормальная работа функции, но она ограничивается третьим уровнем.
т.е. могут выполняться моторные действия. Они не пройдут на Пульт, но в остальном функционал сохранится:
будет тот же период ожидания, выборка Правил, переактивации, все, без исключения.
и все идет как обычно с учетом Целей и поиском решений, если кончился цикл по последнему стимулу с пульта.

Уже было infoFunc10(cycle) //вспомнить самое значимое, что было для начала размышения
*/
func restoreDreamingProcess(cycle *cycleInfo) bool {
	cycle.lastProcessID = "dreamingProcess"
	//		infoFunc20(cycle)// выбрать последний пустой кадр эпиз.памяти dreamingEpisodeHistoryID и запустить процесс осмысления

	// при возникновении - cycle.log+="Начать сновидение или мечтание(пассивное размышление, не в ответ на Стимул).<br>"

	//  если было прерывание, то восстановить инфу о прерванном цикле InterruptImage
	if InterruptMemory != nil && len(InterruptMemory) > 0 {
		// закончился цикл мышления и есть прерванные циклы - нужно вернутьсЯ к последнему.
		lastImg := rememberInterruptImage(cycle) // выборка прерванного размышления из стека и начало размышления об этом
		if lastImg != nil {
			cycle.log += "Вспоминание прерванного шага мышления: " + strconv.Itoa(cycle.ID) + "<br>"
			//return false // был перезапуск
		}
	}
	return false
}

////////////////////////////////////////////////////////

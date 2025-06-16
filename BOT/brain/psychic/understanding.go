/*
Описание принципов:
Про циклы мышления Beast fornit.ru/66140
Схема процесса осознания Beast fornit.ru/66141

По каждому стимулу активируется функция 1 и2 уровней осознания:
func consciousnessElementary()
Если остается проблема с решением, как ответить, то запускаются циклы мышдения:
func consciousnessThinking(cID int)

С каждым новым стимулом появляется свой цикл мышления типа type cycleInfo struct
а уже имеющиеся становятся неглавными (фоновыми, неосознаваемыми).
Все циклы вызываются с помощью диспетчера циклов func dispetchConsciousnessThinking(),
проходя поочередно. Очередь запускается по каждому пульсу.
*/

package psychic

import (
	"BOT/brain/gomeostas"
	"BOT/lib"
	"sort"
	"strconv"
)

// ////////////////////////////
// для тестирования режим: проверка только третьего уровня, без вмешательства первых
var only3ltvelMind = false

// для передачи в пульт при ответе бота - на каком уровне осмысления был дан ответ
var levelOfRunAutomatizm = 0

// only3ltvelMind =true

//////////////////////////////////

// при включении и просыпании - 1 раз, в том числе при просыпении от Сторожевого стимула
var IsFirstActivation = true // первый вызов consciousnessElementary при включении или просыпании -> false
func WakeUpping() { // из sleep/sleep_process.go
	prepareImportanceObjectMenory() // очистить все
	//	IsFirstActivation =true  // только что проснулся - есть в func PsychicCountPuls(
	ReadiStatus = 1 //готовность Beast Для пульта:
}

// сбор строк о проессе осмысления дял Пульта
var conscienceStatus = "" // начинается с каждой объективной активацией вписывается в начало лога cycle.log

//////////////////////////////////////////////

// сколько раз был стимул от оператора после последнего запуска Ответа
var stimulCount = 0

// /////////////////////////////////////////////////////////
var idlenessType = 0 //0 - нет лени, 1 - гомео-лень, 2- осознанная лень Только для вывода на Пульт

var oldThemeImageType = 0 // старый Type образа темы

//var cikleExtremImportanceObject *extremImportance // объект внимания в цикле (перекрывается новым, очистка - при запуске цикла)

/*
	automatizmStatus значения:

0 - сброс рассматривания автоматизма
1 - если автоматизм заблокирован
2 - нет автоматизма или есть старый автоматизм, например, при игнорировании Ответа оператором
*/
var automatizmStatus = 0 // для передачи действий в новый цикл ()

var atmtzmActualTreeNode *Automatizm // отслеживание штатного автоматизма
var atmtzmActualTreeNodeID = 0

var mentAtmzmActualFuncs []int // из функции getFavoritInfoFunc() - выдать номера инфо-фукнции, привычные для данных условий

var isConfusion = false //не было моторного ответа на прошлый стимул, а уже последовавл новый

var show_all_logs = true //false - скрывать холостые, true - показывать все сообщения лога

/*
	эффект от запуска автоматизма сначала == его успешности и меняется в периоде ожидания при Стимуле.

От этого зависит активность мыслей:
при позитивном эффект нет необходимость искать новые действия.
при негативном (после ответа оператора) - нужно искать новые решения.

	if motorActionEffect>0 && c.dreaming==0{// не нужно искать решение, раз уже был нормальный ответ
		return false
	}
*/
var motorActionEffect = 0

//////////////////////////////////////////////////////

/////////////////////////////////////////////////////
/* Главная, активная с каждым ориентировочным рефлексом функция циклов осмысления
для поддержания информационной среды и произвольности.

*/
func consciousnessElementary() bool { // return true

	//levelOfRunAutomatizm = 0

	if EvolushnStage < 4 {
		return false
	}

	mentalInfoStruct.noOperatorStimul = false

	// первый проход - при включении и при просыпании func WakeUpping() из sleep/sleep_process.go
	if IsFirstActivation {

		IsFirstActivation = false
		initMentalMemories()
		clinerCycleLogsFiles()
	} else { //if IsFirstActivation  {
		/* Не дергать func consciousnessElementary()
		при незначительных изменениеях гомео-параметров
		*/
		if ActivationTypeSensor == 1 { // активация по изменению гомео-параметров
			// насколько изменились параметры
			if !gomeostas.GetGomeoParsDiff() {
				return false // недостаточно сильное изменение, чтобы привлечь внимание
			}
			if (PulsCount - curActiveActionsPulsCount) > 10 { // прошло > 10 пульсов со времени последнего стимула от оператора
				/* но измение гомео-параметров не могло произойти из-за стимула оператора
				так что можно включать режим провокации
				*/
				mentalInfoStruct.noOperatorStimul = true // не было стимула от оператора > 10 сек при значительном изменении гомео-параметров
			}
		}
	}

	lib.NoReflexWithAutomatizm = false // пока еще можно показывать акции рефлексов с автоматизмами в одной плашке
	///////////////////////////////////////////////////////////

	/* освободить запрет на запуск автоматизма (обычно запущенного ментально)
	wasRunTreeStandardAutomatizm=true устанавливается при любом запуске автоматизма в func RumAutomatizm
	ТЕПЕРЬ всегда после запуска автоматизма LastRunAutomatizmPulsCount >0
	*/
	//wasRunTreeStandardAutomatizm = false

	//initInfoFunc8pars()

	if IsSleeping { // - в процессе сна без сновидения
		return false
	}
	endDereamsCycles(true) // погасить все главные циклы дремы, оставив фоновые

	// с каждым вызовом func consciousnessElementary начинать новый главный цикл мышления
	mainС := resetMineCycleAndBeginAsNew()
	if mainС != nil {
		/* нужно чтобы сразу появилось detectedActiveLastProblemNodID для работы с эпиз.памятью на первых двух уровнях
		В infoFunc8 происходит активация дерева понимания проблемы
		*/
		infoFunc8(mainС)
	}

	if stimulCount > 1 { //не было моторного ответа на прошлый стимул, а уже последовавл новый
		isConfusion = true
		//conscienceStatus+="Произошел конфуз: не было моторного ответа на прошлый стимул, а уже последовавл новый<br>"
	}

	// начать новую последовательность ID инфо-функций, вызываемых после очередной активации
	// НЕТ! infoFuncSequence=nil

	if !only3ltvelMind { //проверка только третьего уровня, без вмешательства первых Включается наверху!

		/*Образ, не имеющий модели понимания, вызывает беспокойство и активирует пассивный режим – исследовать признаки образа, найти ассоциации к ним в личном опыте.
		Это определяется при каждом стимуле.
		*/
		CurrentInformationEnvironment.IsUnknownActionsImageID = 0 // сбросить внимание к прежнему такому образу
		if curActiveActions != nil {
			if isUnknownActionsImage(curActiveActions.ID) { //знаком ли стимул для данных условий? Новизна.
				CurrentInformationEnvironment.IsUnknownActionsImageID = curActiveActions.ID
			}
		}
		//////////////////////////////  ПЕРВЫЙ УРОВЕНЬ  /////////////////////////////////////////////
		/* если не первом уровне автоматизм запускается, то return false
		если не запускается то проблема передается на второй уровень,
		т.е. return true не должно быть на 1-м уровне, кроме случаев:
		1) если запускается альтернативный штатному автоматизм и нужно заблокировать штатный,
		2) были использованы Правила, но бесрезультатно и нет смысла снова пробовать на втором уровне.
		*/

		if CurrentInformationEnvironment.veryActualSituation || CurrentInformationEnvironment.danger {
			// только если важная ситуация нужно беречь вычислительные ресурсы
			isDreamInterrupt = true // прерывание процесса размышления с новой актуальной проблемой после стимула.
		}

		conscienceStatus += "Отработка первого уровня осмысления<br>"
		motorActionEffect = 0
		mentalInfoStruct.prognoseEffect = 0

		/// экстремально важные объекты внимания для мышления о них
		if detectedActiveLastProblemNodID > 0 { //без detectedActiveLastProblemNodID невозможно искать getExtremImportanceObject()
			getExtremImportanceObject() // найти extremImportanceObject
		}

		//  1 и 2 уровни - только для обработки рвущегося на выполнение автоматизма

		/*!!Не всегда currentAutomatizmAfterTreeActivatedID определяеется ДО understandingSituation(1)
		поэтому в func consciousnessElementary есть свой atmtzmActualTreeNodeID := getAutomatizmFromNodeID(detectedActiveLastNodID)
		и тогда atmtzmActualTreeNodeID оказывается из предыдущей активации дерева автоматизмов.
		*/
		atmtzmActualTreeNodeID = 0 // atmtzmActualTreeNodeID определяется только здесь!
		if wasCurrentAutomatizmAfterTree > 0 && currentAutomatizmAfterTreeActivatedID > 0 {

			atmtzmActualTreeNodeID = currentAutomatizmAfterTreeActivatedID
			// это значит, что 	atmtzmActualTreeNodeID - из старой активации, не требует осмысления как актуальный рвущийся автоматизм

			/*распознать: не навязывают ли стимулом то, что не соотвествует текущей Теме и Цели?
			Меняет состояние CurrentInformationEnvironment.isStimulToForce:
			true - текущий стмул навязывает то, что не соотвествует текущей Теме и Цели
			*/
			recognizerStimulToForce()
		} else {
			atmtzmActualTreeNodeID = 0 // если не нашелся автоматизм на стимул, нужно обнулить ID предыдущего активного автоматизма, иначе он полезет как рвушийся на исполнение и не даст запустить infoFunc()13
			//wasCurrentAutomatizmAfterTree=0// на всякий случай
			//atmtzmActualTreeNodeID = getAutomatizmFromNodeID(detectedActiveLastNodID)
		}
		////////////////////////////////////

		if isIgnoreAutomatizmID(atmtzmActualTreeNodeID) {
			atmtzmActualTreeNodeID = 0
			conscienceStatus += "Игнорирующий штатный автоматизм - расценивается как отсуствие реакции.<br>"
		}

		if atmtzmActualTreeNodeID == 0 {
			// атасная ситуация
			if (EvolushnStage == 4 || CurrentInformationEnvironment.veryActualSituation) && CurrentInformationEnvironment.danger {
				// срочно, не думая, найти аварийный автоматизм, хоть что-то из привязанных к ветке или новый привязать, сделать штатным
				atmtzmActualTreeNodeID = findAlternativeAtmtzm()
			}
		}

		if atmtzmActualTreeNodeID > 0 && CurrentInformationEnvironment.isStimulToForce {
			//текущий стмул навязывает то, что не соотвествует текущей Теме и Цели
			atmtzmActualTreeNodeID = 0
			conscienceStatus += "Автоматизм не соответствует текущей теме и цели.<br>"
			/*позволить пройти дальше чтобы найти альтернативу на 2 уровне ПРАВИЛА
			а если нет, то блокировать всеболее низкоуровневое, перейдя к циклам мышления.
			*/
		}

		// есть ли штатный мот.автоматизм и нужно ли его менять или задумываться
		if atmtzmActualTreeNodeID > 0 { // есть рвущийся на выполнение штатный автоматизм
			//ниже точно такой же вызов - зачем 2 раза одно и тоже?
			//conscienceStatus+="Есть штатный автоматизм <b> <span style='cursor:pointer;color:blue' onClick='show_automatizms(" + strconv.Itoa(atmtzmActualTreeNodeID) + ")'>" + strconv.Itoa(atmtzmActualTreeNodeID) + "</span>" + "</b><br>"
			mentalInfoStruct.noStaffAtmzmID = false
			mentalInfoStruct.motorAtmzmID = atmtzmActualTreeNodeID // для последующего использования в инфо-фукнциях
			/* Период ожидания ответа LastRunAutomatizmPulsCount при поочередном Стимуле-Ответе есть всегда.
			   А здсь - поиск Ответа именно после каждого Стимула. Так что LastRunAutomatizmPulsCount в функции не учитываем.
			*/
			////////////////////////////// 1 уровень ////////////////////
			// ПЕРВЫЙ УРОВЕНЬ, самый примитивный уровень:

			//НЕТ!!! atmtzmActualTreeNode := GetBelief2AutomatizmListFromTreeId(detectedActiveLastNodID)
			/* учитывается именно тот автоматизм, что рвется на выполнение при активации дерева автоматизмов,
			   даже если он подобран "мягким алгоритмом" в getAutomatizmFromNodeID
			   При атасе он выполняется не раздумявая, иначе подвергается сомнению в infoFunc6()
			*/

			//			atmtzmActualTreeNode = AutomatizmFromId[atmtzmActualTreeNodeID]
			atmtzmActualTreeNode, ok := ReadeAutomatizmFromId(atmtzmActualTreeNodeID)
			if ok && atmtzmActualTreeNode != nil { // есть незаблокированный автоматизм
				if atmtzmActualTreeNode.Usefulness >= 0 { //незаблокированный, м.б. нештатным!

					//	isDangerousModelAutomatism(atmtzmActualTreeNode.ActionsImageID) // проверка функции

					conscienceStatus += "Есть моторный автоматизм <b><span style='cursor:pointer;color:blue' onClick='show_automatizms(" + strconv.Itoa(atmtzmActualTreeNodeID) + ")'>" + strconv.Itoa(atmtzmActualTreeNodeID) + "</span>" + "</b><br>"
					// Если период сомнений или важная ситуация и нет опасности, то - ПОДВЕРГНУТЬ СОМНЕНИЮ автоматизм.
					if (EvolushnStage == 4 || CurrentInformationEnvironment.veryActualSituation) &&
						!CurrentInformationEnvironment.danger {

						//ПОДВЕРГНУТЬ СОМНЕНИЮ автоматизм, на время размышления заблокировать выплнение штатного
						/* Здесь можно, в зависимости от наработанной инфо-базы, подставлять более подходящий автоматизм, если он
						обнаруживается и привязать его к дереву.
						При этом другие позитивные автоматизмы дерева не блокируются.  В бессознанке выбирается штатный или самый привычный из них.
						Использотьва модели понимания и правила, м.б. значимости для этого в checkAutomatizm(atmtzmActualTreeNode)
						*/
						conscienceStatus += "Подвергнуть сомнению автоматизм в спокойной ситуации.<br>"
						//infoFunc6(cycle) // если все в порядке, он в infoFunc6() запустится, если запущен альтарнативный, то он стал штатным,
						// func checkAutomatizm - бездумная версия проверки
						resA := checkAutomatizm(atmtzmActualTreeNode)
						if resA == nil {
							mentalInfoStruct.noStaffAtmzmID = true
							mentalInfoStruct.motorAtmzmID = 0 // чтобы в usualThinkProcess/infoFunc8 не вернуло true, иначе в usualThinkProcess не дойдет до infoFunc2, а в самой infoFunc2 не дойдет до func13
							atmtzmActualTreeNode = nil        // чтобы сначало отработал поиск в ментальных правилах findSuitableMentalFunc(), а потом уже func13
							//были использованы Правила, но бесрезультатно и нет смысла снова пробовать на втором уровне
							levelOfRunAutomatizm = 1 // для передачи в пульт при ответе бота - на каком уровне осмысления был дан ответ
							conscienceStatus += "1 уровень мышления. Автоматизм ID=" + strconv.Itoa(atmtzmActualTreeNodeID) + " подвергнут сомнению и остановлен.<br>"
							isDreamInterrupt = false // прерывание процесса размышления с новой актуальной проблемой после стимула.
							return true              // блокировать старый штатный автоматизм потому, что nil означает опасную ситуацию и отсутствие в правилах положительной перспективы
						}
						if resA.ID == atmtzmActualTreeNodeID { // не изменился выбор
							/* если есть значимая новизна или атас, нужно проверять по правилам прогноз последствий
							   и если для данного действия есть групповое правило, кончающееся негативом, то не выполнять такое действие.
							   Пример:
							      Оператор: «ты кто?»
							      Beast: «бот»
							      Оператор: «уверен?»
							      Beast: «да»
							      Оператор: [Поощрить]
							      Beast: [радуется, смеется, симпатизирует]
							      Оператор: жмет на плашку индикации времени, чтобы прервать диалог и начать новый. Либо ждет, пока индикация сама погаснет.
							      Оператор: «кто ты?»
							      Beast: «Бог»
							      Оператор: «уверен?»
							      Beast: «да»
							      Оператор: [Наказать]
							      Beast: [плачет, прощает]
							      В итоге получатся 2 групповых правила, одно с положительным эффектом, другое с отрицательным.
							   	На Оператор: «кто ты?» нужно прогнозировать отрицательный ответ и не выдавать "бог".
							*/
							//							if EvolushnStage > 4 { //5-я сталия развития.
							// прогноз на основе правил. TODO
							needRulesPrognoz := false
							// при экстремальной отрицательной значимости
							if extremImportanceObject != nil && extremImportanceObject.extremVal < 0 {
								needRulesPrognoz = true
							}
							// в спокойной ситуации
							if !CurrentInformationEnvironment.veryActualSituation && !CurrentInformationEnvironment.danger && !isUnrecognizedPhraseFromAtmtzmTreeActivation {
								needRulesPrognoz = true
							}
							motorActionEffect = atmtzmActualTreeNode.Usefulness
							if needRulesPrognoz {
								// прогноз по данному типу действий автоматизма и правила из последнего кадра эпиз.памяти
								accuracy, effect := getPrognoze(resA)
								mentalInfoStruct.prognoseEffect = effect
								if accuracy > 0 && effect < 0 { //есть негативное предсказание
									//// пусть ищет другой
									motorActionEffect = 0
									mentalInfoStruct.motorAtmzmID = 0 // чтобы в usualThinkProcess/infoFunc8 не вернуло true, иначе в usualThinkProcess не дойдет до infoFunc2, а в самой infoFunc2 не дойдет до func13
									//return true // запретить штатный	НЕТ пусть пройдет до конца функции

									// понизить такой автоматизм, но правило не удалять!
									passivationAutomatizm(resA, -1)
									/* TODO нужно как-то учитывать в infoFunc2 mentalInfoStruct.prognoseEffect, особенно отрицательный.
									т.е. нужно учитывать прогноз (как информированность об эффекте) в соотвествующих инфо-функциях и не только
									*/
									conscienceStatus += "1 уровень мышления. Автоматизм ID=" + strconv.Itoa(resA.ID) + " есть негативный прогноз.<br>"
								}
							}

							//							}
							////////////////

							//motorActionEffect = atmtzmActualTreeNode.Usefulness
							if mentalInfoStruct.prognoseEffect >= 0 && !isUnrecognizedPhraseFromAtmtzmTreeActivation {
								return false // запустить штатный автоматизм
							} else {
								CurrentInformationEnvironment.needThinkingAboutAutomatizm = true
							}
						}
						if resA.ID != atmtzmActualTreeNodeID { // был выбран альтернативный из привязанных к ветке
							runConsciousnessAutomatizm(resA)
							levelOfRunAutomatizm = 1 // для передачи в пульт при ответе бота - на каком уровне осмысления был дан ответ
							conscienceStatus += "1 уровень мышления. Автоматизм ID=" + strconv.Itoa(atmtzmActualTreeNodeID) + " остановлен и запущен альтернативный ID=" + strconv.Itoa(resA.ID) + ".<br>"
							isDreamInterrupt = false // прерывание процесса размышления с новой актуальной проблемой после стимула.
							return true              // блокировать старый штатный автоматизм т.к. запущен альтернативный.
						}

						// если есть extremImportanceObject - посмотреть автоматизмы у extremImportanceObject как категории
						/*  старая поддержка категорий убрана, можно посмотреть в categorization.go и categorization_functions.go
						в архиве 2024_01_14_try_episodic_tree_rekease_FINAL_OF_3_VERSION
						ВНИМАНИЕ! любые взаимосвязи теперь находятся по дереву эпиз.памяти.
						т.к. поддержка категорий пока еще не применена нигде,
						то ОСТАВЛЯЕТСЯ НА БУДУЩЕЕ при необходимости.
						Тогда и будут написаны все необходимые функции.
						Пока что удаляется набор категорий и запись файла /memory_psy/category_model.txt
						который не нужен т.к. поиск по дереву эпиз.памяти достаточно быстрый.

												if extremImportanceObject!=nil{
													// выбрать лучший автоматизм из массива ActionsImage категорий экстремального объекта внимания
													res:=getBestAtmtzmFormCategoryExtremImportanceObject()
													if res{ // найден и запущен автоматизм
														return true
													}
												}
						*/

						/*попробовать найти подходящее в дочках дерева (категории - более ранний узел)
						  Категории по фразам уже смотрили выше:
						  alternatives:=getPhraseCategoryChildrens(wordSensor.CurretWordsIDarr)

						   посмотреть категории - все дочки более общего уровня - ФАКТИЧЕСКИ СМОТРИМ БУКВУ АЛФАВИТА...

						    старая поддержка категорий убрана, можно посмотреть в categorization.go и categorization_functions.go
						    в архиве 2024_01_14_try_episodic_tree_rekease_FINAL_OF_3_VERSION
						    ВНИМАНИЕ! любые взаимосвязи теперь находятся по дереву эпиз.памяти.
						    т.к. поддержка категорий пока еще не применена нигде,
						    то ОСТАВЛЯЕТСЯ НА БУДУЩЕЕ при необходимости.
						    Тогда и будут написаны все необходимые функции.
						    Пока что удаляется набор категорий и запись файла /memory_psy/category_model.txt
						    который не нужен т.к. поиск по дереву эпиз.памяти достаточно быстрый.


						  	alternatives := getStimulCategoryChildrens(atmtzmActualTreeNode.BranchID)
						  	if alternatives != nil {
						  		// просмотреть дочки категории и если есть такие экстремальные объекты с позитивом, сделать действие по наилучшему

						  		res := reseachAtmtzmFromStimulCategory(alternatives)

						  		if res { // найден и запущен автоматизм
						  			return true
						  		}
						  	}
						*/
					} else { // при опасности некогда думать, нужно действовать привычно
						conscienceStatus += "Опасная ситуация - запустить штатный автоматизм <b><span style='cursor:pointer;color:blue' onClick='show_automatizms(" + strconv.Itoa(atmtzmActualTreeNodeID) + ")'>" + strconv.Itoa(atmtzmActualTreeNodeID) + "</span>" + "</b>.<br>"
						motorActionEffect = atmtzmActualTreeNode.Usefulness

						// ЗАПИСАТЬ В ДЕРЕВО НОВЫЙ ЭПИЗОД со стимулом и ответом, хотя почему-то было newEpisodeMemory(0,0) без Правила
						/*	saveNewEpisodic(curActions.ID,atmtzmActualTreeNode.ActionsImageID,motorActionEffect) */
						// записать нулевое Правило:
						saveNewEpisodic(curActions.ID, 0, 0, 0)
						isDreamInterrupt = false // прерывание процесса размышления с новой актуальной проблемой после стимула.
						return false             // выполнить страрый штатный автоматизм
					} // если нет - далее искать альтернативу
				} else { // если автоматизм заблокирован
					/*удалить авторитарное Правило с таким действием.
										Иначе никогда не сработает checkForUnbolokingAutomatizm, см. ниже об этом
					НЕЛЬЗЯ УДАЛЯТЬ КАДР ЭПИЗЮПАМЯТИ!!!
					*/
					//conscienceStatus+="Удалить авторитарное Правило с действием заблокированного автоматизма<br>"
					/* НО при создании авторитарного правила (func fixNewRules(lastCommonDiffValue int)) определяется,
										есть ли автоматизм с действием оператора curActiveActionsID, и если у него atmtzmActualTreeNode.Usefulness<0 -
										снять блокировку и сделать штатным (checkForUnbolokingAutomatizm(curActiveActionsID))

					В func searchingRules(trigger int,rImg []int,condType int )(int,int){ делается проверка:
					если действие правила имеется в заблокированном автоматизме, то такое правило исключается и ищутся другие:
					blockExist:=checkBlockedAutomatizm(curR)
					*/

					//Нужно подумать о проблеме автоматизма или проявить инициативу, в общем, запустить func infoFunc25()
					CurrentInformationEnvironment.needThinkingAboutAutomatizm = true

					if EvolushnStage > 4 { // вместо плохого автоматизма - из успешной доминанты
						automatizmStatus = 1
					}
					/*заблокированный автоматизм посылается в RumAutomatizm() не смотря на то, что будет там гарантировано остановлен потому, что так будет виден ответ
					на пульте: автоматизм найден, но заблокирован. При этом mentalInfoStruct.motorAtmzmID = atmtzmActualTreeNodeID позволит потом разобраться с ним с помощью осмысления
					Если заблокировать и сделать переход на 2 уровень, это ничего не даст: там создается автоматизм по учительскому правилу, а не прямому, которое фиксирует существующую
					отработанную реакцию.
					То путь переходит на второй уровень осмысления правил для создания автоматизмов по учительскому правилу и при условии, что нет никакого автоматизма на такой стимул
					Это отзеркаливание старого, не опробованного опыта.*/
					// не позволять рефлексов! return true
				}
			}
			//if currentAutomatizmAfterTreeActivatedID > 0{ // есть рвущийся на выполнение автоматизм
		} else { // нет автоматизма или есть старый автоматизм, например, при игнорировании Ответа оператором
			conscienceStatus += "Нет штатного моторного автоматизма.<br>"
			mentalInfoStruct.noStaffAtmzmID = true
			mentalInfoStruct.motorAtmzmID = 0
			if EvolushnStage > 4 {
				automatizmStatus = 2
			}
			// conscienceStatus+="Нет штатного моторного автоматизма.<br>"

			//if isUnrecognizedPhraseFromAtmtzmTreeActivation { //при активации была нераспознанная фраза
			/* должна быть не только фраза isUnrecognizedPhraseFromAtmtzmTreeActivation,
			а любой стимул, на который нет автоматизма!
			и в таком контексте чтобы сработал диспетчер func infoFunc2
			*/
		}

		/////////////////////////////////////////////////////////////////

		//return false //для тестирования

		//////////////////////////////// 2 уровень ПРАВИЛА ////////////////////////////
		conscienceStatus += "Отработка второго уровня осмысления<br>"
		// не сброшена проблема оценки штатного автоматизма или нет штатного автоматизма
		if mentalInfoStruct.motorAtmzmID > 0 || mentalInfoStruct.noStaffAtmzmID {
			// ВТОРОЙ УРОВЕНЬ - попытка использования примитивных Правил

			/* нужно учитывать, что в правилах действия могут быть типа AmtzmNextString.ID
			- если ID действия превышает по величине prefixActionIdValue
			*/
			var rule Rule // найти Правило
			isFoundRule := false
			/*Найти в расчете на метод GPT последнее известное Правило по цепочке последних limit кадров эпиз.памяти
			  Найти самое свежее Правило, имеющее схожесть с последними кадрами эпизодической памяти
			  для текущей ветки дерева проблем detectedActiveLastUnderstandingNodID, а значит, для текущей совокупности условий.
			  В эпиз.памяти записываются Правила с учетом последовательности образов действий типа AmtzmNextString.ID

			  Последний Стимул при активации Дерева автоматизмов: curActiveActionsID, найти для него наиболее преспективную цепочку Правил,
			  а если нет, то единственное Правило и по нему - рекомендуемое действие.
			*/
			limitN := getLimitCountEM()
			getTargetEpisodicStrIdArr(curActiveActionsID, limitN)
			//выбранная целевая цепочка Правил откуда берем следующее Действие.
			if targetEpisodicStrIdArr != nil {
				/* есть цепочка с конечным плюсовым эффектом, значит можно так действвоать
				а первым членом может быть негативный эффект и если он слишком большой, то не годится
				*/
				rule = targetEpisodicStrIdArr[0]
				if rule.Effect > -3 {
					effect0 := rule.Effect
					if rule.Effect < 0 {
						// есть ли позитив в конце, превышающий негатив
						ruleLast := targetEpisodicStrIdArr[len(targetEpisodicStrIdArr)-1]
						if ruleLast.Effect > effect0 {
							isFoundRule = true
						}
					}
				}
			} else { // не найдено по GPT без учета учительский Правил
				// тогда просто ищем по любым Правилам
				rule = getSingleBestRule(3, curActiveActionsID)
				if rule.Action > 0 {
					isFoundRule = true
				}
			}
			if isFoundRule {
				//по правилу найти или создать (в случае AmtzmNextString) автоматизм и запустить его
				res := makeActionFromRooles(rule) //	!!!!!ПЕРЕДЕЛАТЬ для type Rule struct {
				isDreamInterrupt = false          // прерывание процесса размышления с новой актуальной проблемой после стимула.
				return res
			}

		} //if mentalInfoStruct.motorAtmzmID > 0{ // есть рвущийся на выполнение автоматизм

		// не сброшена проблема оценки штатного автоматизма
		if mentalInfoStruct.motorAtmzmID > 0 { // все еще есть рвущийся на выполнение штатный автоматизм
			conscienceStatus += "Остается проблема оценки штатного автоматизма на втором уровне.<br>"

			//atmtzmActualTreeNode := AutomatizmFromId[mentalInfoStruct.motorAtmzmID]
			atmtzmActualTreeNode, ok := ReadeAutomatizmFromId(mentalInfoStruct.motorAtmzmID)
			// если автоматизм уверенный и ситуация не опасна, то сбросить проблему
			if ok && atmtzmActualTreeNode.Usefulness > 1 && atmtzmActualTreeNode.Count > 3 &&
				CurrentInformationEnvironment.veryActualSituation &&
				!CurrentInformationEnvironment.danger {
				mentalInfoStruct.motorAtmzmID = 0 // пусть выполняется штатный автоматизм
				mentalInfoStruct.noStaffAtmzmID = false
				conscienceStatus += "На втором уровне запущен штатный автоматизм " + strconv.Itoa(mentalInfoStruct.motorAtmzmID) + " т.к. автоматизм уверенный и ситуация не опасна.<br>"
				isDreamInterrupt = false // прерывание процесса размышления с новой актуальной проблемой после стимула.
				return false
			} else {
				conscienceStatus += "Нужно подумать о проблеме оценки штатного автоматизма.<br>"
				// подумать об этом
				if CurrentInformationEnvironment.veryActualSituation || CurrentInformationEnvironment.danger {
					runNewTheme(12, 5)
				}
				CurrentInformationEnvironment.needThinkingAboutAutomatizm = true
			}
		}
		/////////////////////////////////////////
	} //if !only3ltvelMind{//проверка только третьего уровня, без вмешательства первых Включается наверху!

	// нужно решать проблему в цикле мышления в func consciousnessThinking
	conscienceStatus += "На 1-м и 2-м уровнях осмысления проблема не решена.<br>"
	blockingNewProblemTryCount = false // от быстрого увеличения счетчика проблем в цикле

	// есть ли штатный мент.автоматизм в узле дерева проблем, если есть сравнить полезность
	// ....

	if CurrentInformationEnvironment.isStimulToForce {
		levelOfRunAutomatizm = 2 // для передачи в пульт при ответе бота - на каком уровне осмысления был дан ответ
		conscienceStatus += "2 уровень мышления. Автоматизм ID=" + strconv.Itoa(atmtzmActualTreeNodeID) + " остановлен, так как не соответствует текущей цели.<br>"
		return true // блокировать всеболее низкоуровневое.
	}

	if currentAutomatizmAfterTreeActivatedID > 0 {
		// если это плохой автоматизм, или плохой прогноз в правилах по нему, то не запускать
		a, ok := ReadeAutomatizmFromId(currentAutomatizmAfterTreeActivatedID)
		if !ok || a.Usefulness < 0 || mentalInfoStruct.prognoseEffect < 0 {
			currentAutomatizmAfterTreeActivatedID = 0 // не запускать автоматизм, но разрешить рефлексы
		}
	}

	isDreamInterrupt = false // прерывание процесса размышления с новой актуальной проблемой после стимула.
	return false             // не блокировать последующий код ориентировочного рефлекса.
}

///////////////////////////////
/* сразу после первого прохода consciousnessElementary() запускается певый цикл
позволяющий оценить ситуацию и совершить действия даже без Стимула.
*/
func beginMentalCycle() {
	createNewCycleIteration()
}

////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////
/* Постоянная циркуляция циклов мышления, по каждому пульсу запуск dispetchConsciousnessThinking()
запуск имеющихся циклов по очереди, начиная с главного по каждому пульсу
Неглавные циклы запускаются раз в три пульса?
*/
var curMainCyckle *cycleInfo // текущий главный цикл
func dispetchConsciousnessThinking() {
	if EvolushnStage < 4 {
		return
	}
	if len(cyclesArr) == 0 {
		return
	}
	// сортировать по убыванию
	var mainID = 0
	keys := make([]int, 0, len(cyclesArr))

	noAnyCycle := true
	for id, v := range cyclesArr {
		if v == nil {
			continue
		}
		noAnyCycle = false
		if v.isMainCycle {
			curMainCyckle = v
			mainID = id
		} else {
			keys = append(keys, id)
		}
	}

	if noAnyCycle { // все циклы оказались погашены, нужно запустить главный
		cyclesArr = nil    // убираем следы погашенных циклов
		beginMentalCycle() // начать первый, главный цикл мышлления с вызовом func infoFunc8()
		return
	}

	//	lib.MapFree(MapGwardCyclesArr)
	//sort.Ints(keys)
	// отсортировать по убыванию
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	// запускать циклы по очереди, начиная с главного
	if mainID > 0 {
		cycle, ok := ReadecyclesArr(mainID)
		if !ok || cycle == nil {
			return
		}
		consciousnessThinking(mainID, cycle)
		cycle.count++
	}
	for _, id := range keys {
		//Неглавные циклы запускаются раз в три пульса?
		if PulsCount%3 == 0 {
			cycle, ok := ReadecyclesArr(id)
			if !ok || cycle == nil {
				return
			}
			consciousnessThinking(id, cycle)
			cycle.count++
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////

var curThinkingPulsCount = 0 // только для контроля
////////////////////////////////////////////////////////////////////////////////////////
/* ментальная часть consciousnessElementary - мышление: 3,4 уровни мышления

Сначала выясняется, стоит ли решать проблему и если да, то
если есть текущая проблема (нет автоматизма) то выбирается цель infoFunc8()
если нет проблемы с автоматизмом и позволяет атас, то смотрятся доминанты.
*/
func consciousnessThinking(cID int, cycle *cycleInfo) {

	if IsSleepingDream { // нужно запускать сновидение.
		/*Сон начинается с перевода главного цикла мышления в пассивный режим (сновидения)
		Если есть инсайт, то просырание и все заново,
		если нет инсайта, начинать по очереди фоновые циклы переводить в главные в пассивном режиме.
		Это - на уровне не func consciousnessElementary(), а в func dispetchConsciousnessThinking()->func consciousnessThinking

		Тут будут проходить все активные циклы мышления, но снивидение всегда будет для главного.
		*/
		if !cycle.isMainCycle {
			// перейти в главный
			cycle = curMainCyckle
			//начать сновидение
			GotoDreaming(cycle) // после исчерпания сюжетя сновижения цикл будет закрыт и установится новый главный.
		}

		return
	}

	if EvolushnStage < 4 {
		return
	}

	//	detectedActiveLastProblemNodID=0// может и не быть проблемы (автоматизм запущен или лень) и тогда пусть остается старая.

	// пустые циклы запускать в 5 раз реже Главнй не может замедляться.
	if !cycle.isMainCycle && cycle.idle {
		if PulsCount%5 != 0 {
			return
		}
	}

	// ступорный цикл не запускать
	if cycle.isStupor {
		return
	}

	if curThinkingPulsCount != PulsCount { // только для контроля

		curThinkingPulsCount = PulsCount
	}
	//////////////////////////////////////
	if correctRepeatCycleFuncStr(cycle) {
		/* предыдущий фрагмент итерации повторяется, ЧТО С ЭТИМ ДЕЛАТЬ?
		в природе в голове часто крутятся назойливые фрагменты...
		Уменьшить скорость вызовов.
		*/
		if PulsCount%5 != 0 {
			return
		}
	}
	cycle.func0Arr = nil        // сначала лог вызовов в течении 1 итерации func consciousnessThinking()
	updateCycleLogsFiles(cycle) // сохранить лог в файле

	//////////////////////////////////////////////////////////////////////

	/// ЭТО - ГЛАВНЫЙ ЦИКЛ - обновлять инфо-среду с каждым проходом
	if cycle.isMainCycle {
		refreshCurrentInformationEnvironment()
		if !IsSleeping { // - без сна
			//ТОЛЬКО ДЛЯ ГЛАВНОГО ЦИКЛА - период ожидания ответа с пульта на действие
			cycle.isWaitingPeriod = CurrentInformationEnvironment.IsWaitingPeriod
		}
	}
	//////////////////////////////////////////

	//////////////////////////////////////////
	if cycle.dreaming {
		//если было прерывание, то восстановить инфу о прерванном цикле InterruptImage
		if restoreDreamingProcess(cycle) {
			cycle.idle = false
			return
		}
	}

	// ПАССИВНЫЙ РЕЖИМ РАЗМЫШЛЕНИЯ
	if (PulsCount - curActiveActionsPulsCount) > waitingTimeForDreamRejim { // прошло > waitingTimeForDreamRejim пульсов со времени последнего стимула
		/* при эмоции с Базовыми контекстами Поиск, Игра, Гон, Агрессия, Страх
		начать размышлеие
		*/
		if !cycle.dreaming {
			if existsBaseContext(2) ||
				existsBaseContext(3) ||
				existsBaseContext(4) ||
				existsBaseContext(8) ||
				existsBaseContext(9) {

				GotoDreaming(cycle) // начать процесс ПАССИВНЫЙ РЕЖИМ РАЗМЫШЛЕНИЯ
				//процесс далее уже не идет
				return
			}
		}
		// блокировка запуска пассивного размышлнения до следующего стимула, который установит curActiveActionsPulsCount=PulsCount
		curActiveActionsPulsCount = 10000000000
	}

	/////////////////////////////////////////////////////////////

	//////////////////////  ЛЕНЬ? (осоловелость)
	/* детекция первичного ленивого состояния
	isIdleness() Учитывает Доминанту.
	*/
	idleness := isIdleness()

	//ТЕСТИРОВАНИЕ только третьего уровня, без вмешательства первых Включается наверху!
	if only3ltvelMind {
		idleness = false
		automatizmStatus = 2 // нет автоматизма
		mentalInfoStruct.motorAtmzmID = 0
		CurrentInformationEnvironment.needThinkingAboutAutomatizm = true // думать как сделать действие
	}

	if idleness { // ЛЕНЬ
		lookForUnactualDominant() // просмотрт доминант для закрытия неактуальных
		if lenessProcess(cycle) {
			cycle.idle = false
			return
		}
	} //if isIdleness()
	////////////////////////////////////

	/////////////////////////  НЕТ ЛЕНИ
	if mentalInfoStruct.ExtremObjID > 0 { // после возврата прерванного в func rememberInterruptImage
		extremImportanceObject = getExtremObjFromID(mentalInfoStruct.ExtremObjID)
		mentalInfoStruct.ExtremObjID = 0
	}

	if extremImportanceObject != nil {
		cycle.impObjID = extremImportanceObject.objID
		weight := lib.Abs(extremImportanceObject.extremVal)
		if weight > cycle.weight {
			cycle.weight = weight
		}
	}
	//////////////////////////////////////////

	/////////////////////////////// для нового цикла:
	if cycle.count == 0 {
		// освободить запрет на запуск автоматизмов.
		wasRunPurposeActionFunc = false

		// cycle.log+="Отработка третьего уровня осмысления  "+strconv.Itoa(cycle.ID)+" мышления с ID циикла "+onClickStr(cycle.ID,"show_cyckle","")+"<br>"
		/* Ментальное определение ближайшей Цели без текущей темы - УЖЕ БЫЛО В НАЧАЛЕ: resetMineCycleAndBeginAsNew()
		   и Активация дерева понимания проблемы ProblemTreeActivation().
		   Если досюда не дошло detectedActiveLastProblemNodID пусть остается старая проблема

		if cycle.isMainCycle { // только для главного цикла (не перезапускать для всех циклов!)
			// только в начале цикла 1 раз
			//if cycle.funcArr!=nil && !lib.ExistsValInArr(cycle.funcArr, 8) {
			// только здеь запускается infoFunc8 !
			 infoFunc8(cycle)
			//}
		}*/

		if isUnrecognizedPhraseFromAtmtzmTreeActivation { //при активации была нераспознанная фраза
			cycle.isWaitingPeriod = false
			if infoFunc2(cycle) {
				return
			}
		}

		if automatizmStatus == 1 { //если автоматизм заблокирован
			// запускать решение доминанты в подходящих условиях т.е. взять опыт из ранее решенных проблм
			res := runDominantaAction(cycle, true) //если подходящий автоматизм найден в успешной Доминанте то он будет запущен
			if res {                               // найден и запущен
				return
			}
		}
		if automatizmStatus == 2 { //нет автоматизма или есть старый автоматизм, например, при игнорировании Ответа оператором
			// запускать решение доминанты в подходящих условиях т.е. взять опыт из ранее решенных проблм
			// не только закрытая доминанта, но и точная по условиям, но не проверенная.
			res := runDominantaAction(cycle, false) //если подходящий автоматизм найден в успешной Доминанте то он будет запущен
			if res {                                // найден и запущен
				return
			}
		}
		if automatizmStatus > 0 { // проблема автоматизма не решена
			//automatizmStatus = 0
			if infoFunc2(cycle) {
				return
			}
			//return
		}
	}

	/////////////////////////////////////////////////////////////

	//////////////////////////////////////////////
	// только если ВАЖНАЯ ИЛИ ОПАСНАЯ СИТУАЦИЯ  не в дреме Блок повышенной акутальности
	if motorActionEffect == 0 && !cycle.dreaming && (CurrentInformationEnvironment.veryActualSituation ||
		CurrentInformationEnvironment.danger) {
		//счетчик нерешенных проблем и открытие доминанты createDominanta(cycle)
		if dangerActualProcess(cycle) {
			cycle.idle = false
			return
		}
	}
	///////////////////////////////////////////////////////////

	//////////////////////// ВАЖНЫЙ ОБЪЕКТ ВНИМАНИЯ
	if extremImportanceObject != nil { // подумать об этом и только об этом
		if EvolushnStage > 4 {
			getCurrentDominant() // по текущему extremImportanceObject
		}
		/* выявляется проблемный объект с отрицательным эффектом - основа мук творчества
		problemExtremImportanceObject = extremImportanceObject
		и делаются первые шаги решения
		*/
		if extremImportanceObjectProcess(cycle) {
			cycle.idle = false
			return
		}
	}
	/////////////////////////////////////////////////////////////

	////////////
	/* ОБЫЧНАЯ СИТУАЦИЯ, О ЧЕМ ДУМАТЬ продолжение шага мышления
	уже запускалось infoFunc8 выше

	if usualThinkProcess(cycle) {
		cycle.idle = false
		return
	}*/
	/////////////////////////////////////////////////////////////

	// чтобы срабатывало не только в func infoFunc2 - неудержимый порыв общения с провокацией
	if EvolushnStage > 3 && mentalInfoStruct.noOperatorStimul {
		infoFunc31pulsCount = 0       // чтобы снова проверить условие isNeedForCommunication()
		if isNeedForCommunication() { // нужно провоцировать оператора
			infoFunc31(cycle)
			return
		}

	}

	//////////////////////////////// мышление о Доминанте
	if EvolushnStage > 4 { //5-я сталия развития. Творчество.
		if dominantsProcess(cycle) { // TODO еще не закончена ф-ция!

			// подавление мешающих стимулов при серьезном поиске решений проблемы
			repressionStimulsNoises() // устанавливает isRepressionStimulsNoise=true

			return
		}
	} //if EvolushnStage > 4
	/////////////////////////////////////////////////////////

	if cycle.idle == false {
		cycle.log += "Цикл - в ХОЛОСТОМ ожидании.<br>"

		//Не было действий Beast  более 100 пульсов, возможно есть условия проявить инициативу
		if CurrentInformationEnvironment.IsIdleness100pulse {
			infoFunc25(cycle) //Посмотреть условия чтобы проявить инициативу
		}
	}
	cycle.idle = true
}

///////////////////////////////////////////////////////////////////////////////////////////

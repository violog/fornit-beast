/*
индексная страница http://localhost:8181/index

В проекте много глобальных переменных, что привычно раздражает Свидетелей инкапсуляции и непорочного пространства имен,
но ТАК НУЖНО (спорно) для одганизации среды,
схожей с организацией линкующих указателей в мозге (т.е. связей с одного распознавателя к целому ансаблю - объекту).
Ну и есть немало других вещей, нарушающих Порядок и Традиции Golang.
Попытки использовать горутины оказались просто неуместными (спорно) и просто ненужными, учитывая вряд ли в чем-то могущий быть выигрыш.
Короче, код предоставляется на вольное растерзание и свободное экспериментирование, без претензий, сорри за возможный негатив.
Везде много пространных комментариев, которые запутывают даже меня, но они НУЖНЫ.
*/

package main

import (
	"BOT/brain"
	"BOT/brain/action_sensor"
	"BOT/brain/gomeostas"
	"BOT/brain/psychic"
	"BOT/brain/reflexes"
	"BOT/brain/sleep"
	"BOT/brain/transfer"
	"BOT/brain/update"
	"BOT/brain/words_sensor"
	"BOT/lib"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var xxxxxxx = 0 // в дебаге иногда начинается циклический вызов, это - защелка

// true - остановка всей активности для совершения критических глобальных операций
var isGlobalStopAllActivnost = false

var noСloseWeght = false

func receiveSend(resp http.ResponseWriter, r *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// посылается постоянно раз в 1сек (из /common/linking.php) или с запросом или для подтверждения связи,
	// а так же для передачи по текущему пульсу информации от Beast, например WritePultConsol()
	if r.Method == "POST" {
		if gomeostas.CheckBeastDeath() {
			isGlobalStopAllActivnost = true
			_, _ = fmt.Fprint(resp, "!!!")
		}

		if !isGlobalStopAllActivnost {
			// текстовый блок для набивки дерева слов-фраз из http://go/pages/words.php
			raw_text_block := r.FormValue("raw_text_block")
			if len(raw_text_block) > 0 {
				brain.IsPultActivnost = true
				res := word_sensor.SetNewTextBlock(raw_text_block)
				brain.IsPultActivnost = false
				_, _ = fmt.Fprint(resp, res)
				return
			}
			// текст из окна ввода с пульта
			text_split := r.FormValue("text_split")
			if len(text_split) > 0 {
				brain.IsPultActivnost = true
				is_forced_mode, _ := strconv.Atoi(r.FormValue("is_forced_mode")) // режим быстрого формирования у-рефлексов
				if is_forced_mode == 0 {                                         // наоборот
					reflexes.IsUnlimitedMode = 1
				} else {
					reflexes.IsUnlimitedMode = 0
				}
				toneID, _ := strconv.Atoi(r.FormValue("control_tone"))
				pultMood := r.FormValue("control_mood") // тон сообщения
				moodID, _ := strconv.Atoi(pultMood)     // настроение сообщения
				res := word_sensor.VerbalDetection(text_split, is_forced_mode, toneID, moodID)
				if moodID > 0 {
					// учесть мотивирующий эффект
					action_sensor.UpdateMoodEffectFromMessage(moodID)
				}
				// если добавлены пусковые стимулы (нажаты кнопки на пульте)
				actions_image := r.FormValue("actions_image")
				action_sensor.IsPultAction()
				if len(actions_image) > 0 {
					// brain.IsPultActivnost = true

					enegry, _ := strconv.Atoi(r.FormValue("food_portion"))
					action_sensor.SetActionFromPult(actions_image, enegry)
					/*
						// активировать дерево действием
						reflexes.ActiveFromAction()
						brain.IsPultActivnost = false
					*/
				}

				reflexes.ActiveFromPhrase() // активировать дерево рефлексов фразой - только для условных рефлексов
				brain.IsPultActivnost = false
				// и если есть ответ, то тут же выдать на Пульт
				var answerStr = ""
				if len(lib.ActionsForPultStr) > 5 && !sleep.IsSleeping {
					lib.ActionsForPultStr = lib.SharedReflexWithAutomatism()
					if len(lib.ActionsForPultStr) > 5 {
						if false { // проверка акций
							arts := strings.Split(lib.ActionsForPultStr, "||")
							for i := 0; i < len(arts); i++ {
								art := strings.Split(arts[i], "|")
								if len(art) != 2 {
									return
								}
							}
						}
					}
					//могут быть несколько одинаковых сообщений с разной энергичностью. Для одинаковых действий ставится максимальная Энергичность из имеющихся.
					answerStr = lib.DoublesRemove(lib.ActionsForPultStr)

					lib.ActionsForPultStr = "" // очистка для новой порции
				}
				_, _ = fmt.Fprint(resp, res+"|&|"+answerStr)
				return
			}

			// отправить на пульт состояние гомеостаза Beast и его базовые контексты
			// not present in PHP = never used
			sincronism := r.FormValue("sincronism")
			if len(sincronism) > 0 {
				// выполнить цикл действий по пульсу перед отправкой результата на Пульт
				brain.SincroTic()
				_, _ = fmt.Fprint(resp, "sincronism")
				return
			}

			// отправить на пульт состояние гомеостаза Beast и его базовые контексты
			getParams := r.FormValue("get_params")
			if len(getParams) > 0 {
				brain.IsPultActivnost = true
				outStr := gomeostas.GetCurGomeoParams()

				var waitingPeriod = "_"
				// был Стиму оператора, но 2 пульса как нет ответного автоматизма
				if psychic.NoautomatismAfterStimul == 2 {
					// активация дерева не вызвала автоматизм и не было периода ожидания
					waitingPeriod = "zzz"
					/*	psychic.NoautomatismAfterStimul=0 // не повторять, а дать сбросить сообщение
						потому как все сложно из-за синхронизации пульсов (js:sincronisationWithGo()),
						котоая тоже вызывает "actions_image"
						Cброс NoautomatismAfterStimul происходит по js:endNoautomatismAfterStimul()
						после отработки сообщения в GET обработчике: r.FormValue("endNoautomatismAfterStimul")
					*/
				}
				res, timeWait := psychic.WaitingPeriodForActions()
				if res {
					waitingPeriod = strconv.Itoa(timeWait)
				}
				var gaimMode = "0"
				if transfer.IsPsychicGameMode {
					gaimMode = "1"
				}
				var teachingMode = "0"
				if transfer.IsPsychicТeachingMode {
					teachingMode = "1"
				}
				p1 := gomeostas.GetCurGomeoStatus()
				p2 := gomeostas.GetCurContextActive()
				p3 := reflexes.GetCurrentConditionsStr()
				p4 := strconv.Itoa(psychic.CurrentMoodCondition)

				outStr += "#%#" + p1 + "#%#" + p2 +
					"#%# " + p3 + // чтобы постоянно была инфа о сочетаниях контекстов
					"#%#" + strconv.Itoa(brain.LifeTime) +
					"#%#" + reflexes.NoUnconditionRefles +
					"#%#" + waitingPeriod +
					"#%#" + psychic.GetPsichicReady() +
					"#%#" + p4 +
					"#%#" + gaimMode +
					"#%#" + psychic.GetExtandIndoForPult() +
					"#%#" + teachingMode

				brain.IsPultActivnost = false

				if strings.Contains(outStr, "zzz") { // для точки прерывания
					_, _ = fmt.Fprint(resp, outStr)
					return
				}

				_, _ = fmt.Fprint(resp, outStr)
				return
			}
			//////////////////////////////////////////////

			// установка параметров гомеостаза с Пульта:
			// задать параметры гомеостаза Beast
			setParamsId := r.FormValue("set_params")
			if len(setParamsId) > 0 {
				brain.IsPultActivnost = true
				id, _ := strconv.Atoi(setParamsId)
				gomeostas.SetCurGomeoParams(id, r.FormValue("params_list"))
				brain.IsPultActivnost = false
				_, _ = fmt.Fprint(resp, "1")
				return
			}

			//  передача информации от Beast в Пульт различается идентификатором назначения информации перед самой информацией
			// информация для консоли ничинается с идентификатора назначения: "CONSOL:"
			if len(lib.WritePultConsolStr) > 0 {
				_, _ = fmt.Fprint(resp, "CONSOL:"+lib.WritePultConsolStr)
				lib.WritePultConsolStr = "" // очистка для новой порции
				return
			}

			/* Если есть непустой ActionsForPultStr, то выдать на Пульт действия
			информация о действиях Beast ничинается с идентификатора назначения: "ACTION:"
			*/
			if len(lib.ActionsForPultStr) > 0 {
				sctionsStr := lib.SharedReflexWithAutomatism()
				if len(sctionsStr) > 0 {
					_, _ = fmt.Fprint(resp, "ACTION:"+sctionsStr)
				}
				lib.ActionsForPultStr = "" // очистка для новой порции
				return
			}

			// если ничего выше не было, то:
			// передача на Пульт сигнала готовности - когда нет других запросов - посылаетс сигнал на Пульт в function bot_answer(res)
			if word_sensor.IsReadyWordSensorLevel() {
				//идентификатор назначения информации: "READY"
				_, _ = fmt.Fprint(resp, "READY")
				return
			}
			_, _ = fmt.Fprint(resp, "POST")
		} else {
			brain.NotAllowAnyActions = true

			// Сформировать условные рефлексы на основе списка фраз-синонимов
			file_for_cond_reflexes := r.FormValue("file_for_cond_reflexes")
			if len(file_for_cond_reflexes) > 0 {
				reflexes.FormingConditionsRefleaxFromList(file_for_cond_reflexes)
				_, _ = fmt.Fprint(resp, "OK")
			}
		}
		//fmt.Println("EMPTY")
	}

	if r.Method == "GET" {
		// проверка активности Beast, аозвращает текущий brain.PulsCount
		brain.IsPultActivnost = true
		check_activity := r.FormValue("check_activity")
		if check_activity == "1" {
			_, _ = fmt.Fprint(resp, brain.PulsCount)
			return
		}
		// остановка любой активности Beast
		brain.IsPultActivnost = true
		stop_activity := r.FormValue("stop_activity")
		if stop_activity == "1" {
			isGlobalStopAllActivnost = true
			_, _ = fmt.Fprint(resp, "stop")
			return
		}
		// восстановление активности Beast
		start_activity := r.FormValue("start_activity")
		if start_activity == "1" {
			isGlobalStopAllActivnost = false
			brain.IsPultActivnost = false
			_, _ = fmt.Fprint(resp, "active")
			return
		}
		// ЗОНА ОСОБЫХ ДЕЙСТВИЙ в период остановленной активности Beast:

		// Сохранить текущее состояние Beast
		save_all_memory := r.FormValue("save_all_memory")
		if save_all_memory == "1" {
			brain.IsPultActivnost = true
			if brain.SaveAll() {
				_, _ = fmt.Fprint(resp, "yes")
				brain.IsPultActivnost = false
				return
			}
			_, _ = fmt.Fprint(resp, "no")
			brain.IsPultActivnost = false
			return
		}
		// корректное выключение Beast
		bot_shutdown := r.FormValue("bot_shutdown")
		if len(bot_shutdown) > 0 {
			brain.IsPultActivnost = true
			cleanupFunc(bot_shutdown)
			return
		}
		// выключение по крестику окна браузера
		beforeunload := r.FormValue("beforeunload")
		if len(beforeunload) > 0 {
			testSTR := r.FormValue("testSTR")
			if len(testSTR) > 0 {
				// просто для тестирования
			}
			if noСloseWeght {
				return
			}
			brain.IsPultActivnost = true
			cleanupFunc("1") // "1" - сохранить все
			return
		}
		// на время перезагрузки НЕ закрывать прогу
		do_not_close_on_restart := r.FormValue("do_not_close_on_restart")
		if len(do_not_close_on_restart) > 0 {
			noСloseWeght = true
			_, _ = fmt.Fprint(resp, "!")
			//time.Sleep(2 * time.Second) // ждем 2 сек
			time.AfterFunc(2*time.Second, func() {
				noСloseWeght = false
			})
			return
		}

		// Формирование зеркальных автоматизмов на основе списка ответов
		mirror_making_fool := r.FormValue("mirror_making_fool")
		if len(mirror_making_fool) > 0 {
			res := psychic.FormingMirrorAutomatismFromList(mirror_making_fool)
			_, _ = fmt.Fprint(resp, res)
			return
		}
		// Формирование зеркальных автоматизмов на основе общего шаблона
		mirror_making_temp := r.FormValue("mirror_making_temp")
		if len(mirror_making_temp) > 0 {
			res := psychic.FormingMirrorAutomatismFromTempList(mirror_making_temp)
			_, _ = fmt.Fprint(resp, res)
			return
		}

		if !isGlobalStopAllActivnost {
			setExpParam := r.FormValue("export_data") // экспорт данных
			if len(setExpParam) > 0 {
				brain.IsPultActivnost = true
				if setExpParam == "1" {
					IsExpTrue, expTxt := update.ExportFileUpdate([]int{1, 2, 3, 4, 5})
					if IsExpTrue == true {
						setExpParam = "yes|" + expTxt
					} else {
						setExpParam = "no|" + expTxt
					}
				}
				brain.IsPultActivnost = false
				_, _ = fmt.Fprint(resp, setExpParam)
				return
			}

			setImpParam := r.FormValue("import_data") // импорт данных
			if len(setImpParam) > 0 {
				brain.IsPultActivnost = true
				if setImpParam == "1" {
					IsImpParam, impTxt := update.ImportFileUpdate([]int{1, 2, 3, 4, 5})
					if IsImpParam == true {
						setImpParam = "yes|" + impTxt
					} else {
						setImpParam = "no|" + impTxt
					}
				}
				brain.IsPultActivnost = false
				_, _ = fmt.Fprint(resp, setImpParam)
				return
			}

			get_reflex_tree := r.FormValue("get_reflex_tree")
			if get_reflex_tree == "1" {
				brain.IsPultActivnost = true
				tree := reflexes.GetReflexesTreeForPult()
				brain.IsPultActivnost = false
				if tree == "!!!" {
					return // запрет показа карты во время распознавания и записи
				}
				_, _ = fmt.Fprint(resp, tree)
				return
			}

			get_phrase_tree := r.FormValue("get_phrase_tree")
			if get_phrase_tree == "1" {
				brain.IsPultActivnost = true
				phTree := word_sensor.GetPhraseTreeForPult()
				brain.IsPultActivnost = false
				if phTree == "!!!" {
					return // запрет показа карты во время распознавания и записи
				}
				_, _ = fmt.Fprint(resp, phTree)
				return
			}

			get_words_list := r.FormValue("get_words_list")
			if get_words_list == "1" {
				brain.IsPultActivnost = true
				phTree := word_sensor.GetWordsListForPult()
				brain.IsPultActivnost = false
				if phTree == "!!!" {
					return // запрет показа карты во время распознавания и записи
				}
				_, _ = fmt.Fprint(resp, phTree)
				return
			}

			deleting_word := r.FormValue("deleting_word")
			if deleting_word == "1" {
				delete_word := r.FormValue("delete_word")
				deleteWord, _ := strconv.Atoi(delete_word)
				brain.IsPultActivnost = true
				word_sensor.DeleteWord(deleteWord)
				brain.IsPultActivnost = false
				_, _ = fmt.Fprint(resp, "OK")
				return
			}

			get_word_tree := r.FormValue("get_word_tree")
			if get_word_tree == "1" {
				brain.IsPultActivnost = true
				phTree := word_sensor.GetWordTreeForPult()
				brain.IsPultActivnost = false
				if phTree == "!!!" {
					return // запрет показа карты во время распознавания и записи
				}
				_, _ = fmt.Fprint(resp, phTree)
				return
			}

			//
			/*xxxxxxx - заглушка от множестенных вызовов (в том числе при возврате по стеку от паники)


			 */
			set_action := r.FormValue("set_action")
			if len(set_action) > 0 && xxxxxxx == 0 {
				xxxxxxx = 1
				// нажимать кнопки не чаще раз в 1 секунду
				time.AfterFunc(1*time.Second, func() {
					xxxxxxx = 0
				})

				brain.IsPultActivnost = true
				//ActionsForPultOldStr = lib.ActionsForPultOldStr
				action_sensor.IsPultAction()
				enegry, _ := strconv.Atoi(r.FormValue("food_portion"))
				// здесь включение игрового режима IsPsychicGameMode при нажатии на Поиграть
				action_sensor.SetActionFromPult(set_action, enegry)

				// активировать дерево действием
				reflexes.ActiveFromAction()
				brain.IsPultActivnost = false
				// и если есть ответ, то тут же выдать на ПультL
				var answerStr = ""
				if len(lib.ActionsForPultStr) > 5 {
					lib.ActionsForPultStr = lib.SharedReflexWithAutomatism()
					if len(lib.ActionsForPultStr) > 5 {
						answerStr = lib.ActionsForPultStr
					}
					lib.ActionsForPultStr = "" // очистка для новой порции
				}
				_, _ = fmt.Fprint(resp, answerStr)
				return
			}

			// установка-снятие игрового режима в pult_base_contexts.php по function end_game_mode()
			is_game_mode := r.FormValue("is_game_mode")
			if len(is_game_mode) > 0 {
				if is_game_mode == "1" {
					transfer.IsPsychicGameMode = true // переключатель игрового режима
				} else {
					transfer.IsPsychicGameMode = false // переключатель игрового режима
				}
				_, _ = fmt.Fprint(resp, "ok")
				return
			}

			// установка-снятие учительского режима в pult_base_contexts.php по function end_teaching_mode()
			is_teaching_mode := r.FormValue("is_teaching_mode")
			if len(is_teaching_mode) > 0 {
				if is_teaching_mode == "1" {
					transfer.IsPsychicТeachingMode = true // переключатель учительского режима
				} else {
					transfer.IsPsychicТeachingMode = false // переключатель учительского режима
				}
				_, _ = fmt.Fprint(resp, "ok")
				return
			}

			// снятие состояния Хорошо и Плохо
			clear_well_bad := r.FormValue("clear_well_bad")
			if len(clear_well_bad) > 0 {
				// Снять режимы Плохо и Хорошо
				gomeostas.CliarWellBad()
				_, _ = fmt.Fprint(resp, "ok")
				return
			}

			/* // проверка игрового режима
			check_game_mode := r.FormValue("check_game_mode")
			if len(check_game_mode) > 0 {
				if transfer.IsPsychicGameMode{
					_, _ = fmt.Fprint(resp, "1")
				}else{
					_, _ = fmt.Fprint(resp, "0")
				}
				return
			}*/

			clear_cond_reflex_timer := r.FormValue("clear_cond_reflex_timer")
			if len(clear_cond_reflex_timer) > 0 {
				isGlobalStopAllActivnost = true
				ret := reflexes.ClinerTimeConditionReflex()
				isGlobalStopAllActivnost = false
				_, _ = fmt.Fprint(resp, ret)
				return
			}

			get_condition_reflex_info := r.FormValue("get_condition_reflex_info")
			if len(get_condition_reflex_info) > 0 {
				base := r.FormValue("limitBasicID")
				limitBasicID, _ := strconv.Atoi(base)
				ref := reflexes.GetConditionReflexInfo(limitBasicID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_automatism_list_info := r.FormValue("get_automatism_list_info")
			if len(get_automatism_list_info) > 0 {
				lpage := r.FormValue("limitBasicID")
				limitPage, _ := strconv.Atoi(lpage)
				ref := psychic.GetAutomatismInfo(limitPage)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			unblock_all_automatisms := r.FormValue("unblock_all_automatisms")
			if len(unblock_all_automatisms) > 0 {
				psychic.UnblockingAllAtmtzms()
				_, _ = fmt.Fprint(resp, "ok")
				return
			}

			get_next_actions := r.FormValue("get_next_actions")
			if len(get_next_actions) > 0 {
				ref := psychic.GetNextActionsInfoList()
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_trigger_info := r.FormValue("get_trigger_info")
			if len(get_trigger_info) > 0 {
				triggerID, _ := strconv.Atoi(r.FormValue("triggerID"))
				ref := reflexes.GetTreeAutomatismTriggersInfo(triggerID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_sequence_info := r.FormValue("get_sequence_info")
			if len(get_sequence_info) > 0 {
				autmzmID, _ := strconv.Atoi(r.FormValue("autmzmID"))
				ref := psychic.GetAutomatismSequenceInfo(autmzmID, false)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_show_problem_tree_info := r.FormValue("get_show_problem_tree_info")
			if len(get_show_problem_tree_info) > 0 {
				pID, _ := strconv.Atoi(r.FormValue("mautmzmID"))
				ref := psychic.ShowProblemTreeInfo(pID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_emotion_info := r.FormValue("get_emotion_info")
			if len(get_emotion_info) > 0 {
				emotionID, _ := strconv.Atoi(r.FormValue("emotionID"))
				ref := psychic.GetStrnameFromBaseImageID(emotionID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_next_actions_info := r.FormValue("get_next_actions_info")
			if len(get_next_actions_info) > 0 {
				nextID, _ := strconv.Atoi(r.FormValue("nextID"))
				ref := psychic.GetNextActionsInfo(nextID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_object_info := r.FormValue("get_object_info")
			if len(get_object_info) > 0 {
				objectID, _ := strconv.Atoi(r.FormValue("objectID"))
				ref := psychic.GetStrnameFromobjectID(objectID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			clear_atmtzm_block := r.FormValue("clear_atmtzm_block")
			if len(clear_atmtzm_block) > 0 {
				atmtzmID, _ := strconv.Atoi(r.FormValue("atmtzmID"))
				ref := psychic.UnblockAutomatismID(atmtzmID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_dominant_info := r.FormValue("get_dominant_info")
			if len(get_dominant_info) > 0 {
				dID, _ := strconv.Atoi(r.FormValue("atmtzmID"))
				ref := psychic.DominantaInfoStr(dID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_self_perception_info := r.FormValue("get_self_perception_info")
			if len(get_self_perception_info) > 0 {
				ref := psychic.GetSelfPerceptionInfo()
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_self_conscience_info := r.FormValue("get_self_conscience_info")
			if len(get_self_conscience_info) > 0 {
				ref := psychic.GetConscienceInfo()
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_cycle_log_info := r.FormValue("get_cycle_log_info")
			if len(get_cycle_log_info) > 0 {
				cID, _ := strconv.Atoi(get_cycle_log_info)
				ref := psychic.GetCycleLocInfo(cID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_automatism_tree := r.FormValue("get_automatism_tree")
			if len(get_automatism_tree) > 0 {
				base := r.FormValue("limitBasicID")
				limitBasicID, _ := strconv.Atoi(base)
				ref := psychic.GetAutomatismTreeForPult(limitBasicID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}
			get_node_automatisms := r.FormValue("get_node_automatisms")
			if len(get_node_automatisms) > 0 {
				nodeID, _ := strconv.Atoi(r.FormValue("autNodeID"))
				ref := psychic.GetNodesAutomatismsInfo(nodeID) //GetAutomatismForNodeInfo(nodeID)- пишет в лог
				_, _ = fmt.Fprint(resp, ref)
				return
			}
			get_automatism := r.FormValue("get_automatism")
			if len(get_automatism) > 0 {
				ID, _ := strconv.Atoi(r.FormValue("autID"))
				atmz, ok := psychic.ReadeAutomatismFromId(ID)
				if ok {
					ref := psychic.GetAutomotizmActionsString(atmz, false, true)
					_, _ = fmt.Fprint(resp, ref)
				}
				return
			}

			get_problem_tree_node := r.FormValue("get_problem_tree_node")
			if len(get_problem_tree_node) > 0 {
				nodeID, _ := strconv.Atoi(r.FormValue("autNodeID"))
				ref := psychic.GetProblemTreeForNodeInfo(nodeID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_node_aut_tree_node := r.FormValue("get_node_aut_tree_node")
			if len(get_node_aut_tree_node) > 0 {
				nodeID, _ := strconv.Atoi(r.FormValue("autNodeID"))
				ref := psychic.GetAutomatismNodeTreeForPult(nodeID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_mental_automatism_tree := r.FormValue("get_mental_automatism_tree")
			if len(get_mental_automatism_tree) > 0 {
				base := r.FormValue("limit")
				limit, _ := strconv.Atoi(base)
				ref := psychic.GetMentalAutomatismTreeForPult(limit)
				_, _ = fmt.Fprint(resp, ref)
				return
			}
			get_node_mental_situations := r.FormValue("get_node_mental_situations")
			if len(get_node_mental_situations) > 0 {
				nodeID, _ := strconv.Atoi(r.FormValue("autNodeID"))
				ref := psychic.GetMentalSituationsForNodeInfo(nodeID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_node_mental_automatism := r.FormValue("get_node_mental_automatism")
			if len(get_node_mental_automatism) > 0 {
				//maID,_:=strconv.Atoi(r.FormValue("autNodeID"))
				ref := psychic.GetMentalAutomatismForPult()
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_mental_model := r.FormValue("get_mental_model")
			if len(get_mental_model) > 0 {
				maID, _ := strconv.Atoi(r.FormValue("autNodeID"))
				ref := psychic.GetModelExtremImportanceInfo(maID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_node_mental_purpose := r.FormValue("get_node_mental_purpose")
			if len(get_node_mental_purpose) > 0 {
				nodeID, _ := strconv.Atoi(r.FormValue("autNodeID"))
				ref := psychic.GetMentalPurposeForNodeInfo(nodeID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_node_theme_image := r.FormValue("get_node_theme_image")
			if len(get_node_theme_image) > 0 {
				nodeID, _ := strconv.Atoi(r.FormValue("autNodeID"))
				ref := psychic.GetMentalThemeForNodeInfo(nodeID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_mental_problem_tree := r.FormValue("get_mental_problem_tree")
			if len(get_mental_problem_tree) > 0 {
				base := r.FormValue("limit")
				limit, _ := strconv.Atoi(base)
				ref := psychic.GetMentalPriblemTreeForPult(limit)
				_, _ = fmt.Fprint(resp, ref)
				return
			}
			/*
				get_mental_understanding_models_info := r.FormValue("get_mental_understanding_models_info")
				if len(get_mental_understanding_models_info) > 0 {
					ref := psychic.GetMentalUndestandingModelsForPult()
					_, _ = fmt.Fprint(resp, ref)
					return
				}
			*/
			get_rules_list_info := r.FormValue("get_rules_list_info")
			if get_rules_list_info == "1" {
				rules := psychic.GetCur100lastRules(0)
				_, _ = fmt.Fprint(resp, rules)
				return
			}

			get_understanding_model := r.FormValue("get_understanding_model")
			if get_understanding_model == "1" {
				ids := r.FormValue("objID")
				id, _ := strconv.Atoi(ids)
				model := psychic.Get_understand_model_from_object(id)
				_, _ = fmt.Fprint(resp, model)
				return
			}

			clear_episodic_memory := r.FormValue("clear_episodic_memory")
			if clear_episodic_memory == "1" {
				psychic.ClianEpisodicMemory()
				_, _ = fmt.Fprint(resp, "did")
				return
			}

			clian_ment_episodic_memory := r.FormValue("clian_ment_episodic_memory")
			if clian_ment_episodic_memory == "1" {
				psychic.ClianMentalEpisodicMemory()
				_, _ = fmt.Fprint(resp, "did")
				return
			}

			get_mental_rules_list_info := r.FormValue("get_mental_rules_list_info")
			if get_mental_rules_list_info == "1" {
				rules := psychic.GetMentallastRules()
				_, _ = fmt.Fprint(resp, rules)
				return
			}

			get_mental_action_info := r.FormValue("get_mental_action_info")
			if len(get_mental_action_info) > 0 {
				//actID,_:=strconv.Atoi(get_mental_action_info)
				rules := "НЕТ ПОДДЕРЖКИ psychic.GetMentalActionInfo - устарело." //psychic.GetMentalActionInfo(actID)
				_, _ = fmt.Fprint(resp, rules)
				return
			}

			get_mental_goNext_info := r.FormValue("get_mental_goNext_info")
			if len(get_mental_goNext_info) > 0 {
				//actID,_:=strconv.Atoi(get_mental_goNext_info)
				rules := "НЕТ ПОДДЕРЖКИ psychic.GetMentalActionsString - устарело." //psychic.GetMentalActionsString(actID)
				_, _ = fmt.Fprint(resp, rules)
				return
			}

			// инфа о крутящихся активных циклах
			get_cycle_info := r.FormValue("get_cycle_info")
			if len(get_cycle_info) > 0 {
				cID, _ := strconv.Atoi(get_cycle_info)
				rules := psychic.GetCycleStrInfo(cID)
				_, _ = fmt.Fprint(resp, rules)
				return
			}

			// инфо о инфо-функции
			get_mental_func_info := r.FormValue("get_mental_func_info")
			if len(get_mental_func_info) > 0 {
				fID, _ := strconv.Atoi(get_mental_func_info)
				rules := psychic.GetgetInfoFuncInfoStr(fID)
				_, _ = fmt.Fprint(resp, rules)
				return
			}

			get_automatism_tree_info := r.FormValue("get_automatism_tree_info")
			if get_automatism_tree_info == "1" {
				objID, _ := strconv.Atoi(r.FormValue("objID"))
				ref := psychic.GetAtmzmTreeInfo(objID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}
			get_understanding_tree_info := r.FormValue("get_understanding_tree_info")
			if get_understanding_tree_info == "1" {
				objID, _ := strconv.Atoi(r.FormValue("objID"))
				ref := psychic.GetUndstgTreeInfo(objID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}
			get_mental_automatism_info := r.FormValue("get_mental_automatism_info")
			if get_mental_automatism_info == "1" {
				objID, _ := strconv.Atoi(r.FormValue("objID"))
				ref := psychic.GetMentAtmzmInfo(objID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			make_automatisms_from_reflexes := r.FormValue("make_automatisms_from_reflexes")
			if len(make_automatisms_from_reflexes) == 1 {
				ref := reflexes.RunMakeAutomatismsFromReflexes()
				_, _ = fmt.Fprint(resp, ref)
				return
			}
			make_automatisms_from_genetic_reflexes := r.FormValue("make_automatisms_from_genetic_reflexes")
			if len(make_automatisms_from_genetic_reflexes) == 1 {
				ref := reflexes.RunMakeAutomatismsFromGeneticReflexes()
				_, _ = fmt.Fprint(resp, ref)
				return
			}
			// объект значимости
			get_mental_importance_list_info := r.FormValue("get_mental_importance_list_info")
			if get_mental_importance_list_info == "1" {
				ref := psychic.GetImportanceToPult()
				_, _ = fmt.Fprint(resp, ref)
				return
			}
			get_show_cycle_info := r.FormValue("get_show_cycle_info")
			if get_show_cycle_info == "1" {
				cID, _ := strconv.Atoi(r.FormValue("objID"))
				ref := psychic.GetCycleInfo(cID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_importance_object_info := r.FormValue("get_importance_object_info")
			if get_importance_object_info == "1" {
				objID, _ := strconv.Atoi(r.FormValue("objID"))
				//objType, _ := strconv.Atoi(r.FormValue("objType"))
				ref := psychic.GetImportanceObjectInfo(objID)
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			get_dominants_list_info := r.FormValue("get_dominants_list_info")
			if get_dominants_list_info == "1" {
				ref := psychic.GetDominantsListToPult()
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			//получить фразы, используемые в автоматизмах для иконки выбора Пульта
			conditions_words_basic := r.FormValue("conditions_words_basic")
			if conditions_words_basic == "1" {
				// против concurrent map iteration and map write
				brain.IsPultActivnost = true
				bID := r.FormValue("basicID")
				bID = strings.Trim(bID, " ")
				basicID, _ := strconv.Atoi(bID)
				contexts := r.FormValue("contexts")
				ref := psychic.GetAutomatismPraseList(basicID, contexts)
				brain.IsPultActivnost = false
				_, _ = fmt.Fprint(resp, ref)
				return
			}

			// обнулить параметры гомеостаза Beast
			clear_gomeo_pars := r.FormValue("clear_gomeo_pars")
			if len(clear_gomeo_pars) > 0 {
				brain.IsPultActivnost = true
				value, _ := strconv.ParseFloat(clear_gomeo_pars, 64)
				gomeostas.ClinerAllGomeoParams(value)
				brain.IsPultActivnost = false
				_, _ = fmt.Fprint(resp, "1")
				return
			}

			// завершить период ожидания по щелчку на плашке
			stop_waiting_period := r.FormValue("stop_waiting_period")
			if len(stop_waiting_period) > 0 {
				psychic.StopWaitingWeriodFromOperator() // завершить цепочку темы
				_, _ = fmt.Fprint(resp, "1")
				return
			}

			end_no_automatism := r.FormValue("end_no_automatism")
			if len(end_no_automatism) > 0 {
				psychic.NoautomatismAfterStimul = 0
				_, _ = fmt.Fprint(resp, "ok_end")
				return
			}

			_, _ = fmt.Fprint(resp, "GET")
		}
	}

}

// инициализация
func init() {
	lib.GetMainPathExeFile()
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// тестирование комбинаций. Если бы время работы было приемлемо,
	//то можно было бы запускать процесс из Пульта в меню Инструменты (шестеренка)
	// 	tools.MakeContextCombinations()
	/*
		res:=lib.DoublesRemove("3|Действие: <b><span style=\"font-size:14px;\">кричит</span>, <span style=\"font-size:14px;\">дерется</span>, <span style=\"font-size:14px;\">опасается</span></b><br><span style=\"font-size:15px;color:#000000\">Энергичность: <b>Средне (сила=8)</b></span><br>||3|Действие: <b><span style=\"font-size:14px;\">ОРЕТ</span>, <span style=\"font-size:14px;\">КУСАЕТСЯ</span>, <span style=\"font-size:14px;\">опасается</span></b><br><span style=\"font-size:15px;color:#000000\">Энергичность: <b>Средне (сила=6)</b></span><br>||3|Действие: <b><span style=\"font-size:14px;\">кричит</span>, <span style=\"font-size:14px;\">дерется</span>, <span style=\"font-size:14px;\">опасается</span></b><br><span style=\"font-size:15px;color:#000000\">Энергичность: <b>Средне (сила=6)</b></span><br>")
		if len(res)>0{}
	*/
}

func main() {

	address := lib.ReadFileContent(lib.GetMainPathExeFile() + "/common/linking_address.txt")
	address = strings.TrimSpace(address)[7:]

	brain.RunInitialisation() // init.go
	brain.Puls()

	http.HandleFunc("/", receiveSend)
	_ = http.ListenAndServe(address, nil)
	//	fmt.Println("Сервер запущен...")

}

// отключение Beast - по запросу со страницы Пульта
func cleanupFunc(typeClosing string) {
	lib.WritePultConsol("Beast вырубается.")
	fmt.Print("ПОСЛЕДНИЕ ДЕЙСТВИЯ ПЕРЕД ЗАКРЫВАНИЕМ ПРОГРАММЫ")

	if typeClosing == "1" {
		// записать текущее состояние Дерева Моделей и Эпизодическую память
		//	brain.PrepareBeforCloseTreeModel()
		/* если внезапно откобчить мозг человека, то из памяти пропадет все, что было в последние полчаса
		так что просто записывать PrepareBeforCloseTreeModel() раз в 10 минут и при создании нового узла дерева
		*/
		brain.SaveAll()
	}

	os.Exit(1)
}

// здесь могут быть функции для обеспечения связи между пакетами чтобы избегать цикличного импорта

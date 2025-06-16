/*  обработка во сне

Во время сновидений мы можем видеть и слышать вещи,
которые не существуют в реальном мире, испытывать эмоции и переживать события,
которые могут быть несвязанными с нашей повседневной жизнью.

Наиболее известным примером был его сон о гильотине: длинный сон о французской революции,
кульминацией которого был приговор Мори к смертной казни. Он проснулся в момент падения ножа гильотины,
обнаружив, что доска с изголовья кровати упала ему на шею.
Внешний раздражитель запустил последовательность воспоминаний, которые образовали сновидения.

Освобождать неиспользуемую память: v=nil

в understanding.go есть func WakeUpping()
*/

package psychic

/*
	можно ли заснуть и спать - вызывается из sleep.SleepPuls()

Если хочется спать
*/
func IsPossibleToSleep() bool {
	/*
		//Если уже сон:  НЕ ПРОСЫПАТЬСЯ ОТ ПСИХ СТИМУЛОВ!
		if IsSleeping {
			// нужно проснуться т.к. при dreamingEpisodeHistoryID в цикле возникла высокая значимость опасности
			if true{
				return false
			}
		}else{// можно ли заснуть?
	*/
	if extremImportanceObject != nil && extremImportanceObject.extremVal <= 5 {
		return false // нельзя засыпать, мучиться негативной проблемой
	}
	if CurrentInformationEnvironment.veryActualSituation ||
		CurrentInformationEnvironment.danger {
		return false // нельзя засыпать, опасно
	}

	return true // да можно заснуть или продолжать спать
}

////////////////////////////////////////////

func sleepingProcess() {
	// удаление временных массивов пробных действий (отрицательные значяения индеков) - во сне
	RemoveTemporeryTryActionArr()

	// элементы кратковременной памяти:
	InterruptMemory = nil
	saveBaseLinksCicleSynthesis = nil
	SynthesizeImportanceMemory = nil

	preparePurposeGeneticObject()
	prepareInformationEnvironment()
	prepareEpisodicMenory()
	//prepareImportanceObjectMenory() - просто очистить при просыпании

	lookForUnactualDominant() // просмотрт доминант для закрытия неактуальных
}

//////////////////////////////////////////////

// ////////////////////////////////////////////////////
func preparePurposeGeneticObject() {
	wCount := len(PurposeGeneticObject)
	if wCount == 0 {

	}
}

////////////////////////////////////////////////////////
/*т.к. объекты хранятся до Сна, то после обработки инфа должна перейти
в другие структуры (предположительные автоматизмы)

// реальный объект инфосреды, от которых нужно освобождать память
*/
func prepareInformationEnvironment() {
	// начать смотреть с начала, не доходя до последних элементов, записываемых во сне
	for i := 0; i < len(InformationEnvironmentObjects); i++ {
		// TODO
		//v=nil // освободить память просмотренного
	}
	// м.б. при просыпании просто забыть всю информацию ?
	//InformationEnvironmentObjects=nil

}

////////////////////////////////////////////////////////
/*  обработка в сновидениях, редуцирование СТАРЫХ использованных эпизодов
будут запоминаться только ВАЖНЫЕ эпизоды или последние 10000 эпизодов
*/
func prepareEpisodicMenory() {

}

//////////////////////////////////////////////////////////

/*
 */
func prepareImportanceObjectMenory() {

}

////////////////////////////////////////////

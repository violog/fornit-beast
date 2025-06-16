/*
Сторожевой центр сна
Заставляет проснуться при важных, опасных стимулах c дерева рефлексов.

*/

package sleep

import (
	"BOT/brain/action_sensor"
	"BOT/brain/gomeostas"
	"BOT/lib"
)

/////////////////////////////////////////////////////

/*
	сторожевой центр сна - в reflexes func activeReflexTree(

sleep.QuardingSleepCenter()
*/
func QuardingSleepCenter() {
	if !IsSleeping {
		return
	}

	// просыпание
	if sleepNecessityValue < 60 {
		WakeUpping()
	}
}

//////////////////////////////////////////////////////
///////////////////////////////////////////////////////
/*e


Если хочется спать
*/
func IsPossibleToSleep() bool {

	if sleepNecessityValue > 120 { // если слишком большая потребность спать, то заснуть
		return true // заснуть
	}
	if gomeostas.CommonBadNormalWell == 1 { // плохо
		return false // нельзя
	}

	// массив текущих пусковых стимулов
	curPultActionsArr := action_sensor.CheckCurActions()
	if lib.ExistsValInArr(curPultActionsArr, 1) || // непонятно
		lib.ExistsValInArr(curPultActionsArr, 3) || // наказать
		lib.ExistsValInArr(curPultActionsArr, 10) || // сделать больно
		lib.ExistsValInArr(curPultActionsArr, 15) { // испугаться
		return false // нельзя
	}

	return true // да можно заснуть
}

/////////////////////////////////////
////////////////////////////////////////////

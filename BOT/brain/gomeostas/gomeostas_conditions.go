/* функции использования текущих условий гомесотаза
 */

package gomeostas

import "sort"

/*
	выявить ID параметров гомеостаза как цели для улучшения в данных условиях

Возвращает PurposeGenetic.veryActual gomeostas.FindTargetGomeostazID
сортировка по уменьшению важности
*/
func FindTargetGomeostazID() (bool, []int) {
	var veryActual = false
	var idArr []int
	// BadNormalWell - состояние каждого параметра гомеостаза: 1 - Похо, 2 - Норма, 3 - Хорошо
	// отсортировать по убыли важности GomeostazParamsWeight
	badNormalWellImp := sortingForImpotents()

	for k, pID := range badNormalWellImp {
		if pID == 1 { // плохо для данного параметра гомеостаза
			idArr = append(idArr, k)
			if k == 1 || k == 2 || k == 7 || k == 8 {
				veryActual = true
			}
		}
	}
	return veryActual, idArr
}

// Сортировка ID параметров гомеостаза по убыванию значимости GomeostazParamsWeight
/* // какой-то левый вариант, было по-другому, н непонятно, почему сменил на этот. Правильный -внизу.
func sortingForImpotents() []int {
	var impC = make([]int, 110)//make(map[int]int)
	//for id, _ := range BadNormalWell {
	for id := 1; id < 9; id++ {
		impC[GomeostazParamsWeight[id]] = id
	}

	vals := make([]int, 0, len(impC))
	for k := range impC {
		vals = append(vals, k)
	}
	// СОРТИРОВКА ПО УБЫВАНИЮ
	sort.Slice(vals, func(i, j int) bool {
		return vals[i] > vals[j]
	})

	var arr = make([]int, 10)//make(map[int]int)
	for k,v := range vals {
		arr[impC[v]] = k//BadNormalWell[impC[v]]
	}

	return arr
}*/

/* // варианта Алексея
func sortingForImpotents() map[int]int{
	var arr=make(map[int]int)

	var impC = make([]int, 10)
	for id := 1; id < 9; id++ {
		impC[id] = GomeostazParamsWeight[id]
	}
	sort.Slice(impC, func(i, j int) bool {
		return impC[i] > impC[j]
	})
//	var arr = make([]int, 10)
	//for k,_ := range impC {
	for id := 1; id < 10; id++ {
		arr[id] = BadNormalWell[id]
	}

	return arr
}
*/

/*
на выходе func sortingForImpotents должна быть карта:
индекс = id параматра по убыванию значимости
значение - BadNormalWell[id]
*/
func sortingForImpotents() map[int]int {
	var impC = make(map[int]int)
	for id, _ := range BadNormalWell {
		impC[GomeostazParamsWeight[id]] = id
	}

	vals := make([]int, 0, len(impC))
	for k := range impC {
		vals = append(vals, k)
	}
	//СОРТИРОВКА ПО убыванию
	sort.Slice(vals, func(i, j int) bool {
		return vals[i] > vals[j]
	})

	var arr = make(map[int]int)

	for _, v := range vals {
		arr[impC[v]] = BadNormalWell[impC[v]]
	}

	return arr
}

/////////////////////////////////////////////////////////

// потребность общаться: гон, общение, обучение, поиск
func IsNeedForCommunication() bool {
	if GomeostazParams[3] > 85 || GomeostazParams[4] > 50 || GomeostazParams[5] > 50 || GomeostazParams[6] > 50 {
		return true
	}
	return false
}

///////////////////////////////////////////////

/* разные вспомогательные функции
 */

package lib

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	/*
		a0Arr:=[]int{2,1,3}
		aArr:=[]int{3,4,2,2,5,3,2,1,3}
		len0F:=len(a0Arr)
		lenF:=len(aArr)
		// есть ли в предшествующем фрагменте cycle.funcArr повторение cycle.func0Arr
		if lenF > 2*len0F {
			// предшествующй фрагмент cycle.funcArr
			fArr:=aArr[:(lenF-len0F)]
			prevF:=fArr[(len(fArr)-len0F):]
			if(EqualArrs(prevF, a0Arr)){
				return
			}
		}
		return
	*/
}

// удалить повторяющиеся значения
func RemoveDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// удалить повторяющиеся значения строк
/*
func RemoveDuplicateStringValues(intSlice []string) []string {
	keys := make(map[int]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
*/

// удалить из массива значение с индексом ID, сохранив порядок
func RemoveArrIndex(intSlice []int, id int) []int {
	if len(intSlice)-1 == id {
		intSlice = intSlice[:len(intSlice)-1]
	} else {
		copy(intSlice[id:], intSlice[id+1:])
	}
	return intSlice
}

// случайное из диапазона чисел
func RandChooseDiapazonQuest(first int, last int) int {
	rand.Seed(time.Now().UnixNano())
	n := first + rand.Intn(last-first+1)
	return n
}

// случайное из нескольких чисел
func RandChooseQuest(nums ...int) int {
	var count = 0
	var num []int
	for _, n := range nums {
		num = append(num, n)
		count++
	}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(count + 1)
	return num[n]
}

// случайное из массива
func RandChooseIntArr(n []int) int {
	var count = 0
	var num []int

	if n == nil {
		return 0
	}
	for _, n := range n {
		num = append(num, n)
		count++
	}
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(count)
	return num[v]
}

// сравнение массивовв int на идентичность
func EqualArrs(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// вернуть наиболее повторяющееся по числу значение массива и среди них первое из равных
func GetMaxCountVal(arr []int) int {
	// группируем значения в a
	var a = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		a[arr[i]]++
	}
	// находим максимальное или первое из равных
	var max = 0
	var ind = 0
	for k, v := range a {
		if max < v {
			max = v
			ind = k
		}
	}
	return ind
}

// индекс значения в массиве
func IndexValInArr(arr []int, val int) (bool, int) {
	if arr == nil {
		return false, 0
	}
	for k, n := range arr {
		if n == val {
			return true, k
		}
	}
	return false, 0
}

// есть такое значение в массиве
func ExistsValInArr(arr []int, val int) bool {
	if arr == nil {
		return false
	}
	for _, n := range arr {
		if n == val {
			return true
		}
	}
	return false
}

// есть такое значение в массиве строк
func ExistsValInStringArr(arr []string, val string) bool {
	if arr == nil {
		return false
	}
	for _, n := range arr {
		if n == val {
			return true
		}
	}
	return false
}

// есть такое значение в массиве с учетом его сортировки
func ExistsValInArrSort(arr []int, val int) bool {
	if arr == nil {
		return false
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return true
		}
	}
	return false
}

// есть полное строковое значение в строке значений
func ExistsValStrInList(List string, val string, razd string) bool {
	var sArr []string

	sArr = strings.Split(List, razd)
	for _, n := range sArr {
		if n == val {
			return true
		}
	}
	return false
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// имеют ли два числа одинаковые знаки.
func EqualSigns(num1, num2 int) bool {
	return (num1 >= 0 && num2 >= 0) || (num1 < 0 && num2 < 0)
}

// числа имеют разные знаки (одно положительное, другое отрицательное)
func IsDiffersOfSign(n1 int, n2 int) bool {
	if (n1 > 0 && n2 < 0) || (n1 < 0 && n2 > 0) {
		return true
	}
	return false
}

func AbsFloate(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// сохранить массив int в переменной
func SaveArrToVar(arr []int, to []int) []int {
	for a := 0; a < len(arr); a++ {
		to = append(to, arr[a])
	}
	return to
}

// текстовый массив в числовой
func IntArrToStrArr(sArr []string) []int {
	var out []int
	var id int

	for i := 0; i < len(sArr); i++ {
		id, _ = strconv.Atoi(sArr[i])
		if id != 0 {
			out = append(out, id)
		}
	}

	return out
}

// числовой массив в текстовый
func StrArrToIntArr(sArr []int) []string {
	var out []string
	var id string

	for i := 0; i < len(sArr); i++ {
		id = strconv.Itoa(sArr[i])
		out = append(out, id)
	}

	return out
}

// объединить 2 массива в один
func SumArr(arr1 []int, arr2 []int) []int {
	var out = append(arr1, arr2...)
	return UniqueArr(out)
}

// убрать дублеры в массиве
func UniqueArr(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// убрать дублеры в массиве 2D
func UniqueArr2D(intSlice [][]int) [][]int {
	var list [][]int
	listBuf := [][]int{{0}}

	for _, arr := range intSlice {
		isFind := true
		for _, arr1 := range listBuf {
			if EqualArrs(arr, arr1) == true {
				isFind = false
				break
			}
		}
		if isFind {
			list = append(list, arr)
			listBuf = append(listBuf, arr)
		}
	}
	return list
}

/*
	найти какие значения m1 есть в m2

т.е. m2 должен быть частью m1 или полностью совпадать с ним
*/
func GetExistsIntArs(m1 []int, m2 []int) []int {
	var found []int
	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m2); j++ {
			if m1[i] == m2[j] {
				found = append(found, m1[i])
			}
		}
	}
	return found
}

/*
	найти каких значений m1 нет в m2

т.е. m2 должен быть частью m1 или полностью совпадать с ним
*/
func GetDifferentIntArs(m1 []int, m2 []int) []int {
	var diff []int
	var exists = false
	for i := 0; i < len(m1); i++ {
		exists = false
		for j := 0; j < len(m2); j++ {
			if m1[i] == m2[j] {
				exists = true
				break
			}
			if !exists {
				diff = append(diff, m1[i])
			}
		}
	}
	return diff
}

// ////////////////////////////////////////
//
//	возвращает ближайшее целочисленное значение типа float64.
func RoundToFloate(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}

// возвращает ближайшее целочисленное значение типа int.
func RoundToInt(x float64) int {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return int(t + math.Copysign(1, x))
	}
	return int(t)
}

//////////////////////////////////////////

// Найти в массиве int индексы всех членов по значению
func FindIndexes(arr []int, value int) []int {
	indexes := []int{}
	for i, num := range arr {
		if num == value {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

/*
	не выдавать уже имеющиеся цифры в arr, т.е. вернет только уникальный ряд (сгруппирует)

Из idArr := []int{2, 25, 26, 2, 25, 26, 3, 5, 7, 2, 25, 26, 3, 3, 3, 12, 12, 2, 25, 26}
вернет: 1,25,26,2,5,6,12
*/
func RemoveDuplicates(arr []int) []int {
	result := []int{}
	for i := 0; i < len(arr); i++ {
		if ExistsValInArr(result, arr[i]) {
			continue
		}
		result = append(result, arr[i])
	}
	return result
}

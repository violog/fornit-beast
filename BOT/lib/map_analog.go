/* иммитация map на основе обычных массивов для работы с массивами типа nodeFromID
Паники типа "concurrent map writes" уже не будет.
Гонки же лоступа оказываются не опасными и повлияют на стабильность.
Смысл в том, что при записи проверяется существование индекса и если его нет, массив наращивается недостающими членами.

Для кажого массива задается:
var AFromID []*aNode  // сам массив
//var AFromID = make([]*aNode, len(strArr))//задать сразу имеющиеся в файле число при загрузке
// запись члена
func WriteAFromID(index int, value *aNode) {
	addAFromID(index)
	AFromID[index] = value
}
func addAFromID(index int){
	if index >= len(AFromID) {
		newSlice := make([]*aNode, index+1)
		copy(newSlice, AFromID)
		AFromID = newSlice
	}
}
// считывание члена
func ReadeAFromID(index int) (*aNode,bool){
	if index >= len(AFromID) || AFromID[index]==nil {
		return nil,false
	}
	return AFromID[index],true
}
Шаблон использования:
var AFromID []*anode
AFromID = make([]*anode, len(str))//задать сразу имеющиеся в файле число при загрузке из файла

WriteAFromID(index, value)
addAFromID(index)
node,ok:=ReadeAFromID(id)
if !ok {return }
В случае передобра range НУЖНО ПРОВЕРЯТЬ if v==nil{continue}:
for _, v := range{
if v==nil{continue}
////////////////////////////////////////////

В случае массива массивов (например, var importanceConditinArr [][]*importance)
перед операией append()
// если нет еще индекса ProblemID, то нарастить
	addImportanceConditinArr(ProblemID)
	importanceConditinArr[ProblemID]=append(importanceConditinArr[ProblemID],&node)
///////////////////////////

Можно спосокйно использовать for id, v := range AutomatizmFromId {
и там просто проверять if v!=nil для массивов, возвращающих указатель

	for k, v := range AutomatizmFromId {
		if v==nil{
			continue
		}
}

Не следует забивать большее число членов, чем их реально т.к. вознкнет проблема опрделения числа реальных членов.

*/

package lib

//////////////////////////  ТЕСТИРОВАНИЕ
/*
func init() {
	MapExperiment() // задержка всегда 10 сек !
	return
}
var yyyArr=make([]*aNode,100)
func MapExperiment(){
	fmt.Println("ПРОВЕРКА")
	WriteAFromID(1,&aNode{1,333})
	go www()
	go rrr()

	time.Sleep(10 * time.Second)// подождем когда все горутины пройдут

	return
}
//////////////////////////////////////
func rrr(){
	for i := 0; i < 10000; i++ {
		for k, _ := range yyyArr {
			_,ok:=ReadeAFromID(k)
			if 	ok{
			}
		}
	}

	for i := 0; i < 10000; i++ {
		_,ok:=ReadeAFromID(i)
		if 	ok{

		}
	}
	return
}
func www(){
	for i := 0; i < 10000; i++ {
		//WriteAFromID(i,&aNode{1,i})
	}
}

////////////////////////////////////////////////

type aNode struct {
ID int
val int
}

var AFromID []*aNode
//var AFromID = make([]*aNode, 20000)//задать сразу имеющиеся в файле число
func WriteAFromID(index int, value *aNode) {
	if index >= len(AFromID) {
		// Увеличиваем размер среза до нужного индекса
		newSlice := make([]*aNode, index+1)
		// Копируем значения из старого среза в новый
		copy(newSlice, AFromID)
		AFromID = newSlice
	}
	// Записываем значение в ячейку с указанным индексом
	AFromID[index] = value
}
func ReadeAFromID(index int) (*aNode,bool){
	if index >= len(AFromID) {
		return nil,false
	}
	return AFromID[index],true
}
*/
/////////////////////////////////////////////

///////// GPT предложило
/*
type MyStruct struct {
	Value int
}

type DynamicArray struct {
	data map[int]*MyStruct
}

// NewDynamicArray creates and initializes a new DynamicArray
func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		data: make(map[int]*MyStruct),
	}
}

// AddElement adds an element at a specified index
func (da *DynamicArray) AddElement(index int, element *MyStruct) {
	da.data[index] = element
}

// GetElement retrieves an element at a specified index
func (da *DynamicArray) GetElement(index int) (*MyStruct, error) {
	if val, ok := da.data[index]; ok {
		return val, nil
	}
	return nil, fmt.Errorf("index %d not found", index)
}

// RemoveElement removes an element at a specified index
func (da *DynamicArray) RemoveElement(index int) error {
	if _, ok := da.data[index]; ok {
		delete(da.data, index)
		return nil
	}
	return fmt.Errorf("index %d not found", index)
}

func main() {
	myArray := NewDynamicArray()

	myArray.AddElement(100, &MyStruct{Value: 10})
	myArray.AddElement(200, &MyStruct{Value: 20})
	myArray.AddElement(300, &MyStruct{Value: 30})

	fmt.Println("Map after adding elements:", myArray.data)

	element, err := myArray.GetElement(200)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Element at index 200:", element)
	}

	err = myArray.RemoveElement(200)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Map after removing element at index 200:", myArray.

*/

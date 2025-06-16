/* Вспомогательные функции

 */

package psychic

///////////////////////////////////////////////////////////////////////////////

// сила эффекта правил PARAMS[0] в зависимости от count PARAMS[1]
func getWpower(effect int, count int) int {
	if count < 3 {
		return effect
	}
	if count < 6 {
		return effect * 2
	}
	return effect * 3
}

// сила значимости стимула PARAMS[2] в зависимости от count PARAMS[1]
func getOpower(stimulsEffect int, count int) int {
	if count < 3 {
		return stimulsEffect
	}
	if count < 6 {
		return stimulsEffect * 2
	}
	return stimulsEffect * 3
}

func getLimitCountEM() int {
	limitN := 20
	if EvolushnStage < 4 {
		limitN = 2
	}
	if EvolushnStage == 4 {
		limitN = 5
	}
	return limitN
}

//////////////////////////////////////////////////

// функция, находящая член массива ruleArr с максимально высоким сочетанием Effect и Count
func findBestRule(ruleArr []Rule) (int, Rule) {
	if ruleArr == nil {
		return -1, Rule{0, 0, 0, 0, 0}
	}
	maxEffectCount := -1 // чтобы нулевые эффекты тоже учитывались
	var bestRule Rule
	id := 0
	for idr, rule := range ruleArr {
		wEffect := getWpower(rule.Effect, rule.Count)
		if wEffect > maxEffectCount {
			maxEffectCount = wEffect
			bestRule = rule
			id = idr
		}
	}
	return id, bestRule
}

// /////////////////////////////////////////////////////
// функция, находящая член массива ruleArr с максимально низким сочетанием Effect и Count
func findWorseRule(ruleArr []Rule) (int, Rule) {
	if ruleArr == nil {
		return -1, Rule{0, 0, 0, 0, 0}
	}
	maxEffectCount := 0 // нулевые эффекты - не вредны
	var bestRule Rule
	id := 0
	for idr, rule := range ruleArr {
		wEffect := getWpower(rule.Effect, rule.Count)
		if wEffect < maxEffectCount {
			maxEffectCount = wEffect
			bestRule = rule
			id = idr
		}
	}
	return id, bestRule
}

///////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////
/* Вытащить из эпизод.памяти посленюю цепочку кадров по числу limit
чтобы использовать ее как шаблон для поиска правил.
Не стоит задавать limit больше 5-10 (оптимизируется)
*/
func getLastSequenceFromEpisodeMemory(limit int) []int {
	ln := len(EpisodicHistoryArr)
	if ln < limit {
		limit = ln - 1
	}
	ip := EpisodicHistoryArr[(ln - limit):]
	var out []int
	for n := 0; n < len(ip); n++ {
		out = append(out, ip[n].ID)
	}
	return out
}

/////////////////////////////////////////////////////////////////

// последний индекс массива EpisodicHistoryArr []int Если
func getEpisodeFromLastMemory(from int) int {
	ln := len(EpisodicHistoryArr)
	if ln == 0 {
		return 0
	}
	return ln - 1
}

//////////////////////////////////////////

// это кадр с учительским правилом
func isTeacherRuleNode(node *EpisodicTreeNode) bool {
	if node.PARAMS[0] == 100 {
		return true
	}
	return false
}

//////////////////////////////////////////////////////////////////

/*
	собрать в targetEpisodicStrIdArr массив только позитивных правил типа Rule из массива кадров эпиз.памяти

kind 0 - любой эффект, 1 - только позитивные правила, 2 - только негативные правила
*/
func ruleFromEpisodeIdArr(nArr []int, kind int) {
	targetEpisodicStrIdArr = nil
	for i := 0; i < len(nArr); i++ {
		node, ok := ReadeEpisodicTreeNodeFromID(nArr[i])
		if !ok {
			continue
		}
		pars := node.PARAMS
		if pars == nil {
			continue
		}

		effect := pars[0]
		if effect == 100 {
			effect = 1
		}
		if kind == 1 && effect < 0 {
			continue
		}
		if kind == 2 && effect >= 0 {
			continue
		}
		// добавить Правило
		rule := Rule{node.Trigger, node.Action, effect, pars[1], 0}
		targetEpisodicStrIdArr = append(targetEpisodicStrIdArr, rule)
	}
	return
}

/////////////////////////////////////////////////////////////////////////////

// собрать в targetEpisodicStrIdArr массив только негативных правил типа Rule из массива кадров эпиз.памяти
func ruleNegativeFromEpisodeIdArr(nArr []int) {
	targetEpisodicStrIdArr = nil
	for i := 0; i < len(nArr); i++ {
		node, ok := ReadeEpisodicTreeNodeFromID(nArr[i])
		if !ok {
			continue
		}
		pars := node.PARAMS
		if pars == nil {
			continue
		}

		effect := pars[0]
		//		if effect==100{effect=1}
		if effect < 0 {
			// добавить Правило
			rule := Rule{node.Trigger, node.Action, pars[0], pars[1], 0}
			targetEpisodicStrIdArr = append(targetEpisodicStrIdArr, rule)
		}
	}
	return
}

/////////////////////////////////////////////////////////////////////////////

/*
	искать цепочки подряд идущих эпизодов, начинающиеся с данного nArr[i],
	 кончающиеся хорошим эффектом и собирать их targetEpisodicStrIdArr[]

заранее должно быть уже обнулено:
targetEpisodicStrIdArr=nil
episodicStrIdArr=nil

typeRule - 1 - искать только прямые Правила, 2 - искать только учительские Правила, 3 - искать все виды Правил
*/
func getPositiveChainsFromPositivArs(typeRule int, nArr []int) {
	//тут скорей всего надо делать перебор с конца, чтобы искать сначала в самой свежей цепочке
	for i := 0; i < len(nArr); i++ {
		// с позитивным концом
		flgBreak := getEpisodicStrIdArr(typeRule, 1, nArr[i], getLimitCountEM())
		if flgBreak {
			// ищем строго в одной цепочке, чтобы не оборвать смысловую (причинно-следственную связь) при поиске.
			// Здесь (2 уровень сознания) причина-следствие определяется тупо по принципу принадлежности кадров памяти одной цепочке:
			// если не было обрыва времени ожидания (внимания) - значит все как то связано, а если был обрыв - это уже другое событие
			// связь  между цепочками или кадрами разных цепочек установливается произвольно на следующих уровнях мышления
			break
		}
	}

	if episodicStrIdArr != nil {
		if len(episodicStrIdArr) > 1 { // нашли несколько альтернативных цепочек, нужно выбрать лучшую из episodicStrIdArr
			// сравниваем Правила и выбираем лучшее по эффекту и уверенности
			var ruleArr []Rule
			for n := 0; n < len(episodicStrIdArr); n++ {
				curArr := episodicStrIdArr[n]
				last := curArr[len(curArr)-1]
				ruleArr = append(ruleArr, last)
			}
			_, rule := findBestRule(ruleArr)
			if rule.Effect >= 0 {
				targetEpisodicStrIdArr = []Rule{rule}
			}

		} else { // единственная
			targetEpisodicStrIdArr = append(targetEpisodicStrIdArr, episodicStrIdArr[0][0])
		}
	}
}

//////////////////////////////////////////////////////////////////

/*
	искать цепочки подряд идущих эпизодов, начинающиеся с данного nArr[i],

кончающиеся большим негативом и собирать их targetEpisodicStrIdArr[]

заранее должно быть уже обнулено:
targetEpisodicStrIdArr=nil
episodicStrIdArr=nil
*/
func getNegativeChainsFromPositivArs(typeRule int, nArr []int) {

	for i := 0; i < len(nArr); i++ {
		//искать цепочки с негативным концом
		flgBreak := getEpisodicStrIdArr(typeRule, 2, nArr[i], getLimitCountEM())
		if flgBreak {
			// ищем строго в одной цепочке, чтобы не оборвать смысловую (причинно-следственную связь) при поиске.
			// Здесь (2 уровень сознания) причина-следствие определяется тупо по принципу принадлежности кадров памяти одной цепочке:
			// если не было обрыва времени ожидания (внимания) - значит все как то связано, а если был обрыв - это уже другое событие, никак не связанное с предыдущим
			// связь  между цепочками или кадрами разных цепочек установливается произвольно на следующих уровнях мышления
			break
		}
	}

	if episodicStrIdArr != nil {
		if len(episodicStrIdArr) > 1 { // нашли несколько альтернативных цепочек, нужно выбрать лучшую из episodicStrIdArr
			// сравниваем Правила и выбираем лучшее по эффекту и уверенности
			var ruleArr []Rule
			for n := 0; n < len(episodicStrIdArr); n++ {
				curArr := episodicStrIdArr[n]
				last := curArr[len(curArr)-1]
				ruleArr = append(ruleArr, last)
			}
			_, rule := findWorseRule(ruleArr)
			if rule.Effect >= 0 {
				targetEpisodicStrIdArr = []Rule{rule}
			}

		} else { // единственная
			targetEpisodicStrIdArr = append(targetEpisodicStrIdArr, episodicStrIdArr[0][len(episodicStrIdArr[0])-1])
		}
	}
}

///////////////////////////////////////////

/*
	начиная с кадра ID, искать цепочку подряд идущих эпизодов исторической последовательности эпизодов EpisodicHistoryArr

начинающихся с данного nArr[i], и кончающегося пустым кадром (id=-1)
и вернуть их в массиве targetEpisodicStrIdArr[]
Находится суммарные негативные и позитивные эффекты в цепочке.
При typeEffect==1 собираются только цепочки с превышающим позитивным ээфектом.
При typeEffect==2 собираются только цепочки с превышающим негативным ээфектом.
При typeEffect==0 - любые суммарные эффекты

typeRule - 1-искать только прямые Правила, 2-искать только учительские Правила, 3- искать все виды Правил
*/
func getEpisodicStrIdArr(typeRule int, typeEffect int, id int, limit int) bool {
	if id == -1 {
		return true // не обрабатывать пустой кадр
	}
	// последний индекс исторической памяти EpisodicHistoryArr[]
	lastI := len(EpisodicHistoryArr) - 1
	var node *EpisodicTreeNode
	var strIdArr []Rule // для формирования цепочки Правил
	effectSum := 0      // максимальный эффект промежуточных звеньев
	endStr := false
	// отмотывать назад для формирования цепочек для каждого EpisodicHistoryArr[i].ID == id
	for i := lastI; i >= 0; i-- {
		// пустой кадр
		if endStr {
			break // конец цепочки, начало следующей  для кадра id
		}
		// как только наткнулись на кадр с целевым id смотрим для него цепочку
		if EpisodicHistoryArr[i].ID == id {
			// начиная с i посмотреть что там дальше и сформировать цепочку
			for n := i; n < lastI || ((n-i) < limit && n < lastI); n++ { //во второй части условия нужно учесть n < lastI, иначе периодически выдает выход за пределы индексов массива
				if EpisodicHistoryArr[n].ID == -1 {
					endStr = true
					break
				}

				//				node=EpisodicTreeNodeFromID[EpisodicHistoryArr[i]]
				node0, ok := ReadeEpisodicTreeNodeFromID(EpisodicHistoryArr[n].ID)
				if !ok {
					continue
				}
				node = node0

				pars := node.PARAMS
				if pars == nil {
					continue
				}

				effect := pars[0]
				count := pars[1]
				if typeRule == 1 && effect == 100 {
					continue
				} // не учитывать учительскаие правила
				if typeRule == 2 && effect != 100 {
					continue
				}

				effectSum += getWpower(effect, count)

				// добавить Правило
				rule := Rule{node.Trigger, node.Action, pars[0], pars[1], 0}
				strIdArr = append(strIdArr, rule)
			} //for n := i; n < lastI || (n-i) < limit; n++ {
			canAddString := false
			if strIdArr != nil && len(strIdArr) > 0 {
				if typeEffect == 1 && effectSum > 0 {
					canAddString = true
				}
				if typeEffect == 2 && effectSum < 0 {
					canAddString = true
				}
				if typeEffect == 0 {
					canAddString = true
				}
			}
			if canAddString {
				episodicStrIdArr = append(episodicStrIdArr, strIdArr)
			}
		} //if EpisodicHistoryArr[i] == id {
	} //for i := lastI; i > 0; i-- {

	return endStr
}

////////////////////////////////////////////////////////////////////

// найти ID исторической памяти для эпизода с акциями Trigger или Action, начиная с номера beginID
func getFrameFromTrigger(beginID int, actID int) int {

	for i := beginID; i < len(EpisodicHistoryArr); i++ {

		node, ok := ReadeEpisodicTreeNodeFromID(EpisodicHistoryArr[i].ID)
		if !ok {
			continue
		}
		if node.PARAMS == nil { // на всякий случай, хотя такого не должно быть
			continue
		}
		if node.Trigger == actID || node.Action == actID {
			return i
		}

	}
	return 0
}

////////////////////////////////////////////////

// от данного ID исторической памяти пройти назад до пустого кадра и вернуть его ID
func returnEmptyHistoryID(beginID int) int {
	for i := beginID; i >= 0 && i > 0; i-- {

		node, ok := ReadeEpisodicTreeNodeFromID(EpisodicHistoryArr[i].ID)
		if !ok {
			continue
		}
		if node.PARAMS == nil { // на всякий случай, хотя такого не должно быть
			continue
		}
		if node.Action == 0 {
			return i
		}
	}
	return 0
}

//////////////////////////////////////////////////

/*
начиная с узла дерева собрать все конечные узлы рекурсивно в episodiCadresArr[]

еще есть
func getIdArr(typeRule int, node *EpisodicTreeNode, Trigger int, Action int) - вытащить все ID узлов, в которых есть условие, начиная с node
*/
var episodiCadresArr []*EpisodicTreeNode

func getAllChidsFromNode(rt *EpisodicTreeNode) {
	episodiCadresArr = nil
	getAllChidsFromNodeProc(rt)
}
func getAllChidsFromNodeProc(rt *EpisodicTreeNode) {
	if len(rt.Children) == 0 { // конец ветки
		episodiCadresArr = append(episodiCadresArr, rt)
		return
	}
	for i := 0; i < len(rt.Children); i++ {
		getAllChidsFromNodeProc(&rt.Children[i]) // рекурсия на слудующие дочки
	}
}

//////////////////////////////////////////////////////

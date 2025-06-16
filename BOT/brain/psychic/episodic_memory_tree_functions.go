/*  функции Дерева эпизодической памяти

TODO
Кроме работы с кадрами памяти для 1 и 2 уровней осмывсления, где все дается функциями в готовом виде
и поэтому очень механично, нужно ввести инфо-функции более детального рассмотрения цадров с выводами.

*/

package psychic

import (
	word_sensor "BOT/brain/words_sensor"
	"BOT/lib"
)

///////////////////////////////////////////////////////////////////////

/*
Найти все ID узлов EpisodicTreeNode для которых выполняется условие с любым эффектом.

level == 0 - для всех трех условий (NodeAID,NodeSID,NodePID)
level == 1 - для двух первых условий (NodeAID,NodeSID)
level == 2 - только для NodeAID

Значение Trigger определеяется обязательно, обычно это curStimulImageID
Значение Action может быть ==0 и тогда результат покажет возможные варианты ответов на Trigger
Если определены Trigger и Action, то результат оказывается предсказанием того, какой эффект будет при Action.

- ДЛЯ ТЕКУЩИХ СТИМУЛА (Trigger) Ответ (Action) не определен - его и нужно найти подходящий

т.к. пишется как прямое так и учительское правило, то можно подставлять Ответ в Стимул (Trigger int)

typeRule - 1-искать только прямые Правила, 2-искать только учительские Правила, 3- искать все виды Правил
*/
var ePsrsIdArr []int

func getEpisodesArrFromConditions(typeRule int, level int, Trigger int, Action int) []int {
	if Trigger == 0 {
		if !wasRunProvocationFunc { // для func infoFunc31 Trigger равен 0
			return nil
		}
	}

	ePsrsIdArr = nil

	switch level {
	case 0:
		nodePID := detectedActiveLastProblemNodID
		if nodePID == 0 { // еще нет detectedActiveLastProblemNodID
			return nil
		}
		cond := []int{CurrentCommonBadNormalWell, CurrentEmotionReception.ID, nodePID}
		//	 cond= []int{67,38,107,1,}; //Action=3
		id, _ := findEpisodicBrange(0, cond, &EpisodicTree) // ищет ветку, которая полностью соотвествует условиям
		if id == 0 {
			return nil
		}
		node, ok := ReadeEpisodicTreeNodeFromID(id)
		if !ok {
			return nil
		}
		// если в ветке нет одного из условий, то - не найдено
		if node.BaseID != cond[0] || node.EmotionID != cond[1] || node.NodePID != cond[2] {
			return nil
		}
		getIdArr(typeRule, node, Trigger, Action)
		return ePsrsIdArr

	case 1:
		cond := []int{CurrentCommonBadNormalWell, CurrentEmotionReception.ID} // нужно выделять только те, у которых Trigger==curStimulImageID
		//cond= []int{67,38,}; Trigger=1; Action=3
		id, _ := findEpisodicBrange(0, cond, &EpisodicTree)
		if id == 0 {
			return nil
		}
		//node, ok := EpisodicTreeNodeFromID[id]
		node, ok := ReadeEpisodicTreeNodeFromID(id)
		if !ok {
			return nil
		}
		if node.BaseID != cond[0] || node.EmotionID != cond[1] {
			return nil
		}
		getIdArr(typeRule, node, Trigger, Action)
		return ePsrsIdArr

	case 2:
		cond := []int{CurrentCommonBadNormalWell} // нужно выделять только те, у которых Trigger==curStimulImageID
		//cond= []int{67,38,}; Trigger=1; Action=3
		id, _ := findEpisodicBrange(0, cond, &EpisodicTree)
		if id == 0 {
			return nil
		}
		//node, ok := EpisodicTreeNodeFromID[id]
		node, ok := ReadeEpisodicTreeNodeFromID(id)
		if !ok || node.BaseID != cond[0] {
			// найти при любых условиях по всему дереву
			getIdArr(typeRule, &EpisodicTree, Trigger, Action)
			return ePsrsIdArr
		}

		getIdArr(typeRule, node, Trigger, Action)
		return ePsrsIdArr
	}

	return nil
}

/*
	вытащить все ID узлов, в которых есть условие, начиная с node

еще есть
func getAllChidsFromNode(rt *EpisodicTreeNode) -начиная с узла дерева собрать все конечные узлы рекурсивно в episodiCadresArr[]
*/
func getIdArr(typeRule int, node *EpisodicTreeNode, Trigger int, Action int) {
	for _, child := range node.Children {

		if child.PARAMS != nil { // узел с прописанным PARAMS
			if Trigger > 0 && child.Trigger != Trigger {
				continue
			}
			if Action > 0 && child.Action != Action {
				continue
			}
			if typeRule == 1 && isTeacherRuleNode(&child) {
				continue
			}
			if typeRule == 2 && !isTeacherRuleNode(&child) {
				continue
			}
			if child.ID == -1 {
				continue
			}

			ePsrsIdArr = append(ePsrsIdArr, child.ID)
		}
		// продолжим рекурсивно искать в нем далее
		getIdArr(typeRule, &child, Trigger, Action)
	}

	return
}

////////////////////////////////////////////

/*
	Поиск ответного действия по Правилам, принцип GPT: fornit.ru/66139.

Здесь не учитываются учительские правила чтобы цепочки были только из прямых правил
Получить приемлемую последовательность id EpisodicTreeNode от данного Стимула Trigger
в виде глобального массива targetEpisodicStrIdArr цепочек Правил.
Назовет ее целевой последовательностью, т.е. если на данный Trigger отвечать targetEpisodicStrIdArr[0]
а опертор ответит EpisodicTreeNodeFromID[targetEpisodicStrIdArr[1]].Trigger и т.д.
то придем к тому позитивному результату, который прогнозировался при получении targetStrIdArr.
Если оператор ответит непредсказуемо, то нужно снова вызвать getTargetEpisodicStrIdArr(EpisodicTreeNodeFromID[targetEpisodicStrIdArr[1]].Trigger)
и получить скорректированную последовательность.

targetStrIdArr состояит из зафиксированных ранее непрерывных цепочек в EpisodicHistoryArr,
между которыми пустой кадр

limit int - исторические цепочки не более чем в limit эпизодов
При поиске избегать узлов пустых кадров.

В конце найденных цепочек должно быть позитивное Правило, так их можно сравнивать по эффектам и выбрать лучшую.

Получаем цепочку (или одно правило), заканчивающееся позитивом.

	В цепочке только последний член кончается удачей, остальные отрицательные,
	но смысл цепочки в том, что они приводит к удаче и поеэтому в данных конкретных условиях
	следует выполнить первую акций из этой цепочки.

Пример вызова:
getTargetEpisodicStrIdArr(3,curActiveActionsID, 5)

	//Результат получаем в глобальном targetEpisodicStrIdArr []Rule
	//- окончательно выбранная целевая цепочка Правил откуда берем следующее Действие.
	if targetEpisodicStrIdArr != nil {
		// есть цепочка с конечным плюсовым эффектом, значит можно так действвоать
		rules = targetEpisodicStrIdArr[len(targetEpisodicStrIdArr)-1]
	}
	if rules.Action > 0 {
		res := makeActionFromRooles(rules) //	!!!!!ПЕРЕДЕЛАТЬ для type Rule struct {
		return res
	}
*/
var targetEpisodicStrIdArr []Rule // окончательно выбранная целевая цепочка Правил. Всегда есть.
// из цепочки Правил с конечным плюсовым эффектом можно вытащить следующее действие на исполнение.
var episodicStrIdArr [][]Rule // несколько найденных целевых цепочек Правил. Может не быть массива альтернативных цепочек.

func getTargetEpisodicStrIdArr(Trigger int, limit int) {

	targetEpisodicStrIdArr = nil
	episodicStrIdArr = nil

	// попытка найти в точном соотвествии с условиями
	nArr := getEpisodesArrFromConditions(1, 0, Trigger, 0)
	if nArr == nil {
		nArr = getEpisodesArrFromConditions(1, 1, Trigger, 0)
	}
	if nArr == nil {
		nArr = getEpisodesArrFromConditions(1, 2, Trigger, 0)
	}
	if nArr == nil { // нет никакого опыта с Trigger
		return
	}
	/////////////////////////////////////////

	/* искать цепочки подряд идущих эпизодов, начинающиеся с данного nArr[i],
	  кончающиеся хорошим эффектом и собирать их targetEpisodicStrIdArr[]
	Значение limit пока ==5 но может оптимизироваться (TODO)
	*/
	getPositiveChainsFromPositivArs(1, nArr)
	if targetEpisodicStrIdArr != nil {
		return
	}
	return
}

//////////////////////////////////////////////////////////////////

/*
	Выбрать лучшее Правило, как можно с более точным учетом условий НО БЕЗ СТИМУЛА

для провокации func infoFunc31
ДОЛЖНО БЫТЬ wasRunProvocationFunc = true
*/
func getProvocationBestRule() Rule {
	targetEpisodicStrIdArr = nil
	episodicStrIdArr = nil
	rule := Rule{0, 0, 0, 0, 0}
	if !wasRunProvocationFunc {
		return rule
	}

	// попытка найти в точном соотвествии с условиями
	nArr := getEpisodesArrFromConditions(3, 0, 0, 0)
	if nArr == nil {
		nArr = getEpisodesArrFromConditions(3, 1, 0, 0)
	}
	if nArr == nil {
		nArr = getEpisodesArrFromConditions(3, 2, 0, 0)
	}
	if nArr == nil { // нет никакого опыта с Trigger
		return rule
	}
	// собрать в targetEpisodicStrIdArr массив позитивных правил типа Rule из массива кадров эпиз.памяти
	ruleFromEpisodeIdArr(nArr, 1)
	// Выбрать одно лучшее Правило
	_, rule = findBestRule(targetEpisodicStrIdArr)

	return rule
}

////////////////////////////////////////////////////////////////////

/*
	Выбрать лучшее Правило, как можно с более точным учетом условий

typeRule - 1-искать только прямые Правила, 2-искать только учительские Правила, 3- искать все виды Правил
результат - все найденные Правила в targetEpisodicStrIdArr
и возвращает лучшее Правило
*/
func getSingleBestRule(typeRule, Trigger int) Rule {

	targetEpisodicStrIdArr = nil
	episodicStrIdArr = nil
	rule := Rule{0, 0, 0, 0, 0}

	// попытка найти в точном соотвествии с условиями
	nArr := getEpisodesArrFromConditions(typeRule, 0, Trigger, 0)
	if nArr == nil {
		nArr = getEpisodesArrFromConditions(typeRule, 1, Trigger, 0)
	}
	if nArr == nil {
		nArr = getEpisodesArrFromConditions(typeRule, 2, Trigger, 0)
	}
	if nArr == nil { // нет никакого опыта с Trigger
		return rule
	}
	// собрать в targetEpisodicStrIdArr массив позитивных правил типа Rule из массива кадров эпиз.памяти
	ruleFromEpisodeIdArr(nArr, 1)
	// Выбрать одно лучшее Правило
	_, rule = findBestRule(targetEpisodicStrIdArr)

	return rule
}

////////////////////////////////////////////////////////////////////

/*
	При оценке возможности запуска автоматизма

посмотреть в Правилах, м.б. после плохого эффекта последует следующее Правило с хорошим эффектом
(не только для GPT, а при typeRule==3 - любые Правила)
и тогда можно допустить Usefulness<0 в расчете на последующий успех.
Не в качестве волевого усилия, а чисто автоматически использовать такую информацию.

Для точного соблюдения условий 3-х деревьев.

Смотрим, было ли раньше, что после данного ответа Action получилось удачно,
причем не в каком-то одном случае, а в целом статистика.

GPT. ЗДЕСЬ не смотрится строго последовательность правил.

typeRule - 1-искать только прямые Правила, 2-искать только учительские Правила, 3- искать все виды Правил
*/
func isNextWellEffectFromActonRules(typeRule int, Trigger int, Action int) bool {
	/* Найти все ID узлов EpisodicTreeNode для которых выполняется условия 3-х деревьев
	между которыми есть пустой кадр
	*/
	targetEpisodicStrIdArr = nil
	episodicStrIdArr = nil
	episNodes := getEpisodesArrFromConditions(typeRule, 0, Trigger, Action)
	/* смотреть всегда точно !!!
	if episNodes==nil{// исключаем из условий 3-е дерево
		episNodes=getEpisodesArrFromConditions(typeRule,1,Trigger,Action)
	}
	if episNodes==nil{// исключаем из условий 2-е и 3-е дерево
		episNodes=getEpisodesArrFromConditions(typeRule,2,Trigger,Action)
	}*/

	/* искать цепочки подряд идущих эпизодов, начинающиеся с данного nArr[i],
	  кончающиеся хорошим эффектом и собирать их targetEpisodicStrIdArr[]
	Значение limit пока ==5 но может оптимизироваться (TODO)
	*/
	getPositiveChainsFromPositivArs(typeRule, episNodes)
	if targetEpisodicStrIdArr != nil {
		return true
	}

	/*
		if episNodes == nil {
			return false
		}
		for i := 0; i < len(episNodes); i++ {
			node, ok := ReadeEpisodicTreeNodeFromID(episNodes[i])
			if !ok {
				continue
			}
			effect:=node.PARAMS[0]
				if effect==100{effect=1}
				if effect > 0 { // [хоть один вариант позитивный
					return true
				}
		}*/
	return false
}

/////////////////////////////////////////////////////////////////

/*
	выдать []Rule для данного ID фразы с позитивным Эфектом

чтобы из полученных действий составить цепочку для автоматизма
exactlyCond - true - правила для текущих условий, false - только для 1-го дерева
*/
func getWellRoolesFromPhraseId(phraseID int, exactlyCond bool) []Rule {
	var rArr []Rule
	var cond []int
	if exactlyCond {
		cond = []int{CurrentCommonBadNormalWell, CurrentEmotionReception.ID}
	} else {
		cond = []int{CurrentCommonBadNormalWell}
	}
	/////////////////////////
	id, _ := findEpisodicBrange(0, cond, &EpisodicTree)
	if id == 0 {
		return nil
	}
	node, ok := ReadeEpisodicTreeNodeFromID(id)
	if !ok {
		return nil
	}
	// вытащить все конечные ветки, начиная с node
	getAllChidsFromNode(node)
	if episodiCadresArr == nil {
		return nil
	}

	for i := 0; i < len(episodiCadresArr); i++ {
		act, ok := ReadeActionsImageArr(episodiCadresArr[i].Trigger)
		if !ok {
			continue
		}
		if act.PhraseID[0] != phraseID {
			continue
		}
		effectCompare := 0
		params := episodiCadresArr[i].PARAMS
		effect := params[0]
		if effect == 100 {
			effect = 1
		}
		if effect > effectCompare {
			rArr = append(rArr, Rule{episodiCadresArr[i].Trigger, episodiCadresArr[i].Action, effect, params[1], 0})
			break
		}
	}

	return rArr
}

/////////////////////////////////////////////////////////////////

/*
выдать []Rule для данного ID слова с позитивным Эфектом

чтобы из полученных действий составить цепочку для автоматизма
exactlyCond - true - правила для текущих условий, false - только для 1-го дерева

ф-ция вызывается когда не распознана фраза
!!! если нераспознана фраза, то detectedActiveLastNodID НЕВЕРЕН его нельзя использовать как поиск по эп.памяти

	поэтому нужно использовать lastWellPhrasedetectedActiveLastNodID
*/
func getWellRulesFromWordId(wordID int, exactlyCond bool) []Rule {
	var rArr []Rule
	var cond []int
	if exactlyCond {
		cond = []int{CurrentCommonBadNormalWell, CurrentEmotionReception.ID}
	} else {
		cond = []int{CurrentCommonBadNormalWell}
	}
	/////////////////////////
	id, _ := findEpisodicBrange(0, cond, &EpisodicTree)
	if id == 0 {
		return nil
	}
	node, ok := ReadeEpisodicTreeNodeFromID(id)
	if !ok {
		return nil
	}
	// вытащить все конечные ветки, начиная с node
	getAllChidsFromNode(node)
	if episodiCadresArr == nil {
		return nil
	}

	for i := 0; i < len(episodiCadresArr); i++ {
		eArr, ok := ReadeActionsImageArr(episodiCadresArr[i].Trigger)
		if !ok {
			continue
		}
		// есть ли такое слово во фразе
		if eArr.PhraseID == nil {
			continue
		}
		aArr := word_sensor.GetWordsArrFromPraseID(eArr.PhraseID[0])

		//if act.PhraseID[0] != phraseID {
		if !lib.ExistsValInArr(aArr, wordID) { // нет такого слова во фразе
			continue
		}
		effectCompare := 0
		params := episodiCadresArr[i].PARAMS
		effect := params[0]
		if effect == 100 {
			effect = 1
		}
		if effect > effectCompare {
			rArr = append(rArr, Rule{episodiCadresArr[i].Trigger, episodiCadresArr[i].Action, effect, params[1], 0})
			//break
		}
	}

	return rArr
}

////////////////////////////////////////////////////////////////////////

/*
	Пpедсказание позитивного эффекта от Ответных действий для данных условий

Сначала ищется конец цепочки с позитивным эффектом, который превышает негативы предыдущих звеньев
и возвращает эффект конечного звена так, что действие дает позивные предсказания даже если сначала приводит к негативу.
Если таких цепочек не находится, то просто делается статистика действия после данного стимула.
Если нет такой статистики, что выдается результат статистики действия с любыми стимулами.

Возвращает точность предсказания (1-самое точное,) если найдено и величину вероятного эффекта.
Если не найдено возвращает 0,0

Отимистическое предсказание, т.е. из найденных правил выбирается наилучший жффект.
*/
func positiveFromActionAfterStimul(Trigger int, Action int) (int, int) {
	targetEpisodicStrIdArr = nil
	episodicStrIdArr = nil

	// попытка найти в точном соотвествии с условиями  цепочки прямых правил с решающим позитивом на конце
	nArr := getEpisodesArrFromConditions(1, 0, Trigger, Action)
	if nArr == nil {
		nArr = getEpisodesArrFromConditions(1, 1, Trigger, Action)
	}
	if nArr == nil {
		nArr = getEpisodesArrFromConditions(1, 2, Trigger, Action)
	}
	if nArr != nil { // есть кадры
		/* искать цепочки подряд идущих эпизодов, начинающиеся с данного nArr[i],
		  кончающиеся хорошим эффектом и собирать их targetEpisodicStrIdArr[]
		Значение limit пока ==5 но может оптимизироваться (TODO)
		*/
		getPositiveChainsFromPositivArs(1, nArr)
		if targetEpisodicStrIdArr != nil {
			return 1, targetEpisodicStrIdArr[0].Effect
		}
		// раз нет цепочек, учтем просто правила для условий эмоций
		// собрать в targetEpisodicStrIdArr массив позитивных правил типа Rule из массива кадров эпиз.памяти
		ruleFromEpisodeIdArr(nArr, 1)
		// Выбрать одно лучшее Правило
		_, rule := findBestRule(targetEpisodicStrIdArr)
		return 2, rule.Effect
	}

	/* попытка найти в точном соотвествии с условиями, но не цепочки, а отдельные правила, только прямые (без учительских)
	Если не найдено для точных условий (эмоции), будет искать при любых условиях по всему дереву.
	*/
	nArr = getEpisodesArrFromConditions(1, 0, Trigger, Action)
	if nArr == nil {
		nArr = getEpisodesArrFromConditions(1, 1, Trigger, Action)
	}
	if nArr == nil {
		nArr = getEpisodesArrFromConditions(1, 2, Trigger, Action)
	}
	if nArr != nil {
		// собрать в targetEpisodicStrIdArr массив позитивных правил типа Rule из массива кадров эпиз.памяти
		ruleFromEpisodeIdArr(nArr, 1)
		// Выбрать одно лучшее Правило
		_, rule := findBestRule(targetEpisodicStrIdArr)
		return 2, rule.Effect
	}

	// искать без учета стимула
	ePsrsIdArr = nil
	getIdArr(1, &EpisodicTree, Trigger, Action)
	if ePsrsIdArr != nil {
		// собрать в targetEpisodicStrIdArr массив позитивных правил типа Rule из массива кадров эпиз.памяти
		ruleFromEpisodeIdArr(ePsrsIdArr, 1)
		// Выбрать одно лучшее Правило
		_, rule := findBestRule(targetEpisodicStrIdArr)
		return 3, rule.Effect
	}

	return 0, 0
}

/////////////////////////////////////////////////////////////////////

/*
	Пpедсказание негативного эффекта от Ответных действий для данных условий

Пессимистическое предсказание, т.е. из найденных правил выбирается наихудший жффект.
*/
func negativeFromActionAfterStimul(Trigger int, Action int) (int, int) {
	targetEpisodicStrIdArr = nil
	episodicStrIdArr = nil

	nArr := getEpisodesArrFromConditions(1, 0, Trigger, Action)
	if nArr == nil {
		nArr = getEpisodesArrFromConditions(1, 1, Trigger, Action)
	}
	if nArr == nil {
		nArr = getEpisodesArrFromConditions(1, 2, Trigger, Action)
	}
	if nArr != nil {
		// собрать в targetEpisodicStrIdArr массив негативных правил типа Rule из массива кадров эпиз.памяти
		ruleFromEpisodeIdArr(nArr, 2) // 2 - массив негативных правил
		// Выбрать одно худшее Правило
		_, rule := findWorseRule(targetEpisodicStrIdArr)
		return 2, rule.Effect
	}
	return 0, 0
}

///////////////////////////////////////////////////////////////////////////

/*
	предсказание последствий выполнения действия автоматизма

для understanding_functions.go func getPrognoze(atmtzm *Automatizm)

Алгоритм:
Поиск цепочек, заканчивающихся большим негативом (превышающим позитив промежуточных)
Если не найдены цепочки, то поиск негатива в правилах с учетом эмоций и Стимула Trigger
Если нет таких, то поиск правил с участием Action

т.е. в первую очередь смотрим нет ли негатива и если нет, то автоматизм может запускаться.

Возвращает accuracy,effect:
accuracy == 1 - точное предсказание для действия, точное предсказание для действияно потом будет позитив
accuracy == 2 - менее точное предсказание для действия, совершенное после Стимула curStimulImageID в данных условиях
accuracy == 3 - неточное предсказание для действия при любом стимуле и любых условиях
*/
///////////////////////
/* СТАРАЯ ВЕРСИЯ
func getPrognoseFromAutmtzmAction(Trigger int, Action int) (int, int) {
	targetEpisodicStrIdArr = nil
	episodicStrIdArr = nil

	// Все кадры, точно отвечающие условиям - в nArr[]
	nArr := getEpisodesArrFromConditions(1, 0, Trigger, Action)
	if nArr == nil {
		nArr = getEpisodesArrFromConditions(1, 1, Trigger, Action)
	}
	if nArr == nil {
		nArr = getEpisodesArrFromConditions(1, 2, Trigger, Action)
	}
	if nArr != nil { // есть кадры
		// искать цепочки подряд идущих эпизодов, начинающиеся с данного nArr[i],
		//кончающиеся большим негативом и собирать их targetEpisodicStrIdArr[]

targetEpisodicStrIdArr = nil
episodicStrIdArr = nil
getNegativeChainsFromPositivArs(1, nArr)
if targetEpisodicStrIdArr != nil {
return 1, targetEpisodicStrIdArr[0].Effect
} ////////////////////

// раз нет цепочек, учтем просто правила для условий эмоций
targetEpisodicStrIdArr = nil
episodicStrIdArr = nil
// собрать в targetEpisodicStrIdArr массив позитивных правил типа Rule из массива кадров эпиз.памяти
ruleFromEpisodeIdArr(nArr, 2)
// Выбрать одно худшее Правило
_, ruleМ := findWorseRule(targetEpisodicStrIdArr)

// вдруг есть позитив
targetEpisodicStrIdArr = nil
episodicStrIdArr = nil
getPositiveChainsFromPositivArs(1, nArr)
ruleFromEpisodeIdArr(nArr, 1)
// Выбрать одно худшее Правило
_, ruleP := findBestRule(targetEpisodicStrIdArr)
if ruleP.Effect > ruleМ.Effect {
return 2, ruleP.Effect
}
return 2, ruleМ.Effect
}

// искать без учета стимула
targetEpisodicStrIdArr = nil
episodicStrIdArr = nil
ePsrsIdArr = nil
getIdArr(1, &EpisodicTree, Trigger, Action)
if ePsrsIdArr != nil {
// собрать в targetEpisodicStrIdArr массив негативных правил типа Rule из массива кадров эпиз.памяти
ruleFromEpisodeIdArr(nArr, 2) // 2 - массив негативных правил
// Выбрать одно лучшее Правило
_, ruleМ := findWorseRule(targetEpisodicStrIdArr)

// вдруг есть позитив
targetEpisodicStrIdArr = nil
episodicStrIdArr = nil
getPositiveChainsFromPositivArs(1, nArr)
ruleFromEpisodeIdArr(nArr, 1)
// Выбрать одно худшее Правило
_, ruleP := findBestRule(targetEpisodicStrIdArr)
if ruleP.Effect > ruleМ.Effect {
return 2, ruleP.Effect // 2 т.к. найдено в цепочках
}
return 3, ruleМ.Effect
}
return 0, 0
}
*/

/*
	предсказание последствий выполнения действия автоматизма

для understanding_functions.go func getPrognoze(atmtzm *Automatizm)

Алгоритм:
Поиск цепочек, заканчивающихся большим негативом (превышающим позитив промежуточных)
Если не найдены цепочки, то поиск негатива в правилах с учетом эмоций и Стимула Trigger
Если нет таких, то поиск правил с участием Action

т.е. в первую очередь смотрим нет ли негатива и если нет, то автоматизм может запускаться.

Возвращает accuracy,effect:
accuracy == 1 - точное предсказание для действия, точное предсказание для действияно потом будет позитив
accuracy == 2 - менее точное предсказание для действия, совершенное после Стимула curStimulImageID в данных условиях
accuracy == 3 - неточное предсказание для действия при любом стимуле и любых условиях
*/
func getPrognoseFromAutmtzmAction(Trigger int, Action int) (int, int) {
	targetEpisodicStrIdArr = nil
	episodicStrIdArr = nil

	// Все кадры, точно отвечающие условиям - в nArr[]
	nArr := getEpisodesArrFromConditions(1, 0, Trigger, Action)
	if nArr != nil {
		ef := finalCommonResult(nArr)
		return 1, ef
	}
	nArr = getEpisodesArrFromConditions(1, 1, Trigger, Action)
	if nArr != nil {
		ef := finalCommonResult(nArr)
		return 2, ef
	}
	nArr = getEpisodesArrFromConditions(1, 2, Trigger, Action)
	if nArr != nil { // есть кадры
		ef := finalCommonResult(nArr)
		return 3, ef
	} ////////////////////

	// искать без учета стимула
	targetEpisodicStrIdArr = nil
	episodicStrIdArr = nil
	ePsrsIdArr = nil
	getIdArr(1, &EpisodicTree, Trigger, Action)
	if ePsrsIdArr != nil {
		ef := finalCommonResult(nArr)
		return 3, ef
	}
	return 0, 0
}

// вернет положительное или отрицательное значение эффекта
func finalCommonResult(nArr []int) int {
	targetEpisodicStrIdArr = nil
	episodicStrIdArr = nil
	// собрать в targetEpisodicStrIdArr массив позитивных правил типа Rule из массива кадров эпиз.памяти
	ruleFromEpisodeIdArr(nArr, 2)
	// Выбрать одно худшее Правило
	_, ruleМ := findWorseRule(targetEpisodicStrIdArr)

	// вдруг есть позитив
	targetEpisodicStrIdArr = nil
	episodicStrIdArr = nil
	getPositiveChainsFromPositivArs(1, nArr)
	ruleFromEpisodeIdArr(nArr, 1)
	// Выбрать одно лучшее Правило
	_, ruleP := findBestRule(targetEpisodicStrIdArr)

	if ruleP.Effect > ruleМ.Effect {
		return ruleP.Effect
	}
	return ruleМ.Effect

	//return ruleP.Effect, ruleМ.Effect
}

///////////////////////////////////////////////////////////

/*
выдать все последующие кадры в цепочках, после данного (memFr *EpisodicTreeNode)
до пустых кадров (ID==-1)
*/
func getNextHistiryEpisodicFrame(memFr *EpisodicTreeNode) []*EpisodicTreeNode {
	if len(EpisodicHistoryArr) == 0 {
		return nil
	}
	startID := memFr.ID
	var out []*EpisodicTreeNode
	for n := 0; n < len(EpisodicHistoryArr)-2; n++ {
		if EpisodicHistoryArr[n].ID == startID && EpisodicHistoryArr[n+1].ID != -1 {
			frm, ok := ReadeEpisodicTreeNodeFromID(EpisodicHistoryArr[n+1].ID)
			if ok {
				out = append(out, frm)
			}
		}
	}

	return out
}

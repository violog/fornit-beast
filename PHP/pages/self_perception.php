<?
/*   Страница текущего состояния психики Beast
http://go/pages/self_perception.php  

*/
$page_id=7;
$title="Психика Beast";
include_once($_SERVER['DOCUMENT_ROOT']."/common/header.php");
include_once($_SERVER['DOCUMENT_ROOT']."/common/show_waiting.php");
include_once($_SERVER['DOCUMENT_ROOT']."/common/spoiler.php");

?>
<div  style='position:absolute;top:38px;left:180px;font-family:courier;font-size:16px;cursor:pointer;' onClick="location.reload(true)"><b>Обновить</b></div>
<!-- div  class='navigator_button' style='position:absolute;top:38px;left:300px;cursor:pointer;' onClick="get_autimat_table()" title='Ментальные автоматизмы'>Таблица автоматизмов</div -->

<div class='navigator_button' style='position:absolute;top:38px;left:510px;cursor:pointer;' onClick="get_tree1()" title='Дерево текущей ситуации с произвольной активацией'>Дерево ситуации</div>

<div class='navigator_button' style='position:absolute;top:38px;left:670px;cursor:pointer;' onClick="get_tree_problem()" title='Дерево проблем'>Дерево проблем</div>


<div id='div_id' style='font-family:courier;font-size:16px;'>Нужен коннект с Beast.</div>
</div>
</center>
<div style='margin-top:10px;text-align:left;'>

<div class='navigator_button' style='cursor:pointer;color:red;' onClick="clian_memory_psy();" title='Очистить всю память психики в папке memory_psy'><b>Очистить всю память психики (папка memory_psy)</b></div>


<hr style='width:90%' align='left'>


<span class="spoiler_header"  onclick="open_close('lib_block1_id',1)" style="cursor:pointer;font-size:16px"><?=set_sopiler_icon('lib_block1_id')?><b>Дерево эпизодической памяти, Правила</b></span>
<div id="lib_block1_id" class="spoiler_block spoiler" style="position:relative;z-index:10;top:0px;left:0px;padding-left:15px;background-color:#ffffff;width:1100px;height:0px;">
Дерево сохраняет Правила и Модели понимания (см. ниже спойлер Модели понимания). Здесь показываются 100 последних Правил, записанных в кадрах эпизодической памяти, включая "учительские".<br>
<br>
Можно вручную очистить эпизодическую памяь.<br> 
Для этого нужно выключить Beast,<br>
и обнулисть файлы /memory_psy/episodic_tree.txt и /memory_psy/episodic_history.txt<br>
ОБЯЗАТЕЛЬНО ОБА ЭТИХ ФАЙЛА!
</div>
<div class='navigator_button' style='cursor:pointer;color:red;' onClick="clian_episodic();" title='Очистить эпизодическую память ТОЛЬКО ПОСЛЕ ВКЛЮЧЕНИЯ Beast!'><b>Очистить эпизодическую память</b></div>
<div  class='navigator_button' style='cursor:pointer;' onClick="get_rules()" title='Таблица моторынх Правил'><b>Правила</b></div>
<hr style='width:90%' align='left'>

<!-- span class="spoiler_header" onclick="open_close('lib_block4_id',1)" style="cursor:pointer;font-size:16px"><?=set_sopiler_icon('lib_block4_id')?><b>Дерево эпизодической памяти, Модели понимания</b></span>
<div id="lib_block4_id" class="spoiler_block spoiler" style="position:relative;z-index:10;top:0px;left:0px;padding-left:15px;background-color:#ffffff;width:1100px;height:0px;">
Модель понимания для данного объекта - это все кадры эпизодической памяти, где он участвует.<br>
В карту модели понимания объекта (с индексом == extremImportance.extremObjID )
записывается массив всех ID эпизодической памяти, где он участвует.<br>
Ключом (индексом массива) карты моделей явояктся ID образа типа extremImportance.extremObjID.
<br -->
</div>
<div  class='navigator_button' style='cursor:pointer;' onClick="get_map_uundstd_obj()" title=''><b>Карта моделей понимания</b></div>
<hr style='width:90%' align='left'>


<span class="spoiler_header" onclick="open_close('lib_block11_id',1)" style="cursor:pointer;font-size:16px"><?=set_sopiler_icon('lib_block11_id')?><b>Дерево ментальной эпизодической памяти. Ментальные правила (автоматизмы).</b></span>
<div id="lib_block11_id" class="spoiler_block spoiler" style="position:relative;z-index:10;top:0px;left:0px;padding-left:15px;background-color:#ffffff;width:1100px;height:0px;">
Носитель ментальных правил: какая цепочка инфо-функций в условиях данной проблемы и темы привела к данному эффекту.<br>
Выполняет роль Ментальных автоматизмов при поиске мент.правила (последователньости инфо-функций) для уверенного запуска в данных условиях.

<br>
</div>
<div class='navigator_button' style='cursor:pointer;color:red;' onClick="clian_ment_episodic();" title='Очистить эпизодическую память ТОЛЬКО ПОСЛЕ ВКЛЮЧЕНИЯ Beast!'><b>Очистить мент. эпизодическую память</b></div>
<div  class='navigator_button' style='cursor:pointer;' onClick="get_ment_rules()" title='Таблица ментальных Правил'><b>Ментальные правила</b></div>
<hr style='width:90%' align='left'>




<span class="spoiler_header"  onclick="open_close('lib_block2_id',1)" style="cursor:pointer;font-size:16px"><?=set_sopiler_icon('lib_block2_id')?><b>Циклы осмысления</b></span>
<div id="lib_block2_id" class="spoiler_block spoiler" style="position:relative;z-index:10;top:0px;left:0px;padding-left:15px;background-color:#ffffff;width:1100px;height:0px;">
Цикл осмысления - субъективный ориентировочный рефлекс, каждой шаг которого основывается на информации, даваемой предыдущим шагом
с целью найти подходящие действия для данной ситуации, что дает возможность снова сориентироваться.<br>
После каждой активации по объективному оринентировочному рефлексу возникает цикл ментальной обработки информации, который заключается в поиске подходящей информационной функции, запуска ментального автоматизма этой функции, получения новой информации, использования ее последущей инфо-функцией и так - до создания ментального автоматизма запуска моторного действия, которое проверяется на успешность и прописывается как моторный автоматизм.<br>
</div>
<div  class='navigator_button' style='cursor:pointer;' onClick="get_cicles()" title='Ментальные циклы - кратковременная память'><b>Циклы</b></div>
<hr style='width:90%' align='left'>



<span class="spoiler_header"  onclick="open_close('lib_block33_id',1)" style="cursor:pointer;font-size:16px"><?=set_sopiler_icon('lib_block33_id')?><b>Цепочки действий (программы действий)</b></span>
<div id="lib_block33_id" class="spoiler_block spoiler" style="position:relative;z-index:10;top:0px;left:0px;padding-left:15px;background-color:#ffffff;width:1100px;height:0px;">
Последовательность отдельных известных образов действий образует программу действий с пусковым начальным звеном. Такая цепочка действий, хранящаяся в оперативной памяти, может быть запущена как целая программа последовательности действий и тогда по оценке результата первое звено становится автоматизмом, а последующие - продолжением его действия.<br>Становится возможной осознанная корректировка сложных программ действия и выполнение этой программы уже без размышления.<br>Так же становится возможным осознанный прогноз возможных результатов этих действий.<br>
</div>
<div  class='navigator_button' style='cursor:pointer;' onClick="get_nexts()" title=''><b>Цепочки действий</b></div>
<hr style='width:90%' align='left'>




<span class="spoiler_header"  onclick="open_close('lib_block3_id',1)" style="cursor:pointer;font-size:16px"><?=set_sopiler_icon('lib_block3_id')?><b>Значимость элементов восприятия</b></span>
<div id="lib_block3_id" class="spoiler_block spoiler" style="position:relative;z-index:10;top:0px;left:0px;padding-left:15px;background-color:#ffffff;width:1100px;height:0px;">
Значимость элементов восприятия - как объекта произвольного внимания: 
того из всего воспринимаемого, что имеет наибольшую значимость
т.к. именно наибольшая значимость должна осмысливаться.<br>

Кроме того, значимости объектов в связке с кадром эпиз.памчти - это и есть модель понимания данного объекта внимания -
его значимость в разных условиях и то, какие действия могут быть совершены при этом (UnderstandingModel).<br>

Значимость - величина от -10 0 до 10, приобретаемая объектов внимания в данной ситуации
- берется из результата пробных действий и связывается ос всеми компонентами воспринимаемого в этих условиях.<br>

Оценке значимости подлежат элементы действия оператора:
кнопки воздействия, фразы и отдельные слова, принимающие значимость фразы.<br>
</div>
<div  class='navigator_button' style='cursor:pointer;' onClick="get_imp_obj()" title=''><b>Объекты значимости</b></div>


<hr style='width:90%' align='left'>
<span class="spoiler_header" onclick="open_close('lib_block5_id',1)" style="cursor:pointer;font-size:16px"><?=set_sopiler_icon('lib_block5_id')?><b>Доминанты</b></span>
<div id="lib_block5_id" class="spoiler_block spoiler" style="position:relative;z-index:10;top:0px;left:0px;padding-left:15px;background-color:#ffffff;width:1100px;height:0px;">
Доминанта нерешенной проблемы – постоянно существующая структура,  где запоминается проблема, условия и попытки ее решения.<br>
Доминанта активируется во многих местах адаптивных механизмов и позволяет находить решение новым способом – по аналогии.<br>
При удачном решении и его реальной проверки, доминанты закрывается (закрытие гештальта), но остается в базе опыта решений.<br>
Доминанта — это просто пролонгированная Цель. Повышенное время существования Доминанты повышает шансы найти решение по аналогии func checkRelevantAction
Активная Доминанта, если она есть, ВСЕ ВРЕМЯ сопровождает появление новой инфы восприятия так, чтобы была возможность найти решение по аналогии.<br>
Решенная Доминанта является третьим типом формирования автоматизмов на уровне пятой ступени развития: творческое формирование автоматизма.<br>
Решенные Доминанты являются источников личного опыта и аналогов решения проблем.
<br>
</div>
<div  class='navigator_button' style='cursor:pointer;' onClick="get_dominant()" title=''><b>Доминанты</b></div>

<hr style='width:90%' align='left'>


</div>
<br>
<br>
<br>
<br>




<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var linking_address='<?include($_SERVER["DOCUMENT_ROOT"]."/common/linking_address.txt");?>';

// ждем пока не включат бестию
check_Beast_activnost(6);// после 4-го пульса И запускается get_info()

function get_info()
{
wait_begin();
var AJAX = new ajax_support(linking_address+"?get_self_perception_info=1",sent_info);
AJAX.send_reqest();
function sent_info(res)
{
//alert(res);
wait_end();
document.getElementById('div_id').innerHTML=res;

setTimeout("chech_new_info()",2000);
}
}
//get_info();
//setTimeout("chech_new_info()",2000);
/*
function get_autimat_table()
{
open_anotjer_win("/pages/mental_automatizm_table.php");
}*/

function get_tree1()
{
//alert("Дерево автоматизмов");
open_anotjer_win("/pages/mental_automatizm_tree.php");
}
function get_tree_problem()
{
//alert("Дерево автоматизмов");
open_anotjer_win("/pages/mental_problem_tree.php");
}
function get_rules()
{
//alert("Дерево понимания");
open_anotjer_win("/pages/rules.php");
}

function get_ment_rules()
{
open_anotjer_win("/pages/mental_rules.php");
}

function get_mental_rules()
{
open_anotjer_win("/pages/mental_rules.php");
}
function get_cicles(){
open_anotjer_win("/pages/mental_cicles.php");
}

function get_map_uundstd_obj(){
open_anotjer_win("/pages/mental_undastending_models.php");
}

function get_dominant()
{
open_anotjer_win("/pages/mental_dominants.php");
}
//////////////////////////////

var old_size=0;
function chech_new_info()
{
var AJAX = new ajax_support("/pages/self_perception_checher.php",sent_size_info);
AJAX.send_reqest();
function sent_size_info(res)
{
	//alert(res);
if(old_size!=res)
{
get_info();
}
old_size=res;
setTimeout("chech_new_info()",2000);
}
}


function get_imp_obj()
{
open_anotjer_win("/pages/mental_importance.php");
}



function show_object(type,id)
{
var AJAX = new ajax_support(linking_address + "?objID="+id+"&objType="+type+"&get_importance_object_info=1", sent_undstg_info);
AJAX.send_reqest();
function sent_undstg_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}

function show_dominant(id) // DominantaInfoStr(dID)
{
var AJAX = new ajax_support(linking_address + "?objID="+id+"&objType="+type+"&get_dominant_info=1", sent_undstg_info);
AJAX.send_reqest();
function sent_undstg_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}

}


function get_problem_tree(id)
{
var AJAX = new ajax_support(linking_address + "?autNodeID="+id+"&get_problem_tree_node=1", sent_purpose_info);
AJAX.send_reqest();
function sent_purpose_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}

}

function show_object(type,id)
{
var AJAX = new ajax_support(linking_address + "?objID="+id+"&objType="+type+"&get_importance_object_info=1", sent_undstg_info);
AJAX.send_reqest();
function sent_undstg_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}

function get_nexts()
{
//список цепочек действий AmtzmNextString
open_anotjer_win("/pages/automatizm_next_actions_table.php");
}
function show_next_actions(nextID)//func GetNextActionsInfo(nextID int)string{
{
var AJAX = new ajax_support(linking_address + "?nextID="+nextID+"&get_next_actions_info=1", sent_emotion_info);
AJAX.send_reqest();
function sent_emotion_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}


function clian_episodic()
{ 
show_dlg_confirm("Точно очистить эпизодическую память?",1,-1,clian_episodic2);
}
var clian_episodic_3=0;
function clian_episodic2()
{
if (clian_episodic_3)
{
show_dlg_alert("Уже запущено.",1000);
}
clian_episodic_3=setTimeout("clian_episodic3()",2000);
var AJAX = new ajax_support(linking_address + "?clian_episodic_memory=1", sent_clian_episodic);
AJAX.send_reqest();
function sent_clian_episodic(res) {
clearTimeout(clian_episodic_3);
if(res=="did")
	{
show_dlg_alert("Эпизодическая память очищена.",2000);
	}else{
show_dlg_alert("Ошибка с очисткой...",0);
	}
}
}
function clian_episodic3()
{
show_dlg_alert("Нужно выключить Beast и подождать 6 секунд!",2000);
clian_episodic_3=0;
}



function clian_ment_episodic()
{ 
show_dlg_confirm("Точно очистить эпизодическую память?",1,-1,clian_ment_episodic2);
}
var clian_ment_episodic_3=0;
function clian_ment_episodic2()
{
if (clian_ment_episodic_3)
{
show_dlg_alert("Уже запущено.",1000);
}
clian_episodic_3=setTimeout("clian_ment_episodic3()",2000);
var AJAX = new ajax_support(linking_address + "?clian_ment_episodic_memory=1", sent_clian_ment_episodic);
AJAX.send_reqest();
function sent_clian_ment_episodic(res) {
clearTimeout(clian_ment_episodic_3);
if(res=="did")
	{
show_dlg_alert("Эпизодическая память очищена.",2000);
	}else{
show_dlg_alert("Ошибка с очисткой...",0);
	}
}
}
function clian_ment_episodic3()
{
show_dlg_alert("Нужно выключить Beast и подождать 6 секунд!",2000);
clian_ment_episodic_3=0;
}



function clian_memory_psy()
{
show_dlg_confirm("Точно очистить все файлы памфти в /memory_psy/?", "Да", "Нет", clian_memory_psy2);
}
function clian_memory_psy2()
{
	var AJAX = new ajax_support("/tools/cliner_mempry.php", sent_cliner_p_info);
		AJAX.send_reqest();

		function sent_cliner_p_info(res) {
			//alert(res);
			if (res[0] != "!") {
				show_dlg_alert(res.substr(1), 0);
				return;
			}
			show_dlg_alert("все файлы памяти в /memory_psy/ очищены.<br><br>Если Beast включена, ее нужно перезапустить!", 0);
		}
}

</script>

</div>
</body>
</html>
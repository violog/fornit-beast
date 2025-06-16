<?
/*   Стрнаница автоматизмов Beast
http://go/pages/automatizm.php  

*/
$page_id=6;
$title="Автоматизмы Beast";
include_once($_SERVER['DOCUMENT_ROOT']."/common/header.php");
include_once($_SERVER['DOCUMENT_ROOT']."/common/show_waiting.php");


?>
<div  style='position:absolute;top:38px;left:250px;font-family:courier;font-size:16px;cursor:pointer;color:blue;' onClick="get_autimat_table"><b>Обновить</b></div>

<div  style='position:absolute;top:38px;left:370px;font-family:courier;font-size:16px;cursor:pointer;color:blue;' onClick="get_autimat_table()">Таблица автоматизмов</div>
<div  style='position:absolute;top:38px;left:590px;font-family:courier;font-size:16px;cursor:pointer;color:blue;' onClick="get_tree1()">Дерево автоматизмов</div>
<div  style='position:absolute;top:38px;left:800px;font-family:courier;font-size:16px;cursor:pointer;color:blue;' onClick="get_rulles()">Правила</div>


<!--  НЕТ СРАЗУ ЗАГРУЖАЕМОГО КОНТЕНТА div id='div_id' style='font-family:courier;font-size:16px;'>Нужен коннект с Beast.</div>
</div -->


<div style="width:1000px;">
Предоставляемые здесь инструменты возможно использовать не только для ознакомления и тестирования (рекомендации - в самом низу), но в этом случае они не могут заменить полноценное воспитание, вот почему.<br>
Если уже созданы впрок все возможные сочетания безусловных и условных рефлексов (их по 6000 каждого вида), то при запуске создания автоматизмов создается настолько большое и <b>избыточное</b> количество автоматизмов, что работа подсмотрщика дерева автоматизма оказывается очень долгой и совершенно бесполезной (хотя дерево автоматизмов работает вполне оперативно). Кроме того, в ходе развития последующих стадий образуется лавинообразно еще большее число вторичных автоматизмов (это какая-то шизофрения получается), что совершенно не адекватно нормальному процессу воспитания.<br>
<hr>

<?
///// стадии развития 
$stages=read_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/stages.txt");
$stages=trim($stages);
if ($stages > 3) {
	echo "<div style='color:red;border:solid 1px #8A3CA4;padding:10px;background-color:#DDEBFF;'>Инструменты ниже <b>НЕ СЛЕДУЕТ ИСПОЛЬЗОВАТЬ</b> потому, что уже пройдена 3-я стадия развития. Использовать эти можно, только ясно представлявляя происходящее и умея работать с файлами памяти.</div>";
}
?>

<div  style='display:inline-block;font-family:courier;font-size:18px;cursor:pointer;color:blue;border:solid 1px #8A3CA4;border-radius: 7px;padding-left:4px;padding-right:4px;
background-color:#eeeeee;' title="Для тестирования различных конфигураций автоматизмов - очистка предудыщих." onClick="cliner_all_automatizm()">Удалить все автоматизмы и очистить дерево автоматизмов</div> - для создания новой порции автоматизмов. Или можно самим удалить содержимое в файлах /memory_psy/automatizm_images.txt и /memory_psy/automatizm_tree.txt и перезагрузить Beast.<br>
<br>


Для тестирования возможно избежать долгий период воспитания с формированием автоматизмов и просто сгенерировать автоматизмы на основе существующих рефлексов (с приоритетом условных рефлексов).<br>
При этом у автоматизмов будут установлены опции уже проверенного автоматизма с полезностью, равной 1 (вполне полезно). Это правомерно потому, что рефлексы создавались уже полезными для своих условий.<br>
В дальнейшем такие автоматизмы будут проверяться в зависимости от реакции оператора и изменения состояния Beast, корректируясь настолько эффективно, насколько позволяет текущая стадия развития.
<br><br>
<div id="make_genetic_automatizms_id" style='display: inline-block;relative;font-family:courier;font-size:16px;cursor:pointer;
border:solid 1px #8A3CA4;border-radius: 7px;padding-left:4px;padding-right:4px;
background-color:#eeeeee;' onClick="make_genetic_automatizms()">Создать автоматизмы на основе существующих <b>безусловных</b> рефлексов</div> - связываются с 3-м уровнем ветки дерева<br>

<br>
<div id='res_div_id' style='font-family:courier;font-size:21px;color:green;font-weight:bold;'></div>
<div id='div_id' style='font-family:courier;font-size:21px;color:red;font-weight:bold;'></div>
</div>
<div id="make_automatizms_id" style='display:inline-block;relative;font-family:courier;font-size:16px;cursor:pointer;
border:solid 1px #8A3CA4;border-radius: 7px;padding-left:4px;padding-right:4px;
background-color:#eeeeee;' onClick="make_automatizms()">Создать автоматизмы на основе существующих <b>условных</b> рефлексов</div> - связываются с 7-м (вербальным) уровнем ветки дерева. Если условные рефлексы были сгенерированы искусственно, то будет создано большое число автоматизмов. Но <span style='color:red'>нет смысла создавать автоматизмы из условных рефлексов, если затем воспользоваться кнопкой "Сформировать зеркальные автоматизмы для всех таблиц сочетаний контекстов"</span> потому, что вторые перекроют первых.<br>
<br>

</div>

<hr>
<b>Для стадии 3</b> так же необходимо длительное время для отзеркаливания реакций оператора в различных ситуациях. Это время возможно сократить для тестирования, запустив редактор создания зеркальных автоматизмов (первичного жизненного импринтингового опыта) на основе существующих автоматизмов.<br>
При этом уже будут вставлены умолчательные значения в виде предположительной инверсии реакций автоматизма.
<br><br>
<div  style='display: inline-block;relative;font-family:courier;font-size:16px;cursor:pointer;
border:solid 1px #8A3CA4;border-radius: 7px;padding-left:4px;padding-right:4px;
background-color:#eeeeee;' onClick="open_anotjer_win('/pages/mirrors_automatizm.php')" >Начать набивку зеркальных автоматизмов на основе существующих</div>


<?
// сколько файлов уже есть в mirror_reflexes_basic_phrases
$m_file_count=0;
$tdir=$_SERVER["DOCUMENT_ROOT"]."/lib/mirror_reflexes_basic_phrases/";
$filesArr="var filesArr = new Array();";
$n=0;
if($dh = opendir($tdir)) 
{ //exit("!!!");
while(false !== ($file = readdir($dh))) 
{		
if($file=="." || $file=="..")
	continue;
if(filesize($tdir.$file)>0)
	{
$m_file_count++;
	}
}
closedir($dh);
}

echo "&nbsp;&nbsp;&nbsp;&nbsp;количество созданных для отзеркаливания файлов: <span style='font-size:20px'><b>".$m_file_count."</b></span>";
if($m_file_count)
{

echo "<br><br>
<div  style='display: inline-block;relative;font-family:courier;font-size:16px;cursor:pointer;
border:solid 1px #8A3CA4;border-radius: 7px;padding-left:4px;padding-right:4px;
background-color:#eeeeee;' onClick='open_anotjer_win(\"/pages/mirrors_automatizm_maker.php\")' >Сформировать зеркальные автоматизмы для всех таблиц сочетаний контекстов</div> - очень большое число автоматизмов, которые, впрочем, будут нормально отрабатывать.";
echo "";
}


if(file_exists($_SERVER["DOCUMENT_ROOT"]."/lib/mirror_basic_phrases_common.txt") && filesize($_SERVER["DOCUMENT_ROOT"]."/lib/mirror_basic_phrases_common.txt")>20)
{
echo "<br><br><div  style='display: inline-block;relative;font-family:courier;font-size:16px;cursor:pointer;
border:solid 1px #8A3CA4;border-radius: 7px;padding-left:4px;padding-right:4px;
background-color:#eeeeee;' onClick='open_anotjer_win(\"/pages/mirrors_automatizm_maker_from_template.php\")' >Сформировать зеркальные автоматизмы только для таблицы общего шаблона</div> - это наиболее разумное решение, т.к. будет создано относительно недольшое число автоматизмов. И в этом случае не требуется заполнять таблицы по каждому сочетанию Базовых контекстов.";
}
?>
<br>
<br>
<b>Рекомендуемая последовательность для искусственного создания рабочей коллекции автоматизмов.</b><br>
1. Заполнить общий шаблон пусковых символов <a href="/pages/condition_reflexes_basic_phrases_common.php">/pages/condition_reflexes_basic_phrases_common.php</a><br>
2. Нажать кнопку “Удалить все автоматизмы и очистить дерево автоматизмов”<br>
3. Нажать кнопку “Создать автоматизмы на основе существующих безусловных рефлексов”<br>
4. Нажать кнопку “Создать автоматизмы на основе существующих услоных рефлексов”<br>
5. Заполнить общий шаблон ответов для имитации отзеркаливания <a href="/pages/mirror_basic_phrases_common.php">/pages/mirror_basic_phrases_common.php</a><br>
6. Нажать кнопку “Сформировать зеркальные автоматизмы только для таблицы общего шаблона”.<br> 
Даже в таком усеченном варианте получится более 12 тыс. автоматизмов, которые будут грузиться при включении несколько секунд.<br>
<br>
<b>Альтернативный вариант (неподьемный для персонального компьютера):</b><br>
7. Заполнить таблицы для всех сочетаний Базовых контекстов <a href="/pages/mirrors_automatizm.php">/pages/mirrors_automatizm.php (очень большая работа)</a><br>
8. Нажать кнопку “Сформировать зеркальные автоматизмы для всех таблиц сочетаний контекстов”.<br>
Такой вариант сделает загрузку при включении очень долгой и НЕ РЕКОМЕНДУЕТСЯ.
<br>
Для тестирования же возможно использовать любые варианты.<br>

<br>
<b>Правильнее всего:</b><br>
Не делать автоматизмы с огромной избыточностью, а формировать их естесенным общением с Beast, но это потребует очень много времени.
<br>
<br>
<br>
<br>

</div>

</div>



<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var linking_address='<?include($_SERVER["DOCUMENT_ROOT"]."/common/linking_address.txt");?>';

var bot_is_connected=0;
var AJAX = new ajax_support(linking_address + "?check_Beast_activnost=1", check_conn_info);
AJAX.send_reqest();
function check_conn_info(res) {
	bot_is_connected=1;
}


function cliner_all_automatizm()
{ 
show_dlg_confirm("Дерево автоматизмов и сами автоматизмы будут очищены.<br>Вы уверены?","Да, очистить!","Отмена",cliner_all_automatizm2);
}
function cliner_all_automatizm2()
{
var AJAX = new ajax_support("/lib/cliner_all_automatizm_memory.php", sent_cliner_reflex_memory);
AJAX.send_reqest();
function sent_cliner_reflex_memory(res) {

if(bot_is_connected==0)
show_dlg_alert("Файлы памяти автоматизвом очищены.",0);
else// нужно обесточить GO
{
show_dlg_alert("Файлы памяти автоматизмов очищены.<br>Best выключается для очистки памяти.",0);

/*
var server = "/kill.php";
		var AJAX = new ajax_support(server, sent_end_answer);
		AJAX.send_reqest();
		function sent_end_answer(res) {
			show_dlg_alert("Beast выключен.", 2000);
		}*/
//Выключить без сохранения памяти (bot_closing=2), просто погасить исполняемый файл
		var AJAX = new ajax_support(linking_address + "?bot_closing=2", sent_info);
		AJAX.send_reqest();
		function sent_info(res) {
			// не будет ответа
		}
}
}
}
////////////////////////////////////////


function get_autimat_table()
{
open_anotjer_win("/pages/automatizm_table.php");
}

function get_tree1()
{
//alert("Дерево автоматизмов");
open_anotjer_win("/pages/automatizm_tree.php");
}
function get_tree2()
{
alert("Дерево понимания");
}
// процесс идет в ГО
var type_reqwest_go=0;// 1 - создать автоматизмы из рефлексов
function make_automatizms()
{
document.getElementById('make_automatizms_id').disabled=true
//exists_connection(); // если нет коннекта, будет сообщение
// ждем пока не включат бестию
check_Beast_activnost(6);// после 4-го пульса И запускается get_info()
type_reqwest_go=1; //alert(type_reqwest_go);
}
function get_info() 
{
if(type_reqwest_go==1)/////////////////////////////
{
var linking_address = '<? include($_SERVER["DOCUMENT_ROOT"] . "/common/linking_address.txt"); ?>';
wait_begin();
var AJAX = new ajax_support(linking_address + "?make_automatizms_from_reflexes=1", sent_process_info);
AJAX.send_reqest();
function sent_process_info(res) {
			//alert(res);
wait_end();
document.getElementById('res_div_id').innerHTML = res;
document.getElementById('make_automatizms_id').disabled=false
}
}/////////////////////////////

if(type_reqwest_go==2)/////////////////////////////
{
var linking_address = '<? include($_SERVER["DOCUMENT_ROOT"] . "/common/linking_address.txt"); ?>';
wait_begin();
var AJAX = new ajax_support(linking_address + "?make_automatizms_from_genetic_reflexes=1", sent_process_info);
AJAX.send_reqest();
function sent_process_info(res) {
			//alert(res);
wait_end();
document.getElementById('res_div_id').innerHTML = res;
document.getElementById('make_genetic_automatizms_id').disabled=false
}
}/////////////////////////////

}
////////////////////////////////
function make_genetic_automatizms()
{
document.getElementById('make_genetic_automatizms_id').disabled=true
//exists_connection(); // если нет коннекта, будет сообщение
// ждем пока не включат бестию
check_Beast_activnost(4);// после 4-го пульса И запускается get_info()
type_reqwest_go=2; //alert(type_reqwest_go);
}







// станица правил
function get_rulles()
{
open_anotjer_win("/pages/rules.php");
}
</script>

</body>
</html>
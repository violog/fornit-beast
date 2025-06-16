<?
/*   Стадии развития Beast
http://go/pages/stages.php  

*/
$page_id=10;
$title="Стадии развития Beast";
include_once($_SERVER['DOCUMENT_ROOT']."/common/header.php");

include_once($_SERVER['DOCUMENT_ROOT']."/common/common.php");

//////////////////////////////////////// САБМИТЫ
if(isset($_POST['gogogo'])&&$_POST['gogogo']==1)
{
//var_dump($_POST);exit("<hr>!!!!!!! ".$_POST['next']);
$next=$_POST['next'];

//exit($next);
write_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/stages.txt",$next);

echo "<form name=\"refresh\" method=\"post\" action=\"/pages/stages.php\"></form>";
echo "<script language=\"JavaScript\">document.forms['refresh'].submit();</script>";
exit();
}
///////////////////////////////////////////////



// считать файл 
$stages=read_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/stages.txt");
$stages=trim($stages);

$chbx="/img/chekbox0.png";
function set_img($num)
{
global $stages;
$img="/img/chekbox0.png";
if($num==$stages)
	$img="/img/chekbox1.png";
return $img;
}
///////////////////////////////////////////////////



echo "<div style='max-width:1000px;'>
<hr style='width:990px;'><h3>Редактор Стадии развития </h3>(stages.txt)<br>";

echo "Каждая стадия развития запрещает операции изменений состояния предыдущих и последующих стадий.<br>
Перейдя на новую стадию развития <b>уже НЕ СЛЕДУЕТ возвращаться к прежним</b> (потому как все изменения последующих стадий основаны на предыдущих).<br>
Поэтому переключение стадий развития является очень ответственной операций и должно выполняться только при тщательном тестировании достигнутого на текущей стадии.<br>
НО это важно только для формирования последовательности безусловных рефлексов и не влияет на схемы Beast у которого будем различать пассивное и инициативные развитие.<br>
Для возвращения к прежним стадиям развития можно использовать сохраненные архивы памяти (шестеренка справа-сверху на Пульте).<br><br>
При нормальной плотности общения и воспитания указаны приблизительные сроки перехода на следующую стадию развития – в годах.<hr>
<br>
";


echo "<b>Переключатель стадий развития</b>  :<br>";
?>
<style>
.stages_img
{
cursor:pointer;
}
</style>


<form id="goto1" name="form" method="post" action="/pages/stages.php" >

<div style="position:relative;">
<div style="position:absolute;top:-4px;right:0px;"><a href="https://scorcher.ru/beast/main_help.htm#_Toc122883048" target="_blank">О стадии № 0</a></div>
<b>0</b> <img class="stages_img" src="<?echo set_img(0);?>" onClick="goto_stages(0)"> - Подготовка к рождению: прошивка наследственных особенностей:</div>
<li><a href="/pages/gomeostaz.php">РедакторФормирование жизненных параметров гомеостаза</a></li>
<li><a href="/pages/words.php">Заливка Фраз для дерева слов и дерева фраз</a></li>
<li><a href="/pages/terminal_actions.php">Редактор возможных Действий</a></li>
<li><a href="/pages/reflexes.php">Редактор безусловных рефлексов</a> 
<span style="font-size:16px;cursor:pointer;color:#7E58FF;" onClick="cliner_reflex_times_stadies()"  title="Чтобы рефлексы не просрочили свое время жизни, нужно обновлять его перед началом использования." ><b>Обновить время жизни рефлексов</b></span></li>
В общем-то эта стадия уже предложена в головом виде и можно просто перейти к следубщей, а можно изменить базовую сущность этого живого существа.
<br><br>
<div style="position:relative;">

<div style="position:absolute;top:-4px;right:0px;"><a href="https://scorcher.ru/beast/main_help.htm#_Toc122883053" target="_blank">О стадии № 1</a></div>
<b>1</b> <img class="stages_img" src="<?echo set_img(1);?>" onClick="goto_stages(1)"> - Рождение Beast. Формирование набора условных рефлексов. Период взаимодействия с Beast любым образом, с разными сочетаниями действий и очень простых фраз в различных состояниях его Базовых параметров (для этой стадии развития можно устанавливать параметры гомеостаза слайдерами Пульта для тестирования).</div>
Чем дольше этот период, тем более эффективные навыки получит Beast.

<br><br>
<div style="position:relative;">

<div style="position:absolute;top:-4px;right:0px;"><a href="https://scorcher.ru/beast/main_help.htm#_Toc122883054" target="_blank">О стадии № 2</a></div>
<b>2</b> <img class="stages_img" src="<?echo set_img(2);?>" onClick="goto_stages(2)"> - Формирование базовых автоматизмов через отзеркаливание чужих действий и случайные попытки связать их между собой. Взаимодействия с Beast любым образом, с разными сочетаниями действий и очень простых фраз, предваряя фразы действиями и бездействием, в различных состояниях его Базовых параметров (для этой стадии развития можно устанавливать параметры гомеостаза слайдерами Пульта для тестирования). Так же Beast будет экспериментировать с простыми действиями и фразами.<br></div>
Чем дольше этот период, тем более эффективные навыки получит Beast.

<br><br>
<div style="position:relative;">

<div style="position:absolute;top:-4px;right:0px;"><a href="https://scorcher.ru/beast/main_help.htm#_Toc122883055" target="_blank">О стадии № 3</a></div>
<b>3</b> <img class="stages_img" src="<?echo set_img(3);?>" onClick="goto_stages(3)"> - Формирование базовых автоматизмов через наблюдения за чужими действиями с выявлением причинно-следственных связей в наблюдаемых диалогах в виде учительских автоматизмов.</div>
Чем дольше этот период, тем более эффектиные навыки получит Beast. 
<br><br>
<div style="position:relative;">

<div style="position:absolute;top:-4px;right:0px;"><a href="https://scorcher.ru/beast/main_help.htm#_Toc785941235" target="_blank">О стадии № 4</a></div>
<b>4</b> <img class="stages_img" src="<?echo set_img(4);?>" onClick="goto_stages(4)"> - Начало работы дерева понимания или дерева ментальных автоматизмов. Накопление внутренних автоматизмов мышления и произвольности.<br>
Период преступной инициативы. Продолжается интенсивное и продолжительное общение, но Beast начнет подвергать сомнению то, что раньше принималось безусловно и осмысливать результаты таких действий.</div>
Чем дольше этот период, тем более эффектиные навыки получит Beast.
<br><br>
<div style="position:relative;">

<div style="position:absolute;top:-4px;right:0px;"><a href="https://scorcher.ru/beast/main_help.htm#_Toc297531259" target="_blank">О стадии № 5</a></div>
<b>5</b> <img class="stages_img" src="<?echo set_img(5);?>" onClick="goto_stages(5)"> - Инициативное и творческое развитие Beast в среде общения. </div>
<br><br>


<br><br>
<input type='hidden' name='gogogo' value='1'>
<input type='hidden' name='next' value=''>
</form>
</div>
<!-- hr><br>
<b>Вернуться на предыдущую стадию развития - <span style='color:red'>Будет удалено все, что сформировалось после стадии возврата!</span></b --><br>
<br>
<br>
<br>
<br>
<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var linking_address = '<? include($_SERVER["DOCUMENT_ROOT"] . "/common/linking_address.txt"); ?>';


var c_timer=0;
var bot_is_connected=0;
function check_activnost()
{
	bot_is_connected=0;
var AJAX = new ajax_support(linking_address + "?check_Beast_activnost=1", check_conn_info);
AJAX.send_reqest();
function check_conn_info(res) {
	bot_is_connected=1;
	c_timer=setTimeout("check_activnost()",1000);  //show_dlg_alert("!!!!",0);
}
c_timer=setTimeout("check_activnost()",1000);
}
check_activnost();



/////////////////////////////
var stages='<?=$stages?>';
var next_level=0;
function goto_stages(next)
{
if(bot_is_connected)
{
show_dlg_alert("Сначала выключите Beast.",0);
return;
}
next_level=next;   
if(next<stages)
{
//show_dlg_confirm("Точно вернуться к предыдущему уровню развития? Будет удалена память, зависимая от стадии До рождения.","ДА","НЕТ",gotonextleve0);
show_dlg_confirm("Очистить память, зависимую от стадии &quot;"+next_level+"&quot;?<br>Или вернуться к предыдущему уровню развития без очистки?","Да, очистить","Вернуться без очистки",gotonextlevel);
return;
}
if(next>stages)
{
	if(next - stages > 1){
		var next_stag = stages
		next_stag ++
show_dlg_alert("Нельзя пропускать стадии развития!</br>Следующей стадией должна быть: " + next_stag + "!",0);
		return;
	}
gotonextlevel2();
}
/////////////////
if(next==stages)
{
	return;
}
//gotonextlevel(next);
}
///////////////////////
function closed_dlg_confirm()// если выбрано "Вернуться без очистки"
{ 
//	alert("1");
gotonextlevel2()
}

function gotonextlevel()
{
// в зависимости от того, куда возвращаемся
var AJAX = new ajax_support("/lib/cliner_stadies_memory.php?next_level="+next_level, sent_cliner_reflex_memory);
AJAX.send_reqest();
function sent_cliner_reflex_memory(res) {
show_dlg_alert("Память, зависимая от стадии "+next_level+", очищена.",2000);
setTimeout("gotonextlevel2();",2000);
}
}
function gotonextlevel2()
{
var form=document.getElementById("goto1");
form.next.value=next_level; 
form.submit();
next_level=0;
}

var cliner_reflex_timer=0;
function cliner_reflex_times_stadies()
{
if(!bot_is_connected)
{
show_dlg_alert("Включите Beast для обновления времени жизни рефлексов.",0);
return;
}
 cliner_reflex_timer=setTimeout("cliner_reflex_times_stadies2()",2000);
var AJAX = new ajax_support(linking_address + "?cliner_time_condition_reflex=1", sent_get_info);
AJAX.send_reqest();
function sent_get_info(res) 
{
show_dlg_alert('Время жизни условных рефлексов очищено.',2000);
clearTimeout(cliner_reflex_timer);
}
}
function cliner_reflex_times_stadies2()
{
show_dlg_alert('Функция не срабатывает, сначала нужно включить Beast.',0);
}
</script>
<?
////////////////////////////////////////////////////////





/////////////////////////////////////////////////
?>
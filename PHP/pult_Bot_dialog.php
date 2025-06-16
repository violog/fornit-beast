<?
/*  диалог общения с Beast
 include_once($_SERVER['DOCUMENT_ROOT']."/pult_Bot_dialog.php");
*/

echo "<hr>";
echo "<div id='stimul_action_id' style='background-color:#ffffff;'>";

echo "<div style='position:relative;margin-top:10px;height:18px;'>";
echo '<div id="about_bot_ready" style="position:absolute;left:0;top:0;color:red;">Beast еще не пришел в себя, общение невозможно.</div>';
echo '
<div id="diff_condition_str" style="position:absolute;right:150px;top:2;display:none;" title="Изменение состояния после ответа Оператора на действия Beast.">Состояние после действия оператора: </div>
<div id="diff_condition_id" style="position:absolute;right:0;top:0;" title="Изменение состояния после ответа Оператора на действия Beast."></div>';
echo "</div>";



if($stages>0)
{
echo '<div style="margin-top:10px;color:;">По возможности предваряйте фразы подходящими действиями, сначала нажав кнопку (или несколько подряд) внизу.</div>';
}


?>
<div style="position:relative;border:solid 1px #8A3CA4;background-color:#cccccc;width:1000px;padding:10px;
margin-top:10px;">

<b>Поcлать сообщение Beast</b>: <span id="stadia_warn" style="color:red;"></span><br>
<div style='position:absolute;top:15px;right:10px;' title='Режим форсированного формирования вербальных распознавателей (без отсеивания мусорных слов) и условных рефлексов.'><nobr><input type="checkbox" value="1" onChange="switch_input_rejim(this)"> - режим форсированной обработки</nobr></div>

<div id="conditions_words_id" style='position:absolute;top:4px;left:250px;display:none;' title='Список слов, для которых есть условный рефлекс в этих условиях.'><img src='/img/words.png' onClick='get_conditions_words()' ></div>

<div id="basic_words_id" style='position:absolute;top:4px;left:300px;display:none;' title='Список фраз, для которых есть автоматизм Beast в этих условиях.'><img src='/img/words.png' onClick='get_conditions_words_basic()' ></div>


<div id="note_rejim_id" style='position:absolute;top:0px;right:0px;color:red;display:none;cursor:pointer;' title='Режим форсированного формирования вербальных распознавателей (без отсеивания мусорных слов) и условных рефлексов.' onClick="show_dlg_alert('<div style=`font-size:14px;font-weight:200;text-align:left;`>Включен режим форсированного формирования вербальных распознавателей (без отсеивания мусорных слов) и условных рефлексов.<br><br>Следует использовать для набивки фраз словарного запаса, для тестирования условных рефлексов и т.п. когда не требуется отсев случайного путем повторений.',0)"><nobr>Включен режим форсированной обработки.</nobr></div>
<script>
var is_input_rejim=1;
function switch_input_rejim(ch)
{
	if(ch.checked==true)
	{
document.getElementById('note_rejim_id').style.display="block";
is_input_rejim=0;
	}
	else
	{
document.getElementById('note_rejim_id').style.display="none";
is_input_rejim=1;
	}
}
function mood_cliner()
{
document.getElementById('radio_2').checked=true;
document.getElementById('radio2_4').checked=true;
}
</script>

<div style="position:relative;">
<div style="position:absolute;top:10px;left:-20px;color:red;cursor:pointer;padding:4px;border:solid 1px #8A3CA4;border-radius:50%;background-color:#ffffff" title="Очистить окно ввода" onClick="cliner_textarea()"><b>X</b></div>
<textarea id="input_id"  style="width:calc(100% - 10px);margin-top:10px;" rows="6" maxlength="500" onMouseDown="click_textarea()" onKeyDown="click_textarea()" disabled>Привет</textarea><br>
<b>Тон:</b> 
<input id='radio_1' type='radio' name='rdi' value='4' >повышенный 
<input id="radio_2" type='radio' name='rdi' value='0' checked>нормальный 
<input id="radio_3" type='radio' name='rdi' value='3'>вялый 

&nbsp;&nbsp;<b>Передать контекст своего настроения:</b><br>

<input id='radio2_4' type='radio' name='rdi2' value='0' checked>Нормальное &nbsp;
<?
// получить эффекты от нажатий чтобы показать их в titles
$effectsArr=array();// строки эффектив
$progs = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/moode_dialog_effects.txt");
$strArr = explode("\r\n", $progs);
$n=5;
foreach ($strArr as $str) {
	$par = explode("|", $str);
	$id = $par[0];
	$e=$par[1];
	$znak=substr($e,0,1);
	$effect="".$znak.substr($e,1);

$bgcolor="";
	if($znak=="+")
		$bgcolor="#9FF0B3";
	if($znak=="-")
		$bgcolor="#FFC9D0";

echo "<span style='background-color:".$bgcolor.";padding-right:2px;' title='Мотивационный эффект: ".$effect."'><input id='radio2_".$n."' type='radio' name='rdi2' value='".$id."'>".$par[2]."</span> &nbsp;";
$n++;
} 
?>

<input id="reset_button_id" type="button"  value="Сброс" onClick="reset_go()" title="Разорвать цепочку кадров ЭП" style="position:absolute;bottom:-8px;right:80px;padding:4px;" disabled> 
<input id="input_button_id" type="button"  value="Послать" onClick="sent_go()" style="position:absolute;bottom:-8px;right:0px;padding:4px;" disabled> 
</div>
<div id='recognized_block_id' style="position:relative;margin-top:10px;background-color:#eeeeee;padding:4px;min-height:40px;">

<div style="position:absolute;top:-48px;left:750px;font-size:12px;cursor:pointer;
line-height: 0.8;padding-left:3px;padding-right:3px;border-radius: 7px;
border:solid 1px #8A3CA4;background-color:#efefef;" onClick="mood_cliner()" title="Очистить Тон и настроение."><br><img src='/img/delete.gif'><br>&larr;<br><br></div>


<!-- div style="position:absolute;top:4px;right:170px;font-size:12px;"><nobr><a href='/pages/words_tree.php' target='showpage2' style='position:absolute;top:0px;right:0px;'>Дерево слов</a></nobr></div>
<div style="position:absolute;top:4px;right:80px;font-size:12px;"><nobr><a href='/pages/phrase_tree.php' target='showpage2' style='position:absolute;top:0px;right:0px;'>Дерево фраз</a></nobr></div>
<div style="position:absolute;top:4px;right:4px;font-size:12px;cursor:pointer;" onClick="tree_cliner()" title="Очистить дерево слов и фраз чтобы начать заново.">- очистить</div -->

<span style="color:#666666;font-size:15px;">Распознаное:</span><br><span id="pult_result_id" style="margin-top:10px;height:20px;"></span>

<div style="position:absolute;top:-47px;right:10px;color:green;cursor:pointer;"  onClick='explane_sending()' title="По возможности предваряйте фразы подходящими действиями, сначала нажав кнопку (или несколько подряд) внизу."><b>Как добавить действия</b></div>
</div>

</div>

<?
//////// индикация периода ожидания реакции оператора на действие автоматизма
// "Осталось времени на ответ:"
echo "<div style='position:relative;'>";
echo "<div id='time_limit_id' class='luminous_text' style='position:absolute;z-index:1000;top:-10px;left:50%;transform: translate(-50%, 0);
background-color:#ffffff;border-radius: 7px;
padding:3px;
font-size: 18px;color:#BF0000;cursor:pointer;
display: none;
'
title='Завершить по щелчку.' onClick='stop_waiting_period()'></div>";// <nobr>Осталось времени на ответ:  сек</nobr>
echo "</div>";



echo "</div>";//<div id='stimul_action_id'>
?>

<script Language="JavaScript" src="/ajax/ajax_post.js"></script>
<script>
function stop_waiting_period()// завершить период ожидания по щелчку на плашке
{
var AJAX = new ajax_support(linking_address + "?stop_waiting_period=1", sent_stop_waiting_period); 
AJAX.send_reqest();
function sent_stop_waiting_period(res) {
document.getElementById('time_limit_id').style.display="none";
old_period_val=0;
allowShowWaightStr=0;
}
}
////////////////////////////////
function explane_sending()
{
show_dlg_alert("Чтобы добавить действия (кнопки внизу) нужно просто нащелкать нужные (не нажимая треугольнички) тогда эти действия добавятся в образ сообщения.",0);
}

function tree_cliner()
{
	show_dlg_confirm("Вам придется заново набивать фразы.<br>Точно очистить детектор фраз?",1,1,cliner_continue);

}
function cliner_continue()
{
var AJAX = new ajax_support("/lib/tree_cliner_server.php",sent_tree_cliner);
AJAX.send_reqest();
function sent_tree_cliner(res)
{
show_dlg_alert("Деревья слов и фраз очищены.",0);
}
}
function cliner_textarea()
{
end_dlg_alert();
//end_ReceiveAnsvetFromPult();
document.getElementById('input_id').value="";
document.getElementById('input_id').focus();
}

function reset_go(){
	var AJAX = new ajax_support(linking_address + "?stop_waiting_period=1", sent_stop_waiting_period); 
	AJAX.send_reqest();
	function sent_stop_waiting_period(res) {
		document.getElementById('time_limit_id').style.display="none";
		old_period_val=0;
		allowShowWaightStr=0;
		show_dlg_alert("Цепочка кадров ЭП прервана.",0);
}
}

function sent_go()
{
end_dlg_bot_action(); //alert("!!!!");

var txt=document.getElementById('input_id').value;
if (txt.length==0)
{
show_dlg_alert("Пустое сообщение.",1500);
	return;
}

var tone=0;
var moode=0;
var allr=document.getElementsByName('rdi');
for(var i=0; i<allr.length; i++)
{
    if (allr[i].checked) 
	{
		tone=allr[i].value;
      break; 
	}
 }
//alert(tone);
allr=document.getElementsByName('rdi2');
for(var i=0; i<allr.length; i++)
{
    if (allr[i].checked) 
	{
		moode=allr[i].value;
      break; 
	}
 }
//alert(moode);

///////////////////////////////////////////
var server="/lib/separate_words_str_server.php";
var params="text="+txt;    //alert(txt);return;
var AJAX = new ajax_post_support(server,params,sent_request_bot,1);
AJAX.send_reqest();
function sent_request_bot(res)
{
// если выбраны действия для добавления
//alert(allow_sent_to_beast);
params="is_input_rejim="+is_input_rejim+"&pult_tone="+tone+"&pult_mood="+moode+"&text_dlg="+res;

var triggers_str="";
if(allow_sent_to_beast)
{		
// получить список пусковых стимулов (class='actions', но есть actionsArr)
for(i=0;i<actionsArr.length;i++)
	{
triggers_str+=actionsArr[i]+"|";//! нельзя разделять ; или ,
	}
var food_portion = document.getElementById("food_portion_id").selectedIndex + 1;
params+="&set_img_action=" + triggers_str + "&food_portion=" + food_portion;

//addInfoToHistory(3,""+triggers_str);

desactivationAll();
}

var str=tone+"||"+moode+"||"+res+"||"+triggers_str;
addInfoToHistory(2,str); // в /index.php

//alert(params);return;
bot_contact(params,text_bot_answer);
document.getElementById('input_id').value = "";
}
}

function text_bot_answer(res)
{ 
//	alert(res);
if(res=="POST")
{
show_dlg_alert('Не воспринят текст...',3000);
//setTimeout(`document.forms['refresh'].submit();`,2000);
return;
}

var par=res.split("|&|");   // alert(par[0]+" | "+par[1]);

document.getElementById('pult_result_id').innerHTML=par[0];

if(par[1].length>5)// уже готов ответ Beast в том же пульсе
			{ //alert(par[1]);
// выдать его на Пульт
new_bot_action(par[1]);
			}

//show_dlg_alert(res,2000);
//setTimeout(`document.forms['refresh'].submit();`,2000);
}
////////////////////////////
function click_textarea()
{
end_dlg_alert();
//end_ReceiveAnsvetFromPult();
}
/* при каждом изменении условий проверка заготовленных в http://go/pages/condition_reflexes_basic_phrases.php слов - условных рефлексов.
*/
var cur_conditions_words="";
var old_contexts="";
var curBasicPar="";
var curContextsPar=""
function check_cur_conditions_words(basic,contexts)
{
//	alert(old_contexts+" | "+contexts);
	if(old_contexts==contexts)
		return;
old_contexts=contexts;
curBasicPar=basic;
curContextsPar=contexts;

//alert(basic+" | "+contexts);exit();
var AJAX = new ajax_support("/lib/get_exclamations_for_conditions.php?basic="+basic+"&contexts="+contexts, sent_ch_words_info);
AJAX.send_reqest();
function sent_ch_words_info(res) {
//show_dlg_alert(res, 0);
if(res.length>0)
{
cur_conditions_words=res;
document.getElementById('conditions_words_id').style.display="block";
document.getElementById('basic_words_id').style.display="block";
}
else
{ // не гасить
//cur_conditions_words="";
//document.getElementById('conditions_words_id').style.display="none";
//document.getElementById('basic_words_id').style.display="none";
}
}
}
function get_conditions_words()
{
show_dlg_alert2("<br><span style='font-weight:normal;'>Щелкните по слову:<br>" + cur_conditions_words + "<br>", 0);
}
/////////////////////////////
function get_conditions_words_basic()
{
	// разрешать только после готовности ГО
	if(beast_ready!=2)
	{
show_dlg_alert("Будет доступно когда beast будет готов к общению", 0);
return;
	}

//	alert(curBasicPar+" || "+curContextsPar);return; //  1 || 1,2,9
var AJAX = new ajax_support(linking_address + "?conditions_words_basic=1&basicID="+curBasicPar+"&contexts="+curContextsPar, sent_words_basic_info); 
AJAX.send_reqest();
function sent_words_basic_info(res) {
			//alert(res);
show_dlg_alert2("<br><span style='font-weight:normal;'>" + res + "<br>", 2);
}

}
//////////////////////////////
function insert_pult_word(word)
{
	end_dlg_alert2();
	document.getElementById('input_id').value=word;
}

</script>
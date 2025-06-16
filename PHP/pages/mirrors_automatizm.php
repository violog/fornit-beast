<?
/* набивальщик базы зеркальных автоматизмов 

Сохранение - в /pages/mirrors_automatizm_saver.php

http://go/pages/mirrors_automatizm.php

*/
$page_id = -1;
$title = "Создание базы зеркальных автоматизмов";
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");
//include_once($_SERVER['DOCUMENT_ROOT']."/pult_js.php");
//////////////////////////////////////////////////////////////

$bsID=0;
if(isset($_GET['bsID']))
$bsID=$_GET['bsID'];            //  exit("> $bsID");

$id_list="";
$get_list="";
if(isset($_GET['id_list']))
{
$id_list=$_GET['id_list'];

$get_list=explode(";",$id_list); 
}


?>
<div style="position:absolute;top:35px;right:10px;color:red"> - только после формирования шаблонов условных рефлексов</div>


<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var is_table_shoved=0;// 1 таблица показывается
function get_table(kind)
{
wait_show();
var link="/pages/mirrors_automatizm_table.php?"+cur_condition_choose+"&kind="+kind;
//alert(link);
var AJAX = new ajax_support(link, sent_table_info);
AJAX.send_reqest();
function sent_table_info(res)
{
	wait_end();
//show_dlg_alert(res,0);
if(res[0]!='!')
{ //alert(res);
show_dlg_alert(res,0);
return;
}
document.getElementById('table_id').innerHTML=res.substr(1);
//document.getElementById('insert_from_common_id').style.display="block";

is_table_shoved=1;
wait_end();
}
}


</script>
<?
//////////////////////////////////////////////////////////////
echo "<div id='hr_table_id' style='position:relative;display:none;'>
<hr>
<div style='position:absolute;top:-10px;left:50%;transform: translate(-50%, 0);background-color:#ffffff;padding-left:10px;padding-right:10px;'><b>Задать условия для таблицы ввода фраз-ответоа</b>
</div>";

echo "<div style='position:absolute;z-index:10;top:20px;right:20px;'><a href='/pages/condition_reflexes_basic_phrases.htm'>Страница пояснений</a></div>";

echo "<div style='position:absolute;top:60px;right:10px;border:solid 1px #8A3CA4;border-radius: 7px;padding:10px;box-shadow: 8px 8px 8px 0px rgba(122,122,122,0.3);background-color:#efefef;max-width:65%;font-size:14px;'>
<b>Поясненния:</b><br>
Колонка &quot;Ответная фраза&quot; - это ответ оператора на фразу в колонке слева в данных условиях (строка над таблицей). Как человек, оператор может ответить, не затрудняясь с подбором, свободно, так что не стоит особенно заморачиваться с Ответной фразой.<br>НО нужно понимать, что эту фразу Beast станет использовать бездумно, полагаясь на авторитет оператора (раз он так отвечает, то и мне можно попробовать). Желательно писать короткие и ясные фразы.<br>

<b>Использование:</b><br>
первым делом следует заполнить Общий шаблон ответов (по ссылке ниже). Затем станет проще набивать таблицы для каждого сочетания Базовых контекстов. После заполнения таблицы следует нажать под ней “Сохранить фразы” или просто нажать Ctrl+S. 

</div>";

echo "<div style='position:absolute;top:240px;;left:450px;font-size:16px;' >Фразы могут быть <b>НЕ уникальны</b> для таблицы.</div>";


echo "<div style='position:absolute;top:270px;right:10px;font-size:18px;cursor:pointer;color:#885CFF' onClick='location.href=`/pages/mirror_basic_phrases_common.php`'>Общий шаблон ответов</div>";

// onChange='choode_base_cond(this)' - нет определенной зависимости...
echo "<b>Базовое состояние:</b><br>
<select id='base_id' onChange='refresh_context_combo(this)'> 
<option value='1' "; if($bsID==1)echo "selected"; echo ">Плохо</option>
<option value='2' "; if($bsID==2)echo "selected"; echo ">Норма</option>
<option value='3' "; if($bsID==3)echo "selected"; echo ">Хорошо</option>
</select><span title='Общее Базовое состояние формируется из отдельных состояний Базовых параметров гомеостаза и при этом никак не коррелирует с диапазонами состояний параметров гомеостаза.'> - Общее Базовое состояние</span><br>
";

//exit("> $bsID");






// Базовые контексты $baseContextArr
include_once($_SERVER['DOCUMENT_ROOT'] . "/lib/base_context_list.php");

/* Все возможные сочетания активных контекстов выбираются из таблицы "Активности Базовых стилей" (минуса игнорируются и идет проверка на антагонистов).
*/
$contextsArr=array();// ID выбранных контекстов без антагонистов
echo "<b>Выбрать сочетания контекстов:</b><br> 
<div id='context_variations_id'></div>";

?>
<script>
var current_base_condition=1;// при обновлении страницы - Плохо
function refresh_context_combo(combo)
{
current_base_condition=combo.options[combo.selectedIndex].value; //alert(current_base_condition);
get_context_variations(1);
}

function get_context_variations(bc)
{ 
	wait_begin();
	// base_condition="+bc+"& 
//	alert("/pages/reflexes_maker_b_contexts.php?get_list=<?=$id_list?>");
var AJAX = new ajax_support("/pages/reflexes_maker_b_contexts.php?current_base_condition="+current_base_condition+"&get_list=<?=$id_list?>", send_context_variations);
		AJAX.send_reqest();

function send_context_variations(res) 
{ // alert(res);
			wait_end();
if(res[0]!='!')
{ //alert(res);
show_dlg_alert(res,0);
return;
}
res=res.substr(1);
// удалить (число рефлексов)
res=res.replace(/\(\/?[^>]+\)/g, '');

document.getElementById('context_variations_id').innerHTML=res; 
//alert(res.substr(1));

//show_dlg_alert(res.substr(1),0);
document.getElementById('button_table_id').style.display="";
document.getElementById('hr_table_id').style.display="block";

}
}
get_context_variations(1);// сразу показать 
/*
function choode_base_cond(combo)
{
var bc=combo.options[combo.selectedIndex].value; //alert(bc);
get_context_variations(bc);
}
*/
</script>
<?
///////////////////////////////////////////////////////////////////////



echo "<br><input id='button_table_id' type='button' value='Создать таблицу для заполнения фразами' onClick='choose_0()' style='display:none;'>";
//////////////////////////////////////////////////////////////

echo "<div style='position:relative;'>
<hr>
<div style='position:absolute;top:-10px;left:50%;transform: translate(-50%, 0);background-color:#ffffff;padding-left:10px;padding-right:10px;'><b>Таблица для заполнения фразами-ответами</b>
</div>";

//echo "<div id='insert_from_common_id' style='position:absolute;top:-10px;right:10px;background-color:#efefef;font-size:18px;border:solid 1px #8A3CA4;border-radius: 7px;padding-left:10px;padding-right:10px;cursor:pointer;display:none;' onClick='get_table(1)'>Заполнить из общего шаблона</div>";
//var_dump($contextsArr);exit();
//////////////////////////////////////////////////////////////

echo "<div id='conditions_block_id' style='position:relative;display:none'>";
echo "<div style='position:absolute;top:10px;right:0px;'><span onClick='prases_saver()' style='color:#AE55FF;cursor:pointer;font-size:18px;'>Сохранение фраз</span> - по <b>Ctrl+S</b></div>";
echo "<b>Выбранные условия:</b><br>";

echo "Базовое состояние: <b><span id='base_cond_id'></span></b>";
echo "</b>&nbsp;&nbsp;&nbsp;&nbsp;Сочетания контекстов: <b><span id='base_context_name'></span></b>";

echo "</div>";


$ton_moode_dlg="<br><div style='text-align:left;'>
<b>Тон:</b> &nbsp; 
<input id='radio_0' type='radio' name='rdi' value='0' checked>0 нормальный &nbsp;
<input id='radio_1' type='radio' name='rdi' value='1'>1 вялый &nbsp;
<input id='radio_2' type='radio' name='rdi' value='2' >2 повышенный <br>

<br><b>Настроение:</b> &nbsp;
<input id='radio2_0' type='radio' name='rdi2' value='0' checked>0 Нормальное &nbsp;
<input id='radio2_1' type='radio' name='rdi2' value='1' >1 Хорошее &nbsp;
<input id='radio2_2' type='radio' name='rdi2' value='2' >2 Плохое &nbsp;
<input id='radio2_3' type='radio' name='rdi2' value='3'>3 Игровое &nbsp;
<input id='radio2_4' type='radio' name='rdi2' value='4'>4 Учитель &nbsp;
<input id='radio2_5' type='radio' name='rdi2' value='5'>5 Агрессивное&nbsp;
<input id='radio2_6' type='radio' name='rdi2' value='6'>6 Защитное &nbsp;
<input id='radio2_7' type='radio' name='rdi2' value='7'>7 Протест </div>";


//////////////////////////////////////// ТАБЛИЦА
echo "<div id='table_id'></div>";




//////////////////////////////////////////////////////////////
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/alert2_dlg.php");
?>
<script Language="JavaScript" src="/ajax/ajax_post.js"></script>
<script>
var cur_bcond_choose="";
var cur_bcontex_choose="";
var cur_condition_choose="";
function choose_0()
{
var bsID=document.getElementById('base_id').selectedIndex +1;
switch(bsID)
{
case 1: cur_bcond_choose="Плохо"; break;
case 2: cur_bcond_choose="Норма"; break;
case 3: cur_bcond_choose="Хорошо"; break;
}
//alert(bsID);

var combo=document.getElementById('base_context_id'); 
//alert(combo.selectedIndex);
if(combo.selectedIndex==-1)
	{
	show_dlg_alert("Нужно выбрать сочетание Базовых контекстов в списке.",0);
	return;
	}
var id_list=combo.options[combo.selectedIndex].value; // 1;3
cur_bcontex_choose=combo.options[combo.selectedIndex].text;


cur_condition_choose='bsID='+bsID+'&id_list='+id_list; 
// alert(cur_condition_choose);

document.getElementById('conditions_block_id').style.display="block";
document.getElementById('base_cond_id').innerHTML=cur_bcond_choose;
document.getElementById('base_context_name').innerHTML=cur_bcontex_choose;

get_table(0);


//location.href='/pages/mirrors_automatizm.php?bsID='+bsID+'&id_list='+id_list;
}
//////////////////////////////////////////////

function prases_saver()
{
//	alert("ЗАПИСЬ");	return;
var saveStr="";
var tr =0;
var nodes = document.getElementsByClassName('r_table'); // alert(nodes.length);
for(var i=0; i<nodes.length; i++) 
{
tr=nodes[i]; //alert(tr.cells[0].innerHTML+"|"+tr.cells[3].childNodes[0].value);return;

// пропускаем все, что не содержит фраз и действий
if(tr.cells[1].childNodes[0].value==0 && tr.cells[3].childNodes[0].value==0)
	continue;

//alert(tr.cells[0].innerHTML+"|"+tr.cells[3].childNodes[0].value);return;
saveStr+=tr.cells[0].childNodes[0].value+"|"+tr.cells[1].childNodes[0].value+"|"+tr.cells[2].childNodes[0].value+"|"+tr.cells[3].childNodes[0].value+"|||";
}//for(
//  alert(saveStr);return;

if(saveStr.length==0)
{
show_dlg_alert("Нет фраз-ответов.",2000);
return;
}
//alert(saveStr); // return;
/////////////////////////
saveStr=cur_condition_choose+"&saveStr="+saveStr;  // alert(saveStr);return;
var link="/pages/mirrors_automatizm_saver.php";
//alert(link);
var AJAX = new ajax_post_support(link,saveStr, sent_table_save,1);
AJAX.send_reqest();
function sent_table_save(res)
{
//show_dlg_alert(res,0);
if(res[0]!='!')
{ //alert(res);
show_dlg_alert(res,0);
return;
}
show_dlg_alert("Записаны новые фразы.",2000);
// перегрузить таблицу чтобы показывало дубликаты фраз
get_table(0);
}
}
////////////////////////////////

/////////////////////////////////////////
	// сохранение по Ctrl+S
var is_press_strl = 0;
document.onkeydown = function(event) { 
	
		var kCode = window.event ? window.event.keyCode : (event.keyCode ? event.keyCode : (event.which ? event.which : null))

		//alert(kCode);
		if (kCode == 17) // ctrl
			is_press_strl = 1;

		if (is_press_strl) {
			if (kCode == 83) {
				event.preventDefault();
//alert(is_table_shoved);
	if(!is_table_shoved)
		return false;
				//alert("!!!!! ");
				reflex_saver();
				is_press_strl = 0;
				return false;
			}
		}
}
////////////////////
document.onmouseup = function(event) { 
var t = event.target || event.srcElement;    
while(t)
{ 
if(t.id == "div_dlg_alert2")
	 return;	    
t = t.offsetParent;
}
end_dlg_alert2(); 
}
///////////////////////////
function reflex_saver()
{
show_dlg_confirm("Сохранить список?",1,-1,prases_saver);
}
function close_all_dlg()
{
// просто чтобы была такая пустая и не было варннига при закрытии по фону
}
	////////////////////////////
function set_sel(tr, id) {
		//	alert(id);
		var nodes = document.getElementsByClassName('highlighting'); //alert(nodes.length);
		for (var i = 0; i < nodes.length; i++) {
			nodes[i].style.border = "solid 1px #000000";
		}
		tr.style.border = "solid 2px #000000";
}



////////////////////////////////////////
// сработает перед закрытием show_dlg_confirm
function onw_dlg_exit_proc()
{
//alert("onw_dlg_exit_proc");
var allr=document.getElementsByName('rdi'); //alert(allr.length);
var ton=0;
for(var i=0; i<allr.length; i++)
{
    if (allr[i].checked) 
	{
		ton=allr[i].value;
      break; 
	}
}
var allr=document.getElementsByName('rdi2');
var moode=0;
for(var i=0; i<allr.length; i++)
{
    if (allr[i].checked) 
	{
		moode=allr[i].value;
      break; 
	}
}
var inp=document.getElementById('insert_'+cur_ton_moode_id).value=ton+","+moode;

}
var cur_ton_moode_id=0;
function show_ton_mood(id)
{  
cur_ton_moode_id=id;
var cont=`<?=$ton_moode_dlg?>`;
is_onw_dlg_exit_proc=1; //предопределенная переменнная
show_dlg_alert("<br><span style='font-weight:normal;'>Выберите Тон и настроение:<br>" + cont + "<br>", 0);
// проставить выделение старого выбора
var inp=document.getElementById('insert_'+cur_ton_moode_id).value;
if(inp.length>0)
{
var tm=inp.split(","); //alert(tm[0]+" | "+tm[1]+" || "+tm.length);
  //alert(document.getElementById('radio_2').checked);
document.getElementById('radio_'+tm[0]).checked=true;
document.getElementById('radio2_'+tm[1]).checked=true;
}

//is_onw_dlg_exit_proc=0 - при закрытии само очистится
}
////////////////////////////////
function show_actions_list(nid)
{
event.stopPropagation();
var selected=document.getElementById("insert2_" + nid).value;
//show_dlg_alert(selected,0);
event.stopPropagation();
		var AJAX = new ajax_support("/lib/get_actions_list.php?selected="+selected, sent_act_info);
		AJAX.send_reqest();

		function sent_act_info(res) {
			show_dlg_alert2("<br><span style='font-weight:normal;'>Выберите значения:<br>" + res + "<br><input type='button' value='Выбрать значения' onClick='set_input_list("+nid + ")'>", 2);
		}
}
/////////////////////////////
function set_input_list(nid) {
var aStr = "";
var nodes = document.getElementsByClassName('chbx_identiser'); //alert(nodes.length);
for(var i=0; i<nodes.length; i++) 
{
if(nodes[i].checked)
	{
if(aStr.length > 0)
	aStr += ",";
aStr += nodes[i].value;
	}
}
		//alert(aStr);
		document.getElementById("insert2_" + nid).value = aStr;

		end_dlg_alert2();
}
/////////////////////////////////////////
</script>


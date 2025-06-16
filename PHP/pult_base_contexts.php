<?
/*  Индикация базовых контекстов
include_once($_SERVER['DOCUMENT_ROOT']."/pult_base_contexts.php");
*/

?>
<script Language="JavaScript" src="/ajax/ajax.js"></script>
<br>
<div style="position:relative;">

<!-- div id="game_mode_id" class='context' style="position:absolute;z-index:10;top:60px;left:400px;
width:100px;background-color:#EEEEEE;text-align:center;
cursor:pointer;" title="Действия в этом режиме не дают гомео-эффект, все как бы понарошку, но Наказать и Поощрить работают." onClick="game_mode_switcher()">Режим игры</div -->
<div id="game_mode_id"  style="position:absolute;z-index:10;top:60px;left:400px;
text-align:center;background-color:#D3FFF2;font-size:16px;padding:4px;" title="Действия в этом режиме не дают гомео-эффект, все как бы понарошку, но Наказать и Поощрить работают."></div>
<div id='game_mode_exit_id' class='alert_exit' style='position:absolute;z-index:10;top:62px;left:380px;display:none;' title='Снять игровой режим' onClick='end_game_mode();'><span >&#10006;</span></div>


<div style="margin-bottom:10px;"><b>Состояние базовых контекстов:</b> <span id="contect_list_id"></span></div>
<div id='context_1' class='context action_poz4' >Пищевой</div>
<div id='context_2' class='context action_poz4' >Поиск</div>
<div id='context_3' class='context action_poz4' >Игра</div>
<div id='context_4' class='context action_poz4' >Гон</div>
<div id='context_5' class='context action_poz4' >Защита</div>
<div id='context_6' class='context action_poz4' >Лень</div>
<div id='context_7' class='context action_poz4' >Ступор</div>
<div id='context_8' class='context action_poz4' >Страх</div>
<div id='context_9' class='context action_poz4' >Агрессия</div>
<div id='context_10' class='context action_poz4' >Злость</div>
<div id='context_11' class='context action_poz4' >Доброта</div>
<div id='context_12' class='context action_poz4' >Сон</div>

<div style="position:absolute;top:0px;right:0px;color:#66B15C;"><b>
<?
switch($stages) {
case 0:
echo "Нулевая стадия развития: до рождения.";
	break;
case 1:
echo "ПЕРВАЯ стадия развития: рождение Beast, условные рефлексы.";
	break;
case 2:
echo "ВТОРАЯ стадия развития: Формирование базовых автоматизмов.";
	break;
case 3:
echo "ТРЕТЬЯ стадия развития: Период подражания.";
	break;
case 4:
echo "ЧЕТВЕРТАЯ стадия развития: Период преступной инициативы.";
	break;
case 5:
echo "ПЯТАЯ стадия развития: Инициативное и творческое развитие.";
	break;
}
?>
</b></div>


<div id='condition_button_id' style='position:absolute;bottom:-10px;right:0px;
background-color:#FFFFCC;
padding-left:5px;padding-right:5px;
border:solid 1px #8A3CA4;
border-radius: 7px;
text-align:center;
cursor:pointer;
display:none;'
title='Вывести список текущих условий в верхний-правый угол страницы в желтом блоке.'
onClick='show_condition_info()'>Текущие<br>условия</div>
</div>

<script>

var cf_color = {
1:"#E4FFEB",
2:"#FFE8FF",
3:"#FFE8E8",
4:"#FFC9D0",
5:"#CCC9FF",
6:"#D0CCD0",
7:"#FFFFD3",
8:"#FFC2C5",
9:"#FFA7DD",
10:"#FFA7DD",
11:"#C5F5FF",
12:"#D0CCD0"
}; 

for(i=1;i<13;i++)
{
document.getElementById("context_"+i).style.backgroundColor=cf_color[i];
}


function get_context_info(i_str)// в /pult_gomeo.php
{ 
//alert(i_str);
var pars=i_str.split("|");    //alert(pars[1]);
//alert("! "+pars[0]+" ! "+pars[1]+" ! "+pars[2]); 
for(i=0;i<pars.length;i++) // bcntx_1
{
if(pars[i].length<3)
	continue;
var p=pars[i].split(";");
var id=p[0];
var val=p[1];
set_active_id(id,val);
}
}

function set_active_id(id,set)
{  
	if(!document.getElementById("context_"+id))
		return;
//	if(id==1)alert(id+" | "+set);
if(set==1)
document.getElementById("context_"+id).style.boxShadow="0px 0px 6px 6px "+LightenDarkenColor(cf_color[id],-20);
else
if(document.getElementById("context_"+id))
document.getElementById("context_"+id).style.boxShadow="";
} 
//set_active_id(4,1);

function show_condition_info()
{
	event.stopPropagation();

	if(current_condition.length<3)
	{alert('Не передана информация по текущим условиям.');return;}


document.getElementById("helper_pult_id").style.display = "block";
document.getElementById("helper_pult_id").innerHTML = "<div class='alert_exit' style='top:0; right:0;' title='закрыть' onClick='end_helper_dlg();'><span style='position:relative; top:-1px; left:1px;'>&#10006;</span></div><br>"+ful_condition_str;
/*
var AJAX = new ajax_support("/lib/condition_info_translate.php?current_condition="+current_condition, sent_condition_info);
AJAX.send_reqest();
//alert("1");
setTimeout("rebooting()",1000);// выждать завершения процессов
function sent_condition_info(res)
	{  
document.getElementById("helper_pult_id").style.display = "block";
document.getElementById("helper_pult_id").innerHTML = "<div class='alert_exit' style='top:0; right:0;' title='закрыть' onClick='end_helper_dlg();'><span style='position:relative; top:-1px; left:1px;'>&#10006;</span></div><br>"+res;
	}
*/
}
function end_helper_dlg()
{
document.getElementById("helper_pult_id").style.display = "none";
}

var is_game_mode=0;
function switchGameMpda(mode)
{
	is_game_mode=mode;
set_game_moda(is_game_mode);
if(is_game_mode==1)
	{
show_dlg_alert2("Включен игровой режим:<br><div style='font-weight:normal;text-align:left;'>Действия с Пульта и Beast не будут изменять гомеопатические параметры.<br>Только кнопки Наказать и Поощрить будут давать непосредственное воздействие на полезность последнего выполненного автоматизма. При этом оператору дается 15 секунд на нажатие кнопок.<br>На стадии развития > 3-й пишутся Правила.</div>",0);
	}
}
/////////////////////////////////////////
var is_teaching_mode=0;
function switchTeachingMpda(mode)
{
	is_teaching_mode=mode;
set_teaching_moda(is_teaching_mode);
if(is_teaching_mode==1)
	{
show_dlg_alert2("Включен учительский режим:<br><div style='font-weight:normal;text-align:left;'>Преимущественное отзеркаливание (инфо-функция 13) на стадиях развития от 4-й.</div>",0);
	}
}
/////////////////////////////////////////
/*
function game_mode_switcher()
{
if (!exists_connect)
{
show_dlg_alert2("Нужно включить Beast.",0);
return;
}

if(is_game_mode==0)
{
	is_game_mode=1;
} 
else
{
	is_game_mode=0;  // context 
}
set_game_moda(is_game_mode);

var AJAX = new ajax_support(linking_address + "?is_game_mode=" + is_game_mode, sent_action);
AJAX.send_reqest();
function sent_action(res) 
{		
	if(is_game_mode==1)
	{
show_dlg_alert2("Включен игровой режим:<br><div style='font-weight:normal;text-align:left;'>Действия с Пульта и Beast не будут изменять гомеопатические параметры.<br>Только кнопки Наказать и Поощрить будут давать непосредственное воздействие на полезность последнего выполненного автоматизма. При этом оператору дается 15 секунд на нажатие кнопок.<br>На стадии развития > 3-й пишутся Правила.</div>",0);
	}
}
}*/

// 120 секунд держится игровой режим без обновлдения по каждому периоду ожидания ответа
var limit_game_moda_time=120000; 
var game_moda_timer_id=0; //id таймера пролонгации игрового режима
function game_moda_prolongate()
{
clearTimeout(game_moda_timer_id);
game_moda_timer_id = setTimeout("end_game_mode()",limit_game_moda_time);
}
function set_game_moda(isGame)
{
game_moda_prolongate();
var div_txt=document.getElementById("game_mode_id");
var div_exit=document.getElementById("game_mode_exit_id");
var div_fone=document.getElementById("stimul_action_id");
if(isGame)
{
div_txt.innerHTML="Игровой режим";
div_fone.style.backgroundColor="#D3FFF2";
div_exit.style.display="block";
} 
else
{
div_txt.innerHTML="";
div_fone.style.backgroundColor="#ffffff";
div_exit.style.display="none";
}
}
////////////////////////////
function end_game_mode()
{
var AJAX = new ajax_support(linking_address + "?is_game_mode=0", sent_game_mode_action);
AJAX.send_reqest();
function sent_game_mode_action(res) 
{		
set_game_moda(0);
}
}
///////////////////////////////////////////
// 120 секунд держится игровой режим без обновлдения по каждому периоду ожидания ответа
var limit_teaching_moda_time=120000; 
var teaching_moda_timer_id=0; //id таймера пролонгации игрового режима
function teaching_moda_prolongate()
{
clearTimeout(teaching_moda_timer_id);
teaching_moda_timer_id = setTimeout("end_teaching_mode()",limit_teaching_moda_time);
}
function set_teaching_moda(isTich)
{
teaching_moda_prolongate();
var div_txt=document.getElementById("teaching_mode_div");
if(isTich)
{
div_txt.style.display="block";
} 
else
{
div_txt.style.display="none";
}
}
function end_teaching_mode()
{
var AJAX = new ajax_support(linking_address + "?is_teaching_mode=0", sent_teaching_mode_action);
AJAX.send_reqest();
function sent_teaching_mode_action(res) 
{		
set_teaching_moda(0);
}
}

//////////////////////////////////////
//alert(res);
//if(typeof(exists_connect)=='number' && exists_connect)
//{ 

// посмотреть, в каком режиме находится is_game_mode на сервере, м.б. страница обновилась, а на сервере ГО стоит игровой режим
/*  инфа получается каждую секунду в 
function IsGameMode(){
	var AJAX = new ajax_support(linking_address + "?check_game_mode=1", sent_check_game_mode);
	AJAX.send_reqest();
	function sent_check_game_mode(res) 
	{	
//alert(res);
		set_game_moda(parseInt(res));
	}	
}

IsGameMode();
*/

//}
//////////////////////
</script>
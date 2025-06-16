<?
/* 
include_once($_SERVER['DOCUMENT_ROOT']."/common/alert2_dlg.php");

Диалоги вспомогательных алертов (чтобы их не гасили базовые адерты)

Применение:
show_dlg_alert2(mess,autoclose);   end_dlg_alert2()
	show_dlg_alert2("<br>"+res,2);//только крестик
	show_dlg_alert2("<br>"+res,3);//и крестик и OK
show_dlg_confirm2("Уверены?",1,-1,save_continue);// save_continue - если нажали на ДА
!!! если нажали на др.кнопку, то будет вызов умолчательной функции closed_dlg_confirm2()!!!
!!! ВСЕГДА НУЖНО ОПРЕДЕЛЯТЬ function closed_dlg_confirm2() пусть пустую, чтобы не срабатывало где-то в кеше!!!
end_dlg_confirm2(0); - зактыть confirm2


Только один алерт или конфирм может быть в данный момент на странице. 
Последующий алерт просто перекрывает предыдущий, так что легко можно сделать автосопровождение процесса пояснениями.
*/


?>
<style>
.alert2_exit /* кнопка крестика выхода */
{
position:absolute;
top:1px;
right:2px;
border:0;
border-radius:4px;
color:#FFFFFF;
text-align:center;
cursor:pointer;
background:linear-gradient(180deg, #774E9D, #3D146B);
font-size:15px;
padding:1px;
width:20px;
}
.alert2_exit:hover /* кнопка крестика выхода */
{
background:linear-gradient(180deg, #F08612, #F08612);
}
.alert2s_dlg_botton  					/* кнопка алерта и конфирма */
{
font-size:16px;
}
</style>
<div id="div_dlg_alert2"
style="
position:fixed;
z-index: 10000;
max-height:600px;
#min-width:300px;
overflow:auto;
display:none;
padding: 20px;
top: 50%;left: 50%;
transform: translate(-50%, -50%);
border:solid 2px #8A3CA4;
color:#000000;background-color:#eeeeee;
font-size:16px;font-family:Arial;font-weight:bold;
box-shadow: 8px 8px 8px 0px rgba(122,122,122,0.3);border-radius: 10px;
text-align:center;
"
onmouseup='event.preventDefault();'></div>

<div id='blanck2_div_id' style='position:fixed;Z-INDEX:1000;top:0px;left:0px;width:100%;height:100%;background: rgba(64,64,64,0.7);display:none;' onClick="close_all_dlg(1)"></div>

<script>
// autoclose: 0 - без кнопки ОК, 1 или время в сек - самозакрытие, 2 - без кнопки OK и не самозакрываться
function show_dlg_alert2(mess,autoclose)
{   //alert2(autoclose);
if(autoclose)
{
var def_time=3000;
if(autoclose>100)
	def_time=autoclose;

//var bexit="<div class='alert2_exit' style='top:0.5vw; right:0.5vw;' title='закрыть' onClick='end_dlg_alert2();'><span style='position:relative; top:-1px; left:1px;'>&#10006;</span></div>";

if(autoclose==2)// только крестик
{ 
def_time=1000000000;// фактически не самозакрываться  И с крестиком выхода
document.getElementById('div_dlg_alert2').innerHTML=mess+"<div class='alert2_exit' style='top:0.5vw; right:0.5vw;' title='закрыть' onClick='end_dlg_alert2();'><span style='position:relative; top:-1px; left:1px;'>&#10006;</span></div><div style='height:20px;'>";
}
else
if(autoclose==3)// и крестик и OK
{
def_time=1000000000;// фактичсеки не самозакрываться  И с крестиком выхода
document.getElementById('div_dlg_alert2').innerHTML=mess+"<div class='alert2_exit' style='top:0.5vw; right:0.5vw;' title='закрыть' onClick='end_dlg_alert2();'><span style='position:relative; top:-1px; left:1px;'>&#10006;</span></div><div style='height:20px;'></div><input type='button' class='alert2s_dlg_botton' value='OK' onClick='end_dlg_alert2()' style='position:absolute;bottom:5px;left: 50%;transform: translate(-50%, 0);'>";
}
else
{
	document.getElementById('div_dlg_alert2').innerHTML=mess;
}
	setTimeout("end_dlg_alert2()",def_time);
}
else
document.getElementById('div_dlg_alert2').innerHTML=mess+
"<div style='height:20px;'></div><input type='button' class='alert2s_dlg_botton' value='OK' onClick='end_dlg_alert2()' style='position:absolute;bottom:5px;left: 50%;transform: translate(-50%, 0);'>";


// alert2(mess);
document.getElementById('div_dlg_alert2').style.display="block";
//alert2(document.getElementById('div_dlg_alert2').style.display);
//alert2(document.getElementById('div_dlg_alert2').style.display);
}
function end_dlg_alert2()
{ //alert2("!!!!!!");
document.getElementById('div_dlg_alert2').style.display="none";
}
//////////////////////////////////////////////////////
var is_show_confirm2=0;
var own_proc=0;
//show_dlg_confirm2("Очистить текущий список? Вы уверены?","Уверен","Вернуться",confitm_res);
function show_dlg_confirm2(mess,yes,no,own_function)
{
is_show_confirm2=1;
own_proc=own_function; //alert2(typeof(own_proc));

var name_yes="Да";
if(yes.length>0)
	name_yes=yes;
var name_no="Нет";
if(no.length>0)
	name_no=no;

var buttons="<div style='height:20px;'></div>";


if(no!=-1)
	{
buttons+="<input type='button' value='"+name_yes+"' class='alert2s_dlg_botton' onClick='end_dlg_confirm2(1)' style='position:absolute;bottom:5px;left: 30%;transform: translate(-50%, 0);'>";

buttons+="<input type='button' value='"+name_no+"' class='alert2s_dlg_botton' onClick='def_func();' style='position:absolute;bottom:5px;left: 70%;transform: translate(-50%, 0);'>";

	}
	else//!! ЕСЛИ ВТОРАЯ КНОПКА==-1 то ее не показывать
	{
buttons+="<input type='button' value='"+name_yes+"' class='alert2s_dlg_botton' onClick='end_dlg_confirm2(1)' style='position:absolute;bottom:5px;left: 50%;transform: translate(-50%, 0);'>";
	}


var cntn="<span style='font-size:15px;'>"+mess+"</span>"+buttons;
document.getElementById('div_dlg_alert2').innerHTML=cntn;
document.getElementById('div_dlg_alert2').style.display="block";

bkanking_dlg(1);
}
// закрыто второй кнопкой - если есть умолчательная функция у клиента - отработать
function def_func()
{ //alert2("1");
if(typeof(closed_dlg_confirm2)=='function')
{  //alert2("2");
closed_dlg_confirm2();
}
end_dlg_confirm2(0);
}
function end_dlg_confirm2(type)
{
is_show_confirm2=0;
//	alert2(typeof(own_proc));
if(type)
{
if(typeof(own_proc)=='function')
{
	//alert2(typeof(own_proc));
	//own_proc();
// нужно дать выйти из функции end_dlg_confirm2 чтобы нормально выполнить хоть какую own_proc
	setTimeout("goto_own_proc()",100);
}
}

document.getElementById('div_dlg_alert2').style.display="none";
bkanking_dlg(0);
}
///////////////////////////////////////////////
function goto_own_proc()
{
if(typeof(own_proc)=='function')
{
own_proc();
}
own_proc=0;
}
///////////////////////////////////////////////
/// бланкирование фона при показе диалога
function bkanking_dlg(set)
{
var cy=document.getElementsByTagName( 'body' )[0].offsetHeight;
if(cy<screen.height)
	cy=screen.height;
document.getElementById('blanck2_div_id').style.height=cy+"px";

//alert2(set);
if(set)
document.getElementById('blanck2_div_id').style.display="block";
else
{
if(!is_show_confirm2)
{
document.getElementById('blanck2_div_id').style.display="none";
//close_all_dlg();

}
}
}
function cancel_blocking()
{
document.getElementById('blanck2_div_id').style.display="none";
}

// закрывается по window.onmouseup  в /common/header.php
</script>
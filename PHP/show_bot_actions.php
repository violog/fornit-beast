<?
/* вывод на Пульт действий Beast 
include_once($_SERVER['DOCUMENT_ROOT']."/show_bot_actions.php");

Каждая акция приходит в формате: вид действия (1 - действие рефлекса, 2 - фраза) затем строка акции,
например: "1|Предлогает поиграть" или "2|Привет!"
Может передаваться неограниченная последовательность акций, разделенных "||"
например: "1|Предлогает поиграть||2|Привет!"
*/



?>
<div id="div_bot_action"
style="
position:fixed;
z-index: 100;
bottom: 10px;right: 10px;
max-height:600px;
min-width:250px;
overflow:auto;
display: none;
padding: 10px;padding-top:4px;padding-bottom:4px;
border:solid 2px #8A3CA4;
color:#000000;background-color:#eeeeee;
font-size:16px;font-family:Arial;font-weight:normal;
box-shadow: 8px 8px 8px 0px rgba(122,122,122,0.3);
border-radius: 10px;
text-align:left;
"></div>


<script>
var isAutomatizmShow=0;// для определения начала показа "Осталось времени..."
/*
Каждая акция - в формате: вид действия (1 - действие рефлекса, 2 - фраза) затем строка акции,
например: "1|Предлогает поиграть" или "2|Привет!"
Можно передавать неограниченную последовательность акций, разделяя их "||"
например: "1|Предлогает поиграть||2|Привет!"
*/
function new_bot_action(act_str)
{
isAutomatizmShow=0;
var aOut="";
var actArr=act_str.split("||");
var act="";
var actKind=0;
var actStr="";
for(n=0;n<actArr.length;n++)
{
	if(actArr[n].length==0)
		continue;
act=actArr[n].split("|");  // show_dlg_alert2(act_str,0);
if(act.length!=2)
	{
alert("Неверно прописана акция: "+act_str);
return; // кривая акция
	}
//	alert(actKind);
actKind=act[0];   //actKind=3;
actStr=act[1];

switch(actKind)
{
case "0":
aOut+="<b>Простейший инстинкт:</b><br>"+actStr+"<br>";
break;
case "1":
aOut+=""+actStr+"<br>";// уже есть (в reflex_action.go) <b>БЕССМЫСЛЕННЫЙ безусловный рефлекс:</b>
break;
//case "2":  ГОВОРИТ - ТОЖЕ АВТОМАТИЗМ
//aOut+="<b>Beast говорит:</b><br>"+actStr+"<br>";
//break;
case "3":// моторн.автоматизм
isAutomatizmShow=1;
aOut+="<div style=\"position:relative;padding:10px;background-color:#CCE8FF;\">Бессознательный <b>Автоматизм:</b><br>"+actStr+"</div>";
break;
case "4":// ментальный запуск моторного автоматизма
isAutomatizmShow=1;
aOut+="<div style=\"position:relative;padding:10px;background-color:#CCE8FF;\"><b>Осознанно:</b><br>"+actStr+"</div>";
break;
case "10":// непонимание, растерянность - в случае отсуствия пси-реакций но Лени.
isAutomatizmShow=1;
aOut+="<div style=\"position:relative;padding:10px;background-color:#FFE8E8;\"><b>Непонимание, растерянность:</b><br>"+actStr+"</div>";
break;
//break;
default:
alert("Неверно прописана акция: "+act_str);
break; // кривая акция
}
}
// зеленое сияние вокруг на 1,5 сек
document.getElementById('div_bot_action').style.boxShadow="0px 0px 83px 27px rgba(33, 180, 8, 0.69)"; 
setTimeout("end_effect_bot_action()",1500);

//if(isAutomatizmShow==1){alert(aOut);}
show_dlg_bot_action(aOut);

log=aOut.replace( /(<([^>]+)>)/ig, '' );
log=log.replace(/\</g,'');
log=log.replace(/\>/g,'');
log=log.replace(/&nbsp;/g,' ');// иначе не передает POST
//log+="\r\n";
//alert(log);
set_consol(log); 
}
function end_effect_bot_action()
{
document.getElementById('div_bot_action').style.boxShadow="8px 8px 8px 0px rgba(122,122,122,0.3)";
}
////////////////////////
var actTimer=0;
function show_dlg_bot_action(str)
{
end_dlg_bot_action();
setTimeout("show_dlg_bot_action2('"+str+"')",250);
}
function show_dlg_bot_action2(str)
{
//крестик
var exit="<div class='alert_exit' style='top:2px; right:2px;' title='закрыть' onClick='end_dlg_bot_action();'><span style='position:relative; top:-1px; left:1px;'>&#10006;</span></div>";

document.getElementById('div_bot_action').innerHTML=exit+
"<div style='margin-bottom:10px;'><b>Действия Beast:</b></div>"+str;
document.getElementById('div_bot_action').style.display="block";

// запись в историю
//alert(str);
let pattern = /<img[^>]*>/g;
str = str.replace(pattern, '');  //alert(newTxt);

/*
if(str.indexOf("Ответь сам на")>=0)
alert(str);
*/
var str="<div style='margin-bottom:10px;'><b>Действия Beast:</b></div>"+str;
addInfoToHistory(1,str); // в /index.php

actTimer=setTimeout("end_dlg_bot_action()",10000);
}
////////////////////////////

function end_dlg_bot_action()
{
clearTimeout(actTimer);// если закрыли крестиком, тут же готовность получить новую акцию
document.getElementById('div_bot_action').style.display="none";
}

//new_bot_action("1|Предлогает поиграть||2|Привет!");


//////////////////////////////////
</script>
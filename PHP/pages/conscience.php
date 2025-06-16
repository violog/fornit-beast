<?
/*   Стрнаница текущего состояния функции осознавания Beast
http://go/pages/conscience.php  

Пауза-Пуск для останова просмотра на стр.Сознания
*/
$page_id=8;
$title="Информационная среда осознания Beast";
include_once($_SERVER['DOCUMENT_ROOT']."/common/header.php");
include_once($_SERVER['DOCUMENT_ROOT']."/common/show_waiting.php");
//include_once($_SERVER['DOCUMENT_ROOT']."/common/spoiler.php");

echo "Сохраненные циклы мышления можно смотреть на странице <a href='/pages/mental_cicles.php' target='_blank'>Сохраненные Циклы мышления</a>.<br>";
//include_once($_SERVER['DOCUMENT_ROOT']."/pages/conscience_log.php");echo $conscienceArr[1];
?>   


<div id='puls_id' class='puls_passive' style='position:absolute;top:40px;left:450px;
'></div>

<div id='switcher_id' class='context' style='position:absolute;top:40px;left:900px;
width:100px;background-color:#D0FFD7;text-align:center;
cursor:pointer;
display:none;' onClick="show_switcher()"
title="Временная остановка смены информации.">Пауза</div>

<!-- div id='cliner_all_id' class='context' style='position:absolute;top:70px;left:900px;
width:100px;background-color:#FFC2BF;text-align:center;
cursor:pointer;
display:none;' onClick="cliner_all()"
title="Временная остановка смены информации.">Очистить</div -->

<div id='archive_id' style='min-height:30px;max-height:70px;overflow:auto;'></div>

<div id='div_id' style='font-family:courier;font-size:16px;'>Нужен коннект с Beast.</div>
</div>
</center>
<div style='margin-top:10px;text-align:left;'>



</div>
<br>
<br>
<br>
<br>




<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script Language="JavaScript" src="/ajax/ajax_post.js"></script>
<script>
//clinerlog();

var linking_address='<?include($_SERVER["DOCUMENT_ROOT"]."/common/linking_address.txt");?>';

// ждем пока не включат бестию
check_Beast_activnost(6);// после 4-го пульса И запускается get_info()

var c_timer=0;
var info_timer=0;  //clearTimeout(info_timer);
function get_info()
{
c_timer=setTimeout("connecting()",5000);// проверка коннекта
wait_begin();
var AJAX = new ajax_support(linking_address+"?get_self_conscience_info=1",sent_info);
AJAX.send_reqest();
function sent_info(res)
{
clearTimeout(c_timer);  // show_dlg_alert("есть",0);
//alert(res);
wait_end();

var p=res.split("#|#"); 
showlog(p[0]);// кнопки циклов

document.getElementById('div_id').innerHTML="<b>Текущая Информационная картина:</b> "+p[1];
document.getElementById('puls_id').className="puls_active";

setTimeout("pulse_close()",500);
//setTimeout("chech_new_info()",2000);
info_timer=setTimeout("get_info()",1000);

btn=document.getElementById("switcher_id");
btn.style.display="block";


}
}
//////////////////////////////////////////////
// кнопки циклов: <div id='sh_info_'+index>
function showlog(res)
{
	//   alert(res);
	var list="";
	var p=res.split("|");  
	for(i=0; i<p.length; i++)
	{
		if(p[i].length==0)
			continue;
		//alert(p[i]);
var c_info=p[i].split(",");

var isMain=0;
if(c_info[0]=='m')
	isMain=1;

var id=c_info[1];  
var order=c_info[2];  //alert(p[i]+" | "+id+" | " +order);

var name=i;
var bg="#eeeeee";
if(isMain)
	bg="#BFF5C2";

list+="<div  id='sh_info_"+id+"' class='archive_info' style='background-color:"+bg+";' onClick='get_short_info("+id+")'>"+order+"</div>";

	}
	document.getElementById("archive_id").innerHTML=list;
}
////////////////////  выдать лог выбранного цикла
var c_timer2=0;
function get_short_info(index) // psychic.GetCycleLocInfo()
{
c_timer2=setTimeout("warning_cycle()",1000);// проверка коннекта
wait_end();
show_mode=0;
show_switcher();
var nodes = document.getElementsByClassName('archive_info'); //alert(nodes.length);
for(var i=0; i<nodes.length; i++) 
{ 
nodes[i].className="archive_info";
}
document.getElementById('sh_info_'+index).className="archive_info luminous_box_green";
//alert(document.getElementById('sh_info_'+index).className);

var AJAX = new ajax_support(linking_address+"?get_cycle_log_info="+index,sent_cycle_info);
AJAX.send_reqest();
function sent_cycle_info(res)
{
	clearTimeout(c_timer2); 
	if(res.length>4)
document.getElementById('div_id').innerHTML="<b>Текущая Информационная картина:</b> "+res;
	else
document.getElementById('div_id').innerHTML="Нет инфо-картины для этого цикла.";

end_dlg_alert();
}
}
function warning_cycle()
{
show_dlg_alert("Нет коннекта с Beast...",0);
}
//////////////////////



function pulse_close()
{
document.getElementById('puls_id').className="puls_passive";
}
function connecting()
{
// нет коннекта 5 секунд
//get_info(); // попытка дозвониться
show_dlg_alert("Нет коннекта с Beast.<br>Включите Beast и <span style='color:blue;cursor:pointer;' onClick='location.reload(true)'>повторите</span> операцию.",0);
}


/////////////////////////////////////////////
var show_mode=0;// 0- работа, 1 - пауза
function show_switcher()
{
btn=document.getElementById("switcher_id");
if(show_mode==0)
{
	wait_end();
	show_mode=1;
btn.style.backgroundColor="#EEEEEE";
btn.style.fontWeight="normal";
btn.innerHTML="Возобновить";
clearTimeout(info_timer);
clearTimeout(c_timer);
} 
else
{
	wait_end();
	show_mode=0;  // context 
btn.style.backgroundColor="#D0FFD7";
btn.style.fontWeight="bold";
btn.innerHTML="Пауза";
get_info();
}
}
///////////////////////////////////////////////////


function show_automatizms(id)
{
var AJAX = new ajax_support(linking_address + "?autNodeID="+id+"&get_node_automatizms=1", sent_automatizms_info);
AJAX.send_reqest();
function sent_automatizms_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}

function get_situation(id)
{
var AJAX = new ajax_support(linking_address + "?autNodeID="+id+"&get_node_mental_situations=1", sent_situation_info);
AJAX.send_reqest();
function sent_situation_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}

function get_purpose(id)
{
var AJAX = new ajax_support(linking_address + "?autNodeID="+id+"&get_node_mental_purpose=1", sent_purpose_info);
AJAX.send_reqest();
function sent_purpose_info(res) {
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

function show_cyckle(id) // GetCyckleInfo()
{
var AJAX = new ajax_support(linking_address + "?objID="+id+"&get_show_cyckle_info=1", sent_undstg_info);
AJAX.send_reqest();
function sent_undstg_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}

function show_mental_actions(actID)
{
var AJAX = new ajax_support(linking_address + "?get_mental_action_info="+actID, sent_action_info);
AJAX.send_reqest();
function sent_action_info(res) {

show_dlg_alert(res,0); 
}
}

function show_mental_goNext(id) // GetMentalActionsString
{
var AJAX = new ajax_support(linking_address + "?get_mental_goNext_info="+id, sent_action_info);
AJAX.send_reqest();
function sent_action_info(res) {

show_dlg_alert(res,0); 
}
}


</script>

</div>
</body>
</html>
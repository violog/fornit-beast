<?
/*  Модели понимания объектов восприятия
http://go/pages/mental_undastending_models.php  
*/



$page_id = 7;
$title = "Модели понимания объектов восприятия";
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");
//include_once($_SERVER['DOCUMENT_ROOT']."/pult_js.php");

include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/show_waiting.php");

// из-за конфликтов чтение-запись убрал авто обновление
?>
<div id='div_id' style='font-family:courier;font-size:16px;'>Нужен коннект с Beast.</div>
</div>


<div id='puls_id' class='puls_passive' style='position:absolute;top:40px;left:450px;'></div>

<div id='rules_info_id' style='font-family:courier;font-size:16px;'></div>


<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var linking_address = '<? include($_SERVER["DOCUMENT_ROOT"] . "/common/linking_address.txt"); ?>';

// ждем пока не включат бестию
check_Beast_activnost(6);// после 6-го пульса И запускается get_info()
//get_info();

var old_size = 0;
var limitBasicID=0;//>0 - лимитировать показ только одним из базовых состояний Плохо,Норма,Хорошо

var c_timer=0;
function get_info() {  
c_timer=setTimeout("connecting()",5000);// проверка коннекта

var AJAX = new ajax_support(linking_address + "?get_mental_undestanding_models_info=1", sent_get_info);
AJAX.send_reqest();
function sent_get_info(res) { //alert(res);

clearTimeout(c_timer);
if(res.length<5)
	{
document.getElementById('rules_info_id').innerHTML = "Еще нет информации.";
	}
document.getElementById('div_id').innerHTML=res;
document.getElementById('puls_id').className="puls_active";
setTimeout("pulse_close()",500);
//show_dlg_alert("!!!!!!!!!!!",0); 

setTimeout("get_info()",1000);
}
}
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
/////////////////////////////////////////////////


function show_atmzm_tree(id)
{
var AJAX = new ajax_support(linking_address + "?objID="+id+"&get_atmzm_tree_info=1", sent_atmzm_tree_info);
AJAX.send_reqest();
function sent_atmzm_tree_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}
function show_unde_tree(id)
{
var AJAX = new ajax_support(linking_address + "?objID="+id+"&get_undstg_tree_info=1", sent_undstg_info);
AJAX.send_reqest();
function sent_undstg_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}

function show_ment_atmzm(id)
{
var AJAX = new ajax_support(linking_address + "?objID="+id+"&get_ment_atmzm_info=1", sent_ment_atmzm_info);
AJAX.send_reqest();
function sent_ment_atmzm_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}




function get_ment_model_index(id)//образ типа extremImportance
{
var AJAX = new ajax_support(linking_address + "?autNodeID="+id+"&get_ment_model_index=1", sent_purpose_info);
AJAX.send_reqest();
function sent_purpose_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}

// НЕТ ТАКОГО И НЕ ИСПОЛЬЗУЕТСЯ
function get_epiz_memory_info(id)//кадр эпиз.памяти
{
var AJAX = new ajax_support(linking_address + "?autNodeID="+id+"&get_epiz_memory_info=1", sent_purpose_info);
AJAX.send_reqest();
function sent_purpose_info(res) {
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


function get_trigger_action_info(id)
{
var AJAX = new ajax_support(linking_address + "?autNodeID="+id+"&get_trigger_action_info=1", sent_purpose_info);
AJAX.send_reqest();
function sent_purpose_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}

}

</script>

</body>
</html>
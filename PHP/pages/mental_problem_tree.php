<?
/*  Дерево проблем
http://go/pages/mental_problem_tree.php  
*/


$page_id = 7;
$title = "Дерево проблем";
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/show_waiting.php");
?>

<div style='position:absolute;top:40px;left:300px;font-family:courier;font-size:16px;cursor:pointer;' onClick="location.reload(true)">Обновить</div>

<div id='div_id' style='font-family:courier;font-size:16px;'>Нужен коннект с Beast.</div>
</div>

<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var linking_address = '<? include($_SERVER["DOCUMENT_ROOT"] . "/common/linking_address.txt"); ?>';

// ждем пока не включат бестию
check_Beast_activnost(6);// после 6-го пульса И запускается get_info()
////////////////////////////////////////////////////////////////////////

//show_dlg_alert("Ожидаем 4 секунды...",0);
//setTimeout("get_info()",4000);
var limit=0;//>0 - лимитировать показ только одним из базовых состояний Плохо,Норма,Хорошо
function get_info() { //alert("!!!!");
		wait_begin();
var AJAX = new ajax_support(linking_address + "?limit="+limit+"&get_mental_priblem_tree=1", sent_info);
AJAX.send_reqest();
function sent_info(res) {
			//alert(res);
		wait_end();
	document.getElementById('div_id').innerHTML = res;
}
}
function show_level(base)
{
	limit=base;
get_info();
}
///////////////////////////////////////////

////////////////
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
function show_node_automatizms(id)
{
var AJAX = new ajax_support(linking_address + "?autNodeID="+id+"&get_node_aut_tree_node=1", sent_automatizms_info);
AJAX.send_reqest();
function sent_automatizms_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}


function get_theme(id) // psychic.GetMentalThemeForPult(tID)
{
var AJAX = new ajax_support(linking_address + "?autNodeID="+id+"&get_node_theme_image=1", sent_purpose_info);
AJAX.send_reqest();
function sent_purpose_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}

function get_ment_automatizm(id) // psychic.GetMentalAutomatizmForPult(tID)
{
var AJAX = new ajax_support(linking_address + "?autNodeID="+id+"&get_node_mental_automatizm=1", sent_purpose_info);
AJAX.send_reqest();
function sent_purpose_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}

function show_actions(id)
{
//alert(id);
var AJAX = new ajax_support(linking_address + "?autmzmID="+id+"&get_sequence_info=1", sent_sequence_info);
AJAX.send_reqest();
function sent_sequence_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}

</script>

</div>
</body>

</html>
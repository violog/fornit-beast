<?
/* Ментальные Правила
http://go/pages/mental_rules.php  
*/

$page_id = 7;
$title = "Правила ментальные, Стимул-Ответ-Эффект";
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");
//include_once($_SERVER['DOCUMENT_ROOT']."/pult_js.php");

include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/show_waiting.php");

// из-за конфликтов чтение-запись убрал авто обновление
?>
<style>
.frameEP
{
display:inline-block;
//width:16px;
height:16px;
padding:2px;
border:solid 1px #000000;
border-radius: 7px;
background-color:#DDE1E4;
}
.frameEm
{
display:inline-block;
//width:16px;
height:16px;
padding:2px;
border:solid 1px #000000;
border-radius: 7px;
background-color:#ffffff;
}
</style>


<div style='position:absolute;top:40px;left:450px;font-family:courier;font-size:16px;cursor:pointer;' onClick="location.reload(true)"><b>Обновить</b> - нет автоматического обновления<span style="padding-left:100px"><span> 
</div>

<div id='div_id' style='font-family:courier;font-size:16px;'>Нужен коннект с Beast.</div>
</div>

<div id='rules_info_id' style='font-family:courier;font-size:16px;'></div>


<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var linking_address = '<? include($_SERVER["DOCUMENT_ROOT"] . "/common/linking_address.txt"); ?>';

// ждем пока не включат бестию
check_Beast_activnost(6);// после 6-го пульса И запускается get_info()
//get_info();

var old_size = 0;
var limitBasicID=0;//>0 - лимитировать показ только одним из базовых состояний Плохо,Норма,Хорошо


function get_info() {  
var AJAX = new ajax_support(linking_address + "?get_mental_rulles_list_info=1", sent_get_info);
AJAX.send_reqest();
function sent_get_info(res) {
if(res.length<10)
	{
document.getElementById('rules_info_id').innerHTML = "Еще нет правил.";
	}

document.getElementById('div_id').innerHTML=res;

//show_dlg_alert("!!!!!!!!!!!",0); 
}

//setTimeout("get_info()",2000);из-за конфликтов чтение-запись убрал авто обновление
}



function show_problem_tree(id)
{

var AJAX = new ajax_support(linking_address + "?mautmzmID="+id+"&get_show_problem_tree_info=1", sent_sequence_info);
AJAX.send_reqest();
function sent_sequence_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}

function show_mental_func(fID)
{
var AJAX = new ajax_support(linking_address + "?get_mental_func_info="+fID, sent_func_info);
AJAX.send_reqest();
function sent_func_info(res) {

show_dlg_alert(res,0); 
}
}

function show_theme_strings(tID)
{
//	alert(cickle);
var AJAX = new ajax_support(linking_address + "?autNodeID="+tID+"&get_node_theme_image=1", sent_purpose_info);
AJAX.send_reqest();
function sent_purpose_info(res) {
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
</script>

</body>
</html>
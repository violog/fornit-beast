<?
/*   Правила для объекта с ID=nn
http://go/pages/rules_model.php  


*/
$objID=$_GET['id'];


$page_id = 666;
$title = "Правила для объекта с ID=".$objID;
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/show_waiting.php");

if(empty($objID))
{
echo "Нужно задать параметр ID в адресной строке:<br>http://go/pages/rules_model.php?id=nn";
exit();
}

// из-за конфликтов чтение-запись убрал авто обновление
?>
<div style='position:absolute;top:40px;left:350px;font-family:courier;font-size:16px;cursor:pointer;' onClick="location.reload(true)"><b>Обновить</b> - нет автоматического обновления<span style="padding-left:100px"><span> 
</div>

<div id='div_id' style='font-family:courier;font-size:16px;'>Нужен коннект с Beast.</div>
</div>



<div id='rules_info_id' style='font-family:courier;font-size:16px;'></div>

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
.frameEv
{
display:inline-block;
//width:16px;
height:16px;
padding:2px;
border:solid 1px #000000;
border-radius: 7px;
background-color:#33ccff;
}
</style>

<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var linking_address = '<? include($_SERVER["DOCUMENT_ROOT"] . "/common/linking_address.txt"); ?>';

// ждем пока не включат бестию
check_Beast_activnost(6);// после 6-го пульса И запускается get_info()
//get_info();

var old_size = 0;
var limitBasicID=0;//>0 - лимитировать показ только одним из базовых состояний Плохо,Норма,Хорошо


function get_info() {  
var AJAX = new ajax_support(linking_address + "?objID=<?=$objID?>&get_undestand_model=1", sent_get_info);
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


function show_emotion(emotionID)
{
var AJAX = new ajax_support(linking_address + "?emotionID="+emotionID+"&get_emotion_info=1", sent_emotion_info);
AJAX.send_reqest();
function sent_emotion_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
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

function get_undestand_model(id)
{ //alert(id);
var AJAX = new ajax_support(linking_address + "?objID="+id+"&get_undestand_model=1", sent_undestand_model_info);
AJAX.send_reqest();
function sent_undestand_model_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}
</script>

</body>
</html>
<?
/* сохраненные от последнего пробуждения Циклы осознания
http://go/pages/mental_cicles.php  
*/

$page_id = 7;
$title = "Сохраненные Циклы мышления";
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");
//include_once($_SERVER['DOCUMENT_ROOT']."/pult_js.php");


echo "В динамике изменений циклы мышления можно смотреть на странице <a href='/pages/conscience.php' target='_blank'>Информационная среда осознания Beast</a>.<br>";

$files=array();

$isEmptyList=true;
if($dh = opendir($_SERVER['DOCUMENT_ROOT']."/cycle_logs/")) 
{ //exit("!!!");
while(false !== ($file = readdir($dh))) 
{		
if($file=="." || $file=="..")
	continue;
	
array_push($files,$file);
$isEmptyList=false;

}
closedir($dh);
}

if($isEmptyList)
{
echo "Нет сохраненных циклов мышления от последнего периода бодрствования.";
}
else
{
sort($files, SORT_NUMERIC);
reset($files);

	echo "Циклы, сохраненные от последнего периода бодрствования:<br><br>";
	foreach($files as $file)
	{
$num=substr($file,0,strlen($file)-4);

echo "<span id='".$num."_id'  class='archive_num' onClick='show_cycle(".$num.",\"".$file."\")'>".$num."</span>";

	}

}
?>
<div id='cycle_info_id' style='font-family:courier;font-size:16px;'></div>





<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var linking_address='<?include($_SERVER["DOCUMENT_ROOT"]."/common/linking_address.txt");?>';
function show_cycle(num,file)
{
//alert(file);
var AJAX = new ajax_support("/pages/mental_cycle_file.php?file="+file, sent_get_info);
AJAX.send_reqest();
function sent_get_info(res) 
{
document.getElementById('cycle_info_id').innerHTML = res;

var nodes = document.getElementsByClassName('archive_num');
for(var i=0; i<nodes.length; i++)
{ 
nodes[i].style.backgroundColor="#ffffff";
}
document.getElementById(num+'_id').style.backgroundColor="#A8FFAA";
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

</body>
</html>

<?
/*   список ментальных автоматизмов
http://go/pages/ mental_automatizm_table.php  

Формат записи:
id|Usefulness|ActionsImageID|Count
*/

$page_id = 7;
$title = "Список ментальных автоматизмов";
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/show_waiting.php");

$out_str_for_del = ""
?>
<div style='position:absolute;top:40px;left:350px;font-family:courier;font-size:16px;cursor:pointer;' onClick="location.reload(true)"><b>Обновить</b></div>

<div id='div_id' style='font-family:courier;font-size:16px;'>Нужен коннект с Beast.</div>
</div>


<div id='automatizm_info_id' style='font-family:courier;font-size:16px;'></div>



<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var linking_address = '<? include($_SERVER["DOCUMENT_ROOT"] . "/common/linking_address.txt"); ?>';

// ждем пока не включат бестию
check_Beast_activnost(6);// после 6-го пульса И запускается get_info()

var old_size = 0;
var limitPage=0;//>0 - лимитировать показ постранично


function get_info() {
	end_dlg_alert();
	wait_begin();
		var AJAX = new ajax_support(linking_address + "?limitPage="+limitPage+"&get_mental_automatizm_list_info=1", sent_get_info);
		AJAX.send_reqest();

		function sent_get_info(res) {
			//alert(res);
			wait_end();
			document.getElementById('automatizm_info_id').innerHTML = res;
document.getElementById('div_id').innerHTML="Информация - по щелчку на Пусковой стимул или Действия автоматизма.";
					}
	}
	

function show_page(nPage)
{ //alert(nPage);
	limitPage=nPage;
get_info();
}
//////////////////////////////////////////////


function show_actions(id)
{

var AJAX = new ajax_support(linking_address + "?mautmzmID="+id+"&get_mental_actiom_img_info=1", sent_sequence_info);
AJAX.send_reqest();
function sent_sequence_info(res) {
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


function show_automatizm(id)
{
var AJAX = new ajax_support(linking_address + "?autID="+id+"&get_automatizm=1", sent_automatizms_info);
AJAX.send_reqest();
function sent_automatizms_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}
</script>

</body>
</html>
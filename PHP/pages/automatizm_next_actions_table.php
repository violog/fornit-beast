<?
/*   список цепочек действий AmtzmNextString 
http://go/pages/automatizm_next_actions_table.php  


*/

$page_id = 7;
$title = "Список цепочек действий AmtzmNextString";
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/show_waiting.php");

$out_str_for_del = "";
?>
<div style='position:absolute;top:40px;left:450px;font-family:courier;font-size:16px;cursor:pointer;' onClick="location.reload(true)"><b>Обновить</b></div>



<div id='div_id' style='font-family:courier;font-size:16px;'>Нужен коннект с Beast.</div>
</div>


<div id='next_info_id' style='font-family:courier;font-size:16px;'></div>



<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var linking_address = '<? include($_SERVER["DOCUMENT_ROOT"] . "/common/linking_address.txt"); ?>';

// ждем пока не включат бестию
check_Beast_activnost(6);// после 4-го пульса И запускается get_info()

var old_size = 0;
var limitBasicID=0;//>0 - лимитировать показ только одним из базовых состояний Плохо,Норма,Хорошо


function get_info() {
	end_dlg_alert();
	wait_begin();
		var AJAX = new ajax_support(linking_address + "?limitBasicID="+limitBasicID+"&get_next_actions_info_list=1", sent_get_info);
		AJAX.send_reqest();

		function sent_get_info(res) {
			//alert(res);
			wait_end();
			document.getElementById('next_info_id').innerHTML = res;
//document.getElementById('div_id').innerHTML="Информация - по щелчку на Действия автоматизма.";
					}
	}
	

function show_level(base)
{
	limitBasicID=base;
get_info();
}
//////////////////////////////////////////////



function show_next_actions(nextID)//func GetNextActionsInfo(nextID int)string{
{
var AJAX = new ajax_support(linking_address + "?nextID="+nextID+"&get_next_actions_info=1", sent_emotion_info);
AJAX.send_reqest();
function sent_emotion_info(res) {
			//alert(res);
res=res.replace(/\<\/b\>/g,'</b><br>');
show_dlg_alert("<div style='text-align:left;font-weight:normal;'>"+res+"</div>",0);
}
}
</script>

</body>
</html>
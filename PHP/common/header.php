<?
/*      
$page_id=0;
$title="Пульт связи с Beast";
include_once($_SERVER['DOCUMENT_ROOT']."/common/header.php");
*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");
mb_http_input('UTF-8');
mb_http_output('UTF-8');
mb_internal_encoding("UTF-8");




?>
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title><?=$title?></title>
  <link rel="stylesheet" href="/common/common.css">
</head>
<body style="margin-left:20px;font-family:arial;font-size:14px;">
<form name="open_another_win" method="post" action="" target="_blank"></form>
<?
include_once($_SERVER['DOCUMENT_ROOT']."/common/common.php"); 

include_once($_SERVER["DOCUMENT_ROOT"]."/common/top_menu.php");
?>
<h1 style='font-size:18px;'><?=$title?></h1>

<script>
function open_anotjer_win(link)
{  
document.forms.open_another_win.action=link;
document.forms.open_another_win.submit();
}
window.onmouseup = function(e)
{
var t = e.target || e.srcElement;
while(t)
{
if(t.id == "div_dlg_alert")
	return;	
t = t.offsetParent;
}
end_dlg_confirm(0);//   зкрывать конфирм при щелчке по серому.
//!!!!! end_dlg_control();

// закрыть спойлер инструментов
if(document.getElementById("block_id"))
if(parseInt(document.getElementById("block_id").style.height)!=0)
	open_close(`block_id`,0);

}
// для поддержки страниц информации от 
// ждем пока не включат бестию
// !!!! ЕСЛИ ГДЕ-ТО БЕСКОНЕЧНЫЙ ЦИКЛ, ТО brain.PulsCount остановится !!!!
var begin_activnost_after=1; // после какого пульса запускать функцию get_info()
function check_Beast_activnost(begin)
{
if(typeof(get_info)!="function")
{
	alert("Нужно определить функцию получения информации get_info()");
	return;
}
begin_activnost_after=begin;
get_Beast_connection();
}
var check_info_timer=0
var check_info_timer_cheker=0
function get_Beast_connection()
{	
check_info_timer_cheker=setTimeout("get_Beast_connection_checker()",1000);
var AJAX = new ajax_support(linking_address + "?check_Beast_activnost=1", check_info);
AJAX.send_reqest();
function check_info(res) 
{ 
//	alert(res);
clearTimeout(check_info_timer_cheker);
document.getElementById('div_id').innerHTML = "Beast активна, но еще не готова, подождите "+(begin_activnost_after-res)+" секунд.";

res=parseInt(res);  //alert(res +" | "+begin_activnost_after);
			if(res > begin_activnost_after)
			{  //alert(res);
clearTimeout(check_info_timer);
clearTimeout(check_info_timer_cheker);
document.getElementById('div_id').innerHTML = "Beast активна";
get_info();
//return;
			}
		}
		check_info_timer=setTimeout("get_Beast_connection()",1000);
}
// есть ли коннект с ГО

function exists_connection()
{
check_info_timer_cheker=setTimeout("get_Beast_connection_checker()",2000);
var AJAX = new ajax_support(linking_address + "?check_Beast_activnost=1", check_go_connction);
AJAX.send_reqest();
function check_go_connction(res) {
//	alert(res);
clearTimeout(check_info_timer_cheker);
}
}
function get_Beast_connection_checker()
{
//	alert();
clearTimeout(check_info_timer);
show_dlg_alert("Нет коннекта с Beast.<br>Включите Beast и <span style='color:blue;cursor:pointer;' onClick='location.reload(true)'>повторите</span> операцию.",0);
wait_end();
}
</script>
<?
///// стадии развития 
$stages=read_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/stages.txt");
$stages=trim($stages);
switch($page_id)
{
case 1:case 3:case 4: // кроме слов case 2:
	if($stages>0)
	{
echo "<div style='font-size:21px;color:red;'><b>Это - пройденная стадия развития! Не следует редактировать данные на этой станице!</b></div>";
	}
break;



}

?>
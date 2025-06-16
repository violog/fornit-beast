<?
/*  список распознаваемых слов в алфавитном порядке
http://go/pages/words_list.php  

*/
$page_id=-1;
$title="Список слов";
include_once($_SERVER['DOCUMENT_ROOT']."/common/header.php");
include_once($_SERVER['DOCUMENT_ROOT']."/common/show_waiting.php");
?>
<div  style='position:absolute;top:40px;left:200px;font-family:courier;font-size:16px;cursor:pointer;' onClick="get_info()">Обновить</div>

<div id='div_id' style='font-family:courier;font-size:16px;'>Нужен коннект с Beast.</div>
</div>
<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var linking_address='<?include($_SERVER["DOCUMENT_ROOT"]."/common/linking_address.txt");?>';
var old_size=0;
function get_info()
{
wait_begin();
var AJAX = new ajax_support(linking_address+"?get_words_list=1",sent_info);
AJAX.send_reqest();
function sent_info(res)
{
//alert(res);
wait_end();
document.getElementById('div_id').innerHTML=res;
}
}
get_info();

var cur_delete_word_id=0;
function delete_word(id)
{
cur_delete_word_id=id;
show_dlg_confirm("Точно удалить это слово?<br>Будут удалены эти слова из Дерева Слов и Дерева Фраз. Операция критическая.","Удалить","Нет",delete_word2);
}
function delete_word2(id)
{
wait_begin();
var AJAX = new ajax_support(linking_address+"?deleting_word=1&delete_word="+cur_delete_word_id,sent_del_info);
AJAX.send_reqest();
function sent_del_info(res)
{
//alert(res);
wait_end();
if(res=="OK")
	{
show_dlg_alert("Слово было удалено.<br><br>Необходимо <a href='javascript:reload_beast()'>перезагрузить</a> Beast.",0);
	}
	else
	{
show_dlg_alert("Не удалось удалить слово с ID="+cur_delete_word_id+". Нужно обратить к разработчику.",0);
	}
}
}

// выключить и снова включить исполняемый файл
function reload_beast()
{
end_dlg_alert();
wait_begin();
var AJAX = new ajax_support("/kill.php", sent_reb1_answer);
AJAX.send_reqest();
//alert("1");
setTimeout("rebooting()",1000);// выждать завершения процессов
function sent_reb1_answer(res)
	{

	}
}
function rebooting()
{
//	alert("2");
var AJAX = new ajax_support("/run.php", sent_reb2_answer);
AJAX.send_reqest();
setTimeout("rebooting2()",500);
function sent_reb2_answer(res)
	{

	}
}
function rebooting2()
{
wait_end();  //alert("333");
show_dlg_alert("Beast перезагружен.",1500);
setTimeout("location.reload(true)",1500);
}
</script>

</div>
</body>
</html>
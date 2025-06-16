<?
/*  дерево слов
http://go/pages/words_tree.php  

*/
$page_id=-1;
$title="Дерево слов";
include_once($_SERVER['DOCUMENT_ROOT']."/common/header.php");
include_once($_SERVER['DOCUMENT_ROOT']."/common/show_waiting.php");
?>
<div  style='position:absolute;top:40px;left:200px;font-family:courier;font-size:16px;cursor:pointer;' onClick="get_info()">Обновить</div>

<div id='div_id' style='font-family:courier;font-size:16px;'></div>
</div>
<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var old_size=0;
function get_info()
{
wait_begin();
// Все же нужна связь с Ботом для получения инфы о дереве.
var AJAX = new ajax_support("/pages/words_tree_server.php?old_size="+old_size,sent_info);
AJAX.send_reqest();
function sent_info(res)
{
	wait_end();
//alert(res);
var p=res.split("|##|");
/*
if(p[0]=="not")
{
	setTimeout(`get_info()`,1000);
	return;
}
*/
old_size=p[0];
document.getElementById('div_id').innerHTML=p[1];
//alert(old_size);
//setTimeout(`get_info()`,1000);
}
}
get_info();
</script>

</div>
</body>
</html>
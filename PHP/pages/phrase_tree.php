<?
/*  дерево фраз
http://go/pages/phrase_tree.php  

*/
$page_id=-1;
$title="Дерево фраз";
include_once($_SERVER['DOCUMENT_ROOT']."/common/header.php");
include_once($_SERVER['DOCUMENT_ROOT']."/common/show_waiting.php");
?>
<div  style='position:absolute;top:40px;left:200px;font-family:courier;font-size:16px;cursor:pointer;' onClick="get_info()">Обновить</div>

<div id='div_id' style='font-family:courier;font-size:16px;'>Нужен коннект с Beast.</div>
</div>
<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var linking_address='<?include($_SERVER["DOCUMENT_ROOT"]."/common/linking_address.txt");?>';
// ждем пока не включат бестию
check_Beast_activnost(4);
var old_size=0;
function get_info()
{
/*
var AJAX = new ajax_support("/pages/phrase_tree_server.php?old_size="+old_size,sent_size_info);
AJAX.send_reqest();
function sent_size_info(res)
{
if(res!="not")
{
	*/
wait_begin();
var AJAX = new ajax_support(linking_address+"?get_phrase_tree=1",sent_info);
AJAX.send_reqest();
function sent_info(res)
{
//alert(res);
wait_end();
document.getElementById('div_id').innerHTML=res;
}
//old_size=res;
//setTimeout(`get_info()`,1000);

//}
//alert(old_size);
//setTimeout(`get_info()`,1000);
//}
}
get_info();
</script>

</div>
</body>
</html>
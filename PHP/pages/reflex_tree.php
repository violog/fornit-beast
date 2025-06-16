<?
/*  дерево рефлексов
http://go/pages/reflex_tree.php  

*/
$page_id=-1;
$title="Дерево рефлексов";
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
var AJAX = new ajax_support(linking_address+"?get_reflex_tree=1",sent_info);
AJAX.send_reqest();
function sent_info(res)
{
//alert(res);
wait_end();
document.getElementById('div_id').innerHTML=res;
}
}
get_info();
</script>

</div>
</body>
</html>
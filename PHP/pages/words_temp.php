<?
/*  Накопитель слов-фраз
http://go/pages/words_temp.php  

*/
$page_id=-1;
$title="Накопитель слов-фраз";
include_once($_SERVER['DOCUMENT_ROOT']."/common/header.php");
?>
В скобках - число повторений. Сначала наиболее частые.
<div id='div_id'></div>
</div>
<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var old_size=0;
function get_info()
{
var AJAX = new ajax_support("/pages/words_temp_server.php?old_size="+old_size,sent_info);
AJAX.send_reqest();
function sent_info(res)
{
//alert(res);
var p=res.split("|##|");
if(p[0]=="not")
{
	setTimeout(`get_info()`,1000);
	return;
}
old_size=p[0];
document.getElementById('div_id').innerHTML=p[1];
//alert(old_size);
setTimeout(`get_info()`,1000);
}
}
get_info();
</script>
</body>
</html>
<?
/* крутящаяся гифка при ожидании
include_once($_SERVER['DOCUMENT_ROOT']."/common/show_waiting.php");

*/
?>
<div id="waite_loading" style="display:none;position:absolute;z-index:1001;width:128px;height:128px;top: 50%;left: 50%;transform: translate(-50%, -50%);"><img src="/common/loading.gif" border=0></div>

<script>
var timer_wait=0;
function wait_begin()// начало показа гифки ожидания через 0,5сек
{
document.getElementById('waite_loading').style.display="none";
timer_wait=setTimeout("wait_show()",500);
}
function wait_show()
{
document.getElementById('waite_loading').style.display="block";
}
function wait_end()//гашение гифки
{
clearTimeout(timer_wait);
document.getElementById('waite_loading').style.display="none";
}
</script>
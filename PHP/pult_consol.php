<?
/*  Консоль событий Beast для Пульта

*/


?>


<div style="position:relative;
margin-top:30px;">
<div style="position:absolute;top:0px;right:0px;cursor:pointer;" onClick='show_consol_win()'>Показывать в отдельном окне</div>
<b>Консоль событий Beast (тех инфа):</b>
<div style="position:absolute;top:30px;left:-10px;color:red;cursor:pointer;padding:4px;border:solid 1px #8A3CA4;border-radius:50%;background-color:#ffffff" title="Очистить окно ввода" onClick="cliner_consol()"><b>X</b></div>

<div id="pult_consol_id" style="border:solid 1px #666666;
min-height:60px;max-height:200px;
padding-left:10px;padding-right:10px;
overflow:auto;
"><span style="color:#888888;">Пока нет событий.</span></div>


</div>
<script>
var list_consol_info="";
function cliner_consol()
{
document.getElementById('pult_consol_id').innerHTML="";
list_consol_info="";
var AJAX = new ajax_support('/pages/consol_cliner_info.php',sent_request_info);
AJAX.send_reqest();
function sent_request_info(res)
{
//alert(res);
}
}
function set_consol(str)
{
var list_consol_info=str+"<br><separ>\r\n"; // alert(list_consol_info);
document.getElementById('pult_consol_id').innerHTML=list_consol_info;

// записывать в файл /pult_consol.txt - последнее - наверх
param="info="+list_consol_info;
var AJAX = new ajax_post_support('/pages/consol_server.php',param,sent_consol_info,1);
AJAX.send_reqest();
function sent_consol_info(res)
{
//alert(res);
}
}
//////////////////////////////////
function show_consol_win()
{
open_anotjer_win("/pages/consol.php");

}
</script>
<?
/*   движок связи с Beast
Если на данный странице нужно передавать и принимать данные для Beast, то:
<script Language="JavaScript" src="/ajax/ajax_post.js"></script>
include_once($_SERVER['DOCUMENT_ROOT']."/common/linking.php");

вызвать bot_contact(params,own_proc) с params в виде GET-запроса НАЧИНАЯ С &, а own_proc - имя фнкции для получения данных от Beast.
*/



?>
<script>
var linking_address='<?include($_SERVER["DOCUMENT_ROOT"]."/common/linking_address.txt");?>';
//alert(linking_address)
/*
function bot_contact111(params,own_proc)
{
if(typeof(own_proc)=='function')
	{
own_proc("WWWW");
	}
}
*/

// 0 - блокирование контактов на время операций с памятью для функций Инструментов
var actived_contact=1;


// bot_contact посылается раз 1сек постоянно из pult.php
var bot_cur_answer="";
function bot_contact(params,own_proc)
{
	if(actived_contact==0)// блокировать все запросы
	{
		return;
	}

 // show_dlg_alert(t_count,500);
  //alert(params);return;
var AJAX = new ajax_post_support(linking_address,params,sent_request_bot,1);
AJAX.send_reqest();
function sent_request_bot(res)
{
//alert(res);
if(typeof(own_proc)=='function')
	{
	//setTimeout("own_proc()",1);
	own_proc(res);
	}
}
}

function bot_contact_get(params,own_proc){
	if(actived_contact==0){
		return;
	}
	var AJAX = new ajax_support(linking_address+"?"+params,sent_request_bot);
	AJAX.send_reqest();
	function sent_request_bot(res){
	if(typeof(own_proc)=='function'){
		own_proc(res);
		}
	}
}

</script>

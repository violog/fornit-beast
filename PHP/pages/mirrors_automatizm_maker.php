<?
/* Сформировать автоматизмы для всех списков ответов.

http://go/pages/mirrors_automatizm_maker.php

перебирает файлы в /lib/mirror_reflexes_basic_phrases посылая запросы в ГО 
*/

$page_id = -1;
$title = "Формирование зеркальных автоматизмов на основе списка ответов";
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");
//include_once($_SERVER['DOCUMENT_ROOT']."/pult_js.php");
//////////////////////////////////////////////////////////////
$dir=$_SERVER["DOCUMENT_ROOT"]."/lib/mirror_reflexes_basic_phrases/";
$filesArr="var filesArr = new Array();";
$n=0;
if($dh = opendir($dir)) 
{ //exit("!!!");
while(false !== ($file = readdir($dh))) 
{		
if($file=="." || $file=="..")
	continue;
if(filesize($dir.$file)>0)
	{
$filesArr.="filesArr[".$n."] = '/lib/mirror_reflexes_basic_phrases/".$file."';";
$n++;
	}

}
closedir($dh);
}
$filesArr.="var fileCount=$n;";
//  var_dump($contents);exit();
if($n==0)
{
echo "<div style='font-family:courier;font-size:16px;display:block;'><span style='font-size:18px;color:red;'><b>Нет файлов списков.</b></span><br> Сначала нужно в <a href='/pages/mirrors_automatizm.php'>редакторе</a> заготовить фразы - ответы.</div>";
exit();
}


echo "<div id='div_id' style='font-family:courier;font-size:18px;'><b>Нужен коннект с Beast.</b></div>";


include_once($_SERVER['DOCUMENT_ROOT']."/common/linking.php");
?>
Сначала нужно:<br>
1. Включить Beast и запустить процесс: <span style="font-size:21px;border:solid 1px #8A3CA4;border-radius: 7px;padding-left:4px;padding-right:4px;cursor:pointer;" onClick='location.reload(true)' title='Если Beast включен, то можно нажимать.'>Поехали</span><br>
2. Дождаться окончания и автоматического выключения Beast.<br>
<br>



<div id='div_id' style='font-family:courier;font-size:16px;display:block;'><span style="font-size:18px;color:red;"><b>Нужен коннект с Beast.</b></span> Включите Beast на Пульте и <a href='/pages/mirrors_automatizm_maker.php'>перезагрузите эту страницу</a>.</div>


<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
// ждем пока не включат бестию
check_Beast_activnost(4);// после 4-го пульса И запускается get_info()

function get_info()
{
var linking_address='<?include($_SERVER["DOCUMENT_ROOT"]."/common/linking_address.txt");?>';
//wait_begin(); // wait_end();
var AJAX = new ajax_support(linking_address+"?stop_activnost=1",sent_blocing);
AJAX.send_reqest();
//var check_working_timer=setTimeout("check_working()",2000);
function sent_blocing(res)
{
//clearTimeout(check_working_timer);
document.getElementById('div_id').innerHTML="Идет процесс формирования зеркальных автоматизмов.";
wait_begin();
processing();
}
}
///////////////////////
var next=0;
function processing()
{
<?=$filesArr?>
//	alert("типа идет процесс");
/*
var AJAX = new ajax_support("/lib/get_file_content.php?file="+filesArr[next],sent_blocing);
AJAX.send_reqest();
//var check_working_timer=setTimeout("check_working()",2000);
function sent_blocing(res)
{
*/
//alert(res);
show_dlg_alert("<span style='font-size:20px;'><b>"+(next+1)+"</b> файл из "+fileCount,0);

//bot_contact("file_for_automayizm_making="+res,sent_process_mess);
//alert(filesArr[next]);
var AJAX = new ajax_support(linking_address + "?mirror_making_fool="+filesArr[next], sent_process_info);
AJAX.send_reqest();
function sent_process_info(res)
{
//alert(res);
if(res!="OK")
{
show_dlg_alert(res+"<hr>Ошибка файла "+filesArr[next]+"<hr>Процесс остановлен.",0);
wait_end();
return;
}

// следующий файл
next++;
if(next==filesArr.length)
end();
else
processing();
}

}

/////////////////////////////////////////

function end()
{
wait_end();
document.getElementById('div_id').innerHTML="Закончен процесс формирования зеркальных автоматизмов.";
show_dlg_alert("Beast выключается для корректного сохранения информации.",2000);
var AJAX = new ajax_support(linking_address+"?bot_closing=1",sent_bot_closing);
AJAX.send_reqest();
function sent_bot_closing(res)
{
	// не будет ответа

}
}
///////////////////////////////////

</script>


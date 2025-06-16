<?
/* Заполнить общий шаблон пусковых символов для http://go/pages/condition_reflexes_basic_phrases.php

http://go/pages/condition_reflexes_basic_phrases_common.php
*/
$page_id = -1;
$title = "Заполнить общий шаблон пусковых символов";
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");
//include_once($_SERVER['DOCUMENT_ROOT']."/pult_js.php");
//////////////////////////////////////////////////////////////


if(isset($_POST['gogogo'])&&$_POST['gogogo']==1)
{
$out="";
foreach($_POST['ids'] as $id => $str)
{
$out.=$str."|".$_POST['phrase'][$id]."\r\n";
}

//exit("$out");
writing_file($_SERVER["DOCUMENT_ROOT"]."/lib/condition_reflexes_basic_phrases_common.txt",$out);

echo "<form name=\"refresh\" method=\"post\" action=\"/pages/condition_reflexes_basic_phrases_common.php\"></form>";
echo "<script language=\"JavaScript\">document.forms['refresh'].submit();</script>";
exit();
}
//////////////////////////////////////////////

// Пусковые стимулы
$progs = reading_file($_SERVER["DOCUMENT_ROOT"] . "/pages/combinations/list_triggers.txt");
$progs=substr($progs,strpos($progs,"\r\n")+2); // exit("$progs");
$aArr = explode("\r\n", $progs);
$triggerArr=array();
foreach ($aArr as $str) {
	if(empty($str))
		continue;
$p = explode("|", $str);  
$triggerArr[$p[0]]=$p[1];
}
// var_dump($triggerArr);exit();


///////////////////////////////////////
// имеющиеся фразы
$id_list = str_replace(";",",",$id_list);
$file=$_SERVER["DOCUMENT_ROOT"]."/lib/condition_reflexes_basic_phrases_common.txt";
//exit("$file");
$progs = reading_file($file);
$strArr = explode("\r\n", $progs);
$phraseArr=array();
	foreach ($strArr as $str) {
		if (empty($str))
			continue;
		$p = explode("|", $str);
		if (empty($p[1]))
			continue;
		$phraseArr[$p[0]]=$p[1];
	}
//  var_dump($phraseArr);exit();


// проверка неповторяемости слов, иначе у.рефлекс будет неопределнным
$wArr=array();
$badArr="";
$repeatedArr=array();
foreach($phraseArr as $str)
{
if(in_array($str,$wArr))
{
$badArr.=$str."; ";
array_push($repeatedArr,$str);
}
array_push($wArr,$str);
}
if(!empty($badArr))
{
echo "<b><span style='color:red'>Есть повторяющиеся фразы: ".$badArr."</span></b>";
}
///////////////////////////////////////////////////////////////////////


$out="<table class='main_table' cellpadding=0 cellspacing=0 border=1 width='700px'>
		<tr>
			<th width=360  class='table_header'>Пусковые стимулы рефлекса</th>
			<th  class='table_header'>Фраза-синоним</th>
		</tr>";

$nid=0;
foreach ($triggerArr as $index => $resArr)
{
//	var_dump($resArr);exit();
$out.="<tr class='r_table highlighting' style='background-color:#eeeeee;' onClick='set_sel(this,`" . $index . "`)'>";

// пусковые стимулы
$out.="<td ><input type='hidden' name='ids[]' value='".$index."'><nobr>".$resArr."</nobr></td>";

// фраза-синоним
$phrase=get_prase_exists($index); // exit("$phrase");
$bg="";
if(in_array($phrase,$repeatedArr))
{
$bg="style='background-color:#FFFFAA;'";
}
$out.="<td  class='table_cell'><input id='insert_".$nid."' name='phrase[]' class='table_input' type='text' value='".$phrase."' ".$bg."><img src='/img/down17.png' class='select_control' onClick='show_word_list(".$nid.")' title='Выбор слов'></td>";

$out.="</tr>";
$nid++;
}
$out.="</table>";

$out.="<br><input type='submit' value='Сохранить' >";


/////////////////////////////////////////////////////////
echo "<div style='font-size:16px;' onClick='location.href=`/pages/condition_reflexes_basic_phrases_common.php`'><b>Фразы должны быть уникальны</b> для таблицы, иначе условный рефлекс окажется неопределенным!</div>";
echo 'Cохранение по Ctrl+S<form id="form_id" name="form" method="post" action="/pages/condition_reflexes_basic_phrases_common.php">';
echo $out;
echo "<input type='hidden' name='gogogo' value='1'>
</form>";




//////////////////////////////////////
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/alert2_dlg.php");
?>
<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
function set_sel(tr, id) {
		//	alert(id);
		var nodes = document.getElementsByClassName('highlighting'); //alert(nodes.length);
		for (var i = 0; i < nodes.length; i++) {
			nodes[i].style.border = "solid 1px #000000";
		}
		tr.style.border = "solid 2px #000000";
	}
function show_word_list(id)
{  
var AJAX = new ajax_support("/lib/get_exclamations.php?id=" + id, sent_words_info);
AJAX.send_reqest();
function sent_words_info(res) {
			//alert(res);
show_dlg_alert2("<br><span style='font-weight:normal;'>Щелкните по слову:<br>" + res + "<br>", 0);
}
}
function insert_word(id,word)
{
end_dlg_alert2();
	word="    <br>"+word+" ";
	word=word.trim();
	word=word.replace(/<\/?[^>]+>/g,'');
end_dlg_alert2();
var inp=document.getElementById('insert_'+id);
inp.setRangeText(word, inp.selectionStart, inp.selectionEnd, "end");
/*
if(inp.value.length>0)
inp.value+=" ";
document.getElementById('insert_'+id).value=inp+word;
*/
//alert(word);
}


	// сохранение по Ctrl+S
var is_press_strl = 0;
document.onkeydown = function(event) { 
		var kCode = window.event ? window.event.keyCode : (event.keyCode ? event.keyCode : (event.which ? event.which : null))

		//alert(kCode);
		if (kCode == 17) // ctrl
			is_press_strl = 1;

		if (is_press_strl) {
			if (kCode == 83) {
				event.preventDefault();
				//alert("!!!!! ");
				save_CTRRLS();
				is_press_strl = 0;
				return false;
			}
		}
}
////////////////////
document.onmouseup = function(event) { 
var t = event.target || event.srcElement;    
while(t)
{ 
if(t.id == "div_dlg_alert2")
	 return;	    
t = t.offsetParent;
}
end_dlg_alert2(); 
}
///////////////////////////
function save_CTRRLS()
{
show_dlg_confirm("Сохранить список?",1,-1,prases_saver);
}
function prases_saver()
{
document.forms.form_id.submit();
}
</script>
<?
///////////////////////////////////////////
function get_prase_exists($index)
{
global $phraseArr; //exit("$id");
//echo "$bsID | $id_list | $actions<br>";

if(isset($phraseArr[$index]))
	return $phraseArr[$index];// вернуть фразу

return "";
}
///////////////////////////////////////////////////
function get_actions($trArr)
{
	global $rActionsArr;
$acts="";
$aArr=explode(",",$trArr); 
foreach($aArr as $a)
{
	if(empty($a))
		continue;
	if(!empty($acts))
		$acts.=", ";
$acts.=$a." ".$rActionsArr[$a]."";
}
return $acts;
}
///////////////////////////////////////////////////
function reading_file($file)
{
if(!file_exists($file))
	return "";
if(filesize($file)==0)
	return "";
$hf=fopen($file,"rb");
if($hf)
{
$contents=fread($hf,filesize($file));
fclose($hf);
return $contents;
}//if($hf)
return "";
}
///////////////////////////////////////////////////
///////////////////////////////////////////////////
function writing_file($file,$content)
{
$hf=fopen($file,"wb+");
if($hf)
{
fwrite($hf,$content,strlen($content));
fclose($hf);
chmod($file, 0666);
return 1;
}
return 0;
}
//////////////////////////////////
?>
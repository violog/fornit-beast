<?
/* сформировать таблицу для http://go/pages/reflexes.php 

/pages/reflexes_maker_table.php?bsID=1&id_list=1;8
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$bsID=$_GET['bsID'];
$id_list=$_GET['id_list'];




////// Собрать данные по существующим рефлексам
$progs = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/dnk_reflexes.txt");
$strArr = explode("\r\n", $progs);  //var_dump($strArr);exit();
$reflexArr=array();
foreach ($strArr as $str) {
	if (empty($str))
		continue;

$par = explode("|", $str);
$par[2]=str_replace(",",";",$par[2]);

if($par[1]!=$bsID || $par[2]!=$id_list)
	continue;
array_push($reflexArr,$par);
}
//var_dump($reflexArr);exit();

//$resArr=get_reflex_exists("1","1;8","1;3");var_dump($resArr);exit();
////////////////////////////////////////////////////////////////////

/*
///////////////////// список возможных действий
$rActionsArr=array();
$progs = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/terminal_actons.txt");
	$strArr = explode("\r\n", $progs);
	foreach ($strArr as $str) {
if (empty($str) || $str[0] == '#')
			continue;
$p = explode("|", $str);
$rActionsArr[$p[0]]=$p[1];
	}
// var_dump($rActionsArr);exit();
////////////////////////////////////////////////////////////////////
*/



$out="<table class='main_table' cellpadding=0 cellspacing=0 border=1 width='100%'>
		<tr>
		<th width=120 class='table_header'>ID (если есть)</th>
			<th  class='table_header'>ID (3) пусковых стимулов через запятую</th>
			<th  class='table_header'>ID действий через запятую</th>
		</tr>";

// строка без пускового стимула (общий рефлекс)

// есть ли такой рефлекс?
$resArr=get_reflex_exists($bsID,$id_list,"");
// класс r_table - только для идентификции при сохранении
$out.="<tr class='r_table' style='background-color:#eeeeee;'>";
$out.="<td >".$resArr[0]."</td>";
$out.="<td style='background-color:#FFB8B8;' title='Рефлекс - БЕЗ ДЕЙСТВИЙ!'><input type='hidden' value=''></td>";
//!!! не делать переносов!!! в <td ></td> должн быть только <input >!!!! т.к. при сохранении смотрим tr.cells[2].childNodes[0]
if(empty($resArr[0]))// еще нет рефлекса
$out.="<td class='table_cell'><input id='input_0' class='table_input' type='text'  " . only_numbers_and_Comma_input() . "  value='" . $resArr[1] . "' ><img src='/img/down17.png' class='select_control' onClick='show_actions_list(0)' title='Выбор действий'></td>";
else // старый рефлекс - нередактируемый список действий
$out.="<td ><input type='hidden' value='" . $resArr[1] . "'>" . $resArr[2] . "</td>";
$out.="</tr>";
////////////////////////////////////////////

include_once($_SERVER['DOCUMENT_ROOT'] . "/lib/actions_from_pult.php");

////// антагонисты пусковых стимулов из /lib/actions_from_pult.php
$antFromId=$actionsFromPultAntagonistsArr;  
//var_dump($antFromId);exit();
//var_dump($actionsFromPultArr);exit();
////////////////////////////////////////////

// все рабочие сочетания Пусковых стимулов
$nNumbers=count($actionsFromPultArr);  





// все рабочие сочетания Пусковых стимулов НЕ РАБОТАЕТ
//
//




// var_dump($triggersComb);exit();


// Пусковые стимулы
include_once($_SERVER['DOCUMENT_ROOT'] . "/lib/get_ubicum_combination.php");
$triggersComb=get_ubicum_combination($nNumbers);
////////////////////////////////////////////
// пустые строки - заготовки рефлексов: все возможные сочетания пусковых стимулов
$actionArr=array();// ID выбранных акций
foreach ($triggersComb as $nTArr)
{
$aList="";// набор акций в виде строки
$aArr=array();// набор акций в виде массива
foreach ($nTArr as $nT)
{
$idT=(int)($nT)+1;// id пускового стимула

// убрать антагонистов
$isAntagonist=0;
foreach($aArr as $g)
{ 
	//exit("$idT <hr> ".var_dump($antFromId[$g]));
// есть ли антагонизм к $idT в уже имеющихся $aArr
if(in_array($idT,$antFromId[$g]))
{  
$isAntagonist=1; //exit("!!!!!!!!!!! $idT");
}
}
if($isAntagonist)
	continue;

if(!empty($aList))
	$aList.=",";
$aList.=$idT;
array_push($aArr,$idT);

// убирать использованную $k1 из последующих (начинать следующую уже без $k1)
if(in_array($aList,$actionArr))
{
	continue;
}
// не более 3-х действий подряд: оставляем первый и 2 последних. Это еще и добавит сочетаний.
	$aList=reduce_list($aList);
	if(substr_count($aList, ',')<2)
		array_push($actionArr,$aList);
}
}
// var_dump($actionArr);exit();
/////////////////////////////////////////////////////////



// сохранять сочетания Пусковых стимулов в раб.файлах для треугольничков /pages/reflexes.php, вызывающих диалог выбора 
$list="";
foreach($actionArr as $str)
{
$s="";
$p = explode(",", $str);
foreach($p as $a)
{
	if(empty($a) || $a<0)
		continue;
if(!empty($s))
	{
	$s.=", ";
	}

	$s.=$a." ".$actionsFromPultArr[$a][0];
}
$list.=$str."|".$s."\r\n";
}
// записывать только если изменилось
$new=md5($list);
$hash=$new."\r\n";
$oldLisr=read_file($_SERVER["DOCUMENT_ROOT"]."/pages/combinations/list_triggers.txt");
$old=md5(substr($oldLisr,strpos($oldLisr,"\r\n")+2));
//exit("$new<br>$old");
if($new!=$old)
{ 
write_trigger_file($_SERVER["DOCUMENT_ROOT"]."/pages/combinations/list_triggers.txt",$hash.$list);
}
/*  ИСПОЛЬЗОВАНИЕ:
$progs = read_file($_SERVER["DOCUMENT_ROOT"] . "/pages/combinations/list_triggers.txt");
$progs=substr($progs,strpos($progs,"\r\n")+2); // exit("$progs");
$aArr = explode("\r\n", $progs);
$triggerArr=array();
//$triggerArr["_"]="";
foreach ($aArr as $str) {
	if(empty($str))
		continue;
$p = explode("|", $str);  
$triggerArr[$p[0]]=$p[1];
}
// var_dump($triggerArr);exit();
*/
//////////////////////////////////////////////////







////////////////////////////////////// вывод таблицы
$nid=1;

foreach ($actionArr as $list)
{
// есть ли такой рефлекс?
$resArr=get_reflex_exists($bsID,$id_list,$list);// вернуть ID и действия рефлекса
//var_dump($resArr);exit();
//if($nid==2)exit("$bsID, $id_list, $list");
//if($bsID=="1" && $id_list=="1;8" && $list=="1;3")exit("!!!!!");

$out.="<tr class='r_table highlighting' style='background-color:#eeeeee;' onClick='set_sel(this," . $id . ")'>";
$out.="<td >" . $resArr[0] . "</td>";
$out.="<td ><input type='hidden' value='" . $list . "'>".get_actions_names_list($list)."</td>";
if(empty($resArr[0]))// еще нет рефлекса
$out.="<td  class='table_cell'><input id='input_".$nid."' class='table_input' type='text'  " . only_numbers_and_Comma_input() . "  value='" . $resArr[1] . "' ><img src='/img/down17.png' class='select_control' onClick='show_actions_list(".$nid.")' title='Выбор действий'></td>";
else // старый рефлекс - нередактируемый список действий
$out.="<td >" . $resArr[2] . "</td>";
$out.="</tr>";

$nid++;
}
$out.="</table>";

$out.="<br><input type='button' value='Сохранить рефлексы' onClick='reflex_saver()'>";

echo "!".$out;
////////////////////////////////////////////////////////////////////


///////// не более 3-х действий подряд: оставляем первый и 2 последних.
function reduce_list($list)
{
	$e=explode(",",$list);
	$len=count($e);
	if($len<4)
		return $list;
	return $e[0].",".$e[$len-2].",".$e[$len-1];
}

///////////////////////////////// // есть ли такой рефлекс?
function get_reflex_exists($bsID,$id_list,$actions)
{
global $reflexArr;
//echo "$bsID | $id_list | $actions<br>";

foreach ($reflexArr as $reflex)
{
	if($reflex[1]==$bsID && $reflex[2]==$id_list && $reflex[3]==$actions)
		return array($reflex[0],$bsID,$reflex[4]);// вернуть ID и действия рефлекса
}

return array("","");
}

////////////////////////////////////////
// позволяет вводить только цифры,  и запятую
function only_numbers_and_Comma_input($limit=0)
{
$out = <<<EOD
onKeyDown='only_numbers_and_Comma_input(this,$limit)' onKeyUp='only_numbers_and_Comma_input(this,$limit)' onMouseUp='only_numbers_and_Comma_input(this,$limit)'
EOD;
return $out;
}
?>
<script>
function only_numbers_and_Comma_input(inp,limit)
{  
var val=inp.value;
inp.value=val.replace(/[^0-9,]/g,'');
if(limit>0)
	{
inp.value=inp.value.substr(0,limit);
	}
}
</script>

<?
/////////////////////////////////////////////////
function get_actions_names_list($list)
{
	global $actionsFromPultArr;
$out="";
	$arr=explode(",",$list);
	foreach($arr as $a)
	{
if(!empty($out))
	$out.=",&nbsp;&nbsp;&nbsp;&nbsp;";
$out.=$a."&nbsp;".$actionsFromPultArr[$a][0]."";
	}
return $out;
}
///////////////////////////////////////////////////
function read_file($file)
{
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
function write_trigger_file($file,$content)
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
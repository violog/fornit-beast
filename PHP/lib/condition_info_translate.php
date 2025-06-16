<?
/* расшифровать из
/lib/condition_info_translate.php?current_condition=1|2,5,8|11


текущее состояние current_condition в виде
1|2,5,8|11
Базовое состояние
Активные контексты
Пусковые стимулы

*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");


$current_condition=$_GET['current_condition'];  //exit("$current_condition");

$cArr=explode("|",$current_condition); // var_dump($cArr);exit();

$lev1=get_info_1($cArr[0]);
$lev2=get_info_2($cArr[1]);
$lev3=get_info_3($cArr[2]);

$inf=$lev1."<br>".$lev2."<br>".$lev3; 

echo $inf;





/////////////////////////////////////////////
function get_info_1($str)
{
	$str=trim($str);
	switch($str)
	{
case 1: return "1 Плохо"; break;
case 2: return "2 Норма"; break;
case 3: return "3 Хорошо"; break;
	} 
}
////////////////////////////////////////

function get_info_2($str)
{
include_once($_SERVER['DOCUMENT_ROOT'] . "/lib/base_context_list.php");
$inf="";
$pArr=explode(",",$str);
foreach($pArr as $s)
{
$s=trim($s);
	if(empty($s))
		continue;
if(!empty($inf))
	$inf.="; ";
$inf.=$s."&nbsp;".$baseContextArr[$s][0];
}
return $inf;
}
//////////////////////////////////////////

function get_info_3($str)
{
	if(empty($str))
	{
return "Любые действия или без действий.";
	}
include_once($_SERVER['DOCUMENT_ROOT'] . "/lib/actions_from_pult.php");
$inf="";
$pArr=explode(",",$str);
foreach($pArr as $s)
{
$s=trim($s);
	if(empty($s))
		continue;
if(!empty($inf))
	$inf.="; ";
$inf.=$s."&nbsp;".$actionsFromPultArr[$s][0];
}
return $inf;
}
?>
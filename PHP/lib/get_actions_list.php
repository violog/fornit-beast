<?
/*  Выдать контрол для выбора действий из списка ДЛЯ РЕФЛЕКСОВ
/lib/get_actions_list.php?selected=1,3
*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');

$selected=$_GET['selected'];

include_once($_SERVER['DOCUMENT_ROOT']."/common/common.php");


// 4-й уровень - ID Действий рефлекса: 
$vArr=array();
$idArr=explode(",",$selected);    //  var_dump($idArr);exit();
foreach($idArr as $s)
{
	$s=trim($s);
	if(empty($s))
		continue;
array_push($vArr,$s);
}

$out="<table border=0 style='width:800px;font-size:14px;'><tr>";

$progs=read_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/terminal_actons.txt");
$strArr=explode("\r\n",$progs);

	$nCol = 0;
	$n = 0;
foreach($strArr as $str)
{
if(empty($str) || $str[0]=='#')
	continue;
if ($nCol == 6) { 
	$out.="</tr><tr>";
	$nCol = 0;
}									//exit($str);
$p=explode("|",$str);
$id=$p[0];
$bg = "";
if ($id < 30) {
	$bg = "style='color:#B16DB4;'";
}

$out.="<td " . $bg . " align='left'><nobr><input class='chbx_identiser' type='checkbox' value='".$id."' "; if(in_array($id,$vArr))$out.="checked"; $out.=">-".$id."&nbsp;".$p[1]."</nobr></td>";
	
$nCol++;
$n++;
}
$out.="</tr></table>";

exit($out);

?>
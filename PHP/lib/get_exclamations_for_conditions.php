<?
/*  Выдать контрол для выбора слов для окна ввода фраз в пульте
http://go/lib/get_exclamations_for_conditions.php
*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');

$bsID=trim($_GET['basic']);
$id_list=trim($_GET['contexts']);
// exit("$bsID | $id_list");

//$bsID=1;$id_list="1,2,8";


$file=$_SERVER["DOCUMENT_ROOT"]."/lib/condition_reflexes_basic_phrases/".$bsID."_".str_replace(",","_",$id_list).".txt"; 
// exit("$file");
if(!file_exists($file))
exit("");

$progs=read_file($file);  //exit("$progs");
$strArr=explode("\r\n",$progs);
$wArr=array();
foreach($strArr as $str)
{
if(empty($str))
	continue;
							//	exit($str);
$p=explode("|",$str);
array_push($wArr,$p[5]);
}
$wArr = array_unique($wArr);
sort($wArr, SORT_STRING);
reset($wArr);                //var_dump($wArr);exit();

$out="<table border=0 style='width:800px;font-size:14px;'><tr>";
$nCol=0;
foreach($wArr as $word)
{
if(empty($word))
	continue;
if ($nCol == 6) { 
	$out.="</tr><tr>";
	$nCol = 0;
}									//exit($word);
$out.="<td align='left' style='cursor:pointer;' onClick='insert_pult_word(`".$word."`)'><nobr>".$word."</nobr></td>";
	
$nCol++;
$n++;
}
$out.="</tr></table>";


echo $out;




///////////////////////////////////////////////////
function read_file($file)
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
?>


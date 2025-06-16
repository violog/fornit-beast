<?
/* получить инфу из words_temp_arr.txt
/pages/words_temp_server.php?old_size=0
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$file=$_SERVER['DOCUMENT_ROOT']."/memory_reflex/words_temp_arr.txt";
$size=filesize($file);
$old_size=$_GET['old_size'];

if($size==$old_size)
exit("not|##|");
///////////////////////////////////////////////////

$wt=read_file($file);

//$wt=str_replace("\r\n","<br>",$wt);
$wtArr=explode("\r\n",$wt); //var_dump($wtArr);exit();
rsort($wtArr, SORT_NUMERIC);reset($wtArr);

$out="<table class='main_table'  cellpadding=4 cellspacing=0 border=1 width='100%'>";
$n=0;
foreach($wtArr as $line)
{
	if(empty($line))
		continue;
	$p=explode("|#|",$line);
	$inf=$p[1]."&nbsp;(".$p[0].")";
if($n%3==0)
{
if($n)
	$out.="</tr>";
$out.="<tr>";
}
$out.="<td>".$inf."</td>";

$n++;
}
$out.="</tr></table>";

echo $size."|##|".$out;

//////////////////////////////////////////////
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
?>
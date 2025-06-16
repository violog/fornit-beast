<?
/*  установить возраст. (AJAX)
/lib/set_life_time.php

http://go/lib/set_life_time.php?yeas=1&month=2&days=3
*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$yeas=$_GET['yeas']*3600*24*365;
$month=$_GET['month']*3600*24*30;
$days=$_GET['days']*3600*24;

$life_time=$yeas+$month+$days;

if($life_time>0)
{
write_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/life_time.txt",$life_time);
write_file($_SERVER["DOCUMENT_ROOT"]."/memory_psy/episod_memory.txt","");
write_file($_SERVER["DOCUMENT_ROOT"]."/memory_psy/dominanta.txt","");
// Обновить время жизни условных рефлексов
$content=read_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/condition_reflexes.txt");
$content=trim($content);
$list=explode("\r\n",$content);  // var_dump($list);exit(); 
$curP=(int)$life_time/(3600*24); // exit("! $life_time | ".$curP);
$out="";
foreach($list as $str)
{
if(empty($str) || $str[0]=='#')
	continue;

$p=explode("|",$str);
$p[6]=$p[7]=$curP;
$out.=$p[0]."|".$p[1]."|".$p[2]."|".$p[3]."|".$p[4]."|".$p[5]."|".$p[6]."|".$p[7]."\r\n";
}
write_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/condition_reflexes.txt",$out);





// TODO: Обновить время рождения Доминант  /lib/set_life_time.php!!!!!!!!!





exit("1");
}

echo "0";

///////////////////////////////////////////////////
function write_file($file,$content)
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
//////////////////////////////////
?>
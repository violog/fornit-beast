<?
/* Сохранить рефлексы табцы из http://go/pages/reflexes_maker.php 

/pages/reflexes_maker_saver.php
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");


if(0)// тестирование
{
$bsID=1;
$id_list="1,3";
$saveStr="1|12,15,18";
}
else
{
$bsID=$_POST['bsID'];
$id_list=$_POST['id_list'];
$saveStr=$_POST['saveStr'];
}
//exit("$bsID | $id_list | $saveStr");


$id_list = str_replace(";",",",$id_list); 
$id=0;// последний существующий iD
$progs = read_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/dnk_reflexes.txt");
$progs=trim($progs);
$strArr = explode("\r\n", $progs);
/*
for($n=count($strArr)-1;$n>0;$n--)// нужен перебор снизу т.к. м.б. пустые строки внизу
{
	$lastRstr=trim($strArr[$n]);
	if(empty($lastRstr))
		continue;
$p = explode("|", $lastRstr);  
$id=$p[0]+1;
break;
}  уже не нужен :) т.к. $progs=trim($progs);
*/
$p = explode("|", $strArr[count($strArr)-1]);  
$id=$p[0]+1;
//exit("! $id");

/////////////////////////////////////////////////////
$out=$progs;  //exit("! ".$progs[strlen($progs)-1]);
//if($progs[strlen($progs)-1]!="\r\n")
$out.="\r\n";
/////////////////////////////////////////////////////

$new="";
$rArr=explode("||",$saveStr);
foreach($rArr as $rp)
{
	if(empty($rp))
		continue;
$p=explode("|",$rp);

$new.=$id."|".$bsID."|".$id_list."|".$p[0]."|".$p[1]."\r\n";

$id++;
}
//exit("$new");

$out.=$new;
write_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/dnk_reflexes.txt",$out);

write_file($_SERVER["DOCUMENT_ROOT"]."/pages/dnk_reflexes_seved.txt","1");

echo "!";


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
//////////////////////////////////

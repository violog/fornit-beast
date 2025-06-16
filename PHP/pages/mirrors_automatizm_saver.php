<?
/* Сохранить рефлексы табцы из http://go/pages/mirrors_automatizm.php

trigger=xxxx&answer=rrrr&tonMood=|1,2|actions=23,43||
http://go/pages/mirrors_automatizm_saver.php
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$bsID=$_POST['bsID'];
$id_list=$_POST['id_list'];
$saveStr=$_POST['saveStr'];       

if(0)
{
$bsID=1;
$id_list="2";
$saveStr="непонятно|ну и что|1,4|23,45||";
}

// создать /lib/mirrors_basic_phrases если его нет
$dir=$_SERVER["DOCUMENT_ROOT"]."/lib/mirror_reflexes_basic_phrases/";
if(!is_dir($dir))
{
$mod=0755;
mkdir($dir, $mod);
}

$id_list = str_replace(";",",",$id_list); // exit("$id_list");

$out="#utf8 bom\r\n"; // первая строка пустая, чтобы в ней осталась метка кодировки "\xEF\xBB\xBF" utf8 bom
$rArr=explode("|||",$saveStr);
foreach($rArr as $rp)
{
	if(empty($rp))
		continue;
$p=explode("|",$rp);
// triggPhrase|baseID|ContID_list|answerPhrase|Ton,Mood|actions1,...
$out.=$p[0]."|".$bsID."|".$id_list."|".$p[1]."|".$p[2]."|".$p[3]."\r\n";

}
$out=trim($out);
$out="\xEF\xBB\xBF".$out; // utf8 bom
//exit("$out");

$file=$_SERVER["DOCUMENT_ROOT"]."/lib/mirror_reflexes_basic_phrases/".$bsID."_".str_replace(",","_",$id_list).".txt"; 
// exit("$file");

$old=read_file($file);

//exit(md5($old)."<br>".md5($out));

if(md5($old)==md5($out))
{
echo "Новых фраз нет.";
}
else
{
write_file($file,$out);
echo "!";
}
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
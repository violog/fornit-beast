<?
/*  все используемые фразы в шаблонах условных рефлексов

include_once($_SERVER['DOCUMENT_ROOT'] . "/pages/mirrors_automatizm_get_all_phrases.php");

foreach ($triggerPhraseArr as $tArr)
*/
$triggerPhraseArr=array();
$tdir=$_SERVER["DOCUMENT_ROOT"]."/lib/condition_reflexes_basic_phrases/";

/*
$n=0;
if($dh = opendir($tdir)) 
{ //exit("!!!");
while(false !== ($file = readdir($dh))) 
{		
if($file=="." || $file=="..")
	continue;
if(filesize($tdir.$file)>0)
	{
$tstr=read_t_file($tdir.$file);
$str=explode("\r\n",$tstr);
foreach($str as $s)
{
$p=explode("|",$s);
array_push($triggerPhraseArr,$p[5]);
}
$n++;
	}
}
closedir($dh);
}
*/
// просто считать файл
$tstr=read_t_file($_SERVER["DOCUMENT_ROOT"]."/lib/mirror_basic_phrases_common.txt");
$str=explode("\r\n",$tstr);
foreach($str as $s)
{
$p=explode("|",$s);
array_push($triggerPhraseArr,$p[0]);
}

$triggerPhraseArr=array_unique($triggerPhraseArr);  // var_dump($triggerPhraseArr);exit();


///////////////////////////////////////////////////
function read_t_file($file)
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
<?
/* подключение всего пакета функций
include_once($_SERVER['DOCUMENT_ROOT']."/common/common.php");
*/
include_once($_SERVER['DOCUMENT_ROOT']."/common/alert_confirm.php");
include_once($_SERVER['DOCUMENT_ROOT']."/common/show_waiting.php");
include_once($_SERVER['DOCUMENT_ROOT']."/common/spoiler.php");
include_once($_SERVER['DOCUMENT_ROOT']."/common/input.php");
include_once($_SERVER['DOCUMENT_ROOT']."/common/js.php");

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

?>
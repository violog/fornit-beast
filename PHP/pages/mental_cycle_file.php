<?
/* считать файл лога цикла
/pages/mental_cycle_file.php?file=
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$file=$_SERVER['DOCUMENT_ROOT']."/cycle_logs/".$_GET['file']; 

if(!file_exists($file))
{
echo "Нет файла лога для цикла ".$_GET['file'];
}
else
{
echo read_file($file);
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
?>
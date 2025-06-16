<?
/* получить инфу из phrase_tree.txt
/pages/phrase_tree_server.php?old_size=0
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$file=$_SERVER['DOCUMENT_ROOT']."/memory_reflex/word_tree.txt";
$size=filesize($file);
$old_size=$_GET['old_size'];

if($size==$old_size)
exit("not");
///////////////////////////////////////////////////




echo $size;

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
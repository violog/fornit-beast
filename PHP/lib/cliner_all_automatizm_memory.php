<?
/*  удалить содержимое в файлов /memory_psy/automatizm_images.txt и /memory_psy/automatizm_tree.txt

/lib/cliner_all_automatizm_memory.php

*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");
mb_http_input('UTF-8');
mb_http_output('UTF-8');
mb_internal_encoding("UTF-8");


cliner_file($_SERVER["DOCUMENT_ROOT"]."/memory_psy/automatizm_images.txt");
write_file($_SERVER["DOCUMENT_ROOT"]."/memory_psy/automatizm_tree.txt","1|0|1|0|0|0|0|0
2|0|2|0|0|0|0|0
3|0|3|0|0|0|0|0
");


///////////////////////////////////////////////////
function cliner_file($file)
{
$hf=fopen($file,"wb+");
if($hf)
{
fwrite($hf,"",0);
fclose($hf);
return 1;
}
return 0;
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
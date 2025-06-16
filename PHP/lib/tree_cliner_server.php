<?
/*   очистка деревьев слов и фраз
/lib/tree_cliner_server.php
*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");

write_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/phrase_tree.txt","");
write_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/word_tree.txt","");
write_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/words_temp_arr.txt","");


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
?>
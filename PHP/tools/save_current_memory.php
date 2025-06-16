<?
/*        запись архива памяти Beast
http://go/tools/save_current_memory.php


*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');


// запись текущего состояния памяти в архив в папку /bot_files_save/
$arc_file="CurrentMemory.zip";

// удалить прежнюю запись
@unlink($_SERVER["DOCUMENT_ROOT"]."/tools/bot_files_save/CurrentMemory.zip");

include_once($_SERVER["DOCUMENT_ROOT"]."/lib/pclzip_lib.php");
$fileout="bot_files_save/".$arc_file;
$archive = new PclZip($fileout);
$v_list = $archive->add("../memory_reflex/,../memory_psy/");


echo "!".$fileout;
?>
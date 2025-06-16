<?
/*       удаление файла архива
http://go/tools/delete_archive_server.php?file=


*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');

$file=$_SERVER["DOCUMENT_ROOT"]."/tools/bot_files_save/".$_GET['file'];

$res=unlink($file);

if($res)
echo "!";
else
echo $res;

?>
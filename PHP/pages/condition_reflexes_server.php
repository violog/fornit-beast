<?
/* Удаление записей рефлексов после выключения ГО
/pages/condition_reflexes_server.php
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$out = $_GET['out'];
$out = str_replace("|", "\r\n", $out);
write_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/condition_reflexes.txt", $_GET['out']);
echo $out;

function write_file($file, $content)
{
  $hf = fopen($file, "wb+");
  if ($hf) {
    fwrite($hf, $content, strlen($content));
    fclose($hf);
    chmod($file, 0666);
    return 1;
  }
  return 0;
}
?>
<?
/* Запись лога в файл /pult_consol.txt   - последнее - наверх
/pages/consol_server.php
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$info = urldecode($_POST['info']); //exit($info);

write_file($_SERVER["DOCUMENT_ROOT"] . "/pult_consol.txt", $info);

 echo $info;

function write_file($file, $content)
{
  $hf = fopen($file, "ab+");
  if ($hf) {
    fwrite($hf, $content, strlen($content));
    fclose($hf);
    chmod($file, 0666);
    return 1;
  }
  return 0;
}
?>
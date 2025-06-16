<?
/* Проверка перед удалением рефлексов
/pages/terminal_actions_delete.php
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

if (isset($_GET['akt'])) {
  $akt = (int)$_GET['akt'];
  $ref = IsUsedInReflexes($akt);
  if ($ref != 0) {
    echo "Действие [" . $akt . "] используется в рефлексе [" . $ref . "], удаление запрещено!";
    exit();
  }
}

if (isset($_GET['delete_id'])) {
  $deln = (int)$_GET['delete_id'];
  $id = (int)$_GET['id'];
  $str = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/terminal_actons.txt");
  $list = explode("\r\n", $str);
  $out = "";
  $n = 0;
  foreach ($list as $s) {
    if (empty($s)) {
      $n++;
      continue;
    }
    if ($n == $deln) {
      $n++;
      continue;
    }
    $out .= $s . "\r\n";
    $n++;
  }
  write_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/terminal_actons.txt", $out);
  echo "!";
}

function IsUsedInReflexes($akt)
{
  $str = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/dnk_reflexes.txt");
  $list = explode("\r\n", $str);
  foreach ($list as $s) {
    if (empty($s)) continue;

    $lst = explode("|", $s);
    $akt_list = explode(",", $lst[4]);
    foreach ($akt_list as $akt_dnk) {
      if (empty($akt_dnk)) continue;
      if ($akt == $akt_dnk) return $lst[0];
    }
  }
  return 0;
}

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

function read_file($file)
{
  if (filesize($file) == 0)
    return "";
  $hf = fopen($file, "rb");
  if ($hf) {
    $contents = fread($hf, filesize($file));
    fclose($hf);
    return $contents;
  }
  return "";
}

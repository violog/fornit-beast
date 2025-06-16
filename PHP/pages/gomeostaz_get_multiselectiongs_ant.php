<?
/*  Для заполнения полей ввода "Несовместимость активностей Базовых стилей" в http://go/pages/gomeostaz.php
/pages/gomeostaz_get_multiselectiongs_ant.php?id=1
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');

$contextID = $_GET['id'];

include_once($_SERVER['DOCUMENT_ROOT'] . "/common/common.php");

$mArr = array(1 => "1 Пищевой", 2 => "2 Поиск", 3 => "3 Игра", 4 => "4 Гон", 5 => "5 Защита", 6 => "6 Лень", 7 => "7 Ступор", 8 => "8 Страх", 9 => "9 Агрессия", 10 => "10 Злость", 11 => "11 Доброта", 12 => "12 Сон",);

$out = "";

// считать файл 
$progs = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/base_context_antagonists.txt");
$strArr = explode("\r\n", $progs); // var_dump($strArr);exit();

foreach ($strArr as $str) {
  $par = explode("|", $str);
  $id = $par[0];
  if ($contextID == $id) {
    $parsAtr = explode(",", $par[1]);
    $aList = array();
    foreach ($parsAtr as $p) {
      array_push($aList, $p);
    }
    //var_dump($aList);exit();
    $out .= "Антагонисты для контекста &quot;" . $mArr[$contextID] . "&quot;:<br><select id='select_antagonist' multiple='multiple' size=8 style='width:300px;padding:4px;'>";
    foreach ($mArr as $id => $name) {
      $out .= "<option id='" . $id . "' value='" . $id . "'";
      if (in_array($id, $aList)) $out .= "selected";
      $out .= ">" . $name . "</option>";
    }
    $out .= "</select>";
    exit($out);
  }
}
?>
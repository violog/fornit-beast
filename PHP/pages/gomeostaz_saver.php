<?
/* сабмит форм http://go/pages/gomeostaz.php 2 и 3
/pages/gomeostaz_saver.php
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$n_form=$_POST['gogogo'];   //exit("! $n_form");

/// Несовместимость активностей Базовых стилей
if($n_form==4)
{
$out = "";
foreach ($_POST['id'] as $id => $str) {
	$id = trim($str);
	$id = preg_replace('/[^0-9]/', '', $id);

$out .= $id;
		$out .= "|" . $_POST['ant'][$id];
		$out .= "\r\n";

}
write_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/base_context_antagonists.txt", $out);
}
// Активности Базовых стилей
if($n_form==3)
{
$out = "";
foreach ($_POST['id'] as $id => $str) {
	$id = trim($str);
	$id = preg_replace('/[^0-9]/', '', $id);

$out .= $id;
		$out .= "|" . $_POST['bad'][$id];
		$out .= "|" . $_POST['well'][$id];
		$out .= "|" . $_POST['d1'][$id];
		$out .= "|" . $_POST['d2'][$id];
		$out .= "|" . $_POST['d3'][$id];
		$out .= "|" . $_POST['d4'][$id];
		$out .= "|" . $_POST['d5'][$id];
		$out .= "\r\n";

}
write_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/base_context_activnost.txt", $out);
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

echo "";
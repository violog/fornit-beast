<?
/* получить инфу из /memory_psy/automatizm_tree.txt
/pages/automatizm_tree_checker.php
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$file = $_SERVER['DOCUMENT_ROOT'] . "/memory_psy/automatizm_tree.txt";
$size = filesize($file);
echo $size;

function read_file($file)
{
	if (filesize($file) == 0)
		return "";
	$hf = fopen($file, "rb");
	if ($hf) {
		$contents = fread($hf, filesize($file));
		fclose($hf);
		return $contents;
	} //if($hf)
	return "";
}

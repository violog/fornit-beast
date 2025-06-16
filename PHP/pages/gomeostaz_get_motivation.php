<?
/* /pages/gomeostaz_get_motivation.php?id=1 */

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$actID = $_GET['id'];

$fa = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/Gomeostaz_pult_actions.txt");
$strArr = explode("\r\n", $fa);
foreach ($strArr as $str) {
	$par = explode("|", $str);
	$id = $par[0];
	if ($actID == $id) {
		exit("" . $par[2]);
	}
}
echo "";

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
?>
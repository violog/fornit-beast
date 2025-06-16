<?
/* Получение лога из файла /pult_consol.txt
/pages/consol_server.php
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$info = read_file($_SERVER["DOCUMENT_ROOT"] . "/pult_consol.txt");

echo $info;

function read_file($file)
{
	if (filesize($file) == 0)
		return "";
	$hf = fopen($file, "rb");
	if ($hf) {
		$contents = fread($hf, filesize($file));
		fclose($hf);

// последние фразы вверху
$out="";
$lines=explode("<br><separ>\r\n",$contents);
for($i=count($lines)-1;$i>=0;$i--)
{
$out.=$lines[$i]."<hr style='width:800px;border:solid 1px #dddddd;background-color:#dddddd;' align='left'>";
}


		return $out;
	} //if($hf)
	return "";
}
?>
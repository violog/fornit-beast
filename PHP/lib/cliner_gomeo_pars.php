<?
/*  обнулить гомео-параметры. (AJAX)
/lib/cliner_gomeo_pars.php

*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$value=$_GET['value'];


write_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/GomeostazParams.txt",
"1|".$value."
2|0
3|0
4|0
5|0
6|0
7|0
8|0");

echo "1";


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
?>
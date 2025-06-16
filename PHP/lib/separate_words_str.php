<?
/*
разделить строку на слова: 
prepare_str($str)

Нужно иметь:
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");
mb_http_input('UTF-8');
mb_http_output('UTF-8');
mb_internal_encoding("UTF-8");

*/

function prepare_str($text)
{
// выделить слова
$text=mb_strtolower($text);  // exit("! $text");

// символы, которые нужно выделять
$text=preg_replace('/([\(\)\+\=\-\?\!\[\]\{\}\<\>\.\,\/: ])/u','|\\1|',$text);

$text=preg_replace('/\r\n/','|#|',$text);
$text=preg_replace('/\n/','|#|',$text);

$text=str_replace("||","|",$text);

//exit("!!!! $text");

return $text;
}
///////////////////////////////
?>
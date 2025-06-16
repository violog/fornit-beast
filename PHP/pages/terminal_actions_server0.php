<?
/* ПРоверка перед записью и запись рефлексов
Проверяется уникальность сочетания условий чтобы не дублировались.
/pages/terminal_actions_server.php
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");


$out=""; 
//var_dump($_POST);exit();
$n=0;
foreach($_POST['id'] as $id => $str)
{
$id=trim($str);
$id=preg_replace('/[^0-9]/','',$id);

$decr=trim($_POST['decr'][$id]);

$val=trim($_POST['val'][$id]);
$val=preg_replace('/[^0-9>.;-]/','',$val);

$target=trim($_POST['target'][$id]);
$target=preg_replace('/[^0-9>,]/','',$target);

// проверка синтаксиса для строки действий:
                                             //  $val="1>-1.0;2>2.>";
$as=explode(";",$val);
$m=0;
foreach($as as $act)
{
	if(empty($act))
		continue;
if(strpos($act,">")===false)
{
exit("Ошибка в строке действий c ID=".$id.": неверный синтаксис");
}
$n1=substr_count($act, '>'); 
if($n1>1)
{
exit("Ошибка в строке действий c ID=".$id.": лишние символы &gt;.");
}

// if($m==1)exit("$act");
$p=explode(">",$act); 
$ext=preg_replace('/[0-9]/','',$p[0]); //exit("$ext");
if(strlen($ext))
{
exit("Ошибка в строке действий c ID=".$id.": лишние символы номера Параметра гомеостаза.");
}
$ext=preg_replace('/[0-9.-]/','',$p[1]); //exit("$ext");
if(strlen($ext))
{
exit("Ошибка в строке действий c ID=".$id.": лишние символы в числе Параметра гомеостаза.");
}
// if($m==1)exit("{$p[0]} | {$p[0]}");
if(empty($p[0]) || empty($p[1]))
{
exit("Ошибка в строке действий c ID=".$id.": неверный синтаксис.");
}
$m++;
}

$out.=$id."|".$decr."|".$val."|".$target."\r\n";

}



write_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/terminal_actons.txt",$out);

echo "!";

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
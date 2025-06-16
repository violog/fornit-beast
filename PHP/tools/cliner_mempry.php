<?
/* сбросить память до младенческого состяония, т.е. до безусловных рефлексов.
/tools/cliner_mempry.php
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");


/*
// очистить файлы:
write_empty($_SERVER['DOCUMENT_ROOT']."/memory_reflex/base_style_images.txt");
write_empty($_SERVER['DOCUMENT_ROOT']."/memory_reflex/condition_reflexes.txt");
write_empty($_SERVER['DOCUMENT_ROOT']."/memory_reflex/trigger_stimuls_images.txt");

// скопировать файлы:
copy($_SERVER['DOCUMENT_ROOT']."/tools/memory_reflex0/dnk_reflexes.txt", $_SERVER['DOCUMENT_ROOT']."/memory_reflex/dnk_reflexes.txt");
copy($_SERVER['DOCUMENT_ROOT']."/tools/memory_reflex0/phrase_tree.txt", $_SERVER['DOCUMENT_ROOT']."/memory_reflex/phrase_tree.txt");
copy($_SERVER['DOCUMENT_ROOT']."/tools/memory_reflex0/word_tree.txt", $_SERVER['DOCUMENT_ROOT']."/memory_reflex/word_tree.txt");
copy($_SERVER['DOCUMENT_ROOT']."/tools/memory_reflex0/terminal_actons.txt", $_SERVER['DOCUMENT_ROOT']."/memory_reflex/terminal_actons.txt");
copy($_SERVER['DOCUMENT_ROOT']."/tools/memory_reflex0/GomeostazParams.txt", $_SERVER['DOCUMENT_ROOT']."/memory_reflex/GomeostazParams.txt");


// теперь - просто залить архив /tools/bot_files_save/birthday.zip
*/

// очистить все в папке memory_psy
$dir=$_SERVER['DOCUMENT_ROOT']."/memory_psy/";
if($dh = opendir($dir)) 
{ //exit("!!!");
while(false !== ($file = readdir($dh))) 
{		
if($file=="." || $file=="..")
	continue;

	if(!is_dir($dir.$file))
	{	
write_empty($dir.$file);
	}
}
closedir($dh);
}



echo "!";

///////////////////////////////////////////////////
function write_empty($file)
{
$hf=fopen($file,"wb+");
if($hf)
{
fwrite($hf,"",0);
fclose($hf);
chmod($file, 0666);
return 1;
}
return 0;
}
///////////////////////////////

?>
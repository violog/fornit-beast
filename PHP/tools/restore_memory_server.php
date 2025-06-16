<?
/*       восстановление файлов памяти из файл архива
http://go/tools/restore_memory_server.php?file=2022_06_11_09_45.zip


*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');

$file=$_GET['file']; 

// очистить папки
scan_dir_del($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/");
scan_dir_del($_SERVER["DOCUMENT_ROOT"]."/memory_psy/");

include_once($_SERVER["DOCUMENT_ROOT"]."/lib/pclzip_lib.php");
$archive = new PclZip("bot_files_save/".$file);  //exit("! $archive");

//залить файлы прямо в /tools/ с теми дерикториями, что есть в архиве,
$list = $archive->extract(PCLZIP_OPT_PATH,"../tools/");  

//var_dump($list);exit(); 

/*
scan_dir_del($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/");
copy_works_files("/bot_files_TEMP/memory_reflex/","/memory_reflex/");

scan_dir_del($_SERVER["DOCUMENT_ROOT"]."memory_psy/");
copy_works_files("/bot_files_TEMP/memory_psy/","/memory_psy/");
*/

echo "!";
/////////////////////////////////////////////


/////////////////////////////////////////////////
/*
function copy_works_files($dir,$dirto)
{
if(!is_dir($_SERVER['DOCUMENT_ROOT'].$dir))
	return "";
if($dh = opendir($_SERVER['DOCUMENT_ROOT'].$dir)) 
{ 
while(false !== ($file = readdir($dh))) 
{		
if($file=="." || $file=="..")
	continue;
if(!is_dir($_SERVER['DOCUMENT_ROOT'].$dir.$file))
{ 
//$out.="<a href='".$dir.$file."' target='_blank'>".$dir.$file."</a><br>";
//$fc=read_file($_SERVER['DOCUMENT_ROOT'].$dir.$file);
//write_file($dirto.$file,$fc);
copy($_SERVER['DOCUMENT_ROOT'].$dir.$file, $_SERVER['DOCUMENT_ROOT'].$dirto.$file);
}
}
closedir($dh);
}
}
*/
//////////////////////////////////////////////
function scan_dir_del($dirname) 
{ 
    // Открываем текущую директорию 
$dir = opendir($dirname); 
    // Читаем в цикле директорию 
while(false !== ($file = readdir($dir))) 
{ 
      // Если файл обрабатываем его содержимое 
      if($file != "." && $file != "..") 
      { 
        // Если имеем дело с файлом - удаляем его 
        if(is_file($dirname."/".$file)) 
        { 
          unlink($dirname."/".$file); 
        } 
        // Если перед нами директория, вызываем рекурсивно 
        // функцию scan_dir_del 
        if(is_dir($dirname."/".$file)) 
        { 
          scan_dir_del($dirname."/".$file); 
          // После чего удаляем пустую директорию 
          rmdir($dirname."/".$file); 
        } 
      } 
} 
    // Закрываем директорию 
closedir($dir); 
}
?>
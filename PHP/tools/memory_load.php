<?
/*        перезапись памяти Beast из файла
http://go/tools/memory_load.php


*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');

//exit("!!!");

$sortArr=array();
if($dh = opendir($_SERVER["DOCUMENT_ROOT"]."/tools/bot_files_save/")) 
{ //exit("!!!");
$n=0;
while(false !== ($file = readdir($dh))) 
{		
if($file=="." || $file=="..")
	continue;
if(!is_dir($file))
{
$ext=substr($file,strrpos($file,'.'));
$ext=strtolower($ext);
if($ext!=".zip" && $ext!=".ZIP" )
{
continue;
}

//if($n==1) exit("! $file");
if($file=="CurrentMemory.zip")
{
$str= "<tr class='highlighting'><td class='archive_list' onClick='restore_archive(`".$file."`)' title='Залить эту память Beast'><li>CurrentMemory</td><td>

&nbsp;&nbsp;&nbsp;&nbsp;</td><td>
&nbsp;</td></tr>";
}
else
{
$arr=explode("_",$file);
if(count($arr)>5)
{
$time=mktime((int)$arr[4], 0, (int)$arr[6], (int)$arr[2], (int)$arr[1], (int)$arr[2] );
$txt="<li>".$arr[0]."-".$arr[1]."-".$arr[2]." &nbsp;&nbsp;".$arr[3].":".$arr[4];
}else
$txt="<li>".$file;
$txt=str_replace(".zip","",$txt);

$str= "<tr class='highlighting'><td class='archive_list' onClick='restore_archive(`".$file."`)' title='Залить эту память Beast'>".$txt."</td><td>

&nbsp;&nbsp;&nbsp;&nbsp;</td><td>
<span class='archive_list' title='удалить данный архив' onclick='remove_archive(`".$file."`)');'>
<img src='/img/delete.gif' border=0 title='Удалить архив памяти Beast'  style='cursor:pointer;'></span></td></tr>";
}

$sortArr[$n]=$time."|".$str;
$n++;
}

}
closedir($dh);
}
///////////////////
rsort($sortArr, SORT_NUMERIC);
reset($sortArr);                  //var_dump($sortArr);exit();
echo "!<table border=0>";
foreach($sortArr as $str)
{ 
echo substr($str,strpos($str,'|')+1);

}
echo "</table>";
///////////////////////////////////////////////////







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
</BODY>
</HTML>
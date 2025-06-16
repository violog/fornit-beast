<?
/* История общения показ

go/pages/history_show.php

*/
if(!isset($_GET['history_file']))
{
$title = "История общения";
}
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");


if(!isset($_GET['history_file']))
{
$dir=$_SERVER["DOCUMENT_ROOT"]."/history/"; 
$list="<div style='max-width:800px;
display: flex;
flex-wrap: wrap;
flex-direction: row;'>";//border:solid 1px #000000;
if($dh = opendir($dir)) 
{  
while(false !== ($file = readdir($dh))) 
{		
if($file=="." || $file=="..")
	continue;

if(!is_dir($dir.$file))
{
$f=substr($file,0,strrpos($file,'.'));
$list.="<div class='divF' onClick='show_file(`".$f."`)'>".$f."</div>";

}
}
closedir($dh);
}
$list.="</div>";
echo $list;
}
else// history_file
{
//echo $_GET['history_file'];
$file=$_SERVER["DOCUMENT_ROOT"]."/history/".$_GET['history_file'].".htm"; //exit("! $file");
$hf=fopen($file,"rb");
if($hf)
{
$contents=fread($hf,filesize($file));
fclose($hf);

$contents.="<img src='/img/history.png' style='position:fixed;z-index:1000;top:10px;right:10px;cursor:pointer;' title='Вернуться в список' onClick='location.href=\"/pages/history_show.php\"'>";

echo $contents;

}//if($hf)

}
/////////////////////////////////////////////////////////


?>
<style>
.divF{
#width: 30%;
margin-right:20px;
white-space: nowrap;
cursor:pointer;
}

</style>
<script>
function show_file(file)
{
//alert(file);
location.href='/pages/history_show.php?history_file='+file;
}
</script>
<?
/*  создать новый безусловный рефлекс
http://go/lib/create_reflex.php?aStr=1,3&conditions=2|6|

*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');

$aStr=$_GET['aStr'];
$conditions=$_GET['conditions'];
$conditions=str_replace(" ","",$conditions);


$content=read_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/dnk_reflexes.txt");
$content=trim($content);
// последний ID:
$list=explode("\r\n",$content);  //var_dump($list);exit(); 
$last=$list[count($list)-1]; //exit("! $last");
$p=explode("|",$last);  
$lastID=$p[0];  //exit("! $lastID");

// просмотреть все рефлексы и если найден с теми же условиями, то перекрыть его,
// а если нет, то дописать новый в конце.
$n=0;
$rNewList="";
$isRepaled=0;
foreach($list as $str)
{
if(empty($str) || $str[0]=='#')
	continue;

$p=explode("|",$str);
$cnd=$p[1]."|".$p[2]."|".$p[3];
if($cnd==$conditions)
{
$new=$p[0]."|".$conditions."|".$aStr."\r\n"; 
$rNewList.=$new; 
//echo "Было: ".$str." Стало: ".$new."<br>";
$isRepaled=1;
}
else
$rNewList.=$str."\r\n";
$n++;
}

if($isRepaled)
{
$content=$rNewList;
}
else
{
$newr=($lastID+1)."|".$conditions."|".$aStr;
//exit("! $newr");

$content.="\r\n".$newr; 
}

//exit("<hr>".$isRepaled."<hr>".$content);
$progs=write_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/dnk_reflexes.txt",trim($content));

echo "!";
///////////////////////////////////////////////////
function read_file($file)
{
if(filesize($file)==0)
	return "";
$hf=fopen($file,"rb");
if($hf)
{
$contents=fread($hf,filesize($file));
fclose($hf);
return $contents;
}//if($hf)
return "";
}
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
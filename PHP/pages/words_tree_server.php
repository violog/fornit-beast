<?
/* получить инфу из words_temp_arr.txt
/pages/words_tree_server.php?old_size=0
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$file=$_SERVER['DOCUMENT_ROOT']."/memory_reflex/word_tree.txt";
$size=filesize($file);
$old_size=$_GET['old_size'];

//if($size==$old_size)
//exit("not|##|");
///////////////////////////////////////////////////

$wt=read_file($file);

//$wt=str_replace("\r\n","<br>",$wt);
$wtArr=explode("\r\n",$wt); //var_dump($wtArr);exit();

$symbolsArr = [" ","а","б","в","г","д","е","ё","ж","з","и","й","к","л","м","н","о","п","р","с","т","у","ф","х","ц","ч","ш","щ","ъ","ы","ь","э","ю","я","a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z","!","?","@","#","$","%","^","&","*","(",")","+","=","-","1","2","3","4","5","6","7","8","9","0","[","]","{","}","<",">",".",",","/"];

$tree=array();
$n=0;
foreach($wtArr as $line)
{
	if(empty($line))
		continue;
	$p=explode("|#|",$line);
	$word=$p[1];
	$d=explode("|",$line);
	 /// первая цифра - ID узла, вторая - ID родителя
	$id=$d[0];
	$chid=$d[1];

//$tree[$n][$id]=$word;
$tree[$n]=array(0=>$chid,1=>$id,2=>$word);
$n++;
}
//var_dump($tree);exit();

$out="";
$level=0;
foreach($symbolsArr as $id => $s)
{  //if($id==2)exit("!!!!! $id");

$out.=get_word($id,$level,$tree,"");

}

function get_word($id,$level,$tree,$pre)
{   
//	if($id==2)exit("!!!!! $pre");
global $symbolsArr;
	$out="";
	$level0=$level;
foreach($tree as $ti)
	{	
$level=$level0;// для прохода данного уровня восстанавливаем $level
if($ti[0]==$id)
	{   //exit("!!!!!");

//if($id<count($symbolsArr) )echo("$level | {$ti[2]} <br>");

if($level==0)
{
	$out.=$symbolsArr[$id-1];//$id-1 т.к. дерево нициируется начиная с id==1
}

$str=$ti[2]."(".$ti[1].",".$ti[0].")";

// отсутп
$out.=str_repeat("&nbsp;",10*$level);
//$out.=str_repeat("&nbsp;",strlen($str));
//$out.="<span style='color:dddddd'>|".str_repeat("_",4)."</span>";

$level++;

$out.=$str."<br>";
$out.=get_word($ti[1],$level,$tree,$pre.$ti[2]);
	}

//$out.="#";
	}
return $out;
}
//exit($out);


echo $size."|##|".$out;

//////////////////////////////////////////////
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
?>
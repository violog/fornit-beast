<?
/*  Выдать контрол для выбора Пусковых стимулов из списка ДЛЯ РЕФЛЕКСОВ
/lib/get_triggers_list.php?nid=1&selected=1,3
*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');

$nid=$_GET['nid'];
$selected=$_GET['selected'];


// Пусковые стимулы
// Пусковые стимулы
$progs = read_file($_SERVER["DOCUMENT_ROOT"] . "/pages/combinations/list_triggers.txt");
$progs=substr($progs,strpos($progs,"\r\n")+2); // exit("$progs");
$aArr = explode("\r\n", $progs);
$triggerArr=array();
//$triggerArr["_"]="";
foreach ($aArr as $str) {
	if(empty($str))
		continue;
$p = explode("|", $str);  
$triggerArr[$p[0]]=$p[1];
}
// var_dump($triggerArr);exit();


foreach($triggerArr as $ids => $str)
{
	if(substr_count($str, ',')>1)// не более 2-х сочетаний контектосв!
	continue;

	$bg="";
	if($nsel==$n)
	{
		$bg="#cccccc";
//		exit("> $nsel");
	}
$out.="<div style='text-align:left;cursor:pointer;background-color:".$bg.";' onClick='set_input3_list(".$nid.",`".$ids."`)'>".$str."</div>";
$n++;
}

exit($out);

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
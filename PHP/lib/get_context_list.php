<?
/*  Выдать контрол для выбора Базовых контекстов из списка ДЛЯ РЕФЛЕКСОВ
/lib/get_context_list.php?selected=1,3
*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');

$nid=$_GET['nid'];  // exit("> $nid");
$selected=$_GET['selected'];


// реально возможные сочетания контекстов
$c_list = read_file($_SERVER["DOCUMENT_ROOT"] . "/pages/combinations/combo_contexts_str.txt");
$c_list=str_replace(";",",",$c_list);
$ContextIdArr=explode("\r\n",$c_list);  //var_dump($ContextIdArr);exit();
$nsel=0;
$n=0;
foreach($ContextIdArr as $str)
{
//	echo "$selected==$str <br>";
if($selected==$str)
	{
$nsel=$n; // exit("> $nsel");
	}

$n++;
}


$c_list = read_file($_SERVER["DOCUMENT_ROOT"] . "/pages/combinations/combo_contexts_names.txt");
$c_list=str_replace(";",",",$c_list);
$ContextnamesArr=explode("\r\n",$c_list); // var_dump($ContextnamesArr);exit();
$out="";
$n=0;
foreach($ContextnamesArr as $str)
{
	if(substr_count($str, ',')>2)// не более 3-х сочетаний контектосв!
	continue;
	$bg="";
	if($nsel==$n)
	{
		$bg="#cccccc";
		//exit("> $nsel");
	}
$out.="<div style='text-align:left;cursor:pointer;background-color:".$bg.";' onClick='set_input2_list(".$nid.",`".$ContextIdArr[$n]."`)'>".$str."</div>";
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
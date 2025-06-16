<?
/*  Выдать контрол для выбора действий из списка
/lib/get_action_choose.php?id=115

*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');

$id=$_GET['id'];

$vArr=array();
if(isset($_GET['id'])&&$_GET['id']>0)//список выделений акций
{
// считать файл 
$progs=read_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/dnk_reflexes.txt");
$strArr=explode("\r\n",$progs); // var_dump($strArr);exit();

$lev4="";
foreach($strArr as $str)
{      
$par=explode("|",$str);	 // exit("$id | ".$par[0]);
	
if($id==$par[0])
{
$lev4=$par[4];  // var_dump($par);exit();
}

$idArr=explode(",",$lev4);    //  var_dump($idArr);exit();
foreach($idArr as $s)
{
	$s=trim($s);
	if(empty($s))
		continue;
array_push($vArr,$s);
}
}
/////////////////////////////////////////////

}


$out="<select id='actions_combo' name='actions_combo[]' multiple='multiple' size=8 style='width:300px;padding:4px;'>";

$progs=read_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/terminal_actons.txt");
$strArr=explode("\r\n",$progs);
$n=0;
foreach($strArr as $str)
{
if(empty($str) || $str[0]=='#')
	continue;
$p=explode("|",$str);
$id=$p[0];
//$out.="<option id='".$p[0]."' value='".$p[0]."' selected>".$p[0].". ".$p[1]."</option>";
$out.="<option id='".$id."' value='".$id."'"; if(in_array($id,$vArr))$out.="selected";$out.=">".$id." ".$p[1]."</option>";

$n++;
}
$out.="</select>";

echo $out;


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
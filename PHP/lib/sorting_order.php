<?
/*
Сортировка чиселовых полей типа 2,5,8,-3,-4,-6,-11
Сначала идут положительные по возрастанию, потом - отрицательные - в порядке возрастания цифры

include_once($_SERVER["DOCUMENT_ROOT"]."/lib/sorting_order.php");


*/

// $nstr="9,2,5,8,-3,-4,-6,-11";$res=sorting_order($nstr);exit("$res");

function sorting_order($nstr)
{
$nArr = explode(",", $nstr);
$pArr=array();
$mArr=array();
foreach($nArr as $n)
{
if($n>=0)
	array_push($pArr,$n);
else
	array_push($mArr,$n);

}

sort($pArr, SORT_NUMERIC);
reset($pArr);
rsort($mArr, SORT_NUMERIC);
reset($mArr);
$out="";
foreach($pArr as $n)
{
	$out.=$n.",";
}
foreach($mArr as $n)
{
	$out.=$n.",";
}

$out=substr($out,0,strlen($out)-1);// убрать последнюю запятую
//exit("$out");
return $out;
}

/*
function sorting_order($nstr)
{
$nArr = explode(",", $nstr);
sort($nArr, SORT_NUMERIC);
reset($nArr);
$out="";
$mout="";
foreach($nArr as $n)
{
if($n>=0)
	$out.=$n.",";
else
	$mout.=$n.",";

}
$ret=$out.$mout;  
$ret=substr($ret,0,strlen($ret)-1);// убрать последнюю запятую
//exit("$ret");
return $ret;
}
*/
?>
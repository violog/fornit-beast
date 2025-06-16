<?
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

function ExistsValInArr($arr, $val){
	foreach($arr as $str){
		if($str==$val){
			return true;
		}
	}
	return false;
}

$aktArr=array(1=>"Непонятно (1)",2=>"Понятно (2)",3=>"Наказать (3)",4=>"Поощрить (4)",5=>"Накормить (5)",6=>"Успокоить (6)",7=>"Предложить поиграть (7)",8=>"Предложить поучить (8)",9=>"Игнорировать (9)",10=>"Сделать больно (10)",11=>"Сделать приятно (11)",12=>"Заплакать (12)",13=>"Засмеяться (13)",14=>"Обрадоваться (14)",15=>"Испугаться (15)",16=>"Простить (16)",17=>"Вылечить (17)",18=>"Разозлить (18)",);

$rArr=array();
$aArr=array();
$flg_break=false;

foreach($_POST['id'] as $id => $str){
	if($id==5) continue;
	if($flg_break==true) break;
	$rArr=trim($_POST['effect'][$id]);
	$bArr=explode(",",$rArr);
	$out1 ="В списке действий ".$aktArr[$id];

	for ($i=0; $i<count($bArr); $i++){
		$tArr=explode(">",$bArr[$i]);
		$ida=$tArr[0];
		if($ida>0 && !ExistsValInArr(array(1,2,3,4,5,6,7,8),$ida)){
			$out =$out1." указан несуществующий базовый параметр (".$ida.")!";
			$flg_break=true;
			break;
		}
		if(ExistsValInArr($aArr, $ida)){
			$out = $out1." есть дублирующее воздействие на базовый параметр ".$ida."!";
			$flg_break=true;
			break;
		}
		array_push($aArr, $ida);
	}
	$aArr=array();
}
if($out=="") $out="*";
exit($out);
?>
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

$kontArr=array(1=>"Пищевой (1)",2=>"Поиск (2)",3=>"Игра (3)",4=>"Гон (4)",5=>"Защита (5)",6=>"Лень (6)",7=>"Ступор (7)",8=>"Страх (8)",9=>"Агрессия (9)",10=>"Злость (10)",11=>"Доброта (11)",12=>"Сон (12)",);

$rArr=array();
$aArr=array();
foreach($_POST['id'] as $id => $str){
	if($flg_break==true) break;
	$rArr=trim($_POST['ant'][$id]);
	$antArr=explode(",",$rArr);
	for ($i=0; $i<count($antArr); $i++){
		$out1 ="В списке антагонистов ".$kontArr[$id];
		if(!ExistsValInArr(array(1,2,3,4,5,6,7,8,9,10,11,12),$antArr[$i])){
			exit($out1." указан не существующий контекст ".$antArr[$i]."!");
			$flg_break=true;
			break;
		}
		if($id==$antArr[$i]){
			exit("Нельзя указывать антагонистом контекста ".$kontArr[$id]." этот же контекст!");
			$flg_break=true;
			break;
		}
		if(ExistsValInArr($aArr, $antArr[$i])){
			exit($out1." есть дублер (".$antArr[$i].")!");
			$flg_break=true;
			break;
		}
		array_push($aArr, $antArr[$i]);
	}
	$aArr=array();
}
echo "*";
?>
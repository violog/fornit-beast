<?
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

function ExistsValInArr($arr, $val)
{
	foreach ($arr as $str) {
		if ($str == $val) {
			return true;
		}
	}
	return false;
}

function read_file($file)
{
	if (filesize($file) == 0)
		return "";
	$hf = fopen($file, "rb");
	if ($hf) {
		$contents = fread($hf, filesize($file));
		fclose($hf);
		return $contents;
	}
	return "";
}

function ChekValue($kArr, $baze_name, $zone_name)
{
	$aArr = array();
	$kontArr = array(1 => "Пищевой (1)", 2 => "Поиск (2)", 3 => "Игра (3)", 4 => "Гон (4)", 5 => "Защита (5)", 6 => "Лень (6)", 7 => "Ступор (7)", 8 => "Страх (8)", 9 => "Агрессия (9)", 10 => "Злость (10)", 11 => "Доброта (11)", 12 => "Сон (12)",);
	$progs = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/base_context_antagonists.txt");
	$strArr = explode("\r\n", $progs);

	$out="Для " . $baze_name . " в разделе \"" . $zone_name . "\"";
	$out_ant = "";
	for ($i = 0; $i < count($kArr); $i++) {
		$ap=abs($kArr[$i]);
		if (!ExistsValInArr(array(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12), $ap)) {
			if($ap<20 || $ap >40)
			exit($out . " указан не существующий контекст (" . $kArr[$i] . ")!|0");
		}
		if (ExistsValInArr($aArr, $kArr[$i])) {
			exit($out . " есть дублер (" . $kArr[$i] . ")!|0");
		}
		$val = $kArr[$i] * -1;
		if (ExistsValInArr($aArr, $val)) {
			exit($out . " указан один и тот же контекст (" . $kArr[$i] . "/" . $val . ") с разными знаками!|0");
		}
		array_push($aArr, $kArr[$i]);
	}

	for ($i = 0; $i < count($aArr); $i++) {
		$ida = $aArr[$i];
		if ($ida > 0) {
			$num = (int)$ida - 1;
			$AntArr = explode("|", $strArr[$num]);
			$antArr = explode(",", $AntArr[1]);
			foreach ($antArr as $val) {
				if (ExistsValInArr($aArr, $val)) {
					//return $out." указан контекст \"".$kontArr[$val]."\", при этом он же указан антагонистом в таблице Несовместимость активностей Базовых стилей для контекста ".$kontArr[$ida]."!";
					exit($out." указан контекст \"" . $kontArr[$val] . "\", заданный антагонистом в таблице \"Несовместимость активностей Базовых стилей\" для контекста \"" . $kontArr[$ida] . "\"!<br>");
				}
			}
		}
	}
	if ($out_ant != "") {
		return $out_ant . "|1";
	}
	return  "";
}

$zArr = array();
$zoneArr = array("bad", "well", "d1", "d2", "d3", "d4", "d5");
$baseArr = array(1 => "Энергия (1)", 2 => "Стресс (2)", 3 => "Гон (3)", 4 => "Потребность в общении (4)", 5 => "Потребность в обучении (5)", 6 => "Поиск (6)", 7 => "Самосохранение (7)", 8 => "Повреждения (8)",);
$baseZoneArr = array(0 => "Выход из нормы", 1 => "Возврат в норму", 2 => "Норма", 3 => "Слабое отклонение", 4 => "Значительное отклонение", 5 => "Сильное отклонение", 6 => "Критически опасное отклонение",);

foreach ($_POST['id'] as $id => $str) {
	if ($flg_break == true) break;
	for ($i = 0; $i < count($zoneArr); $i++) {
		$zArr = trim($_POST[$zoneArr[$i]][$id]);
		$kArr = explode(",", $zArr);
		$out_buf = ChekValue($kArr, $baseArr[$id], $baseZoneArr[$i]);
		if ($out_buf != "") {
			$strOut = explode("|", $out_buf);
			if ($strOut[1] == "0") {
				$flg_break = true;
				$out = $out_buf;
				break;
			} else {
				$out .= $strOut[0];
			}
		}
	}
}

echo "*";
?>
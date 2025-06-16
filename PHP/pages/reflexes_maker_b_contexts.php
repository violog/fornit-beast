<?
/* выдать возможные сочетания Базовых контекстов в зависимости от Базового состояния
для http://go/pages/reflexes.php 

/pages/reflexes_maker_b_contexts.php?base_condition=1

Cостояния отдельных параметров гомеостаза не коррелируют с общим Базовым состоянием (определяемым в ГО func commonBadDetecting()) и тут сложности, например, Базовое состояние Хорошо возникает после улучшения после плохих диапазонов гомео-параметров.
Поэтому невозможно вычислить сочетаний активных Базовых контекстов, дающие Базовые состояния.
Так что параметр $base_condition НЕ ИСПОЛЬЗУЕТСЯ.

Алгоритм:
1.	Создать массив всех возможных без внутренних повторений) сочетаний ячеек (СЯ) таблицы "Активности Базовых стилей".
2.	Начать перебор каждого из сочетаний СЯ и для каждого из них суммируем ID сочетаний контекстов этих ячеек. Получаем все возможные сочетания контекстов таблицы таблицы "Активности Базовых стилей". Из кажого суммарного сочетания убираем антагонистов, отрицательные контексты (которые должны госиться) и переводим сочетания контекстов в строки, оставляя только уникальные.
3.	В результате получаем все возможные сочетания контекстов в виде строк с разделителем “;”.

*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");


set_time_limit(0);

$current_base_condition=0;
if(isset($_GET['current_base_condition']))
$current_base_condition=$_GET['current_base_condition']; //  exit("> $current_base_condition");


$base_condition=$_GET['base_condition']; // НЕ ИСПОЛЬЗУЕТСЯ т.к. нет отпределенной зависимости
$get_list=$_GET['get_list'];  //exit($get_list);


// сгенерировать рабочие сочетания Базовых контекстов НЕ ПОЛУЧАЕТСЯ КОРРЕКТНО
//include_once($_SERVER['DOCUMENT_ROOT'] . "/pages/reflexes_maker_context_combinations.php");
/*
В ДАННЫЙ МОМЕНТ СКРИПТ /pages/reflexes_maker_context_combinations.php НЕ ИСПОЛЬЗУЕТСЯ вот почему.

Проверено немало алгоритмов фомрирования списков (наиболее эффективные для PHP собраны в array_combinations.php), но до сих пор ни один не используется вот почему:
1. время выполнения оказывается неприемлемо долгим при создании неповторяющихся комбинайиций из 8х7 ячеек. Многие алгоритмы вызывают ошибки недостатка памяти (даже, использующие yield PHP).
Даже в ГО профессиональный алгоритм работает неприемлемое время (tools.combinations_maker.go).
2. Наличие антагонистов и гасящих контекстов делает результат зависимым от способа обработки.

Поэтому сейчас используются файлы готовых списков в папке /pages/combinations/ составленные на основе ранее полученных вариантов и проверненные эвристически.
!!! В случае изменения таблицы "Активности Базовых стилей"(в http://go/pages/gomeostaz.php) 
наобходимо пересматривать списки 
/pages/combinations/combo_contexts_str.txt
и
/pages/combinations/combo_contexts_names.txt

В случае, если будет сделан генератор сочетаний, то он должен срабатывать при запуске из Публта, из менб Инструментов (шестеренка) и обновлять списки, а при каждом редактировании таблицы "Активности Базовых стилей" должно быть предупреждение о необходимости обновления списков.
*/

// использовать имеющиеся combo_contexts_str.txt и combo_contexts_names.txt
$idText = read_file($_SERVER["DOCUMENT_ROOT"] . "/pages/combinations/combo_contexts_str.txt");
//$nameText = read_file($_SERVER["DOCUMENT_ROOT"] . "/pages/combinations/combo_contexts_names.txt");

$contextsArr=array();
$cList = explode("\r\n", $idText);    // var_dump($cList);exit();
$n=0;
foreach($cList as $c)
{
	if(empty($c))
		continue;
if(substr_count($c, ',')>2)// не более 3-х сочетаний контектосв!
	continue;

//$contextsArr[$n]=array();
//$p = explode(";", $aArr);
$c=str_replace(",",";",$c);
array_push($contextsArr,$c);
}
//var_dump($contextsArr);exit();
///////////////////////////////////////////////////////////////
// Базовые контексты $baseContextArr - только для получения имен базовых контекстов
include_once($_SERVER['DOCUMENT_ROOT'] . "/lib/base_context_list.php");
//var_dump($baseContextArr);exit();



$rtxts = read_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/dnk_reflexes.txt");
$rtxts=trim($rtxts);
$rtArr = explode("\r\n", $rtxts);

// собрать комбобокс
$out="<select id='base_context_id' size=12 style='max-width:360px;'>";// multiple='multiple' 
foreach($contextsArr as $aArr)
{
	$str="";
	$p = explode(";", $aArr); //var_dump($p);exit();
	foreach($p as $a)
{
	if(empty($a) || $a<0)
		continue;
if(!empty($str))
	{
	$str.=",&nbsp;";
	}
$a=(int)$a;
	$str.=$a."&nbsp;".$baseContextArr[$a][0];
}

// сколько рефлексов сделано в этом сочетании
$combStr=preg_replace('/[^0-9,]/','',$str); //echo $combStr."<br>";
$rcount=0;
foreach($rtArr as $cur)
{
	$p = explode("|", $cur); 
	if($p[1]==$current_base_condition && $combStr==$p[2])
		$rcount++;
}
// exit($p[1]."==$current_base_condition && $combStr==".$p[2]);
//echo $combStr." $rcount <br>";






$out.="<option  value='".$aArr."' ";   //exit($get_list."<hr>".$aArr);
if(!empty($get_list) && $get_list==$aArr)
{
$out.="selected";
}

$out.=" title='".$str." (рефлексов: $rcount)'>".$str." (рефлексов: $rcount)</option>";
//	array_push($contextsNameArr,$str);
}
$out.="</select><br>";

//exit($out);
echo "!".$out;

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
<?
/*  сгенерировать рабочие сочетания Базовых контекстов

include_once($_SERVER['DOCUMENT_ROOT'] . "/pages/reflexes_maker_context_combinations.php");
для http://go/pages/reflexes_maker.php 

В ДАННЫЙ МОМЕНТ ЭТОТ СКРИПТ НЕ ИСПОЛЬЗУЕТСЯ вот почему.

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

//ini_set('memory_limit', '2024M');

///////////////////////////////////////////////////////
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
exit("НЕ ИСПОЛЬЗУЕТСЯ");

// антагонисты
$progs = read_combo_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/base_context_antagonists.txt");
$strArr = explode("\r\n", $progs);  //exit("$progs");
$antFromId = array();// антагонисты для каждого выбранного в списке $get_list ID контекста
foreach ($strArr as $str) {
	$par = explode("|", $str);
	$id = $par[0];
	$as = explode(",", $par[1]); 
	$antFromId[$id]=array();
	foreach ($as as $a)
	{			
	array_push($antFromId[$id],$a);
	}
}
// var_dump($antFromId);exit();

// Базовые контексты $baseContextArr - только для получения имен базовых контекстов
include_once($_SERVER['DOCUMENT_ROOT'] . "/lib/base_context_list.php");



////////// таблица Активности Базовых стилей
$progs = read_combo_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/base_context_activnost.txt");
$strArr = explode("\r\n", $progs);  //exit("$progs");
$contextArr = array();
foreach ($strArr as $str) {
	$par = explode("|", $str);
	$id = $par[0];

$contextArr[$id]=array();
	for($n=1;$n<8;$n++)
	{
	array_push($contextArr[$id],$par[$n]);
	}
}
// var_dump($contextArr);exit();
//////////////////////////////////////////////////////////////////////


//!!!! сделать все возможные сочетаия 12 контекстов, и потом из каждой поудалять антагонистов,
//??? в список антагонистов включить минусы из таблицы

ДОБАВЛЯТЬ ПЕРЕБОРЫ К УЖЕ СУЩЕСТВУЮЩЕМУ ПРЕДЫДУЩЕМУ, записывая результат в вызодной массив !!!

???
$outArr=array();// в выходной массив добавляется (array_push) $curArr с каждой итерацией, т.е. после каждой итерации он возрастает на 7*8=56 элементов. Всего в нем станет 7*56 элементов.
$curArr=array();// текущий суммирующий массив 56 элементов - с каждой итерацией накапливает цифры ID при сдвиге следующей строки с уборкой повторов

/* попытка повторить то, что в ГО
$nid=array();
$outArr=array();
$gCArr=$contextArr;
$cols=count($gCArr[1]);
$rows=8;
// собираются все значения переборов ячеек в одну строку
for ($i = 1; $i < $rows; $i++) {
	iterate($i);
}
function iterate($nrow)
{
	global $outArr,$gCArr,$cols;
	$outArr[$nrow]=array();
//for ($n = 0; $n < $cols; $n++) {
		$out="";
		for ($i = 0; $i < $cols; $i++) {
			if ($i>0){$out .=",";}
			$out .=$gCArr[$nrow][$i];
			// собираются все значения данного набора ячеек в одну строку
			//exit("$out");
			array_push($outArr[$nrow],$out);
		}
	//}
}
//var_dump($outArr);exit();

// теперь нужно перебрать эти 7 $outArr[$nrow] по сочетаниям без повторений 7 значений

/// все комбинации по 7 без повторов
function comb($m, $words) {
    if (!$m) {
        yield [];
        return;
    }
    if (!$words) {
        return;
    }
    $h = $words[0];
    $t = array_slice($words, 1);
    foreach(comb($m - 1, $t) as $c)
        yield array_merge([$h], $c);
    foreach(comb($m, $t) as $c)
        yield $c;
}
$words = array();
for($m=0;$m<7;$m++)
array_push($words,$m);

$cur_comb=array();
foreach(range(1, 7) as $n)
{
    foreach(comb($n, $words) as $c)
	{
       // echo join(' ', $c), "\n";
	   //$cur_comb.=implode('|', $c)."\r\n";
	   array_push($cur_comb,$c);
	}
}
var_dump($cur_comb);exit();


foreach($cur_comb as $cur)
{

}




$contextsArr0=array();// сочетания контекстов
*/



$combArr = array_unique($combArr); 
var_dump($combArr);exit();

$out="";
foreach($combArr as $ccomb)
{
$out.=$ccomb."\r\n";
}
$out=md5($str)."\r\n".$out;
write_combo_file($_SERVER["DOCUMENT_ROOT"]."/lib/contexts_combin.txt",$out);
exit("1111");


$list=read_combo_file($_SERVER["DOCUMENT_ROOT"]."/lib/contexts_combin.txt");
$combArr = explode("\r\n", $list);   
var_dump($combArr);exit();

// по каждому сочетанию готовим суммы строк
$contextsArr0=array();// сочетания контекстов
$n=0;
foreach($combArr as $ccomb)
{
$sumArr=array();// сумматор значений ячеек данного сочетания $ccomb
$pArr = explode("|", $ccomb);
foreach($pArr as $cell)
{
$sumArr=array_merge($sumArr,$cell);
}
//var_dump($pArr);exit();
$sumArr=array_unique($sumArr);
sort($sumArr, SORT_NUMERIC);reset($sumArr);
//if($n==10) {var_dump($sumArr);exit("<hr> $col1 | $row1 || $col2 | $row2 || $curComb1 | $curComb2");}
array_push($contextsArr0,$sumArr);
$n++;
}
var_dump($contextsArr0);exit();


// убрать антагонистов, отрицательнеы контексты (которые должны госиться) и перевести сочетания контекстов в строки, оставить только уникальные
$contextsArr=array();// сочетания контекстов
foreach($contextsArr0 as $comb)
{
$str="";  
$minusArr=array();
$antArr=array();

foreach($comb as $a)// подготовка к удалению отрицательных
{
	if($a<0){
		array_push($minusArr,-$a);
	}
}
$antArr=array();// для проверки антагонистов
foreach($comb as $a)
{
	if($a<0){
		continue;
	}
// убрать отрицательнеы контексты (которые должны гаситься)
if(in_array($a,$minusArr))
{
continue;
}

// исключить антагонистов, проверка для каждого выбранного ID кроме уже прошедших проверку
if(1)
{
$isAntagonist=0;
foreach($antArr as $g)
{ 
if(in_array($a,$antFromId[$g]))
{  
$isAntagonist=1;  // var_dump($antArr); exit("<hr> $a");
}
}
if($isAntagonist)
continue;
}

	if(!empty($str))
			{
				$str.=";";
			}
			$str.=$a;
			array_push($antArr,$a);
}
array_push($contextsArr,$str);
}
$contextsArr=array_unique($contextsArr);// Число сочетаний - 35 :)

//var_dump($contextsArr);exit("<hr>Число сочетаний: ".count($contextsArr));






///////////////////////////////////////////
// расположить по возрастанию чиcла контекстов
uasort($contextsArr, "cmpare");
function cmpare($a, $b) 
{ 
    if (strlen($a) == strlen($b)) {
        return 0;
    }
    return (strlen($a) < strlen($b)) ? -1 : 1;
}

// var_dump($contextsArr);exit("<hr>Число сочетаний: ".count($contextsArr));


// сохранять строки комбо контекстов в раб.файле  combo_contexts_str.txt
$list_id="";
$list_name="";
foreach($contextsArr as $str)
{
$list_id.=$str."\r\n";

$s="";
	$p = explode(";", $str);
foreach($p as $a)
{
	if(empty($a) || $a<0)
		continue;
if(!empty($s))
	{
	$s.=", ";
	}

	$s.=$a." ".$baseContextArr[$a][0];
}
$list_name.=$s."\r\n";
}
write_combo_file($_SERVER["DOCUMENT_ROOT"]."/pages/combinations/combo_contexts_str.txt",$list_id);
write_combo_file($_SERVER["DOCUMENT_ROOT"]."/pages/combinations/combo_contexts_names.txt",$list_name);


///////////////////////////////////////////////////
function write_combo_file($file,$content)
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
///////////////////////////////////////////////////
function read_combo_file($file)
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
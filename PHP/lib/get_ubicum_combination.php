<?
/* выдать массив неповторяющихся сочетений от 0 до таксимального числа

include_once($_SERVER['DOCUMENT_ROOT'] . "/lib/get_ubicum_combination.php");

$cellComb=get_ubicum_combination(56);
var_dump($cellComb);exit();
*/
function get_ubicum_combination($max)
{

function permutations($arr,$n)
{
     $res = array();

     foreach ($arr as $w)
     {
           if ($n==1) $res[] = $w;
           else
           {
                 $perms = permutations($arr,$n-1);

                 foreach ($perms as $p)
                 {
                      $res[] = $w."|".$p;
                 } 
           }
     }

     return $res;
}

$words = array();
for($m=0;$m<$max;$m++)
array_push($words,"$m");

$pe = permutations($words,3); // var_dump($pe);exit();


$cellComb = array();
foreach($pe as $p)
{
	$arr=explode("|",$p);  //var_dump($arr);exit();
	sort($arr, SORT_NUMERIC);reset($arr);
	$arr=array_unique($arr);
if(count($arr))
	array_push($cellComb,$arr);
}
// убрать повторения
//$cellComb=array_unique($cellComb, SORT_REGULAR);// array_unique для многомерных массивов оставляет некоторые
$cellComb = array_map("unserialize", array_unique(array_map("serialize", $cellComb)));// убирает дубли идеально

return $cellComb;
}




/* не хватает памяти: Fatal error: Out of memory (allocated 316669952) (tried to allocate 201326600 bytes) 
set_time_limit(0);
ini_set('memory_limit', '4024M');  

$cellComb = array();
$cellStr= array();
for($m=0;$m<56;$m++)
array_push($cellStr,"$m");

function placing($chars, $from=0, $to = 0){
	global $cellStr;
    $cnt = count($chars);
    if(($from == 0) && ($to == 0)){
        $from = 1;
        $to = $cnt;
    }
    if($from == 0) $from = 1;
    if($to == 0) $to = $from;
    if($from < $to){
        $plac = [];
        for($num = $from; $num <= $to; $num++){
            $plac = array_merge($plac, placing($cellStr, $num));
        }
    }else{
        $plac = [""];   
        for($n = 0; $n < $from; $n++){
            $plac_old = $plac;
            $plac = [];
            foreach($plac_old as $item){
                $last = strlen($item)-1;
                for($m = $n; $m < $cnt; $m++){
                    if($chars[$m] > $item[$last]){
                        $plac[] = $item.$chars[$m];
                    }
                }
            }
        }
    }
    return $plac;
}

$cellComb = placing($cellStr);
var_dump($cellComb);exit();


/*
// получается только две строки
$tComb=array();
$nNumbers=5;

$nStr= array();
for($m=0;$m<$nNumbers;$m++)
array_push($nStr,$m);

array_push($tComb,$nStr);
for($n=0;$n<$nNumbers;$n++)
{
$arr=array();
for($m=0;$m<$nNumbers;$m++)
{
array_push($arr,array($n,$nStr[$m])); 

}
//echo "<hr><hr>";
array_push($tComb,$arr);
}
*/

/*
foreach($tComb as $comb)
{
	if(count($comb)<2)
		continue;
	//var_dump($comb);exit();
	$arr=$comb;
	for($m=0;$m<count($comb)-1;$m++)
	{ 
		var_dump($arr);echo "<hr>";
      unset($arr[$m]);
	  //$arr = array_shift($arr);
	  array_push($tComb,$arr);
	}
}
 var_dump($tComb);exit();
*/








/*
$cellComb = array();
$cellStr= array();
for($m=0;$m<10;$m++)
array_push($cellStr,$m);

function placing($chars, $from=0, $to = 0){
	global $cellStr;
    $cnt = count($chars);
    if(($from == 0) && ($to == 0)){
        $from = 1;
        $to = $cnt;
    }
    if($from == 0) $from = 1;
    if($to == 0) $to = $from;
    if($from < $to){
        $plac = [];
        for($num = $from; $num <= $to; $num++){
            $plac = array_merge($plac, placing($cellStr, $num));
        }
    }else{
        $plac = [""];   
        for($n = 0; $n < $from; $n++){
            $plac_old = $plac;
            $plac = [];
            foreach($plac_old as $item){
                $last = strlen($item)-1;
                for($m = $n; $m < $cnt; $m++){
                    if($chars[$m] > $item[$last]){
                        $plac[] = $item.$chars[$m];
                    }
                }
            }
        }
    }
    return $plac;
}

$cellComb = placing($cellStr);
var_dump($cellComb);exit();
*/



/* //как мой алгоритм без вычитаний средних
function permutation(array $arr)
{
    while($ele=array_shift($arr))
    {
        $x=$ele;
        echo $x."<br>";
        foreach($arr as $rest)
        {
            $x.=" $rest";
            echo $x."<br>";
        }
    }
}
permutation(array("1","2","3","4"));
exit();
*/


/*  мой последний
$tComb=array();
$nNumbers=5;
$tComb=comb($nNumbers);

function comb($max)
{
$out=array();
for($n=0;$n<$max;$n++)
{
array_push($out,array($n));
$a2=array();
array_push($a2,$n);
for($m=$n+1;$m<$max;$m++)
{
array_push($a2,$m);// добавляется по одному
array_push($out,$a2);
}
$a2=comb(count($a2));
$out=array_merge($out,$a2);
}
return $out;
}


 var_dump($tComb);exit();
*/




/*
// тестирование правильности алгоритма выборки рабочих сочетаний
$tComb=array();
$nNumbers=5;
for($n=0;$n<$nNumbers;$n++)
{
array_push($tComb,array($n));
$a2=array();
array_push($a2,$n);
for($m=$n+1;$m<$nNumbers;$m++)
{
array_push($a2,$m);// добавляется по одному
array_push($tComb,$a2);
} 
}
// из каждого сочетания убираем по 1 кроме крайних
foreach($tComb as $comb)
{
	if(count($comb)<3)
		continue;
	$arr=$comb;
	for($m=1;$m<count($comb)-1;$m++)
	{
      unset($arr[$m]);
	  array_push($tComb,$arr);
	}

}
 var_dump($tComb);exit();



// https://ru.stackoverflow.com/questions/482847/%D0%93%D0%B5%D0%BD%D0%B5%D1%80%D0%B0%D1%86%D0%B8%D1%8F-%D1%81%D0%BE%D1%87%D0%B5%D1%82%D0%B0%D0%BD%D0%B8%D0%B9-%D0%B1%D0%B5%D0%B7-%D0%BF%D0%BE%D0%B2%D1%82%D0%BE%D1%80%D0%B5%D0%BD%D0%B8%D0%B9 все рабочие сочетания (без перестановочных повторений) номеров ячеек таблицы "Активности Базовых стилей" 29316 сочетаний для проверки моего алгоритма (предыдущий по коду)
$cellComb = array();
$cellStr= array();
for($m=0;$m<5;$m++)
array_push($cellStr,$m);

function placing($chars, $from=0, $to = 0){
	global $cellStr;
    $cnt = count($chars);
    if(($from == 0) && ($to == 0)){
        $from = 1;
        $to = $cnt;
    }
    if($from == 0) $from = 1;
    if($to == 0) $to = $from;
    if($from < $to){
        $plac = [];
        for($num = $from; $num <= $to; $num++){
            $plac = array_merge($plac, placing($cellStr, $num));
        }
    }else{
        $plac = [""];   
        for($n = 0; $n < $from; $n++){
            $plac_old = $plac;
            $plac = [];
            foreach($plac_old as $item){
                $last = strlen($item)-1;
                for($m = $n; $m < $cnt; $m++){
                    if($chars[$m] > $item[$last]){
                        $plac[] = $item.$chars[$m];
                    }
                }
            }
        }
    }
    return $plac;
}

$cellComb = placing($cellStr);
var_dump($cellComb);exit();
*/
?>
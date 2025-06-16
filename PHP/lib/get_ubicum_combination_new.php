<?
/* выдать массив неповторяющихся сочетений от 0 до таксимального числа

http://go/lib/get_ubicum_combination.php?max=12



сочетания без повторений +yield +php

https://stackoverflow.com/questions/25610919/all-combinations-without-repetitions-with-specific-cardinality
*/
header('Content-Type: text/html; charset=UTF-8');
set_time_limit(0);

$max=$_GET['max'];
$step=1;
if(isset($_GET['step'])) $step=$_GET['step']; // exit("! $step");


$words = array();
for($m=0;$m<$max;$m++)
array_push($words,$m);


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

//$range=range(1, count($words)); 
//$count=count($words);  // exit("! $step");
echo "Шаг ".$step;

$cur_comb="";                        //  $step=1;
//$cur_comb[$step]=array();
    foreach(comb($step, $words) as $c)
	{
       // echo join(' ', $c), "<br>";
	   //array_push($cur_comb,$c);
	   $cur_comb.=implode('|', $c)."\r\n";
	}
	//exit($cur_comb);

if($step==1)
{   //exit("! /lib/".$max."_combin.txt");
add_file($_SERVER["DOCUMENT_ROOT"]."/lib/".$max."_combin.txt",$cur_comb,"wb+"); //exit("! $step");
}
else
add_file($_SERVER["DOCUMENT_ROOT"]."/lib/".$max."_combin.txt",$cur_comb,"ab+");

if($step<$max)
{
	$step++;
echo "<script>window.top.location.href =\"/lib/get_ubicum_combination.php?max=".$max."&step=".$step."\";</script>";
exit();
}
exit();



///////////////////////////////////////////////////
function add_file($file,$content,$param)
{
$hf=fopen($file,$param);
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
?>
<?
/* Пpоверка перед записью и запись рефлексов
Проверяется уникальность сочетания условий чтобы не дублировались.
/pages/reflexes_server.php

До $lastID - просто корректировать старые с имеющимся ID, а после - добавлять новые
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$removeNotAllowe=0;// 1 - При сохранении очистить таблицу от рефлексов с невозможными сочетаниями Базовых контекстов.
if($_POST['removeNotAllowe']==1)
$removeNotAllowe=1;
// exit("> $removeNotAllowe");
//exit("> ".count($_POST['id1']));
//var_dump($_POST['id1']);exit();



// считать имеющиеся рефлексы
$str = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/dnk_reflexes.txt");
$list = explode("\r\n", $str);

$chererArr=array(); // будут сравниваться эти суммарные строки условий
$rArr=array();// сбор данных для записи рефлексов
$n=0;
foreach ($list as $str) 
{
	if (empty($str)) {
		continue;
	}
	$p=explode("|",$str);
	$id=(int)$p[0];

	if(isset($_POST['id1'][$id]))// скорректировать
	{
getCorrectStr($n,$id,$str);
	}
	else
	{
		$rArr[$n][0]=$id;
		$rArr[$n][1]=$p[1];
		$rArr[$n][2]=$p[2];
		$rArr[$n][3]=$p[3];
		$rArr[$n][4]=$p[4];
	}
//echo $id.": ".$rArr[$n][0]."|".$rArr[$n][1]."|".$rArr[$n][2]."|".$rArr[$n][3]."|".$rArr[$n][4]."<br>";
$n++;
}
//var_dump($rArr);exit();
//exit();


// добавить новые рефлексы
if(isset($_POST['id1']))
{
$lastID=$_POST['lastID'];   // exit("> $lastID");
foreach($_POST['id1'] as $id)
{
	$id=trim($id);
	if($id<$lastID)
		continue;
	//exit("> $lastID");
	getCorrectStr($n,$id,$str);
$n++;
}
}
//var_dump($rArr);exit();




//extract($_POST, EXTR_SKIP);
function getCorrectStr($n,$id,$str)
{
global $chererArr,$rArr;

$rArr[$n][0]=$id;

$str1=trim($_POST['id2'][$id]);
$rArr[$n][1]=trim($str1);

$str2=trim($_POST['id3'][$id]);
$str2=sorting_str($str2); 
$rArr[$n][2]=$str2;

$str3=trim($_POST['id4'][$id]);
$str3=sorting_str($str3);
$rArr[$n][3]=$str3;
$sum=$str1."_".$str2."_".$str3;

check_cond_str($id,$sum);
$chererArr[$id]=$sum;

$str=trim($_POST['id5'][$id]);

check_action_id($id,$str);// проверка валидности id действий

$str=sorting_str($str);
$rArr[$n][4]=$str;

//echo $id.": ".$rArr[$n][0]."|".$rArr[$n][1]."|".$rArr[$n][2]."|".$rArr[$n][3]."|".$rArr[$n][4]."<br>";


}
//var_dump($chererArr);exit();
//var_dump($rArr);exit();
//////////////////////////////
function sorting_str($str)
{
$p=explode(",",$str);
sort($p, SORT_NUMERIC);reset($p);// implode(",", $p);
$s="";
for($i=0;$i<count($p);$i++)
{
$p[$i]=trim($p[$i]);
if(empty($p[$i]))
	continue;
if($i)
$s.=",";
$s.=$p[$i];
}
return $s;
}
////////////////////////////////
function check_cond_str($id,$sum)
{
global $chererArr;
foreach($chererArr as $id0 => $ss)
{
if($id==$id0)// саму себя не смотреть
	continue;
	if($sum==$ss)
	{
exit("Строка с ID=$id имеет такие же условия как более ранняя строка с ID=$id0.<br><span style='color:red;'>Условия разных рефлексов не должны совпадать.</span>");
	}
}
}
/////////////////////////////////////////////////
function check_action_id($id,$str)// проверка валидности id действий
{
$val=(int)$str; if(empty($str) || $val==0){
exit("Строка с ID=$id НЕ ИМЕЕТ ID действий, что недопустимо.<br><span style='color:red;'>Рефлекс всегда должен иметь действие.</span>");
}
// существуют ли такие ID действий?
$aList=explode(",",$str);
$progs = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/terminal_actons.txt");
$strArr = explode("\r\n", $progs);  //var_dump($strArr);exit();
$actIDarr=array();
foreach ($strArr as $str) {
if (empty($str) || $str[0] == '#')
	continue;
$par = explode("|", $str); //var_dump($par);exit();
array_push($actIDarr,$par[0]);
}
foreach ($aList as $a)
{
if(!in_array($a, $actIDarr))
	{
exit("В строке с ID = $id есть несуществующее ID действия: $a.");
	}
}
}
////////////////////////////////////////////////




/////////////////// запись
if($removeNotAllowe)
{
// реально возможные сочетания контекстов
$c_list = read_file($_SERVER["DOCUMENT_ROOT"] . "/pages/combinations/combo_contexts_str.txt");
$c_list=str_replace(";",",",$c_list);
$allowContextArr=explode("\r\n",$c_list); // var_dump($allowContextArr);exit();
}


$out=""; 
//var_dump($_POST);exit();
//sort($rArr, SORT_NUMERIC);reset($rArr);
uasort($rArr, "cmp");
function cmp($a, $b) 
{ 
    if ($a[0] == $b[0]) {
        return 0;
    }
    return ($a[0] < $b[0]) ? -1 : 1;
}
//var_dump($rArr);exit();
$n=0;
$back="";// чисто для контроля
foreach($rArr as $str)
{
$s1=$str[0];
$s2=$str[1];
$s3=$str[2];
$s4=$str[3];
$s5=$str[4];
if($removeNotAllowe)
{
if(!in_array($s3,$allowContextArr))
{
continue;
}
}

$out.=$s1."|".$s2."|".$s3."|".$s4."|".$s5."\r\n";
}

//  exit("$out");
write_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/dnk_reflexes.txt",$out);
write_file($_SERVER["DOCUMENT_ROOT"]."/pages/dnk_reflexes_seved.txt","1");

echo "!";

///////////////////////////////////////////////////
function write_file($file,$content)
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
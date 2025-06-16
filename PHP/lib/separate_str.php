<?
/*
разделить строку на корни, окончания и символы: 
prepare_str($str)

Нужно иметь:
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");
mb_http_input('UTF-8');
mb_http_output('UTF-8');
mb_internal_encoding("UTF-8");

Нужно включить:
include_once($_SERVER["DOCUMENT_ROOT"]."/lib/separate_str.php");
include_once($_SERVER["DOCUMENT_ROOT"]."/lib/indexer_stemmer_UTF8.php");
$stemmer = new Lingua_Stem_Ru();

*/

function prepare_str($text,$stemmer)
{
// выделить слова
$text=mb_strtolower($text);  // exit("! $text");

// не здесь $text=str_replace('-',' ',$text);

//$text=preg_replace('/[^а-яА-Я ]/','',$text);
$text=preg_replace('/\r\n/','|##|',$text); 
$text=preg_replace('/\s/','|#|',$text);

$text=preg_replace('/[^|а-яa-z0-9\!\?\@\#\$\%\^\&\*\(\)\+\=\-1234567890\[\]\{\}\<\>\.\,\/ ]/u','',$text);
//$text=preg_replace('/[^абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ\!\?\@\#\$\%\^\&\*\(\)\+\=\-1234567890\[\]\{\}\<\>\.\,\/ /','',$text);

$text=str_replace(' ','',$text);
// все виды пробелов - в один пробел
$text=str_replace('|#|',' ',$text);   //exit("!!!! $text");
$text=str_replace('|##|',"\r\n",$text);
//exit("!!!! $text");

// разбить на слова и символы
$wordArr=array();
$str="";
$save=0;
for($n=0;$n<strlen($text);$n++)
{
//if(ord($text[$n])==13)exit("!!!!!!!!!!");
//||ord($text[$n])==13
if($text[$n]==' ' ||$text[$n]=='.' ||$text[$n]==',' || $text[$n]=='!' ||$text[$n]=='?'  ||$text[$n]=='(' ||$text[$n]==')' ||$text[$n]=='[' ||$text[$n]==']' || ord($text[$n])==13)
{
if(ord($text[$n])==13)
	{
	array_push($wordArr,"|##|");
	$save=1;
continue;
	}
	
$str=trim($str);
if(!empty($str))// если это не одно слово
{
	array_push($wordArr,$str);
	$save=1;
}
if($text[$n]==' ' ||$text[$n]=='.' ||$text[$n]==',' || $text[$n]=='!' ||$text[$n]=='?' ||$text[$n]=='(' ||$text[$n]==')' ||$text[$n]=='[' ||$text[$n]==']' )
{
array_push($wordArr,$text[$n]);// записать знак тоже !!!В ТОМ ЧИСЛЕ ПРОБЕЛ!!!!
$save=1;
}
$str="";
$save=0;
continue;
}

$str.=$text[$n];
}
//exit("!!! $save | ".$str);
if(!$save)
array_push($wordArr,$str);

//var_dump($wordArr);exit();

// разбить на корни, окончания и выдать в такой последовательности
$sArr=array();

$n=0;
foreach($wordArr as $w)
{ 


//$w=trim($w);	
//	exit("!!!! |$w| ".strlen($w));
if(mb_strlen($w)==1)// отдельные символы
	{ 
array_push($sArr,$w);
continue;
	}
////////////////////


// слова с тире: все-таки так-то все-все
if(strpos($w,"-")!==false)// разделить на отдельные слова и отдельно '-'
{ //exit("$w");
$list=explode("-",$w); //var_dump($list);exit();
$m=0;
foreach($list as $s)
{ 
if($m)
array_push($sArr,"-");
$res=separate($s,$sArr,$stemmer);
if(!$res)
	continue;
$n++;
$m++;
}
continue;
}

// нормальные слова
$res=separate($w,$sArr,$stemmer);
if(!$res)
	continue;

$n++;
}

//var_dump($sArr);exit();

// разбить на фразы,превратить в строку
$out="";
$n=0;
foreach($sArr as $s)
{
//  ||$s==','
if($s=='.' || $s=='...' || $s=='!' || $s=='?' || $s=='|##|' )// конец фразы
{
if($s=='|##|')
$out.='|#'; 
else
$out.="|".$s.'|#'; 
continue;
}

if($n)
$out.="|";

// отделить приставки
//$s=predlog_selecting($s) дело неподьемное и вредное, просто не нужно отделять значащую часть слова.


$out.=$s;

$n++;

}
$out=str_replace("|#|#","|#",$out);// убрать лишний конец фразы (по точке и т.п.)
$out=str_replace("#| ","#",$out);// убрать пробел в начале фразы

return $out;
}
//////////////////////// выделение окончаний и корней
function separate($w,&$sArr,$stemmer)
{
$core =$stemmer->stem_word($w);  
//$core=preg_replace('/[^а-яА-Я]/','',$core);
if(empty($core))
	return 0;

array_push($sArr,$core);   //exit("! $w | $core");

if(mb_strlen($w,'UTF-8') > mb_strlen($core,'UTF-8'))
{ 
$len1=mb_strlen($w,'UTF-8');
$len2=mb_strlen($core,'UTF-8');
$end=mb_substr($w,$len2,$len1-$len2,'UTF-8'); 
//exit("$w | $core | $end <br> $len1 | $len2 <br>".mb_strlen($core,'UTF-8'));
array_push($sArr,$end);
}

return 1;
}
////////////////////////////// отделение приставок и корней
function predlog_selecting($s)
{
// чтобы выделить приставку нужно сравнивать с другими словаи имеющими такой же корень (общую часть всех слов) Гиблое дело....

}
///////////////////////////////
?>
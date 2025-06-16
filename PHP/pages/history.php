<?
/* История общения запись

go/pages/history.php

*/
$curHistoryFile="";// имя файла для записи истории

////////////////////////////////////////////////////
if($init_history)
{
$dir=$_SERVER["DOCUMENT_ROOT"]."/history/"; 
$dayliName="".date("Y-m-d").".htm";
$existed=false;
if($dh = opendir($dir)) 
{  
while(false !== ($file = readdir($dh))) 
{		
if($file=="." || $file=="..")
	continue;

if(!is_dir($dir.$file))
{
//$t=filemtime($dir.$file); //exit("! $t");
if($file == $dayliName)
	{
$existed=true;
break;
	}else{
	continue;
	}

}
}
closedir($dh);
}
if(!$existed)
	{
// если дата файла != сегодняшней, открыть новый
$dDayly=mktime(0, 0, 0, date("m"), date("d"), date("Y") );  // exit("! $dDayly");
$dFile=mktime(0, 0, 0, date("m",$time), date("d",$time), date("Y",$time) );
//exit("$dDayly != $dFile");
createNewHistFile($dayliName);  
//
	}
$curHistoryFile=$dayliName;
// exit("! $curHistoryFile");
}
/////////////////////////////////////
function createNewHistFile($dayliName)
{
$beastTime="";
$file=$_SERVER["DOCUMENT_ROOT"]."/memory_reflex/life_time.txt";
$hf=fopen($file,"rb");
if($hf)
{
$beastTime=fread($hf,filesize($file));
fclose($hf);
}//if($hf)
if(!empty($beastTime))
{
$yeas = (int)($beastTime / (3600 * 24 * 365));
$eStr="лет";
if($yeas==1){$eStr="год";}
if($yeas>1 && $yeas<5){$eStr="года";}

$month = (int)( ($beastTime - $yeas*3600*24*365)/ (3600*24*30)  );
$mStr="месяцев";
if($month==1){$mStr="месяц";}
if($month>1 && $month<5){$mStr="месяца";}

$days = (int)(($beastTime - $yeas*3600*24*365  - $month*3600*24*30)/ (3600*24)); 
$dStr="дней";
if($days==1){$dStr="день";}
if($days>1 && $days<5){$dStr="дня";}
$beastTime="Возраст: ".$yeas." ".$eStr.", ".$month." ".$mStr.", ".$days." ".$dStr."";
}


$content="<h2>История за ".date("d-m-Y")." ".$beastTime."</h2>\r\n";
$file=$_SERVER["DOCUMENT_ROOT"]."/history/".$dayliName;
$hf=fopen($file,"wb+");
if($hf)
{
fwrite($hf,$content,strlen($content));
fclose($hf);
chmod($file, 0666);
}
}
//////////////////////////////////////////////////////////////////
//$_GET['BeastOld']="Возраст лет: 1, мес: 2, дней: 3";
//$_GET['histoty_file']="";
if(isset($_GET['BeastOld']))//нужно прописать возраст в заголовке истории
{
$file=$_SERVER["DOCUMENT_ROOT"]."/history/".$_GET['histoty_file'];
//addToHistFile($file,"<hr>".$_GET['BeastOld'],"<hr>");

}

//////////////////////////////////////////////////////////////////






/////////////////////////////////////////////////////////////////
if(isset($_POST['newInfoHist'])&&!empty($_POST['newInfoHist']))
{
$curTime="<b>".date("H:i:s")."</b> ";
$file=$_POST['histoty_file'];  // exit("!! $file");
$str=$_POST['newInfoHist'];  //exit("$str");
$str=str_replace("[x1]","&",$str);
$context=""; if(isset($_POST['context']))$context=$_POST['context'];

switch($_POST['type'])
{
case 1:// ответ бота
$str=$curTime.$context."<br><div style='padding:10px;border:solid 1px #8A3CA4;border-radius: 7px;max-width:800px;margin-top:10px;'>".$str."</div>\r\n\r\n";
break;

case 2:// текст оператора
$pArr=explode("||",$str);
$ton="нормальный"; if($pArr[0]==4)$ton="Повышенный";if($pArr[0]==3)$ton="Вялый";
$mood="Нормальное";
switch($pArr[1])
{
	case 20: $mood="Хорошее"; break;
	case 21: $mood="Плохое"; break;
	case 22: $mood="Игровое"; break;
	case 23: $mood="Учитель"; break;
	case 24: $mood="Агрессивное"; break;
	case 25: $mood="Защитное"; break;
	case 26: $mood="Протест"; break;
}
$actions="";
if(!empty($pArr))
{
$actions=getActions($str);
$actions="<br><b>Действия оператора:</b><br><b><span style='background-color:#ffffff;padding:4px;'>".$actions."</span></b>";
}

$str=$curTime.$context."<br><div style='padding:10px;border:solid 1px #000000;border-radius: 7px;max-width:800px;margin-top:10px;background-color:#EEEBEB;'><b>Текст оператора:</b><br><b><span style='background-color:#ffffff;padding:4px;'>".$pArr[2]."</span>".$actions."</b><br>Тон: ".$ton." Настороение: ".$mood."</div>\r\n\r\n";
break;

case 3:// действия оператора
$actions=getActions($str);
$str=$curTime.$context."<br><div style='padding:10px;border:solid 1px #000000;border-radius: 7px;max-width:800px;margin-top:10px;background-color:#EEEBEB;'><b>Действия оператора:</b><br><b><span style='background-color:#ffffff;padding:4px;'>".$actions."</span></b></div>\r\n\r\n";
break;

case 4:// Есть связь
$str=$curTime."<br><div style='padding:10px;border:solid 1px green;border-radius: 7px;max-width:800px;margin-top:10px;background-color:#EEFFEB;color:green;'>Beast включился</div>\r\n\r\n";
break;

case 5:// Нет связи
$str=$curTime."<br><div style='padding:10px;border:solid 1px red;border-radius: 7px;max-width:800px;margin-top:10px;background-color:#FFEEEB;color:red;'>Выключение Beast</div>\r\n\r\n";
break;

}
addToHistFile($file,$str);
exit("ok");
}
///////////////////////////////
function getActions($str)
{
$actions="";
$pArr=explode("|",$str);
$n=0;
foreach($pArr as $a)
{
	if(empty($a)) continue;
if($n>0) $actions.=" ";
switch($a)
{
	case 1: $actions.="Непонятно"; break;
	case 2: $actions.="Понятно"; break;
	case 3: $actions.="Наказать"; break;
	case 4: $actions.="Поощрить"; break;
	case 5: $actions.="Накормить"; break;
	case 6: $actions.="Успокоить"; break;
	case 7: $actions.="Поиграть"; break;
	case 8: $actions.="Поучить"; break;
	case 9: $actions.="Игнорировать"; break;
	case 10: $actions.="Сделать больно"; break;
	case 11: $actions.="Сделать приятно"; break;
	case 12: $actions.="Заплакать"; break;
	case 13: $actions.="Засмеяться"; break;
	case 14: $actions.="Обрадоваться"; break;
	case 15: $actions.="Испугаться"; break;
	case 16: $actions.="Простить"; break;
	case 17: $actions.="Вылечить"; break;
}
$n++;
}
return $actions;
}
////////////////////////////////////////////////////////////////



///////////////////////////////////////////////////////////////////
function addToHistFile($file,$str)
{
$file=$_SERVER["DOCUMENT_ROOT"]."/history/".$file;
$hf=fopen($file,"ab+");
if($hf)
{
fwrite($hf,$str,strlen($str));
fclose($hf);
chmod($file, 0666);
}
}
?>
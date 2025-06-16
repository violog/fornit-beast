<?
/*  Выдать контрол для выбора действий из списка ДЛЯ РЕФЛЕКСОВ
Для заполнения полей ввода.
/lib/get_multiselectiong.php?kind=1&id=1

В зависимости от kind предлагаются соотвествующие жанному полю ввода значения.

*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');

$id=$_GET['id'];
$kind=$_GET['kind'];

include_once($_SERVER['DOCUMENT_ROOT']."/common/common.php");


// считать файл 
$progs=read_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/dnk_reflexes.txt");
$strArr=explode("\r\n",$progs); // var_dump($strArr);exit();
$n=0;
$lastID=1;
$lev1="";
$lev2="";
$lev3="";
$lev4="";
foreach($strArr as $str)
{      
$par=explode("|",$str);	 // exit("$id | ".$par[0]);
	
if($id==$par[0])
{
$lev1=$par[1];
$lev2=$par[2];
$lev3=$par[3];
$lev4=$par[4];  // var_dump($par);exit();
}
}
/////////////////////////////////////////////



/////////////////////////////////////////////////////
//Первый уровень - ID базовых состояний 
if($kind==1)
{


}
/////////////////////////////////////////////////////


/////////////////////////////////////////////////////
//Второй уровень - ID актуальных Базовых Контекстов: 
if($kind==2)
{
$vArr=array();
$idArr=explode(",",$lev2);    //  var_dump($idArr);exit();
foreach($idArr as $s)
{
	$s=trim($s);
	if(empty($s))
		continue;
array_push($vArr,$s);
}

$out="<select id='select_combo' name='actions_combo[]' multiple='multiple' size=8 style='width:300px;padding:4px;'>";

$mArr=array(1=>"1 Пищевой",2=>"2 Поиск",3=>"3 Игра",4=>"4 Гон",5=>"5 Защита",6=>"6 Лень",7=>"7 Ступор",8=>"8 Страх",9=>"9 Агрессия",10=>"10 Злость",11=>"11 Доброта",12=>"12 Сон",);
foreach($mArr as $id => $name)
{
$out.="<option id='".$id."' value='".$id."'"; if(in_array($id,$vArr))$out.="selected";$out.=">".$name."</option>";
}

$out.="</select>";

exit($out);
}
/////////////////////////////////////////////////////


/////////////////////////////////////////////////////
// 3-й уровень - Третий уровень - ID пусковых стимулов: 
if($kind==3)
{
$vArr=array();
$idArr=explode(",",$lev3);    //  var_dump($idArr);exit();
foreach($idArr as $s)
{
	$s=trim($s);
	if(empty($s))
		continue;
array_push($vArr,$s);
}

$out="<select id='select_combo' name='actions_combo[]' multiple='multiple' size=8 style='width:300px;padding:4px;'>";

$mArr=array(1=>"1 Непонятно",2=>"2 Понятно",3=>"3 Наказать",4=>"4 Поощрить",5=>"5 Накормить",6=>"6 Успокоить",7=>"7 Предложить поиграть",8=>"8 Предложить поучить",9=>"9 Игнорировать",10=>"10 Сделать больно",11=>"11 Сделать приятно",12=>"12 Заплакать",13=>"13 Засмеяться",14=>"14 Обрадоваться",15=>"15 Испугаться",16=>"16 Простить",17=>"17 Вылечить",);
foreach($mArr as $id => $name)
{
$out.="<option id='".$id."' value='".$id."'"; if(in_array($id,$vArr))$out.="selected";$out.=">".$name."</option>";
}

$out.="</select>";

exit($out);
}
/////////////////////////////////////////////////////


// 4-й уровень - ID Действий рефлекса: 
if($kind==4)
{
$vArr=array();
$idArr=explode(",",$lev4);    //  var_dump($idArr);exit();
foreach($idArr as $s)
{
	$s=trim($s);
	if(empty($s))
		continue;
array_push($vArr,$s);
}

$out="<select id='select_combo' name='actions_combo[]' multiple='multiple' size=8 style='width:300px;padding:4px;'>";

$progs=read_file($_SERVER["DOCUMENT_ROOT"]."/memory_reflex/terminal_actons.txt");
$strArr=explode("\r\n",$progs);
foreach($strArr as $str)
{
if(empty($str) || $str[0]=='#')
	continue;
									//exit($str);
$p=explode("|",$str);
$id=$p[0];

$out.="<option id='".$id."' value='".$id."'"; if(in_array($id,$vArr))$out.="selected";$out.=">".$id." ".$p[1]."</option>";
	
}

$out.="</select>";

exit($out);
}

exit("sdfs fssdfgs f");
?>
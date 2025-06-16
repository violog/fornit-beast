<?
/* сформировать таблицу для http://go/pages/condition_reflexes_basic_phrases.php 

/pages/condition_reflexes_basic_phrases_table.php?bsID=1&id_list=1;8
*/

header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
setlocale(LC_ALL, "ru_RU.UTF-8");

$bsID=$_GET['bsID'];
$id_list=$_GET['id_list'];

$kind=$_GET['kind'];


////// Собрать данные по существующим рефлексам
$progs = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/dnk_reflexes.txt");
$strArr = explode("\r\n", $progs);  //var_dump($strArr);exit();
$reflexArr=array();
foreach ($strArr as $str) {
	if (empty($str))
		continue;
$par = explode("|", $str);  // var_dump($par);exit();
$par[2]=str_replace(",",";",$par[2]);  

if($par[1]!=$bsID || $par[2]!=$id_list)
	continue;
//array_push($reflexArr,$par);
$reflexArr[$par[0]]=array($par[1],$par[2],$par[3],$par[4]);
}
// var_dump($reflexArr);exit();

////////////////////////////////////////////////////////////////////

///////////////////// список возможных действий
$rActionsArr=array();
$progs = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/terminal_actons.txt");
	$strArr = explode("\r\n", $progs);
	foreach ($strArr as $str) {
if (empty($str) || $str[0] == '#')
			continue;
$p = explode("|", $str);
$rActionsArr[$p[0]]=$p[1];
}
/////////////////////////////////////

// Пусковые стимулы
$progs = read_file($_SERVER["DOCUMENT_ROOT"] . "/pages/combinations/list_triggers.txt");
$progs=substr($progs,strpos($progs,"\r\n")+2); // exit("$progs");
$aArr = explode("\r\n", $progs);
$triggerArr=array();
//$triggerArr["_"]="";
foreach ($aArr as $str) {
	if(empty($str))
		continue;
$p = explode("|", $str);  
$triggerArr[$p[0]]=$p[1];
}
// var_dump($triggerArr);exit();

///////////////////////////////////////////


$out="";


///////////////////////////////////////
// имеющиеся фразы
$id_list = str_replace(";",",",$id_list);
$file=$_SERVER["DOCUMENT_ROOT"]."/lib/condition_reflexes_basic_phrases/".$bsID."_".str_replace(",","_",$id_list).".txt";
//exit("$file");
$progs = read_file($file);
$strArr = explode("\r\n", $progs);
$phraseArr=array();
foreach ($strArr as $str) {
		if (empty($str) || $str[0] == '#')
			continue;
		$p = explode("|", $str); 
// т.к. добавляли метку для придания файлу кода UTF, нужно ее очистить
		$k=(int)preg_replace('/[^0-9]/','',$p[0]);            //  exit($p[0]." | ".$k);
		$phraseArr[$k]=$p[5];
}
//  var_dump($phraseArr);exit();



// Заполнить из общего шаблона
if($kind)
{
$progs = read_file($_SERVER["DOCUMENT_ROOT"]."/lib/condition_reflexes_basic_phrases_common.txt");
$strArr = explode("\r\n", $progs);
$commonArr=array();
foreach ($strArr as $str) {
		if (empty($str) || $str[0] == '#')
			continue;
		$p = explode("|", $str); 
		$commonArr[$p[0]]=$p[1];
}
//  var_dump($commonArr);exit();
}
//////////////////////////////////////////////


// проверка неповторяемости слов, иначе у.рефлекс будет неопределнным
$badArr="";
$repeatedArr=array();
$wArr=array();
foreach($phraseArr as $str)
{
if(in_array($str,$wArr))
{
$badArr.=$str."; ";
array_push($repeatedArr,$str);
}
array_push($wArr,$str);
}
if(!empty($badArr))
{
	
$out="<b><span style='color:red'>Есть повторяющиеся фразы: ".$badArr."</span></b>. Исправьте, сохраните и нажмите на кнопку &quot;Создать таблицу для заполнения фразами&quot;.";
}
///////////////////////////////////////////////////////////////////////



/////////////////////////////////////////////////////////
////////////////////////////////////// вывод таблицы
$out.="<table class='main_table' cellpadding=0 cellspacing=0 border=1 width='100%'>
		<tr>
		<th width=70 class='table_header'>рефлекс</th>
			<th width=300  class='table_header'>Пусковые стимулы рефлекса</th>
			<th  class='table_header'>Действия рефлекса</th>
			<th  class='table_header'>Фраза-синоним</th>
		</tr>";


$nid=0;   
foreach ($reflexArr as $id => $resArr)
{
//var_dump($resArr);exit();
// рефлексы с пустыми пусковыми не использовать
if(empty($resArr[2]))
	continue;

$out.="<tr class='r_table highlighting' style='background-color:#eeeeee;' onClick='set_sel(this," . $id . ")'>";
$out.="<td >" . $id . "</td>";

// пусковые стимулы
$out.="<td ><input type='hidden' value='".$resArr[2]."'>".$triggerArr[$resArr[2]]."</td>";

// действия рефлекса
$out.="<td ><input type='hidden' value='" . $resArr[3] . "'>".get_actions($resArr[3])."</td>";

// фраза-синоним
$phrase=get_prase_exists($id); // exit("$phrase");
if($kind && empty($phrase))
{
$phrase=$commonArr[$resArr[2]];  // exit("$id | $phrase");
}
$bg="";
if(in_array($phrase,$repeatedArr))
{
$bg="style='background-color:#FFFFAA;'";
}
$out.="<td class='table_cell'><input id='insert_".$nid."' class='table_input' type='text' value='".$phrase."'  ".$bg."><img src='/img/down17.png' class='select_control' onClick='show_word_list(".$nid.")'  title='Выбор слов'></td>";

$out.="</tr>";
$nid++;
}
$out.="</table>";

$out.="<br><input type='button' value='Сохранить фразы' onClick='reflex_saver()'>";

echo "!".$out;
////////////////////////////////////////////////////////////////////



///////////////////////////////////////////////
function get_actions($trArr)
{
	global $rActionsArr;
$acts="";
$aArr=explode(",",$trArr); 
foreach($aArr as $a)
{
	if(empty($a))
		continue;
	if(!empty($acts))
		$acts.=", ";
$acts.=$a." ".$rActionsArr[$a]."";
}
return $acts;
}






///////////////////////////////// // есть ли такой рефлекс?
function get_prase_exists($id)
{
global $phraseArr; //exit("$id");
//echo "$bsID | $id_list | $actions<br>";

if(isset($phraseArr[$id]))
	return $phraseArr[$id];// вернуть фразу

return "";
}

///////////////////////////////////////////////////
function read_file($file)
{
if(!file_exists($file))
	return "";
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
///////////////////////////////////////////////////
function write_trigger_file($file,$content)
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
?>
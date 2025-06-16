<?
/* сформировать таблицу для http://go/pages/mirror_reflexes_basic_phrases.php 

/pages/mirrors_automatizm_table.php?bsID=1&id_list=1;8
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


// Пусковые стимулы
// все используемые фразы в шаблонах условных рефлексов
include_once($_SERVER['DOCUMENT_ROOT'] . "/pages/mirrors_automatizm_get_all_phrases.php");
// var_dump($triggerPhraseArr);exit();

////////////////////////////////////////////////////////////////////

// сохраненный общий шаблон фразы-ответы   trigg|answers|ton,mood|actions
$file=$_SERVER["DOCUMENT_ROOT"]."/lib/mirror_basic_phrases_common.txt";
//exit("$file");
$progs = read_file($file); //var_dump($commonArr);exit();
$strArr = explode("\r\n", $progs);
$commonArr=array();
	foreach ($strArr as $str) {
		if (empty($str))
			continue;
		$p = explode("|", $str);
		$commonArr[$p[0]][0]=$p[1];
		$commonArr[$p[0]][1]=$p[2];
		$commonArr[$p[0]][2]=$p[3];
	}
//  var_dump($commonArr);exit();
/////////////////////////////////////



///////////////////////////////////////
// имеющиеся ответы   baseID|contextsID|trigg|answers|ton,mood|actions
$id_list = str_replace(";",",",$id_list);
$file=$_SERVER["DOCUMENT_ROOT"]."/lib/mirror_reflexes_basic_phrases/".$bsID."_".str_replace(",","_",$id_list).".txt";
//exit("$file");
$progs = read_file($file);
 $progs = substr($progs, 3);
$strArr = explode("\r\n", $progs);
$phraseArr=array();    
foreach ($strArr as $str) {
	// т.к. добавляли метку для придания файлу кода UTF, нужно ее очистить
		if (empty($str) || $str[0] == '#')
			continue;
		$p = explode("|", $str); 

		$k=trim($p[0]);           //  exit($p[0]." | ".$k);
		$phraseArr[$k][0]=$p[3]; // exit("$k | ".$phraseArr[$k][0]);
		$phraseArr[$k][1]=$p[4];
		$phraseArr[$k][2]=$p[5];
}

//  var_dump($phraseArr);exit("<hr>$first");
/////////////////////////////////////////



///////////////////////////////////////////////////////////////////////

$out="";

/////////////////////////////////////////////////////////
////////////////////////////////////// вывод таблицы
$out.="<table class='main_table' cellpadding=0 cellspacing=0 border=1 width='100%'>
		<tr>
		<th width=300 class='table_header'>Пусковая фраза</th>
			<th width=300  class='table_header'>Ответная фраза</th>
			<th  class='table_header'><nobr>Тон и настроение</nobr></th>
			<th  class='table_header'>Ответные действия</th>
		</tr>";


$nid=0;   
foreach ($triggerPhraseArr as $id => $tArr)
{
// exit("$tArr");
// фразы-ответы
$answ="";
$tm="0,0"; 
$actn="";    // exit("$tArr | ".$phraseArr[$tArr][0]);
if(isset($phraseArr[$tArr]))
{    
$answ=$phraseArr[$tArr][0];  // var_dump($phraseArr[$tArr]);exit();
$tm=$phraseArr[$tArr][1];
$actn=$phraseArr[$tArr][2];
}
else
{ 
if(isset($commonArr[$tArr]))
{
$answ=$commonArr[$tArr][0];
$tm=$commonArr[$tArr][1];
$actn=$commonArr[$tArr][2];
}
}


$out.="<tr class='r_table highlighting' style='background-color:#eeeeee;' onClick='set_sel(this," . $id . ")'>";

// пусковые стимулы
$out.="<td  class='table_cell' style='background-color:#eeeeee;'><input type='hidden'  name='trigg[]' value='".$tArr."'><nobr>".$tArr."</nobr></td>";

//Ответная фраза
$out.="<td  class='table_cell'><input  name='answ[]' class='table_input' type='text' value='".$answ."' ></td>";
//Тон и настроение
$out.="<td  class='table_cell'><input id='insert_".$nid."' name='ton_mood[]' class='table_input' type='text' value='".$tm."' ><img src='/img/down17.png' class='select_control' onClick='show_ton_mood(".$nid.")' title='Выбор Тона и Настроения'></td>";
//Ответные действия
$out.="<td  class='table_cell'><input id='insert2_".$nid."' name='actn[]' class='table_input' type='text' value='".$actn."' ><img src='/img/down17.png' class='select_control' onClick='show_actions_list(".$nid.")' title='Выбор действий'></td>";



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
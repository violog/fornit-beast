<?
/* набивальщик базы простейщих фраз для заливки базы условных рефлексов

http://go/pages/condition_reflexes_basic_phrases.php
*/
$page_id = -1;
$title = "Создание базы простейщих фраз для заливки базы условных рефлексов";
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");
//include_once($_SERVER['DOCUMENT_ROOT']."/pult_js.php");
//////////////////////////////////////////////////////////////

$bsID=0;
if(isset($_GET['bsID']))
$bsID=$_GET['bsID'];            //  exit("> $bsID");

$id_list="";
$get_list="";
if(isset($_GET['id_list']))
{
$id_list=$_GET['id_list'];

$get_list=explode(";",$id_list); 
}


?>
<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
var is_table_shoved=0;// 1 таблица показывается
function get_table(kind)
{
wait_show();
var link="/pages/condition_reflexes_basic_phrases_table.php?"+cur_condition_choose+"&kind="+kind;
//alert(link);
var AJAX = new ajax_support(link, sent_table_info);
AJAX.send_reqest();
function sent_table_info(res)
{
//show_dlg_alert(res,0);
if(res[0]!='!')
{ //alert(res);
show_dlg_alert(res,0);
return;
}
document.getElementById('table_id').innerHTML=res.substr(1);
document.getElementById('insert_from_common_id').style.display="block";

is_table_shoved=1;
wait_end();
}
}


</script>
<?
//////////////////////////////////////////////////////////////
echo "<div id='hr_table_id' style='position:relative;display:none;'>
<hr>
<div style='position:absolute;top:-10px;left:50%;transform: translate(-50%, 0);background-color:#ffffff;padding-left:10px;padding-right:10px;'><b>Задать условия для таблицы ввода фраз-синонимов</b>
</div>";

echo "<div style='position:absolute;z-index:10;top:20px;right:20px;'><a href='/pages/condition_reflexes_basic_phrases.htm'>Страница пояснений</a></div>";

echo "<div style='position:absolute;top:60px;right:10px;border:solid 1px #8A3CA4;border-radius: 7px;padding:10px;box-shadow: 8px 8px 8px 0px rgba(122,122,122,0.3);background-color:#efefef;max-width:65%;font-size:14px;'>
<b>Поясненния:</b><br>
Редактор позволяет создавать фразы - синонимы безусловных рефлексов.<br>
Остается дополнить правую колонку Фразой-синонимом.<br>

<b>Использование:</b><br>
В верхнем выпадающем списке выбрать Базовый контекст и под ним – выбрать одно из сочетаний Базовых контекстов.<br>
После нажатия кнопки “Создать таблицу для заполнения фразами” будет сформирована таблица, в правой колонке которой нужно ввести фразу-синоним рефлекса. После заполнения таблицы следует нажать под ней “Сохранить фразы” или просто нажать Ctrl+S. 

</div>";

echo "<div style='position:absolute;top:240px;;left:450px;font-size:16px;' ><b>Фразы должны быть уникальны</b> для таблицы, иначе условный рефлекс окажется неопределенным!</div>";


echo "<div style='position:absolute;top:270px;right:10px;font-size:18px;cursor:pointer;color:#885CFF' onClick='location.href=`/pages/condition_reflexes_basic_phrases_common.php`'>Общий шаблон пусковых символов</div>";

// onChange='choode_base_cond(this)' - нет определенной зависимости...
echo "<b>Базовое состояние:</b><br>
<select id='base_id' onChange='refresh_context_combo(this)'> 
<option value='1' "; if($bsID==1)echo "selected"; echo ">Плохо</option>
<option value='2' "; if($bsID==2)echo "selected"; echo ">Норма</option>
<option value='3' "; if($bsID==3)echo "selected"; echo ">Хорошо</option>
</select><span title='Общее Базовое состояние формируется из отдельных состояний Базовых параметров гомеостаза и при этом никак не коррелирует с диапазонами состояний параметров гомеостаза.'> - Общее Базовое состояние</span><br>
";

//exit("> $bsID");






// Базовые контексты $baseContextArr
include_once($_SERVER['DOCUMENT_ROOT'] . "/lib/base_context_list.php");

/* Все возможные сочетания активных контекстов выбираются из таблицы "Активности Базовых стилей" (минуса игнорируются и идет проверка на антагонистов).
*/
$contextsArr=array();// ID выбранных контекстов без антагонистов
echo "<b>Выбрать сочетания контекстов:</b><br> 
<div id='context_variations_id'></div>";

?>
<script>
var current_base_condition=1;// при обновлении страницы - Плохо
function refresh_context_combo(combo)
{
current_base_condition=combo.options[combo.selectedIndex].value; //alert(current_base_condition);
get_context_variations(1);
}

function get_context_variations(bc)
{ 
	wait_begin();
	// base_condition="+bc+"& 
//	alert("/pages/reflexes_maker_b_contexts.php?get_list=<?=$id_list?>");
var AJAX = new ajax_support("/pages/reflexes_maker_b_contexts.php?current_base_condition="+current_base_condition+"&get_list=<?=$id_list?>", send_context_variations);
		AJAX.send_reqest();

function send_context_variations(res) 
{ // alert(res);
			wait_end();
if(res[0]!='!')
{ //alert(res);
show_dlg_alert(res,0);
return;
}
document.getElementById('context_variations_id').innerHTML=res.substr(1); 
//alert(res.substr(1));

//show_dlg_alert(res.substr(1),0);
document.getElementById('button_table_id').style.display="";
document.getElementById('hr_table_id').style.display="block";

}
}
get_context_variations(1);// сразу показать 
/*
function choode_base_cond(combo)
{
var bc=combo.options[combo.selectedIndex].value; //alert(bc);
get_context_variations(bc);
}
*/
</script>
<?
///////////////////////////////////////////////////////////////////////



echo "<br><input id='button_table_id' type='button' value='Создать таблицу для заполнения фразами' onClick='choose_0()' style='display:none;'>";
//////////////////////////////////////////////////////////////

echo "<div style='position:relative;'>
<hr>
<div style='position:absolute;top:-10px;left:50%;transform: translate(-50%, 0);background-color:#ffffff;padding-left:10px;padding-right:10px;'><b>Таблица для заполнения фразами-синонимами</b>
</div>";

echo "<div id='insert_from_common_id' style='position:absolute;top:-10px;right:10px;background-color:#efefef;font-size:18px;border:solid 1px #8A3CA4;border-radius: 7px;padding-left:10px;padding-right:10px;cursor:pointer;display:none;' onClick='get_table(1)'>Заполнить из общего шаблона</div>";
//var_dump($contextsArr);exit();
//////////////////////////////////////////////////////////////

echo "<div id='conditions_block_id' style='position:relative;display:none'>";
echo "<div style='position:absolute;top:10px;right:0px;'><span onClick='prases_saver()' style='color:#AE55FF;cursor:pointer;font-size:18px;'>Сохранение фраз</span> - по <b>Ctrl+S</b></div>";
echo "<b>Выбранные условия:</b><br>";

echo "Базовое состояние: <b><span id='base_cond_id'></span></b>";
echo "</b>&nbsp;&nbsp;&nbsp;&nbsp;Сочетания контекстов: <b><span id='base_context_name'></span></b>";

echo "</div>";

//////////////////////////////////////// ТАБЛИЦА
echo "<div id='table_id'></div>";




//////////////////////////////////////////////////////////////
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/alert2_dlg.php");
?>
<script Language="JavaScript" src="/ajax/ajax_post.js"></script>
<script>
var cur_bcond_choose="";
var cur_bcontex_choose="";
var cur_condition_choose="";
function choose_0()
{
var bsID=document.getElementById('base_id').selectedIndex +1;
switch(bsID)
{
case 1: cur_bcond_choose="Плохо"; break;
case 2: cur_bcond_choose="Норма"; break;
case 3: cur_bcond_choose="Хорошо"; break;
}
//alert(bsID);

var combo=document.getElementById('base_context_id'); 
//alert(combo.selectedIndex);
if(combo.selectedIndex==-1)
	{
	show_dlg_alert("Нужно выбрать сочетание Базовых контекстов в списке.",0);
	return;
	}
var id_list=combo.options[combo.selectedIndex].value; // 1;3
cur_bcontex_choose=combo.options[combo.selectedIndex].text;


cur_condition_choose='bsID='+bsID+'&id_list='+id_list; 
// alert(cur_condition_choose);

document.getElementById('conditions_block_id').style.display="block";
document.getElementById('base_cond_id').innerHTML=cur_bcond_choose;
document.getElementById('base_context_name').innerHTML=cur_bcontex_choose;

get_table(0);


//location.href='/pages/reflexes_maker.php?bsID='+bsID+'&id_list='+id_list;
}
//////////////////////////////////////////////

function prases_saver()
{
//	alert("ЗАПИСЬ");	return;
var saveStr="";
var tr =0;
var nodes = document.getElementsByClassName('r_table'); //alert(nodes.length);
for(var i=0; i<nodes.length; i++) 
{
tr=nodes[i]; //alert(tr.cells[0].innerHTML+"|"+tr.cells[3].childNodes[0].value);return;

// пропускаем все, что не содержит фраз
if(tr.cells[3].childNodes[0].value==0)
	continue;

//alert(tr.cells[0].innerHTML+"|"+tr.cells[3].childNodes[0].value);return;
saveStr+=tr.cells[0].innerHTML+"|"+tr.cells[1].childNodes[0].value+"|"+tr.cells[2].childNodes[0].value+"|"+tr.cells[3].childNodes[0].value+"|||";
}//for(
//  alert(saveStr);return;

if(saveStr.length==0)
{
show_dlg_alert("Нет новых фраз-синонимов.",2000);
return;
}
//alert(saveStr); // return;
/////////////////////////
saveStr=cur_condition_choose+"&saveStr="+saveStr;  // alert(saveStr);return;
var link="/pages/condition_reflexes_basic_phrases_saver.php";
//alert(link);
var AJAX = new ajax_post_support(link,saveStr, sent_table_save,1);
AJAX.send_reqest();
function sent_table_save(res)
{
//show_dlg_alert(res,0);
if(res[0]!='!')
{ //alert(res);
show_dlg_alert(res,0);
return;
}
show_dlg_alert("Записаны новые фразы.",2000);
// перегрузить таблицу чтобы показывало дубликаты фраз
get_table(0);
}
}
////////////////////////////////

/////////////////////////////////////////
	// сохранение по Ctrl+S
var is_press_strl = 0;
document.onkeydown = function(event) { 
	
		var kCode = window.event ? window.event.keyCode : (event.keyCode ? event.keyCode : (event.which ? event.which : null))

		//alert(kCode);
		if (kCode == 17) // ctrl
			is_press_strl = 1;

		if (is_press_strl) {
			if (kCode == 83) {
				event.preventDefault();
//alert(is_table_shoved);
	if(!is_table_shoved)
		return false;
				//alert("!!!!! ");
				reflex_saver();
				is_press_strl = 0;
				return false;
			}
		}
}
////////////////////
document.onmouseup = function(event) { 
var t = event.target || event.srcElement;    
while(t)
{ 
if(t.id == "div_dlg_alert2")
	 return;	    
t = t.offsetParent;
}
end_dlg_alert2(); 
}
///////////////////////////
function reflex_saver()
{
show_dlg_confirm("Сохранить список?",1,-1,prases_saver);
}
function close_all_dlg()
{
// просто чтобы была такая пустая и не было варннига при закрытии по фону
}
	////////////////////////////
function set_sel(tr, id) {
		//	alert(id);
		var nodes = document.getElementsByClassName('highlighting'); //alert(nodes.length);
		for (var i = 0; i < nodes.length; i++) {
			nodes[i].style.border = "solid 1px #000000";
		}
		tr.style.border = "solid 2px #000000";
}



////////////////////////////////////////
function show_word_list(id)
{  

var AJAX = new ajax_support("/lib/get_exclamations.php?id=" + id, sent_words_info);
AJAX.send_reqest();
function sent_words_info(res) {
			//alert(res);
show_dlg_alert2("<br><span style='font-weight:normal;'>Щелкните по слову:<br>" + res + "<br>", 0);
}
}
function insert_word(id,word)
{
	word="    <br>"+word+" ";
	word=word.trim();
	word=word.replace(/<\/?[^>]+>/g,'');
end_dlg_alert2();
var inp=document.getElementById('insert_'+id);
inp.setRangeText(word, inp.selectionStart, inp.selectionEnd, "end");
/*
if(inp.value.length>0)
inp.value+=" ";
document.getElementById('insert_'+id).value=inp+word;
*/
//alert(word);
}
</script>


<?
/*  Примеры использования функций папки common
http://go/common/0_samples.php
*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");

header('Content-Type: text/html; charset=windows-1251');

include_once($_SERVER["DOCUMENT_ROOT"]."/check_prompt.php");
?><!DOCTYPE html>
<HTML>
<HEAD>
<TITLE>Примеры</TITLE>
<meta http-equiv="Content-Type" content="text/html; charset=windows-1251">
<meta http-equiv="Cache-Control" content="no-cache">
<meta http-equiv="pragma" content="no-cache">
</HEAD>
<BODY>
<?include_once($_SERVER['DOCUMENT_ROOT']."/common/common.php");?>

<h2>Примеры алерта</h2>

<script>
function get_alrtt_1()
{
show_dlg_alert("Самозакрывающийся алерт",1500);
}
function get_alrtt_2()
{
show_dlg_alert("Закройте алерт когда прочтете.",0);
}
function get_alrtt_3()
{ 
show_dlg_alert("Только крестик в алерте.",2);
}
function get_alrtt_4()
{ 
show_dlg_alert("И крестик и ОК в алерте.",3);
}
</script>
<input type='button' value='Самозакрывающийся алерт' onClick='get_alrtt_1()'>
<input type='button' value='Обычный алерт с ОК без крестика' onClick='get_alrtt_2()'> 
<input type='button' value='Только крестик в алерте' onClick='get_alrtt_3()'> 
<input type='button' value='И крестик и ОК в алерте' onClick='get_alrtt_4()'> 

<h2>Примеры конфирма</h2>

<script>
function get_confirm_1()
{
show_dlg_confirm("Уверены?",1,-1,todo_continue);
}
function todo_continue()
{
alert("Тогда удаляем!");
}


function get_confirm_2()
{
show_dlg_confirm("Сделать это?","Конечно","Ни в коем случае!",todo_continue2);
}
function todo_continue2()
{
alert("Делаем Это...");
}
</script>
<input type='button' value='Одна кнопка' onClick='get_confirm_1()'>
<input type='button' value='Две кнопки' onClick='get_confirm_2()'> 


<h2>Крутящаяся гифка ожидания</h2>
<script>
function get_waite_1()
{
wait_begin();
}
function get_waite_2()
{
wait_show();
}
function get_waite_3()
{
wait_end();
}
</script>
<input type='button' value='Начать прцесс с индикацией через 500ms' onClick='get_waite_1()'>
<input type='button' value='Начать прцесс  с индикацией сразу' onClick='get_waite_2()'>
<input type='button' value='Кончить прцесс' onClick='get_waite_3()'>


<h2>Спойлер со вложенными спойлерами</h2>

<span class="spoiler_header" onclick="open_close('block_id',1)" style="cursor:pointer;font-size:16pt"><?=set_sopiler_icon('block_id')?><b>Главный спойлер</b></span>
<div id="block_id" class="spoiler_block spoiler" style="height:0px;">
11111111111111<br>
222222222222222<br>

<span class="spoiler_header" onclick="open_close('block2_id',1)" style="cursor:pointer;font-size:14pt"><?=set_sopiler_icon('block2_id')?><b>Внутренний спойлер 1</b></span>
<div id="block2_id" class="spoiler_block spoiler" style="height:0px;">
aaaaaaaaaaa<br>
ssssssssss<br>

<span class="spoiler_header" onclick="open_close('block3_id')" style="cursor:pointer;"><b>Внутренний спойлер 2</b></span>
<div id="block3_id" class="spoiler_block spoiler" style="height:0px;">
EEEEEEEEE<br>
FFFFFFFFFFFF<br>
VVVVVVVVVVVV<br>
WWWWWWWWWWWWWWWW<br>
</div>

dddddddddd<br>
xxxxxxxxxxxx<br>
</div>

3333333333333<br>
44444444444444<br>
</div>


<br>
открывать-закрывать внутренние дивы: <span onclick="open_close('block2_id')" style="cursor:pointer;font-size:12pt"><b>open_close2</b></span> <span onclick="open_close('block3_id')" style="cursor:pointer;font-size:12pt"><b>open_close3</b></span><br>




<h2>Поддержка свойств полей ввода</h2>
<?include_once($_SERVER['DOCUMENT_ROOT']."/common/input.php");

echo "<input class='input_folder' type='text' style='color:#808F9E;border:solid 1px #808F9E;width:260px;'  ".set_input_mask("Введите число, не более 10 символов","")." ".only_numbers_input(10).">";

echo " <input class='input_folder' type='text' style='color:#808F9E;border:solid 1px #808F9E;width:260px;'  ".set_input_mask("Введите строку латинских символов","")." ".only_latin_input().">";

?>



</BODY>
</HTML>
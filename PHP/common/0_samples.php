<?
/*  ������� ������������� ������� ����� common
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
<TITLE>�������</TITLE>
<meta http-equiv="Content-Type" content="text/html; charset=windows-1251">
<meta http-equiv="Cache-Control" content="no-cache">
<meta http-equiv="pragma" content="no-cache">
</HEAD>
<BODY>
<?include_once($_SERVER['DOCUMENT_ROOT']."/common/common.php");?>

<h2>������� ������</h2>

<script>
function get_alrtt_1()
{
show_dlg_alert("����������������� �����",1500);
}
function get_alrtt_2()
{
show_dlg_alert("�������� ����� ����� ��������.",0);
}
function get_alrtt_3()
{ 
show_dlg_alert("������ ������� � ������.",2);
}
function get_alrtt_4()
{ 
show_dlg_alert("� ������� � �� � ������.",3);
}
</script>
<input type='button' value='����������������� �����' onClick='get_alrtt_1()'>
<input type='button' value='������� ����� � �� ��� ��������' onClick='get_alrtt_2()'> 
<input type='button' value='������ ������� � ������' onClick='get_alrtt_3()'> 
<input type='button' value='� ������� � �� � ������' onClick='get_alrtt_4()'> 

<h2>������� ��������</h2>

<script>
function get_confirm_1()
{
show_dlg_confirm("�������?",1,-1,todo_continue);
}
function todo_continue()
{
alert("����� �������!");
}


function get_confirm_2()
{
show_dlg_confirm("������� ���?","�������","�� � ���� ������!",todo_continue2);
}
function todo_continue2()
{
alert("������ ���...");
}
</script>
<input type='button' value='���� ������' onClick='get_confirm_1()'>
<input type='button' value='��� ������' onClick='get_confirm_2()'> 


<h2>���������� ����� ��������</h2>
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
<input type='button' value='������ ������ � ���������� ����� 500ms' onClick='get_waite_1()'>
<input type='button' value='������ ������  � ���������� �����' onClick='get_waite_2()'>
<input type='button' value='������� ������' onClick='get_waite_3()'>


<h2>������� �� ���������� ����������</h2>

<span class="spoiler_header" onclick="open_close('block_id',1)" style="cursor:pointer;font-size:16pt"><?=set_sopiler_icon('block_id')?><b>������� �������</b></span>
<div id="block_id" class="spoiler_block spoiler" style="height:0px;">
11111111111111<br>
222222222222222<br>

<span class="spoiler_header" onclick="open_close('block2_id',1)" style="cursor:pointer;font-size:14pt"><?=set_sopiler_icon('block2_id')?><b>���������� ������� 1</b></span>
<div id="block2_id" class="spoiler_block spoiler" style="height:0px;">
aaaaaaaaaaa<br>
ssssssssss<br>

<span class="spoiler_header" onclick="open_close('block3_id')" style="cursor:pointer;"><b>���������� ������� 2</b></span>
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
���������-��������� ���������� ����: <span onclick="open_close('block2_id')" style="cursor:pointer;font-size:12pt"><b>open_close2</b></span> <span onclick="open_close('block3_id')" style="cursor:pointer;font-size:12pt"><b>open_close3</b></span><br>




<h2>��������� ������� ����� �����</h2>
<?include_once($_SERVER['DOCUMENT_ROOT']."/common/input.php");

echo "<input class='input_folder' type='text' style='color:#808F9E;border:solid 1px #808F9E;width:260px;'  ".set_input_mask("������� �����, �� ����� 10 ��������","")." ".only_numbers_input(10).">";

echo " <input class='input_folder' type='text' style='color:#808F9E;border:solid 1px #808F9E;width:260px;'  ".set_input_mask("������� ������ ��������� ��������","")." ".only_latin_input().">";

?>



</BODY>
</HTML>
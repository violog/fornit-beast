<?
/*
поддержка свойств полей ввода
include_once($_SERVER['DOCUMENT_ROOT']."/common/input.php");
*/

/* серая подсказка в поле ввода
НУЖНО ЗАРАНЕЕ ПРОСТАВИТЬ style='color:#808F9E;' И НЕ СТАВИТЬ value=!!!
*/
function set_input_mask($text_mask,$value="")
{
if(empty($value))
{// СТРОГО В ОДНУ СТРОКУ!
$out = <<<EOD
	ONFOCUS="if(this.value == '$text_mask'){this.value ='';this.style.color='#000000';}" ONBLUR="if(this.value == '') {this.value = '$text_mask';this.style.color='#808F9E';}"  value="$text_mask"
EOD;
}

return "value='".$value."'";
}
///////////////////////////

//////////////////////
// позволяет вводить только целые цифры
function only_int_inp($limit=0)
{
	// СТРОГО В ОДНУ СТРОКУ!
$out = <<<EOD
onKeyDown='only_int_inp(this,$limit)' onKeyUp='only_int_inp(this,$limit)' onMouseUp='only_int_inp(this,$limit)'
EOD;
return $out;
}
// позволяет вводить только цифры и точку  <input ".only_numbers_input().">
function only_numbers_input($limit=0)
{
	// СТРОГО В ОДНУ СТРОКУ!
$out = <<<EOD
onKeyDown='only_numbers_inp(this,$limit)' onKeyUp='only_numbers_inp(this,$limit)' onMouseUp='only_numbers_inp(this,$limit)'
EOD;
return $out;
}
// позволяет вводить только цифры, точку, запятую, | и ;
function only_numbers_and_sybm_input($limit=0)
{
$out = <<<EOD
onKeyDown='only_numbers_and_sybm_inp(this,$limit)' nKeyUp='only_numbers_and_sybm_inp(this,$limit)' nMouseUp='only_numbers_and_sybm_inp(this,$limit)'
EOD;
return $out;
}

// позволяет вводить только цифры,  и запятую
function only_numbers_and_Comma_input($limit=0)
{
$out = <<<EOD
onKeyDown='only_numbers_and_Comma_input(this,$limit)' onKeyUp='only_numbers_and_Comma_input(this,$limit)' onMouseUp='only_numbers_and_Comma_input(this,$limit)'
EOD;
return $out;
}
////////////////////////////
// позволяет вводить только латинские">
function only_latin_input()
{
$out = <<<EOD
onKeyDown='only_latin_inp(this)' onKeyUp='only_latin_inp(this)' onMouseUp='only_latin_inp(this)'
EOD;
return $out;
}
// позволяет вводить только латинские, символ '-' и цифры  <input ".only_latin_input().">
function only_latin_and_num_input()
{
$out = <<<EOD
onKeyDown='only_latin_and_num_inp(this)' onKeyUp='only_latin_and_num_inp(this)' onMouseUp='only_latin_and_num_inp(this)'
EOD;
return $out;
}

// обрезает пробелы слева и справа
function trim_input()
{
$out = <<<EOD
onKeyDown='trim_inp(this)' onKeyUp='trim_inp(this)' onMouseUp='trim_inp(this)'
EOD;
return $out;
}
?>
<style>
.input_folder
{
border:0;
outline:none;
box-sizing:border-box;
padding-left:4px;
padding-right:4px;
}
</style>

<script>
function only_int_inp(inp,limit)
{
var val=inp.value;
inp.value=val.replace(/[^0-9,\-]/g,'');
if(limit>0)
	{
inp.value=inp.value.substr(0,limit);
	}
}
function only_numbers_inp(inp,limit)
{
var val=inp.value;
inp.value=val.replace(/[^0-9.\-]/g,'');
if(limit>0)
	{
inp.value=inp.value.substr(0,limit);
	}
}
function only_numbers_and_sybm_inp(inp,limit)
{ 
var val=inp.value;
inp.value=val.replace(/[^0-9.,-:;|\>\<]/g,'');
if(limit>0)
	{
inp.value=inp.value.substr(0,limit);
	}
}
function only_numbers_and_Comma_input(inp,limit)
{  
var val=inp.value;
inp.value=val.replace(/[^0-9,]/g,'');
if(limit>0)
	{
inp.value=inp.value.substr(0,limit);
	}
}
//////////////////////////////
function only_latin_inp(inp)
{
var val=inp.value;
inp.value=val.replace(/[^a-zA-Z]/g,'');
}
//////////////////////
function only_latin_and_num_inp(inp)
{
var val=inp.value;
inp.value=val.replace(/[^0-9a-zA-Z\-]/g,'');
}
//////////////////////
function trim_inp(inp)
{
var val=inp.value;
inp.value=val.trim();
}
</script>
<?
/* плавный спойлер
include_once($_SERVER['DOCUMENT_ROOT']."/common/spoiler.php");


<span class="spoiler_header" onclick="open_close('lib_block1_id',1)" style="cursor:pointer;font-size:16px"><?=set_sopiler_icon('lib_block1_id')?><b>ПРАВИЛА</b></span>
<div id="lib_block1_id" class="spoiler_block spoiler" style="position:relative;z-index:10;top:0px;left:0px;padding-left:15px;background-color:#ffffff;width:1100px;height:0px;">

ТЕЛО СПОЙЛЕРА

</div>
*/
function set_sopiler_icon($block_id)
{
echo "<span id='".$block_id."_icon'>&#9660;</span>";
}
?>
<style>
.spoiler_header 			/* заголовок спойлера */
{
margin-left:0px;
background-color:;
cursor:pointer;
padding:4px 4px 4px 0px;
color:#00769F!important;
font-size:19px;
font-weight:bold;
letter-spacing:1px;
}
.spoiler_header:hover
{
background-color:#B6E8F2;
color: #000000!important;
}
.spoiler_block 			/* спойлер  */
{
#border:solid 1px #8A3CA4;
overflow: hidden;
transition: height 500ms ease;
}
</style>

<script>
function open_close(block_id,incon=0)
{  
var elBlock=document.getElementById(block_id); 
if(parseInt(elBlock.style.height)==0)
{
elBlock.clientHeight;
var h=elBlock.scrollHeight; //alert(h);
elBlock.style.height=h+"px";

if(incon)
document.getElementById(block_id+"_icon").innerHTML="&#9650;";

setTimeout("free_height('"+block_id+"')",500);// время == transition: height 500ms
}
else
{
elBlock.clientHeight;
var h=elBlock.scrollHeight;
elBlock.style.height=h+"px"; //теперь transition: height имеет числовой параметр

if(incon)
document.getElementById(block_id+"_icon").innerHTML="&#9660;";

setTimeout("close_height('"+block_id+"')",10);// время не играет роли, лишь бы выйти из стека функции
}

}
function free_height(block_id)
{
var elBlock=document.getElementById(block_id);
// отпускаем фиксированную высоту: блок будет раскрываться по содержимому и внутренние спойлеры будут нормально работать
elBlock.style.height="";
}
function close_height(block_id)
{
var elBlock=document.getElementById(block_id);
elBlock.style.height="0px";
}
</script>
<?
/* 
include_once($_SERVER['DOCUMENT_ROOT']."/common/alert_control.php");

всплавает рядом с родителем - для показа контролов выбора.

Применение:
show_dlg_control(mess,this);   end_dlg_control()
	
Закрывается только крестиком или по end_dlg_control()
*/


?>
<style>
.dlg_control_exit /* кнопка крестика выхода */
{
position:absolute;
top:1px;
right:2px;
border:0;
border-radius:4px;
color:#FFFFFF;
text-align:center;
cursor:pointer;
background:linear-gradient(180deg, #774E9D, #3D146B);
font-size:15px;
padding:1px;
width:20px;
}
.dlg_control_exit:hover /* кнопка крестика выхода */
{
background:linear-gradient(180deg, #F08612, #F08612);
}
</style>
<div id="div_dlg_control"
style="
position:fixed;
z-index: 10000;
max-height:600px;
#min-width:300px;
overflow:auto;
display:none;
padding: 20px;
#top: 50%;left: 50%;
#transform: translate(-50%, -50%);
border:solid 2px #8A3CA4;
color:#000000;background-color:#eeeeee;
font-size:16px;font-family:Arial;
#font-weight:bold;
box-shadow: 8px 8px 8px 0px rgba(122,122,122,0.3);border-radius: 10px;
text-align:center;
"
onmouseup='event.preventDefault();'></div>


<script>
function show_dlg_control(mess,parent)
{  
var dlg_control=document.getElementById('div_dlg_control');
dlg_control.innerHTML=mess+"<div class='dlg_control_exit' style='top:0.0vw; right:0.0vw;' title='закрыть' onClick='end_dlg_control();'><span style='position:relative; top:-1px; left:1px;'>&#10006;</span></div>";

var width  = window.innerWidth || document.documentElement.clientWidth || 
document.body.clientWidth;
var height = window.innerHeight|| document.documentElement.clientHeight|| 
document.body.clientHeight;

dlg_control.style.display="block";  
var p_box = parent.getBoundingClientRect(); //alert(box['top']+" | "+box['left']);
var m_box = dlg_control.getBoundingClientRect();
// есть ли место справа
var m_width=m_box['right']-m_box['left'];
//alert(m_width+" || "+width);
if (p_box['right']+m_width > width)
{
	dlg_control.style.left=(width -m_width)+"px";
}
else
dlg_control.style.left=p_box['left']+"px";

// есть ли место снизу
var m_height=m_box['bottom']-m_box['top'];
if (p_box['top']+m_height > height)
{
	dlg_control.style.top=(height -m_height)+"px";
}
else
dlg_control.style.top=p_box['top']+"px";

}
function end_dlg_control()
{ //dlg_control("!!!!!!");
document.getElementById('div_dlg_control').style.display="none";
}
//////////////////////////////////////////////////////

// закрывается по window.onmouseup  в /common/header.php
</script>
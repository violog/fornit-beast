<?
/* слайдер страниц для выбора из ряда. Только ОДИН такой слайдер может быть на данной странице

Подключение: include_once($_SERVER['DOCUMENT_ROOT']."/common/page_slider.php");
Пример использования:
_____________________________________________________
if(isset($_GET['page_txt'])) $page_txt=$_GET['page_txt']; else $page_txt=799;

include_once($_SERVER['DOCUMENT_ROOT']."/common/page_slider.php");
$link="/common/page_slider_SAMPLE.php?avtor=aaa&page_txt=[N]&id_desk=11";
$page_str = new page_slider;
$page_str->init(800,$page_txt,$link,1,0,"font-famaly:arial;font-size:14px;");
$page_str->show();// верхняя строка страниц

echo "fadfa dasfasdfasdf asdfasdfa<br>
asdf asdfasdfasdfasdfasdfas<br>";
$page_str->show();// нижняя строка страниц
_____________________________________________________

Параметры функции init($max_count,$cur_num,$link_temp,$k,$shift,$style=""):
-$max_count - число страниц 
-$cur_num -  текущий номер в ряду номеров страниц
-$link_temp - href ссылок на страницу в виде: "xxxx=[N]"
-$k - множитель для N, например, в ссылках форуме $k==15 (число сообщений на странице), так что для N==4 будет подставляться 4*15 -> 45.
При $max_count==1 строка ссылок страниц не выводится т.к. - всего одна страница.
-$shist - число, добавляемое к N
т.е. в результате  в ссылке будет посдстваляться вместо [N] -> ($i+$shist)*$k
-$style - стиль вывода номеров страниц
*/

class page_slider
{
private $id=0;   
private $max;
private $cur;
private $tlink;
private $shift;
private $k;
private $style;


public function init($max_count,$cur_num,$link_temp,$k,$shift,$style="")
{
$this->max=$max_count;
$this->cur=$cur_num;
$this->tlink=$link_temp;
$this->k=$k;
$this->shift=$shift;
$this->style=$style;
}

public function show()
{
$this->id++; 
$id=$this->id;
//$pages_str="".$this->id;

if($this->max==1)
	return "";

$nShow=20;
$isSlider=0;
if($this->max>$nShow)// ставим слайдер
{
$isSlider=1;
$width=200;
$curSlayderVal=$this->max-$nShow;
$diapazon=$this->max-$nShow;
$val=$this->cur;
//exit("$diapazon");
$min=0;

$pages_str="
<div style='position:relative;'>Страницы: 
&nbsp;<input id='range_page_slider_".$id."'
type='range' class='page_slider".$id."'
style='width:".$width."px'
min='0' max='".$diapazon."' step='1' value='".$val."' 
title='Диапазон страниц для выбора'  
onChange='cSlader".$id.".set(this,".$id.")'
onInput='cSlader".$id.".set(this,".$id.")'
>&nbsp;
";// onInput='cSlader".$id.".set(this,".$id.")'
// блок для смены диапазона страниц
$pages_str.="<span id='div_page_slider_".$id."'>";

$js_arr="var pagesArr".$id." = new Array();";
for($i=0; $i<$this->max; $i++)
{
$link=str_replace("[N]", ($i+$this->shift)*$this->k, $this->tlink);
$cur="<a class='noUn".$id."' href='".$link."'> ".($i+1)." </a>";
// пропускаем цифры после первой - до предпоследней+$nShow
if( $i<$this->cur-$nShow/2 || $i>$this->cur+$nShow/2)
{
$js_arr.="pagesArr".$id."[".$i."] = \"".$cur."\";";
continue;
}

if($this->cur == $i+$this->shift) 
{	
	$pages_str.="<span  class='noUn".$id."'><b> ".($i+1)." </b><span>"; 
	$js_arr.="pagesArr".$id."[".$i."] = '<b> ".($i+1)." </b>';";
} 
else 
{
$pages_str.=$cur;
$js_arr.="pagesArr".$id."[".$i."] = \"".$cur."\";";
}

//$page_last++;// чтобы отматывалось в конец:
}//for(

if($isSlider)
$pages_str.="</span>";
$pages_str.="</div>";

$pages_str.="<script>
".$js_arr."
var cSlader".$id." = new page_num_control".$id."(pagesArr".$id.",'".$val."','".$this->max."','".$nShow."','".$this->cur."');
</script>";
}
else// слайдер не нужен
{
$pages_str="";
for($i=0; $i<$this->max; $i++)
{
$link=str_replace("[N]", ($i+$this->shift)*$this->k, $this->tlink);
$cur="<a class='noUn".$id."' href='".$link."'> ".($i+1)." </a>";
// пропускаем цифры после первой - до предпоследней+$nShow
if( $i<$this->cur-$nShow/2 || $i>$this->cur+$nShow/2)
{
//$js_arr.="pagesArr".$id."[".$i."] = \"".$cur."\";";
continue;
}

if($this->cur == $i+$this->shift) 
{	
	$pages_str.="<span  class='noUn".$id."'><b> ".($i+1)." </b><span>"; 
} 
else 
{
$pages_str.=$cur;
}

//$page_last++;// чтобы отматывалось в конец:
}//for(
}

//////////////////////
$styles="";
if(!empty($this->style))
{
$sArr=explode(";",$this->style);
foreach($sArr as $s)
{
$s=trim($s);
if(empty($s))
	continue;
$styles.=$s.";";
}
}
?>
<style>
.page_slider<?=$id?>
{
position:relative;
top:4px;
#width:100px;
cursor: pointer;
}
.noUn<?=$id?>
{
text-decoration:none;
<?=$styles?>
}
</style>

<script>
function page_num_control<?=$id?>(pagesArr,begin,max,count,cursel)
{
//alert(typeof(pagesArr));
this.set=function(slider,exemplar)
{
//	alert(slider+" | "+exemplar);
var begin=slider.value;
begin=1*begin;
max=1*max;
count=1*count;
cursel=1*cursel;

// alert(begin+" | "+max+" | "+count);
if(begin>max-count)
	begin=max-count;

var out="";
for(i=begin; i<max; i++)
{ //alert(pagesArr[i]);
//	if(i==300)alert(begin+count);
if(i>=begin+count)
	break;
if(i + <?=$this->shift?> == cursel)
out+="<span  class='noUn<?=$id?>'><b> "+(i+1)+" </b><span>"; 
else
out+=pagesArr[i];
}
//alert(out);
document.getElementById('div_page_slider_'+exemplar).innerHTML=out;
}

}
</script>
<?
echo $pages_str;

}




}
//////////////////////////////////////////////////////////
?>

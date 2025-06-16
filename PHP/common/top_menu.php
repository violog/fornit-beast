<?
/* верхнее меню на все старницы

$page_id=0;- из headr.php
include_once($_SERVER["DOCUMENT_ROOT"]."/common/top_menu.php");
*/
echo "
<style>
.menu_item
{
font-size:14px;
border:solid 1px #8A3CA4;
border-radius: 5px;
box-shadow: 2px 2px 2px 0px rgba(122,122,122,0.3);
margin-left:10px;
padding-left:4px;
padding-right:4px;
cursor:pointer;
}
</style>
";
if(isset($page_id))
{
// все страницы меню
$p_nameArr=array(
0 => array("Пульт","/","Пульт связи с Beast и управления его состоянием"),	
1 => array("Гомеостаз","/pages/gomeostaz.php","Редактор параметров гомеостаза"),
2 => array("Слова","/pages/words.php","Заливка текстов для дерева слов"),
3 => array("Действия","/pages/terminal_actions.php","Редактор возможных Действий"),
4 => array("Рефлексы","/pages/reflexes.php","Редактор безусловных рефлексов"),
5 => array("Ус. рефлексы","/pages/condition_reflexes.php","Условные рефлексы Beast"),
6 => array("Автоматизмы","/pages/automatizm.php","Автоматизмы Beast"),
7 => array("Психика","/pages/self_perception.php","Страница текущего состояния психики Beast"),
8 => array("Сознание","/pages/conscience.php","Страница текущего состояния функции осознавания Beast"),

10 => array("Стадии","/pages/stages.php","Стадии развития Beast"),

);

?>
<script>
var is_menu_click=false;
function isMenuClick()
{
//alert("isMenuClick");
is_menu_click=true;// is_menu_click=false; - в /index.php
}
// обновить без закрытия го-сервера
function page_refresh()
{
var AJAX = new ajax_support(linking_address+'?no_close_weght=1',sent_no_close);
AJAX.send_reqest();
function sent_no_close(res)
{
location.reload(true);
}
}
</script>
<?

echo "<div style='width:1000px'>";

// кнопка обновления без закрытия го-сервера
echo "<span class='menu_item' style='margin-left:-20px;background-color:#D7E8FF;' title='Обновить страницу без закрытия Beast' onClick='page_refresh();' style='".$mstyle."' >&empty;</span>";


foreach($p_nameArr as $k => $m)
{ 
$mstyle="";
if($k == $page_id)
$mstyle="background-color:#3D146B;color:#ffffff;";

echo "<span class='menu_item' title='".$m[2]."' onClick='isMenuClick();location.href=`".$m[1]."`' style='".$mstyle."'>".$m[0]."</span>";
}
echo "</div>";
}

?>
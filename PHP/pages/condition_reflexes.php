<?
/*   Условные рефлексы
http://go/pages/condition_reflexes.php  

Формат записи:
ID|lev1|lev2 через ,|lev3 типа lev3TriggerStimulsID|ActionIDarr через ,|rank|lastActivation|activationTime
*/

$page_id = 5;
$title = "Условные рефлексы Beast";
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/show_waiting.php");

$out_str_for_del = ""
?>


<form name="refresh" method="post" action="/pages/condition_reflexes.php"></form>
<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
	var linking_address = '<? include($_SERVER["DOCUMENT_ROOT"] . "/common/linking_address.txt"); ?>';
	//alert(linking_address);
	function end_deleting() {
		//show_dlg_alert("Beast выключается...", 2000);
		
		//Выключить без сохранения памяти (bot_closing=2), просто погасить исполняемый файл
		var AJAX = new ajax_support(linking_address + "?bot_closing=2", sent_info);
		AJAX.send_reqest();
		function sent_info(res) {
			// не будет ответа
		}
		/*		
		var server = "/kill.php";
		var AJAX = new ajax_support(server, sent_end_answer);
		AJAX.send_reqest();
		function sent_end_answer(res) {
			show_dlg_alert("Beast выключен.", 2000);
		}*/
	}
</script>
<?

// удаление рефлексов
if (isset($_POST['rdelID'])) {  //var_dump($_POST['rdelID']);exit();
	$delArr = explode("|", $_POST['rdelID']);
	$dArr = array();
	foreach ($delArr as $s) {
		if (empty($s)) {
			continue;
		}
		array_push($dArr, substr($s, 6));  //var_dump($dArr);exit();

	}
	$dCount = count($dArr);   //var_dump($dArr);exit();

	$str = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/condition_reflexes.txt");
	$list = explode("\r\n", $str);  //exit("! $str | ".$_GET['delete_id']);
	$out = "";
	$isDeleted = 0;
	foreach ($list as $s) {
		if (empty($s)) {
			continue;
		}
		$r = explode("|", $s);
		$id = $r[0];
		$isDeleting = 0;
		for ($n = 0; $n < $dCount; $n++) { //if($id=="6252") exit("WWWWWWWW");
			if ($id == $dArr[$n]) {
				$isDeleting = 1;
				break;
			}
		}
		if (!$isDeleting) {
			$out .= $s . "|\r\n"; // exit($out);
		}
	}
	//exit(": ".$out);
	// !!! удалять нужно после выключения Beast т.к. там запоминаются файлы при выходе
	write_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/condition_reflexes.txt", $out);
	//exit(": ".$out);
	echo "<script>";
	// вырубить Beast и записать рефлексы после этого
	echo "show_dlg_alert('Выбранные рефлексы удалябтся.<br>Beast выключается.',1500);
		setTimeout(`end_deleting();`,2000);";

	echo "</script>";
	exit();
}
?>
<div style='position:relative;margin-top:-10px;height:70px;width:1000px;'>

<div style='position:absolute;top:-20px;right:0px;font-family:courier;font-size:16px;cursor:pointer;' onClick="open_anotjer_win('/pages/condition_reflexes.htm')"><b>Пояснения</b></div>

<div style='position:absolute;top:-20px;left:250px;font-family:courier;font-size:16px;cursor:pointer;' onClick="location.reload(true)"><b>Обновить</b></div>

<div style='position:absolute;top:-20px;left:360px;font-size:16px;cursor:pointer;color:#7E58FF;' onClick="open_anotjer_win('/pages/condition_reflexes_basic_phrases.php')" title="Создание базы простейщих фраз для заливки базы условных рефлексов."><b>Набить базовые фразы</b></div>

<div style='position:absolute;top:-20px;right:100px;font-size:16px;cursor:pointer;color:#7E58FF;background-color:#eeeeee;padding-left:4px;padding-right:4px;border:solid 1px #8A3CA4;border-radius: 7px;' onClick="open_anotjer_win('/pages/condition_reflexes_basic_phrases_maker.php')" title="Сформировать условные рефлексы на основе списка фраз-синонимов.
ПОСЛЕ ПОЛНОЙ ГОТОВНОСТИ ФРАЗ-СИНОНИМОВ!"><b>Сформировать условные рефлексы</b></div>

<div style='position:absolute;top:0px;left:360px;font-size:16px;cursor:pointer;color:#7E58FF;' onClick="cliner_reflex_times()" title="Чтобы рефлексы не просрочили свое время жизни, нужно обновлять его перед началом использования."><b>Обновить время жизни рефлексов</b></div>

<div id='div_id' style='position:absolute;top:20px;left:0px;font-family:courier;font-size:18px;'><b>Нужен коннект с Beast.</b></div>

<?
if($stages>1)
echo "<div style='position:absolute;top:40px;left:0px;font-size:21px;color:red;'><b>Это - пройденная стадия развития! Не следует редактировать данные на этой станице!</b></div>";
?>

</div>

Рефлексы можно только удалить, после чего нужно перезапустить Beast. Чтобы создался новый рефлекс нужно не менее 3-х раз повторить воздействие пусковых стимулов, не обязательно подряд, - этим предотвращаются случайные, мусорные сочетания.
Для формирования условных рефлексов нужно потратить немало времени (в случае ребенка около года). Это - период взаимодействия с Beast любым образом, с разными сочетаниями действий и очень простых фраз <b>в различных состояниях его Базовых параметров</b> (для этой стадии развития можно устанавливать слайдерами Пульта).
Чем дольше этот период, тем более эффективные навыки получит Beast.

<div id='reflex_info_id' style='font-family:courier;font-size:16px;'></div>
<div style="position:relative;">
<input id="del_btn_id" type="button" value="Удалить выбранные рефлексы" style="position:absolute;top:0;right:5px;display:none;" onclick='delete_reflexes()'></div>

<form id="form_del" name="form_del" method="post" action="/pages/condition_reflexes.php">
	<input type="hidden" name="rdelID" value="">
</form>

<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script>
	var linking_address = '<? include($_SERVER["DOCUMENT_ROOT"] . "/common/linking_address.txt"); ?>';

// ждем пока не включат бестию
check_Beast_activnost(4);// после 4-го пульса И запускается get_info()

	var old_size = 0;
var limitBasicID=0;//>0 - лимитировать показ только одним из базовых состояний Плохо,Норма,Хорошо 
	function get_info() {
		wait_begin();

		var AJAX = new ajax_support(linking_address + "?limitBasicID="+limitBasicID+"&get_condition_reflex_info=1", sent_get_info);
		AJAX.send_reqest();

		function sent_get_info(res) {
			//alert(res);
			wait_end();
			document.getElementById('reflex_info_id').innerHTML = res;
			document.getElementById('del_btn_id').style.display = "block";
		}
	}
//	get_info();

function show_level(base)
{
	limitBasicID=base;
get_info();
}
//////////////////////////////////////////////

	function delete_reflexes(id) {
		show_dlg_confirm("Точно удалить?", "Да", "Нет", delete_reflexes2);
	}

	function delete_reflexes2() {
		var del_str = "";
		var nodes = document.getElementsByClassName('deleteCHBX'); //alert(nodes.length);
		for (var i = 0; i < nodes.length; i++) {
			if (nodes[i].checked) {
				del_str += nodes[i].id + "|";
			}
		}
		if (del_str.length == 0) {
			show_dlg_alert("Не выбраны рефлексы для удаления.", 0);
			return;
		}
		document.forms.form_del.rdelID.value = del_str;
		document.forms.form_del.submit();
	}

function cliner_reflex_times()
{
var AJAX = new ajax_support(linking_address + "?cliner_time_condition_reflex=1", sent_get_info);
AJAX.send_reqest();
function sent_get_info(res) 
{
show_dlg_alert('Время жизни условных рефлексов очищено.',2000);
		setTimeout(`end_deleting();`,2000);
}
}
</script>

</body>
</html>
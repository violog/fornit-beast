<?
/* Редактор возможных Действий Beast
http://go/pages/terminal_actions.php */

$page_id = 3;
$title = "Редактор возможных Действий Beast";
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");

if (filesize($_SERVER['DOCUMENT_ROOT'] . "/memory_reflex/condition_reflexes.txt") > 6) {
	echo "<div style='color:red;border:solid 1px #8A3CA4;padding:10px;background-color:#DDEBFF;'>
	Этот редактор <b>НЕ СЛЕДУЕТ ИСПОЛЬЗОВАТЬ</b> потому, что уже есть условные рефлексы.<br>
	Чтобы использовать редактор, нужно сбросить память Beast (на странице Пульса справа вверху нажать шестеренку и выбрать &quot;
	Сбросить память&quot;) <br>или <b>просто удалить содержимое в файле /memory_reflex/condition_reflexes.txt</b></div>";
}
?>

<div class="main_page_div">
	<script Language="JavaScript" src="/ajax/ajax_form_post.js"></script>
	<script Language="JavaScript" src="/ajax/ajax.js"></script>

	Практически все описанные в виде слов и фраз действия - в природе реализуются цепочками безусловных рефлексов (инстинктами).
	Поэтому отпадает необходимость в редакторе безусловных рефлексов делать реакции на основе других реакций потому как можно
	просто написать фразу, соответствующую любому такому действию.<br>
	Данный редактор связывает действие с тем, какие гомео-параметры улучшает данное действие. Фактически это - самый простой безусловный рефлекс.
	Более сложные безусловные рефлексы редактируются на странице Рефлексы и если там для данных условий не оказывается прописанного рефлекса
	(и нет более высокоуровневых автоматизмов), то срабатывает рефлекс данного редактора. В этом проявляется иерархическая конкурентность рефлексов.
	<h2 class="header_h2">ID параметров <a href="/pages/gomeostaz.php">гомеостаза</a>:</h2>
	1 <b>Энергия</b> (Уменьшается со временем и расходовании)<br>
	2 <b>Стресс</b> (Накапливается в течении дня и снимается во время сна. Увеличивается при стрессовых ситуациях)<br>
	3 <b>Гон</b> (Жизненный параметр данного вида. Постепенно нарастает и требует разрядки)<br>
	4 <b>Потребность в общении</b> (Жизненный параметр данного вида. Постепенно нарастает и требует разрядки)<br>
	5 <b>Потребность в обучении</b> (Зависит от ситуации, но нарастает пока не будет разрядки)<br>
	6 <b>Поиск</b> (Основа поискового поведения. Зависит от ситуации, но нарастает в депривации)<br>
	7 <b>Самосохранение</b> (Жадность, эгоизм, самозащита, страх смерти. Зависит от ситуации, может сам уменьшаться при благополучии)<br>
	8 <b>Повреждения</b> (Параметр общего состояния организма. Повреждения нарастают со временем.)<br>

	<hr>
	<div style="position:relative">
		<h2 class="header_h2">Возможные действия:</h2>
		<div style="position:absolute;top:0;left:200px;color:red;" title="Первые строки зеркально отражают действия с Пульта чтобы Beast 
	мог их отзеркаливать как чужой опыт.">Серые строки в начале должны быть обязательно!</div>
		<a style="position:absolute;top:0px;right:0px;" href="/pages/terminal_actions_help.htm" target="_blank">Пояснение как заполнять таблицу</a>
	</div>

	<form id="form_id" name="form" method="post" action="/pages/terminal_actions.php">
		<table id="main_table" class="main_table" cellpadding=0 cellspacing=0 border=1 width='100%' style="font-size:14px;">
			<tr>
				<th width=70 class='table_header'>ID<br>действия</th>
				<th width='' class='table_header'>Описание действия</th>
				<th width='20%' class='table_header' title='Затраты разделять точкой с запятьй'>Гомеостатические затраты<br> на действие средней силы</th>
				<th width='20%' class='table_header' title='ID разделять запятыми'>Какие ID гомео-параметров<br>улучшает действие</th>
				<th width='30' class='table_header' title="Удалить рефлекс">Х</th>
			</tr>
			<?
			// считать файл 
			$progs = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/terminal_actons.txt");
			$strArr = explode("\r\n", $progs);  //var_dump($strArr);exit();
			$n = 0;
			$lastID = 1;
			$lastAdminNum = 22; // последний ID котрый должен редактировать только специсалист

			foreach ($strArr as $str) {
				if (empty($str) || $str[0] == '#')
					continue;
				$par = explode("|", $str); //var_dump($par);exit();
				$id = $par[0];

				$bg = "";
				if ($id <= $lastAdminNum)
					$bg = "style='background-color:#eeeeee;'";
				if ($id == 1) {
					echo "<tr ><td colspan=5 align='center'>редактируется только специалистом (они жестко прописаны в теле Best)</td></tr>";
				}
				if ($id == 30) {
					echo "<tr ><td colspan=5 align='center'>начало свободно редактируемых действий</td></tr>";
				}
				echo "<tr >";
				echo "<td class='table_cell' style='width:40px;background-color:#eeeeee;'><input type='hidden' name='id[" . $id . "]' value='" . $par[0] . "'  >" . $par[0] . "</td>
				<td class='table_cell'><input class='table_input' " . $bg . " type='text' name='decr[" . $id . "]'   value='" . $par[1] . "'  ></td>
				<td class='table_cell'><input class='table_input' " . $bg . " type='text' name='val[" . $id . "]' " . only_numbers_and_sybm_input() . "  value='" . $par[2] . "'  ></td>
				<td class='table_cell'><input class='table_input' " . $bg . " type='text' name='target[" . $id . "]' " . only_numbers_and_sybm_input() . "  value='" . $par[3] . "'  ></td>
				<td class='table_cell' align='center' title='Удалить действие'><img src='/img/delete.gif' onClick='check_and_del(" . $n . "," . $id . ")'></td>
				</tr>";

				$n++;
				$lastID = $id + 1;
			}
			?>
		</table>
		<div style="position:relative;">
			<input type='hidden' name='gogogo' value='1'>
			<input style="position:absolute;top:0px;right:0px;" type="button" name="save" value="Сохранить" onClick="check_and_save()">
			<input style="position:absolute;top:0px;left:0px;" type="button" name="addnew" value="Добавить новую строку" onClick="add_new_line()">
		</div>
	</form>

	<script>
		function check_and_del(delete_id, id) {
			var server = "/pages/terminal_actions_delete.php?akt=" + id;
			var AJAX = new ajax_support(server, sent_request);
			AJAX.send_reqest();

			function sent_request(res) {
				if (res != "") {
					alert(res);
					return;
				}else{
					if (confirm("Точно удалить?")){
						var server = "/pages/terminal_actions_delete.php?delete_id=" + delete_id + "&id=" + id;
						var AJAX = new ajax_support(server, sent_request_res);
						AJAX.send_reqest();
						function sent_request_res(res) {
							if (res != '!') {
								show_dlg_alert(res, 0);
								return;
							}
							setTimeout("location.reload(true)", 100);
						}
					}
				}
			}
		}

		function check_and_save() {
			wait_begin();
			var AJAX = new ajax_form_post_support('form_id', '/pages/terminal_actions_server.php', sent_request_res);
			AJAX.send_form_reqest();

			function sent_request_res(res) {
				wait_end();
				if (res != '!') {
					show_dlg_alert(res, 0);
					return;
				}
				show_dlg_alert("Сохранено", 1500);
				setTimeout("location.reload(true)", 1500);
			}
		}
		var lastID = <?= $lastID ?>;

		function add_new_line() {
			var tbl = document.getElementById('main_table');
			var currow = tbl.rows.length;
			tbl.insertRow(currow);
			tbl.rows[currow].insertCell(0);
			tbl.rows[currow].cells[0].style.backgroundColor = "#eeeeee";
			tbl.rows[currow].cells[0].innerHTML = "<input type='hidden' value='" + lastID + "' name='id[" + lastID + "]'>" + lastID + "";
			tbl.rows[currow].insertCell(1);
			tbl.rows[currow].cells[1].innerHTML = "<input class='table_input' type='text' name='decr[" + lastID + "]'  value='' >";
			tbl.rows[currow].insertCell(2);
			tbl.rows[currow].cells[2].innerHTML = "<input class='table_input' type='text' name='val[" + lastID + "]' <? echo only_numbers_and_sybm_input(); ?>  value='' >";
			tbl.rows[currow].insertCell(3);
			tbl.rows[currow].cells[3].innerHTML = "<input class='table_input' type='text' name='target[" + lastID + "]' <? echo only_numbers_and_sybm_input(); ?>  value='' >";

			lastID++;
		}

		function set_sel(tr) {
			var nodes = document.getElementsByClassName('highlighting'); //alert(nodes.length);
			for (var i = 0; i < nodes.length; i++) {
				nodes[i].style.border = "solid 1px #000000";
			}
			tr.style.border = "solid 2px #000000";
		}
	</script>

</div>
<br><br><br><br><br>
</body>

</html>
<?
/*       
http://go/  связь с GO <?include($_SERVER["DOCUMENT_ROOT"]."/common/linking_address.txt");?>
ПУЛЬТ связи с Beast
///////////////////// главная Пульта test ///////////////
*/

$page_id = 0;
$title = "Пульт связи с Beast";
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");

include_once($_SERVER['DOCUMENT_ROOT'] . "/common/spoiler.php");
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/alert_confirm.php");
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/alert2_dlg.php");

// набор инструментов (загрузка и сохранение памяти Beast)
include_once($_SERVER["DOCUMENT_ROOT"] . "/tools/tools.php");

///// стадии развития 
$stages = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/stages.txt");
$stages = trim($stages);

echo "<div class='main_page_div' style=''>";

echo "<div style='position:absolute;top:-30px;left:300px;'>
<span id='link_warning_id' style=''></span>

<span style='padding-left:40px;'>Пульс:</span> 
<div id='puls_id' class='puls_passive' style='position:absolute;top:-2px;right:-30px;'  title='Нормальный пульс'></div>

<div id='dander_warn_id' style='position:absolute;top:2px;right:-280px;color:red;font-size:17px;display:none;'  title='Повреждения критические!'><nobr><b>Best скоро умрет!</b></nobr></div>

</div>";

echo "<div id='common_status_id' type='button' style='position:absolute;top:-34px;left:560px;
border-radius: 7px;padding:4px;padding-left:8px;padding-right:8px;'></div>";

// общий движок связи с Beast
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/linking.php");

?>
<input id="bot_switcher" type='button' style="position:absolute;top:-30px;right:10px;display:block" value='&nbsp;Включить Beast&nbsp;' onClick="bot_switcher()">
<div id="close_note_id" style="position:absolute;top:-47px;right:10px;display:none;">Корректное выключение сохраняет текущую память!</div>
<input id="bot_switcher2" type='button' title='Некорректное выключение (память может не сохраниться), лучше пользоваться Выключить Beast справа-сверху шестеренка.' style="position:absolute;top:-30px;right:10px;display:none;color:red;" value='&nbsp;Выключить Beast&nbsp;' onClick="bot_switcher()">

<?
// возраст:
echo "<div id='life_time_id' style='position:absolute;top:0px;right:0px;'></div>";

echo "<div style='position:absolute;top:-66px;right:10px;cursor:pointer;color:blue;' onClick='open_anotjer_win(`/pages/main_help.htm`)'><b>Как использовать Пульт</b></div>";

/// блок включает то, что показывается при коннекте с Beast
echo "<div id='linking_block_div' style='display:;'>";

//задатчики жизненных параметров
include_once($_SERVER["DOCUMENT_ROOT"] . "/pult_gomeo.php");

// Индикация активных базовых контекстов
include_once($_SERVER['DOCUMENT_ROOT'] . "/pult_base_contexts.php");

// Диалог общения с Beast
include_once($_SERVER['DOCUMENT_ROOT'] . "/pult_Bot_dialog.php");

include_once($_SERVER['DOCUMENT_ROOT'] . "/pult_actions.php");

// показ дейстивий и фраз от Beast
include_once($_SERVER['DOCUMENT_ROOT'] . "/show_bot_actions.php");

// консоль событий
include_once($_SERVER['DOCUMENT_ROOT'] . "/pult_consol.php");

echo "</div>";

?>
<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script Language="JavaScript" src="/ajax/ajax_post.js"></script>
<script>
	//show_dlg_alert2("qwqweqeqwe",0);
	<?
	echo "var stages_dev='" . $stages . "';";
	?>

	// адрес для обращения по ajax к ГО-серверу
	var linking_address = '<? include($_SERVER["DOCUMENT_ROOT"] . "/common/linking_address.txt"); ?>';
	setTimeout("get_bot_connect()", 1000);

	var exists_connect = 0; // 1 - есть ответ Beast

	// чтобы запрос прошел нужно заполнить bot_params какой-то спецификой запроса НАЧИНАЯ С &
	var bot_params = "";
	var del_tymer = 0;
	var t_count = 0; // просто счетчик пульсов
	var isBotReady = 0;
	var main_puls_timer = 0;
	var not_allow_get_gomeo = 0; // 1-блокировка изменений при установки новых значений

	////////////////////////// постоянный пульс Пульта 1 сек
	function get_bot_connect() {
		clearTimeout(main_puls_timer);
		main_puls_timer = setTimeout("get_bot_connect()", 1000); // постоянно опрашивать

		// show_dlg_alert(t_count,500);

		// если был коннект, то снова опрашивать, если exists_connect==0 - значит коннект прекращен
		// если сработает after_answer_server() то exists_connect=1;
		if (exists_connect)
			del_tymer = setTimeout("check_bot_connect()", 5000);

		//if(not_allow_get_gomeo)alert("!!!!!!!!!!!!");

		// делаем запрос если нет других текущих запросов (is_active_connect в /common/linking.php)
		var req_bot_answer = 0; // alert(new_gomeo_par_id);
		// этот запрос на состояние бота должен быть всегда, и когда он есть, остальные ниже блокируются на этом пульсе.   
		if (t_count % 2 == 0 && !not_allow_get_gomeo) // опрос состояния Beast раз в 2 сек для /pult_gomeo.php
		{
			req_bot_answer = 1;
			bot_params = "get_params=1";
			bot_contact(bot_params, sent_get_params);
			bot_params = ""; //alert("!!!1");
			//return;
		}

		if (new_gomeo_par_id > 0) {
			//	alert(new_gomeo_par_id);
			bot_params = "set_params=" + new_gomeo_par_id + "&params_val=" + cur_gomeo_val;
			new_gomeo_par_id = 0;
			req_bot_answer = 1;
			bot_contact(bot_params, bot_answer);
			bot_params = ""; //alert("!!!2222");
			//not_allow_get_gomeo=0;
			//return;
		}

		// молотит вхолостую чтобы определять есть ли коннект
		if (!req_bot_answer) { //alert("!!!");
			bot_contact(bot_params, bot_answer);
		}
		bot_params = "";
		t_count++;
	}

	// готовность к новому запросу после ответа Beast
	function after_answer_server() { //alert("!!!");
		clearTimeout(del_tymer);
		exists_connect = 1;
		warning_connect();
		var type = 1;
		pulsing(type); //  show_dlg_alert(t_count,500);
	}
	// ответ Beast при холостом отпросе
	function bot_answer(res) {
		// alert(res);
		//set_Bot_Ready(0);
		if (res == "READY") {
			isBotReady = 1;
			set_Bot_Ready(1);
		}
		// инфа для консоли
		if (res.indexOf("CONSOL:") == 0) { //alert(res);
			set_consol(res.substr(7));
		}
		// инфа о действиях Beast
		if (res.indexOf("ACTION:") == 0) { //alert(res);
			new_bot_action(res.substr(7));
		}
		after_answer_server();
	}

	function check_bot_connect() // не ответил за 5 тактов
	{
		tools_show(0);
		exists_connect = 0;
		warning_connect();
	}

	function warning_connect() {
		if (exists_connect) {
			document.getElementById('link_warning_id').innerHTML = "<span style='color:green;'>Есть связь с Beastм</span>";
			document.getElementById('linking_block_div').style.display = "block";
			document.getElementById('bot_switcher').style.display = "none";
			document.getElementById('bot_switcher2').style.display = "";
			document.getElementById('close_note_id').style.display = "block";
		} else {
			document.getElementById('link_warning_id').innerHTML = "<span style='color:red;'>Нет связи с Beast</span>";
			document.getElementById('linking_block_div').style.display = "none";
			document.getElementById('bot_switcher').style.display = "";
			document.getElementById('bot_switcher2').style.display = "none";
			document.getElementById('close_note_id').style.display = "none";
		}
	}

	function pulsing(type) {
		var cn = "puls_active"; //  show_dlg_alert(t_count,500);
		if (type == 2)
			cn = "puls_stress";
		if (is_stressing)
			cn = "puls_stress";
		if (IsBeastDeath) {
			cn = "puls_death";
			document.getElementById('puls_id').className = cn;
			bot_closing(); //bot_switcher();// выключить 
			return;
		}
		// show_dlg_alert(cn,500);
		document.getElementById('puls_id').className = cn;
		setTimeout("pulsing_end()", 200);
	}

	function pulsing_end() {
		document.getElementById('puls_id').className = "puls_passive";
	}

	function bot_switcher() {
		if (document.getElementById('bot_switcher2').style.display == "none") {
			if (!confirm("Запустить исполняемый файл Beast?"))
				return;
			var server = "/run.php";
			var AJAX = new ajax_support(server, sent_run_answer);
			AJAX.send_reqest();

			// не ожидая sent_run_answer т.к. /run.php зависает при запуске файла 
			document.getElementById('bot_switcher').style.display = "none";
			document.getElementById('close_note_id').style.display = "block";
			document.getElementById('bot_switcher2').style.display = ""; //document.getElementById('bot_switcher2').value="&nbsp;Выключить Beast&nbsp;";
			show_dlg_alert("Beast включен.", 1);

			function sent_run_answer(res) {

			}
		} else {
			// сохранить память в файлах
			var AJAX = new ajax_support(linking_address + "?save_all_memory=1", sent_info_1);
			AJAX.send_reqest();

			function sent_info_1(res) {
				if (res != "yes") {
					show_dlg_alert("Не удалось сохранить память Beast. Выключение отменено.", 0);
					bot_closing();
					return;
				}
				var server = "/kill.php";
				var AJAX = new ajax_support(server, sent_end_answer);
				AJAX.send_reqest();

				function sent_end_answer(res) {
					document.getElementById('bot_switcher2').style.display = "none";
					document.getElementById('bot_switcher').style.display = "";
					show_dlg_alert("Beast выключен.", 1);
					set_Bot_Ready(0);
				}
			}
		}
	}

	function set_Bot_Ready(ready) {
		//alert(ready);
		if (ready == 0) {
			//При потере коннекта сбрасывать таймер 10сек удержания Акции
			clearTimeout(actionTimerID);
			desactivationAll();

			document.getElementById('about_bot_ready').innerHTML = "Beast еще не пришел в себя, нужно немного подождать";
			document.getElementById('input_id').disabled = true;
			document.getElementById('input_button_id').disabled = true;
			document.getElementById('stadia_warn').innerHTML = "";

			tools_show(0)
		} else // ВКЛЮЧЕН
		{
			document.getElementById('about_bot_ready').innerHTML = " "; //alert(typeof(stages_dev));
			if (stages_dev != '0') {
				document.getElementById('input_id').disabled = false;
				document.getElementById('input_button_id').disabled = false;
			} else {
				document.getElementById('stadia_warn').innerHTML = "Нулевая стадия - бессловестная.";
			}
			tools_show(1)
		}
	}
	// сохранение текущей памяти по Ctrl+S
	var is_press_strl = 0;
	document.onkeydown = function(event) {
		var kCode = window.event ? window.event.keyCode : (event.keyCode ? event.keyCode : (event.which ? event.which : null))

		//alert(kCode);
		if (kCode == 17) // ctrl
			is_press_strl = 1;

		if (is_press_strl) {
			if (kCode == 83) {
				event.preventDefault();
				//alert("!!!!! ");
				save_current_memory();
				is_press_strl = 0;
				return false;
			}
		}
	}

	// контроль события критического повреждения (из /pult_gomeo.php function sent_get_params(res))
	var IsBeastDeath = false;
	var is_stressing = false;

	function check_dander_warn_event(val) {
		//!!! работает pulsing(type) 	
		//alert(val);
		if (val >= 80) // начать менять фон
		{
			is_stressing = true;
			var color = "#FFC9C9";
			if (val >= 85) {
				color = "#CC818B";
			}
			if (val >= 90) {
				color = "#CC818B";
			}
			if (val >= 95) {
				color = "#A34747";
			}
			document.body.style.backgroundColor = color;
			document.getElementById('dander_warn_id').style.display = "block";
			document.getElementById('dander_warn_id').innerHTML = "<nobr><b>Best скоро умрет!</b></nobr>";
		}

		if (val < 80) {
			is_stressing = false
			document.body.style.backgroundColor = "#ffffff";
			document.getElementById('dander_warn_id').style.display = "none";
		}
	}

	function is_Beast_Death() {
		IsBeastDeath = true;
		document.body.style.backgroundColor = "#000000";
		document.getElementById('dander_warn_id').style.display = "block";
		document.getElementById('dander_warn_id').innerHTML = "<nobr><b>Beast умер.</b></nobr>";
	}

	var oldActipnStr = "";
	var stopReflexCreate = 0;

	function dialog_no_reflex(conditions, ignor) {
		if (stopReflexCreate) {
			return;
		}
		if (oldActipnStr == conditions) // не повторять идентичное
		{
			return;
		}
		oldActipnStr = conditions;
		//alert(conditions);

		var reason = "Для данных условий нет безусловного рефлекса";
		if (ignor) {
			var reason = "Для данных условий есть только рефлекс игнорирования";
		}

		show_dlg_alert2("<br><span style='font-weight:normal;'>" + reason + "<br>(редактор http://go/pages/reflexes.php)</span><br><br><span onClick='choose_actions(`" + conditions + "`)' style='cursor:pointer;color:blue;'>Создать подходящий рефлекс</span><br><br><span onClick='stop_reflex_create()' style='cursor:pointer;color:blue;'>Больше не показывать этот диалог</span>", 2);
	}

	function choose_actions(conditions) {
		//alert(conditions);
		var AJAX = new ajax_support("/lib/get_action_choose.php?id=0", sent_act_info);
		AJAX.send_reqest();

		function sent_act_info(res) {
			//alert(res);
			show_dlg_alert2("<br><span style='font-weight:normal;'>Выберите до 4-х действий рефлекса:<br>(Lkz dsltktybq используйте Ctrl+клик и Shift+клик)" + res + "<br><input type='button' value='Создать рефлекс' onClick='create_reflex(`" + conditions + "`)'>", 2);
		}
	}

	function create_reflex(conditions) {
		var aStr = ""; // <option class='a_option'
		var combo = document.getElementById('actions_combo');
		var len = combo.options.length;
		for (var n = 0; n < len; n++) {
			if (combo.options[n].selected == true) {
				if (aStr.length > 0)
					aStr += ",";
				aStr += combo.options[n].id;
			}
		}
		var qw = "/lib/create_reflex.php?aStr=" + aStr + "&conditions=" + conditions;
		//alert(qw);return;
		var AJAX = new ajax_support(qw, sent_cr_info);
		AJAX.send_reqest();

		function sent_cr_info(res) {
			if (res != '!') {
				show_dlg_alert2("Ошибка: " + res, 2);
				return;
			}
			//alert(res);
			show_dlg_alert2("Рефлекс создан. Нужно выключить и включить Beast чтобы были восприняты новые рефлексы.", 2);
		}
	}

	function stop_reflex_create() {
		stopReflexCreate = 1;
		show_dlg_alert2("Чтобы снова появлялись диалоги о рефлексах нужно нажать F5 (обновить страницу Пульта).", 0);
	}

	////////// редактировать действия рефлекса
	function edit_b_reflex(id) {
		//alert(conditions);
		var AJAX = new ajax_support("/lib/get_action_choose.php?id=" + id, sent_act_info);
		AJAX.send_reqest();

		function sent_act_info(res) {
			//alert(res);
			show_dlg_alert2("<br><span style='font-weight:normal;'>Выберите до 4-х действий рефлекса ID=" + id + ":<br>(Lkz dsltktybq используйте Ctrl+клик и Shift+клик)" + res + "<br><input type='button' value='Изменить действия рефлекса' onClick='correct_reflex(" + id + ")'>", 2);
		}
	}

	function correct_reflex(id) {
		var aStr = ""; // <option class='a_option'
		var combo = document.getElementById('actions_combo');
		var len = combo.options.length;
		for (var n = 0; n < len; n++) {
			if (combo.options[n].selected == true) {
				if (aStr.length > 0)
					aStr += ",";
				aStr += combo.options[n].id;
			}
		}
		var qw = "/lib/correct_reflex.php?id=" + id + "&aStr=" + aStr;
		//alert(qw);return;
		var AJAX = new ajax_support(qw, sent_cr_info);
		AJAX.send_reqest();

		function sent_cr_info(res) {
			if (res != '!') {
				show_dlg_alert2("Ошибка: " + res, 2);
				return;
			}
			//alert(res);
			show_dlg_alert2("Рефлекс скорректирован.<br><br>Нужно выключить и включить Beast чтобы были восприняты изменения.<br>Это можно сделать после всех коррекций.", 2);
		}
	}
	//new_bot_action("1|БЕССМЫСЛЕННЫЙ безусловный рефлекс<br>Предлогает поиграть");
</script>
</div>
</body>

</html>
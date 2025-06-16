<?
/*       
http://go/  связь с GO include($_SERVER["DOCUMENT_ROOT"]."/common/linking_address.txt");
ПУЛЬТ связи с Beast
///////////////////// главная Пульта test ///////////////
*/
$reflex_quick_working=0;
if(isset($_GET['reflex_quick_working'])&&$_GET['reflex_quick_working']==1)
$reflex_quick_working=1;// режим быстрой набивки б.рефлексов в зависимости от выставляемых условий без коннекта с Beast.

$page_id = 0;
$title = "Пульт связи с Beast";
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/header.php");

// инициация записи истории общения
$init_history=true;
include_once($_SERVER['DOCUMENT_ROOT'] . "/pages/history.php");//  exit("! $curHistoryFile");
$init_history=false;

include_once($_SERVER['DOCUMENT_ROOT'] . "/common/spoiler.php");
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/alert_confirm.php");
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/alert2_dlg.php");
include_once($_SERVER['DOCUMENT_ROOT']."/common/alert_control.php");

// набор инструментов (загрузка и сохранение памяти Beast)
include_once($_SERVER["DOCUMENT_ROOT"] . "/tools/tools.php");

///// стадии развития 
$stages = read_file($_SERVER["DOCUMENT_ROOT"] . "/memory_reflex/stages.txt");
$stages = trim($stages);

echo '<div id="helper_pult_id" style="position:fixed;z-index:1000;top:0px;right:0px;
background-color:#FFFFCC; height:;
padding:6px;
box-shadow: 8px 8px 8px 0px rgba(122,122,122,0.3);
border:solid 1px #81853D; border-radius: 7px;
display:none;"></div>';

echo "<div class='main_page_div' style=''>";

if($reflex_quick_working)
{
echo "<div style='position:absolute;top:-30px;left:300px;font=size:19px;color:#00C200;'><b>Это - режим быстрой набивки безусловных рефлексов</b></div>";
}
else
{
echo "<div style='position:absolute;top:-30px;left:300px;'>
<span id='link_warning_id' style=''></span>

<span style='padding-left:40px;'>Пульс:</span> 
<div id='puls_id' class='puls_passive' style='position:absolute;top:-2px;right:-30px;'  title='Нормальный пульс'></div>

<div id='dander_warn_id' style='position:absolute;top:2px;right:-280px;color:red;font-size:17px;display:none;'  title='Повреждения критические!'><nobr><b>Best скоро умрет!</b></nobr></div>

</div>";

echo "<div id='common_status_id' type='button' style='position:absolute;top:-34px;left:560px;
border-radius: 7px;padding:4px;padding-left:8px;padding-right:8px;'></div>

<div id='common_status_exit_id' class='alert_exit' style='position:absolute;z-index:10;top:-34px;left:622px;display:none;' title='Снять состояния Хорошо и Плохо.\n\nМожно использовать клавишу ESC.' onClick='end_whell_bad();'><span >&#10006;</span></div>


";

// общий движок связи с Beast
include_once($_SERVER['DOCUMENT_ROOT'] . "/common/linking.php");

?>
<input id="bot_switcher" type='button' style="position:absolute;top:-30px;right:10px;display:block" value='&nbsp;Включить Beast&nbsp;' onClick="bot_switcher()">
<div id="close_note_id" style="position:absolute;top:-47px;right:10px;display:none;">Корректное выключение сохраняет текущую память!</div>
<input id="bot_switcher2" type='button' title='Некорректное выключение (память может не сохраниться), лучше пользоваться Выключить Beast справа-сверху шестеренка.' style="position:absolute;top:-30px;right:10px;display:none;color:red;" value='&nbsp;Выключить Beast&nbsp;' onClick="bot_switcher()">

<?
// возраст:
echo "<div style='position:absolute;top:0px;right:0px;'>";
echo "<span id='life_time_id' style=''></span>";

echo '&nbsp;<img src="/img/edit.png" style="cursor:pointer;" title="Установить возраст" onClick="set_life_time(this)">';

echo "</div>";
}


echo "<div style='position:absolute;top:-66px;right:10px;cursor:pointer;color:blue;' onClick='open_anotjer_win(`https://scorcher.ru/beast/main_help.htm`)'><b>Как использовать Пульт</b></div>";







/// блок включает то, что показывается при коннекте с Beast
echo "<div id='linking_block_div' style='display:;'>";

//задатчики жизненных параметров
include_once($_SERVER["DOCUMENT_ROOT"] . "/pult_gomeo.php");

// Индикация активных базовых контекстов
include_once($_SERVER['DOCUMENT_ROOT'] . "/pult_base_contexts.php");




///////////////////////  НАБИВКА РЕФЛЕКСОВ
if($reflex_quick_working)
{
echo "<div style='margin-top:10px;background-color:#eeeeee;padding:10px;border:solid 1px #8A3CA4;border-radius: 7px;box-shadow: 8px 8px 8px 0px rgba(122,122,122,0.3);'><b>Это - режим быстрой набивки безусловных рефлексов в зависимости от выставляемых условий без коннекта с Beast.</b>";

echo "<br>Пояснение:
<br>
<br>
<br>
";

echo "<br><a href='/pult.php'>Переключить нормальный режим Пульта</a>";

echo "</div>";
}
else
{

// Диалог общения с Beast
include_once($_SERVER['DOCUMENT_ROOT'] . "/pult_Bot_dialog.php");
}

include_once($_SERVER['DOCUMENT_ROOT'] . "/pult_actions.php");

// показ дейстивий и фраз от Beast
include_once($_SERVER['DOCUMENT_ROOT'] . "/show_bot_actions.php");

if(!$reflex_quick_working)
{
echo "<div id='extend_info_id' style='height:25px;padding-left:10px;margin-top:10px;'></div>";

// консоль событий
include_once($_SERVER['DOCUMENT_ROOT'] . "/pult_consol.php");
}


echo "</div>";

?>
<script Language="JavaScript" src="/ajax/ajax.js"></script>
<script Language="JavaScript" src="/ajax/ajax_post.js"></script>
<script>

function addInfoToHistory(type,str)
{
// после &nbsp; все обрезает из-за '&' (разделитель POST параметров), маскируем:
str=str.replace(/&/g,'[x1]');

//	alert(type);
//str="<b>!!! БЕССМЫСЛЕННЫЙ безусловный рефлекс:</b><br>&nbsp;Предлагает поиграть, &nbsp;улыбается<br>";
var eInfo=document.getElementById('extend_info_id').innerHTML;
eInfo=eInfo.substr(0,eInfo.indexOf("("));

var bS=document.getElementById('common_status_id').innerHTML;

var conStr=document.getElementById('contect_list_id').innerHTML;

var context="Состояние: <b>"+bS+"</b></span>. Базовые контексты: <b>"+conStr+"</b></span> "+eInfo+"</span>\r\n";

param="histoty_file=<?=$curHistoryFile?>&type="+type+"&context="+context+"&newInfoHist="+str; 
// alert(param);
var AJAX = new ajax_post_support('/pages/history.php',param,sent_history_mess,1);
AJAX.send_reqest();
function sent_history_mess(res)
{
//alert(res);

}
}

// адрес для обращения по ajax к ГО-серверу
var linking_address = '<? include($_SERVER["DOCUMENT_ROOT"] . "/common/linking_address.txt"); ?>';


window.addEventListener('beforeunload', function (event) {
/* нет способа отличить просто обновление страницы от выхода по крестику....
но можно сделать кнопку для обновления по которой будет идти запрет на выключение сервера.
*/

if(is_menu_click)// переключалась вкладка
{
	is_menu_click=false;
	return;
}

  // выполнение действий перед закрытием окна
  //event.preventDefault();
  //event.returnValue = '';
var testSTR=document.referrer+" | "+window.location.href;
var AJAX = new ajax_support(linking_address+'?beforeunload=1&testSTR='+testSTR,sent_brauser_close);
AJAX.send_reqest();
function sent_brauser_close(res)
{
  event.preventDefault();
  event.returnValue = '';
}
});


	//show_dlg_alert2("qwqweqeqwe",0);
	<?
	echo "var stages_dev='" . $stages . "';";
	?>

	

/////////////////////////////////////////////
/* Синхронизация пульса с ГО:
Как только пришел ответ из ГО sent_sincronisation(res) происходит новый запуск через 1 сек.
При этом в main.go происходит подмена автогенератора синхрозапуском: brain.SincroTic()
Так что пуль на Пульте и в ГО синхронизирован.
Времени 5 сек типа должно хватить, чтобы в ГО завершились процессы данного Пульса.
*/
var main_puls_timer = 0;
var main_scan_timer = 0;
// поехали...
main_puls_timer = setTimeout("sincronisationWithGo()", 1000);
function sincronisationWithGo()
{
		clearTimeout(main_puls_timer);
		clearTimeout(main_scan_timer);
// все время опрашивает, но в раб.режиме успевает очиститься clearTimeout(main_puls_timer);
		main_scan_timer = setTimeout("sincronisationWithGo()", 5000);
		//////////////
if(is_sent_params==0)// заглушка чтобы не посылать запросы пока не обработается текущий
{
bot_contact("get_params=1", sent_sincronisation);
}
function sent_sincronisation(res)
{
//   alert("!!!!!!!!!!!! 22222");
// для синхронизации автоопроса и пульса ГО
		clearTimeout(main_puls_timer);
		clearTimeout(main_scan_timer);
		main_puls_timer = setTimeout("sincronisationWithGo()", 1000);
}
// вызов при Включено - раз в секунду, при выключено - раз в 5 сек.
		get_bot_connect();
}
//////////////////////////////////////////////


	var exists_connect = 0; // 1 - есть ответ Beast

	// чтобы запрос прошел нужно заполнить bot_params какой-то спецификой запроса НАЧИНАЯ С &
	var bot_params = "";
	var del_tymer = 0;
	var t_count = 0; // просто счетчик пульсов
	var isBotReady = 0;
	
	var not_allow_get_gomeo = 0; // 1-блокировка изменений при установки новых значений
	////////////////////////// постоянный пульс Пульта 1 сек
function get_bot_connect() {
			// show_dlg_alert(t_count,500);

// Детектор: если был коннект, то снова опрашивать check_bot_connect(), если exists_connect==0 - значит коннект прекращен
		// если сработает after_answer_server() то exists_connect=1;
		if (exists_connect)
			del_tymer = setTimeout("check_bot_connect()", 5000);

		//if(not_allow_get_gomeo)alert("!!!!!!!!!!!!");

		// делаем запрос если нет других текущих запросов (is_active_connect в /common/linking.php)
		var req_bot_answer = 0; // alert(new_gomeo_par_id);
		// этот запрос на состояние бота должен быть всегда, и когда он есть, остальные ниже блокируются на этом пульсе.   
		if (t_count % 2 == 0 && !not_allow_get_gomeo) // опрос состояния Beast раз в 2 сек для /pult_gomeo.php
		{   //alert("!!!!!!!!!!!! 0000");
			req_bot_answer = 1;
			
			if(is_sent_params==0)// заглушка чтобы не посылать запросы пока не обработается текущий
			{
//is_sent_params=1;  // может быть нужно при тестировании обеспечивать толь 1 вызов "get_params=1"
			bot_params = "get_params=1";  
			bot_contact(bot_params, sent_get_params);
			bot_params = ""; //alert("!!!1");
			}
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
		//blockingAutoScan=0;
		tools_show(0);
		exists_connect = 0;
		warning_connect();
	}

var runningBeast=false;
var stopingBeast=false;
	function warning_connect() {
		if (exists_connect) {  
		if(!runningBeast){ // alert("Есть связь");
			runningBeast=true;stopingBeast=false;
			addInfoToHistory(4,"xxxx"); // в /index.php
			}
			document.getElementById('link_warning_id').innerHTML = "<span style='color:green;'>Есть связь с Beast</span>";
			document.getElementById('linking_block_div').style.display = "block";
			document.getElementById('bot_switcher').style.display = "none";
			document.getElementById('bot_switcher2').style.display = "";
			document.getElementById('close_note_id').style.display = "block";
		} else { 
		if(!stopingBeast){ // alert("Нет связи");
			stopingBeast=true;runningBeast=false;
			addInfoToHistory(5,"xxxx"); // в /index.php
			}
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

function bot_switcher() { //  alert("bot_switcher");
		if (document.getElementById('bot_switcher2').style.display == "none") {
			if (stages_dev=='0'){
				show_dlg_alert("Активация возможна только с 1 стадии. Включение отменено.", 0);
				return;				
			}
//if (!confirm("Запустить исполняемый файл Beast?"))
//return;
show_dlg_alert("Включаем...", 1500);
			var server = "/run.php";    
			var AJAX = new ajax_support(server, sent_run_answer);
			AJAX.send_reqest();
			function sent_run_answer(res) {  // alert(res);
// Довольно долго запускается файл ГО, так что не нужно ждать окончания, а сразу
// location.reload(true);
//show_dlg_alert("Beast включился.", 1500);
			}
/* нужно перегрузить страницу, иначе после загрузки сохраненного архива и Включения не появляется страница (как при отсуствии связи). Пока не исследовал почему это, просто поставил обновление страницы:
*/
			location.reload(true);
/*
			// не ожидая sent_run_answer т.к. /run.php зависает при запуске файла 
			document.getElementById('bot_switcher').style.display = "none";
			document.getElementById('close_note_id').style.display = "block";
			document.getElementById('bot_switcher2').style.display = ""; //document.getElementById('bot_switcher2').value="&nbsp;Выключить Beast&nbsp;";
			show_dlg_alert("Beast включен.", 1);
*/
		} else {
			// сохранить память в файлах
			var AJAX = new ajax_support(linking_address + "?save_all_memory=1", sent_save_answer);
			AJAX.send_reqest();
			function sent_save_answer(res) {  //alert(res);
				if (res != "yes") {
					show_dlg_alert("Не удалось сохранить память Beast. Выключение отменено.", 0);
					bot_closing();
					return;
				}
				/*
				var server = "/kill.php";
				var AJAX = new ajax_support(server, sent_end_answer);
				AJAX.send_reqest();
				function sent_end_answer(res) {  //alert(res);
					close_end()
				}*/
				bot_closing();
			}
		}
	}
function close_end()
{
document.getElementById('bot_switcher2').style.display = "none";
document.getElementById('bot_switcher').style.display = "";
show_dlg_alert("Beast выключен. Память сохранена.", 1);
set_Bot_Ready(0);
}

var startPsichicNow=0;// 1- психика уже готова
	function set_Bot_Ready(ready) {
		//alert(ready);
		if (ready == 0) {
			//При потере коннекта сбрасывать таймер 10сек удержания Акции
			//clearTimeout(actionTimerID);
			desactivationAll();

//			document.getElementById('about_bot_ready').innerHTML = "Beast еще не пришел в себя, нужно немного подождать";
			document.getElementById('input_id').disabled = true;
			document.getElementById('input_button_id').disabled = true;
			document.getElementById('reset_button_id').disabled = true;
			document.getElementById('stadia_warn').innerHTML = "";

			tools_show(0)
		} 
		else // ВКЛЮЧЕН
		{
			if (stages_dev != '0') {
				document.getElementById('input_id').disabled = false;
				document.getElementById('input_button_id').disabled = false;
				if (stages_dev=='4' || stages_dev=='5') {
					document.getElementById('reset_button_id').disabled = false;
				} else {
					document.getElementById('reset_button_id').disabled = true;
				}
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
	function dialog_no_reflex(conditions, ful_condition_str,ignor) {  // return;
		if (stopReflexCreate) {
			return;
		}
		if (oldActipnStr == conditions) // не повторять идентичное
		{
			return;
		}
		oldActipnStr = conditions;
		// alert(conditions);
		var reason = "Для данных условий:<br>";

		reason +="<div style='font-weight:normal;text-align:left;font-size:12px;background-color:#D7DAFF;padding:8px;'>"+ful_condition_str+"</div>";

		reason += "нет безусловного рефлекса";
		if (ignor) {
			reason += "есть только рефлекс игнорирования";
		}

		show_dlg_alert2("<br><span style='font-weight:normal;'>" + reason + "<br>(редактор http://go/pages/reflexes.php)</span><br><br><span onClick='choose_actions(`" + conditions + "`)' style='cursor:pointer;color:blue;'>Создать подходящий рефлекс</span><br><br><span onClick='stop_reflex_create()' style='cursor:pointer;color:blue;'>Больше не показывать этот диалог</span>", 2);
	}

function choose_actions(conditions) {
		//alert(conditions);
/*		
		var AJAX = new ajax_support("/lib/get_action_choose.php?id=0", sent_act_info);
		AJAX.send_reqest();

		function sent_act_info(res) {
			//alert(res);
			show_dlg_alert2("<br><span style='font-weight:normal;'>Выберите до 4-х действий рефлекса:<br>(Для выделений используйте Ctrl+клик и Shift+клик)" + res + "<br><input type='button' value='Создать рефлекс' onClick='create_reflex(`" + conditions + "`)'>", 2);
		}
		
*/
//show_dlg_alert(nid,0);
event.stopPropagation();
		var AJAX = new ajax_support("/lib/get_actions_list.php?selected=", sent_act_info);
		AJAX.send_reqest();

		function sent_act_info(res) { //alert(res);
			show_dlg_alert2("<br><span style='font-weight:normal;'>Выберите значения:<br>" + res + "<br><input type='button' value='Выбрать значения' onClick='create_reflex(`" + conditions + "`)'>", 2);
		}
		
}

function create_reflex(conditions) {  
/*	
		var aStr = ""; // <option class='a_option'
		var combo = document.getElementById('actions_combo');  // alert(res);
		var len = combo.options.length;
		for (var n = 0; n < len; n++) {
			if (combo.options[n].selected == true) {
				if (aStr.length > 0)
					aStr += ",";
				aStr += combo.options[n].id;
			}
		}
alert(aStr);
		document.getElementById("lev4_" + nid).value = aStr;

		end_dlg_alert2();
*/
var aStr = "";
var nodes = document.getElementsByClassName('chbx_identiser'); //alert(nodes.length);
for(var i=0; i<nodes.length; i++) 
{
if(nodes[i].checked)
	{
if(aStr.length > 0)
	aStr += ",";
aStr += nodes[i].value;
	}
}
//alert(aStr);
//document.getElementById("lev4_" + nid).value = aStr;
end_dlg_alert2();

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
			show_dlg_alert2("Рефлекс создан. Нужно выключить и включить Beast чтобы были восприняты новые рефлексы.<br><span style='color:blue;cursor:pointer;' onClick='reload_beast()'>Перезагрузить Beast</span> (~1,5 сек)", 2);
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
			show_dlg_alert2("<br><span style='font-weight:normal;'>Выберите до 4-х действий рефлекса ID=" + id + ":<br>" + res + "<br><input type='button' value='Изменить действия рефлекса' onClick='correct_reflex(" + id + ")'>", 2);
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

// выключить и снова включить исполняемый файл
function reload_beast()
{
end_dlg_alert();
end_dlg_alert2();
wait_begin();
var AJAX = new ajax_support("/kill.php", sent_reb1_answer);
AJAX.send_reqest();
//alert("1");
setTimeout("rebooting()",1000);// выждать завершения процессов
function sent_reb1_answer(res)
	{

	}
}
function rebooting()
{
//	alert("2");
var AJAX = new ajax_support("/run.php", sent_reb2_answer);
AJAX.send_reqest();
setTimeout("rebooting2()",500);
function sent_reb2_answer(res)
	{

	}
}
function rebooting2()
{
wait_end();  //alert("333");
show_dlg_alert("Beast перезагружен.",1500);
}

function set_life_time(parent)
{
var str="<div style='color:red;text-align:left;width:300px;'>При изменении возраста  эпизодическая память и доминанты будут удалены! Файл условных рефлексов будет обновлен.</div>";
str+="<b>Установить возраст:</b><br><table border=0>";
str+="<tr><td>Число лет: </td><td><input id='set_yeas' type='text' value='' class='control_input'></td></tr>";
str+="<tr><td>Число месяцев: </td><td><input id='set_month' type='text' value='' class='control_input'></td></tr>";
str+="<tr><td>Число дней: </td><td><input id='set_days' type='text' value='' class='control_input'></td></tr></table>";
str+="<input type='button' value='Установить возраст' onClick='set_life_time2()'>";
show_dlg_control(str,parent);
}
function set_life_time2()
{
var yeas=parseInt(document.getElementById('set_yeas').value);
var month=parseInt(document.getElementById('set_month').value);
var days=parseInt(document.getElementById('set_days').value);
//alert(yeas+" | "+month+" | "+days);
if (isNaN(yeas) && isNaN(month) && isNaN(days))	
{
show_dlg_alert("Нужно определить хотя бы одно из значений.<br>Можно задать только число дней.",0);
	return;
}
var server = "/lib/set_life_time.php?yeas="+yeas+"&month="+month+"&days="+days;    alert(server);
var AJAX = new ajax_support(server, sent_cliner_gomeo);
AJAX.send_reqest();
function sent_cliner_gomeo(res)
{
	if (res=="0")
	{
show_dlg_alert("Возраст не был установлен т.к. переданы пустые значения.",0);
		return;
	}
	var warn="";
	if (exists_connect)
		warn="<br><br>Beast выключается.";

	show_dlg_alert("Установлен новый возраст.<br>теперь все условные рефлексы - новые.<br>Удалена эпизодическая память.<br>Нужно пройти заново 4-ю ступень развития."+warn,0);

	if (exists_connect)
	{
	var server = "/kill.php";
		var AJAX = new ajax_support(server, sent_end_answer);
		AJAX.send_reqest();
		function sent_end_answer(res) {
			show_dlg_alert("Beast выключен.", 2000);
		}
	}

}
end_dlg_control();
}
</script>
</div>
<style>
.control_input
{
width:100%;
border:0;
outline:none;
text-align:left;
padding-left:4px;
padding-top:2px;
box-sizing:border-box;
}
</style>

</body>

</html>
<?
/*  Действия по отношению к Beast
 include_once($_SERVER['DOCUMENT_ROOT']."/pult_actions.php");
*/
$food_portion = '<select id="food_portion_id" title="Порция энергии в процентах от полного насыщения энергии." style="border:0;background-color:#E4FFEB;" onClick="event.stopPropagation();"> 
<option value="1">20%</option>
<option value="2" selected>50%</option>
<option value="3">80%</option>
</select>';
?>
<br>
<div id="action_block_id" style='position:relative;padding:5px;background-color:#ffffff;'>
	<div style='position:relative;'>
		<b>(Де)мотивирующие дйствия в ответ на действия Beast</b> (Пусковые стимулы):
		<a href='/pages/reflex_tree.php' style='position:absolute;top:0px;right:0px;'>Дерево рефлексов</a>
	</div>
	Можно набрать несколько действий подряд для отсылки общего Пускового стимула (но не стоит наборать более 3-х действий). Или <b>просто нажать на треугольник</b> чтобы это действие сразу отработало (когда не нужны несколько действий).<br>
<?
// рамка фона под мотивирующими кнопками Наказать-Поощрить
echo "<div style='position:absolute;z-index:0;
top:50px;left:125px;
width:115px;
height:65px;
background-color:#E8EBFF;cursor:pointer;
text-align:center;
' title='Сильное воздействие кнопок - для поощрения и наказания.' onClick='help_motivation()'><span style='cursor:pointer;font-size:12px;color:#2C288B;'  >Мотивировать</span></div>";

echo "<div style='position:absolute;z-index:0;
top:50px;left:245px;
width:130px;
height:65px;
background-color:#C2C5FF;cursor:pointer;
text-align:center;
' title='Предельно сильно воздействие - для разрешения и запрета.' onClick='help_motivation2()'><span style='font-size:12px;color:#2C288B;'  >Принудить</span></div>";


echo "<div style='position:absolute;z-index:0;
top:50px;left:380px;
width:115px;
height:65px;
background-color:#E8EBFF;cursor:pointer;
text-align:center;
' title='Значительное воздействие - для разрешения и запрета.' onClick='help_motivation3()'><span style='cursor:pointer;font-size:12px;color:#2C288B;'  >Эмпатировать</span></div>";


echo "<div id='teaching_mode_div' style='position:absolute;z-index:0; display:none;
top:50px;left:885px;
width:120px;
height:30px;
background-color:#C3FFB0;
box-shadow: 0px 0px 10px 10px rgba(180,255,174,0.8);
border-radius: 20px;
text-align:center;
' title='Включен учительский режим.' >
<span style='position:relative;top:-10px;'>режим</span>
<div id='teaching_mode_exit_id' class='alert_exit' style='position:absolute;z-index:10;top:-10px;left:100px;' title='Снять учительский режим' onClick='end_teaching_mode();'><span >&#10006;</span></div></div>";

// вокруг Успокоить
/*
echo "<div style='position:absolute;z-index:0;
top:110px;left:5px;
width:115px;
height:40px;
background-color:#E4FFD7;cursor:pointer;
text-align:center;
' title='Погашение удерживаемых режимов Хорошо и Плохо.' onClick='help_motivation4()'><span style='position:relative;top:25px;cursor:pointer;font-size:12px;color:green;'  >Нормализовать</span></div>";
*/

// вставка треугольничка отправки единственного действия
function setGo($id)
{
echo "<img style='position:absolute;top:50%;transform: translate(0, -50%);right:-4px;' src='/img/go.png' onClick='event.stopPropagation();single_action(".$id.")' title='Сразу совершить это одно действие'>";
}



include_once($_SERVER['DOCUMENT_ROOT'] . "/lib/actions_from_pult.php");
//var_dump($actionsFromPultArr);exit();
$trigg=$actionsFromPultArr;
?>
	<div style='position:relative;margin-top:5px;margin-top:10px;'>
	<div id="act_1" class='actions actions_red action_poz1' onClick="to_action(1)" title="<?=$trigg[1][1]?>" ><?=$trigg[1][0]?><?=setGo(1)?></div>
	<div id="act_3" class='actions actions_red action_poz1' onClick="to_action(3)" title="<?=$trigg[3][1]?>"><?=$trigg[3][0]?><?=setGo(3)?></div>
	
	<div id="act_10" class='actions actions_red action_poz3' onClick="to_action(10)" title="<?=$trigg[10][1]?>"><?=$trigg[10][0]?><?=setGo(10)?></div>
	<div id="act_12" class='actions actions_red action_poz1' onClick="to_action(12)" title="<?=$trigg[12][1]?>"><?=$trigg[12][0]?><?=setGo(12)?></div>
	<div id="act_15" class='actions actions_red action_poz3' onClick="to_action(15)" title="<?=$trigg[15][1]?>"><?=$trigg[15][0]?><?=setGo(15)?></div>
	<div id="act_5" class='actions actions_gray action_poz3' onClick="to_action(5)" title='<?=$trigg[5][1]?>'><?= $food_portion ?><?=$trigg[5][0]?><?=setGo(5)?></div>
	<div id="act_7" class='actions actions_gray action_poz1' onClick="to_action(7)" title="<?=$trigg[7][1]?>"><?=$trigg[7][0]?><?=setGo(7)?></div>
	<div id="act_8" class='actions actions_gray action_poz1' onClick="to_action(8)" title="Увеличение потребности учиться - задать контекст обучающего общения.">Поучить<?=setGo(8)?></div>
	<br><div style='font-size:5px;'>&nbsp;</div>
	<div id="act_2" class='actions actions_green action_poz1' onClick="to_action(2)" title="<?=$trigg[2][1]?>"><?=$trigg[2][0]?><?=setGo(2)?></div>
	
	<div id="act_4" class='actions actions_green action_poz2' onClick="to_action(4)" title="<?=$trigg[4][1]?>"><?=$trigg[4][0]?><?=setGo(4)?></div>
	<div id="act_11" class='actions actions_green action_poz3' onClick="to_action(11)" title="<?=$trigg[11][1]?>"><?=$trigg[11][0]?><?=setGo(11)?></div>
	<div id="act_13" class='actions actions_green action_poz1' onClick="to_action(13)" title="<?=$trigg[13][1]?>"><?=$trigg[13][0]?><?=setGo(13)?></div>

	<div id="act_14" class='actions actions_green action_poz3' onClick="to_action(14)" title="<?=$trigg[14][1]?>"><?=$trigg[14][0]?><?=setGo(14)?></div>
	<div id="act_9" class='actions actions_gray action_poz3' onClick="to_action(9)" title="<?=$trigg[9][1]?>"><?=$trigg[9][0]?><?=setGo(9)?></div>
	<br><div style='font-size:5px;'>&nbsp;</div>
	<div id="act_6" class='actions actions_blue action_poz1' onClick="to_action(6)" title="<?=$trigg[6][1]?>"><?=$trigg[6][0]?><?=setGo(6)?></div>
	<div id="act_16" class='actions actions_blue action_poz1' onClick="to_action(16)" title="<?=$trigg[16][1]?>"><?=$trigg[16][0]?><?=setGo(16)?></div>
	<div id="act_17" class='actions actions_blue action_poz1' onClick="to_action(17)" title="<?=$trigg[17][1]?>"><?=$trigg[17][0]?><?=setGo(17)?></div>

<div id="cliner_trigger_stimuls_id" class='actions' style="position:absolute;bottom:25px;right:0px;background-color:#eeeeee;color:grey;text-align:center;" onClick="desactivationAll()" title="Отменить выбранные сочетания пусковых стимулов.">Отменить выборку</div>


	<div id="sent_trigger_stimuls_id" class='actions' style="position:absolute;bottom:0px;right:0px;background-color:#eeeeee;color:grey;text-align:center;" onClick="sent_trigger_stimuls()" title="Послать выбранные сочетания пусковых стимулов.">Послать&nbsp;для&nbsp;Beast</div>
	</div>


</div>
<script>

	<? // массив антагонистов для JS:
	include_once($_SERVER['DOCUMENT_ROOT'] . "/lib/actions_from_pult.php");
	echo "var actionsFromPultAntagonistsArr = new Array();\r\n";
	foreach ($actionsFromPultAntagonistsArr as $k => $v) {
		echo "actionsFromPultAntagonistsArr[" . $k . "]=new Array();\r\n";
		foreach ($v as $a) {
			echo "actionsFromPultAntagonistsArr[" . $k . "].push(" . $a . ");\r\n";
		}
	}
	//exit("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!");
	?>

	function in_array(value, array) {
		if (typeof(array) != 'object')
			return false;
		for (var i = 0; i < array.length; i++) {
			if (value == array[i]) return true;
		}
		return false;
	}

///////////////////////////////// 
// список нажатых действий
var actionsArr = new Array();
var allow_sent_to_beast=0;// разрешается посылать 
function to_action(id) {


		// Не позволять включать антагонистов
		var antagonst = 0;
		// есть ли среди нажатых кнопок actionsArr антагонисты
		for (i = 0; i < actionsArr.length; i++) {
			//alert(actionsArr[i]);
			//alert(actionsFromPultAntagonistsArr[actionsArr[i]]);
			if (in_array(id, actionsFromPultAntagonistsArr[actionsArr[i]])) {
				antagonst = 1;
				break;
			}
		}
		if (antagonst == 1) {
			end_dlg_alert();
			end_dlg_alert2();
			show_dlg_alert("Уже действует антагонист.", 2000);
			return;
		}
		
		

		var trigBtn=document.getElementById("act_" + id); 
	if(trigBtn.className.indexOf('selButton')<0)
	{
	if(actionsArr.length==2)
	{
			end_dlg_alert();
			end_dlg_alert2();
			show_dlg_alert("Не более ДВУХ выбранных действий.", 2000);
			return;
	}
actionsArr.push(id);
trigBtn.className+=' selButton'; // рамка вокруг
	}
	else// убрать класс рамки
	{
actionsArr.splice(actionsArr.indexOf(id),1);
trigBtn.className=trigBtn.className.substr(0,trigBtn.className.indexOf(' selButton'));
	}

		//МЕНЯТЬ ФОН (розовый или голубой) БЛОКА ДЕЙСТВИЙ /pult_actions.php ПРИ НАЖАТИЯХ НА копки действий
		var AJAX = new ajax_support("/pages/gomeostaz_get_motivation.php?id=" + id, sent_motivation_action);
		AJAX.send_reqest();

		// action_block_id
		function sent_motivation_action(res) {
			if (res == "+")
				document.getElementById("action_block_id").style.backgroundColor = "#DDEBFF";
			if (res == "-")
				document.getElementById("action_block_id").style.backgroundColor = "#FFE4E1";
		}

	if(actionsArr.length>0)
	{
allow_sent_to_beast=1;
document.getElementById("sent_trigger_stimuls_id").style.outline="solid 1px #000000";
document.getElementById("sent_trigger_stimuls_id").style.color="#000000";
document.getElementById("cliner_trigger_stimuls_id").style.outline="solid 1px #000000";
document.getElementById("cliner_trigger_stimuls_id").style.color="#000000";
	}
	else
	{
allow_sent_to_beast=0;
document.getElementById("sent_trigger_stimuls_id").style.outline="solid 0px #000000";
document.getElementById("sent_trigger_stimuls_id").style.color="grey";
document.getElementById("cliner_trigger_stimuls_id").style.outline="solid 0px #000000";
document.getElementById("cliner_trigger_stimuls_id").style.color="grey";
	}

}
//////////////////////////////////////////////////////////
function sent_trigger_stimuls()
{ 
	if(!allow_sent_to_beast)
	{
		show_dlg_alert("Не выбрано ни одного действия.", 0);
		return;
	}
var triggers_str="";
// получить список пусковых стимулов (class='actions', но есть actionsArr)
for(i=0;i<actionsArr.length;i++)
	{
triggers_str+=actionsArr[i]+"|";//! нельзя разделять ; или ,
	}

	sending_trigg(triggers_str);
}
////////////////////////////////
function single_action(id)
{
sending_trigg(id);
}
///////////////////////////////////
function sending_trigg(triggers_str)
{
//alert(triggers_str);desactivationAll();return;
//alert(triggers_str);
end_dlg_bot_action(); //alert("!!!!");

//alert(triggers_str);
addInfoToHistory(3,""+triggers_str);

var linking_address = '<? include($_SERVER["DOCUMENT_ROOT"] . "/common/linking_address.txt"); ?>';
		var food_portion = document.getElementById("food_portion_id").selectedIndex + 1;
		//alert(food_portion);
//alert(linking_address + "?set_action=" + triggers_str);
		var AJAX = new ajax_support(linking_address + "?set_action=" + triggers_str + "&food_portion=" + food_portion, sent_action);
		AJAX.send_reqest();

		function sent_action(res) {		
			if(res.length>5)// уже готов ответ Beast в том же пульсе
			{ //alert(res);
// выдать его на Пульт
new_bot_action(res);
			}else{
			show_dlg_alert("Пусковой стимул принят Beast.", 2000);
			}
			desactivationAll();
		}
}
/////////////////////////////////////


//////////////////////////////////////////////////////////



function desactivationAll() {
		actionsArr.length = 0;// очистить массив
		var nodes = document.getElementsByClassName('actions'); //alert(nodes.length);
		for (var i = 0; i < nodes.length; i++) {
			if(nodes[i].className.indexOf('selButton')>=0)
			nodes[i].className=nodes[i].className.substr(0,nodes[i].className.indexOf(' selButton'));
		}
allow_sent_to_beast=0;
document.getElementById("sent_trigger_stimuls_id").style.outline="solid 0px #000000";
document.getElementById("sent_trigger_stimuls_id").style.color="grey";
document.getElementById("cliner_trigger_stimuls_id").style.outline="solid 0px #000000";
document.getElementById("cliner_trigger_stimuls_id").style.color="grey";
// убрать фон действий с пульта
document.getElementById("action_block_id").style.backgroundColor = "#ffffff";
}

function help_motivation()
{
show_dlg_alert2("Кнопки для авторитарного запрета или поощрения только что сделанного.",0);
}
function help_motivation2()
{
show_dlg_alert2("Кнопки для предельно сильного <br>запрета или поощрения только что сделанного.",0);
}
function help_motivation3()
{
show_dlg_alert2("Кнопки с целью вызвать эмпатийный отклик<br>для запрета или поощрения только что сделанного.",0);
}
/*
function help_motivation4()
{
show_dlg_alert2("Приводит в Норму.<br>Снижает Стресс (но не меняет другие гомео-параметры).<br>Гасит удерживаемые режимы Хорошо и Плохо (когда фон стрнаицы становится розовым или голубым).<br>Имеет небольшой позитивный эффект.",0);
}*/
</script>
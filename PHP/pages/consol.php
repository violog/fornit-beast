<?
/*   консоль Бота - отдельное окно
/pages/consol.php
*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');
?>

<!DOCTYPE html>
<html>

<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Консоль событий Бота</title>
</head>

<body style="margin-left:50px;">
  <style>

  </style>
  <div style="position:relative;">
    <h3 style="margin-bottom:2px;">Консоль событий Бота:</h3> последнее - наверху<br>
    <div style="position:absolute;top:0px;left:300px;cursor:pointer;" onClick='cline_log()'><b>Очистить лог</b></div>
  </div>

  <div id='consol_div_id' style="font-family:'Courier New';font-size:15px;"></div>

  <script Language="JavaScript" src="/ajax/ajax.js"></script>
  <script>
    //alert(opener);
    //alert(typeof(opener));
    //if(typeof(opener)!='undefined')  
    function get_info() {
      var AJAX = new ajax_support('/pages/consol_get_info.php', sent_request_info);
      AJAX.send_reqest();

      function sent_request_info(res) {
        //alert(res);
        var out = res;
        if (res.length == 0)
          out = "<span style='color:#999999;font-size:18px;'>Пока нет событий.</span>";
        //alert(out);
        document.getElementById('consol_div_id').innerHTML = out;
      }
      //alert("1");
      setTimeout("get_info()", 2000);

    }
    setTimeout("get_info()", 100);

    function cline_log() {
      var AJAX = new ajax_support('/pages/consol_cliner_info.php', sent_request_info);
      AJAX.send_reqest();

      function sent_request_info(res) {
        document.getElementById('consol_div_id').innerHTML = "";
      }
    }
  </script>
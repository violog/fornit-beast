<?
include_once($_SERVER['DOCUMENT_ROOT']."/get_global_vars.php");

// http://scorcher.ru/ajax/sample_post.php
// http://scorcher/ajax/sample_post.php

header("Cache-Control: no-cache, must-revalidate");
header("Pragma: no-cache");

?>
<html>
<head>
<title>Пример использования</title>
<meta http-equiv="Content-Type" content="text/html; charset=windows-1251">
</head>
<body bgcolor="#FFFFFF" style="margin: 0px 0px 0px 0px;">
<script Language="JavaScript" src="/ajax/ajax_post.js"></script>

<script Language='JavaScript' src='/sys/url_encode.js'></script>
<script language="JavaScript" type="text/javascript">

var text=url_encode("1+2"); //alert(text);
var param='reqvest='+text;
//alert(param);
var AJAX = new ajax_post_support('/ajax/server_script.php',param,sent_reqvest);
AJAX.send_reqest();

function sent_reqvest(res)
{
alert(res);
}
</script>


</body>
</html>
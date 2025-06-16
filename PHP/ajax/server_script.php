<?
include_once($_SERVER['DOCUMENT_ROOT']."/get_global_vars.php");

/*
Обработчик запросов ajax для sample.php


*/

header("Cache-Control: no-store, no-cache, must-revalidate");
header("Cache-Control: post-check=0, pre-check=0", false);
header("Content-type: text/plain; charset=windows-1251");





if(empty($reqvest))
echo "Пустой запрос, м.б. ограничен размер для передачи POST";
else
{
$messOut="Нормально обработан запрос размером ".strlen($reqvest)." байт";
echo $messOut;
}

?> 
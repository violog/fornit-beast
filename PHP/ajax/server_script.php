<?
include_once($_SERVER['DOCUMENT_ROOT']."/get_global_vars.php");

/*
���������� �������� ajax ��� sample.php


*/

header("Cache-Control: no-store, no-cache, must-revalidate");
header("Cache-Control: post-check=0, pre-check=0", false);
header("Content-type: text/plain; charset=windows-1251");





if(empty($reqvest))
echo "������ ������, �.�. ��������� ������ ��� �������� POST";
else
{
$messOut="��������� ��������� ������ �������� ".strlen($reqvest)." ����";
echo $messOut;
}

?> 
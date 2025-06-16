<?
/*  Запуск сервера golang    
http://go/run.php

$command = sprintf(
    'server.exe %d %d',
    intval($_GET['a']),
    intval($_GET['b'])
);
$ret= exec($command, $output); var_dump($output);exit("! $ret | ");

Чтобы запустить фоновый процесс, используйте команду, подобную этой:
exec('nohup java jar "/selenium-server-standalone-3.8.0.jar" > /dev/null 2>&1 & echo $!', $pid);
Это перенаправляет stdout в / dev / null, а затем stderr в stdout, чтобы процесс мог продолжаться в фоновом режиме. После этого он возвращает processid нового фонового процесса и сохраняет его в $ pid.
Затем, позже, вы можете остановить это с exec('kill '.$pid);
*/

//$ret= exec("go_build_server_for_php.exe > /dev/null 2>/dev/null &" ); exit("! $ret");
//$ret= pclose(popen("start /B go_build_server_for_php.exe", "r"));   exit("! $ret");


//$output = shell_exec(' Watch.bat '  . htmlspecialchars($_GET["name"]));

/*
$ret= exec('server.exe'); //echo $ret;
if(strpos($ret,escapeshellarg("Server is listening"))===false)
{
exit("Не запустился сервер Golang...");
}
else
echo $ret;

$res = shell_exec('ps -x');
$res = explode("\n",$res);
print_r($res);
//shell_exec('kill -KILL ProcessID');
*/
header("Expires: Tue, 1 Jul 2003 05:00:00 GMT");
header("Last-Modified: " . gmdate("D, d M Y H:i:s") . " GMT");
header("Cache-Control: no-store, no-cache, must-revalidate");
header("Pragma: no-cache");
header('Content-Type: text/html; charset=UTF-8');

/* Нужно убедиться, что имя файла, который генерит GO является тем, 
что импользуется в shell_exec()!!!!!

*/

$ret= exec("go_build_BOT.exe"); // , $pid
//$res=shell_exec('go_build_main_go.exe');

exit("Статус ".$ret);

// убить процесс
//$ret= exec('kill '.$pid); echo "! $ret";



?>
/*
ИСПОЛЬЗОВАНИЕ:
<script Language="JavaScript" src="/sys/url_encode.js"></script> - для колирования
<script Language="JavaScript" src="/ajax/ajax_form_post.js"></script>
var AJAX = new ajax_form_post_support('form_id','/chat_priv/ajax_submit.php',sent_request_mess);
AJAX.send_form_reqest();
function sent_request_mess(resOut)
{
alert(resOut);
}

на сервере нужно будет раскодировать $_POST[] для каждого имени поля:
foreach($_POST as $k => $v)
{
$_POST[$k]=url_my_decode($v); 
}
extract($_POST, EXTR_SKIP);

function url_my_decode($var_name)
{
$var_name=urldecode($_POST[$var_name]);  
$var_name=str_replace("|#1#|","+",$var_name);// передача плюса
return $var_name;
}

*/
function ajax_form_post_support(form_id,script_url,own_function)
{
var form0 = document.getElementById(form_id);//new FormData();

// Создаем простую копию формы чтобы не портить исходную форму перекодировкой
var simpleCopy = form0.cloneNode(true); //alert(simpleCopy);

// кодируем все поля формы типа urlencode
var elements = simpleCopy.elements;
for (var i = 0, element; element = elements[i++];) {
//element.value=url_encode(element.value);
//       console.log(element.value)
}

var form = new FormData(simpleCopy);
form.append('path', '/');



//	alert(form);
var req;
var timer_id=0;

this.send_form_reqest=function ()
{ 
 //alert(own_function);
this.do_form_reqest(script_url);
}


this.do_form_reqest=function (url) 
{
var brauser=0;
    // branch for native XMLHttpRequest object
    if (window.XMLHttpRequest) 
	{ 
     req = new XMLHttpRequest();  
     brauser=0;   
    // branch for IE/Windows ActiveX version
    } 
	else 
	if (window.ActiveXObject) 
	{
        req = new ActiveXObject("Microsoft.XMLHTTP");
        brauser=1;        
    }

if (typeof(req) == 'undefined') 
{
//alert('Cannot create XMLHTTP instance.');
return false;
}

req.onreadystatechange = function(e) 
{
//timer_id = window.setTimeout("req.abort();", 5000);// убирается с объектом класса
    
    if(req.readyState == 4) 
	{
//clearTimeout(timer_id);
        var status = req.status; 
        // req.statusText; - описание ощибки         
        // only if "OK"
if (req.status == 200) 
{
//alert(req.responseText);
if(typeof(own_function)!='function')
	{
//alert("Не найдена функция для приема сообщений (второй параметр ajax_support())");
return;
	}
own_function(req.responseText);

}
else
{			
//own_function("<error_ajax>");      //alert("Не удалось получить данные:\n" + req.statusText);
}
} 
}
//alert(url);		
req.open("POST", url, true);
/*
var boundary = String(Math.random()).slice(2);
var boundaryMiddle = '--' + boundary + '\r\n';
var boundaryLast = '--' + boundary + '--\r\n'

var body = ['\r\n'];
for (var key in data) {
  // добавление поля
  body.push('Content-Disposition: form-data; name="' + key + '"\r\n\r\n' + data[key] + '\r\n');
}

body = body.join(boundaryMiddle) + boundaryLast;

req.setRequestHeader('Content-Type', 'multipart/form-data; boundary=' + boundary);
req.setRequestHeader("Content-type", "application/x-www-form-urlencoded  charset=windows-1251'");
//req.setRequestHeader('Content-Type', 'application/octet-stream; charset=windows-1251');
//req.setRequestHeader('Content-Type', 'text/html; charset=windows-1251');

//req.setRequestHeader('Content-type', 'application/json; charset=windows-1251');
//req.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
// req.setRequestHeader('Content-Type', 'text/plain; charset=windows-1251');
//req.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded; charset=windows-1251')

*/


//req.setRequestHeader("Content-length", param.length);
//req.setRequestHeader("Connection", "close");

req.send(form);  //alert(param);
}

}
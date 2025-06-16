/*
<script Language="JavaScript" src="/ajax/ajax_post.js"></script>
param="isEditing="+document.forms["form"].isEditing.value+"&receiver_name="+receiver_name+"&comment="+text;
var AJAX = new ajax_post_support('/chat_priv/ajax_submit.php',param,sent_request_mess,urlencoded);
AJAX.send_reqest();
function sent_request_mess(resOut)
{
alert(resOut);
}
// если параметров несколько, то они разделяются &
// param="isEditing="+document.forms["form"].isEditing.value+"&receiver_name="+receiver_name+"&comment="+text;;
*/
function ajax_post_support(script_url,param,own_function,urlencoded)
{
	//alert(script_url);
var req;
var timer_id=0;

this.send_reqest=function ()
{ //alert(out_function_name);
this.do_reqest(script_url);
}


this.do_reqest=function (url) 
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
        status = req.status; 
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

if(urlencoded)//- портит когда передаешь объект file
req.setRequestHeader("Content-type", "application/x-www-form-urlencoded"); 

//req.setRequestHeader("Content-length", param.length);
//req.setRequestHeader("Connection", "close");

req.send(param);
}


}
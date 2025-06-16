// использование: <script Language="JavaScript" src="/ajax/url_encode.js"></script>
// str = url_encode(str);
// псле url_encode на сервере: $worker=urldecode($_GET['worker']);  
//$comment=str_replace("|#1#|","+",$comment);// передача плюса


var global_trans_arr = [];
for (var i = 0x410; i <= 0x44F; i++)
  global_trans_arr[i] = i - 0x350; // А-Яа-я
global_trans_arr[0x401] = 0xA8;    // Ё
global_trans_arr[0x451] = 0xB8;    // ё
//global_trans_arr[0x452] = 0xB9; // №

function url_encode(str)
{ 
var maxlen=1000;
var len=str.length;
var summ="";
var n=0;
while(len>0)
{
var sub = str.substr(n*maxlen,maxlen); //alert(sub);return;
frag=setescape(sub);  //alert(frag); return;
summ=summ+frag;

//if(len>maxlen)
len=len-maxlen;
n++;
}
summ=summ.replace(/\+/g, '|#1#|');
return summ;
}
//////////////////////////////////
function setescape(str)
{
str=str.replace(/№/g, '\%B9');

  var ret = [];
  // Составляем массив кодов символов, попутно переводим кириллицу
  for (var i = 0; i < str.length; i++)
  {
    var n = str.charCodeAt(i);  
    if (typeof global_trans_arr[n] != 'undefined')
      n = global_trans_arr[n];
    if (n <= 0xFF)
      ret.push(n);
  }
  res= window.escape(String.fromCharCode.apply(null, ret));
res=res.replace(/%25B9/g, '%B9');

  return res;
}

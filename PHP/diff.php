<?
// http://go/diff.php

// мое.
$a=array(array(7),
array(2),
array(11),
array(2,9),
array(2,5),
array(2,3),
array(1,2),
array(2,11),
array(3,11),
array(6,11),
array(2,5,9),
array(1,2,8),
array(1,2,3),
array(2,3,5),
array(2,3,9),
array(2,5,8),
array(1,2,9),
array(2,5,11),
array(2,9,10),
array(2,3,11),
array(1,2,11),
array(2,3,5,9),
array(1,2,8,9),
array(1,2,3,9),
array(2,5,8,9),
array(1,2,8,10),
array(2,5,9,10),
array(2,3,9,10),
array(1,2,3,11),
array(2,3,5,11),
array(2,5,8,10),
array(1,2,9,10),
array(1,2,3,9,10),
array(2,5,8,9,10),
array(1,2,8,9,10),
array(2,3,5,9,10)
);
for($n=0;$n<count($a);$n++)
{
sort($a[$n], SORT_NUMERIC);reset($a[$n]);
}

// Паларм
$b=array(array(2,3),
array(1,2),
array(6,11),
array(1,2,3),
array(1,2,8),
array(2,3,9),
array(1,2,9),
array(2,3,5),
array(1,2,11),
array(1,6,11),
array(6,8,11),
array(2,3,11),
array(2,3,5,8),
array(1,2,8,9),
array(1,2,3,9),
array(2,3,5,9),
array(1,2,3,8),
array(1,2,3,11),
array(2,3,5,11),
array(1,2,8,11),
array(2,3,9,10),
array(1,6,8,11),
array(1,2,8,10),
array(1,2,9,10),
array(2,3,5,8,9),
array(1,2,3,5,9),
array(1,2,3,5,8),
array(1,2,3,8,9),
array(1,2,3,8,10),
array(2,3,5,9,10),
array(1,2,3,5,11),
array(1,2,3,8,11),
array(2,3,5,8,10),
array(1,2,8,9,10),
array(1,2,3,9,10),
array(2,3,5,8,11),
array(1,2,3,5,8,9),
array(1,2,3,5,8,11),
array(1,2,3,5,8,10),
array(2,3,5,8,9,10),
array(1,2,3,8,9,10),
array(1,2,3,5,9,10),
array(1,2,3,5,8,9,10)
);

$absentInB="";
foreach($a as $p)
{
	if(!in_array($p,$b))
	{
		$str="";
		foreach($p as $s)
		{
$str.=$s.", ";
		}

		$absentInB.=$str."<br>";
	}

}
echo "В новой нет:<br>$absentInB <hr>";

$absentInA="";
foreach($b as $p)
{
	if(!in_array($p,$a))
	{
		$str="";
		foreach($p as $s)
		{
$str.=$s.", ";
		}

		$absentInA.=$str."<br>";
	}

}
exit("В старой нет:<br>$absentInA");

?>
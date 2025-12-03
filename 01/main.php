<?php


$data = file('input.txt');
$r  = [];
foreach ($data as $line) {
    if (trim($line) === '') {
        continue;
    }
    $d = $line[0];
    $n = (int)substr($line, 1);
    switch ($d) {
    case 'L':
        $r[] = -$n;
        break;
    case 'R':
        $r[] = $n;
        break;
    }
}
print_r($r);
$z = 0;
$i = 50;

foreach ($r as $v) {
    $i += $v;
    while ($i < 0) {
        $i += 100;
    }
    while ($i >= 100) {
        $i -= 100;
    }
    if ($i === 0) {
        $z++;
    }
    print_r( "v:$v p:$i z:$z\n" );
}
print_r(  'Res: ' . $z);

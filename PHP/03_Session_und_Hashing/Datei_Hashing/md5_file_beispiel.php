<?php

    // Berechnet den MD5-Hashwert der Datei "lorem.txt"
    // und speichert das Ergebnis in der Variablen $hashwertmd5.
    $hashwertmd5 = md5_file("lorem.txt");
    // Gibt den Namen der Variable '$hashwertmd5' als String aus,
    // gefolgt von dem tatsächlichen MD5-Hashwert der Datei.
    echo '$hashwertmd5' . ": $hashwertmd5 <br>";

    // Berechnet den SHA1-Hashwert der Datei "lorem.txt"
    // und speichert das Ergebnis in der Variablen $hashwertsha1.
    $hashwertsha1 = sha1_file("lorem.txt");
    // Gibt den Namen der Variable '$hashwertsha1' als String aus,
    // gefolgt von dem tatsächlichen SHA1-Hashwert der Datei.
    echo '$hashwertsha1' . ": $hashwertsha1 <br>";
     
?>

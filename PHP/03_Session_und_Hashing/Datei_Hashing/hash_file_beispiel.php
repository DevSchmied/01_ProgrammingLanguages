<?php
    
    // Berechnet den MD5-Hashwert der Datei "lorem.txt"
    // unter Verwendung der Funktion hash_file() und speichert das Ergebnis
    // in der Variablen $hashwertmd5.
    $hashwertmd5 = hash_file("md5", "lorem.txt");
    // Gibt den Namen der Variable '$hashwertmd5' als String aus,
    // gefolgt von dem tatsächlichen MD5-Hashwert der Datei.
    // Die Verkettung erfolgt mit einem Doppelpunkt und einem Zeilenumbruch.
    echo '$hashwertmd5' . ": " . "$hashwertmd5 <br>";

    // Berechnet den SHA1-Hashwert der Datei "lorem.txt"
    // unter Verwendung der Funktion hash_file() und speichert das Ergebnis
    // in der Variablen $hashwertsha1.
    $hashwertsha1 = hash_file("sha1", "lorem.txt");
    // Gibt den Namen der Variable '$hashwertsha1' als String aus,
    // gefolgt von dem tatsächlichen SHA1-Hashwert der Datei.
    // Die Verkettung erfolgt mit einem Doppelpunkt, jedoch ohne Zeilenumbruch.
    echo '$hashwertsha1' . ": " . "$hashwertsha1";

?>

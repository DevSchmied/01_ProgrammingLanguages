<?php
    // Definiert eine Zeichenkette, die gehasht werden soll
    $zeile = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.";

    // Berechnet den MD5-Hash einer Datei namens "lorem.txt"
    // hash_file() ist eine Funktion, die einen Hash-Wert einer Datei basierend auf dem angegebenen Algorithmus berechnet
    $hashwertFileMd5 = hash_file("md5", "lorem.txt");
    // Gibt den MD5-Hash der Datei aus, dabei wird der Hash-Wert in der Variable $hashwertFileMd5 gespeichert
    echo '$hashwertFileMd5' . ": $hashwertFileMd5<br>";

    // Berechnet den SHA1-Hash einer Datei namens "lorem.txt"
    $hashwertFileSha1 = hash_file("sha1", "lorem.txt");
    // Gibt den SHA1-Hash der Datei aus, dabei wird der Hash-Wert in der Variable $hashwertFileSha1 gespeichert
    echo '$hashwertFileSha1' . ": $hashwertFileSha1<br>";

    // Berechnet den MD5-Hash der Zeichenkette $zeile
    // hash() ist eine Funktion, die einen Hash-Wert basierend auf dem angegebenen Algorithmus berechnet
    $hashwertZeile = hash("md5", $zeile);
    // Gibt den MD5-Hash der Zeichenkette aus, dabei wird der Hash-Wert in der Variable $hashwertZeile gespeichert
    echo '$hashwertZeile' . ": $hashwertZeile<br>";

    // Berechnet den MD5-Hash der Zeichenkette "lorem.txt" (nicht der Datei, sondern der Zeichenkette selbst)
    // md5() ist eine Funktion, die den MD5-Hash einer Zeichenkette berechnet
    $hashwertMd5 = md5("lorem.txt");
    // Gibt den MD5-Hash der Zeichenkette "lorem.txt" aus, dabei wird der Hash-Wert in der Variable $hashwertMd5 gespeichert
    echo '$hashwertMd5' . ": $hashwertMd5<br>";

    // Berechnet den SHA1-Hash der Zeichenkette "lorem.txt" (nicht der Datei, sondern der Zeichenkette selbst)
    // sha1() ist eine Funktion, die den SHA1-Hash einer Zeichenkette berechnet
    $hashwertSha1 = sha1("lorem.txt");
    // Gibt den SHA1-Hash der Zeichenkette "lorem.txt" aus, dabei wird der Hash-Wert in der Variable $hashwertSha1 gespeichert
    echo '$hashwertSha1' . ": $hashwertSha1<br>";

    // Berechnet den SHA1-Hash der Zeichenkette $zeile
    $hashwertZeileSha1 = sha1($zeile);
    // Gibt den SHA1-Hash der Zeichenkette $zeile aus, dabei wird der Hash-Wert in der Variable $hashwertZeileSha1 gespeichert
    echo '$hashwertZeileSha1' . ": $hashwertZeileSha1<br>";
?>

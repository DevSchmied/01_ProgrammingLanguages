<?php
try {
    // Überprüfen, ob die Datei "datumZeit.txt" existiert.
    // Wenn sie existiert, wird sie zum Anhängen geöffnet, andernfalls wird sie neu erstellt.
    $meineDatei = file_exists("datumZeit.txt") ? fopen("datumZeit.txt", "a") : fopen("datumZeit.txt", "w"); 

    // Wenn die Datei nicht geöffnet werden kann, wird eine Ausnahme ausgelöst.
    if (!$meineDatei) {
        throw new Exception('Datei kann nicht geöffnet werden. ');
    }

    // Aktuelles Datum und Uhrzeit werden im gewünschten Format erstellt.
    $datum = "Heute ist der " . date("d.m.Y"); 
    $zeit = " Es ist " . date("h:i:s") . " Uhr." . "\r\n";
    
    // Dateiname und Datum/Zeit werden in die Datei geschrieben.
    fwrite($meineDatei, "Datei: " . basename(__FILE__) . "\r\n");
    fwrite($meineDatei, $datum);
    fwrite($meineDatei, $zeit);
    fwrite($meineDatei, "\r\n");
    
    // Die Datei wird geschlossen.
    fclose($meineDatei);

} catch (Exception $e) {
    // Wenn eine Ausnahme auftritt, wird eine Fehlermeldung ausgegeben und das Skript beendet.
    exit("Ein Fehler ist aufgetreten - Sorry!");
}
?>

<!DOCTYPE html>
<html lang="de">
<head>
    <meta charset="utf-8">
    <title>Dateihandling</title>
    <meta name="robots" content="noindex, nofollow">
    <meta name="author" content="Sergey Grinko">
    <meta name="description" content="Grundgerüst einer HTML Seite">
    <meta name="keywords" content="HTML, Grundgerüst, Aufbau">
</head>
<body>
    <h1>DatumZeitDatei auslesen</h1>
    <p>
        <?php
        try {
            // Überprüfen, ob die Datei "datumZeit.txt" existiert.
            if (!file_exists("datumZeit.txt")) {
                throw new Exception('Datei nicht gefunden. ');
            }

            // Die Datei wird zum Lesen geöffnet.
            $meineDatei = fopen("datumZeit.txt", "r");

            // Wenn die Datei nicht geöffnet werden kann, wird eine Ausnahme ausgelöst.
            if (!$meineDatei) {
                throw new Exception('Datei kann nicht geöffnet werden. ');
            }

            // Die Datei wird Zeile für Zeile gelesen und der Inhalt wird ausgegeben.
            while (!feof($meineDatei)) {
                echo fgets($meineDatei) . "<br>";
            }

            // Die Datei wird geschlossen.
            fclose($meineDatei);

        } catch (Exception $e) {
            // Wenn eine Ausnahme auftritt, wird die Fehlermeldung ausgegeben und das Skript beendet.
            echo $e->getMessage();
            exit("Ein Fehler ist aufgetreten - Sorry!");
        }                        
        ?>
    </p>
</body>
</html>

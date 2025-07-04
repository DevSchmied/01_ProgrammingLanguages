<?php
    // Funktion zur Ermittlung der nächsten verfügbaren ID
    function getNextId() {
        $idFile = 'id.txt'; // Name der Datei, in der die letzte ID gespeichert ist
        // Überprüfen, ob die Datei bereits existiert
        if (file_exists($idFile)) {
            // Lesen der letzten ID aus der Datei und in eine Ganzzahl umwandeln
            $id = (int)file_get_contents($idFile);
        } else {
            // Wenn die Datei nicht existiert, Startwert 0 verwenden
            $id = 0; 
        }
        // Berechnung der nächsten ID
        $nextId = $id + 1;
        // Speichern der neuen ID in die Datei
        file_put_contents($idFile, $nextId);
        // Rückgabe der neuen ID
        return $nextId; 
    }

    // Überprüfen, ob das Formular abgesendet wurde und die erforderlichen Felder gesetzt sind
    if (isset($_POST['vorname']) && isset($_POST['nachname']) && isset($_POST['email'])) {
        // Aufruf der Funktion zur Ermittlung der nächsten ID
        $id = getNextId(); 

        // Öffnen der Datei 'user.txt' im Anhängemodus (erstellen, falls nicht vorhanden)
        $datei = fopen('user.txt', 'a'); 
        // Schreiben der Nutzerdaten mit der neuen ID in die Datei
        fwrite($datei, 'ID' . $id . ': ' . $_POST['vorname'] . ' ' . $_POST['nachname'] . ' ' . $_POST['email'] . "\r\n");
        // Schließen der Datei
        fclose($datei);
    }

    try {
        // Überprüfen, ob die Datei 'user.txt' existiert
        if (!file_exists('user.txt')) {
            // Auslösen einer Ausnahme, wenn die Datei nicht gefunden wurde
            throw new Exception('Datei nicht gefunden. ');
        } else {
            // Öffnen der Datei 'user.txt' zum Lesen
            $inhalt = fopen('user.txt', 'r');
            // Schleife zum Durchlesen der Datei bis zum Ende
            while (!feof($inhalt)) {
                // Lesen einer Zeile aus der Datei
                $zeile = fgets($inhalt);
                // Ausgabe der Zeile in sicherer HTML-Formatierung
                echo htmlspecialchars($zeile) . '<br>'; 
            }
            // Schließen der Datei
            fclose($inhalt);
        }
    } catch (Exception $e) {
        // Ausgabe der Fehlermeldung, wenn eine Ausnahme aufgetreten ist
        echo $e->getMessage();
    }  
    
?>
<!DOCTYPE html>
<html lang="de">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <!-- Formular zur Eingabe von Nutzerdaten -->
    <form action="<?=$_SERVER['PHP_SELF'];?>" method="POST">
        <fieldset>
            <legend>Nutzerdaten</legend>
            <p>Vorname: <input type="text" name="vorname"></p>
            <p>Nachname: <input type="text" name="nachname"></p>
            <p>Email: <input type="text" name="email" required="required"></p>
        </fieldset>
        <!-- Absendeknopf des Formulars -->
        <input type="submit" value="abschicken">
    </form>
</body>
</html>

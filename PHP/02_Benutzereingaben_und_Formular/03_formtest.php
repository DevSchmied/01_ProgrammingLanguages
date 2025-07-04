<?php
// Zeigt den Inhalt des $_POST-Arrays an, für Debugging-Zwecke
// var_dump($_POST);

// Überprüft, ob das Formular abgesendet wurde
if (isset($_POST["senden"])) {
    
    // Überprüft, ob das Feld "vorname" nicht leer ist
    if (!empty($_POST["vorname"])) {
        $vorname = $_POST["vorname"]; // Speichert den eingegebenen Vornamen
    } else {
        $vorname = "Unbekanntes Wesen"; // Setzt einen Standardwert, wenn der Vorname nicht eingegeben wurde
    }

    // Überprüft, ob das Feld "sprache" gesetzt und ein Array ist
    if (isset($_POST["sprache"]) AND is_array($_POST["sprache"])) {
        $sprachen = []; // Initialisiert ein leeres Array, um die ausgewählten Sprachen zu speichern
        
        // Durchläuft das Array der ausgewählten Sprachen
        foreach ($_POST["sprache"] as $value) {
            // Überprüft den Wert und fügt die entsprechende Sprache zum Array hinzu
            switch ($value) {
                case "D":
                    $sprachen[] = "Deutsch";
                    break;
                case "EN":
                    $sprachen[] = "Englisch";
                    break;
                case "F":
                    $sprachen[] = "Französisch";
                    break;
                case "PL":
                    $sprachen[] = "Polnisch";
                    break;
            }
        }

        // Ausgabe der Nachricht mit Namen und ausgewählten Sprachen
        echo "<p>Hallo $vorname. Du sprichst die Sprache(n): ";
        
        // Durchläuft das Array der Sprachen und gibt jede Sprache aus
        foreach ($sprachen as $sprache) {
            echo $sprache . " ";
        }
        echo ".</p>"; // Beendet den Satz mit einem Punkt
    } else {
        // Wenn keine Sprachen ausgewählt wurden, Ausgabe der Nachricht mit Namen
        echo "<p>Hallo $vorname. Du hast keine Sprache ausgewählt.</p>";
    }
}
?>


<!DOCTYPE html>
<html lang="de">
<head>
    <meta charset="utf-8">
    <title>Formular</title>
    <meta name="robots" content="noindex, nofollow">
    <meta name="author" content="Sergey">
    <meta name="description" content="Grundgerüst einer HTML Seite">
    <meta name="keywords" content="HTML, Grundgerüst, Aufbau">
</head>
<body>
    <h1>Ein Formular</h1>
    <form action="<?= $_SERVER["PHP_SELF"]; ?>" method="post">
        <input type="text" name="vorname" id="name">
        <label for="name"> Name</label>
        <br>
        <input type="checkbox" id="sprache1" name="sprache[]" value="D">
        <label for="sprache1"> Deutsch</label>
        <br>
        <input type="checkbox" id="sprache2" name="sprache[]" value="EN">
        <label for="sprache2"> Englisch</label>
        <br>
        <input type="checkbox" id="sprache3" name="sprache[]" value="F">
        <label for="sprache3"> Französisch</label>
        <br>
        <input type="checkbox" id="sprache4" name="sprache[]" value="PL">
        <label for="sprache4"> Polnisch</label>            
        <br>
        <input type="submit" name="senden" value="abschicken">
        <input type="reset" value="zurücksetzen">
    </form>
</body>
</html>

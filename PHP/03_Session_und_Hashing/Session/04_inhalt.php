<?php

    // Starten der Session
    session_start();
    

    // Überprüfen, ob das Formular abgeschickt wurde (Überprüfung auf Vorhandensein des Feldes "name" in den POST-Daten)
    if (isset($_POST["name"])) {
        // Überprüfen, ob Benutzername und Passwort den vorgegebenen Werten entsprechen
        if ($_POST["name"] == "Sergey" && $_POST["pwd"] == "qwerty") {
            // Wenn sie übereinstimmen, speichern des Benutzernamens und Passworts in der Session
            $_SESSION["name"] = $_POST["name"];
            $_SESSION["passwort"] = $_POST["pwd"];	
        } 
    }

    // Wenn der Benutzername nicht in der Session gesetzt ist, beenden des Skripts und Anzeigen einer Nachricht
    if (!isset($_SESSION["name"])) {
        exit("Hier geht es nicht weiter. Bitte <a href='04_login.php'>einloggen.</a>");
    }

?>

<!DOCTYPE html>
<html lang="de">

    <head>
        <meta charset="utf-8">
        <title>LogIn</title>

        <!-- Anweisung an Suchmaschinen, die Seite nicht zu indexieren und den Links nicht zu folgen -->
        <meta name="robots" content="noindex, nofollow">
        
        <!-- Angabe des Autors der Seite -->
        <meta name="author" content="Sergey Grinko">

        <!-- Beschreibung der Seite -->
        <meta name="description" content="Grundgerüst einer HTML Seite">
        
        <!-- Schlüsselwörter für Suchmaschinen -->
        <meta name="keywords" content="HTML, Grundgerust, Aufbau">

    </head>

    <body>

        <h1>geschützter Bereich</h1>
        <p>
            <?php
                // Ausgabe des Benutzernamens und des Passworts aus der Session
				echo "Hallo " . $_SESSION["name"] . " dein Passwort lautet: " . $_SESSION["passwort"];
			?>
        </p>
        <a href="04_login.php">LogOut<a>
        <a href="04_seite2.php"><br>Seite2<a>
            
        <form action="04_seite3.php" method="POST">

            <!-- Der folgende Code ist auskommentiert und wird daher nicht ausgeführt:
            <input type="hidden" name="session_name" value="<?= $_SESSION['name']; ?>">
            <input type="hidden" name="session_passwort" value="<?= $_SESSION['passwort']; ?>">
            <button type="submit">seite3</button>  
            -->
            <button onclick="href='04_seite3.php';">seite3</button>
            
        </form>
    </body>

</html>

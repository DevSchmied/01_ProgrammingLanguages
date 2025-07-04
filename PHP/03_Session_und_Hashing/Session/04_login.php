<?php
session_start();
// Starten der Session

// Session-Daten leeren
$_SESSION = array();

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

        <h1>LogIn-Seite</h1>
        <p>

            <!-- Formular zur Anmeldung -->
            <form action="04_inhalt.php" method="POST">

                <!-- Eingabefeld für den Benutzernamen -->
                <input type="text" name="name" required="required">

                <!-- Eingabefeld für das Passwort -->
                <input type="password" name="pwd">
                
                <!-- Absende-Button für das Formular -->
                <input type="submit" value="Login">
            </form>
        </p>
    </body>
    
</html>

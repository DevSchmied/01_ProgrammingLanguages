<!DOCTYPE html>
<html>
    <head>
        <title>Home</title> <!-- Titel der Startseite -->
    </head>
    <?php
    // Füge den Header-Bereich von der header.php-Datei ein
    include("header.php");
    ?>
    <body>
        <!-- Hauptinhalt der Startseite -->
        <p>Ich bin Body</p>
    </body>

    <!-- Anzeige des aktuellen Datums, indem die Funktion datum() aus der functions.php-Datei aufgerufen wird -->
    <p>Das aktuelle Datum: <?php include_once("functions.php"); datum(); ?></p>

    <?php
    // Füge den Footer-Bereich von der footer.php-Datei ein
    require("footer.php");
    ?>
</html>

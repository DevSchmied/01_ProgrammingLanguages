<!DOCTYPE html>
<html>
    <head>
        <title>About</title> <!-- Titel der „Über uns“-Seite -->
    </head>

    <?php
    // Füge den Header-Bereich von der header.php-Datei ein
    include("header.php");
    ?>

    <body>
        <!-- Hauptinhalt der „Über uns“-Seite -->
        <h1>Das Unternehmen XYZ</h1>
        <p>Das sind die Informationen über uns:</p>
    </body>

    <!-- Anzeige des aktuellen Datums, indem die Funktion datum() aus der functions.php-Datei aufgerufen wird -->
    <p>Das aktuelle Datum: <?php include_once("functions.php"); datum(); ?></p>
    
    <?php
    // Füge den Footer-Bereich von der footer.php-Datei ein
    include("footer.php");
    ?>

</html>

<?php
    // Einbinden der PHP-Datei, die die Berechnungslogik enthält
    include_once("05_berechnung.php");
?>
<!DOCTYPE html>
<html>
<head>
    <title>Umsatzsteuerrechner</title>
    <meta charset="utf-8">
    <meta name="author" content="Sergey Grinko">
    <meta name="description" content="Umsatzsteuerrechner">        
    <meta name="keywords" content="Monate">
</head>
<body>
    <h1>Umsatzsteuerrechner</h1>
    <!-- Formular für die Eingabe der Werte und Berechnung -->
    <form action="<?=$_SERVER["PHP_SELF"]; ?>" method="POST">
        <fieldset>
            <legend>Rechner</legend>
            <p>
                <!-- Eingabefeld für den Nettowert -->
                <input type="number" name="netto" id="i1" required="required" value="<?= $netto; ?>" min="0" step=".01">
                <label for="i1">Nettowert</label>                
            </p>
            <p>            
                <!-- Eingabefeld für den Umsatzsteuersatz -->
                <input type="number" name="ustSatz" id="i2" required="required" value="<?= $ustSatz; ?>" min="0">
                <label for="i2">Umsatzsteuersatz</label>            
            </p>
            <p>
                <!-- Ausgabefeld für den Umsatzsteuerbetrag (nur lesbar) -->
                <input type="number" name="ustBetrag" id="i3" readonly="readonly" value="<?= $ustBetrag; ?>">
                <label for="i3">Umsatzsteuerbetrag</label>                
            </p>
            <!-- Ausgabefeld für den Bruttobetrag (nur lesbar) -->
            <input type="number" name="brutto" id="i4" readonly="readonly" value="<?= $brutto; ?>">
            <label for="i4">Bruttobetrag</label>                
            </p>
        </fieldset>
        <fieldset>
            <p>
                <!-- Absende-Button für das Formular -->
                <input type="submit" value="berechne" />
            </p>            
        </fieldset>
    </form>
</body>
</html>

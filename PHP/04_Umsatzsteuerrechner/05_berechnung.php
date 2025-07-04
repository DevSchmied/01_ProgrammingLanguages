<?php
    // Initialisierung der Variablen für Nettowert, Umsatzsteuersatz, Umsatzsteuerbetrag und Bruttobetrag
    $netto = "";
    $ustSatz = "";
    $ustBetrag = "";
    $brutto = "";

    // Überprüfung, ob das Formular abgesendet wurde
    if (!empty($_POST)) {
        // Durchlaufen aller POST-Daten, um die Werte für Netto und Umsatzsteuersatz zu extrahieren
        foreach ($_POST as $key => $wert) {
            // Zuweisen des Wertes von 'netto' zu $netto
            if ($key == "netto") {
                $netto = $_POST["netto"];
            }
            // Zuweisen des Wertes von 'ustSatz' zu $ustSatz
            if ($key == "ustSatz") {
                $ustSatz = $_POST["ustSatz"];
            }
        }
    }

    // Berechnung des Bruttobetrags und des Umsatzsteuerbetrags, wenn sowohl Netto als auch Umsatzsteuersatz vorhanden sind
    if ($netto !== "" && $ustSatz !== "") {
        // Berechnung des Bruttobetrags: Nettowert plus Umsatzsteuer
        $brutto = $netto * (1 + $ustSatz / 100);
        // Berechnung des Umsatzsteuerbetrags: Bruttobetrag minus Nettowert
        $ustBetrag = $brutto - $netto;

        // Formatierung der Ergebnisse mit Komma als Dezimaltrennzeichen und Punkt als Tausendertrennzeichen
        $brutto = number_format($brutto, 2, ',', '.');
        $ustBetrag = number_format($ustBetrag, 2, ',', '.');
    }
?>

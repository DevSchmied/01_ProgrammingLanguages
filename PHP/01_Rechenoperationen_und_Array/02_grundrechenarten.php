<?php

    // Erzwingt strikte Typendeklarationen in PHP, um sicherzustellen, dass Argumente und Rückgabewerte den angegebenen Typen entsprechen.
    declare(strict_types=1);

    // Variable $zeichen wird deklariert (hier nicht initialisiert, aber in der Funktion später genutzt).
    $zeichen;
 
    // Definition der Funktion 'grundrechenarten', die zwei Zahlen und ein Zeichen als Argumente akzeptiert.
    // Diese Funktion führt eine arithmetische Operation basierend auf dem Zeichen aus.
    function grundrechenarten($zahl1, $zahl2, string $zeichen) {
        // Switch-Anweisung, die den Wert von $zeichen überprüft.
        switch($zeichen) {
            // Fall für Addition.
            case "+":
                return $zahl1 + $zahl2;
                break; // break wird nach return hier nicht benötigt, da die Funktion bereits verlassen wurde.
            
            // Fall für Subtraktion.
            case "-":
                return $zahl1 - $zahl2;
                break; // break wird nach return hier nicht benötigt.

            // Fall für Multiplikation.
            case "*":
                return $zahl1 * $zahl2;
                break; // break wird nach return hier nicht benötigt.

            // Fall für Division.
            case "/":
                // Überprüfung, ob die zweite Zahl 0 ist, da Division durch 0 nicht erlaubt ist.
                if ($zahl2 === 0) {
                    echo "Durch 0 darf man nicht teilen!";
                    break; // Beendet die Ausführung der Funktion, gibt aber keinen Wert zurück.
                }
                return $zahl1 / $zahl2;
                break; // break wird nach return hier nicht benötigt.

            // Standardfall, wenn das Zeichen nicht erkannt wird.
            default:
                echo "Das habe ich noch nicht gelernt!";
        }
    }

    // Initialisierung der Variablen $zahl1, $zahl2 und $zeichen mit den Werten 5, 3 und "+".
    $zahl1 = 5;
    $zahl2 = 3;
    $zeichen = "+";
    
    // Aufruf der Funktion 'grundrechenarten' mit den Argumenten $zahl1, $zahl2 und $zeichen.
    // Ausgabe des Ergebnisses der Funktion in einem HTML-Format.
    echo "<br>$zahl1 $zeichen $zahl2 = " . grundrechenarten($zahl1, $zahl2, $zeichen);

?>

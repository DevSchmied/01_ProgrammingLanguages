<?php
    // Initialisierung eines 2D-Arrays, das Daten über Hunde enthält.
    // Jedes Unterarray repräsentiert einen Hund und enthält seinen Namen, seine Größe und sein Alter.
    $hunde = array(
        array("Fiffi", 35, 4),
        array("Waldi", 25, 8),
        array("Wilma", 42, 19)
    );
    $hunde[] = 33;

    // Eine foreach-Schleife, die über jedes Element des Arrays $hunde iteriert.
    foreach ($hunde as $hund) {
        // Überprüfung, ob das aktuelle Element ein Array ist.
        if(is_array($hund)) {
            echo "<br>";
            // Eine for-Schleife, die über jedes Element des aktuellen Hund-Arrays iteriert.
            for ($i = 0; $i < count($hund); $i++) {
                // Überprüfung, ob der aktuelle Index 0 ist, um den Namen des Hundes auszugeben.
                if ($i == 0) {
                    echo "Der Name: " . $hund[$i] . "<br>";
                }
                // Überprüfung, ob der aktuelle Index 1 ist, um die Größe des Hundes auszugeben.
                if ($i == 1) {
                    echo "Die Größe: " . $hund[$i] . "<br>";
                }
                // Überprüfung, ob der aktuelle Index 2 ist, um das Alter des Hundes auszugeben.
                if ($i == 2) {
                    echo "Das Alter: " . $hund[$i] . "<br>";
                }
            }
        } else {
            echo "<br>Der Wert des Elementes ist: ".  $hund . "<br>";
        }
    }
?>

<?php
include_once 'c_form.php';

// Die Klasse Kreis erbt von der Klasse Form und fügt spezifische Eigenschaften für einen Kreis hinzu
class Kreis extends Form {
    public $radius;

    // Konstruktor der Klasse Kreis, um Farbe und Radius beim Erstellen festzulegen
    public function __construct(string $farbe, float $radius = 1.1) {
        parent::__construct($farbe);
        $this->radius = $radius;
    }

    // Methode zur Berechnung der Fläche des Kreises
    public function berechneFlaeche(float $radius) : float {
        return M_PI * ($radius ** 2);
    }

    // Überschriebene Methode, um eine textuelle Beschreibung des Kreises zurückzugeben
    public function __toString() : string {
        return "Ein Kreis mit Farbe: " . $this->farbe . " und Radius: " . $this->radius . "<br>";
    }
}
?>

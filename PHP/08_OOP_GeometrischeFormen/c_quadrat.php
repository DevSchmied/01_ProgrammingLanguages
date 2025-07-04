<?php
// Einbinden der Basisklasse Form
include_once 'c_form.php';

// Die Klasse Quadrat erbt von der Klasse Form
class Quadrat extends Form {
    public $kantenlaenge;

    // Konstruktor zur Initialisierung der Kantenlänge
    public function __construct(string $farbe, float $kantenlaenge = 10.0) {
        parent::__construct($farbe);
        $this->kantenlaenge = $kantenlaenge;
    }

    // Berechnet die Fläche des Quadrats
    public function berechneFlaeche(float $kantenlaenge) : float {
        return $kantenlaenge ** 2;
    }

    // Gibt eine Beschreibung des Quadrats als String zurück
    public function __toString() : string {
        return "Ein Quadrat mit Farbe: " . $this->farbe . " und Kantenlänge: " . $this->kantenlaenge . "<br>";
    }
}
?>

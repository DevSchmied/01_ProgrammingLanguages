<?php
// Einbinden der Basisklasse Form
include_once 'c_form.php';

// Die Klasse Rechteck erbt von der Klasse Form
class Rechteck extends Form {
    public $hoehe;
    public $breite;
  
    // Konstruktor zur Initialisierung der Höhe und Breite, Standardwerte werden gesetzt
    public function __construct(string $farbe, float $hoehe = 1, float $breite = 2) {
        parent::__construct($farbe);
        $this->hoehe = $hoehe;
        $this->breite = $breite;
    }

    // Berechnet die Fläche des Rechtecks
    public function berechneFlaeche(float $hoehe = 1, float $breite = 2) : float {
        return $hoehe * $breite;
    }

    // Gibt eine Beschreibung des Rechtecks als String zurück
    public function __toString() : string {
        return "Rechteck mit Farbe: ".$this->farbe.", Höhe: ".$this->hoehe." und Breite: ".$this->breite;
    }
}
?>

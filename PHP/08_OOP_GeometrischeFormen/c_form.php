<?php
    // Die Klasse Form stellt die Basis für geometrische Formen bereit
    class Form {

        // Die geschützte Eigenschaft für die Farbe der Form
        protected $farbe = "schwarz";

        // Konstruktor zur Initialisierung der Farbe
        public function __construct(string $farbe) {
            $this->farbe = $farbe;
        }

        // Getter für die Farbe
        public function getFarbe() {
            return $this->farbe;
        }

        // Setter für die Farbe
        public function setFarbe(string $farbe) {
            $this->farbe = $farbe;
        }

    }
?>

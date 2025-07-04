<?php
// Einbinden der notwendigen Klassen für die geometrischen Formen
include_once 'c_rechteck.php';
include_once 'c_quadrat.php';
include_once 'c_kreis.php';

// Erstellen eines Rechteck-Objekts mit der Farbe "schwarz"
$rechteck = new Rechteck("schwarz");

// Berechnung und Ausgabe der Fläche des Rechtecks mit benutzerdefinierten Werten
echo "Die Fläche des Rechteckes: " . $rechteck->berechneFlaeche(10.5, 20) . "<br>";

// Berechnung und Ausgabe der Fläche des Rechtecks mit Standardwerten
echo "Die Fläche des Rechtecks mit Standardwerten: " . $rechteck->berechneFlaeche() . "<br>"; 

// Ausgabe der Beschreibung des Rechtecks
echo $rechteck . "<br>";

// Ändern der Farbe des Rechtecks auf "blau" und Ausgabe der neuen Farbe
$rechteck->setFarbe("blau");
echo "Die Farbe des Rechtecks: " . $rechteck->getFarbe() . "<br>";

// Erstellen eines Quadrat-Objekts mit der Farbe "gelb"
$quadrat = new Quadrat("gelb");

// Berechnung und Ausgabe der Fläche des Quadrats
echo "Die Fläche des Quadrats: " . $quadrat->berechneFlaeche(11.11) . "<br>";

// Ausgabe der Beschreibung des Quadrats
echo $quadrat;

// Erstellen eines Kreis-Objekts mit der Farbe "gruen"
$kreis = new Kreis("gruen");

// Berechnung und Ausgabe der Fläche des Kreises
echo "Die Fläche des Kreises: " . $kreis->berechneFlaeche(2) . "<br>";

// Ausgabe der Beschreibung des Kreises
echo $kreis;
?>

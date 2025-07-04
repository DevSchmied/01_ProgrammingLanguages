<?php
session_start();
// Start der Session und Überprüfung, ob der Benutzer angemeldet ist
if (!isset($_SESSION["name"])) {
	// Wenn der Benutzer nicht angemeldet ist, wird die Ausführung des Skripts hier abgebrochen und eine Nachricht ausgegeben
	exit("Hier geht's nicht weiter. Bitte <a href='login.php'>einloggen.</a>");
}
?>


<!DOCTYPE html>
<html lang="de">
	<head>
		<meta charset="utf-8">
		<title>Seite3</title>
		<!-- Anweisung an Suchmaschinen, die Seite nicht zu indexieren und den Links nicht zu folgen -->
		<meta name="robots" content="noindex, nofollow">
		<!-- Angabe des Autors der Seite -->
		<meta name="author" content="Sergey Grinko">
		<!-- Beschreibung der Seite -->
		<meta name="description" content="Grundgerüst einer HTML Seite">
		<!-- Schlüsselwörter für Suchmaschinen -->
		<meta name="keywords" content="HTML, Grundgerüst, Aufbau">
	</head>
	<body>
		<h1>Seite3</h1>
		<p>
			<?php
				// Ausgabe der Begrüßungsnachricht mit dem Benutzernamen aus der Session
				echo "Hallo " . $_SESSION["name"] . ". Ich bin Seite3.";
				$_SESSION = array();
			?>
		</p>
	</body>
</html>

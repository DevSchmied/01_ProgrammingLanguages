<?php
// Verbindung zur Datenbank herstellen
$servername = "localhost";
$username = "root"; // Ersetzen Sie "root" durch Ihren MySQL-Benutzernamen
$password = ""; // Ersetzen Sie dies durch Ihr MySQL-Passwort
$dbname = "benutzer_db"; // Datenbankname auf benutzer_db ändern

// Erstellen der Verbindung
$conn = new mysqli($servername, $username, $password, $dbname);

// Überprüfen, ob die Verbindung funktioniert
if ($conn->connect_error) {
    die("Verbindung fehlgeschlagen: " . $conn->connect_error);
}

// Benutzer hinzufügen
if (isset($_POST['name']) && isset($_POST['email'])) {
    $name = $_POST['name'];
    $email = $_POST['email'];

    // Eingaben vor SQL-Injektionen schützen
    $name = $conn->real_escape_string($name);
    $email = $conn->real_escape_string($email);

    // SQL-Query zum Einfügen des Benutzers
    $sql = "INSERT INTO users (name, email) VALUES ('$name', '$email')";

    if ($conn->query($sql) === TRUE) {
        echo "<p>Neuer Benutzer erfolgreich hinzugefügt.</p>";
    } else {
        echo "<p>Fehler: " . $sql . "<br>" . $conn->error . "</p>";
    }
}

// Benutzerliste anzeigen
$sql = "SELECT id, name, email FROM users";
$result = $conn->query($sql);
?>

<!DOCTYPE html>
<html lang="de">
<head>
    <meta charset="UTF-8">
    <title>Benutzerverwaltung</title>
</head>
<body>
    <h1>Benutzer hinzufügen</h1>
    <form action="index.php" method="post">
        <label for="name">Name:</label>
        <input type="text" id="name" name="name" required>
        <br>
        <label for="email">E-Mail:</label>
        <input type="email" id="email" name="email" required>
        <br>
        <input type="submit" value="Hinzufügen">
    </form>

    <h1>Benutzerliste</h1>
    <table border='1'>
    <tr>
        <th>ID</th>
        <th>Name</th>
        <th>E-Mail</th>
    </tr>

    <?php
    if ($result->num_rows > 0) {
        // Daten aus der Datenbank abrufen und in einer Tabelle anzeigen
        while($row = $result->fetch_assoc()) {
            echo "<tr>
            <td>" . $row["id"]. "</td>
            <td>" . $row["name"]. "</td>
            <td>" . $row["email"]. "</td>
            </tr>";
        }
    } else {
        echo "<tr><td colspan='3'>Keine Benutzer gefunden</td></tr>";
    }
    ?>
    </table>
</body>
</html>

<?php
// Verbindung schließen
$conn->close();
?>

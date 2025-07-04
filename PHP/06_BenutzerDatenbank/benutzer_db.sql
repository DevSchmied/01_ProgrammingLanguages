-- Erstellen der Datenbank benutzer_db
CREATE DATABASE benutzer_db;

-- Auswahl der Datenbank benutzer_db
USE benutzer_db;

-- Erstellen der Tabelle users
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL
);

package aufgaben.xml_deserialisierung;

import com.fasterxml.jackson.dataformat.xml.XmlMapper;

import java.io.File;
import java.io.IOException;

/**
 * Die Klasse JacksonXmlExampleFromFile demonstriert das Einlesen und Deserialisieren
 * einer XML-Datei in Java unter Verwendung der Jackson-Bibliothek.
 * Die XML-Daten werden in ein People-Objekt konvertiert, das eine Liste von Person-Objekten enthält.
 */
public class JacksonXmlExampleFromFile {
    public static void main(String[] args) {
        // Erstellen einer XmlMapper-Instanz zum Lesen von XML-Dateien
        XmlMapper xmlMapper = new XmlMapper();

        try {
            // Lesen und Deserialisieren der XML-Datei in ein People-Objekt
            People people = xmlMapper.readValue(new File("aufgaben/xml_deserialisierung/resources/people.xml"), People.class);

            // Ausgabe des Namens der ersten Person in der Liste
            System.out.println(people.getPersonen().get(0).getName());

            // Durchlaufen der Liste von Personen und Ausgabe ihrer Details
            for (Person p : people.getPersonen()) {
                System.out.println(p.getName());
                System.out.println(p.getAlter());
                System.out.println(p.getEmail());
            }
        } catch (IOException e) {
            // Wenn ein Fehler auftritt, wird eine RuntimeException ausgelöst
            throw new RuntimeException(e);
        }
    }
}

package aufgaben.xml_deserialisierung;

import com.fasterxml.jackson.dataformat.xml.annotation.JacksonXmlElementWrapper;
import com.fasterxml.jackson.dataformat.xml.annotation.JacksonXmlProperty;

import java.util.ArrayList;
import java.util.List;

/**
 * Die Klasse People repräsentiert eine Sammlung von Person-Objekten.
 * Diese Klasse wird für die Deserialisierung von XML-Daten mit der Jackson-Bibliothek verwendet.
 */
public class People {
    @JacksonXmlProperty(localName = "person")
    @JacksonXmlElementWrapper(useWrapping = false)
    List<Person> personen = new ArrayList<>(); // Liste von Person-Objekten

    /**
     * Standardkonstruktor.
     * Wird benötigt für die Deserialisierung durch die Jackson-Bibliothek.
     */
    public People() {}

    /**
     * Konstruktor mit Parametern.
     *
     * @param personen Liste von Person-Objekten
     */
    public People(List<Person> personen) {
        this.personen = personen;
    }

    /**
     * Gibt die Liste der Person-Objekte zurück.
     *
     * @return Liste der Person-Objekte
     */
    public List<Person> getPersonen() {
        return personen;
    }

    /**
     * Setzt die Liste der Person-Objekte.
     *
     * @param personen die neue Liste der Person-Objekte
     */
    public void setPersonen(List<Person> personen) {
        this.personen = personen;
    }
}

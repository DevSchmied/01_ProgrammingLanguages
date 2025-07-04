package aufgaben.xml_deserialisierung;

import com.fasterxml.jackson.dataformat.xml.annotation.JacksonXmlProperty;

/**
 * Die Klasse Person repräsentiert eine Person mit den Attributen Name, Alter und E-Mail.
 * Diese Klasse wird für die Deserialisierung von XML-Daten mit der Jackson-Bibliothek verwendet.
 */
public class Person {
    @JacksonXmlProperty(localName = "name")
    private String name; // Der Name der Person
    @JacksonXmlProperty(localName = "alter")
    private int alter; // Das Alter der Person
    @JacksonXmlProperty(localName = "email")
    private String email; // Die E-Mail-Adresse der Person

    /**
     * Standardkonstruktor.
     * Wird benötigt für die Deserialisierung durch die Jackson-Bibliothek.
     */
    public Person() {}

    /**
     * Konstruktor mit Parametern.
     *
     * @param name der Name der Person
     * @param alter das Alter der Person
     * @param email die E-Mail-Adresse der Person
     */
    public Person(String name, int alter, String email) {
        this.name = name;
        this.alter = alter;
        this.email = email;
    }

    /**
     * Gibt den Namen der Person zurück.
     *
     * @return der Name der Person
     */
    public String getName() {
        return name;
    }

    /**
     * Setzt den Namen der Person.
     *
     * @param name der neue Name der Person
     */
    public void setName(String name) {
        this.name = name;
    }

    /**
     * Gibt das Alter der Person zurück.
     *
     * @return das Alter der Person
     */
    public int getAlter() {
        return alter;
    }

    /**
     * Setzt das Alter der Person.
     *
     * @param alter das neue Alter der Person
     */
    public void setAlter(int alter) {
        this.alter = alter;
    }

    /**
     * Gibt die E-Mail-Adresse der Person zurück.
     *
     * @return die E-Mail-Adresse der Person
     */
    public String getEmail() {
        return email;
    }

    /**
     * Setzt die E-Mail-Adresse der Person.
     *
     * @param email die neue E-Mail-Adresse der Person
     */
    public void setEmail(String email) {
        this.email = email;
    }
}

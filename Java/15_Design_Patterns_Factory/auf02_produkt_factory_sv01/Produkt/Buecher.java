package aufgaben.auf02_produkt_factory_sv01.Produkt;

public class Buecher extends Produkt {
    @Override
    public void displayInfo() {
        System.out.println("Hier ist ein Produkt von dem Typ: " + this.getClass().getSimpleName());
    }
}

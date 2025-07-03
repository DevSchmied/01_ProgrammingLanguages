#include <stdio.h>
#include <windows.h>

int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    int anzahlPersonen;
    float ticketPreis, gesamtpreis;
    float rabatt = 0.9;

    printf("Geben Sie bitte den Einzelpreis pro Ticket ein (in Euro): ");
    scanf("%f", &ticketPreis);

    printf("Geben Sie bitte die Anzahl der Personen ein: ");
    scanf("%d", &anzahlPersonen);

    if (anzahlPersonen <= 4)
    {
        gesamtpreis = anzahlPersonen * ticketPreis;
    } else
    {
        gesamtpreis = anzahlPersonen * ticketPreis * rabatt;
    }

    printf("Gesamtpreis: %.2f Euro", gesamtpreis);
    
    

    printf("\n\n");
    return 0;
}

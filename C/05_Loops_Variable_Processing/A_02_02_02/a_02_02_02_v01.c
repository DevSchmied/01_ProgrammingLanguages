#include <stdio.h>
#include <windows.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    int anzahlRaeume;
    float breite, laenge, flaeche, gesamtFlaeche;

    printf("Geben Sie die Anzahl der R�ume ein: ");
    fflush(stdin);
    scanf("%d", &anzahlRaeume);

    for (int i = 1; i <= anzahlRaeume; i++)
    {
        printf("\nGeben Sie die Breite des Raumes %d ein: ", i);
        fflush(stdin);
        scanf("%f", &breite);
    
        printf("\nGeben Sie die L�nge des Raumes %d ein: ", i);
        fflush(stdin);
        scanf("%f", &laenge);

        flaeche = breite * laenge;
        gesamtFlaeche += flaeche;
    }
    
    printf("\n\nDas ist die Gesamtfl�che der Wohnung: %.2f m2", gesamtFlaeche);

    printf("\n\n");
    return 0;
}

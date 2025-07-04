#include <stdio.h>
#include <windows.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");
    
    float startKapital, zielKapital, aktuellesKapital, prozentSatz;
    
    printf("Geben Sie Bitte die Summe des Startkapitals (in Euro) ein: ");
    fflush(stdin);
    scanf("%f", &startKapital);
    
    printf("Geben Sie Bitte die Summe des Zielkapitals (in Euro) ein: ");
    fflush(stdin);
    scanf("%f", &zielKapital);
    
    printf("Geben Sie Bitte die Gr��e des Prozentsatzes (in Prozent) ein: ");
    fflush(stdin);
    scanf("%f", &prozentSatz);

    aktuellesKapital = startKapital;
    int jahr = 0;

    while (aktuellesKapital < zielKapital)
    {
        aktuellesKapital = aktuellesKapital * ((prozentSatz + 100) / 100);
        jahr ++;
    }

    printf("Nach %d Jahren betr�gt Ihr Kapital %.2f Euro", jahr,aktuellesKapital);
    

    printf("\n\n");
    return 0;
}

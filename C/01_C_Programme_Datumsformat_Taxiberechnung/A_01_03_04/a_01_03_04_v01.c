#include <stdio.h>
#include <windows.h>

int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    float grundgebuehr, preisProKilometer, preisZuBezahlen;
    int anzahlDerGefahrenenKilometer;

    printf("Geben Sie bitte Grundgeb�hr in Euro: ");
    fflush(stdin);
    scanf("%f", &grundgebuehr); 
    

    printf("Geben Sie bitte Anzahl der gefahrenen Kilometer: ");
    fflush(stdin);
    scanf("%d", &anzahlDerGefahrenenKilometer); 
    

    printf("Geben Sie bitte den Preis pro Kilometer in Euro: ");
    fflush(stdin);
    scanf("%f", &preisProKilometer);  

    preisZuBezahlen = grundgebuehr + anzahlDerGefahrenenKilometer * preisProKilometer;
    
    printf("Bei einer Grundgeb�hr von %.2f Euro und einem Kilometerpreis von %.2f Euro, \nkostet eine %d km lange Taxifahrt insgesamt %.2f Euro. ", grundgebuehr, preisProKilometer, anzahlDerGefahrenenKilometer, preisZuBezahlen);

    printf("\n\n");
    return 0;
}

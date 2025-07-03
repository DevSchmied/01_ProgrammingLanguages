#include <stdio.h>
#include <windows.h>

int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    float gesamtpunktzahl, erreichtePunktzahl, prozentsatz;

    printf("Geben Sie bitte die maximal erreichbare Punktzahl ein: ");
    fflush(stdin);
    scanf("%f", &gesamtpunktzahl);

    printf("Geben Sie bitte die erreichte Punktzahl ein: ");
    fflush(stdin);
    scanf("%f", &erreichtePunktzahl);

    prozentsatz = erreichtePunktzahl / gesamtpunktzahl * 100;
    
    if (prozentsatz >= 50)
    {
        printf("Sie haben %.2f Prozent erreicht und bestanden", prozentsatz);
    } else
    {
        printf("Sie haben %.2f Prozent erreicht und NICHT bestanden", prozentsatz);
    }
    
    printf("\n\n");
    return 0;
}

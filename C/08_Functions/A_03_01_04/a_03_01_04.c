#include <stdio.h>
#include <windows.h>

void gesamtpreis(float einzelpreis, float rabatt, int stueckzahl)
{
    printf("Der Gesamtpreis bel�uft sich auf: %.2f.", einzelpreis * stueckzahl * (1.00 - (rabatt * 0.01)));
    return;
}


 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    float einzelpreis, rabatt;
    int stueckzahl;

    printf("Geben Sie den Einzelpreis der Ware ein: ");
    fflush(stdin);
    scanf("%f", &einzelpreis);

    printf("Geben Sie den Rabatt als Prozent f�r die Ware ein: ");
    fflush(stdin);
    scanf("%f", &rabatt);

    printf("Geben Sie die St�ckzahl der Ware ein: ");
    fflush(stdin);
    scanf("%d", &stueckzahl);

    gesamtpreis(einzelpreis, rabatt, stueckzahl);
    
    printf("\n\n");
    return 0;
}


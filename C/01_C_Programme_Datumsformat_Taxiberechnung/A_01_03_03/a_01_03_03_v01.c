#include <stdio.h>
#include <windows.h>

int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");
    
    int tag, monat, jahr;

    printf("Geben Sie bitte den Tag ein: ");
    fflush(stdin);
    scanf("%d", &tag);

    printf("Geben Sie bitte den Monat ein: ");
    fflush(stdin);
    scanf("%d", &monat);

    printf("Geben Sie bitte das Jahr ein: ");
    fflush(stdin);
    scanf("%d", &jahr);

    printf("Deutsche Datumsdarstellung: %.2d.%.2d.%.4d\n", tag, monat, jahr);
    printf("Amerikanische Datumsdarstellung: %.4d-%.2d-%.2d",jahr,monat,tag);

    printf("\n\n");
    return 0;

}
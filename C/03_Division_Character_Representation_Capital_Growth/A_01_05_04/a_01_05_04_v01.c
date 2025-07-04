#include <stdio.h>
#include <windows.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    float x, y = 0, ergebnis;

    printf("\n\n");
    printf("\tGeben Sie eine erste Kommazahl ein: ");
    fflush(stdin);
    scanf("%f", &x);

    while (y == 0)
    {
        printf("\tGeben Sie eine zweite Kommazahl ein: ");
        fflush(stdin);
        scanf("%f", &y);

        if (y == 0)
        {
            printf("\n\tDurch 0 teilen ist nicht zul�ssig\n");
        }
    }

    ergebnis = x / y;
    printf("\n\t%f", ergebnis);
    

    printf("\n\n");
    return 0;
}

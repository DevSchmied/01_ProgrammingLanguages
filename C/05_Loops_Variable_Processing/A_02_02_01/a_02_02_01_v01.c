#include <stdio.h>
#include <windows.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    int anzahl, zahl, min, max;

    printf("Geben Sie Anzahl ein: ");
    fflush(stdin);
    scanf("%d", &anzahl);

    printf("Geben Sie zahl ein: ");
    fflush(stdin);
    scanf("%d", &zahl);

    min = zahl;
    max = zahl;

    for (int i = 2; i <= anzahl; i++)
    {

        printf("Geben Sie zahl ein: ");
        fflush(stdin);
        scanf("%d", &zahl);

        if (zahl < min)
        {
            min = zahl;
        }

        if (zahl > max)
        {
            max = zahl;
        }    
    }

    printf("min ist: %d und max ist: %d\n", min, max);    

    printf("\n\n");
    return 0;
}

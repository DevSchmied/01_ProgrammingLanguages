#include <stdio.h>
#include <windows.h>
#include <time.h>

int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    srand(time(NULL));
    rand();

    int zuf1 = (rand()%(100 - 1 + 1) + 1) + (rand()%(100 - 1 + 1) + 1 - 100);    
    int zuf2 = (rand()%(100 - 1 + 1) + 1) + (rand()%(100 - 1 + 1) + 1 - 100);
    int sumVomFleisch, sumVomRoboter = zuf1 + zuf2;

    printf("Die erste zuf�llige Zahl: %3d, die zweite Zahl: %3d.\nBerechnen Sie und geben Sie die Summe ein: ", zuf1, zuf2);
    fflush(stdin);
    scanf("%d", &sumVomFleisch);

    if (sumVomRoboter == sumVomFleisch)
        printf("Gl�ckwunsch!!!");
    else
        printf("Fehlermeldung: du bist nicht so gut wie ich, Freundchen. Das korrekte Ergebnis ist: %3d.", sumVomRoboter);

    printf("\n\n");
    return 0;
}

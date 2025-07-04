#include <stdio.h>
#include <windows.h>
#include <time.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    int min, max, anzahl;

    printf("Geben Sie die untere Grenze des Bereiches von zuf�lligen Zahlen ein: ");
    fflush(stdin);
    scanf("%d", &min);    

    printf("Geben Sie die obere Grenze des Bereiches von zuf�lligen Zahlen ein: ");
    fflush(stdin);
    scanf("%d", &max);    

    printf("Geben Sie die Anzahl der Wiederholungen ein: ");
    fflush(stdin);
    scanf("%d", &anzahl);    

    srand(time(NULL));
    rand();

    printf("\n\n");
    
    for (int i = 0; i < anzahl; i++)
    {
        printf("Zuf�llige Zahl %2d: %3d\n\n", i+1, rand()%(max - min + 1) + min);
        Sleep(1500);
    }
    

    printf("\n\n");
    return 0;
}

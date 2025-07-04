#include <stdio.h>
#include <windows.h>
#include <time.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    int zuf1, zuf2, zaehler = 0;

    srand(time(NULL));
    rand();

    for (int i = 0; i < 6000; i++)
    {
        zuf1 = rand()%6 + 1;
        zuf2 = rand()%6 + 1;

        if (zuf1 == zuf2)
        {
            zaehler += 1;
        }
    }
    
    printf("\n\n\tSie hatten %3d Mal die identischen W�rfel-Ergebnisse", zaehler);

    printf("\n\n");
    return 0;
}

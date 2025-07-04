#include <stdio.h>
#include <windows.h>
#include <time.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    int zuf, max = 0;
    srand(time(NULL));
    rand();

    for (int i = 0; i < 10; i++)
    {
        zuf = rand()%(100 - 1 + 1) + 1; // zuf = rand()%100 + 1;
        printf("%d\n", zuf);
        if (max < zuf)
        {
            max = zuf;
        }
    }
    
    printf("\nDer Maximumwert ist: %3d", max);

    printf("\n\n");
    return 0;
}

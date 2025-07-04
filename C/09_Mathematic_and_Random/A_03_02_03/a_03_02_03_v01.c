#include <stdio.h>
#include <windows.h>
#include <time.h>

int zufall(int a, int b)
{
    return rand() % (b - a + 1) + a; 
}

int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    int x = 1;
    int b = 100;

    srand(time(NULL));
    rand();

    do
    {
        x = zufall(x,b);
        printf("x = %d\n", x);
    } while (x != 100);
    

    printf("\n\n");
    return 0;
}

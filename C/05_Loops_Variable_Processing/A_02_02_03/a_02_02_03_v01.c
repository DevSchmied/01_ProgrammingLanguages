#include <stdio.h>
#include <windows.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    int ausnahme;

    printf("Geben Sie eine Zahl, die nicht angezeigt wird ein: ");
    fflush(stdin);
    scanf("%d", &ausnahme);

    for (int i = 1; i <= 100; i++)
    {
        if (i != ausnahme)
        {
            printf("%3d\n", i);
        }
        
    }
    

    printf("\n\n");
    return 0;
}

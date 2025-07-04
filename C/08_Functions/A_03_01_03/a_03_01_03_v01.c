#include <stdio.h>
#include <windows.h>
 
    void bewertung(int x)
    {
        if (x == 0)
        {
            printf("Kein Kommentar");
        } else if (x > 0)
        {
            printf("ANSI C ist ");
            for (int i = 0; i < x; i++)
            {
                printf("sehr ");
            }
            printf("toll!\n");
        }
        
        return;
    }

int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    int x;

    printf("Geben Sie einen ganzahligen Wert zwischen 0 und 5 ein: ");
    fflush(stdin);
    scanf("%d", &x);

    bewertung(x);

    printf("Ende des Programms!");

    printf("\n\n");
    return 0;
}

#include <stdio.h>
#include <windows.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    char zeichen[5];

    for (int i = 0; i < 5; i++)
    {
        printf("Geben Sie einen Buchstaben ein: ");
        fflush(stdin);
        scanf("%c", &zeichen[i]);
    }

    char zeichen2[5];

    for (int i = 0; i < 5; i++)
    {
        zeichen2[4 - i] = zeichen[i];
    }

    printf("\n");

    for (int i = 0; i < 5; i++)
    {
        printf("der Array 1: %c\n",zeichen[i]);
    }

    printf("\n");

    for (int i = 0; i < 5; i++)
    {
        printf("Der Array 2: %c\n",zeichen2[i]);
    }
    
    

    printf("\n\n");
    return 0;
}

#include <stdio.h>
#include <windows.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    char c;
    int i, auswahl, start = 0;

    printf("Geben Sie bitte ein Zeichen ein: ");
    fflush(stdin);
    scanf("%c", &c);

    printf("Geben Sie bitte eine positive ganze Zahl ein: ");
    fflush(stdin);
    scanf("%d", &i);

    printf("Dr�cken Sie 1, wenn Sie wollen, dass die Zeichen nebeneinander angezeigt w�rden und 2 - wenn �bereinander: ");
    fflush(stdin);
    scanf("%d", &auswahl);

    printf("\n\t");

    if (auswahl == 1)
    {
        while (start < i)
        {
            printf("%c", c);
            start += 1;
        }   
    }

    printf("\n");

    if (auswahl == 2)
    {
        while (start < i)
        {
            printf("\n\t%c", c);
            start += 1;
        }
        
    }
    
    

    printf("\n\n");
    return 0;
}

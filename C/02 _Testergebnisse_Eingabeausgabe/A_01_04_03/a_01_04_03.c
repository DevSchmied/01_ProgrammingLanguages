#include <stdio.h>
#include <windows.h>

int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    char zeichen1, zeichen2;
    int auswahl;

    printf("Geben Sie bitte den ersten Buchstaben: ");
    fflush(stdin);
    scanf("%c", &zeichen1);

    printf("Geben Sie bitte den zweiten Buchstaben: ");
    fflush(stdin);
    scanf("%c", &zeichen2);

    printf("Wenn Sie wollen, dass die beiden Buchstaben nebeneinander angezeigt w�rden, dr�cken Sie 1, wenn �bereinander - 2    ");
    fflush(stdin);
    scanf("%d", &auswahl);

    if (auswahl == 1)
    {
        printf("%c%c", zeichen1, zeichen2);
    } if (auswahl == 2)
    {
        printf("%c\n%c", zeichen1, zeichen2);
    }
    
    printf("\n\n");
    return 0;
}

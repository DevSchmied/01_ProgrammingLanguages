#include <stdio.h>
#include <windows.h>

int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    int a = 1;
    char zeichen1, zeichen2, zeichen3;

    printf("Geben Sie bitte das “ a „. Zeichen ein: ");
    fflush(stdin);
    scanf("%c", &zeichen1);

    a =+ 1;
    
    printf("Geben Sie bitte das “ a „. Zeichen ein: ");
    fflush(stdin);
    scanf("%c", &zeichen2);

    a =+ 1;
    
    printf("Geben Sie bitte das “ a „. Zeichen ein: ");
    fflush(stdin);
    scanf("%c", &zeichen3);

    printf("%c\n%c\n%c\n", zeichen1, zeichen2, zeichen3);



}

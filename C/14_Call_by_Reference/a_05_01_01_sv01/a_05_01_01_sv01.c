#include <stdio.h>

inputINT(char text[], int *p)
{
    printf("%s", text);

    fflush(stdin);

    scanf("%p", p);
}

int main(void)
{
    system("chcp 1252");
    system("cls");

    int x;

    inputINT("Geben Sie eine ganze Zahl ein: ", &x);

    printf("Kontrollausgabe von der Zahl: %d", x);

    return 0;
}

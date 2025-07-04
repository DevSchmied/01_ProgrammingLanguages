#include <stdio.h>
#include <windows.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    int zaehler = -1, zahlEingabe;
    printf("\n");
    do
    {
        printf("\nGeben Sie ein, wie viel mal das bisher gefragt wurde: \n\n");
        fflush(stdin);
        scanf("%d", &zahlEingabe);
        zaehler += 1;
    } while (zahlEingabe != zaehler);
    
    printf("\nF�r so etwas brauchen Sie, Freundchen, %d Versuch(e).\n\n", zaehler);

    printf("\n\n");
    return 0;
}

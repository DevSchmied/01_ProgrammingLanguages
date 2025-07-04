#include <stdio.h>
#include <windows.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    int breite, laenge, produkt, versuch = 0;
    printf("\n\n");
    printf("\tGeben Sie die Breite eines Rechteckes: \n");
    fflush(stdin);
    scanf("%d", &breite);
    
    printf("\tGeben Sie die L�nge eines Rechteckes: \n");
    fflush(stdin);
    scanf("%d", &laenge);

    do
    {
        printf("\tGeben Sie das Prokukt eines Rechteckes mit der Breite: %d und L�nge: %d\n", breite, laenge);
        fflush(stdin);
        scanf("%d", &produkt);

        versuch += 1;
    } while (produkt != laenge*breite);
    
    printf("\n\tUm so eine einfache Aufgabe zu l�sen, brauchten Sie %d Versuch(e)", versuch);

    printf("\n\n");
    return 0;
}

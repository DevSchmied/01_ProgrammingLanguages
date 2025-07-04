#include<windows.h>
#include<stdio.h>
#include<string.h>
#include "a_06_02_01_sv01.h"

int main(void) {
    system("chcp 1252");
    system("cls");

    int laenge;

    printf("Geben Sie die L�nge eines Char-Arrays ein: ");
    fflush(stdin);
    scanf("%d", &laenge);

    char arr[laenge];
    printf("\nGeben Sie eine Zeile nicht l�nger als L�nge(-1) ein: ");
    fflush(stdin);
    scanf("%[^\n]", arr);


    aufruf(laenge, arr);



    return 0;
}

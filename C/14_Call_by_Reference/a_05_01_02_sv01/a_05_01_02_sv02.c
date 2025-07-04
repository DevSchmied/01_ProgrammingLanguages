#include<windows.h>
#include<stdio.h>

void tausche(int *p1, int *p2){
    int tmp;

    tmp = *p1;
    *p1 = *p2;
    *p2 = tmp;
}


int main(void){
    system("chcp 1252");
    system("cls");

    int a,b;

    printf("Geben Sie die Zahl a ein: ");
    fflush(stdin);
    scanf("%d", &a);
    printf("\n\n");

    printf("Geben Sie die Zahl b ein: ");
    fflush(stdin);
    scanf("%d", &b);
    printf("\n\n");

    tausche(&a, &b);

    printf("Kontrolausgabe a: %d", a);
    printf("\n\n");

    printf("Kontrolausgabe b: %d", b);
    printf("\n\n");



    return 0;
}

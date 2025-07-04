#include<windows.h>
#include<stdio.h>


tausche(int *p1, int *p2){
    int p3 = *p2;
    *p2 = *p1;
    *p1 = p3;
}

int main(void)
{
    system("chcp 1252");
    system("cls");

    int zahl1, zahl2;

    printf("Geben Sie die 1. Zahl ein: ");
    fflush(stdin);
    scanf("%d",&zahl1);
    printf("\n\n");

    printf("Geben Sie die 2. Zahl ein: ");
    fflush(stdin);
    scanf("%d",&zahl2);
    printf("\n\n");

    tausche(&zahl1, &zahl2);

    printf("Zahl1: %d\nZahl2: %d\n",zahl1, zahl2);

    return 0;
}

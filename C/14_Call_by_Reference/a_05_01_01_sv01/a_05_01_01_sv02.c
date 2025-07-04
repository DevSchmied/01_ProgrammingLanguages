#include<windows.h>
#include<stdio.h>

inputINT(char arr[], int *p){
    printf("%s\n",arr);
    fflush(stdin);
    scanf("%p", p);
}


int main(void){
    system("chcp 1252");
    system("cls");

    int x = 0;
    char text[101];

    strcpy(text, "Bitte geben Sie eine ganze Zahl ein: ");

    inputINT(text, &x);

    printf("%d", x);


    return 0;
}

#include<windows.h>
#include<string.h>
#include<stdio.h>

int main(void) {

    system("chcp 1252");
    system("cls");

    char zeile[100];
    char zeile2[100];

    printf("Geben Sie eine Zeile ein: ");
    fflush(stdin);
    fgets(zeile, 99, stdin);

    for(int i=0; i<strlen(zeile); i++){
        if(zeile[i]>='a' && zeile[i]<='z'){
            zeile2[i] = zeile[i]-32;
        } else {
            zeile2[i] = zeile[i];
        }
    }

    printf("%s",zeile2);


    return 0;
}

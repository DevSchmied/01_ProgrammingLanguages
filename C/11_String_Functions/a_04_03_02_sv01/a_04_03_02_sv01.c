#include<windows.h>
#include<stdio.h>
#include<string.h>

int main(void) {

    system("chcp.com 1252");
    system("cls");

    char wort[21];
    int x;

    do {
        printf("Geben Sie ein Wort ein: ");
        fflush(stdin);
        scanf("%20s", wort);

        x = strlen(wort);

    } while(wort[0] != wort[x-1] && wort[0]+32 != wort[x-1]);

    printf("wunderbar");
}

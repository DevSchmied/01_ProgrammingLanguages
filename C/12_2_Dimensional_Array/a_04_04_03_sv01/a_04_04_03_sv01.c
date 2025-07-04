
#include<stdio.h>
#include<windows.h>
#include<string.h>

int main(void) {

    system("chcp.com 1252");
    system("cls");

    int x;
    char wort[5][21];

    for(int i=0; i<5; i++){
        printf("Geben Sie das %d. Wort ein: ", i+1);
        fflush(stdin);
        scanf("%s", wort[i]);
    }

    do {
        printf("Geben Sie eine Zahl von 1 bis 5 ein: \n");
        fflush(stdin);
        scanf("%d",&x);

        if(x<1 || x>5)
            printf("Kannst du nicht bis 5 z�hlen?\n");
    } while(x<1 || x>5);

    printf("Das %d. Wort: %s",x, wort[x-1]);




    return 0;
}

#include<string.h>
#include<stdio.h>
#include<windows.h>

int main(void) {

    system("chcp.com 1252");
    system("cls");

    int dublette = 0;
    char wort[5][10];

    for(int i=0; i<5; i++) {
        printf("Geben Sie ein %d Wort ein: ", i+1);
        fgets(wort[i], 9, stdin);
//        fflush(stdin);
//        scanf("[^\n]", wort[i-1]);

        for(int j=0; j<i && dublette==0; j++){
            if(strcmp(wort[i],wort[j]) == 0){
                dublette = 1;
                break;
            }
        }
    }

    if(dublette == 1){
        printf("Mensch...");
    } else {
        printf("Ok");
    }


    return 0;
}

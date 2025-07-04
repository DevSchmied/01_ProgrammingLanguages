
#include<stdio.h>
#include<windows.h>
#include<time.h>

int main(void) {

    system("chcp.com 1252");
    system("cls");

    srand(time(NULL));
    rand();

    int lotto[10][6], dublette, zuf, j;

    for(int i=0; i<10; i++){
        j=0;
        while(j<6){
            dublette = 0;
            zuf = rand()%99 +1;

            for(int k=0; k<j; k++){
                if(zuf == lotto[i][k]){
                    dublette = 1;
                    break;
                }
            }

            if(dublette == 0) {
                lotto[i][j] = zuf;
                j++;
            }
        }
    }


    for(int i=0; i<10; i++){
        printf("%.2d. Tipp: ", i+1);
        for(int j=0; j<6; j++){
            printf("%2d ",lotto[i][j]);
        }
        printf("\n");
    }


    return 0;
}

#include<windows.h>
#include<stdio.h>
#include<time.h>

int main(void) {
    system("chcp 1252");
    system("cls");

    int zahl[11];
    int hashwert = 0;

    for(int i=0; i<11; i++){
        printf("Geben Sie die %2d Zahl zwischen 0 und 9: ",i+1);
        fflush(stdin);
        scanf("%d",&zahl[i]);
        printf("\n");

        if(zahl[i]<0 || zahl[i]>9) {
            printf("Falsche Eingabe!!!\n");
            i--;
        }
    }

    hashwert = zahl[0];
    for(int i=1; i<10; i++){
        if(i%2 == 0){
            hashwert += zahl[i];
        } else {
            hashwert -= zahl[i];
        }
    }

//    hashwert = (zahl[0]-zahl[1]+zahl[2]-zahl[3]+zahl[4]-zahl[5]+zahl[6]-zahl[7]+zahl[8]-zahl[9]) % 10;

    if(hashwert == zahl[10]){
        printf("Alles O.K.!\n");
    } else {
        printf("leider unzul�ssig\n");
    }

    return 0;
}

#include<windows.h>
#include<stdio.h>

int main(void){
    system("chcp.com 1252");
    system("cls");

    int prozentsatz;

    printf("Geben Sie den Prozentsatz von 0 bis 100 ein: ");
    fflush(stdin);
    scanf("%d", &prozentsatz);

    if (prozentsatz >= 92 && prozentsatz <= 100){
        printf("sehr gut");
    } else if(prozentsatz >= 81 && prozentsatz <= 91) {
        printf("gut");
    } else if(prozentsatz >= 67 && prozentsatz <= 80) {
        printf("befriedigend");
    } else if(prozentsatz >= 50 && prozentsatz <= 66) {
        printf("ausreichend");
    } else if(prozentsatz >= 30 && prozentsatz <= 49) {
        printf("mangelhaft");
    } else {
        printf("ungen�gend");
    }
    return 0;
}

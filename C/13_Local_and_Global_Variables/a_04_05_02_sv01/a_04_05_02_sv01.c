#include<windows.h>
#include<stdio.h>

int main(void) {

    int zahl[10];

    for(int i=0; i<10; i++){
        printf("Geben Sie %d. eine einstellige Zahl ein: ",i+1);
        fflush(stdin);
        scanf("%d",&zahl[i]);
        printf("\n\n");
    }

    int hash = zahl[0];
    int i = 1;
    do{
        hash = (hash*zahl[i])%10;
        i++;
    } while(i<10);

    printf("der Hashwert: %d.", hash);

    return 0;
}

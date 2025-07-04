#include<stdio.h>
#include<windows.h>
#include<time.h>
#define Z1 3
#define Z2 6

int main(void) {

    system("chcp.com 1252");
    system("cls");

    srand(time(NULL));
    rand();

    int arr[Z1][Z2], summe;

    for(int i=0; i<Z1; i++) {
        for(int j=0; j<Z2; j++) {
            arr[i][j] = rand()%10 + 1;
        }

        summe = 0;

        for(int j=0; j<Z2; j++){
            summe += arr[i][j];
        }
        arr[i][0] = summe;
        printf("%d\n\n", arr[i][0]);
    }

    return 0;
}

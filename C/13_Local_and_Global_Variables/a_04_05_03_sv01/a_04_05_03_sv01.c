#include<windows.h>
#include<stdio.h>
#include<time.h>

int x;

int zufall(void){
    x = rand()%10 + 1;
//    return x;
}

int main(void) {
    srand(time(NULL));
    rand();

    for(int i=0; i<5; i++){
        zufall();
        printf("Der Wert der globalen Variable x: %d",x);
        printf("\n");
    }

    return 0;
}

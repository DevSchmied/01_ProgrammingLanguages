#include<windows.h>
#include<stdio.h>
#include<time.h>

int main(void){
    system("chcp.com 1252");
    system("cls");
    srand(time(NULL));
    rand();

    int x = rand()%6 + 1;

    if(x > 3) {
        printf("�berdurschnitlich");
    } else {
        switch(x){
            case 1: printf("Oh Mann!"); break;
            case 2: printf("Naja!");    break;
            case 3: printf("O.K.");     break;

            default: printf("H�h??");
        }
    }
    return 0;
}

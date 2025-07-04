#include<stdio.h>
#include<windows.h>
#include<string.h>
#include<time.h>

typedef struct
{
    int vorkomma;
    int nachkomma;
} ZUFALL;

void initZUFALL(ZUFALL *z){
    (*z).vorkomma = rand()%1000;
    (*z).nachkomma = rand()%1000;
}

int main(void) {

    srand(time(NULL));
    rand();

    ZUFALL z;

    initZUFALL(&z);

    printf("%d,%.3d", z.vorkomma, z.nachkomma);


    return 0;
}

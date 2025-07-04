#include<stdio.h>

int global = 1;

void erhoeheUmX(int x) {
    global = global + x;
}

void output(void){
    printf("Globale Variable: %d\n", global);
}

void main(void) {
    int x = 12;

    erhoeheUmX(x);
    output();
}

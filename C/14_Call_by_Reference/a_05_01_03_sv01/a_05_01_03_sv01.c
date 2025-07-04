#include<windows.h>
#include<stdio.h>

void rechne(int a, int b, int array[]) {
    array[0] = a+b;
    array[1] = a-b;
    array[2] = a*b;
    array[3] = a/b;
}

int main(void){
    system("chcp 1252");
    system("cls");

    int a,b;
    int arr[4];

    printf("a-Eingabe: ");
    fflush(stdin);
    scanf("%d", &a);
    printf("\n\n");

    printf("b-Eingabe: ");
    fflush(stdin);
    scanf("%d", &b);
    printf("\n\n");

    rechne(a,b,arr);

    for(int i=0; i<4; i++){
        printf("%d\n", arr[i]);
    }

    return 0;
}

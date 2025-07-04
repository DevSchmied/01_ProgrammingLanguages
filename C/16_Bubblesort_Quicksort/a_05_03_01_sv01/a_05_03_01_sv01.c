#include<windows.h>
#include<stdio.h>
#include<string.h>
#define DIM1 5
#define DIM2 50

bubblesort_string(int dim1, int dim2, char arr[][dim2]){
    char arrTmp[dim2];


    for(int i=dim1;i>1;i--){
        for(int j=0; j<i-1; j++){
            if(strcmp(arr[j],arr[j+1]) == 1){
                strcpy(arrTmp, arr[j]);
                strcpy(arr[j], arr[j+1]);
                strcpy(arr[j+1], arrTmp);
            }
        }
    }
}

int main(void) {
    system("chcp 1252");
    system("cls");

    char arrString[DIM1][DIM2];

    for(int i=0; i<DIM1; i++){
        printf("Geben Sie eine %d. Zeile ein: ", i+1);
        fflush(stdin);
        scanf("%[^\n]",arrString[i]);
    }

    bubblesort_string(DIM1, DIM2, arrString);

    for(int i=0; i<DIM1; i++){
        printf("%s\n",arrString[i]);
    }

    return 0;
}

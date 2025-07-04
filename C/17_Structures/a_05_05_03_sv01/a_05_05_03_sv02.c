#include<stdio.h>
#include<windows.h>
#include<string.h>

typedef struct
{
    char text[101];
    int laenge;
} WORT;

void ermittelWortlaenge(WORT arr[]) {
    for (int i = 0; i < 3; i++) {
        arr[i].laenge = strlen(arr[i].text);
    }
}

int main(void) {
    WORT arr[3];

    for(int i=0; i<3; i++){
        printf("%d. Eingabe : ",i+1);
        fflush(stdin);
        scanf("%[^\n]", arr[i].text);
        printf("\n\n");
    }

    ermittelWortlaenge(arr);


    for(int i=0; i<3; i++) {
        printf("%d\n", arr[i].laenge);
    }




    return 0;
}

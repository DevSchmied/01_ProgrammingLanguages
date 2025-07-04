#include<stdio.h>
#include<windows.h>
#include<string.h>

typedef struct
{
    char text[101];
    int laenge;
} WORT;

ermittelWortlaenge(WORT *w){
    (*w).laenge = strlen((*w).text);
}

int main(void) {
    WORT arr[3];

    for(int i=0; i<3; i++){
        printf("%d. Eingabe : ",i+1);
        fflush(stdin);
        scanf("%[^\n]", arr[i].text);
        printf("\n\n");
    }

    for(int i=0; i<3; i++){
        ermittelWortlaenge(&arr[i]);
    }

    for(int i=0; i<3; i++) {
        printf("%d\n", arr[i].laenge);
    }

    return 0;
}

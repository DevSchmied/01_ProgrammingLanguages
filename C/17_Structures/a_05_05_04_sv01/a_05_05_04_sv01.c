#include<stdio.h>
#include<windows.h>
#include<string.h>

typedef struct
{
    int id;
    char nachname[101];
} KUNDE;

int SELECTnachnameWHEREid_x(int x, int i, KUNDE *p){
    if(x>0 && x<=i){
        printf("%s\n\n", (*p).nachname);
        return 1;
    } else {
        return 0;
    }
}

int main(void) {
    system("chcp 1252");
    system("cls");

    KUNDE arr[3];
    int i;

    for(i=0; i<3; i++) {
        printf("%d Nachname: ",i+1);
        fflush(stdin);
        scanf("%[^\n]", &arr[i].nachname);
        printf("\n\n");
    }

    int x;
    printf("x-Engabe: ");
    fflush(stdin);
    scanf("%d", &x);

    int treffer = SELECTnachnameWHEREid_x(x, i, &arr[x]);

    if(treffer == 1) printf("OK");
    else printf("Nicht OK");

    return 0;
}

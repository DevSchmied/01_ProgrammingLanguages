#include<stdio.h>
#include<windows.h>

int main(void){

    system("chcp.com 1252");
    system("cls");

    char zeichen1, zeichen2;

    printf("Gebe das Zeichen 1 ein: ");
    scanf("%c", &zeichen1);

    printf("Gebe das Zeichen 2 ein: ");
    fflush(stdin);
    scanf("%c", &zeichen2);

    if(zeichen1 == 'Q' && zeichen2 != 'Q') {
        printf("Fall 1");
    } else if(zeichen2 == 'Q' && zeichen1 != 'Q') {
        printf("Fall 2");
    } else {
        printf("Fall 3");
    }

    return 0;
}

#include<windows.h>
#include<string.h>
#include<stdio.h>

int main(void) {

    system("chcp 1252");
    system("cls");

    char s1[101],s2[101],s3[101];

    printf("Geben Sie einen %d. Satz ein (max \"100\" Zeichen): ",1);
    fflush(stdin);
    fgets(s1,100,stdin);
//    fflush(stdin);
//    scanf("%[\n]", s1);

    printf("Geben Sie einen %d. Satz ein (max \"100\" Zeichen): ",2);
    fflush(stdin);
    fgets(s2,100,stdin);

    if(strlen(s1)+strlen(s2)<101){
        strcpy(s3,s1);
        strcat(s3,s2);
        printf("%s", s3);
    } else {
        printf("Text ist zu lang");
    }

    return 0;
}

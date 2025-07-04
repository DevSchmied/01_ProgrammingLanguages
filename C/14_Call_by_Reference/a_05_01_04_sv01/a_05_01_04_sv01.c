#include<windows.h>
#include<stdio.h>

void codiere(char text[], char subtext1[], char subtext2[]){
    int hilfe = 0;
    for(int zaehler = 0; zaehler<100; zaehler = zaehler +2){
        subtext1[hilfe] = text[zaehler];
        hilfe = hilfe + 1;
    }

    hilfe = 0;
    for(int zaehler = 1; zaehler<100; zaehler = zaehler +2){
        subtext2[hilfe] = text[zaehler];
        hilfe = hilfe + 1;
    }

    printf("subtext1 nach Trennung: %s\n\n", subtext1);
    printf("subtext2 nach Trennung: %s\n\n", subtext2);
}

void decodiere(char subtext1[], char subtext2[], char text[]){
    int hilfe = 0;
    for(int zaehler = 0; zaehler<100; zaehler = zaehler +2){
        text[zaehler] = subtext1[hilfe];
        hilfe = hilfe + 1;
    }

    hilfe = 0;
    for(int zaehler = 1; zaehler<100; zaehler = zaehler +2){
        text[zaehler] = subtext2[hilfe];
        hilfe = hilfe + 1;
    }

    printf("text nach Wiederherstellung: %s\n\n", text);
}

int main(void){
    system("chcp 1252");
    system("cls");

    char text[101];
    char subtext1[51];
    char subtext2[51];

    for(int i=0; i<100; i++){
        text[i] = 65 + i;
    }

    printf("%s\n\n", text);

    codiere(text, subtext1, subtext2);
    decodiere(subtext1, subtext2, text);


    return 0;
}

#include<windows.h>
#include<stdio.h>
#include<string.h>
#include"../eigene_bibliothek/a_06_02_02_sv01.h"

int main(void) {
    system("chcp 1252");
    system("cls");

    int breite = 10;
    int hoehe = 5;

    rechteck(breite, hoehe);

    return 0;
}

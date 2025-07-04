#include<stdio.h>
#include<windows.h>
#include<string.h>

typedef struct
{
    char strassenname[51];
    int hausnummer;
    char ortname[30];
    char postleitzahl[7];
} ADRESSE;

int main(void) {
    ADRESSE a;

    strcpy(a.strassenname, "Lerchenweg");
    a.hausnummer = 58;
    strcpy(a.ortname, "Siegen");
    strcpy(a.postleitzahl, "006424");

    printf("%s\n", a.strassenname);
    printf("%d\n", a.hausnummer);
    printf("%s\n", a.ortname);
    printf("%s\n", a.postleitzahl);

    return 0;
}

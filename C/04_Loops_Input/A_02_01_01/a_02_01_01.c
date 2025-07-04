#include <stdio.h>
#include <windows.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    char antwort;

    do
    {
        printf("Hallo!\n");

        printf("Wollen Sie, dass ein weiteres mal Hallo ausgegeben wird? 'j' oder 'n' \n");
        fflush(stdin);
        scanf("%c", &antwort);

    } while (antwort == 'j' || antwort == 'J' || antwort == 'y' || antwort == 'Y');       // BEACHTE: Das ist eine solange-Bedingung

    // das Zeichen || ist das logische (offene) ODER 
    
    printf("Dann Tsch�ss!");

    
    printf("\n\n");
    return 0;
}

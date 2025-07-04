#include <stdio.h>
#include <windows.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    int geheimnummer, bestaetigung, zaehler = 0;
    printf("Sie haben 3 Versuche, um eine ganzzahlige Geheimnummer einzugeben");
    do
    {
        zaehler ++; 

        printf("\n\nVersuch Nummer %d. Geben Sie eine ganzzahlige Geheimnummer ein: ", zaehler);
        fflush(stdin);
        scanf("%d", &geheimnummer);

        printf("\nBest�tigen Sie die Geheimnummer, indem Sie die Zahl erneut eingeben: ");
        fflush(stdin);
        scanf("%d", &bestaetigung);

        if (zaehler == 3)
        {
            printf("\n\nSie hatten nur 3 Versuche. Bis zum n�chsten Mal!");
            return 0;
        }
        
    } while (geheimnummer != bestaetigung);
    
    printf("Das ist Ihre neue Geheimnummer: %d", geheimnummer);

    printf("\n\n");
    return 0;
}

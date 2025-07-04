#include <stdio.h>
#include <windows.h>

float x_hoch_n(float x,int n)
{
    float pr = 1;
    
    if (n < 0)
    {
        for (int i = 0; i < abs(n); i++)
        {
            pr *= x;
        }
        

        printf("Ich sagte doch eine positive Zahl. Aber ich kann das auch berechnen. Hier ist das Ergebnis: %f\n", (1 / pr));

        return 1/pr;
    } else
    {
        for (int i = 0; i < n; i++)
            {
                pr *= x;
            }

        return pr;
    }
}

int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    float x;
    int n;

    printf("Geben Sie bitte eine Kommazahl ein: ");
    fflush(stdin);
    scanf("%f", &x);

    printf("Geben Sie bitte eine positive Ganzzahl ein: ");
    fflush(stdin);
    scanf("%d", &n);

    printf("Die Zahl 1 hoch die Zahl 2 ist: %f",x_hoch_n(x,n));

    printf("\n\n");
    return 0;
}

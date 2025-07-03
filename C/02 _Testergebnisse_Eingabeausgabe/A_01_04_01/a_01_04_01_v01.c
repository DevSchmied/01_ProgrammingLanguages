#include <stdio.h>
#include <windows.h>

int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    int alter;

    printf("Geben Sie Ihr Alter ein: ");
    scanf("%d", &alter);

    if (alter > 17)
    {
        printf("In Deutschland gelten Sie als vollj�hrig.\n");   
    } else
    {
        printf("In Deutschland gelten Sie NOCH NICHT als vollj�hrig.\n");
    }
    
    

}

#include <stdio.h>
#include <windows.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    int i;

    for (i = 1; i <= 10; i++)    
    {
        printf("\n%d", i);

        if (i == 10)
        {
            while (i != 1)
            {
                i -= 1;
                printf("\n%d", i);
            }
            
        }
        
    }
    

    printf("\n\n");
    return 0;
}

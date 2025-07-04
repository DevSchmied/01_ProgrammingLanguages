#include <stdio.h>
#include <windows.h>
 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    float array[30];

    for (int i = 1; i <= 10; i++)
    {
        for (int j = 1; j <= 3; j++)
        {
            array[(i -1) * 3 + j -1] = 1.5 + (float)(j - 1);
        }  
    }

    for (int i = 0; i < 30; i++)
    {
        printf("Element %d: %.2f\n", i,array[i]);
    }
    
    

    printf("\n\n");
    return 0;
}

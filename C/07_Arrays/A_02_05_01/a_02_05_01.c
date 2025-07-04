#include <stdio.h>
#include <windows.h>
#include <time.h> 
int main(void)
{
    SetConsoleOutputCP(1252); SetConsoleCP(1252); system("cls");

    srand(time(NULL));
    rand();

    int array[5];

    for (int i = 0; i < 5; i++)
    {
        array[i] = rand() % 100 + 1;
        printf("Der %3d. Wert | der Index %2d | das Element %2d hat den Wert: %2d\n", i+1,i,i, array[i]);
    }

    int index;
    index = rand()%5;
    printf("\narr[\"%d\"]=%d", index, array[index]);
    

    printf("\n\n");
    return 0;
}

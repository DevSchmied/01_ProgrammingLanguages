#include<windows.h>
#include<stdio.h>


int dezimal_nach_binaer(int dezimal){
    if (dezimal > 1) {
        dezimal_nach_binaer(dezimal/2);
    }
    printf("%d", dezimal%2);
}

int main(void)
{
    system("chcp 1252");
    system("cls");

    int dezimal;

    printf("Geben Sie eine ganze Zahl ein: ");
    fflush(stdin);
    scanf("%d", &dezimal);

    dezimal_nach_binaer(dezimal);

    return 0;
}

#include<windows.h>
#include<stdio.h>

int main(void)
{
    system("chcp.com 1252");
    system("cls");

    int x;
    char fan;

    printf("Geben Sie eine ganze Zahl ein: ");
    fflush(stdin);
    scanf("%d",&x);

    switch(x){
        case  1: printf("Januar");      break;
        case  2: printf("Februar");     break;
        case  3: printf("M�rz");        break;
        case  4: printf("April");       break;
        case  5: printf("Mai");         break;
        case  6: printf("Juni");        break;
        case  7: printf("Juli");        break;
        case  8: printf("August");      break;
        case  9: printf("September");   break;
        case 10: printf("Oktober");     break;
        case 11: printf("November");    break;
        case 12: printf("Dezember");    break;

        default: printf("Die Zahl liegt nicht zwischen 1 und 12 \n\n");
    }

    if (x < 1 || x > 12){
        printf("Sind Sie ein Fan von Modulo? j/n \n\n");
        fflush(stdin);
        scanf("%c",&fan);
    }

    if(fan == 'j') {
        x = abs(x % 12);

        switch(x){
            case  1: printf("Januar");      break;
            case  2: printf("Februar");     break;
            case  3: printf("M�rz");        break;
            case  4: printf("April");       break;
            case  5: printf("Mai");         break;
            case  6: printf("Juni");        break;
            case  7: printf("Juli");        break;
            case  8: printf("August");      break;
            case  9: printf("September");   break;
            case 10: printf("Oktober");     break;
            case 11: printf("November");    break;
            case 12: printf("Dezember");    break;

            default: return 0;
        }
    } else {
        return 0;
    }
}

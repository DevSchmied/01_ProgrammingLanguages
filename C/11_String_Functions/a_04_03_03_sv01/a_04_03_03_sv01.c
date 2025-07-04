#include<windows.h>
#include<string.h>
#include<stdio.h>

int main(void) {

    system("chcp 1252");
    system("cls");

    char zeile[2][21];
    int treffer;

    do{

        treffer = 0;

        for(int i=0; i<3; i++){

            printf("Runde %d\n\n",i+1);

            switch(i){
                case 0: printf("Geben Sie die zwei gleichen Zeilen ein.\n");

                        printf("Geben Sie die %d. Zeile ein: ",1);
                        fflush(stdin);
                        scanf("%20[^\n]", zeile[0]);
                        printf("\n");

                        printf("Geben Sie die %d. Zeile ein: ",2);
                        fflush(stdin);
                        scanf("%20[^\n]", zeile[1]);
                        printf("\n");
                        break;

                case 1: printf("Die erste Zeile muss kleiner sein.\n");

                        printf("Geben Sie die %d. Zeile ein: ",1);
                        fflush(stdin);
                        scanf("%20[^\n]", zeile[0]);
                        printf("\n");

                        printf("Geben Sie die %d. Zeile ein: ",2);
                        fflush(stdin);
                        scanf("%20[^\n]", zeile[1]);
                        printf("\n");
                        break;

                case 2: printf("Die zweite Zeile muss kleiner sein.\n");

                        printf("Geben Sie die %d. Zeile ein: ",1);
                        fflush(stdin);
                        scanf("%20[^\n]", zeile[0]);
                        printf("\n");

                        printf("Geben Sie die %d. Zeile ein: ",2);
                        fflush(stdin);
                        scanf("%20[^\n]", zeile[1]);
                        printf("\n");
                        break;
            }
//            printf("Geben Sie die %d. Zeile ein: ",1);
//            fflush(stdin);
//            scanf("%20[^\n]", zeile[0]);
//            printf("\n");
//
//            printf("Geben Sie die %d. Zeile ein: ",2);
//            fflush(stdin);
//            scanf("%20[^\n]", zeile[1]);
//            printf("\n");

            switch(i){
                case 0: if(strcmp(zeile[0],zeile[1]) == 0) treffer++; break;
                case 1: if(strcmp(zeile[0],zeile[1]) < 0) treffer++; break;
                case 2: if(strcmp(zeile[0],zeile[1]) > 0) treffer++; break;
            }
        }

        printf("Treffer: %d\n",treffer);


    } while(treffer != 3);

    printf("Alles O.K.");

    return 0;
}

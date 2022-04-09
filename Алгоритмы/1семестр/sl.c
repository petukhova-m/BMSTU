#include <stdio.h>
#include <string.h>

int main(){
 	char str[101];
    	//char res[101];
    	//printf("Введите строку:");
    	fgets (str, 101, stdin);
    	int f = strlen(str);
    	int n = 3;
	char mas[n][101];
    	int lent[n];
    	//struct Word mas[n];
    	int i = 0;
    	int j = 0;
 	int k = 0;
 	while(str[i] != '\0' && i < f){
 		while(str[i] != ' '){
 			mas[k][j] = str[i];
 			printf("%c", mas[k][j]);
 			i++;
 			j++;
 		}
 		lent[k] = j;
 		k++;
 		i++;
 		i++;
 		j = 0;
 		printf("%d\n", i);
 	}
return 0;
}

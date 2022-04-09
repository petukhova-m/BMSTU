#include  <stdio.h>
#include <locale.h>
#include <string.h>
#include <stdlib.h>
 

struct Word{
	char word[100];
	int len;
};
//int sort;
 
int to_struct(char *arr , struct Word *mas){
    int i = 0, words = 0 , index = 0;
    while(arr[i] != '\0'){
        char *string = (char*)malloc(128 * sizeof(char));
        int len = 0;
        int j = 0;
        while(arr[i] == ' ' && arr[i] != '\0')
            i++;  
        while(arr[i] != ' ' && arr[i] != '\0'){
            string[j] = arr[i];
            i++;
            len++;
            j++;
        }
        for(int q = 0 ; q < len ; q++)
            mas[index].word[q] = string[q];
        mas[index].len = len;
        index++;
        words++;
        free(string);
    }
    mas[words - 1].len--;
    return words;
}
 
 
void csort(struct Word* arr, struct Word* res, int n){
	int count[n];
	for (int i = 0; i < n; i++)
		count[i] = 0;
	int j = 0;
	int i;
	while(j < n - 1){
		i = j+1;
		while(i < n){
			if (arr[i].len < arr[j].len)
				count[j]++;
			else
				count[i]++;
			i++;
		}
		j++;
	}
	for(int i = 0; i < n; i++){
		int k = count[i];
		char temp[100];
		int l = arr[i].len;
		for (int j = 0; j < l; j++){
			temp[j] = arr[i].word[j];
			res[k].word[j] = temp[j];
		}
		res[k].len = l;
	}
}

 
int main(int argc, char *argv[]) {
    char str[101];
    char res[101];
    struct Word arr[101];
    struct Word mas[101];
    fgets (str, 101, stdin);
    char dest[101];
    int n = to_struct(str, arr);
    csort(arr, mas, n);
    for (int i = 0; i < n; i++){
    	for (int j = 0; j < mas[i].len; j++)
    		printf("%c", mas[i].word[j]);
    	printf(" ");
    }
    return 0;
}

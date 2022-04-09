#include <stdio.h> 
#include <string.h> 
#include <stdlib.h> 

struct Word{ 
	char word[100]; 
	int len; 
}; 

int string_to_word(char *arr , struct Word *mas){ 
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
	return words; 
} 

void csort(int n, struct Word *mass, struct Word *sortedMass){ 
	int k; 
	for(int i = 0; i < n; i++){ 
		k = 0; 
		for(int j = 0; j < n; j++){ 
			if(mass[i].len > mass[j].len) 
				k++; 
		} 
		while(sortedMass[k].len != 100) 
			k++; 
		for(int q = 0 ; q < mass[i].len ; q++) 
			sortedMass[k].word[q] = mass[i].word[q]; 
		sortedMass[k].len = mass[i].len; 
	} 
} 

int main(){ 
	char arr[1000]; 
	int words; 
	fgets(arr); 
	struct Word mas[100]; 
	words = string_to_word(arr , mas); 
	struct Word mas_sort[words]; 
	for(int i = 0 ; i < words ; i++) 
		mas_sort[i].len = 100; 
	csort(words , mas , mas_sort); 
	for(int i = 0 ; i < words ; i++){ 
		for(int j = 0 ; j < mas_sort[i].len ; j++) 
			printf("%c" , mas_sort[i].word[j]); 
		printf(" "); 
	} 
	return 0; 
}

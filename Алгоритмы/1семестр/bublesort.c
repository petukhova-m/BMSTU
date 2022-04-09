#include <stdlib.h> 
#include <stdio.h> 
 
int *array; 
 
int compare(unsigned long i, unsigned long j) 
{ 
        if (i <= j) { 
                printf("COMPARE %ld %ld\n", i, j); 
        } else { 
                printf("COMPARE %ld %ld\n", j, i); 
        } 
 
        if (array[i] == array[j]) return 0; 
        return array[i] < array[j] ? -1 : 1; 
} 
 
void swap(unsigned long i, unsigned long j) 
{ 
        if (i <= j) { 
                printf("SWAP %ld %ld\n", i, j); 
        } else { 
                printf("SWAP %ld %ld\n", j, i); 
        } 
 
        int t = array[i]; 
        array[i] = array[j]; 
        array[j] = t; 
} 
 

void bubblesort(unsigned long nel, int (*compare)(unsigned long i, unsigned long j), void (*swap)(unsigned long i, unsigned long j)) 
{ 
    unsigned long first = 0;
    unsigned long last = nel -1;
    while(first <= last){
	unsigned long pos = -1;
	for (unsigned long i = first; i < last; i++)
		if (compare(i, i + 1) > 0){
			swap(i, i + 1);
			pos = i;
		}
	if (pos != -1){
		last = pos;
		pos = -1;
		for (unsigned long i = last; i > first; i--)
			if (compare(i, i - 1) < 0){
				swap(i, i - 1);
				pos = i;
			}
    		if (pos != -1)
    			first = pos;
    		else
    			first = last + 1;
    	}
    	else
    		first = last + 1;
    	}
} 
int main(int argc, char **argv) 
{ 
        int i, n; 
        scanf("%d", &n); 
 
        array = (int*)malloc(n * sizeof(int)); 
        for (i = 0; i < n; i++) scanf("%d", array+i); 
 
        bubblesort(n, compare, swap); 
        for (i = 0; i < n; i++) printf("%d ", array[i]); 
        printf("\n"); 
 
        free(array); 
        return 0; 
}

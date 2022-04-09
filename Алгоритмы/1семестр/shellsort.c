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
 
int fib(int n){
	if (n == 0) return 1;
	if (n == 1) return 1;
	return fib(n-1) + fib(n-2);
}

void shellsort(unsigned long nel, int (*compare)(unsigned long i, unsigned long j), void (*swap)(unsigned long i, unsigned long j)){
	int s = 1;
	while(fib(s) <= nel)
		s++;
	s--;
	int step = fib(s);
        //int step = nel/2;
        while(step > 0){
		for (int i = step; i < nel; i++) 
			for (int j = i - step; j >= 0; j -= step)
				if(compare(j, j + step) > 0)
					swap(j, j+ step);
				else
					break;
		if (s > 1)
			s--;
		else
			break;
		step = fib(s);
	}
}
 

int main(int argc, char **argv) 
{ 
        int i, n; 
        scanf("%d", &n); 
 
        array = (int*)malloc(n * sizeof(int)); 
        for (i = 0; i < n; i++) scanf("%d", array+i); 
 
        shellsort(n, compare, swap); 
        for (i = 0; i < n; i++) printf("%d ", array[i]); 
        printf("\n"); 
 
        free(array); 
        return 0; 
}

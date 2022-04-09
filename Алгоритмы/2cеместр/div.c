#include<stdio.h>
#include<math.h>
#include<stdlib.h>
#include<time.h>

long dividers(long x, long *del){
	long j = 0;
	for(long i = 1; i*i <= x; i++)
		if (x % i == 0){
			del[j] = i;
			j++;
			//del[j++] = x / i;
		}
	long k = j - 1;
	for(long i = k; i >= 0; i--){
		if (del[i] != x / del[i]){
			del[j] = x / del[i];
			j++;
		}
	}
	return (j - 1);
}

void DIV(long x, long z, long y, unsigned long *a, int *k) {
        a[*k] = y;
	//printf("a[%d] = %ld\n", *k, a[*k]);
	(*k)++;
	if (z != y) {
		for(y++; y * y <= x && x % y != 0; y++);
		if (y*y <= x) 
			DIV(x, x/y, y, a, k);
		a[*k] = z;
		//printf("a[%d] = %ld\n", *k, a[*k]);
		(*k)++;
	}
}

int main(){
	double start = clock();
	long x;
	scanf("%ld", &x);
	long *del = (long*)malloc(10000 * sizeof(long));
	//long k = dividers(x, del);
	int k = 0;
	DIV(x, x, 1, del, &k);
	printf("graph {\n");
	for (long i = k; i >= 0; i--)
		printf("%ld\n", del[i]);
	for (long i = k; i >= 0; i--){
		for (long j = i - 1; j >= 0; j--){
			int u = 0;
			if (del[i] % del[j] == 0)
				u = 1;
			for (long z = i -1; z > j; z--)
				if (del[i] % del[z] == 0 && del[z] % del[j] == 0)
					u = 0;
			if (u)
				printf("%ld -- %ld\n", del[i], del[j]);
		}
	}
	printf("}");
	 printf("%.4lf\n", (clock() - start) / CLOCKS_PER_SEC);
	return 0;
}

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

void Kadane(float *a, int l, int r, int n){
	float maxpr = a[0];
	int start = 0;
	float pr = 0;
	int i = 0;
	while(i < n){
		pr += a[i];
		//printf("%d %f\n", i, pr);
		if (pr > maxpr){
			maxpr = pr;
			l = start;
			r = i;
		}
		i++;
		if (pr < 0){
			pr = 0;
			start = i;
		}
	}
	printf("%d %d\n", l, r);
}

int main(){
	int n = 0;
	scanf("%d", &n);
	float a, b;
	float *arr = (float*)malloc(n * sizeof(float));
	for (int i = 0; i < n; i++){
		scanf("%f%f", &a, &b);
		arr[i] = log(a/b);
		//printf("%f\n", arr[i]);
	}
	int l = 0, r = 0;
	float maxpr = arr[0];
	int start = 0;
	float pr = 0;
	int i = 0;
	while(i < n){
		pr += arr[i];
		//printf("%d %f\n", i, pr);
		if (pr > maxpr){
			maxpr = pr;
			l = start;
			r = i;
		}
		i++;
		if (pr < 0){
			pr = 0;
			start = i;
		}
	}
	printf("%d %d\n", l, r);
	//Kadane(arr, l, r, n);
	//printf("%d %d", l, r);
	free(arr);
	return 0;
}

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int gcd(int a, int b) {
	if (b != 0){
		int k = a % b;
		return abs(gcd(b, k));
	}
	else
		return abs(a);
}

int logs(int a){
    int k = 0;
    while(a > 0){
        a/=2;
        k++;
    }
    return k - 1;   
}

int pows(int a, int b){
	if (b == 0)
		return 1;
	int k = 1;
	for (int i = 0; i < b; i++)
		k *= a;
	return k;
}

//int gcd(int a, int b){
//	while(a != 0 && b!= 0){
//		if (a > b)
//			a = a % b;
//		else
//			b = b % a;
//	}
//	return(a + b);
//}

void ComputeLogarithms(int m, int *lg){
	int i = 1;
	int j = 0;
	while(i < m){
		while(j < pows(2, i)){
			lg[j] = i - 1;
			j++;
		}
		i++;
	}
}

int SparseTable_Query(int **st, int l, int r, int *lg){
	int j = lg[r - l + 1];
	int k = pows(2, j);
	return gcd(st[j][l], st[j][r - k + 1]);
}

void SparseTable_Build(int *arr, int *lg, int **st, int n){
	int m = lg[n]+1;
	int i = 1;
	while(i < n){
		st[0][i] = arr[i];
		i++;
	}
	int j = 1;
	while(j < m){
		i = 0;
		while(i <= n - pows(2, j)){
			int k = pows(2, j - 1);
			st[j][i] = gcd(st[j - 1][i], st[j - 1][i + k]);
			printf("%d %d %d\n", st[j - 1][i], st[j - 1][i + k], st[j][i]);
			i++;
		}
		j++;
	}
}


int main(int argc, char** argv) {
        int n = 0;
        scanf("%d", &n);
        int *a = (int*)malloc(n * sizeof(int));
        for(int i = 0;i < n; i++) {
                scanf("%d", &a[i]);
        }
        int *lg = (int*)malloc(2000000 * sizeof(int));
        ComputeLogarithms(20, lg);
        int **st = (int**)malloc(100000 * sizeof(int));
        for (int i = 0;i <= lg[n]; i++) {
                st[i] = (int*)malloc(300001 * sizeof(int));
        }
        SparseTable_Build(a, lg, st, n);
        int k = 0;
        scanf("%d", &k);
        for (int i = 0; i < k; i++) {   
        	int a, b;     
                scanf("%d %d", &a, &b);
                printf("%d\n", SparseTable_Query(st, a, b, lg));
        }
        free(a);
        free(lg);       
        free(st);
        return 0;
}


#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>

int peaks(int *a, int i, int n){
        if (n == 1)
                return 1;
	if (n > 1 && i == 0 && a[i] >= a[i+1])
		return 1;
	if (n > 1 && i == n-1 && a[i] >= a[i-1])
		return 1;
	if (n > 1 && i != n-1 && i != 0 && a[i] >= a[i+1] && a[i] >= a[i-1])
		return 1;
	else
		return 0;
}

int peak(int *arr, int a, int b, int n){
	int sum = 0;
	int j = a+1;
	if (n == 1)
	        return 1;
	while(a <= b){
		if (n > 1 && a == 0 && arr[a] >= arr[a+1])
			sum++;
		if (n > 1 && a == n-1 && arr[a] >= arr[a-1])
			sum++;
		if (a < n - 1 && a > 0 && arr[a] >= arr[a+1] && arr[a] >= arr[a-1])
			sum++;
		a++;
	}
	return sum;
}

void build (int *a, int v, int tl, int tr, int *t, int n) {
	if (tl == tr)
		t[v] = peak(a, tl, tr, n);
	else {
		int tm = (tl + tr) / 2;
		build (a, v*2 + 1, tl, tm, t, n);
		build (a, v*2+ + 2, tm+1, tr, t, n);
		t[v] = t[v*2 + 1] + t[v*2 + 2];
	}
}

int query(int *t, int l, int r, int a, int b, int ver){
	if (l == a && r ==b)
		return t[ver];
	else{
		int m = (a+b)/2;
		if (r <= m)
			return query(t, l, r, a, m, 2*ver+1);
		else{
			if (l > m)
				return query(t, l, r, m + 1, b, 2*ver + 2);
			else
				return (query(t, l, m, a, m, 2*ver+1) +query(t, m+1, r, m+1, b, 2*ver + 2));
			return query(t, l, r, a, m, 2*ver+1);
		}	
	}
}


void upd(int *t, int i, int v, int a, int b, int j, int *arr, int n){
	if (a == b){
		//arr[i] = v;
		//for (int k = 0; k <n; k++)
		//	printf("%d ", arr[k]);
		//printf("\n");
		//for (int k = 0; k <4*n; k++)
		//	printf("%d ", t[k]);
		//printf("\n");
		t[j] = peaks(arr, a, n);
	}
	else{
		int m = (a + b) /2;
		if (i <= m)
			upd(t, i, v, a, m, 2*j + 1, arr, n);
		else
			upd(t, i, v, m+1, b, 2*j+2, arr, n);
		t[j] = t[2*j + 1] + t[2*j+2];
	}
}


int main(){
    int n;
    scanf("%d" , &n);
    int *arr = (int*)malloc(n * sizeof(int));
    for(int i = 0; i < n ; i++)
        scanf("%d" , &arr[i]);
    int k;
    scanf("%d" ,&k);
    int a , b;
    int *tree = (int*)malloc(4 * n * sizeof(int));
    build(arr , 0, 0, n - 1 ,  tree, n);
    char s[4];
    for(int i = 0 ; i < k ; i++){
        scanf("%s" , s);
        scanf("%d" , &a);
        scanf("%d" , &b);
        if(s[0] == 'P')
            printf("%d\n" , query(tree ,a , b , 0 , n - 1, 0));
        else if(s[0] = 'U'){
        	arr[a] = b; 
        	//for (int k = 0; k < n; k++)
        	//	printf("%d ", arr[i]);
		//printf("\n");
		upd(tree, a, b, 0, n - 1, 0, arr, n); 
		if(a != 0) 
			upd(tree, a - 1, b, 0, n - 1, 0, arr, n); 
		if (a != n - 1) 
				upd(tree, a + 1, b, 0, n - 1, 0, arr, n); 
	}
        
    }
    //free(s);
    free(arr);
    free(tree);
    return 0;
}
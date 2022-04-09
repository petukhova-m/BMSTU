#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>

int cmax(int a, int b){
	if (a >= b)
		return a;
	else
		return b;
}

void build (int *a, int v, int tl, int tr, int *t) {
	if (tl == tr)
		t[v] = a[tl];
	else {
		int tm = (tl + tr) / 2;
		build (a, v*2 + 1, tl, tm, t);
		build (a, v*2+ + 2, tm+1, tr, t);
		t[v] = cmax(t[v*2 + 1], t[v*2 + 2]);
	}
}

int query(int *t, int l, int r, int a, int b, int ver){
	if (l == a && r ==b)
		return t[ver];
	else{
		int m = (a+b)/2;
		if (r <= m)
			return query(t, l, r, a, m, 2*ver+1);
		else
			if (l > m)
				return query(t, l, r, m + 1, b, 2*ver + 2);
			else
				return cmax(query(t, l, m, a, m, 2*ver+1), query(t, m+1, r, m+1, b, 2*ver + 2));
	}
}

void upd(int *t, int i, int v, int a, int b, int j){
	if (a == b)
		t[j] = v;
	else{
		int m = (a + b) /2;
		if (i <= m)
			upd(t, i, v, a, m, 2*j + 1);
		else
			upd(t, i, v, m+1, b, 2*j+2);
		t[j] = cmax(t[2*j + 1], t[2*j+2]);
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
    build(arr , 0, 0, n - 1 ,  tree);
    char *s = (char*)malloc(4 * sizeof(char));
    for(int i = 0 ; i < k ; i++){
                //s[3] = '\0';
                scanf("%s" , s);
        scanf("%d" , &a);
        scanf("%d" , &b);
        if(s[0] == 'M')
            printf("%d\n" , query(tree ,a , b , 0 , n - 1, 0));
        else if(s[0] = 'U')
            upd(tree , a , b , 0 , n - 1, 0);
        
    }
    free(s);
    free(arr);
    free(tree);
    return 0;
}

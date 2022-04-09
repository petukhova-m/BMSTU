#include <stdio.h>
#include <math.h>

int abs(int a){
	if(a >= 0)
		return a;
	else
		return -a;
}

void merge(int first, int mid, int last, int *a){
	int i = first;
	int j = mid + 1;
	int k = 0;
	int b[last - first + 1];
	while(k < last - first + 1){
                if ((i <= mid) && (j == last + 1 || abs(a[i]) <= abs(a[j]))){
                	printf("%d - %d\n", k, b[k]);
                	printf("%d - %d\n", i, a[i]);
                	
                	b[k] = a[i];
                        i++;
                }
                else{
                	if (j <= last){
                		b[k] = a[j];
                		j++;
                	}
                }
                k++;
        }
        int q = 0;
        for(int r = first; r <= last ; r++)
    		if (q < k){
        		a[r] = b[q];
        		q++;
        	}
}


void vstavki(int m, int n, int *a) {
	int i , elem , loc;
    	i = m + 1;
    	while (i <= n){
    		elem = a[i];
        	loc = i - 1;
        	while(loc >= m && abs(a[loc]) > abs(elem)){
            		a[loc + 1] = a[loc];
            		loc--;
        	}
        	a[loc + 1] = elem;
        	i++;
    	}
}

void slianie(int first, int last, int *a){
	if (last > first && last - first <= 5)
		vstavki(first, last, a);
	else{
		if (first < last){
			int mid = (first + last) / 2;
			slianie(first, mid, a);
			slianie(mid + 1, last, a);
			merge(first, mid, last, a);
		}
	}
}

int main(){
	int n;
	scanf("%d", &n);
	int a[n];
	//int b[n];
	for (int i =0; i < n; i++)
		scanf("%d", &a[i]);
	slianie(0, n - 1, a);
	//vstavki(0, n, a);
	for (int i = 0; i < n; i++)
		printf("%d ", a[i]);
	return 0;
}

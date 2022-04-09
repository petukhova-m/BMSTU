#include <stdio.h>
#include <stdlib.h>

struct a{
        int v;
        int index;
};

struct q{
        int count;
        int cap;
        struct a *heap;
};

void init(struct q* w, int n){
	w->cap = n;
        w->count = 0;
        w->heap = malloc(n * sizeof(struct a)); 
}

void swap (int a, int b, struct q* e){
        int v1, index1;
        v1 = e->heap[a].v;
        index1 = e->heap[a].index;
        e->heap[a].v = e->heap[b].v;
        e->heap[a].index = e->heap[b].index;
        e->heap[b].v = v1;
        e->heap[b].index = index1;
}

int compare(struct q* y, int i, int j){
	if (y->heap[i].v > y->heap[j].v)
		return 1;
	return 0;
}

void heapify (int i, int n, struct q* y){
        int l, r, j;
        while (free) {
                l = 2 * i + 1;
                r = l + 1;
                j = i;
                if (l < n && compare(y, i, l) == 1)       
                        i = l;
                if (r < n && compare(y, i, r) == 1)
                        i = r;
                if (i == j)
                        break;
                swap(i, j, y);
        }        
}

void insert(struct q* y,struct a b){   
        int i = y->count;
        y->count += 1;
        y->heap[i] = b;
        while ((i > 0) && (y->heap[(i - 1)/2 ].v > y->heap[i].v)) {
                swap((i - 1)/2, i, y);
                i = (i - 1)/2;
        }
        //y->heap[i]=b;
}

struct a extractmin(struct q* y){
        struct a x = y->heap[0];
        y->count--;
        if (y->count > 0) {
                y->heap[0] = y->heap[y->count];
                heapify(0, y->count, y);
        }
        return x;
}

void mer(int n, int sum, int *kol, int **mas, struct q* w){
	int tm[n];
	for (int i = 0; i < n; i++)
		tm[i] = 0;
	for (int i = 0; i < n; i++) {
                if (kol[i] != 0) {
                        struct a e;
                        e.v = mas[i][0];
                        e.index = i;
                        insert(w, e);
                        tm[i] ++;
                }          
        }
        for (int i = 0; i < sum; i++) {
                struct a r;
                r = extractmin(w);
                printf("%d ", r.v);
                if (tm[r.index] != kol[r.index]) {
                        struct a x;
                        x.v = mas[r.index][tm[r.index]];
                        x.index = r.index;
                        insert(w, x);
                        tm[r.index]++;
                }       
        }
        printf("\n");
}

int main(){
        int n, sum;
        sum = 0;
        scanf("%d", &n);
        int *kol = malloc(n * sizeof(int));
        int **mas = malloc(n * sizeof(int*));  
        struct q* w = malloc(sizeof(struct q));
        init(w, n);
        for (int i = 0; i < n; i++) {
                scanf("%d", &kol[i]);
                sum += kol[i];
                mas[i] = malloc(kol[i] * sizeof(int));
        }
        for (int i = 0; i < n; i++) {
                for(int j = 0; j < kol[i]; j++) {
                        scanf("%d", &mas[i][j]);
                }
        }
        mer(n, sum, kol, mas, w);
        free(w->heap);
        for (int i = 0; i < n; i++)
                free(mas[i]);
        free(mas);
        free(kol);
        free(w);
        return 0;
}

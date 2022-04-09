#include <stdio.h>
#include <string.h>
#include <stdlib.h>

struct Elem {
        struct Elem *next;
        int v;
        int number;
};

void Assign(int i, int v, int m, struct Elem **hashtable)
{
	int n = i % m;
	struct Elem *p, *l;
	p = hashtable[n];
	int u = 1;
	while (p->next != NULL) {
			if (p->number == i) {
			u = 1;
			p->v = v;
			return;
		}
		p = p->next;
	} 
	if (p->number == i) {
		u = 1;
		p->v = v;
		return;
	}
	//if (u == 0){
		p->next = (struct Elem*)malloc(sizeof(struct Elem));
		p->next->next = NULL;
		p->next->v = v;
		p->next->number = i;
	//}
}

void At(int i, int m, struct Elem **hashtable)
{
	int n = i % m;
	struct Elem *p;
	p = hashtable[n];
	int u = 0;
	while (p->next != NULL) {
		if (p->number == i) {
			u = 1;
			printf("%d\n", p->v);
			break;
		}
		else
			p = p->next;
	}
	if (p->number == i && u == 0) {
		u = 1;
		printf("%d\n", p->v);
		//return;
	}
	if (u == 0)
		printf("%d\n", 0);
}

int main(int argc, char **argv)
{
	int n = 0, m = 0;
	scanf("%d", &n);
	scanf("%d", &m);
	struct Elem *hashtable[m];
	for(int i = 0; i < m; i++){
		hashtable[i] = (struct Elem*)malloc(sizeof(struct Elem));
		hashtable[i]->next = NULL;
		hashtable[i]->number = -1;
	}
	
	char a[6];
	int q, v;
	for(int i = 0; i < n; i++){
		scanf("%s", a);
		if (strcmp(a, "AT") == 0) {
			scanf("%d", &q);
			At(q, m, hashtable);
		}
		if (strcmp(a, "ASSIGN") == 0){
			scanf("%d", &q);
			scanf("%d", &v);
			Assign(q, v, m, hashtable);
		}
	}
	
	for (int i = 0; i < m; i++) {
		struct Elem *p = hashtable[i];
		while (p->next != NULL) {
			struct Elem *q = p->next;
			free(p);
                        p = q;
		}
		free(p);
	}
	return 0;
}

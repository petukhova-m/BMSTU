#include <stdlib.h>
#include <stdio.h>
#include <string.h>

struct st{
	int *data;
	int cap;
	int top;
};

void init(struct st *s){
	s->cap = 100000;
	s->top = 0;
	s->data = malloc(s->cap * sizeof(int));
}

void Const(struct st *s, int x){
	//if (s->cap == s->top){
	//	s->cap *= 2;
	//	s->data = realloc(s->data, s->cap * sizeof(int));
	//}
	s->data[s->top] = x;
	s->top++;
}

void Add(struct st *s){
	int temp = s->data[s->top-1] + s->data[s->top - 2];
	s->top--;
	s->data[s->top - 1] = temp;
}

void Sub(struct st *s){
	int temp = s->data[s->top-1] - s->data[s->top - 2];
	s->top--;
	s->data[s->top - 1] = temp;
}

void Mul(struct st *s){
	int temp = s->data[s->top-1] * s->data[s->top - 2];
	s->top--;
	s->data[s->top - 1] = temp;
}

void Div(struct st *s){
	int temp = s->data[s->top-1] / s->data[s->top - 2];
	s->top--;
	s->data[s->top - 1] = temp;
}

int maxs(int a, int b){
	if(a > b)
		return a;
	else
		return b;
}

int mins(int a, int b){
	if(a < b)
		return a;
	else
		return b;
}

void Max(struct st *s){
	int temp = maxs(s->data[s->top-1], s->data[s->top - 2]);
	s->top--;
	s->data[s->top - 1] = temp;
}

void Min(struct st *s){
	int temp = mins(s->data[s->top-1], s->data[s->top - 2]);
	s->top--;
	s->data[s->top - 1] = temp;
}

void Neg(struct st *s){
	s->data[s->top - 1] *= (-1);
}

void Dup(struct st *s){
	int temp = s->data[s->top-1];
	s->top++;
	s->data[s->top - 1] = temp;
}

void Swap(struct st *s){
	int temp = s->data[s->top - 1];
	s->data[s->top - 1] = s->data[s->top - 2];
	s->data[s->top - 2] = temp;
}

int main(){
	struct st s;
	init(&s);
	int n;
	scanf("%d", &n);
	char str[5];
	for (int i = 0; i < n; i++){
		scanf("%s", str);
		if (strcmp(str, "CONST") == 0){
			int x = 0;
			scanf("%d", &x);
			Const(&s, x);
		}
		if (strcmp(str, "ADD") == 0)
			Add(&s);
		if (strcmp(str, "SUB") == 0)
			Sub(&s);
		if (strcmp(str, "MUL") == 0)
			Mul(&s);
		if (strcmp(str, "DIV") == 0)
			Sub(&s);
		if (strcmp(str, "MAX") == 0)
			Max(&s);
		if (strcmp(str, "MIN") == 0)
			Min(&s);
		if (strcmp(str, "NEG") == 0)
			Neg(&s);
		if (strcmp(str, "DUP") == 0)
			Dup(&s);
		if (strcmp(str, "SWAP") == 0)
			Swap(&s);
	}
	printf("%d", s.data[s.top - 1]);
	free(s.data);
	return 0;
}
			
			

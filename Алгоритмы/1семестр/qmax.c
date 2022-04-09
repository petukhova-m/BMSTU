#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct stack{
	int cap;
	int top1;
	int top2;
	int *max;
	int *data;
};

void init(struct stack *s){
	s->cap = 10000;
	s->top1 = 0;
	s->top2 = 9999;
	s->data = malloc(s->cap * sizeof(int));
	s->max = malloc(s->cap * sizeof(int));
}

void empty(struct stack *s){
	if (s->top1 == 0 && s->top2 == s->cap-1)
		printf("true\n");
	else
		printf("false\n");
}

void push1 (struct stack *s, int x) {
        s->data[s->top1]=x;
        //printf("%d\n", s->top1);
        if (s->top1 == 0)
                s->max[s->top1]=x;   
        if (s->top1 > 0 && s->max[s->top1-1]<x)
                s->max[s->top1]=x;
        else 
                if (s->top1 > 0)
                	s->max[s->top1]=s->max[s->top1-1];
        s->top1++;
}

void push2 (struct stack *s, int x){
        s->data[s->top2]=x;
        if (s->top2 == s->cap - 1)
                s->max[s->top2]=x;   
        if (s->top2 < s->cap - 1 && s->max[s->top2+1] > x)
                s->max[s->top2] = s->max[s->top2 + 1];
        else 
                if (s->top2 < s->cap - 1)
                	s->max[s->top2] = x;
        s->top2--;
}

int pop1 (struct stack *s){
        int x;
        s->top1--;
        x=s->data[s->top1];
        return x;
}

int pop2 (struct stack *s){
        int x;
        s->top2++;
        x=s->data[s->top2];
        return x;
}

int deq(struct stack *s){
        int x;
        if (s->top2 == s->cap - 1) {
                while (s->top1 != 0) {
                       push2(s, pop1(s));
               }
        }
        x=pop2(s);
        return x;
}

int max(struct stack *s){
        if (s->top1 != 0 && s->top2 == s->cap - 1)
                return s->max[s->top1-1];
        if (s->top1 != 0 && s->top2 != s->cap - 1 && s->max[s->top1 - 1] >= s->max[s->top2 + 1])
        	return s->max[s->top1-1];
        else 
        	return s->max[s->top2+1];
}

int main(){
	struct stack s;
	init(&s);
	int n = 0;
	scanf("%d", &n);
	char str[5];
	for (int i = 0; i < n; i++){
		scanf("%s", str);
		if (strcmp(str, "ENQ") == 0){
			int x = 0;
			scanf("%d", &x);
			push1(&s, x);
		}
		if (strcmp(str, "MAX") == 0)
			printf("%d\n", max(&s));
		if (strcmp(str, "DEQ") == 0)
			printf("%d\n", deq(&s));
		if (strcmp(str, "EMPTY") == 0)
			empty(&s);
	}
	free(s.data);
	free(s.max);
	return 0;
}

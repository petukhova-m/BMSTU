#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct buf{
        long *data;
        int cap;
        int count;
        int head;
        int tail;
};

void Init(struct buf *q){
        q->cap = 4;
        q->count = 0;
        q->head = 0;
        q->tail = 0;
        q->data = malloc(q->cap * sizeof(long));
}

void Empty(struct buf *q){
        if (q->count == 0){
                printf("true");
                printf("\n");
        }
        else{
                printf("false");
                printf("\n");
        }
}

void New(struct buf *q){
        int k = q->cap;
        q->cap *= 2;
        q->data = realloc(q->data, q->cap * sizeof(long));
        for (int i = q->tail; i < q->count; i++)
                q->data[k + i] = q->data[i];
}

void Enque(struct buf *q, long x){
        if (q->count == q->cap)
                New(&q);
        q->data[q->tail] = x;
        q->tail++;
        if (q->tail == q->cap)
                q->tail = 0;
        q->count++;
}

long Deq(struct buf *q){
        long x = q->data[q->head];
        q->head++;
        if (q->head == q->cap)
                q->head = 0;
        q->count--;
        //printf("%ld\n", x);
        return x;
}

int main(){
        struct buf q;
        Init(&q);
        long n;
        scanf("%ld", &n);
        char str[5];
        long b = n;
        for (long i = 0; i < b; i++){
                scanf("%s", str);
                if (strcmp(str, "EMPTY") == 0)
                        Empty(&q);
                if (strcmp(str, "ENQ") == 0){
                        long x = 0;
                        scanf("%ld", &x);
                        Enque(&q, x);
                }
                if (str[1] == 69 && q.count > 0){
                       // Deq(&q);
                        printf("%ld\n", 
                        Deq(&q));
                }
        }
        free(q.data);
        return 0;
}

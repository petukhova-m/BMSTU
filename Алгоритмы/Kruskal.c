#include <stdio.h>
#include <stdlib.h>
#include <math.h>

struct Vertex{
        int x, y;
        int dep;
        struct Vertex *parent;
    };

int compare(struct Vertex *this, struct Vertex *obj){
	if (this->x > obj->x && this->y > obj->y)
                return 1;
        if (this->x == obj->x && this->y == obj->y)
                return 0;
        else return -1;
        }
        
struct Vertex* find(struct Vertex* this){
        if(compare(this, this->parent) == 0)
                return this;
        else
                return find(this->parent);
}
        
struct Road{
        struct Vertex *A;
        struct Vertex *B;
        double len;
};
    
struct Vertex *nodes;
struct Road *roads;

int compareTo(const struct Road *this, const struct Road *obj) {
            if (this->len > obj->len)
                return 1;
            if (this->len == obj->len)
                return 0;
            else
                return -1;
}

void Union(struct Vertex *v, struct Vertex *u){
        struct Vertex *v1 = find(v);
        struct Vertex *u1 = find(u);
        if(v1->dep > u1->dep)
            u1->parent = v1;
        else{
            v1->parent = u1;
            u1->dep++;
        }
    }

int main() {
	int n;
        scanf("%d", &n);
        nodes = (struct Vertex*)malloc(sizeof(struct Vertex)*n);
        for (int i = 0; i < n; i++) {
            int x, y;
            scanf("%d %d", &x, &y);
            nodes[i].x = x;
            nodes[i].y = y;
            nodes[i].parent = &nodes[i];
            nodes[i].dep = 0;
        }

        int num = n * (n - 1) / 2;
        roads = (struct Road*)malloc(sizeof(struct Road)*num);
        int k = 0;
        for (int i = 0; i < n; i++)
            for (int j = i + 1; j < n; j++) {
                roads[k].A = &nodes[i];
                roads[k].B = &nodes[j];
                int a = pow((roads[k].A->x - roads[k].B->x), 2);
                int b = pow((roads[k].A->y - roads[k].B->y), 2);
                roads[k].len = sqrt(a + b);
                k++;
            }
        qsort(roads,num, sizeof(struct Road), (int(*) (const void *, const void *)) compareTo);
         double summ = 0;
         int edges = 0;
         for(int i = 0; i < num && edges < n - 1; i++) {
             if (compare(find(roads[i].A), find(roads[i].B)) != 0) {
                 summ += roads[i].len;
                 Union(roads[i].A, roads[i].B);
                 edges++;
             }
         }
          printf("%.2f\n", summ);
    }


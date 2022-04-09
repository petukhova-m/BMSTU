#include<stdio.h>
#include<stdlib.h>

struct Vertex{
	int edge[10][2];
	int visit;
	int numComp;
	int koledge;
};

struct Vertexmax{
	int edge[300][2];
	int visit;
	int numComp;
	int koledge;
};

void visitVer(struct Vertex *nodes, int i, int comp, int *countVer, int *countEdge){
	nodes[i].visit = 1;
	nodes[i].numComp = comp;
	(*countVer)++;
	(*countEdge) += nodes[i].koledge;
	for(int j = 0; j < nodes[i].koledge; j++){
		int k = nodes[i].edge[j][0];
		if(nodes[k].visit == 0)
			visitVer(nodes, k, comp, countVer, countEdge);
	}
}

void visitVermax(struct Vertexmax *nodes, int i, int comp, int *countVer, int *countEdge){
	nodes[i].visit = 1;
	nodes[i].numComp = comp;
	(*countVer)++;
	(*countEdge) += nodes[i].koledge;
	for(int j = 0; j < nodes[i].koledge; j++){
		int k = nodes[i].edge[j][0];
		if(nodes[k].visit == 0)
			visitVer(nodes, k, comp, countVer, countEdge);
	}
}

void printVermax(struct Vertexmax *nodes, int i, int maxComp){
	//nodes[i].visit = 1;
	for(int j = 0; j <nodes[i].koledge; j++){
		int k = nodes[i].edge[j][0];
		if(nodes[i].edge[j][1] == 0){
			printf("%d -- %d ", i, k);
			if (nodes[i].numComp == maxComp)
				printf("[color = red]");
			printf("\n");
			nodes[i].edge[j][1] = 1;
			int f = nodes[k].edge[0][0];
			int l = 0;
			for(int h = 0; h < nodes[k].koledge && f != i; h++){	
				f = nodes[k].edge[h][0];
				l = h;
			}
			nodes[k].edge[l][1] = 1;
			
		}
	}
	for(int j = 0; j < nodes[i].koledge; j++){
		int k = nodes[i].edge[j][0];
		if(nodes[i].edge[j][1] == 0)
			printVer(nodes, k, maxComp);
	}
}

void printVer(struct Vertex *nodes, int i, int maxComp){
	//nodes[i].visit = 1;
	for(int j = 0; j <nodes[i].koledge; j++){
		int k = nodes[i].edge[j][0];
		if(nodes[i].edge[j][1] == 0){
			printf("%d -- %d ", i, k);
			if (nodes[i].numComp == maxComp)
				printf("[color = red]");
			printf("\n");
			nodes[i].edge[j][1] = 1;
			int f = nodes[k].edge[0][0];
			int l = 0;
			for(int h = 0; h < nodes[k].koledge && f != i; h++){	
				f = nodes[k].edge[h][0];
				l = h;
			}
			nodes[k].edge[l][1] = 1;
			
		}
	}
	for(int j = 0; j < nodes[i].koledge; j++){
		int k = nodes[i].edge[j][0];
		if(nodes[i].edge[j][1] == 0)
			printVer(nodes, k, maxComp);
	}
}

int main(){
	int n, m;
	scanf("%d%d", &n, &m);
	int maxComp = 0, maxEdge = 0, countVer = 0, countEdge = 0;
	int maxNodes = 0, comp = 0, minVer = -1;
	//if (n > 990000)
		struct Vertex *nodes = (struct Vertex*)malloc(n*sizeof(struct Vertex));
	if (n < 990000)
		struct Vertexmax *nodes = (struct Vertexmax*)malloc(n*sizeof(struct Vertexmax));
	for(int i = 0; i < n; i++){
		nodes[i].koledge = 0;
		nodes[i].visit = 0;
		nodes[i].numComp = -1;
	}
	for(int i = 0; i < m; i++){
		int u, v;
		scanf("%d%d", &u, &v);
		nodes[u].edge[nodes[u].koledge][0] = v;
		nodes[u].edge[nodes[u].koledge][1] = 0;
		nodes[u].koledge++;
		if (u != v){
			nodes[v].edge[nodes[v].koledge][0] = u;
			nodes[v].edge[nodes[v].koledge][1] = u;
			nodes[v].koledge++;
		}
	}
	for(int i = 0; i < n; i++){
		if(nodes[i].visit == 0){
			if (n > 990000)
				visitVer(nodes, i, comp, &countVer, &countEdge);
			else
				visitVermax(nodes, i, comp, &countVer, &countEdge);
			countEdge/=2;
			comp++;
		}
		if(countVer > maxNodes){
			maxNodes = countVer;
			maxEdge = countEdge;
			maxComp = comp - 1;
		}
		if(countVer == maxNodes && countEdge > maxEdge){
			maxEdge = countEdge;
			maxComp = comp - 1;
		}
		countEdge = 0;
		countVer = 0;
	}
	
	printf("graph {\n");
	for(int i = 0; i < n; i++){
		printf("%d ", i);
		if(nodes[i].numComp == maxComp)
			printf("[color = red]");
		printf("\n");
	}
	for(int i = 0; i < n; i++)
		nodes[i].visit = 0;
	for(int i = 0; i < n; i++)
		if (n > 990000)
			printVer(nodes, i, maxComp);
		else
			printVermax(nodes, i, maxComp);

	printf("}");
	free(nodes);
	return 0;
}

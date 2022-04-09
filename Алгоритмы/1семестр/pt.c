#include <stdio.h>
#include <string.h>
#include <stdlib.h>

struct Unit {
        struct Unit *arcs[26];
        long keyw;
        long keys;
} *root;

struct Unit* InitUnit(){
	struct Unit *unit = malloc(sizeof(struct Unit));
	unit->keyw = unit->keys = 0;
	int i;
	for(i = 0; i < 26; i++)
		unit->arcs[i] = NULL;
	return unit;
}

long Find(char k){
	char key[26] = "abcdefghijklmnopqrstuvwxyz";
	int i = 0;
	for (i = 0; i < 26; i++)
		if (k == key[i])
			return i;
}

void Prefix(struct Unit *rt, char *k){
	struct Unit *unit;
	unit = rt;
	long i, lenk = strlen(k), findk;
	for(i = 0; i < lenk; i++) {
		findk = Find(k[i]);
		if (unit->arcs[findk] == NULL) {
			printf("%d\n", 0);
			return;
		}
		unit = unit->arcs[findk];
	}	
	printf("%ld\n", unit->keys);
}

int main(int argc, char **argv){
	root = InitUnit();
	long i, n, m;
	scanf("%ld", &n);
	char a[10];
	char *k = (char*)malloc(sizeof(char)*100000);
	for(i = 0; i < n; i++) {
		scanf("%s", a);
		if (strcmp(a, "INSERT") == 0){
			scanf("%s", k);
			InsertUnit(root, k);
			struct Unit *unit;
			unit = rt;
			long i, lenk = strlen(k), findk;
			for(i = 0; i < lenk; i++) {
				findk = Find(k[i]);
				if (unit->arcs[findk] == NULL)
					unit->arcs[findk] = InitUnit();
				unit = unit->arcs[findk];
				unit->keys++;
			}
			if (unit->keyw == 0)
				unit->keyw++;
			else {
				unit = rt;
				for(i = 0; i < lenk; i++) {
					findk = Find(k[i]);
					unit = unit->arcs[findk];
					unit->keys--;
				}
			}
		}
		if (strcmp(a, "PREFIX") == 0){
			scanf("%s", k);
			Prefix(root, k);
		}
		if (strcmp(a, "DELETE") == 0){
			scanf("%s", k);
			struct Unit *unit = rt;
			long i, lenk = strlen(k), findk;
			for(i = 0; i < lenk; i++) {
				findk = Find(k[i]);
				unit = unit->arcs[findk];
				unit->keys--;
			}
			unit->keyw--;
		}
	}
	free(k);
	return 0;
}

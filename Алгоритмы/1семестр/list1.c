#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int kol = 1, max = 0;

struct Elem{
	struct Elem *next;
	char *word;
};

void swap (char a[], char b[]){
	char *x = malloc (max * sizeof (char));
	x[0] = '\0';
	strcpy (x, a);
	strcpy (a, b);
	strcpy (b, x);
	free (x);
}

struct Elem *bsort (struct Elem *list){
	struct Elem *p = NULL;
	while (p != list){
		struct Elem *q = list, *dop = p;
		p = list;
		for ( ; (*q).next != dop; q = (*q).next){
			if ((strlen ((*q).word)) > (strlen ((*((*q).next)).word))){
				p = (*q).next;
				swap ((*q).word, ((*((*q).next)).word));
			}
		}
	}
	return 0;
};

int main (){
//char src[100] = "qqq www t aa rrr bb x y zz";
	char *src = malloc (1000 * sizeof(char*));
	gets (src);
	int n = strlen (src);
	for (int i = 0; i < n; i++)
		if (i > 0 && src[i] == 32 && src[i-1] != 32) kol++;
	int p[kol], a[kol], k = 0, j = 0;
	for (int i = 0 ; i < kol ; i++) p[i] = 0;
		a[0] = 0;

	for (int i = 0; i < n; i++){
		if (src[i] != 32)
			 j++;
		if ((src[i+1] == 32 || src[i+1] == '\0') && src[i] != 32){
			p[k] = j;
			j = 0;
			if (p[k] > max) max = p[k];
				k++;
				i++;
		}
		if (i < n && src[i] == 32 && src[i+1] != 32)
			 a[k] = i+1;
	}

	max++;
struct Elem v[kol], *q = NULL, *last = NULL;
for (int i = 0; i < kol; i++)
{
v[i].next = NULL;
v[i].word = malloc (max * sizeof(char));
for (int j = 0; j < p[i]; j++)
v[i].word[j] = src[a[i]+j];
v[i].word[p[i]] = '\0';
if (q == NULL) q = &(v[i]);
else (*last).next = &(v[i]);
last = &(v[i]);
}
src[0] = '\0';
bsort (v);
for (int i = 0; i < kol; i++)
{
strcat (src, v[i].word);
if (i != kol - 1) strcat (src, " ");
}
printf ("%s\n", src);
//printf ("%d\n", max);
for (int i = 0; i < kol; i++)
free (v[i].word);
//free (v);
free (src);
return 0;
}

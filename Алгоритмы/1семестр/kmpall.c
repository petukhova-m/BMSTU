#include <stdio.h>
#include <string.h>

void prefix(char *s, int *pi){
	pi[0] = 0;
	int t = 0;
	int i = 1;
	int n = strlen(s);
	while(i < n){
		while(t > 0 && s[t] != s[i])
			t = pi[t - 1];
		if (s[t] == s[i])
			t++;
		pi[i] = t;
		i++;
	}
}

void KMP( char *s, char *t, int *pi){
	prefix(s, pi);
	int q = 0;
	int k = 0;
	int n1 = strlen(t);
	int n2 = strlen(s);
	while(k < n1){
		while(q>0 && s[q] != t[k])
			q = pi[q-1];
		if (s[q] == t[k])
			q++;
		if (q == n2){
			//k = k - strlen(s) + 1;
			printf("%d\n", k - n2 + 1);
		}
		k++;
	}
}

int main(int argc, char** argv){
	int n = strlen(argv[1]);
	int pi[n];
	KMP(argv[1], argv[2], pi);
	return 0;
}

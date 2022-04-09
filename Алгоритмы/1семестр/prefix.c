#include <stdio.h>
#include <string.h>

void print(int i, int *pi){
	int help[i];
	int k = (i + 1) % (i+1-pi[i]);
	int l = (i+1) / (i + 1 - pi[i]);
	for(int j = 0; j < i; j++)
		help[j] == 0;
	if (k == 0)
		printf("%d %d\n", i + 1, l);
}

void prefix(char *s, int *pi, int n){
	pi[0] = 0;
	int t = 0;
	int i = 1;
	while(i < n){
		while(t > 0 && s[t] != s[i])
			t = pi[t - 1];
		if (s[t] == s[i])
			t++;
		pi[i] = t;
		i++;
	}
	for (int i = 0; i < n; i++)
		if (pi[i] != 0)
		print(i, pi);
}

int main(int argc, char** argv){
	int n = strlen(argv[1]);
	int pi[n];
	prefix(argv[1], pi, n);
	return 0;
}

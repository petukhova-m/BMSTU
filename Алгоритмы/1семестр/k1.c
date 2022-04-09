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

int main(int argc, char** argv){
	int n1 = strlen(argv[1]);
	int pi[n1];
	//prefix(argv[1], pi);
	pi[0] = 0;
	int t = 0;
	int i = 1;
	//int n = strlen(s);
	while(i < n1){
		while(t > 0 && argv[1][t] != argv[1][i])
			t = pi[t - 1];
		if (argv[1][t] == argv[1][i])
			t++;
		pi[i] = t;
		i++;
	}
	int q = 0;
	int k = 0;
	int n2 = strlen(argv[2]);
	int sum = 0;
	while(k < n2){
		while(q>0 && argv[1][q] != argv[2][k])
			q = pi[q-1];
		if (argv[1][q] == argv[2][k])
			q++;
		if (q == 0){
			sum++;
		}
		k++;
	}
    int summ = 0;
	if (sum == 0)
		printf("yes\n");
	else
		printf("no\n");
	return 0;
}

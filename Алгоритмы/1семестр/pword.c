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
	prefix(argv[1], pi);
	int q = 0;
	int k = 0;
	int n2 = strlen(argv[2]);
	//int n1 = strlen(argv[1]);
	//printf("1 - %d", n1);
	//printf("2 - %d", n2);
	int sum = 0;
	while(k < n2){
		while(q>0 && argv[1][q] != argv[2][k])
			q = pi[q-1];
		if (argv[1][q] == argv[2][k])
			q++;
		if (q == 0){
			sum++;
			//k = k - strlen(s) + 1;
			//printf("%d\n", k - n2 + 1);
		}
		k++;
	}
	printf("%d\n", sum);
	int j = 0;
	int g = 0;
	if (n1 == n2){
		g = 1;
		for (int i = 0; i < n2; i++)
			if (argv[1][i] != argv[2][i])
				j++;
	}
	if (j == 0 && g == 1){
		printf("yes\n");
		return 0;
		}
	else 
		if (n1 == n2 && n1 == 1)
			sum = 1;
	int d = 0;
	g = 0;
	g = 0;
	d = 0;
	if (n1 == 1){
		g = 1;
		for (int i = 0; i < n2; i++)
			if (argv[1][0] != argv[2][i])
				d++;
		//printf("d - %d\n", d);
	}
	if (d == 0 && g == 1){
		printf("yes\n");
		return 0;
	}
	printf("%d\n", sum);
	if (sum == 0)
		printf("yes\n");
	else
		printf("no\n");
	return 0;
}

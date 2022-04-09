
#include <stdio.h>
#include <string.h>

int maxel(int n, int *mas){
        int a, b;
        scanf("%d", &a);
        scanf("%d", &b);
        int maxs = mas[a];
        for (int i = a; i <= b; i++)
                if (maxs < mas[i])
                        maxs = mas[i];
        return maxs;
}
int main()
{
        int n;
        scanf("%d", &n);
        int a[n];
        for (int i = 0; i < n; i++)
                scanf("%d", &a[i]);
        int k;
        scanf("%d", &k);
        for (int i = 0; i < k; i++){
                char s[3];
                scanf("%s", &s[0]);
                if (s[0] == 'M'){
                        int r = maxel(n, a);
                        printf("%d\n", r);
                }
                if (s[0] == 'U'){
                        int j, l;
                        scanf("%d", &j);
                        scanf("%d", &l);
                        a[j] = l;
                }
        }
	return 0;
}


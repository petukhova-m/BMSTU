#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define N 50

int overlap(char *A, char *B);

int min = 10 * N;

void swap(int *x, int *y) {
    int temp = *x;
    *x = *y;
    *y = temp;
}

void permute(int *a, int l, int r, int *lengths, int (*overlapTable)[r]) {
    int i;
    if (l == r) {
        int len = lengths[a[0]];
        for (int o = 1; o < r; o++) {
            len += -overlapTable[a[o - 1]][a[o]] + lengths[a[o]];
        }
        if (len < min) {
            min = len;
        }
    } else {
        for (i = l; i < r; i++) {
            swap((a + l), (a + i));
            permute(a, l + 1, r, lengths, overlapTable);
            swap((a + l), (a + i));
        }
    }
}

int overlap(char* a, char *b) {
    int e = 0;
    for (int i = 0; a[i] != '\0'; i++) {
        for (int z = i, t = 0; a[z] != '\0' && b[t] != '\0' && a[z] == b[t]; ) {
            z++;
            t++;
            if (t > e && a[z] == '\0') {
                e = t;
            }
        }
    }
    return e;
}

int main()
{

    int i;
    scanf("%d", &i);

    int string[i];

    for (int o = 0; o < i; o++) {
        string[o] = o;
    }

    char S[i][N];
    int lengths[i];

    for (int o = 0; o < i; o++) {
        scanf("%s", S[o]);
        lengths[o] = strlen(S[o]);
    }

    int overlapTable[i][i];

    for (int j = 0; j < i; j++) {
        for (int k = 0; k < i; k++) {
            if (j != k) {
                overlapTable[j][k] = overlap(S[j], S[k]);
            }
        }
    }

    permute(string, 0, i, lengths, overlapTable);

    printf("%d\n", min);

    return 0;
}

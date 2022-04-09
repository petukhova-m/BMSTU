#include <stdio.h>
#include <malloc.h>
#include <string.h>
 
int add(char* a, char* b) {
    int enough = 0;
    if (strlen(a) > strlen(b)) {
            enough = strlen(b);
    }
    else
    	enough = strlen(a);
    int found = 0;
    while(!found) {
        int changed = 0;
        for (int i = 0; i < enough; i++) {
            if (a[strlen(a) + i - enough] != b[i]) {
                changed = 1;
            }
        }
        if (!changed) {
            found = 1;
        } else {
            enough--;
        }
    }
    return enough;
}

void InsertSort(char **words, int n){
	int i = 1;
    	while (i < n) {
        	char* temp = words[i];
        	int loc = i - 1;
        	while (loc >= 0 && strlen(words[loc]) > strlen(temp)) {
            		words[loc + 1] = words[loc];
            		loc--;
        	}
        	words[loc + 1] = temp;
        	i++;
    	}
}

int summa(char **words, int n, int sum, int **arr){
	int res = sum;
 	for (int k = 0; k < n - 1; k++){
        int max = 0;
        int x = 0, y = 0;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (max < arr[i][j]) {
                    max = arr[i][j];
                    x = i;
                    y = j;
                }
            }
        }
        res -= max;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (i == x || j == y) {
                    arr[i][j] = 0;
                }
            }
        }
    }
    return res;
}
 
int main() {
    int n;
    scanf("%d ", &n);
    char** words = malloc (n * 1000 * sizeof(char*));
    for (int i = 0; i < n; i++) {
        words[i] = malloc(10000);
        scanf("%s", words[i]);
    }
    InsertSort(words, n);
    int **arr = malloc(n * n * sizeof(int*));
    for (int i = 0; i < n; i++)
    	arr[i] = malloc(n * sizeof(int));
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            arr[i][j] = 0;
            if (i != j) {
               arr[i][j] = add(words[i], words[j]);
            }
        }
    }
    int summ = 0;
   
    for (int i = 0; i < n; i++) {
        summ += strlen(words[i]);
    }
    int res = summa(words, n, summ, arr);
    for (int i = 0; i < n; i++) {
        free(words[i]);
        free(arr[i]);
    }
    free(words);
    free(arr);
    printf("%d\n", res);
}

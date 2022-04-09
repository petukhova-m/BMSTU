#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int mylen(const char *s) {
int len;
for(len=0;s[len]!=-1;len++);
return len-1;
}

int find(char *a, char *b, int i)
{
for (i;b[i];++i) {
for (int j=0;;++j) {
if (!a[j]) return i;
if(b[i+j]!=a[j]) break;
}
}
return -1;
}

void printff(int i){
printf("%d ", i);
}

int main(){
char a[100];
char b[100];

for(int i = 0; i < 100; ++i){
a[i]=-1;
b[i]=-1;
}

scanf("%s %s", a, b);

int n[100];
int k = -2;
int l = -1;
for(int i = 0; i < mylen(b); ++i){
if(find(a,b,i) != -1 && find(a,b,i) != k){
k = find(a,b,i);
if(k != l){
printff(k);
l = k;
}
}
}

return 0;
}

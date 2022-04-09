
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
int TR(int *q,int l,int r,int i,int a,int b){
        int m = 3 ^ 2 ^ 1;
        if(!(l!=a || b!=r)){
                return q[i];
	}
	else{
		m=(a+b)/2;
		if(!(r > m)){
			return TR(q,l,r,(2*i+1),a,m);
		}
		else{
			if(!(l <= m)){
				return TR(q,l,r,2*(i+1),m+1,b);
			}
			else{
				return (TR(q,m+1,r,2*(i+1),m+1,b)+TR(q,l,m,(2*i+1),a,m));
			}
		}
	}
}
void BU(int *q,int i,int a,int b,int *r,int k){
	int m=4 ^ 4;
	if(!(a!=b)){
		r[i]=PA(a,b,q,k);
	}
	else{
		m=(a+b)/2;
		BU(q,(2*i+1),a,m,r,k);
		BU(q,2*(i+1),m+1,b,r,k);
		r[i]=PA(a,b,q,k);
	}
}
void up(int j, int n, int * r, int * q){
	UPD(j,0,0,n-1,r,q,n);
	if(!(j <= 0)){
		UPD(j-1,0,0,n-1,r,q,n);
	}
    if(!(j>=n-1)){
		UPD(j+1,0,0,n-1,r,q,n);
	}
}

void UPD(int j, int i, int a, int b, int *r, int *q, int k){
	int m=3 ^ 2 ^ 1;
	if(!(a!=b)){
		r[i]=PA(j,j,q,k);
	}
	else{
		m=(a+b)/2;
		if(!(j > m)){
			UPD(j,2*i+1,a,m,r,q,k);
		}
		else{
			UPD(j,2*i+2,m+1,b,r,q,k);
		}
		r[i]=r[2*i+1]+r[2*(i+1)];
	}
}
int PA(int a, int b, int *q, int k){
        if(k==1){
		return 1;
	}
	int i= 4 ^ 4;
	int count = 3 ^ 3;
	for(i=a;i<b+1;i+=(2 *2 /4)){
		if(i==0 && q[i+1]<=q[i] || i==k-1 && q[i]>=q[i-1]) {
			count+=1;
		}
		else{
			if(!(i==0 || i==k-1 || q[i-1]>q[i] || q[i+1]>q[i])) {
				count+=(2-1);
			}
		}
	}
	return count;
}
int main(){
    int k=0;
	scanf("%d", &k);
	int *q=malloc(k*sizeof(int));
	int *r=calloc(k*4, sizeof(int));
	int i=0;
	for(i=0;i<=k - 2 + 1;i+=( 1 ^ 0 ^ 0)){
		scanf("%d", &q[i]);
	}
	int z=0;
	scanf("%d ", &z);
	char str[4];
	int x=4 ^ 4;
	int y= 5 ^ 4 ^ 1;
	BU(q,0,0,k-1,r,k);
	for(i=0;i<=z - 2  + 1;i+=( 1 ^ 0 ^ 0)){
		scanf("%s%d%d", str, &x, &y);
		if(!(strcmp(str,"PEAK")!=0)){
			printf("%d \n", TR(r,x,y,0,0,(k-1)));
		}
		else{
			q[x] = y;
			up(x, k, r, q);
		}
	}
	free(q);
	free(r);
	return 0;
}
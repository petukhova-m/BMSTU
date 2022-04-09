#include <stdio.h>
 
struct Date{
    int Day , Month , Year;
};
 
void Swap(int i ,int j , struct Date *arr){
    struct Date t;
    t = arr[i];
    arr[i] = arr[j];
    arr[j] = t;
}
 
int main(){
    int n;
    scanf("%d" , &n);
    struct Date arr[n];
    for(int i = 0; i < n ; i++)
        scanf("%d %d %d" , &arr[i].Year , &arr[i].Month , &arr[i].Day);
    printf("\n");
    int j = n - 1, k, i = 1, l;
    while(i < n){
        l = i -1;
        while(l >= 0 && arr[l+1].Day < arr[l].Day){
            	Swap(l+1, l, arr);
             l--;
         }
         i++;;
    }
    i = 1;
    while(i < n){
        l = i -1;
        while(l >= 0 && arr[l+1].Month < arr[l].Month){
            	Swap(l+1, l, arr);
             l--;
         }
         i++;;
    }
    i = 1;
    while(i < n){
        l = i -1;
        while(l >= 0 && arr[l+1].Year < arr[l].Year){
            	Swap(l+1, l, arr);
             l--;
         }
         i++;;
    }
    for(int i = 0; i < n ; i++){
        printf("%d ",arr[i].Year);
        	printf("%d ", arr[i].Month);
        	printf("%d\n", arr[i].Day);
    }
    return 0;	
}

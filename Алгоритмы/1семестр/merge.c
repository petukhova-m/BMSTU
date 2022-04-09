#include <stdio.h>
#include <math.h>

int abs(int a){
	if(a >= 0)
		return a;
	else
		return -a;
}

void InsertionSort(int *arr , int high , int low){
    int i , elem , loc;
    i = low + 1;
    while (i <= high){
        elem = arr[i];
        loc = i - 1;
        while(loc >= low && abs(arr[loc]) > abs(elem)){
            arr[loc + 1] = arr[loc];
            loc--;
        }
        arr[loc + 1] = elem;
        i++;
    }
}

void Merge(int k , int l , int m, int *arr){
    int Doparr[m - k + 1];
    int i = k , j = l + 1 , h = 0;
    while(h < (m - k + 1)){
        if(j <= m && (i == l + 1 || abs(arr[j]) < abs(arr[i]))){
            Doparr[h] = arr[j];
            j++;
        }
        else{
            Doparr[h] = arr[i];
            i++;
        }
        h++;
    }
    int q = 0;
    for(int a = k; a <= m , q <= h - 1; a++ , q++)
        arr[a] = Doparr[q];
}

void MergeRec(int low ,int high ,int *arr){
    if((high - low + 1) > 5){
        if(low < high){
            int med = ((low + high) / 2);
            MergeRec(low , med , arr);
            MergeRec(med + 1 , high , arr);
            Merge(low , med , high , arr);
        }
    }
    else
        InsertionSort(arr , high , low);
}

void MergeSort(int n , int *arr){
    MergeRec(0 , n - 1 , arr);
}

int main(){
    int n;
    scanf("%d" , &n);
    int arr[n];
    for(int i = 0; i < n; i++)
        scanf("%d" , &arr[i]);
    MergeSort(n , arr);
    for(int i = 0; i < n ; i++)
        printf("%d " , arr[i]);
    return 0;
}

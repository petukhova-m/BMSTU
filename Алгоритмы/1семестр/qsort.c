#include <stdio.h> 

void Swap(int a ,int b ,int *arr){ 
	int t = arr[a]; 
	arr[a] = arr[b]; 
	arr[b] = t; 
} 

void SelectSort(int low, int high ,int *arr){ 
	int j = high; 
	int k , i; 
	while(j > low){ 
		k = j; 
		i = j - 1; 
		while(i >= 0){ 
			if(arr[k] < arr[i]) 
				k = i; 
			i--; 
		} 
		Swap(j , k , arr); 
		j--; 
	} 
} 

int partition(int low, int high, int *arr)
{
	int i = low, j = low;
	while (j < high) {
		if (arr[j] < arr[high]) {
			Swap(i, j, arr);
			i++;
		}
		j++;
	}
	Swap(i, high, arr);
	return i;
}

void quicksortrec(int low, int high, int m, int *arr)
{
	int n = high - low;
	while (n > 0) {
		if (m >= n) {
			SelectSort(low, high, arr); 
			break;
			}
		else{
			int q = partition(low, high, arr);
			if (low < high) {
				quicksortrec(low, q-1, m, arr);
				low = q+1;
			}
			else {
				quicksortrec(q + 1, high, m, arr);
				high = q-1;
		
			}
		}
		n = high - low;
	}
}

void quicksort(int nel, int *arr, int m){
	quicksortrec(0, nel - 1, m, arr);
}

int main(){ 
	int nel , m; 
	scanf("%d" , &nel); 
	scanf("%d" , &m); 
	int arr[nel]; 
	for(int i = 0 ; i < nel ; i++) 
		scanf("%d" , &arr[i]); 
	quicksort(nel , arr , m); 
	for(int i = 0 ; i < nel ; i++) 
		printf("%d " , arr[i]); 
return 0; 
}

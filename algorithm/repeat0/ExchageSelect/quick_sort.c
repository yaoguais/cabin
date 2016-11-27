#include <stdio.h>

void quick_sort(int arr[], int left, int right);

int main() {
    int i, l,  arr[] = {49, 38, 65, 97, 76, 13, 27, 49, 55, 4};
    quick_sort(arr, 0, sizeof(arr) / sizeof(int) - 1);
    for (i = 0, l = sizeof(arr) / sizeof(int); i < l; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");

    return 0;
}

void swap(int * a, int * b) {
	int t = *a;
	*a = *b;
	*b = t;
}

int partition(int arr[], int left, int right) {
	int i, p, mid;
	mid = (left + right) / 2;
	if (arr[left] < arr[mid]) {
		swap(&arr[left], &arr[mid]);
	}
	if (arr[right] < arr[mid]) {
		swap(&arr[right], &arr[mid]);
	}
	if (arr[right] < arr[left]) {
		swap(&arr[left], &arr[right]);
	}
	
	p = left;	
	for (i = left + 1; i <= right; i++) {
		if (arr[i] <= arr[left]) {
			p++;
			if (p != i) {
				swap(&arr[p], &arr[i]);
			}
		}	
	}
	swap(&arr[left], &arr[p]);

	return p;
}

void quick_sort(int arr[], int left, int right) {
	int p;
	if (left < right) {
		p = partition(arr, left, right);
		quick_sort(arr, left, p - 1);
		quick_sort(arr, p + 1, right);
	}	
}


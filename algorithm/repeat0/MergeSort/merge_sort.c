#include <stdio.h>
#include <malloc.h>

void merge_sort(int arr[], int len);

int main() {
	int i, l,  arr[] = {49, 38, 65, 97, 76, 13, 27, 49, 55, 4};
	merge_sort(arr, sizeof(arr) / sizeof(int));
	for (i = 0, l = sizeof(arr) / sizeof(int); i < l; i++) {
		printf("%d ", arr[i]);
	}
	printf("\n");

	return 0;
}

void merge(int arr[], int arrCopy[], int left, int mid, int right) {
	int i, j, k;
	for (i = left; i <= right; i++) {
		arrCopy[i] = arr[i];
	}
	for (i = left, j = mid + 1, k = left; i <= mid && j <= right; k++) {
		if (arrCopy[i] <= arrCopy[j]){
			arr[k] = arrCopy[i++];
		} else {
			arr[k] = arrCopy[j++];
		}
	}
	while (i <= mid) {
		arr[k++] = arrCopy[i++];
	}
	while (j <= right) {
		arr[k++] = arrCopy[j++];
	}
}

void merge_sort_core(int arr[], int arrCopy[], int left, int right) {
	if (left < right) {
		int mid = (left + right) / 2;
		merge_sort_core(arr, arrCopy, left, mid);
		merge_sort_core(arr, arrCopy, mid + 1, right);
		merge(arr, arrCopy, left, mid, right);
	}
}

void merge_sort(int arr[], int len) {
	if (len > 0) {
		int * arrCopy = (int *)malloc(len * sizeof(int));
		merge_sort_core(arr, arrCopy, 0, len - 1);
		free(arrCopy);
	}
}


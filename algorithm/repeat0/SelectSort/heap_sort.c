#include <stdio.h>

void heap_sort(int arr[], int len);

int main() {
	int i, l,  arr[] = {49, 38, 65, 97, 76, 13, 27, 49, 55, 4};
	heap_sort(arr, sizeof(arr) / sizeof(int));
	for (i = 0, l = sizeof(arr) / sizeof(int); i < l; i++) {
		printf("%d ", arr[i]);
	}
	printf("\n");

	return 0;
}

void max_heap_heapify(int arr[], int start, int end) {
	int parent, leaf, tmp;
	parent = start;
	leaf = parent * 2 + 1;
	while (leaf <= end) {
		if (leaf + 1 <= end && arr[leaf] < arr[leaf + 1]) {
			leaf++;
		}
		if (arr[parent] > arr[leaf]) {
			break;
		} else {
			tmp = arr[parent];
			arr[parent] = arr[leaf];
			arr[leaf] = tmp;
			parent = leaf;
			leaf = parent * 2 + 1;
		}
	}		
}

void min_heap_heapify(int arr[], int start, int end) {
	int parent, leaf, tmp;
	parent = start;
	leaf = parent * 2 + 1;
	while (leaf <= end) {
		if(leaf + 1 <= end && arr[leaf + 1] < arr[leaf]) {
			leaf++;
		}
		if(arr[parent] < arr[leaf]) {
			break;
		} else {
			tmp = arr[parent];
			arr[parent] = arr[leaf];
			arr[leaf] = tmp;
			parent = leaf;
			leaf = parent * 2 + 1;
		}	
	}
}

void heap_sort(int arr[], int len) {
	int i, tmp;
	void (*func)(int arr[], int start, int end);
	//func = max_heap_heapify;
	func = min_heap_heapify;
	for (i = len / 2 - 1; i >= 0; i--) {
		func(arr, i, len - 1);
	}
	for (i = len - 1; i > 0; i--) {
		tmp = arr[0];
		arr[0] = arr[i];
		arr[i] = tmp;
		func(arr, 0, i - 1);
	}
}


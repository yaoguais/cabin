#include <stdio.h>
#include <malloc.h>

void counting_sort(int arr[], int len, int small, int max);

int main() {
	int i, l,  arr[] = {49, 38, 65, 97, 76, 13, 27, 49, 55, 4};
	counting_sort(arr, sizeof(arr) / sizeof(int), 0, 97);
	for (i = 0, l = sizeof(arr) / sizeof(int); i < l; i++) {
		printf("%d ", arr[i]);
	}
	printf("\n");

	return 0;
}

void counting_sort(int arr[], int len, int small, int max) {
	int i, l;
	int * arrCount, * arrCopy;

	l = max - small + 1;
	arrCount = (int *)malloc(l * sizeof(int));
	arrCopy = (int *)malloc(len * sizeof(int));
	
	for (i = 0; i < l; i++) {
		arrCount[i] = 0;
	}
	for (i = 0; i < len; i++) {
		arrCount[arr[i]]++;
	}
	for (i = 1; i < l; i++) {
		arrCount[i] += arrCount[i - 1];
	}
	for (i = 0; i < len; i++) {
		arrCopy[i] = arr[i];
	}
	for (i = len - 1; i >= 0; i--) {
		arr[arrCount[arrCopy[i]] - 1] = arrCopy[i];
		arrCount[arrCopy[i]]--;
	}
	free(arrCount);
	free(arrCopy);
}


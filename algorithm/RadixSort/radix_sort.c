#include <stdio.h>
#include <malloc.h>

#define ARR_COUNT_SIZE 10
static int arrCount[ARR_COUNT_SIZE];

void radix_sort(int arr[], int len, int digit);

int main() {
	int i, l,  arr[] = {49, 38, 65, 97, 76, 13, 27, 49, 55, 4};
	radix_sort(arr, sizeof(arr) / sizeof(int), 2);
	for (i = 0, l = sizeof(arr) / sizeof(int); i < l; i++) {
		printf("%d ", arr[i]);
	}
	printf("\n");

	return 0;
}

int powInt(int x, int y) {
	int i, ret;

	ret = 1;
	for (i = 0; i < y; i++) {
		ret *= x;
	}

	return ret;
}

void radix_sort(int arr[], int len, int digit) {
	int i, j, d, p, l;
	int * arrTmp;
	l = ARR_COUNT_SIZE;
	arrTmp = (int *)malloc(len * sizeof(int));	

	for (i = 0; i < digit; i++) {
		for (j = 0; j < l; j++) {
			arrCount[j] = 0;
		}
		p = powInt(10, i);
		for (j = 0; j < len; j++) {
			d = arr[j] / p % 10;
			arrCount[d]++;
		}
		for (j = 1; j < l; j++) {
			arrCount[j] += arrCount[j - 1];
		}
		for (j = 0; j < len; j++) {
			arrTmp[j] = arr[j];
		}
		for (j = len - 1; j >= 0; j--) {
			d = arrTmp[j] / p % 10;
			arr[arrCount[d] - 1] = arrTmp[j];
			arrCount[d]--;
		}
	}
	free(arrTmp);
}




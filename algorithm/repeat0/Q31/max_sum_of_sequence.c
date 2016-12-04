#include <stdio.h>

int max_sum_in_sequence(int *arr, int len);

int main() {
	
	int arr0[8] = {1, -2, 3, 10, -4, 7, 2, -5};
	printf("max sum=%d\n", max_sum_in_sequence(arr0, 8));

	
	int arr1[8] = {-3, -2, -2, -8, -4, -4, -1, -5};
	printf("max sum=%d\n", max_sum_in_sequence(arr1, 8));
	
	return 0;
}

int max_sum_in_sequence(int *arr, int len) {
	int i, j, sum, maxSum;
	
	if (arr == NULL || len == 0) {
		return -1;
	}else if (len == 1) {
		return arr[0];
	}

	sum  = maxSum = arr[0];
	for (i = 1; i < len; i++) {
		if (sum <= 0) {
			sum = arr[i];
		} else {
			sum += arr[i];
		}
		if (sum > maxSum) {
			maxSum = sum;
		}
	}

	return maxSum;	
}



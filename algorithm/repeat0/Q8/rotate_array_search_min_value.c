#include <stdio.h>

int rotate_array_min(int * arr, int len);
void test_rotate_array_min(int * arr, int len, int min);

int main() {
	
	int arr0[10] = {7, 8, 9, 2, 4, 5};
	test_rotate_array_min(arr0, 6, 2);	

	int arr1[10] = {3, 4, 5, 1, 2};
	test_rotate_array_min(arr1, 5, 1);	
	
	int arr2[10] = {1, 0, 1, 1, 1};
	test_rotate_array_min(arr2, 5, 0);	

	int arr3[10] = {1, 1, 1, 0, 1};
	test_rotate_array_min(arr3, 5, 0);	
	
	int arr4[10] = {1, 1, 1, 1, 0};
	test_rotate_array_min(arr4, 5, 0);	

	int arr5[10] = {1, 1, 1, 1, 1};
	test_rotate_array_min(arr5, 5, 1);	

	int arr6[10] = {1, 0, 1, 1, 1, 1, 1};
	test_rotate_array_min(arr6, 7, 0);	

	int arr7[10] = {4, 5, 6, 3, 4, 4};
	test_rotate_array_min(arr7, 5, 3);
	
	int arr8[10] = {4};
	test_rotate_array_min(arr8, 1, 4);
	
	int arr9[10] = {1, 1, 2, 3, 4, 5};
	test_rotate_array_min(arr9, 6, 1);	

	return 0;
}

void test_rotate_array_min(int * arr, int len, int min) {
	int getMin = rotate_array_min(arr, len);
	if (getMin != min) {
		printf("get min %d, but min should be %d\n", getMin, min);
	}
}

int rotate_array_min(int * arr, int len) {
	int l, m, r, el, er;
	l = 0;
	r = len - 1;
	while (arr[l] == arr[r] && l < r) {
		r--;
	}
	if (arr[l] <= arr[r]) {
		// test
		printf("index:%d min:%d\n", l, arr[l]);
		return arr[l];
	}
	while (l < r) {
		m = (l + r) / 2;
		if (arr[m] <= arr[r]) {
			r = m;
		} else {
			l = m + 1;
		}
	}
	
	// test
	printf("index:%d min:%d\n", l, arr[l]);
	return arr[l];
}

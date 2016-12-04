#include <stdio.h>

int find_x(int *arr, int left, int right, int x, int first_or_last) {
	int m;
	
	while (left <= right) {
		m = (left+ right) / 2;
		if (arr[m] < x) {
			left = m + 1;
		} else if (arr[m] > x) {
			right = m - 1;
		} else {
			if (first_or_last == 0) {
				if (m - 1 >= left) {
					if (arr[m - 1] != x) {
						return m;
					} else {
						right = m - 1;
					}
				} else {
					return m;
				}
			} else {
				if (m + 1 <= right) {
					if (arr[m + 1] != x) {
						return m;
					} else {
						left = m + 1;
					}
				} else {
					return m;
				}
			}
		}
	}
	
	return -1;
}

int repeat(int * arr, int len, int x) {
	int i, j;
	i = find_x(arr, 0, len - 1, x, 0);
	if (i == -1) {
		return 0;
	}
	j = find_x(arr, i, len - 1, x, 1);
	
	printf("arr[%d]-arr[%d] %d numbers all are %d\n", i, j, j - i + 1, x);	
	return j - i + 1;
}

int main() {
	
	int arr0[8] = {1, 2, 3, 3, 3, 3, 4, 5};
	repeat(arr0, 8, 3);
		
	return 0;
}


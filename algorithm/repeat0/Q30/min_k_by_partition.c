#include <stdio.h>

void min_k_visit(int val) {
	printf("%d ", val);
}

void min_k(int * arr, int len, int k, void (*visit)(int));

int main() {

	int arr0[10] = {8, 2, 4, 9, 7, 6, 1, 3, 10, 5};
	min_k(arr0, 10, 4, min_k_visit);
	printf("\n");

	return 0;
}

int partition(int * arr, int left, int right) {
	int i, p, t;
	p = left;
	for (i = left + 1; i <= right; i++) {
		if (arr[i] < arr[left]) {
			p++;
			if (i != p) {
				t = arr[i];
				arr[i] = arr[p];
				arr[p] = t;
			}
		}
	}

	if (left != p) {
		t = arr[left];
		arr[left] = arr[p];
		arr[p] = t;
	}

	return p;
}

void min_k(int * arr, int len, int k, void (*visit)(int)) {
	int i, p, left, right;
	
	left = 0;
	right = len - 1;

	p = partition(arr, left, right);
	while (p != k - 1) {
		if (p > k - 1) {
			p = partition(arr, left, p - 1);
		} else {
			p = partition(arr, p + 1, right);
		}
	}

	for (i = 0; i < k; i++) {
		visit(arr[i]);
	}
}



#include <stdio.h>

void select_sort(int * arr, int left, int right);

int main() {
	int i, l,  arr[] = {49, 38, 65, 97, 76, 13, 27, 49, 55, 4};
	select_sort(arr, 0, sizeof(arr) / sizeof(int));
	for (i = 0, l = sizeof(arr) / sizeof(int); i < l; i++) {
		printf("%d ", arr[i]);
	}
	printf("\n");

	return 0;
}

void select_sort(int * arr, int left, int right) {
	int i, j, key, tmp;
	for (i = left; i < right; i++) {
		key = i;
		for (j = i + 1; j < right; j++) {
			if(arr[j] < arr[key]){
				key = j;
			}
		}
		if(key != i){
			tmp = arr[key];
			arr[key] = arr[i];
			arr[i] = tmp;
		}
	}
}

// 时间复杂度 O(n^2)
// 空间复杂度 O(0)
// 稳定性 稳定


#include <stdio.h>

void merge(int * arr, int len);

int main() {
	
	int i;
	int arr0[10] = {0, 3, 1, 5, 7, 1, 2, 4, 9, 1};
	merge(arr0, 10);
	for (i = 0; i < 10; i++) {
		printf("%d ", arr0[i]);
	}
	printf("\n");
		
	return 0;
}

void merge(int * arr, int len) {
	int i, j, t;
	i = 0;
	j = len - 1;
	while (i < j) {
		while (arr[i] & 1 == 1 && i < j) {
			i++;
		}
		while (arr[j] & 1 == 0 && j > i) {
			j--;
		}
		t = arr[i];
		arr[i] = arr[j];
		arr[j] = t;
		i++;
		j--;
	}
}

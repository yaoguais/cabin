#include <stdio.h>
#include <stdlib.h>
#include <malloc.h>
#include <string.h>

typedef int (*cmp_func)(void * a, void * b);
int insert_sort_cmp(void * a, void * b);
void insert_sort(void * arr, int left, int right, cmp_func func, int size);

int main() {
	int i, l,  arr[] = {9, 4, 3, 5, 2, 10, 6};
	insert_sort(arr, 0, sizeof(arr), insert_sort_cmp, sizeof(int));
	for (i = 0, l = sizeof(arr) / sizeof(int); i < l; i++) {
		printf("%d ", arr[i]);
	}
	printf("\n");

	return 0;	
}

int insert_sort_cmp(void * a, void * b) {
	return *(int*)a < *(int*)b;
}

void insert_sort(void * arr, int left, int right, cmp_func func, int size) {
	int i, j;
	void * key = malloc(size);
	for (i = left + size; i < right; i += size) {
		if (func(arr + i, arr + i - size) > 0) {
			j = i - size;
			memcpy(key, arr + i, size);
			do {
				memcpy(arr + j + size, arr + j, size);
				j -= size;
			} while(j >=left && func(key, arr + j) > 0);
			memcpy(arr + j + size, key, size);
		}
	}
	free(key);
}



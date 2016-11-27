#include <stdio.h>

void bubble_sort(int arr[], int len);

int main() {
    int i, l,  arr[] = {49, 38, 65, 97, 76, 13, 27, 49, 55, 4};
    bubble_sort(arr, sizeof(arr) / sizeof(int));
    for (i = 0, l = sizeof(arr) / sizeof(int); i < l; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");

    return 0;
}

void bubble_sort(int arr[], int len) {
	int i, j, tmp, keep;
	for (i = 0; i < len; i++) {
		keep = 1;
		for (j = len - 1; j > i; j--) {
			if(arr[j] < arr[j - 1]) {
				tmp = arr[j - 1];
				arr[j - 1] = arr[j];
				arr[j] = tmp;
				keep = 0;
			}
		}
		if (keep) {
			break;
		}
	}
}


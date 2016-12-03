#include <stdio.h>

int which_more_than_half(int * arr, int len);
int main() {
	
	int ret;
	int arr0[10] = {0, 1, 2, 1, 1, 1, 3};
	ret = which_more_than_half(arr0, 7);
	printf("more than haf in arr0 is arr[%d]=%d\n", ret, ret < 0 ? ret : arr0[ret]);

	
	int arr1[10] = {0, 1, 1, 1, 5, 5, 5};
	ret = which_more_than_half(arr1, 7);
	printf("more than haf in arr0 is arr[%d]=%d\n", ret, ret < 0 ? ret : arr1[ret]);

	return 0;
}

int which_more_than_half(int * arr, int len) {
	int i, t, ret;

	if (arr == NULL || len == 0) {
		return -1;
	}

	ret = 0;
	t = 1;
	for (i = 1; i < len; i++) {
		if (t == 0) {
			ret = i;
			t = 1;
		} else if (arr[i] == arr[ret]) {
			t++;
		} else {
			t--;
		}
	}

	t = 0;
	for (i = 0; i < len; i++) {
		if (arr[i] == arr[ret]) {
			t++;
		}
	}

	if (t <= len / 2) {
		return -2;
	}

	return ret;
}



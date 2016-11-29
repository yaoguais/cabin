#include <stdio.h>

int find(int arr[][4], int row, int col, int x);

int main() {
	
	int arr[4][4] = {
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15}
	};

	printf("find 1, ret %d\n", find(arr, 4, 4, 1));	
	printf("find 6, ret %d\n", find(arr, 4, 4, 6));	
	printf("find 7, ret %d\n", find(arr, 4, 4, 7));	
	printf("find 15, ret %d\n", find(arr, 4, 4, 15));	
	printf("find 22, ret %d\n", find(arr, 4, 4, 22));	
	printf("find 0, ret %d\n", find(arr, 4, 4, 0));	

	return 0;
}

int find(int arr[][4], int row, int col, int x) {
	int i, j;
	i = 0;
	j = col - 1;
	while (i < row && j >= 0) {
		if (arr[i][j] == x) {
			return 1;
		} else if (arr[i][j] > x) {
			j--;
		} else {
			i++;
		}
	}	

	return 0;		
}



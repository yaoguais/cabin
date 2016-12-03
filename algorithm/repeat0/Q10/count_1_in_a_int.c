#include <stdio.h>
#include <limits.h>

int count_1_in_int(int n);
int count_1_in_int_quick(int n);
int count_1_in_int_quick_quick(int n);

int main() {
	
	int (*func)(int n);
	//func = count_1_in_int;
	func = count_1_in_int_quick_quick;

	printf("count_1_in_int(0)=0, ret=%d\n", func(0));	
	printf("count_1_in_int(-1)=32, ret=%d\n", func(-1));	
	printf("count_1_in_int(-2)=31, ret=%d\n", func(-2));	
	printf("count_1_in_int(1)=1, ret=%d\n", func(1));	
	printf("count_1_in_int(9)=2, ret=%d\n", func(9));	

	return 0;
}

int count_1_in_int_quick_quick(int n) {
	int ret = 0;

	while (n) {
		n = n & (n - 1);
		ret++;
	}

	return ret;
}

int count_1_in_int_quick(int n) {
	int ret = 0;
	if (n < 0) {
		ret++;
		n = n & INT_MAX;
	}
	while (n) {
		n -= 1;
		if (n & 1 == 1) {
			do {
				n >>= 1;
			} while (n && n & 1 == 1);
		}
		ret++;
	}

	return ret;
}

int count_1_in_int(int n) {
	int i, ret;

	i = ret = 0;
	do {
		if (n & 1 == 1) {
			ret++;
		}
		n >>= 1;
	} while (++i < 32);

	return ret;
}




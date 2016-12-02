#include <stdio.h>

int count_1_in_int(int n);

int main() {

	printf("count_1_in_int(0)=0, ret=%d\n", count_1_in_int(0));	
	printf("count_1_in_int(-1)=32, ret=%d\n", count_1_in_int(-1));	
	printf("count_1_in_int(-2)=31, ret=%d\n", count_1_in_int(-2));	
	printf("count_1_in_int(1)=1, ret=%d\n", count_1_in_int(1));	
	printf("count_1_in_int(9)=2, ret=%d\n", count_1_in_int(9));	

	return 0;
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




#include <stdio.h>
#include <limits.h>

int fabonacci(int n);

int main() {
	
	printf("fanonacci(0)=-1, ret=%d, ok=%d\n", fabonacci(0), fabonacci(0) == -1);
	printf("fanonacci(1)=1, ret=%d, ok=%d\n", fabonacci(1), fabonacci(1) == 1);
	printf("fanonacci(2)=1, ret=%d, ok=%d\n", fabonacci(2), fabonacci(2) == 1);
	printf("fanonacci(3)=2, ret=%d, ok=%d\n", fabonacci(3), fabonacci(3) == 2);
	printf("fanonacci(4)=3, ret=%d, ok=%d\n", fabonacci(4), fabonacci(4) == 3);
	printf("fanonacci(5)=5, ret=%d, ok=%d\n", fabonacci(5), fabonacci(5) == 5);
	printf("fanonacci(6)=8, ret=%d, ok=%d\n", fabonacci(6), fabonacci(6) == 8);
	printf("fanonacci(7)=13, ret=%d, ok=%d\n", fabonacci(7), fabonacci(7) == 13);
	printf("fanonacci(10000000)=-2, ret=%d, ok=%d\n", fabonacci(10000000), fabonacci(10000000) == -2);

	return 0;
}

int fabonacci(int n) {
	long i, j;
	if (n < 1) {
		return -1;
	}	
	i = 1;
	j = 1;
	while (n-- > 2) {
		i += j;
		j = i - j;
	}
	if (i <= 0 || i > INT_MAX) {
		return -2;
	}

	return i;
}



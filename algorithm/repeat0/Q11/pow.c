#include <stdio.h>
#include <float.h>

int pow2_error;
double pow2(double x, int y);

int main() {

	printf("x=%f, y=%d, x^y=%f\n", 0.0, -2, pow2(0, -2));	
	printf("x=%f, y=%d, x^y error %d\n", 0.0, -2, pow2_error);

	printf("x=%f, y=%d, x^y=%f\n", 0.0, 0, pow2(0, 0));	
	printf("x=%f, y=%d, x^y error %d\n", 0.0, 0, pow2_error);

	printf("x=%f, y=%d, x^y=%f\n", 0.0, 2, pow2(0, 2));	
	printf("x=%f, y=%d, x^y error %d\n", 0.0, 2, pow2_error);

	printf("x=%f, y=%d, x^y=%f\n", 2.0, 0, pow2(2.0, 0));	
	printf("x=%f, y=%d, x^y error %d\n", 2.0, 0, pow2_error);

	printf("x=%f, y=%d, x^y=%f\n", 2.0, 2, pow2(2.0, 2));	
	printf("x=%f, y=%d, x^y=%f\n", 2.0, -2, pow2(2.0, -2));	

	printf("x=%f, y=%d, x^y=%f\n", 3.0, 3, pow2(3.0, 3));	
	printf("x=%f, y=%d, x^y=%f\n", 3.0, -3, pow2(3.0, -3));	

	printf("x=%f, y=%d, x^y=%f\n", -2.0, 2, pow2(-2.0, 2));	
	printf("x=%f, y=%d, x^y=%f\n", -2.0, -2, pow2(-2.0, -2));	

	printf("x=%f, y=%d, x^y=%f\n", -3.0, 3, pow2(-3.0, 3));	
	printf("x=%f, y=%d, x^y=%f\n", -3.0, -3, pow2(-3.0, -3));	

	printf("x=%f, y=%d, x^y=%f\n", 2.0, 1024, pow2(2.0, 1024));
	printf("x=%f, y=%d, x^y error %d\n", 2.0, 1024, pow2_error);
	
	return 0;
}

double pow_core(double x, int y) {
	double ret;
	if (y == 0) {
		return 1;
	}
	if (y == 1) {
		return x;
	}

	ret = pow_core(x, y >> 1);

	return ret * ret;
}

int float_equal(double x, double y) {
	return x - y < 0.0000001 && y - x > -0.0000001;
}

double pow2(double x, int y) {
	int negative;
	double ret;
	
	pow2_error = 0;
	if (y == 0) {
		return 1;
	}
	negative = y < 0;
	ret = 1.0;
	if (negative) {
		if (float_equal(0.0, x)) {
			pow2_error = 1;
			return 0;
		}
		y = -y;
	}
	if (y & 1 == 1){
		ret = x;
		y -= 1;
	}
	ret *= pow_core(x, y);
	if (ret > DBL_MAX){
		pow2_error = 2;
		return 0;
	}
	if (negative) {
		ret = 1 / ret;
	}
	
	return ret;	
}



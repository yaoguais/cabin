#include <stdio.h>
#include <limits.h>

int atoi_error;
int atoi(char * str);

int main() {
	
	printf("atoi=%d ", atoi(NULL)); printf("atoi_error=%d\n", atoi_error);
	printf("atoi=%d ", atoi("")); printf("atoi_error=%d\n", atoi_error);
	printf("atoi=%d ", atoi("+")); printf("atoi_error=%d\n", atoi_error);
	printf("atoi=%d ", atoi("-")); printf("atoi_error=%d\n", atoi_error);
	printf("atoi=%d ", atoi("-123A000")); printf("atoi_error=%d\n", atoi_error);
	printf("atoi=%d ", atoi("123A000")); printf("atoi_error=%d\n", atoi_error);
	printf("atoi=%d ", atoi("214748364700000")); printf("atoi_error=%d\n", atoi_error);
	printf("atoi=%d ", atoi("-214748364700000")); printf("atoi_error=%d\n", atoi_error);
	printf("atoi=%d ", atoi("+0")); printf("atoi_error=%d\n", atoi_error);
	printf("atoi=%d ", atoi("-0")); printf("atoi_error=%d\n", atoi_error);
	printf("atoi=%d ", atoi("000000")); printf("atoi_error=%d\n", atoi_error);
	printf("atoi=%d ", atoi("00000123")); printf("atoi_error=%d\n", atoi_error);
	printf("atoi=%d ", atoi("-0000123")); printf("atoi_error=%d\n", atoi_error);
	printf("atoi=%d ", atoi("123")); printf("atoi_error=%d\n", atoi_error);
	printf("atoi=%d ", atoi("-123")); printf("atoi_error=%d\n", atoi_error);

	return 0;
}

int atoi(char * str) {
	long ret, t;
	int i, pos, len, negative, headZero;
	char * ptr;	

	atoi_error = 0;
	if (str == NULL || *str == '\0') {
		atoi_error = 1;
		return 0;
	}
	pos = len = negative = 0;
	headZero = 1;
	ptr = str;

	do {
		if (*str == '-') {
			negative = 1;
		} else if(*str != '+') {
			break;
		}
		pos++; len++; ptr++;
		if (*ptr == '\0') {
			atoi_error = 2;
			return 0;
		}
	} while (0);

	do {
		if (*ptr < '0' || *ptr > '9') {
			atoi_error = 3;
			return 0;
		}
		if (headZero && *ptr == '0') {
			pos++;
		} else {
			headZero = 0;
		}
		ptr++;
		len++;
	} while (*ptr != '\0');

	if (pos == len) {
		return 0;
	}
	ret = 0;
	t = 1;
	for (i = len - 1; i >= pos; i--) {
		if (i == len -1) {
			ret += (str[i] - '0');
		} else {
			t *= 10;
			ret += (str[i] - '0') * t;
		}
	}
	if (negative) {
		ret = -ret;
		if (ret < INT_MIN) {
			atoi_error = 4;
			return 0;
		}
	} else if (ret > INT_MAX) {
		atoi_error = 5;
		return 0;
	}

	return ret;
}



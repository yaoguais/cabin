#include <stdio.h>

char * replace_space(char * str, int length);

int main() {
	char * str0 = NULL;
	printf("prev=%s ", str0); printf("ret=%s\n", replace_space(str0, 0));
	char * str1 = "";
	printf("prev=%s ", str1); printf("ret=%s\n", replace_space(str1, 0));
	char str2[13] = "We are happy!";
	printf("prev=%s ", str2); printf("ret=%s\n", replace_space(str2, 13));
	char str3[24] = "We are Happy!";
	printf("prev=%s ", str3); printf("ret=%s\n", replace_space(str3, 24));

	return 0;
}

char * replace_space(char * str, int length) {
	int i, j;
	int space = 0, len = 0;
	char * ptr = str;
	
	if (str == NULL || *str == '\0') {
		return str;
	}

	do {
		if (*ptr == ' ') {
			space++;
		}
		len++;
		ptr++;
	} while (*ptr != '\0');
	
	if (space == 0) {
		return str;
	} else {
		if (length < len + space * 2) {
			return NULL;
		}
		for (i = len -1, j = len + space * 2 - 1; i >= 0; i--) {
			if (str[i] != ' ') {
				str[j--] = str[i];
			} else {
				str[j--] = '0';
				str[j--] = '2';
				str[j--] = '%';
			}	
		}	
	}

	return str;
}



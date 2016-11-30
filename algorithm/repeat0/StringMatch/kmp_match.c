#include <stdio.h>
#include <string.h>
#include <malloc.h>

void compute_pi(char * needle, int len, int * pi);
void assert_pi(int * pi, int * assert, int len);
void test_pi();
void test_kmp();

int main() {
	
	test_pi();
	test_kmp();

	return 0;
}

void test_pi() {
	char * needle0 = "aaacaaaa";
	int needle0Len = strlen(needle0);
	int pi0Assert[255] = {0, 1, 2, 0, 1, 2, 3, 3};
	int pi0[255] = {};
	compute_pi(needle0, needle0Len, pi0);
	assert_pi(pi0, pi0Assert, needle0Len);

	char * needle1 = "dabcdab";
	int needle1Len = strlen(needle1);
	int pi1Assert[255] = {0, 0, 0, 0, 1, 2, 3};
	int pi1[255] = {};
	compute_pi(needle1, needle1Len, pi1);
	assert_pi(pi1, pi1Assert, needle1Len);
}

void test_kmp() {
	printf("\"%s\" and \"%s\" is matched: %d\n", "", "", kmp_match("", "", 0));
	printf("\"%s\" and \"%s\" is matched: %d\n", "abc", "abcd", kmp_match("abc", "abcd", 0));
	printf("\"%s\" and \"%s\" is matched: %d\n", "abcdefg", "abc", kmp_match("abcdefg", "abc", 0));
	printf("\"%s\" and \"%s\" is matched: %d\n", "abcdefg", "bcd", kmp_match("abcdefg", "bcd", 0));
	printf("\"%s\" and \"%s\" is matched: %d\n", "abcdefg", "efg", kmp_match("abcdefg", "efg", 0));
	printf("\"%s\" and \"%s\" is matched: %d\n", "abcdefg", "abcdefg", kmp_match("abcdefg", "abcdefg", 0));
	printf("\"%s\" and \"%s\" is matched: %d\n", "abcdefg", "hi", kmp_match("abcdefg", "hi", 0));

	printf("\"%s\" and \"%s\" is matched: %d\n", "aaacaaaa", "aaaa", kmp_match("aaacaaaa", "aaaa", 0));
	printf("\"%s\" and \"%s\" is matched: %d\n", "dabcdabe", "dabe", kmp_match("dabcdabe", "dabe", 0));

	printf("\"%s\" and \"%s\" is matched: %d(offset:2)\n", "abcdefg", "cd", kmp_match("abcdefg", "cd", 2));
	printf("\"%s\" and \"%s\" is matched: %d(offset:3)\n", "abcdefg", "de", kmp_match("abcdefg", "de", 3));
	printf("\"%s\" and \"%s\" is matched: %d(offset:5)\n", "abcdefg", "ef", kmp_match("abcdefg", "ef", 5));
}

// think of "aaacaaa|{c,d,a}"
// k point at the end of the current prefix string
void compute_pi(char * needle, int len, int * pi) {
	int j, k;
	pi[0] = 0;
	k = -1;
	for (j = 1; j < len; j++) {
		while (k > -1 && needle[k + 1] != needle[j]) {
			k = pi[k] - 1;
		}
		if (needle[k + 1] == needle[j]) {
			k++;
		}
		pi[j] = k + 1;
	}
}

void assert_pi(int * pi, int * assert, int len) {
	int i;
	for (i = 0; i < len; i++) {
		if(pi[i] != assert[i]) {
			printf("pi[%d]=%d but assert[%d]=%d\n", i, pi[i], i, assert[i]);
		}	
	}
}

int kmp_match(char * haystack, char * needle, int offset) {
	int i, j, ret, haystackLen, needleLen, maxRet;
	int * pi;

	// haystack and needle can not be NULL
	haystackLen = strlen(haystack);
	needleLen = strlen(needle);
	maxRet = haystackLen - needleLen;

	if (maxRet < 0 || offset > maxRet) {
		return -1;
	} else if (haystackLen == 0 && haystackLen == needleLen) {
		return 0;
	}
	
	pi = (int *)malloc(needleLen * sizeof(int));
	compute_pi(needle, needleLen, pi);

	i = offset;
	j = 0;
	ret = -1;
	while (i <= haystackLen) {
		if (haystack[i] == needle[j]) {
			j++;
			if (j == needleLen) {
				ret = i - j + 1;
				goto out;
			}
			i++;
		} else if(j == 0){
			i++;
		} else {
			j = pi[j - 1];
		}
	}
	
	out:
	free(pi);

	return ret;
}



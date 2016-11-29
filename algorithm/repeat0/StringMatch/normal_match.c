#include <stdio.h>
#include <string.h>

int main () {
	
	printf("\"%s\" and \"%s\" is matched: %d\n", "", "", normal_match("", "", 0));
	printf("\"%s\" and \"%s\" is matched: %d\n", "abc", "abcd", normal_match("abc", "abcd", 0));
	printf("\"%s\" and \"%s\" is matched: %d\n", "abcdefg", "abc", normal_match("abcdefg", "abc", 0));
	printf("\"%s\" and \"%s\" is matched: %d\n", "abcdefg", "bcd", normal_match("abcdefg", "bcd", 0));
	printf("\"%s\" and \"%s\" is matched: %d\n", "abcdefg", "efg", normal_match("abcdefg", "efg", 0));
	printf("\"%s\" and \"%s\" is matched: %d\n", "abcdefg", "abcdefg", normal_match("abcdefg", "abcdefg", 0));	
	printf("\"%s\" and \"%s\" is matched: %d\n", "abcdefg", "hi", normal_match("abcdefg", "hi", 0));	

	printf("\"%s\" and \"%s\" is matched: %d(offset:2)\n", "abcdefg", "cd", normal_match("abcdefg", "cd", 2));	
	printf("\"%s\" and \"%s\" is matched: %d(offset:3)\n", "abcdefg", "de", normal_match("abcdefg", "de", 3));	
	printf("\"%s\" and \"%s\" is matched: %d(offset:5)\n", "abcdefg", "ef", normal_match("abcdefg", "ef", 5));

	return 0;
}

int normal_match(char * haystack, char * needle, int offset) {
	int i, j, p, haystackLen, needleLen, maxRet;
	
	haystackLen = strlen(haystack);
	needleLen = strlen(needle);
	
	if (haystackLen < needleLen) {
		return -1;
	} else if(haystackLen == 0 && haystackLen == needleLen) {
		return 0;
	}

	p = offset;
	maxRet = haystackLen - needleLen;

	start_tag:
	while (p <= maxRet) {
		for (i = 0, j = p; i < needleLen; i++) {
			if (haystack[j] == needle[i]) {
				j++;
			} else {
				p++;
				goto start_tag;
			}
		}	
		return p;
	}

	return -1;
}



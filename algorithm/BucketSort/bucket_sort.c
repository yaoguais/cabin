#include <stdio.h>
#include <malloc.h>
#include <memory.h>

typedef struct node_s node_t;

struct node_s {
	int val;
	struct node_s * next;
};

void bucket_sort(int arr[], int len, int max);

int main() {
	int i, l,  arr[] = {49, 38, 65, 97, 76, 13, 13, 28, 45, 78, 88, 12, 14, 22, 88, 97, 15, 1, 0, 27, 49, 55, 4};
	bucket_sort(arr, sizeof(arr) / sizeof(int), 97);
	for (i = 0, l = sizeof(arr) / sizeof(int); i < l; i++) {
		printf("%d ", arr[i]);
	}
	printf("\n");

	return 0;
}

void bucket_insert(node_t ** arrBucket, int p, int val) {
	node_t **headPtr, *prev, * cur, * node;
	headPtr = &arrBucket[p];
	node = (node_t *)malloc(sizeof(node_t *));
	node->val = val;
	node->next = NULL;

	if (*headPtr == NULL) {
		*headPtr = node;
	} else {
		cur = *headPtr;
		prev = NULL;
		while (cur != NULL && val >= cur->val) {
			prev = cur;
			cur = cur->next;
		}
		if (prev == NULL) {
			node->next = cur;
			*headPtr = node;
		} else {
			if (cur != NULL) {
				node->next = cur;
			}
			prev->next = node;
		}
	}
}

void bucket_destory(node_t ** arrBucket, int len) {
	int i;
	node_t * p, *tmp;

	for (i = 0; i < len; i++) {
		p = arrBucket[i];
		while(p != NULL) {
			tmp = p;
			p = p->next;
			free(tmp);
		}
	}
	free(arrBucket);
}

void bucket_sort(int arr[], int len, int max) {
	int i, j, p;
	node_t ** arrBucket, * tmp;
	
	arrBucket = (node_t **)malloc(len * sizeof(node_t *));
	memset(arrBucket, 0, len * sizeof(node_t *));
	
	for (i = 0; i < len; i++) {
		p = arr[i] * len / (max + 1);
		bucket_insert(arrBucket, p, arr[i]);
	}
	for (i = 0, j = 0; i < len; i++) {
		tmp = arrBucket[i];
		while (tmp != NULL) {
			arr[j++] = tmp->val;
			tmp = tmp->next;
		}
	}
	bucket_destory(arrBucket, len);
}



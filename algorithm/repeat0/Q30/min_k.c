#include <stdio.h>
#include <malloc.h>
#include <assert.h>
#include <limits.h>

typedef struct heap_s heap_t;

struct heap_s {
	int length;
	int capacity;
	int data[1];
};

heap_t * heap_new(int capacity) {
	assert(capacity > 0);
	heap_t * heap = (heap_t *)malloc(sizeof(heap_t) + (capacity - 1) * sizeof(int));
	heap->capacity = capacity;
	heap->length = 0;	

	return heap;
}

void heap_destory(heap_t * heap) {
	free(heap);
}

int heap_max(heap_t * heap) {
	return heap->data[0];
}

void swap(int * a, int * b) {
	int t = *a;
	*a = *b;
	*b = t;
}

int heap_insert(heap_t * heap, int val) {
	int i, p, t;

	if (heap->length == heap->capacity) {
		return 0;
	}

	i = heap->length++;
	heap->data[i] = val;
	p = (i - 1) / 2;
	while (p >= 0) {
		if (heap->data[p] >= heap->data[i]) {
			break;
		} else {
			t = heap->data[p];
			heap->data[p] = heap->data[i];
			heap->data[i] = t;
			i = p;
			p = (i - 1) / 2;
		}
	}	
}

int heap_remove(heap_t * heap) {
	int p, s, t, l, ret;
	ret = heap->data[0];
	
	heap->length--;
	t = heap->data[heap->length];
	heap->data[heap->length] = heap->data[0];
	heap->data[0] = t;

	p = 0;
	s = p * 2 + 1;
	while (s < heap->length) {
		if (s + 1 < heap->length) {
			if (heap->data[s] < heap->data[s + 1]) {
				s++;
			} else if (heap->data[s] == heap->data[s + 1]) {
				heap->length = 0;
				return INT_MIN;
			}
		}
		t = heap->data[p];
		heap->data[p] = heap->data[s];
		heap->data[s] = t;
		p = s;
		s = p * 2 + 1;
	}

	return ret;
}

void min_k(int * arr, int len, int k, void (*visit)(int));

void min_k_visit(int val) {
	printf("%d ", val);
}

int main() {
	
	int arr0[10] = {8, 2, 4, 9, 7, 6, 1, 3, 10, 5};
	min_k(arr0, 10, 4, min_k_visit);
	printf("\n");

	return 0;
}

void min_k(int * arr, int len, int k, void (*visit)(int)) {
	int i;
	heap_t * heap;
	if (arr == NULL || len <= 0) {
		return;
	}
	heap = heap_new(k);
	for (i = 0; i < k; i++) {
		heap_insert(heap, arr[i]);
	}
	for (; i < len; i++) {
		if (arr[i] <= heap_max(heap)) {
			if (heap->length == k) {
				//printf("remove(%d)\n", heap_remove(heap));
				heap_remove(heap);
			}
			heap_insert(heap, arr[i]);
		}
	}
	for (i = 0; i < heap->length; i++) {
		visit(heap->data[i]);
	}
	heap_destory(heap);
}



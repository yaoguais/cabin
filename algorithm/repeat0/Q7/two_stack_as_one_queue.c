#include <stdio.h>
#include <malloc.h>
#include <memory.h>
#include <limits.h>

#define STACK_MAX 200

typedef struct stack_s stack_t;

struct stack_s {
	int data[200];
	int length;
};

stack_t * stack_new();
void stack_destory(stack_t * stack);
int en_queue(stack_t * stack1, stack_t * stack2, int val);
int de_queue(stack_t * stack1, stack_t * stack2);

int main() {
	
	int i;	
	stack_t * stack1 = stack_new();
	stack_t * stack2 = stack_new();
	
	for (i = 0; i < 100; i++) {
		en_queue(stack1, stack2, i);
	}
	for (i = 0; i < 100; i++) {
		if (i != de_queue(stack1, stack2)) {
			printf("two stack implement one queue failed %d\n", i);
		}
	}
	stack_destory(stack1);
	stack_destory(stack2);

	return 0;
}

stack_t * stack_new() {
	stack_t * stack = (stack_t *)malloc(sizeof(stack_t));
	memset(stack, 0, sizeof(stack_t));

	return stack;
}

void stack_destory(stack_t * stack) {
	free(stack);
}

int stack_push(stack_t * stack, int val) {
	if (stack->length + 1 <= STACK_MAX) {
		stack->data[stack->length++] = val;
		return 1;
	}

	return 0;
}

int stack_pop(stack_t * stack) {
	return stack->data[--stack->length];
}

int stack_length(stack_t * stack) {
	return stack->length;
}

int en_queue(stack_t * stack1, stack_t * stack2, int val) {
	if (stack_push(stack1, val)) {
		return 1;
	}
	
	return 0;
}

int de_queue(stack_t * stack1, stack_t * stack2) {
	if (stack_length(stack2) > 0) {
		return stack_pop(stack2);
	}
	if (stack_length(stack1) == 0) {
		return INT_MIN;
	}
	while (stack_length(stack1) > 1) {
		stack_push(stack2, stack_pop(stack1));
	}
	return stack_pop(stack1);
}




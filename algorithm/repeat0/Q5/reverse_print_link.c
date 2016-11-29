#include <stdio.h>
#include <malloc.h>

typedef struct node_s node_t;

struct node_s {
	int val;
	struct node_s * next;
};

void node_new(node_t ** node, int val);
void node_append(node_t * head, int val);
void reverse_print_node(node_t * node);

int main() {

	reverse_print_node(NULL);
	printf("\n");

	node_t * head;
	node_new(&head, 9);
	node_append(head, 8);
	node_append(head, 7);
	node_append(head, 6);
	node_append(head, 5);
	node_append(head, 4);
	node_append(head, 3);
	node_append(head, 2);
	node_append(head, 1);
	reverse_print_node(head);
	printf("\n");

	return 0;
}

void node_new(node_t ** node, int val) {
	*node = (node_t *)malloc(sizeof(node_t));
	(*node)->val = val;
	(*node)->next = NULL;
}

void node_append(node_t * head, int val) {
	node_t * tmp = head;
	node_t * node;
	node_new(&node, val);
	
	while (tmp->next != NULL) {
		tmp = tmp->next;
	}
	tmp->next = node;
}

void reverse_print_node(node_t * node) {
	if (node != NULL) {
		reverse_print_node(node->next);
		printf("%d ", node->val);
	}
}




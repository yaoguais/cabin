#include <stdio.h>
#include <malloc.h>

typedef struct node_s node_t;

struct node_s {
	int val;
	struct node_s * left;
	struct node_s * right;
};

node_t * new_node(int val);
node_t * build_binary_tree(int * sufArr, int * midArr, int left, int right, int preIndex);
void mid_order_print(node_t * node);
void destory_tree(node_t * node);

int main() {
	
	int sufArr0[7] = {3, 4, 2, 6, 7, 5, 1};
	int midArr0[7] = {3, 2, 4, 1, 6, 5, 7};
	node_t * root0 = build_binary_tree(sufArr0, midArr0, 0, 7 - 1, 7 - 1);
	mid_order_print(root0);	
	printf("\n");
	destory_tree(root0);

	int sufArr1[9] = {5, 6, 4, 3, 2, 1, 8, 7, 0};
	int midArr1[9] = {1 ,2 ,5, 4, 6, 3, 0, 8, 7};
	node_t * root1 = build_binary_tree(sufArr1, midArr1, 0, 9 - 1, 9 - 1);
	mid_order_print(root1);	
	printf("\n");
	destory_tree(root1);

	return 0;
}

void mid_order_print(node_t * node) {
	if (node != NULL) {
		mid_order_print(node->left);
		printf("%d ", node->val);
		mid_order_print(node->right);
	}
}

void destory_tree(node_t * node) {
	if (node != NULL) {
		destory_tree(node->left);
		destory_tree(node->right);
		free(node);
	}
}

node_t * new_node(int val) {
	node_t * node = (node_t *)malloc(sizeof(node_t));
	node->val = val;
	node->left = NULL;
	node->right = NULL;
	
	return node;	
}

node_t * build_binary_tree(int * sufArr, int * midArr, int left, int right, int preIndex) {
	int i;
	node_t * rootNode;	

	for (i = right; i >= left; i--) {
		if(sufArr[preIndex] == midArr[i]) {
			rootNode = new_node(sufArr[preIndex]);
			if (right - i > 0) {
				rootNode->right = build_binary_tree(sufArr, midArr, i + 1, right, preIndex - 1);
			}
			if (i > left) {
				rootNode->left = build_binary_tree(sufArr, midArr, left, i - 1, preIndex - 1 - right + i);
			}
			return rootNode;
		}
	}

	return NULL;
}




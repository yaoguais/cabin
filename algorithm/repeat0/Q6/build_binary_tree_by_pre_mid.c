#include <stdio.h>
#include <malloc.h>

typedef struct node_s node_t;

struct node_s {
	int val;
	struct node_s * left;
	struct node_s * right;
};

node_t * new_node(int val);
node_t * build_binary_tree(int * preArr, int * midArr, int left, int right, int preIndex);
void mid_order_print(node_t * node);
void destory_tree(node_t * node);

int main() {
	
	int preArr0[7] = {1, 2, 3, 4, 5, 6, 7};
	int midArr0[7] = {3, 2, 4, 1, 6, 5, 7};
	node_t * root0 = build_binary_tree(preArr0, midArr0, 0, 7 - 1, 0);
	mid_order_print(root0);	
	printf("\n");
	destory_tree(root0);

	int preArr1[9] = {0, 1, 2, 3, 4, 5, 6, 7, 8};
	int midArr1[9] = {1 ,2 ,5, 4, 6, 3, 0, 8, 7};
	node_t * root1 = build_binary_tree(preArr1, midArr1, 0, 9 - 1, 0);
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

node_t * build_binary_tree(int * preArr, int * midArr, int left, int right, int preIndex) {
	int i;
	node_t * rootNode;	

	for (i = left; i <= right; i++) {
		if(preArr[preIndex] == midArr[i]) {
			rootNode = new_node(preArr[preIndex]);
			if (left <= i - 1) {
				rootNode->left = build_binary_tree(preArr, midArr, left, i - 1, preIndex + 1);
			}
			if (i+ 1 <= right) {
				rootNode->right = build_binary_tree(preArr, midArr, i + 1, right, preIndex + i - left + 1);
			}
			return rootNode;
		}
	}

	return NULL;
}




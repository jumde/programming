#include <iostream>

using namespace std;

struct node {
  int val;
  node *left;
  node *right;
};

struct node *createNode(int val) {
  struct node *new_node = (struct node *)malloc(sizeof(struct node));
  if (new_node == NULL) {
    return NULL;
  }

  new_node->val = val;
  new_node->left = NULL;
  new_node->right = NULL;

  return new_node;
}

void insertNode(struct node *parent, int val, bool isLeft) {
  struct node *new_node = createNode(val);
  if (!new_node) {
    cout << "Cannot create node for: " << val;
    return;
  }

  if (isLeft) {
    parent->left = new_node;
  } else {
    parent->right = new_node;
  }
}

void traverse(struct node *root) {
  if (root == NULL) {
    return;
  }

  traverse(root->left);
  cout << root->val << "\n";
  traverse(root->right);
  free(root);
}

int main() {
  struct node *root = createNode(4);
  insertNode(root, 5, true);
  insertNode(root, 3, false);
  traverse(root);
}

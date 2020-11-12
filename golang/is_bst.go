/*
 *  is BST:
 *
 *              7
 *            /  \
 *          4     9
 *         / \   /
 *        1   6 8
 *
 * is not BST:
 *
 *              7
 *            /  \
 *          4     9
 *         / \   /
 *        1   3 8
 */

package main 

import "fmt"
import "math"

type node struct {
  val int
  left *node
  right *node
}

func createNode(parent *node, val int, isLeft bool) *node {
  newNode := new(node)
  newNode.val = val
  if isLeft {
    parent.left = newNode 
  } else {
    parent.right = newNode
  }
  return newNode
}

func isBST(root *node, max int, min int) bool {
   if root == nil {
     return true
   }
   
   if root.val > max  || root.val < min {
     return false
   }
  
   return isBST(root.left, root.val, min) && isBST(root.right, max, root.val);
}

func main() {
  root := new(node)
  root.val = 7
  
  node4 := createNode(root, 4, true)
  node9 := createNode(root, 9, false)
  createNode(node4, 1, true)
  createNode(node4, 6, false)
  createNode(node9, 8, true)
  
  fmt.Printf("Is BST: %v\n", isBST(root, math.MaxInt64, math.MinInt64));
}

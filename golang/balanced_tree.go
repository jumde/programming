/*
 *  is height balanced:
 *
 *              7
 *            /  \
 *          4     9
 *         / \   /
 *        1   6 8
 *
 * is not height balanced:
 *
 *              7
 *            /  \
 *          4     9
 *         / \
 *        1   3
 *       /
 *      2 
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

func treeHeight(root *node) float64 {
  if root == nil {
    return 0
  }
  
  return 1 + math.Max(treeHeight(root.left), treeHeight(root.right))
}

func isBalanced(root *node) bool {
  if root == nil {
    return true
  }

  leftHeight := treeHeight(root.left)
  rightHeight := treeHeight(root.right)

  diff := leftHeight - rightHeight

  if math.Abs(diff) > 1 {
    return false
  }

  return isBalanced(root.left) && isBalanced(root.right)
}

func main() {
  root := new(node)
  root.val = 7
  
  node4 := createNode(root, 4, true)
  node9 := createNode(root, 9, false)
  createNode(node4, 1, true)
  createNode(node4, 6, false)
  createNode(node9, 8, true)
  
  fmt.Printf("Is Balanced: %v\n", isBalanced(root));
}

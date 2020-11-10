/*
 *  Level order traversal:
 *
 *              8
 *            /  \
 *          7     9
 *         / \   / 
 *        1   3 5
 */

package main 

import "fmt"

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

func levelOrderTraversal(root *node) {
  var nodes []*node
  var currentNode *node
  
  if root == nil {
    return
  }
  
  fmt.Printf("%v ", root.val)
  if root.left != nil {
    nodes = append(nodes, root.left)
  }
  
  if root.right != nil {
    nodes = append(nodes, root.right)
  }

  for len(nodes) > 0 {
    currentNode, nodes = nodes[0], nodes[1:]
    fmt.Printf("%v ", currentNode.val)
    if currentNode.left != nil {
      nodes = append(nodes, currentNode.left)
    }
    if currentNode.right != nil {
      nodes = append(nodes, currentNode.right)
    } 
  }
  fmt.Println("")
}

func main() {
  root := new(node)
  root.val = 8
  
  node7 := createNode(root, 7, true)
  node9 := createNode(root, 9, false)
  createNode(node7, 1, true)
  createNode(node7, 3, false)
  createNode(node9, 5, true)
  
  levelOrderTraversal(root)
}

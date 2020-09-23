/*
  Binary Tree:

        3
      /   \
     4     5
    / \   /
   9   2 7
  /       \
 1         6
          /
         8
*/

package main

import "fmt"

// Node defines the basic unit of the binary tree
type Node struct{
  left *Node
  right *Node
  value int
}

func insert(root *Node, val int, left bool) *Node {
  node := &Node {nil, nil, val}
  if (left) {
    root.left = node
  } else {
    root.right = node
  }
  return node
}

func depthFirstSearch(root *Node) {
  if root == nil {
    return
  }
  fmt.Printf("%v, ", root.value)
  depthFirstSearch(root.left)  
  depthFirstSearch(root.right)  
}

func breadthFirstSearch(root *Node) {
  var elements []*Node
  if root == nil {
    return
  }
  elements = append(elements, root)
  for len(elements) > 0 {
    node := elements[0]
    fmt.Printf("%v, ", node.value)
    if (node.left != nil) {
      elements = append(elements, node.left)
    } 
    if (node.right != nil) {
      elements = append(elements, node.right)
    }
    elements = elements[1:]
  }
}

func main() {
  root := &Node{nil, nil, 3}

  node4 := insert(root, 4, true)
  node5 := insert(root, 5, false)
  node9 := insert(node4, 9, true)
  insert(node4, 2, false)
  node7 := insert(node5, 7, true)
  insert(node9, 1, true)
  node6 := insert(node7, 6, false)
  insert(node6, 8, true)

  depthFirstSearch(root)
  fmt.Println("")

  breadthFirstSearch(root)
  fmt.Println("")
}

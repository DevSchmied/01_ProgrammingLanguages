package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func insertNode(rootNode *Node, node *Node) *Node {
	// Base case: Empty position found or tree is empty
	if rootNode == nil {
		return node
	}

	// Case 1: New value is GREATER than current node. Insert in RIGHT subtree
	if node.Val > rootNode.Val {
		rootNode.Right = insertNode(rootNode.Right, node)

		// Case 2: New value is SMALLER than current node. Insert in LEFT subtree
	} else if node.Val < rootNode.Val {
		rootNode.Left = insertNode(rootNode.Left, node)
	}

	// Case 3: Values are EQUAL (node.Val == rootNode.Val). Do nothing (no duplicates allowed)
	return rootNode
}

func main() {
	var root *Node

	root = insertNode(root, &Node{Val: 5})
	root = insertNode(root, &Node{Val: 3})
	root = insertNode(root, &Node{Val: 7})
	root = insertNode(root, &Node{Val: 2})
	root = insertNode(root, &Node{Val: 4})
	root = insertNode(root, &Node{Val: 6})
	root = insertNode(root, &Node{Val: 8})

	fmt.Println("Root:", root.Val)                    // 5
	fmt.Println("Left:", root.Left.Val)               // 3
	fmt.Println("Right:", root.Right.Val)             // 7
	fmt.Println("Left-Left:", root.Left.Left.Val)     // 2
	fmt.Println("Left-Right:", root.Left.Right.Val)   // 4
	fmt.Println("Right-Left:", root.Right.Left.Val)   // 6
	fmt.Println("Right-Right:", root.Right.Right.Val) // 8
}

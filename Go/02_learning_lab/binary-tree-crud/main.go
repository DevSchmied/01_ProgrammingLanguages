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

func deleteNode(rootNode *Node, node *Node) *Node {
	// Base case: Empty position found
	if rootNode == nil {
		return nil
	}

	// Case 1: Value is GREATER. Delete from RIGHT subtree
	if node.Val > rootNode.Val {
		rootNode.Right = deleteNode(rootNode.Right, node)
		return rootNode
	}

	// Case 2: Value is SMALLER. Delete from LEFT subtree
	if node.Val < rootNode.Val {
		rootNode.Left = deleteNode(rootNode.Left, node)
		return rootNode
	}

	// Case 3: Value FOUND. Delete this node
	// Subcase 1: No children
	if rootNode.Left == nil && rootNode.Right == nil {
		return nil
	}

	// Subcase 2: One child
	if rootNode.Left == nil {
		return rootNode.Right
	}
	if rootNode.Right == nil {
		return rootNode.Left
	}

	// Subcase 3: Two children - replace with inorder successor
	successor := rootNode.Right
	for successor.Left != nil {
		successor = successor.Left
	}
	rootNode.Val = successor.Val
	rootNode.Right = deleteNode(rootNode.Right, &Node{Val: successor.Val})

	return rootNode
}

func searchNode(rootNode *Node, searchedNode *Node) *Node {
	// Base case: Empty position found or value not in tree
	if rootNode == nil {
		return nil
	}

	// Case 1: Value FOUND. Return the current node
	if searchedNode.Val == rootNode.Val {
		return rootNode
	}

	// Case 2: Searched value is GREATER. Search in RIGHT subtree
	if searchedNode.Val > rootNode.Val {
		return searchNode(rootNode.Right, searchedNode)
	}

	// Case 3: Searched value is SMALLER. Search in LEFT subtree
	return searchNode(rootNode.Left, searchedNode)
}

func main() {
	var root *Node

	// ========== INSERT TESTS ==========
	fmt.Println("=== INSERT TESTS ===")

	// Insert values to create a balanced BST
	root = insertNode(root, &Node{Val: 5})
	root = insertNode(root, &Node{Val: 3})
	root = insertNode(root, &Node{Val: 7})
	root = insertNode(root, &Node{Val: 2})
	root = insertNode(root, &Node{Val: 4})
	root = insertNode(root, &Node{Val: 6})
	root = insertNode(root, &Node{Val: 8})

	fmt.Println("Initial tree after inserts:")
	fmt.Println("  Root:", root.Val)                    // 5
	fmt.Println("  Left:", root.Left.Val)               // 3
	fmt.Println("  Right:", root.Right.Val)             // 7
	fmt.Println("  Left-Left:", root.Left.Left.Val)     // 2
	fmt.Println("  Left-Right:", root.Left.Right.Val)   // 4
	fmt.Println("  Right-Left:", root.Right.Left.Val)   // 6
	fmt.Println("  Right-Right:", root.Right.Right.Val) // 8
	fmt.Println()

	// ========== SEARCH TESTS ==========
	fmt.Println("=== SEARCH TESTS ===")

	// Test 1: Search for existing value
	result := searchNode(root, &Node{Val: 5})
	fmt.Println("Search(5):", result.Val) // 5

	// Test 2: Search for non-existent value
	result = searchNode(root, &Node{Val: 10})
	fmt.Println("Search(10):", result) // nil
	fmt.Println()

	// ========== DELETE TESTS ==========
	fmt.Println("=== DELETE TESTS ===")

	// Test 1: Delete leaf node (no children)
	fmt.Println("Test 1: Delete leaf node (2)")
	root = deleteNode(root, &Node{Val: 2})
	fmt.Println("  Root:", root.Val)                  // 5
	fmt.Println("  Left:", root.Left.Val)             // 3
	fmt.Println("  Left-Left:", root.Left.Left)       // nil
	fmt.Println("  Left-Right:", root.Left.Right.Val) // 4
	fmt.Println()

	// Test 2: Delete node with one child (right child)
	fmt.Println("Test 2: Delete node with one child (7)")
	root = deleteNode(root, &Node{Val: 7})
	fmt.Println("  Root:", root.Val)                  // 5
	fmt.Println("  Right:", root.Right.Val)           // 8
	fmt.Println("  Right-Left:", root.Right.Left.Val) // 6
	fmt.Println("  Right-Right:", root.Right.Right)   // nil
	fmt.Println()

	// Test 3: Delete node with two children (using successor)
	fmt.Println("Test 3: Delete node with two children (5)")
	root = deleteNode(root, &Node{Val: 5})
	fmt.Println("  New Root (successor):", root.Val) // 6
	fmt.Println("  Left:", root.Left.Val)            // 3
	fmt.Println("  Right:", root.Right.Val)          // 8
	fmt.Println("  Right-Left:", root.Right.Left)    // nil
	fmt.Println()

	// Test 4: Delete remaining nodes
	fmt.Println("Test 4: Delete remaining nodes")
	root = deleteNode(root, &Node{Val: 3})
	fmt.Println("  After deleting 3:", root.Val) // 6
	fmt.Println("  Left:", root.Left)            // nil

	root = deleteNode(root, &Node{Val: 6})
	fmt.Println("  After deleting 6:", root.Val) // 8

	root = deleteNode(root, &Node{Val: 8})
	fmt.Println("  After deleting 8:", root) // nil
	fmt.Println()

	// Test 5: Delete from empty tree
	fmt.Println("Test 5: Delete from empty tree")
	root = deleteNode(root, &Node{Val: 10})
	fmt.Println("  Result:", root) // nil
}

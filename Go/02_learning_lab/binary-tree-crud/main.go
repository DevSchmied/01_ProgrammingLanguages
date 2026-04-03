package main

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

}

package main

import "fmt"

/*
LeetCode 105: Construct Binary Tree from Preorder and Inorder Traversal
Quelle: LeetCode (https://leetcode.com/)

Problem Statement:

Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

Preorder traversal visits nodes in the order: root, left subtree, right subtree.
Inorder traversal visits nodes in the order: left subtree, root, right subtree.

Examples:
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

Input: preorder = [-1], inorder = [-1]
Output: [-1]

Constraints:
- 1 <= preorder.length <= 3000
- inorder.length == preorder.length
- -3000 <= preorder[i], inorder[i] <= 3000
- preorder and inorder consist of unique values.
- Each value in inorder also appears in preorder.
- preorder is guaranteed to be the preorder traversal of the tree.
- inorder is guaranteed to be the inorder traversal of the tree.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	rootIdx := 0
	for idx, v := range inorder {
		if v == preorder[0] {
			rootIdx = idx
		}
	}

	leftInorder := inorder[:rootIdx]
	rightInorder := inorder[rootIdx+1:]

	leftPreorder := preorder[1 : 1+len(leftInorder)]
	rightPreorder := preorder[1+len(leftInorder):]

	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)

	return root
}

func main() {
	// Test case 1
	preorder1 := []int{3, 9, 20, 15, 7}
	inorder1 := []int{9, 3, 15, 20, 7}
	tree1 := buildTree(preorder1, inorder1)

	fmt.Println("Test 1:")
	fmt.Printf("  Root value: %d (expected: 3)\n", tree1.Val)
	fmt.Printf("  Left child: %d (expected: 9)\n", tree1.Left.Val)
	fmt.Printf("  Right child: %d (expected: 20)\n", tree1.Right.Val)
	fmt.Printf("  Right->Left: %d (expected: 15)\n", tree1.Right.Left.Val)
	fmt.Printf("  Right->Right: %d (expected: 7)\n", tree1.Right.Right.Val)

	// Test case 2
	preorder2 := []int{-1}
	inorder2 := []int{-1}
	tree2 := buildTree(preorder2, inorder2)
	fmt.Printf("\nTest 2:\n")
	fmt.Printf("  Root value: %d (expected: -1)\n", tree2.Val)
}

package main

/*
LeetCode 106: Construct Binary Tree from Inorder and Postorder Traversal
Quelle: LeetCode (https://leetcode.com/)

Problem Statement:

Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree.

Examples:
Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
Output: [3,9,20,null,null,15,7]
Explanation: The binary tree represented by these traversals has root 3, left child 9, right child 20, and 20 has left child 15 and right child 7.

Input: inorder = [-1], postorder = [-1]
Output: [-1]
Explanation: A single node tree with value -1.

Constraints:
- 1 <= inorder.length <= 3000
- postorder.length == inorder.length
- -3000 <= inorder[i], postorder[i] <= 3000
- inorder and postorder consist of unique values.
- Each value of postorder also appears in inorder.
- inorder is guaranteed to be the inorder traversal of the tree.
- postorder is guaranteed to be the postorder traversal of the tree.
*/

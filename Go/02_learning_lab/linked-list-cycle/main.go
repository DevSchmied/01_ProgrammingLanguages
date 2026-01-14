package main

/*
LeetCode 141 — Linked List Cycle

Problem Description:
Given the head of a singly linked list, determine if the linked list contains a cycle.

A cycle exists if there is some node in the list that can be reached again by
continuously following the `next` pointer.

Internally, `pos` is used to indicate the index of the node that the tail's `next`
pointer is connected to. Note that `pos` is NOT passed as a parameter.

Return true if the linked list has a cycle. Otherwise, return false.

Examples:

Example 1:
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: The tail connects to the node at index 1 (0-indexed), forming a cycle.

Example 2:
Input: head = [1,2], pos = 0
Output: true
Explanation: The tail connects to the node at index 0, forming a cycle.

Example 3:
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.

Constraints:
- The number of nodes in the list is in the range [0, 10^4].
- -10^5 <= Node.val <= 10^5
- pos is -1 or a valid index in the linked list.
*/

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

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	visited := make(map[*ListNode]struct{})
	cur := head

	for cur != nil {
		if _, ok := visited[cur]; ok {
			return true
		}
		visited[cur] = struct{}{}
		cur = cur.Next
	}

	return false
}

func main() {
	fmt.Println("=================1. Example=================")

	var ln1, ln2, ln3, ln4 ListNode

	ln4 = ListNode{
		Val:  -4,
		Next: &ln2,
	}

	ln3 = ListNode{
		Val:  0,
		Next: &ln4,
	}

	ln2 = ListNode{
		Val:  2,
		Next: &ln3,
	}

	ln1 = ListNode{
		Val:  3,
		Next: &ln2,
	}

	fmt.Println(hasCycle(&ln1))

	fmt.Println("=================2. Example=================")

	var ln21, ln22 ListNode

	ln21.Val = 1
	ln21.Next = &ln22

	ln22.Val = 2
	ln22.Next = &ln21

	fmt.Println(hasCycle(&ln21))

	fmt.Println("=================3. Example=================")

	ln31 := ListNode{
		Val:  1,
		Next: nil,
	}

	fmt.Println(hasCycle(&ln31))

}

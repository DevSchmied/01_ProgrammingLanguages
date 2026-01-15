package main

import "fmt"

/*
LeetCode 160 — Intersection of Two Linked Lists

Problem Description:
Given the heads of two singly linked lists headA and headB, return the node at which
the two linked lists intersect. If the two linked lists have no intersection, return nil.

An intersection is defined by reference, not by value. In other words, two nodes are
considered intersecting only if they point to the same memory location.

The linked lists must retain their original structure after the function returns.

Important Notes:
- There are no cycles in the entire linked structure.
- The judge internally constructs the lists using the following parameters:
  - intersectVal: the value of the intersecting node (0 if there is no intersection)
  - listA: the first linked list
  - listB: the second linked list
  - skipA: number of nodes before the intersecting node in listA
  - skipB: number of nodes before the intersecting node in listB
- These parameters are NOT passed to the function directly.

Example 1:
Input:
intersectVal = 8
listA = [4,1,8,4,5]
listB = [5,6,1,8,4,5]
skipA = 2
skipB = 3

Output:
Intersected at '8'

Explanation:
The node with value 8 is the first common node by reference.
Nodes with the same value (e.g., 1) are NOT considered intersecting
unless they are the same object in memory.

Example 2:
Input:
intersectVal = 2
listA = [1,9,1,2,4]
listB = [3,2,4]
skipA = 3
skipB = 1

Output:
Intersected at '2'

Example 3:
Input:
intersectVal = 0
listA = [2,6,4]
listB = [1,5]

Output:
No intersection (return nil)

Constraints:
- Number of nodes in listA: m
- Number of nodes in listB: n
- 1 <= m, n <= 3 * 10^4
- 1 <= Node.val <= 10^5
- 0 <= skipA <= m
- 0 <= skipB <= n
- intersectVal == 0 if and only if the lists do not intersect
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	curA := headA
	mapA := make(map[*ListNode]struct{})

	for curA != nil {
		mapA[curA] = struct{}{}
		curA = curA.Next
	}

	curB := headB

	for curB != nil {
		if _, ok := mapA[curB]; ok {
			return curB
		}
		curB = curB.Next
	}

	return nil
}

func getIntersectionNodeOptimized(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB

	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return a
}

func main() {

	fmt.Println("=========================\nExample 1\nIntersection at value 8\n=========================")

	var lnb1, lnb2, lnb3, lna2, lna3, lnab4, lnab5, lnab6 ListNode

	// List B: 5 -> 6 -> 1 -> 8 -> 4 -> 5
	lnb1.Val = 5
	lnb1.Next = &lnb2

	lnb2.Val = 6
	lnb2.Next = &lnb3

	lnb3.Val = 1
	lnb3.Next = &lnab4

	// List A: 4 -> 1 -> 8 -> 4 -> 5
	lna2.Val = 4
	lna2.Next = &lna3

	lna3.Val = 1
	lna3.Next = &lnab4

	// Shared part
	lnab4.Val = 8
	lnab4.Next = &lnab5

	lnab5.Val = 4
	lnab5.Next = &lnab6

	lnab6.Val = 5
	lnab6.Next = nil

	res1 := getIntersectionNode(&lna2, &lnb1)
	if res1 != nil {
		fmt.Println("Intersection found at node with value:", res1.Val)
	} else {
		fmt.Println("No intersection found")
	}

	fmt.Println()

	fmt.Println("=========================\nExample 2\nIntersection at value 2\n=========================")

	// Shared part
	var shared1, shared2 ListNode
	shared1.Val = 2
	shared1.Next = &shared2

	shared2.Val = 4
	shared2.Next = nil

	// List A: 1 -> 9 -> 1 -> 2 -> 4
	var a1, a2, a3 ListNode
	a1.Val = 1
	a1.Next = &a2

	a2.Val = 9
	a2.Next = &a3

	a3.Val = 1
	a3.Next = &shared1

	// List B: 3 -> 2 -> 4
	var b1 ListNode
	b1.Val = 3
	b1.Next = &shared1

	res2 := getIntersectionNode(&a1, &b1)
	if res2 != nil {
		fmt.Println("Intersection found at node with value:", res2.Val)
	} else {
		fmt.Println("No intersection found")
	}

	fmt.Println()

	fmt.Println("=========================\nExample 3\nNo intersection\n=========================")

	// List A: 2 -> 6 -> 4
	var a4, a5, a6 ListNode
	a4.Val = 2
	a4.Next = &a5

	a5.Val = 6
	a5.Next = &a6

	a6.Val = 4
	a6.Next = nil

	// List B: 1 -> 5
	var b2, b3 ListNode
	b2.Val = 1
	b2.Next = &b3

	b3.Val = 5
	b3.Next = nil

	res3 := getIntersectionNode(&a4, &b2)
	if res3 != nil {
		fmt.Println("Intersection found at node with value:", res3.Val)
	} else {
		fmt.Println("No intersection found")
	}

	fmt.Println("=========================\nExample 1 (optimized)\nIntersection at value 8\n=========================")

	res1opt := getIntersectionNodeOptimized(&lna2, &lnb1)
	if res1opt != nil {
		fmt.Println("Intersection found at node with value:", res1opt.Val)
	} else {
		fmt.Println("No intersection found")
	}

	fmt.Println()

	fmt.Println("=========================\nExample 2 (optimized)\nIntersection at value 2\n=========================")

	res2opt := getIntersectionNodeOptimized(&a1, &b1)
	if res2opt != nil {
		fmt.Println("Intersection found at node with value:", res2opt.Val)
	} else {
		fmt.Println("No intersection found")
	}

	fmt.Println()

	fmt.Println("=========================\nExample 3 (optimized)\nNo intersection\n=========================")

	res3opt := getIntersectionNodeOptimized(&a4, &b2)
	if res3opt != nil {
		fmt.Println("Intersection found at node with value:", res3opt.Val)
	} else {
		fmt.Println("No intersection found")
	}
}

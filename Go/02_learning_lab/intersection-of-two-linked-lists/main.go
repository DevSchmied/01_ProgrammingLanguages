package main

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

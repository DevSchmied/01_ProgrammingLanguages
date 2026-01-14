package main

import "testing"

func TestHasCycle(t *testing.T) {
	t.Run("nil list", func(t *testing.T) {
		if hasCycle(nil) {
			t.Fatal("expected false for nil list")
		}
	})

	t.Run("single node without cycle", func(t *testing.T) {
		n1 := &ListNode{Val: 1}
		if hasCycle(n1) {
			t.Fatal("expected false for single node without cycle")
		}
	})

	t.Run("single node with self cycle", func(t *testing.T) {
		n1 := &ListNode{Val: 1}
		n1.Next = n1

		if !hasCycle(n1) {
			t.Fatal("expected true for self cycle")
		}
	})

	t.Run("two nodes without cycle", func(t *testing.T) {
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n1.Next = n2

		if hasCycle(n1) {
			t.Fatal("expected false for two nodes without cycle")
		}
	})

	t.Run("two nodes with cycle to head", func(t *testing.T) {
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n1.Next = n2
		n2.Next = n1

		if !hasCycle(n1) {
			t.Fatal("expected true for cycle to head")
		}
	})

	t.Run("cycle in the middle", func(t *testing.T) {
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n3 := &ListNode{Val: 3}
		n4 := &ListNode{Val: 4}
		n5 := &ListNode{Val: 5}

		n1.Next = n2
		n2.Next = n3
		n3.Next = n4
		n4.Next = n5
		n5.Next = n3

		if !hasCycle(n1) {
			t.Fatal("expected true for cycle in the middle")
		}
	})
}

func TestHasCycleOptimized(t *testing.T) {
	t.Run("nil list", func(t *testing.T) {
		if hasCycleOptimized(nil) {
			t.Fatal("expected false for nil list")
		}
	})

	t.Run("single node without cycle", func(t *testing.T) {
		n1 := &ListNode{Val: 1}
		if hasCycleOptimized(n1) {
			t.Fatal("expected false for single node without cycle")
		}
	})

	t.Run("single node with self cycle", func(t *testing.T) {
		n1 := &ListNode{Val: 1}
		n1.Next = n1

		if !hasCycleOptimized(n1) {
			t.Fatal("expected true for self cycle")
		}
	})

	t.Run("two nodes without cycle", func(t *testing.T) {
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n1.Next = n2

		if hasCycleOptimized(n1) {
			t.Fatal("expected false for two nodes without cycle")
		}
	})

	t.Run("two nodes with cycle to head", func(t *testing.T) {
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n1.Next = n2
		n2.Next = n1

		if !hasCycleOptimized(n1) {
			t.Fatal("expected true for cycle to head")
		}
	})

	t.Run("cycle in the middle", func(t *testing.T) {
		n1 := &ListNode{Val: 1}
		n2 := &ListNode{Val: 2}
		n3 := &ListNode{Val: 3}
		n4 := &ListNode{Val: 4}
		n5 := &ListNode{Val: 5}

		n1.Next = n2
		n2.Next = n3
		n3.Next = n4
		n4.Next = n5
		n5.Next = n3

		if !hasCycleOptimized(n1) {
			t.Fatal("expected true for cycle in the middle")
		}
	})
}

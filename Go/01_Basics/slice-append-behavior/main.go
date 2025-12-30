package main

/*
Task:

Implement two functions that attempt to add a new element to a sequence of values.

One function should receive the sequence directly, while the other should receive a reference to the sequence.

Observe and explain why modifying the sequence succeeds in one case but not in the other.
*/

import "fmt"

func appendN(n []int) {
	n = append(n, 4)
}
func appendNs(n *[]int) {
	*n = append(*n, 4)
}

func main() {
	n := []int{1, 2, 3}

	appendN(n)
	fmt.Println(n)

	appendNs(&n)
	fmt.Println(n)
}

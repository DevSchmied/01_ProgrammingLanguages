package main

import "fmt"

/*
Task: Copying data between dynamic collections

Write a program that works with a dynamic collection of integers and demonstrates how data can be transferred from one collection to another.

The program should:

- Create an initial collection with several integer values.
- Prepare another empty collection intended to receive the data.
- Attempt to transfer the data and record how many elements were actually transferred.
- Show the difference between a collection that only reserves memory and one that truly contains elements.
- Ensure that a final collection is created as a fully independent copy of the original data.
- Output the result of the transfer operation.

The goal is to understand how internal structure and initialization of dynamic collections affect data transfer and memory independence.
*/

func main() {
	a := []int{5, 6, 7, 8}

	b := make([]int, len(a))

	count := copy(b, a)
	fmt.Println("count:", count)
	fmt.Println("b", b)

	c := append([]int(nil), a...)
	fmt.Println("c", c)
}

package main

/*
Task:
Create a function that returns another function which maintains internal state across multiple calls.
Each invocation of the returned function should update and return the stored state.
*/

import "fmt"

func counter() func() int {
	n := 0

	return func() int {
		n++
		return n
	}
}

func main() {
	c := counter()

	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}

package main

import "fmt"

func main() {

	/*
		Task 1 — Integer Pointer
			- Declare an integer variable.
			- Declare a pointer variable to an integer.
			- Assign a value to the integer variable.
			- Assign the address of the integer variable to the pointer.
			- Print the value of the integer variable.
			- Print the address of the integer variable.
			- Print the value of the pointer variable.
			- Print the value stored at the address the pointer points to.
	*/

	var num int
	var numPtr *int

	num = 42
	numPtr = &num

	fmt.Printf("Value of num: %d\n", num)
	fmt.Printf("Address of num: %p\n", &num)
	fmt.Printf("Value of numPtr (address stored): %p\n", numPtr)
	fmt.Printf("Value pointed to by numPtr: %d\n", *numPtr)

}

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
	fmt.Println("---------------------Task 1---------------------")

	var num int
	var numPtr *int

	num = 42
	numPtr = &num

	fmt.Printf("Value of num: %d\n", num)
	fmt.Printf("Address of num: %p\n", &num)
	fmt.Printf("Value of numPtr (address stored): %p\n", numPtr)
	fmt.Printf("Value pointed to by numPtr: %d\n", *numPtr)

	/*
		Task 2 — String Pointer
			- Declare a string variable.
			- Declare a pointer variable to a string.
			- Assign a text value to the string variable.
			- Assign the address of the string variable to the pointer.
			- Print the value of the string variable.
			- Print the address of the string variable.
			- Print the value of the pointer variable.
			- Print the value stored at the address the pointer points to.
	*/

	fmt.Println()
	fmt.Println("---------------------Task 2---------------------")

	var str string
	var strPtr *string

	str = "Text"
	strPtr = &str

	fmt.Printf("Value of str: %s\n", str)
	fmt.Printf("Address of str: %p\n", &str)

	fmt.Printf("Value of strPtr (stored address): %p\n", strPtr)
	fmt.Printf("Value pointed to by strPtr: %s\n", *strPtr)
}

package main

import (
	"fmt"
	"os"
)

/*

Source: Head First Go — Chapter on Variadic Functions and Slice Expansion

Assignment:
Write a Go program that calculates the average of a series of numbers provided as command-line arguments.

Your program must:
1. Read all command-line arguments (excluding the program name).
2. Convert each argument from string to float64.
3. Store all converted numbers in a []float64 slice.
4. Implement a function average(numbers ...float64) float64 that accepts a variable number of arguments and returns their average.
5. Call this function using the slice you created.
Since average expects a variadic list of float64, you must use the slice-expansion syntax (slice...).
6. Print the result with two decimal places.
7. The implementation must be split into several logically structured steps so that they can be committed separately.
*/

func main() {
	args := os.Args[1:]
	fmt.Println("Arguments:", args)
}

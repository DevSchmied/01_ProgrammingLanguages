package main

/*
Task:
Implement a program that constructs a string by repeatedly appending characters in a loop.

The solution should use an efficient approach to string accumulation, minimizing unnecessary memory allocations.
*/

import (
	"fmt"
	"strings"
)

func main() {
	sb := strings.Builder{}

	for i := 0; i < 100; i++ {
		_, err := sb.WriteString("A")
		if err != nil {
			fmt.Printf("Stringbuilder error: %v\n", err)
		}
	}
	str := sb.String()
	fmt.Println(str)
}

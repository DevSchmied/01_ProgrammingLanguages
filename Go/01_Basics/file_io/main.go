package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Task:
Write a Go program that asks the user to enter some text and then saves this text into a file named output.txt.
After saving, the program should read the same file and print the content back to the user.

*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter some text: ")
	text, _ := reader.ReadString('\n')

	fmt.Println("You entered:", text)
}

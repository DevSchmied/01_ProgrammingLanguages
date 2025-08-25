package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	// String exercises in Go
	fmt.Println("--------------------String exercises in Go--------------------\n")

	// Task 1: Create a string and find its length
	originalText := "   Hello, World!   "
	length := len(originalText)
	fmt.Printf("Length of originalText: %d\n", length)

	// Task 2: Check if the string contains "World"
	containsWorld := strings.Contains(originalText, "World")
	fmt.Printf("Does originalText contain \"World\"? %t\n", containsWorld)

	// Task 3: Convert the string to uppercase
	upperText := strings.ToUpper(originalText)
	fmt.Printf("Uppercase: %s\n", upperText)

	// Task 4: Remove leading and trailing spaces
	trimmedText := strings.TrimSpace(originalText)
	fmt.Printf("Trimmed: %s\n", trimmedText)

	// Task 5: Replace "World" with "Go"
	replaced := strings.Replace(originalText, "World", "Go", 1)
	fmt.Printf("Replaced: %s\n", replaced)

	// Task 6: Split the original string by comma
	splitValues := strings.Split(originalText, ",")
	fmt.Println("Split by comma:")
	for _, v := range splitValues {
		fmt.Println(v)
	}

	// Task 7: Join ["Hello", "Go"] into a single string (3 ways)
	words := []string{"Hello", "Go"}

	// 1. Concatenation
	joinedConcat := ""
	for _, w := range words {
		joinedConcat += w + " "
	}
	joinedConcat = strings.TrimSpace(joinedConcat)

	// 2. Standard method
	joinedStringsJoin := strings.Join(words, " ")

	// 3. Using strings.Builder
	var sb strings.Builder
	for i, w := range words {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(w)
	}
	joinedBuilder := sb.String()

	fmt.Println("Joined (concatenation):", joinedConcat)
	fmt.Println("Joined (strings.Join):", joinedStringsJoin)
	fmt.Println("Joined (strings.Builder):", joinedBuilder)

	// Task 8: Check if trimmed string starts with "Hello"
	startsWithHello := strings.HasPrefix(trimmedText, "Hello")
	fmt.Printf("Does trimmedText start with \"Hello\"? %t\n", startsWithHello)

	// Task 9: Repeat the string "Go" 3 times
	repeated := strings.Repeat("Go", 3)
	fmt.Println("Repeated:", repeated)

	// Task 10: Convert trimmed string into a slice of runes
	runes := []rune(trimmedText)
	fmt.Println("Runes:", runes)

	// Task 11: Convert an error to a string
	err := errors.New("Test Error")
	errString := err.Error()
	fmt.Println("Error as string:", errString)

	repeat()
}

func repeat() {

	fmt.Println("\n\n--------------------------Function repeat()--------------------------")

	// String exercises in Go

	// Task 1: Create a string originalText = "   Hello, World!   " and find its length.

	originalText := "   Hello, World!   "
	fmt.Println("Length of originalText:", len(originalText))

	// Task 2: Check if the string originalText contains the substring "World".

	isContain := strings.Contains(originalText, "World")
	fmt.Println("Does originalText contain the substring \"World\"?", isContain)

	// Task 3: Convert the string originalText to uppercase.

	originalTextUp := strings.ToUpper(originalText)
	fmt.Println("originalText in uppercase:", originalTextUp)

	// Task 4: Remove leading and trailing spaces from the string originalText.

	trimmedText := strings.TrimSpace(originalText)
	fmt.Println("originalText without leading and trailing spaces:", trimmedText)

	// Task 5: Replace "World" with "Go" in the string originalText.

	replacedText := strings.Replace(originalText, "World", "Go", 1)
	fmt.Println("Replaced \"World\" with \"Go\":", replacedText)

	// Task 6: Split the trimmed string (from Task 4) by comma.

	wordsSplit := strings.Split(trimmedText, ",")
	fmt.Println("Split trimmedText into words:")
	for i := 0; i < len(wordsSplit); i++ {
		fmt.Printf("%d. word: %s\n", i+1, wordsSplit[i])
	}

	// Task 7: Join the string slice words = ["Hello", "Go"] into a single string separated by a space.
	// (3 ways: Concatenation, Standard Method, StringBuilder)
	words := []string{"Hello", "Go"}

	// 1. Concatenation
	singleString := ""
	for i := 0; i < len(words); i++ {
		if i != 0 {
			singleString += " "
		}
		singleString += words[i]
	}
	fmt.Println("Joined words (Concatenation):", singleString)

	// 2. Using strings.Join
	singleString = strings.Join(words, " ")
	fmt.Println("Joined words (strings.Join):", singleString)

	// 3. Using StringBuilder
	var sb strings.Builder
	for i, w := range words {
		if i != 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(w)
	}
	singleString = sb.String()
	fmt.Println("Joined words (StringBuilder):", singleString)

	// Task 8: Check if the trimmed string (from Task 4) starts with "Hello".

	isStart := strings.HasPrefix(trimmedText, "Hello")
	fmt.Println("Does trimmedText start with \"Hello\"?", isStart)

	// Task 9: Repeat the string "Go" 3 times.

	repeatedGo := strings.Repeat("Go", 3)
	fmt.Println("The string \"Go\" repeated 3 times:", repeatedGo)

	// Task 10: Convert the trimmed string (from Task 4) into a slice of runes.

	runes := []rune(trimmedText)
	fmt.Println("trimmedText as a slice of runes:", runes)

	// Task 11: Convert an error (error type) into a string. !!!!!!!!!!!!!!

	err := errors.New("Test error")
	errString := err.Error()
	fmt.Println("User-defined error:", errString)
}

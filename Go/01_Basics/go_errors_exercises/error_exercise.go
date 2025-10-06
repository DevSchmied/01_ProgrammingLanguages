package main

import (
	"errors"
	"fmt"
	"sync"
)

/*
Goal

In this exercise, you will practice how to create, wrap, check, and unwrap errors in Go.
You will also learn how to define a custom error type and extract it using errors.As().

Step 1 — Create and return basic errors
Write a function:
- func readFile(name string) (string, error)
 * If name is an empty string (""), return an error using errors.New("filename is missing").
 * If the file name is "config.txt", return the string "file loaded successfully" and nil.
 * Otherwise, return a new error using fmt.Errorf("file %s not found", name).

Step 2 — Wrap an error using fmt.Errorf()
Create a new function:
    func loadConfig(name string) error

Inside this function:
- Call readFile(name).
- If an error is returned, wrap it with a new message using:
      fmt.Errorf("failed to load configuration: %w", err)
- Return the wrapped error.
- If no error occurs, return nil.

*/

// Step 1
// readFile simulates reading a file and returns different errors or success messages.
func readFile(name string) (string, error) {
	if name == "" {
		return "", errors.New("filename is missing")
	} else if name == "config.txt" {
		return "file loaded successfully", nil
	} else {
		return "", fmt.Errorf("file %s not found", name)
	}
}

// Step 2
// loadConfig calls readFile() and wraps any returned error with additional context.
func loadConfig(name string) error {
	if _, err := readFile(name); err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}
	return nil
}

func main() {
	var wg sync.WaitGroup

	// Step 1
	// Empty filename -> expect error.
	wg.Add(1)
	go func() {
		defer wg.Done()
		str, err := readFile("")
		fmt.Printf("string: \"%s\", error: %v\n", str, err)
	}()

	// Valid file -> expect success.
	wg.Add(1)
	go func() {
		defer wg.Done()
		str, err := readFile("config.txt")
		fmt.Printf("string: \"%s\", error: %v\n", str, err)
	}()

	// Unknown file -> expect error.
	wg.Add(1)
	go func() {
		defer wg.Done()
		str, err := readFile("foo.txt")
		fmt.Printf("string: \"%s\", error: %v\n", str, err)
	}()

	wg.Wait()

	// Step 2
	// Test 1: Empty filename → expect wrapped error
	err := loadConfig("")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Test 2: Valid file → expect success
	err = loadConfig("config.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Successfully loaded")
	}

	// Test 3: Invalid file → expect wrapped error
	err = loadConfig("name.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

}

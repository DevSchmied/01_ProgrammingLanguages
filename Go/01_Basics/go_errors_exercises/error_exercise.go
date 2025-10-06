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
*/

func readFile(name string) (string, error) {
	if name == "" {
		return "", errors.New("filename is missing")
	} else if name == "config.txt" {
		return "file loaded successfully", nil
	} else {
		return "", fmt.Errorf("file %s not found", name)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		str, err := readFile("")
		fmt.Printf("string: \"%s\", error: %v\n", str, err)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		str, err := readFile("config.txt")
		fmt.Printf("string: \"%s\", error: %v\n", str, err)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		str, err := readFile("foo.txt")
		fmt.Printf("string: \"%s\", error: %v\n", str, err)
	}()

	wg.Wait()
}

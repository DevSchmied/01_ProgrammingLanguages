package main

/*
Task:
Implement a program in which multiple concurrently executing operations modify a shared numeric value.

Demonstrate how the absence of synchronization leads to incorrect results, and how the use of a mutual exclusion mechanism ensures correct updates of the shared state.
*/

import (
	"fmt"
	"sync"
)

var counter = 0

func main() {

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	fmt.Println("Counter WITHOUT mutex after 1000 iterations:", counter)

	counter = 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			counter++
		}()
	}

	wg.Wait()
	fmt.Println("Counter WITH mutex after 1000 iterations:", counter)
}

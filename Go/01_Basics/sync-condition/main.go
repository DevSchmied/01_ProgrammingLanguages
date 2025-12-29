package main

/*
Task:
Implement a program where one task waits for a specific condition to become true before continuing its execution.
Another task should update the shared state and notify the waiting task when the condition is satisfied.
*/

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)

	ready := false

	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		for !ready {
			fmt.Println("Waiting...")
			cond.Wait()
		}
		fmt.Println("Condition met")
		mu.Unlock()
	}()

	time.Sleep(2 * time.Second)

	mu.Lock()
	ready = true
	cond.Signal()
	mu.Unlock()

	wg.Wait()
}

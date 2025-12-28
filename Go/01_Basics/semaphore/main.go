package main

/*
Task:
Implement a program that limits the number of concurrently running tasks.
Multiple tasks should be started in parallel, but only a fixed number of them may execute at the same time.
*/

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	sem := make(chan struct{}, 3)
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(job int) {
			defer wg.Done()
			sem <- struct{}{}
			fmt.Printf("%d. Job. Time after start: %v\n", job+1, time.Since(start))
			time.Sleep(1 * time.Second)
			<-sem
		}(i)
	}

	wg.Wait()
}

package main

/*
Task:
Design a program that processes a series of tasks using a fixed number of workers.
Tasks should be distributed among workers and processed concurrently, while ensuring that all tasks are completed before the program exits.
*/
import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("%d. worker processing %d. job\n", id, job)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	const workers = 3
	const jobsNum = 10

	var wg sync.WaitGroup
	jobs := make(chan int)

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for j := 1; j <= jobsNum; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
}

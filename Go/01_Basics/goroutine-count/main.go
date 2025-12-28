package main

/*
Task:
Create a program that starts multiple concurrent tasks, keeps them running for a short period of time, and reports how many concurrent executions are active at a given moment.
*/
import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(2 * time.Second)
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())

}

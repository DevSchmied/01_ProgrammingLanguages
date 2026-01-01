package main

/*
Task description:

Design a program that processes data coming from two independent asynchronous sources.

Each source produces a sequence of values over time and signals when it has finished producing data. The program must:
- Receive and handle values from both sources as they become available, without blocking on one source while waiting for the other.
- Correctly detect when a source has finished and stop listening to it.
- Continue processing until all sources are exhausted.
- Output each received value along with information about which source it came from.

The solution should demonstrate safe coordination of concurrent data streams, proper termination handling, and avoidance of unnecessary waiting or resource leaks.
*/

import (
	"fmt"
	"time"
)

func main() {

	chInt := make(chan int, 3)
	chStr := make(chan string, 3)

	go func() {
		for i := 0; i < 5; i++ {
			chInt <- i + 1
			time.Sleep(500 * time.Millisecond)
		}
		close(chInt)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			chStr <- string(65 + i)
			time.Sleep(500 * time.Millisecond)
		}
		close(chStr)
	}()

	for chInt != nil || chStr != nil {
		select {
		case valueInt, ok := <-chInt:
			if !ok {
				chInt = nil
				continue
			}
			fmt.Println("Int value:", valueInt)
		case valueStr, ok := <-chStr:
			if !ok {
				chStr = nil
				continue
			}
			fmt.Println("String value:", valueStr)
		}
	}
}

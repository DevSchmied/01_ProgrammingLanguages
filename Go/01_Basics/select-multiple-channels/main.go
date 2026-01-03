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
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	chInt1 := make(chan int, 3)
	chStr1 := make(chan string, 3)

	chInt2 := make(chan int, 3)
	chStr2 := make(chan string, 3)

	go func() {
		for i := 0; i < 5; i++ {
			chInt1 <- i + 1
			time.Sleep(500 * time.Millisecond)
		}
		close(chInt1)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			chStr1 <- string(65 + i)
			time.Sleep(500 * time.Millisecond)
		}
		close(chStr1)
	}()

	for chInt1 != nil || chStr1 != nil {
		select {
		case valueInt, ok := <-chInt1:
			if !ok {
				chInt1 = nil
				continue
			}
			fmt.Println("Int value:", valueInt)
		case valueStr, ok := <-chStr1:
			if !ok {
				chStr1 = nil
				continue
			}
			fmt.Println("String value:", valueStr)
		}
	}

	/*
		Task:

			Design a program that concurrently consumes data from multiple independent input streams.

			Each stream produces values asynchronously and may finish at an unknown time.
			The program should process incoming values as soon as they are available, detect when a stream is exhausted, and terminate gracefully when all streams are completed or when a predefined timeout is reached.
	*/

	go func() {
		for i := 0; i < 5; i++ {
			chInt2 <- i + 1
			time.Sleep(500 * time.Millisecond)
		}
		close(chInt2)
	}()

	go func() {
		for i := 0; i < 5; i++ {
			chStr2 <- string(65 + i)
			time.Sleep(500 * time.Millisecond)
		}
		close(chStr2)
	}()

	fmt.Println("=================repeat()=================")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go repeat(chInt2, chStr2, &wg)
	wg.Wait()
}

func repeat(ch1 <-chan int, ch2 <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for ch1 != nil || ch2 != nil {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Println("repeat: info from ch1", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println("repeat: info from ch2", v)
		case <-ctx.Done():
			ch1, ch2 = nil, nil
			return
		}
	}
}

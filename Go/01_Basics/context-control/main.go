package main

/*
Task:
Implement a program that performs continuous work in the background and can be stopped gracefully when a termination signal is received.
The background process should periodically perform its work and exit cleanly when instructed to stop.
*/

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine is ended")
				return
			case <-time.After(500 * time.Millisecond):
				fmt.Println("The work is processing...")
			}
		}
	}(ctx)

	time.Sleep(1500 * time.Millisecond)
	cancel()
	time.Sleep(500 * time.Millisecond)
}

package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Recieved cancellation signal, task is ended")
			return
		default:
			fmt.Println("Executing task...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// create an empty Context
	ctx := context.Background()

	// create a Context with cancelling signal
	ctxWithCancel, cancel := context.WithCancel(ctx)
	defer cancel() // defer cancel, make sure Context is cancelled when existing the function

	go worker(ctxWithCancel)

	time.Sleep(3 * time.Second)
	cancel() // send cancellation signal
	time.Sleep(1 * time.Second)
}

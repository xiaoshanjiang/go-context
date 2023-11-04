package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// create an empty Context
	ctx := context.Background()

	// create a Context with cancelling signal
	ctxWithCancel, cancel := context.WithCancel(ctx)
	defer cancel() // defer cancel, make sure Context is cancelled when existing the function

	go func() {
		for {
			select {
			case <-ctxWithCancel.Done():
				fmt.Println("Recieved cancellation signal, task is ended")
				return
			// simulate a long-running task and prints out "Task done"
			// if the task is done within 2 seconds
			case <-time.After(2 * time.Second):
				fmt.Println("Task done")
			}
		}
	}()

	time.Sleep(1 * time.Second)
	cancel() // send cancellation signal
	time.Sleep(1 * time.Second)
}

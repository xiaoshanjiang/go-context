package main

import (
	"context"
	"fmt"
)

func main() {
	// create an empty Context
	ctx := context.Background()

	// create a Context with cancelling signal
	_, cancel := context.WithCancel(ctx)
	defer cancel() // defer cancel, make sure Context is cancelled when existing the function

	fmt.Println("Context created successfully.")
}

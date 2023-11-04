package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// Set a 2 seconds timeout for every request
	// if the request is handled within 2 seconds, return "Hello world!"
	// else send the timeout signal and return a timeout error
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Fprintln(w, "Hello, World!")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("Handling request: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

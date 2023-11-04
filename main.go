package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// use Context from http.Request struct
	ctx := r.Context()

	select {
	// set a timeout of 3 seconds
	// if the request is handled within 3 seconds, return "Hello, World!"
	// else return an error
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

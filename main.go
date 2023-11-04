package main

import (
	"context"
	"fmt"
)

type key string

func main() {
	ctx := context.WithValue(context.Background(), key("name"), "Alice")
	value := ctx.Value(key("name"))
	fmt.Println("Name:", value)
}

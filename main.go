package main

import (
	"context"
	"fmt"
	"time"
)

type MyContext struct {
	context.Context
	RequestID string
}

func WithRequestID(ctx context.Context, requestID string) *MyContext {
	return &MyContext{
		Context:   ctx,
		RequestID: requestID,
	}
}

func (c *MyContext) Deadline() (deadline time.Time, ok bool) {
	return c.Context.Deadline()
}

func (c *MyContext) Done() <-chan struct{} {
	return c.Context.Done()
}

func (c *MyContext) Err() error {
	return c.Context.Err()
}

func (c *MyContext) Value(key interface{}) interface{} {
	if key == "requestID" {
		return c.RequestID
	}
	return c.Context.Value(key)
}

func main() {
	ctx := context.Background()

	ctxWithRequestID := WithRequestID(ctx, "12345")

	requestID := ctxWithRequestID.Value("requestID").(string)

	fmt.Println("Request ID: ", requestID)
}

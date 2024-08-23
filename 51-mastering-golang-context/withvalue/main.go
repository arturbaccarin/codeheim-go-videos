package main

import (
	"context"
	"fmt"
)

func main() {
	// Create a new context and store a value in it
	ctx := context.WithValue(context.Background(), "userID", 42)

	// Pass the context to a function that needs access to the value
	ProcessRequest(ctx)
}

// ProcessRequest simulates processing a request with access to context values
func ProcessRequest(ctx context.Context) {
	// Extract the value from the context using the same kwy
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		fmt.Println("userID not found in context")
		return
	}

	// Use the extracted value
	fmt.Printf("Processing request for user ID: %d\n", userID)

	// Simulate passing the context further down the call chain
	FurtherProcessing(ctx)
}

// FurtherProcessing simulates another function that needs access to context value
func FurtherProcessing(ctx context.Context) {
	// Extract the value from the context again
	userID, ok := ctx.Value("userID").(int)
	if !ok {
		fmt.Println("userID not found in context")
		return
	}

	// Use the value to perform some operation
	fmt.Printf("Further processing for user ID: %d\n", userID)
}

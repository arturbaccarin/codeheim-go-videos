package withcancel

import (
	"context"
	"fmt"
	"time"
)

/*
Context carries deadlines, cancellation
signals and request-scoped values
between goroutines
*/

func main() {
	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())

	// Laucnh a goroutine that will listen to the context cancellation
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 1 canceled:", ctx.Err())
				return
			default:
				fmt.Println("Goroutine 1 running")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Laucnh a goroutine that will listen to the context cancellation
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 2 canceled:", ctx.Err())
				return
			default:
				fmt.Println("Goroutine 2 running")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	// Simluate some work in the main function
	fmt.Println("Main function running")
	time.Sleep(2 * time.Second)

	// Cancel the context, whic will signal all goroutines to stop
	fmt.Println("Main function canceled")
	cancel()

	// Give goroutine time to finish
	time.Sleep(1 * time.Second)
	fmt.Println("Main function finished")
}

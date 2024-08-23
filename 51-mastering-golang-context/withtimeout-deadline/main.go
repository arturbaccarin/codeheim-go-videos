package withtimeout

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	// Create a context with a dedline
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancelFunc := context.WithDeadline(context.Background(), deadline)
	defer cancelFunc()

	// Create a context with a timeout of 3 seconds
	// ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancelFunc()

	// Perform a search with the context
	res, err := Search(ctx, "random string")
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("response: %s\n", res)
}

// Search simluates a serach operation that can be canceled via context
func Search(ctx context.Context, query string) (string, error) {
	// Channel to receive the response from the slow function
	resp := make(chan string)
	go func() {
		resp <- RandomSleepAndReturnAPI(query)
		close(resp)
	}()

	// Wait to either the response or the context to be done
	for {
		select {
		case dst := <-resp:
			return dst, nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

// RandomSleepAndReturnAPI simulates a slow API call by sleeping for a random duration
func RandomSleepAndReturnAPI(query string) string {
	// Create a new random number generator with a custom seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random duration
	randomDuration := time.Duration(rng.Int63n(int64(5 * time.Second)))

	// Sleep for the random duration
	time.Sleep(randomDuration)

	// Return a quirky message after sleeping
	return fmt.Sprintf("It took us %v... Hope it was worth the wait!", randomDuration)
}

package main

import (
	"context"
	"fmt"
	"time"
)

func acquireLock(client *redis.client, lockKey string, timeout time.Duration) bool {
	ctx := context.Background()

	// Try to acquire the lock with SETNX command (SET if Not eXists)
	lockAcquired, err := client.SetNX(ctx, lockKey, "1", timeout).Result()
	if err != nil {
		fmt.Println("error acquiring lock: ", err)
		return false
	}

	return lockAcquired
}

func releaseLock(client *redis.client, lockKey string) {
	ctx := context.Background()
	client.Del(ctx, lockKey)
}

func main() {
	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	defer client.Close()

	// Define the lock key and lock timeout
	lockKey := "my-lock"
	lockTimeout := 10 * time.Second

	// Acquire the lock
	if acquireLock(client, lockKey, lockTimeout) {
		fmt.Println("Lock acquired successfully")

		// Simulate some work with the lock
		time.Sleep(20 * time.Second)
		fmt.Println("Work done!")

		// Release the lock
		releaseLock(client, lockKey)
		fmt.Println("Lock released successfully")
	} else {
		fmt.Println("failed to acquire lock. Resource is already locked")
	}
}

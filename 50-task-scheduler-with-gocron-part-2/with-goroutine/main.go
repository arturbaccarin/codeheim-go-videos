package main

import (
	"log"
	"time"
)

func runAtSpecificTime(t time.Time, ch chan bool) {
	log.Println("Goroutine starts!")
	<-time.After(time.Until(t))
	log.Println("Task is running")
	log.Println("Task completed!")
	ch <- true
}

func main() {
	// Specific time you want to run the code
	runTime := time.Now().Add(5 * time.Second)

	// channel to signal when the task is done
	done := make(chan bool)

	// start a goroutine that waits until the specific time
	go runAtSpecificTime(runTime, done)

	// wait for the task to complete
	<-done
}

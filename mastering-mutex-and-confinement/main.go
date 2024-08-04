package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func buyTicket_(wg *sync.WaitGroup, userId int, remainingTickets *int) {
	defer wg.Done()

	mutex.Lock()
	if *remainingTickets > 0 {
		*remainingTickets--
		fmt.Printf("User %d bought a ticket. There are %d tickets left\n", userId, *remainingTickets)
	} else {
		fmt.Printf("User %d cannot buy a ticket. There are no tickets left\n", userId)
	}
	mutex.Unlock()
}

func main_() {
	var tickets int = 500

	var wg sync.WaitGroup

	for userId := 0; userId < 2000; userId++ {
		wg.Add(1)
		go buyTicket_(&wg, userId, &tickets)
	}

	wg.Wait()
}

// go run main.go | grep bought

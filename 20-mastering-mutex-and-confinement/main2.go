package main

import (
	"fmt"
	"sync"
)

func manageTicket(ticketChan chan int, doneChan chan bool, tickets *int) {
	for {
		select {
		case user := <-ticketChan:

			if *tickets > 0 {
				*tickets--
				fmt.Printf("User %d bought a ticket. There are %d tickets left\n", user, *tickets)
			} else {
				fmt.Printf("User %d cannot buy a ticket. There are no tickets left\n", user)
			}
		case <-doneChan:
			fmt.Println("All tickets have been sold")
		}
	}
}

func buyTicket(wg *sync.WaitGroup, ticketChan chan int, userId int) {
	defer wg.Done()
	ticketChan <- userId
}

func main() {
	var wg sync.WaitGroup
	tickets := 500

	ticketChan := make(chan int)
	doneChan := make(chan bool)

	go manageTicket(ticketChan, doneChan, &tickets)

	for userId := 0; userId < 2000; userId++ {
		wg.Add(1)
		go buyTicket(&wg, ticketChan, userId)
	}

	wg.Wait()
	doneChan <- true
}

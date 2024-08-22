package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	faktory "github.com/contribsys/faktory/client"
	"github.com/go-co-op/gocron/v2"
)

func main() {
	s, err := gocron.NewScheduler()
	if err != nil {
		// handle error
		panic(err)
	}

	// add a Cron job to the scheduler
	s.NewJob(
		gocron.CronJob(
			"0 10 * * *",
			false,
		),
		gocron.NewTask(
			func() {
				client, err := faktory.Open()
				if err != nil {
					log.Fatal(err)
				}
				job := faktory.NewJob("report", "test@codeheim.io")
				job.Queue = "critical"
				err = client.Push(job)
				if err != nil {
					log.Println("Error pushing job")
				}
			},
		),
		gocron.WithName("Send report at 1000 hrs everyday"),
	)

	// Start the scheduler
	s.Start()

	// Set up a channel to listen for interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("\nInterrupt signal received. Exiting...")
		_ = s.Shutdown()
		os.Exit(0)
	}()

	for {
	}
}

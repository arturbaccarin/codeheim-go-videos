package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/google/uuid"
)

func main() {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	// add a simple job to the scheduler
	j, err := s.NewJob(
		gocron.DurationJob(
			30*time.Second,
		),
		gocron.NewTask(
			func(a string) {
				log.Println(a)
			},
			"Every 30 seconds",
		),
		gocron.WithName("job: every 30 seconds"),
		gocron.WithEventListeners(
			gocron.BeforeJobRuns(
				func(jobID uuid.UUID, jobName string) {
					log.Printf("Job starting: %s, %s", jobID.String(), jobName)
				},
			),
			gocron.AfterJobRuns(
				func(jobID uuid.UUID, jobName string) {
					log.Printf("Job starting: %s, %s", jobID.String(), jobName)
				},
			),
			gocron.AfterJobRunsWithError(
				func(jobID uuid.UUID, jobName string, err error) {
					log.Printf("Job had an error: %s, %s, %v", jobID.String(), jobName, err)
				},
			),
		),
	)
	if err != nil {
		panic(err)
	}

	log.Println(j.ID())

	// add a cron job to the scheduler
	s.NewJob(
		gocron.CronJob(
			"*/10 * * * *",
			false,
		),
		gocron.NewTask(
			func(a string) {
				log.Println(a)
			},
			"CronJob: Every 10min",
		),
		gocron.WithName("job: CronJob: Every 10min"),
	)

	// add a Daily job to the scheduler
	s.NewJob(
		gocron.DailyJob(
			1, // Runs Every Day
			gocron.NewAtTimes(
				gocron.NewAtTime(23, 10, 00),
				gocron.NewAtTime(05, 30, 00),
			),
		),
		gocron.NewTask(
			func(a string, b string) {
				log.Println(a, b)
			},
			"Dailyjob", "Runs everyday",
		),
		gocron.WithName("job: Dailyjob"),
	)

	// add a One Time job to the shceduler
	s.NewJob(
		gocron.OneTimeJob(
			gocron.OneTimeJobStartDateTime(time.Now().Add(5*time.Minute)),
		),
		gocron.NewTask(
			func(a string, b string) {
				log.Println(a, b)
			},
			"Dailyjob", "Runs everyday",
		),
		gocron.WithName("job: Dailyjob"),
	)

	// add a Random duration job to the shceduler
	s.NewJob(
		gocron.DurationRandomJob(2*time.Minute, 5*time.Minute),
		gocron.NewTask(
			func(a string, b string) {
				log.Println(a, b)
			},
			"Dailyjob", "Runs everyday",
		),
		gocron.WithName("job: Dailyjob"),
	)

	// Start the scheduler
	s.Start()

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

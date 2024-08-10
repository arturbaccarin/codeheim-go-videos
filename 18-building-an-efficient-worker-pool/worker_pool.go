package main

import (
	"fmt"
	"sync"
	"time"
)

// Task definition
type Task interface {
	Process()
}

// Email task definition
type EmailTask struct {
	Email       string
	Subject     string
	MessageBody string
}

// Way to process the tasks
func (t *EmailTask) Process() {
	fmt.Println("Sending email to %d\n", t.Email)

	// Simulate a time consuming process
	time.Sleep(2 * time.Second)
}

// Image processing task
type ImageProcessing struct {
	ImageUrl string
}

func (t *ImageProcessing) Process() {
	fmt.Println("Processing the image %s\n", t.ImageUrl)

	// Simulate a time consuming process
	time.Sleep(5 * time.Second)
}

// Worker pool definition
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

// Functions to execute the worker pool
func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	// Initialize the tasks channel
	wp.tasksChan = make(chan Task, len(wp.Tasks))

	// Start the workers
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	// Send tasks to the tasks channel
	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}
	close(wp.tasksChan)

	// Wait for all tasks to be processed
	wp.wg.Wait()
}

package main

import "fmt"

func main() {
	// Create new tasks
	// tasks := make([]Task, 20)

	// for i := 0; i < 20; i++ {
	// 	tasks[i] = Task{ID: i + 1}
	// }

	tasks := []Task{
		&EmailTask{
			Email:       "a@a.com",
			Subject:     "test",
			MessageBody: "test"},

		&ImageProcessing{
			ImageUrl: "a.a.com"},
	}

	// Create a worker pool
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5, // number of workers that can run at a time
	}

	// Run the worker pool
	wp.Run()
	fmt.Println("done running worker pool")
}

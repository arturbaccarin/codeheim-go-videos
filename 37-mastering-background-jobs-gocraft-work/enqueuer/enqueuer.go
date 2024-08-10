package main

import (
	"log"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

// Redis pool
var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:6379")
	},
}

// Create job enqueuer
var enqueuer = work.NewEnqueuer("demo_app", redisPool)

func main() {
	_, err := enqueuer.Enqueue("email",
		work.Q{
			"userID":  10,
			"subject": "just testing",
		})
	// work.Q{
	// 	"email":   "a@a.com",
	// 	"subject": "test",
	// 	"message": "test",
	// })
	if err != nil {
		log.Fatal(err)
	}

	_, err = enqueuer.Enqueue("report",
		work.Q{
			"userID": 5,
		})
	if err != nil {
		log.Fatal(err)
	}
}

// workwebui -redis="redis://localhost:6379" -ns="demo_app" -listen=":5040"

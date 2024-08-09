package processjobs

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

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

type User struct {
	ID          int64
	Email, Name string
}

type Context struct {
	currentUser *User
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("starting a new job: ", job.Name, " with ID: ", job.ID)
	return next()
}

func (c *Context) FindCurrentUser(job *work.Job, next work.NextMiddlewareFunc) error {
	// If there's a user_id param
	if _, ok := job.Args["userID"]; ok {
		userID := job.ArgInt64("userID")

		// FIXME: query the DB and get the user
		c.currentUser = &User{ID: userID, Email: "test" + strconv.Itoa(int(userID)) + "@test.com", Name: "test"}
		if err := job.ArgError(); err != nil {
			return err
		}
	}

	return next()
}

// Create job enqueuer
var enqueuer = work.NewEnqueuer("demo_app", redisPool)

func main() {
	pool := work.NewWorkerPool(Context{}, 10, "demo_app", redisPool)

	// Middlewares
	pool.Middleware((*Context).Log)
	pool.Middleware((*Context).FindCurrentUser)

	// Name to job map
	// pool.Job("email", SendEmail)
	pool.JobWithOptions("email",
		work.JobOptions{Priority: 10, MaxFails: 1}, (*Context).SendEmail)
	pool.JobWithOptions("report",
		work.JobOptions{Priority: 10, MaxFails: 1}, (*Context).Report)

	pool.Start()

	// Wait for a signal to quit
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	// Stop the pool
	pool.Stop()
}

func (c *Context) SendEmail(job *work.Job) error {
	// addr := job.ArgString("email")
	addr := c.currentUser.Email
	subject := job.ArgString("subject")
	if err := job.ArgError(); err != nil {
		return err
	}

	fmt.Printf("Sending email to %s with subject %s\n", addr, subject)
	time.Sleep(time.Second * 2)
	return nil
}

func (c *Context) Report(job *work.Job) error {
	fmt.Println("Preparing report...")
	time.Sleep(time.Second * 10)
	// send the report via email
	enqueuer.Enqueue("email", work.Q{"email": c.currentUser.Email, "subject": "report"})
	return nil
}

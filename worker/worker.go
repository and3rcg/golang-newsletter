package main

// this is the server, it'll listen to Redis' queue and execute tasks whenever
// usually, it's run via the command go run worker/worker.go from the project's root folder

import (
	"log"
	"newsletter-go/tasks"
	"os"

	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading.env file")
	}

	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     os.Getenv("REDIS_HOST"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		},
		asynq.Config{Concurrency: 2},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeTaskSendNewsletterEmails, tasks.HandlerTaskSendNewsletterEmails)
	// mux.HandleFunc(tasks.TypeDemoTask, tasks.HandlerTaskDemo)
	// ^^^^^^ you can register more tasks here by adding more lines accordingly

	if err := srv.Run(mux); err != nil {
		log.Fatalf("Could not run asynq server: %v", err)
	}
}

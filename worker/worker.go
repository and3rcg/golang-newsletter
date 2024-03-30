package main

// esse aqui é o servidor, ele vai checar o redis para executar tasks
// para rodar o worker, execute go run worker/worker.go a partir da raiz do projeto

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
	mux.HandleFunc(tasks.TypeDemoTask, tasks.HandlerTaskDemo)
	// mux.HandleFunc(tasks.TypeDemoTask, tasks.HandlerTaskDemo)
	// ^^^^^^ you can register more tasks here by adding more lines accordingly

	if err := srv.Run(mux); err != nil {
		log.Fatalf("Could not run asynq server: %v", err)
	}
}

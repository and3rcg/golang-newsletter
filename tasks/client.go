package tasks

import (
	"log"
	"os"
	"sync"

	"github.com/hibiken/asynq"
)

var (
	client *asynq.Client
	once   sync.Once
)

// esse método é rodado no main.go pra iniciar o cliente
func InitWorkerClient() {
	log.Println("Starting worker client...")
	once.Do(func() {
		client = asynq.NewClient(asynq.RedisClientOpt{
			Addr:     os.Getenv("REDIS_HOST"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})
	})
}

// esse carinha aqui é feito pra se usar com defer lá no main.go junto ao InitWorkerClient
func CloseWorkerClient() {
	if client != nil {
		client.Close()
	}
}

// se precisarmos do cliente em algum lugar do código, pegamos ele por aqui
func GetWorkerClient() *asynq.Client {
	return client
}

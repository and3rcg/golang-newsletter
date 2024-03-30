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

// this function must be called in main.go to start the client
func InitWorkerClient() {
	log.Println("Starting worker client...")
	// notice that the client must be set to EXACTLY the same address and DB as the server is listening to
	once.Do(func() {
		client = asynq.NewClient(asynq.RedisClientOpt{
			Addr:     os.Getenv("REDIS_HOST"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})
	})
}

// this one must be used with defer right after initializing the client, to ensure safe closing
func CloseWorkerClient() {
	if client != nil {
		client.Close()
	}
}

// if we need the client anywhere in the code, we call this function
func GetWorkerClient() *asynq.Client {
	return client
}

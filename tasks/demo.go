package tasks

// General structure of the tasks fliles:
// 1. A list of task types. In the worker's server (worker/worker.go), mux.HandleFunc registers tasks by the names of the types,
// and we call them via the client (client.Enqueue(task))
// 2. A struct with the JSON payload of the task's arguments (if applicable).
// If you're familiar with Celery, you might know that the arguments must be JSON serializable. It works just the same here.
// 3. The task's constructor function: This function will generate the task's payload and return the task object to be enqueued.
// 4. Task handler: it's a func(ctx context.Context, t *asynq.Task) that returns an error. This is the task's algorithm itself.

// IMPORTANT: you must pass the payload to the task as an array of bytes (a.k.a use JSON.Marshal)
// We can also register more than one task in a file, just declare more types, payload structs and handlers, registering them accordingly.

// Notice that the tasks in the constructor are called by their type. The type => handler association is made by the server.

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

const (
	TypeDemoTask = "demo_task:execute"
)

type DemoTaskPayload struct {
	WaitTime uint
}

func NewTaskDemo(timeSeconds uint) *asynq.Task {
	payload := DemoTaskPayload{
		WaitTime: timeSeconds,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil
	}

	task := asynq.NewTask(TypeDemoTask, payloadBytes)

	// I can also call the client from here by using client.Enqueue(task) and returning nothing

	return task
}

func HandlerTaskDemo(ctx context.Context, t *asynq.Task) error {
	var args DemoTaskPayload
	err := json.Unmarshal(t.Payload(), &args)
	fmt.Println("Task started!")
	if err != nil {
		fmt.Println(err)
		return err
	}

	for i := range [12]int{} {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Task finished successfully!")
	return nil
}

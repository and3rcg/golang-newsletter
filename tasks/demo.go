package tasks

// Estrutura geral dos arquivos de tasks:
// 1. Uma lista de tipos de task. No worker server (worker/worker.go), o mux.HandleFunc cadastra as tasks com os nomes dos tipos,
// e chamamos eles pelo client
// 2. Um struct com o payload dos argumentos a serem enviados. Lembra que o Celery precisa de argumentos serializáveis? aqui é igual
// 3. Função construtora da task: é essa função que vai gerar o payload que a task vai receber para trabalhar
// 4. Handler da task: é uma func(ctx context.Context, t *asynq.Task) que retorna um erro

// IMPORTANTE: o handler da task só recebe o payload se for no formato array de bytes (ou json.Marshal)
// Também podemos cadastrar mais tasks em um arquivo, basta declarar novos tipos de task, payloads, funções construtoras e handlers

// note que as tasks são chamadas pelo tipo (veja o asynq.NewTask). O server vincula os tipos aos handlers de tasks

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

	// também posso chamar o client direto daqui com client.Enqueue(task)

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

package handler

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
)

func Telegram(ctx context.Context, task *asynq.Task) error {
	b := task.Payload()
	log.Println(string(b))
	return nil
}

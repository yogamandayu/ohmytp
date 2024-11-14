package handler

import (
	"context"
	"github.com/hibiken/asynq"
	"log"
)

func Telegram(ctx context.Context, task *asynq.Task) error {
	b := task.Payload()
	log.Println(string(b))
	return nil
}

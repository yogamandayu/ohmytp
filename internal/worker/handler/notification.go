package handler

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/hibiken/asynq"
	"github.com/yogamandayu/ohmytp/consts"
	"github.com/yogamandayu/ohmytp/internal/app"
	"github.com/yogamandayu/ohmytp/internal/domain/entity"
	"github.com/yogamandayu/ohmytp/pkg/telegram"
)

// Notification is a worker handler for notification.
type Notification struct {
	App *app.App
}

// NewNotificationHandler is a constructor.
func NewNotificationHandler(app *app.App) *Notification {
	return &Notification{
		App: app,
	}
}

// Handler is a notification handler.
func (n Notification) Handler(ctx context.Context, task *asynq.Task) error {
	b := task.Payload()

	var payload entity.WorkerNotification

	err := json.Unmarshal(b, &payload)
	if err != nil {
		return err
	}
	if payload.Via == consts.ViaTelegramWorkerNotification {
		bot := telegram.NewTelegramBot(n.App.Log, n.App.Config.TelegramBot.Config)
		data, ok := payload.Data.(map[string]interface{})
		if !ok {
			n.App.Log.Error("invalid assertion data via telegram")
			return errors.New("worker.error.handler.telegram.invalid_assertion")
		}
		return bot.SendMessage(data["message"].(string))
	}

	return nil
}

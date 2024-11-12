package telegram

import (
	"log/slog"
)

type Bot struct {
	Log *slog.Logger

	UserID string
}

func NewTelegramBot(log *slog.Logger) *Bot {
	return &Bot{
		Log: log,
	}
}

func (b *Bot) SendMessage(message string) error {

	return nil
}

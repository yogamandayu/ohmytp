package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

// Config is a telegram bot config.
type Config struct {
	Token  string
	ChatID string
}

type Bot struct {
	Log *slog.Logger
	Config
}

type MessagePayload struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func NewTelegramBot(log *slog.Logger, config *Config) *Bot {
	return &Bot{
		Log:    log,
		Config: *config,
	}
}

func (b *Bot) SendMessage(message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.Config.Token)

	payload := MessagePayload{
		ChatID: b.Config.ChatID,
		Text:   message,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var respBody interface{}
		_ = json.NewDecoder(resp.Body).Decode(&respBody)
		b.Log.Error("error sending message to telegram bot, err:%v", respBody)
		return errors.New("telegram.error.send_message")
	}

	return nil
}

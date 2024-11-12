package telegram_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yogamandayu/ohmytp/internal/tests"
	"github.com/yogamandayu/ohmytp/pkg/telegram"
)

func TestTelegramBotSendMessage(t *testing.T) {
	testSuite := tests.NewTestSuite()
	defer func() {
		t.Cleanup(testSuite.Clean)
	}()
	testSuite.LoadApp()

	t.Run("Test send message to telegram bot", func(t *testing.T) {
		bot := telegram.NewTelegramBot(testSuite.App.Log, testSuite.App.Config.TelegramBot.Config)
		err := bot.SendMessage(fmt.Sprintf("Test telegram bot at %s", time.Now().Format(time.RFC3339)))
		require.NoError(t, err)
	})
}

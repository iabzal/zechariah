package zechariah

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (b *Bot) ConsultationCmd(upd tgbotapi.Update, questionKey int) {
	_ = questionKey
	message := `https://zechariahc.com/`
	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message))
	reply.ParseMode = "html"

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("failed to send help message", zap.Error(err))
	}
}

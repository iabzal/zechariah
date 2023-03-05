package iskanderzhuma

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (b *Bot) ContactCmd(upd tgbotapi.Update) {
	message := `
Мои контакты

Телефон:
+7 701 111 11 77

Почта
info@zechariah.com

Соц. сети:
INSTAGRAM: https://instagram.com/
WHATSAPP: https://wa.me/
	`
	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message))
	reply.ParseMode = "html"

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("failed to send help message", zap.Error(err))
	}
}

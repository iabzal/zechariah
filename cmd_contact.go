package zechariah

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (b *Bot) ContactCmd(upd tgbotapi.Update, questionKey int) {
	_ = questionKey
	message := `
Мои контакты

Телефон:
+971585876770

Почта:
Z.d770@yahoo.com

Сайт:
https://zechariahc.com/

Соц. сети:
INSTAGRAM: https://instagram.com/zechariah.9?igshid=YmMyMTA2M2Y=
WHATSAPP: https://wa.me/
	`
	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message))
	reply.ParseMode = "html"

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("failed to send help message", zap.Error(err))
	}
}

package zechariah

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (b *Bot) ServiceCmd(upd tgbotapi.Update, questionKey int) {
	_ = questionKey
	message := `
<b>Я предлагаю широкий спектр услуг, чтобы помочь вашей компании достичь успеха и стать еще более процветающей.</b>

1.	Индивидуальные консультации.
2.	Коммерческая консультация бизнеса и коллективов
3.	Семейные консультации
`
	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message))
	reply.ParseMode = "html"

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("failed to send help message", zap.Error(err))
	}
}

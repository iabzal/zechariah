package iskanderzhuma

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (b *Bot) ServiceCmd(upd tgbotapi.Update) {
	message := `
Я предлагаю широкий спектр услуг, чтобы помочь вашей компании достичь успеха и стать еще более процветающей.

— Медиаторство, решение проблем, 
  выстраивание благоприятного климата в компании.

— Работа с генеральными, директорами компании.

— Работа с топ менеджерами на эффективность работы.

— Индивидуальные консультации.
`
	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message))
	reply.ParseMode = "html"

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("failed to send help message", zap.Error(err))
	}
}

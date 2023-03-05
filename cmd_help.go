package iskanderzhuma

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func (b *Bot) HelpCmd(upd tgbotapi.Update) {
	message := `
		<b>О цифровой науке Сюцай.</b>

		Сюцай переводится как «знать на расстоянии» и имеет очень древние корни. Сразу стоит отметить, что Сюцай – это НЕ нумерология о которой знает большинство, хотя она тоже оперирует цифрами.

		Например, число даты рождения — это своего рода код от сейфа души, зная который, вы можете стать более осознанным в реальной жизни и гармонично развиваться во всех сферах.
		
		Основная задача науки Сюцай — это развитие человеком своего сознания и реализация души. Чтобы узнать, как это работает, необходимо ознакомиться с базовыми понятиями.
	`
	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message))
	reply.ParseMode = "html"

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("failed to send help message", zap.Error(err))
	}
}

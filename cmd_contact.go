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

Telegram: @Zechariah8
What’sApp: https://wa.me/971585681122
www.zechariahc.com,
www.zechariah8.com,
www.zechariah8.vip
E- mail: mentor@zechariah8.com
Instagram: https://instagram.com/zechariah.9?igshid=YmMyMTA2M2Y=
https://t.me/zechariah9
@zechariah_help_bot
	`
	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message))
	reply.ParseMode = "html"

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("failed to send help message", zap.Error(err))
	}
}

package zechariah

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

const (
	man   = "man"
	woman = "woman"
)

func (b *Bot) GetAboutMe(upd tgbotapi.Update) {
	//message := `Введите свою дату рождения как на примере <b>18.08.1993</b>:`
	message := `Ваш пол`
	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message))
	reply.ParseMode = "html"

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Мужской", man),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Женский", woman),
		),
	)
	reply.ReplyMarkup = keyboard

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("failed to send help message", zap.Error(err))
	}
}

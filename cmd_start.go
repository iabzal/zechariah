package iskanderzhuma

import (
	"fmt"
	"go.uber.org/zap"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type replyKeyboardValue string

const (
	ReplyCalcDate     = replyKeyboardValue("Сделать расчет")
	ReplyService      = replyKeyboardValue("Мои услуги")
	ReplyContact      = replyKeyboardValue("Мои контакты")
	ReplyConsultation = replyKeyboardValue("Консультация")
)

func (b *Bot) StartCmd(upd tgbotapi.Update) {
	name := upd.Message.From.UserName
	if name == "" {
		name = upd.Message.From.FirstName
	}
	message := `Добро пожаловать в мой бот-помощник, %s!
Здесь вы можете получить предварительный расчет.`
	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message, name))
	reply.ParseMode = "html"

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(string(ReplyCalcDate)),
			tgbotapi.NewKeyboardButton(string(ReplyService)),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(string(ReplyContact)),
			tgbotapi.NewKeyboardButton(string(ReplyConsultation)),
		),
	)
	reply.ReplyMarkup = keyboard
	reply.DisableWebPagePreview = true

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("failed to send start message", zap.Error(err))
	}
}

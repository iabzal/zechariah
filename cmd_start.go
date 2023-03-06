package zechariah

import (
	"fmt"
	"go.uber.org/zap"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type replyKeyboardValue string

const (
	aboutMeTxt      = "Хочу узнать о себе"
	serviceTxt      = "Мои услуги"
	contactTxt      = "Мои контакты"
	consultationTxt = "Консультация"

	ReplyAboutMe      = replyKeyboardValue(aboutMeTxt)
	ReplyService      = replyKeyboardValue(serviceTxt)
	ReplyContact      = replyKeyboardValue(contactTxt)
	ReplyConsultation = replyKeyboardValue(consultationTxt)
)

func (b *Bot) StartCmd(upd tgbotapi.Update, questionKey int) {
	_ = questionKey
	name := upd.Message.From.UserName
	if name == "" {
		name = upd.Message.From.FirstName
	}
	message := `Добро пожаловать, %s!
Здесь вы можете получить информацию о себе.`
	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message, name))
	reply.ParseMode = "html"

	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(string(ReplyAboutMe)),
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

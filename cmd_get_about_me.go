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

var questions = map[int]string{
	1: "Введите свое полное имя:",
	2: "Ваша цель?",
	3: "Что Вас беспокоит?",
	4: "Занимаетесь ли Вы спортом?",
	5: "Как Вы оцениваете качества Вашего сна?",
	6: "Есть ли у Вас хобби?",
	7: "Ваш пол?",
}

func (b *Bot) GetAboutMe(upd tgbotapi.Update, questionKey int) {
	//message := `Введите свою дату рождения как на примере <b>18.08.1993</b>:`
	message := questions[questionKey]
	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message))
	reply.ParseMode = "html"

	if questionKey == 7 {
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Мужской", man),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Женский", woman),
			),
		)
		reply.ReplyMarkup = keyboard
	}

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("failed to send help message", zap.Error(err))
	}
}

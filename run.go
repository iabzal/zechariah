package zechariah

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
	"regexp"
	"strconv"
	"time"
)

// Run listens updates
func (b *Bot) Run() {
	updatesCfg := tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 10,
	}
	var gender, dateOfBirth = "", ""
	for upd := range b.GetUpdatesChan(updatesCfg) {
		if upd.Message != nil {
			if upd.Message.IsCommand() {
				key := upd.Message.Command()
				if cmd, ok := b.commands[commandKey(key)]; ok {
					go cmd.action(upd)
				} else {
					b.logger.Error("command handler not found", zap.String("cmd", key))
				}
				continue
			}

			if regexp.MustCompile(`\d`).MatchString(upd.Message.Text) {
				dateOfBirth = upd.Message.Text
				go b.CalcDateCmd(upd, gender)
			}

			if cmd, ok := b.replyToCommand(upd.Message.Text); ok {
				go cmd.action(upd)
				continue
			}
		} else if upd.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(upd.CallbackQuery.ID, upd.CallbackQuery.Data)
			if _, err := b.Request(callback); err != nil {
				return
			}

			if upd.CallbackQuery.Data == man || upd.CallbackQuery.Data == woman {
				gender = upd.CallbackQuery.Data
				askDateOfBirth(b, upd)
			} else if upd.CallbackQuery.Data == getYearDesc {
				sendYearResult(b, upd, gender, dateOfBirth)
			} else {
				msg := tgbotapi.NewMessage(upd.CallbackQuery.Message.Chat.ID, upd.CallbackQuery.Data)
				if err := b.apiRequest(msg); err != nil {
					return
				}
			}
		}
	}
}

func (b *Bot) Stop() {
}

func askDateOfBirth(b *Bot, upd tgbotapi.Update) {
	msg := tgbotapi.NewMessage(upd.CallbackQuery.Message.Chat.ID, "Дата Вашего рождения,\nнапример: 24.05.1994")
	err := b.apiRequest(msg)
	if err != nil {
		return
	}
}

func sendYearResult(b *Bot, upd tgbotapi.Update, gender string, dateOfBirth string) {
	digits := regexp.MustCompile(`\D+`).ReplaceAllString(dateOfBirth, "")
	yearResult := getYearResult(digits)
	yearResultInt, _ := strconv.Atoi(yearResult)

	res := getYearDecision(yearResultInt, gender)
	msg := tgbotapi.NewMessage(upd.CallbackQuery.Message.Chat.ID, res)
	if err := b.apiRequest(msg); err != nil {
		return
	}

	time.Sleep(1 * time.Second)
	info := "Чтобы еще более подробно узнать о себе, полноценно реализовать Ваш потенциал, совершить ментальное обновление Вашего сознания и квантовый скачок Вашего развития, нажмите кнопку “Готов стать лучше"
	msgInfo := tgbotapi.NewMessage(upd.CallbackQuery.Message.Chat.ID, info)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Готов стать лучше", "https://zechariahc.com/"),
		),
	)
	msgInfo.ReplyMarkup = keyboard

	if err := b.apiRequest(msgInfo); err != nil {
		return
	}
}

package zechariah

import (
	"fmt"
	"go.uber.org/zap"
	"regexp"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const getYearDesc = "year_info"

func sumDigits(number int) int {
	sumResult := 0
	if number == 0 {
		return 0
	}
	sumResult = number%10 + sumDigits(number/10)
	if sumResult > 9 {
		sumResult = sumResult%10 + sumDigits(sumResult/10)
	}
	return sumResult
}

func getYearResult(digits string) string {
	t := time.Now()
	year := strconv.Itoa(t.Year())
	data := digits[0:4] + year
	convertToInt, _ := strconv.Atoi(data)
	return strconv.Itoa(sumDigits(convertToInt))
}

func getMindResult(digits string) string {
	convertToInt, _ := strconv.Atoi(digits[0:2])
	return strconv.Itoa(sumDigits(convertToInt))
}

func (b *Bot) CalcDateCmd(upd tgbotapi.Update, gender string) {
	var message string
	digits := regexp.MustCompile(`\D+`).ReplaceAllString(upd.Message.Text, "")
	if len(digits) == 8 {
		mindResult := getMindResult(digits)
		mindResultInt, _ := strconv.Atoi(mindResult)
		message = getMindDecision(mindResultInt, gender)
	} else {
		message = `Неправильный формат.`
	}

	reply1 := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf("Точная информация на личной консультации."))
	reply1.DisableWebPagePreview = true

	if err := b.apiRequest(reply1); err != nil {
		b.logger.Error("failed to send start message", zap.Error(err))
	}
	time.Sleep(1 * time.Second)

	reply2 := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message))
	reply2.ParseMode = "html"
	reply2.DisableWebPagePreview = true

	if err := b.apiRequest(reply2); err != nil {
		b.logger.Error("failed to send start message", zap.Error(err))
	}
	time.Sleep(1 * time.Second)

	reply3 := tgbotapi.NewMessage(upd.Message.Chat.ID, "📲 хотите Я еще расскажу о вас?")
	reply3.DisableWebPagePreview = true

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Да", getYearDesc),
		),
	)
	reply3.ReplyMarkup = keyboard

	if err := b.apiRequest(reply3); err != nil {
		b.logger.Error("failed to send start message", zap.Error(err))
	}
}

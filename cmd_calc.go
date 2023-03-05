package iskanderzhuma

import (
	"fmt"
	"go.uber.org/zap"
	"regexp"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

func getMissionResult(digits string) string {
	convertToInt, _ := strconv.Atoi(digits)
	return strconv.Itoa(sumDigits(convertToInt))
}

func getMindResult(digits string) string {
	convertToInt, _ := strconv.Atoi(digits[0:2])
	return strconv.Itoa(sumDigits(convertToInt))
}

func (b *Bot) CalcDateCmd(upd tgbotapi.Update) {
	var message, requestMsg string
	digits := regexp.MustCompile(`\D+`).ReplaceAllString(upd.Message.Text, "")
	if len(digits) == 8 {
		mindResult := getMindResult(digits)
		mindResultInt, _ := strconv.Atoi(mindResult)

		missionResult := getMissionResult(digits)
		missionResultInt, _ := strconv.Atoi(missionResult)

		message = `
<b>Ваш результат:</b>

<b>Число Сознания: ` + mindResult + `</b>
` + getMindDecision(mindResultInt) + `

<b>Число Миссии: ` + missionResult + `</b>
` + getMissionDecision(missionResultInt) + `

📲 Если хотите получить больше информации оставьте свой номер телефона.`

	} else if len(digits) == 10 || len(digits) == 11 {
		requestMsg = `
НОВАЯ ЗАЯВКА.
Номер телефона: ` + upd.Message.Text

		message = `Спасибо за обращение! В ближайшее время мы свяжемся с Вами. 😊`
	} else {
		message = `Неправильный формат.`
	}

	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, fmt.Sprintf(message))
	reply.ParseMode = "html"
	reply.DisableWebPagePreview = true

	if err := b.apiRequest(reply); err != nil {
		b.logger.Error("failed to send start message", zap.Error(err))
	}

	if len(requestMsg) > 0 {
		replyRequest := tgbotapi.NewMessage(432101609, fmt.Sprintf(requestMsg))
		replyRequest.ParseMode = "html"
		replyRequest.DisableWebPagePreview = true

		if err := b.apiRequest(replyRequest); err != nil {
			b.logger.Error("failed to send start message", zap.Error(err))
		}
	}
}

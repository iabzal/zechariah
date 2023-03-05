package iskanderzhuma

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	commandEntity struct {
		key    commandKey
		desc   string
		action func(upd tgbotapi.Update)
	}
)

type commandKey string

const (
	StartCmdKey        = commandKey("start")
	CalcCmdKey         = commandKey("calc")
	ServiceCmdKey      = commandKey("service")
	ContactCmdKey      = commandKey("contact")
	ConsultationCmdKey = commandKey("consultation")
)

func (b *Bot) initCommands() error {
	commands := []commandEntity{
		{
			key:    StartCmdKey,
			desc:   "Запустить бота",
			action: b.StartCmd,
		},
		{
			key:    CalcCmdKey,
			desc:   "Сделать расчет",
			action: b.GetDateCmd,
		},
		{
			key:    ServiceCmdKey,
			desc:   "Мои услуги",
			action: b.ServiceCmd,
		},
		{
			key:    ContactCmdKey,
			desc:   "Мои контакты",
			action: b.ContactCmd,
		},
		{
			key:    ConsultationCmdKey,
			desc:   "Консультация",
			action: b.ConsultationCmd,
		},
	}

	tgCommands := make([]tgbotapi.BotCommand, 0, len(commands))
	for _, cmd := range commands {
		b.commands[cmd.key] = cmd
		tgCommands = append(tgCommands, tgbotapi.BotCommand{
			Command:     "/" + string(cmd.key),
			Description: cmd.desc,
		})
	}

	config := tgbotapi.NewSetMyCommands(tgCommands...)
	return b.apiRequest(config)
}

func (b *Bot) replyToCommand(text string) (commandEntity, bool) {
	switch replyKeyboardValue(text) {
	case ReplyCalcDate:
		cmd, ok := b.commands[CalcCmdKey]
		return cmd, ok
	case ReplyService:
		cmd, ok := b.commands[ServiceCmdKey]
		return cmd, ok
	case ReplyContact:
		cmd, ok := b.commands[ContactCmdKey]
		return cmd, ok
	case ReplyConsultation:
		cmd, ok := b.commands[ConsultationCmdKey]
		return cmd, ok
	}

	return commandEntity{}, false
}
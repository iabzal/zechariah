package iskanderzhuma

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type (
	Bot struct {
		*tgbotapi.BotAPI
		logger    *zap.Logger
		commands  map[commandKey]commandEntity
		callbacks map[callbackType]callbackFn
	}
)

// New creates bot instance
func New(logger *zap.Logger, token string, opts ...Option) (*Bot, error) {
	api, aErr := tgbotapi.NewBotAPI(token)
	if aErr != nil {
		return nil, aErr
	}

	var bo options
	for _, optFn := range opts {
		optFn(&bo)
	}
	api.Debug = bo.debug

	logger = logger.Named("bot")
	b := &Bot{
		BotAPI:    api,
		logger:    logger,
		commands:  make(map[commandKey]commandEntity),
		callbacks: make(map[callbackType]callbackFn),
	}

	if err := b.initCommands(); err != nil {
		return nil, err
	}
	b.initCallbacks()

	b.logger.Info("bot created", zap.String("username", api.Self.UserName))
	return b, nil
}

func (b *Bot) apiRequest(c tgbotapi.Chattable) error {
	_, err := b.Request(c)
	return err
}

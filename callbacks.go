package iskanderzhuma

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

type callbackType int

type (
	callbackEntity struct {
		cbType     callbackType
		id         string
		parentType callbackType
		parentIds  []string
		page       int
	}

	callbackFn func(upd tgbotapi.Update, entity callbackEntity)
)

func (c callbackEntity) Clone() callbackEntity {
	return callbackEntity{
		cbType:     c.cbType,
		id:         c.id,
		parentType: c.parentType,
		parentIds:  c.parentIds,
		page:       c.page,
	}
}

func (b *Bot) initCallbacks() {
	b.callbacks = map[callbackType]callbackFn{}
}

func unmarshallCb(data string) callbackEntity {
	d := strings.Split(data, ";")

	var cbType int
	if len(d) > 0 {
		cbType, _ = strconv.Atoi(d[0])
	}

	var id string
	if len(d) > 1 {
		id = d[1]
	}

	var pType int
	if len(d) > 2 {
		pType, _ = strconv.Atoi(d[2])
	}

	var parentIds []string
	if len(d) > 3 {
		parentIds = strings.Split(d[3], ".")
	}

	var page int
	if len(d) > 4 {
		page, _ = strconv.Atoi(d[4])
	}

	return callbackEntity{
		cbType:     callbackType(cbType),
		id:         id,
		parentType: callbackType(pType),
		parentIds:  parentIds,
		page:       page,
	}
}

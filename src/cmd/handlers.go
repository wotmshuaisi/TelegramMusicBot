package main

import (
	"time"

	tba "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func testhandler(ctx *tba.Update) {
	// max
	inlineQueryWG.Add(1)
	inlineTasksCount++
	defer func() {
		inlineTasksCount--
		inlineQueryWG.Done()
	}()

	logrus.Infof("offset: %s content: %s", ctx.InlineQuery.Offset, ctx.InlineQuery.Query)
	time.Sleep(time.Second * 40)
}

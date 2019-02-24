package main

import (
	"github.com/google/uuid"

	tba "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func testhandler(bot *tba.BotAPI, ctx *tba.Update) {
	// task limiter
	inlineQueryWG.Add(1)
	inlineTasksCount++
	defer func() {
		inlineTasksCount--
		inlineQueryWG.Done()
	}()
	// processing event
	u, _ := uuid.NewUUID()
	i := tba.NewInlineQueryResultAudio(u.String(), "http://fs.open.kugou.com/05914e4cdceb1d0e406dac5a0722168c/5c72eabb/G140/M04/00/01/LIcBAFvMQmaAVJNYAEgvmkjJyY0089.mp3", "Hello - Adele")

	config := tba.InlineConfig{
		InlineQueryID: ctx.InlineQuery.ID,
		IsPersonal:    true,
		Results:       []interface{}{&i},
	}

	res, err := bot.AnswerInlineQuery(config)
	if err != nil || !res.Ok {
		logrus.WithError(err).Warnf("content: %s res: %+v", ctx.InlineQuery.Query, res)
	}

}

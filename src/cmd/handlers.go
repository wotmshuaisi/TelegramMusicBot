package main

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/wotmshuaisi/TelegramMusicBot/src/lib/music"
	tba "github.com/wotmshuaisi/telegram-bot-api"
)

//  Music Inline Query handler
func musicInlineQuery(bot *tba.BotAPI, update *tba.Update, api music.API) {
	// counting tasks number for restriction
	inlineQueryWG.Add(1)
	inlineTasksCount++
	defer func() {
		inlineTasksCount--
		inlineQueryWG.Done()
	}()
	// processing event
	l, err := api.ListItem(update.InlineQuery.Query)
	if err != nil {
		logrus.WithError(err).Warnf("id: %s query: %s", update.InlineQuery.ID, update.InlineQuery.Query)
		return
	}

	config := tba.InlineConfig{
		InlineQueryID: update.InlineQuery.ID,
		IsPersonal:    true,
	}

	for _, v := range *l {
		config.Results = append(config.Results, tba.NewInlineQueryResultAudio(v.ID, v.URL, "👍 "+v.Title, v.Performer, v.Duration))
	}

	if config.Results == nil {
		uu, _ := uuid.NewUUID()
		config.Results = append(config.Results, tba.NewInlineQueryResultArticleMarkdown(uu.String(), "no result", "no result"))
	}

	res, err := bot.AnswerInlineQuery(config)
	if err != nil || !res.Ok {
		logrus.WithError(err).Warnf("query: %s res: %+v", update.InlineQuery.Query, res)
	}

}

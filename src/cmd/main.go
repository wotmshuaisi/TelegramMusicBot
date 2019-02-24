package main

import (
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/wotmshuaisi/TelegramMusicBot/src/lib/music/kugou"

	"github.com/sirupsen/logrus"
	tba "github.com/wotmshuaisi/telegram-bot-api"
)

func init() {
	// init logrus
	logFile, err := os.OpenFile("./log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})

	if err != nil {
		logrus.SetOutput(os.Stdout)
		return
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, logFile))
}

func main() {
	bot, err := tba.NewBotAPI(botToken)
	if err != nil {
		logrus.WithError(err).Fatalln("init bot api error")
	}

	bot.Debug, _ = strconv.ParseBool(debugMode)

	logrus.Infof("Current Bot: %s", bot.Self.UserName)

	// set timeout for get new messages
	u := tba.NewUpdate(0)
	u.Timeout = 60
	// get messages
	updates, _ := bot.GetUpdatesChan(u)

	// telegram event loop
	for update := range updates {
		// inline query handling
		if update.InlineQuery != nil {
			// skip empty content
			if strings.TrimSpace(update.InlineQuery.Query) == "" {
				continue
			}
			// inline query task limiter
			if inlineTasksCount >= maxInlineTaskCount {
				inlineQueryWG.Wait()
			}
			// music handler
			go musicInlineQuery(bot, &update, kugou.NewAPI())
		}
		// ....

		// sleep && getin next loop
		time.Sleep(time.Millisecond * 10)
		continue
	}

}

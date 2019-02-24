package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	tba "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/wotmshuaisi/TelegramMusicBot/src/lib/utils"
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
	bot, err := tba.NewBotAPI(utils.GetEnvWithFatal("TELEGRAM_TOKEN"))
	if err != nil {
		logrus.WithError(err).Fatalln("init bot api error")
	}

	bot.Debug, _ = strconv.ParseBool(utils.GetEnvWithDefault("TELEGRAM_DEBUG", "False"))

	logrus.Infof("Current Bot: %s", bot.Self.UserName)

	// set timeout for get new messages
	u := tba.NewUpdate(0)
	u.Timeout = 60
	// get messages
	updates, _ := bot.GetUpdatesChan(u)

	// running in command mode
	for update := range updates {
		// inline query handling
		if update.InlineQuery != nil || strings.TrimSpace(update.InlineQuery.Query) != "" {
			// inline query task limiter
			if inlineTasksCount >= maxInlineTaskCount {
				fmt.Println("current tasks", inlineTasksCount)
				inlineQueryWG.Wait()
			}
			go testhandler(&update)
		}
		// ....

		// sleep && getin next loop
		time.Sleep(time.Millisecond * 10)
		continue
	}

}

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

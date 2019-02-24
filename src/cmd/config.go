package main

import (
	"sync"

	"github.com/wotmshuaisi/TelegramMusicBot/src/lib/utils"
)

var (
	// lock for inline query task
	inlineQueryWG = sync.WaitGroup{}
	// couter for inline tasks
	inlineTasksCount = 0
	// maximum inline query task
	maxInlineTaskCount = 10
	// Bot token
	botToken = utils.GetEnvWithFatal("TELEGRAM_TOKEN")
	// debug mode
	debugMode = utils.GetEnvWithDefault("TELEGRAM_DEBUG", "False")
)

package main

import "sync"

var (
	// lock for inline query task
	inlineQueryWG = sync.WaitGroup{}
	// couter for inline tasks
	inlineTasksCount = 0
	// maximum inline query task
	maxInlineTaskCount = 10
)

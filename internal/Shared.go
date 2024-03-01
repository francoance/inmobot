package internal

import (
	"context"
	"os"
	"os/signal"
)

var Context context.Context
var InitialScrapeNeeded = false

func InitShared() {
	newCtx, _ := signal.NotifyContext(context.Background(), os.Interrupt)
	Context = newCtx

	if DatabaseExists() {
		InitialScrapeNeeded = false
	}
}

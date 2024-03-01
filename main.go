package main

import (
	"inmobot/internal"
	"time"
)

func main() {
	initDeps()
	internal.SendMessage("Arranqué :)")

	if internal.InitialScrapeNeeded {
		internal.SendMessage("Scrape inicial...")
		internal.ScrapePage()
		internal.InitialScrapeNeeded = false
	}

	for {
		internal.ScrapePage()
		time.Sleep(time.Duration(internal.Configuration.ScrapPeriodInMinutes) * time.Minute)
	}
}

func initDeps() {
	internal.InitConfig()
	internal.InitShared()
	internal.InitTelegramBot()
	internal.InitDatabase()
	internal.InitHtmlParser()
}

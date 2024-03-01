package internal

import (
	"github.com/gocolly/colly"
	"strings"
	"time"
)

var scraper *colly.Collector
var retries int64 = 0

func InitHtmlParser() {
	scraper = colly.NewCollector(
		colly.AllowedDomains("inmoclick.com.ar"),
		colly.AllowURLRevisit(),
	)
	scraper.SetRequestTimeout(time.Duration(Configuration.RequestTimeoutInSeconds) * time.Second)
}

func ScrapePage() {
	scraper.OnHTML("article", func(e *colly.HTMLElement) {
		go parseInm(e)
	})

	err := scraper.Visit(Configuration.Url)
	if err != nil {
		retries++
		time.Sleep(time.Second * 30 * time.Duration(Configuration.RequestTimeoutInSeconds))
		go ScrapePage()
	} else {
		retries = 0
	}
}

func parseInm(inm *colly.HTMLElement) {
	kid := inm.Attr("kid")
	precio := inm.Attr("precio")
	supTotal := inm.Attr("sup_t")
	supCub := inm.Attr("sup_c")
	foto := inm.ChildAttr("div.property-image a.cont-photo img", "data-defer-src")
	habitaciones := inm.ChildText("div.property-tags div div.wi-dormitorio")
	banos := inm.ChildText("div.property-tags div div.wi-banio")
	url := inm.ChildAttr("div.property-image a.cont-photo", "href")
	direccion := strings.Replace(inm.ChildText("div.property-data"), "\n", "", -1)
	handleInmueble(Inmueble{
		kid,
		precio,
		supTotal,
		supCub,
		foto,
		habitaciones,
		banos,
		url,
		direccion,
	})
}

func handleInmueble(inm Inmueble) {
	if !InmuebleExists(inm.kid) {
		SaveInmueble(inm)
		if !InitialScrapeNeeded {
			NotifyInmueble(inm)
		}
	}
}

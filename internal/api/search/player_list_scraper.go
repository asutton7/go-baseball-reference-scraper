package search

import (
	"time"

	"github.com/gocolly/colly"
	"github.com/patrickmn/go-cache"
)

type PlayerListScraper struct {
	Cache cache.Cache
}

func (p PlayerListScraper) ScrapeAndCache(url string, cacheKey string) []PlayerProfile {
	c := colly.NewCollector()

	scraperResults := []PlayerProfile{}
	c.OnHTML("#div_players_ p", func(h *colly.HTMLElement) {
		link := h.ChildAttr("a", "href")
		name := h.ChildText("a")
		years := h.Text[len(h.Text)-11:]

		scraperResults = append(scraperResults, PlayerProfile{name, years, link})
	})

	c.Visit(url)
	p.Cache.Add(cacheKey, scraperResults, 24*time.Hour)
	return scraperResults
}

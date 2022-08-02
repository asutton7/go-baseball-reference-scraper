package search

import "github.com/gocolly/colly"

// TODO: Refactor to be a struct that maintains cache
func ScrapePlayerList(url string) []PlayerProfile {
	c := colly.NewCollector()

	scraperResults := []PlayerProfile{}
	c.OnHTML("#div_players_ p", func(h *colly.HTMLElement) {
		link := h.ChildAttr("a", "href")
		name := h.ChildText("a")
		years := h.Text[len(h.Text)-11:]

		scraperResults = append(scraperResults, PlayerProfile{name, years, link})
	})

	c.Visit(url)
	return scraperResults
}

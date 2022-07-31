package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	cols := []string{}

	yearlyStats := []map[string]string{}
	c.OnHTML("#batting_standard thead tr th", func(h *colly.HTMLElement) {
		cols = append(cols, h.Text)
	})

	c.OnHTML("#batting_standard tbody tr", func(h *colly.HTMLElement) {
		statHeaders := cols[1:]
		yearlyStats = append(yearlyStats, make(map[string]string))
		currentPos := len(yearlyStats) - 1
		year := h.ChildText("th")

		yearlyStats[currentPos]["Year"] = year
		h.ForEach("td", func(i int, h *colly.HTMLElement) {
			yearlyStats[currentPos][statHeaders[i]] = h.Text
		})
	})

	c.Visit("https://www.baseball-reference.com/players/w/willite01.shtml")

	for _, yearStats := range yearlyStats {
		fmt.Println(yearStats)
	}
}

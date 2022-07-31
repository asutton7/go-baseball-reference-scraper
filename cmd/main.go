package main

import (
	s "baseball_scraper/internal/api/scrape_player"
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	yearlyStats := s.ScrapePlayer(c, "https://www.baseball-reference.com/players/w/willite01.shtml")

	for _, yearStats := range yearlyStats {
		fmt.Println(yearStats)
	}
}

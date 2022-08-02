package search

import (
	"fmt"
	"strings"
)

type PlayerProfile struct {
	Name        string
	YearsPlayed string
	Link        string
}

func Search(lastName string, baseUrl string, p PlayerListScraper) []PlayerProfile {
	firstChar := strings.ToLower(string(lastName[0]))

	cachedResult, found := p.Cache.Get(firstChar)
	fmt.Print("Was cached?")
	fmt.Print(found)
	if !found {
		cachedResult = p.ScrapeAndCache(baseUrl+"/"+firstChar, firstChar)
	}

	searchResults := []PlayerProfile{}
	for _, player := range cachedResult.([]PlayerProfile) {
		if strings.Contains(strings.ToLower(player.Name), strings.ToLower(lastName)) {
			searchResults = append(searchResults, player)
		}
	}

	return searchResults
}

package search

import (
	"strings"

	"github.com/patrickmn/go-cache"
)

func search(lastName string, collection []PlayerProfile) []PlayerProfile {
	searchResults := []PlayerProfile{}
	for _, player := range collection {
		if strings.Contains(strings.ToLower(player.Name), strings.ToLower(lastName)) {
			searchResults = append(searchResults, player)
		}
	}

	return searchResults
}

func HandleSearch(c cache.Cache, url, lastName string) []PlayerProfile {
	firstChar := strings.ToLower(string(lastName[0]))
	scrapedResults, found := c.Get(firstChar)
	if !found {
		scrapedResults = Scrape(url)
		c.Add(firstChar, scrapedResults, cache.DefaultExpiration)
	}
	return search(lastName, scrapedResults.([]PlayerProfile))
}

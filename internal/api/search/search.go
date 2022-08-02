package search

import (
	"strings"
)

type PlayerProfile struct {
	Name        string
	YearsPlayed string
	Link        string
}

func Search(lastName string, baseUrl string, scrapePlayerList func(url string) []PlayerProfile) []PlayerProfile {
	firstChar := strings.ToLower(string(lastName[0]))

	// TODO: Cache scraper results
	players := scrapePlayerList(baseUrl + "/" + firstChar)

	searchResults := []PlayerProfile{}
	for _, player := range players {
		if strings.Contains(strings.ToLower(player.Name), strings.ToLower(lastName)) {
			searchResults = append(searchResults, player)
		}
	}

	return searchResults
}

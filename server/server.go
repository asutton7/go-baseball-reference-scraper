package configure_server

import (
	"baseball_scraper/internal/api/scrape_player"
	"baseball_scraper/internal/api/search"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

var scrapedPlayerCache = cache.New(24*time.Hour, 24*time.Hour)

func ScrapePlayerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	pathToPlayer := r.URL.Query()["path"][0]

	stats := scrape_player.ScrapePlayer("https://www.baseball-reference.com/" + pathToPlayer)
	json.NewEncoder(w).Encode(stats)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	lastName := r.URL.Query()["lastName"][0]
	firstChar := strings.ToLower(string(lastName[0]))
	searchResults := search.HandleSearch(*scrapedPlayerCache, "https://www.baseball-reference.com/players/"+firstChar, lastName)
	json.NewEncoder(w).Encode(searchResults)
}

func ConfigureServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/player", ScrapePlayerHandler)
	mux.HandleFunc("/search", SearchHandler)
	return mux
}

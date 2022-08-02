package configure_server

import (
	player "baseball_scraper/internal/api/scrape_player"
	search "baseball_scraper/internal/api/search"
	"encoding/json"
	"net/http"
)

func ScrapePlayerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	pathToPlayer := r.URL.Query()["path"][0]

	stats := player.ScrapePlayer("https://www.baseball-reference.com/" + pathToPlayer)
	json.NewEncoder(w).Encode(stats)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	lastName := r.URL.Query()["lastName"][0]
	players := search.Search(lastName, "https://www.baseball-reference.com/players", search.ScrapePlayerList)
	json.NewEncoder(w).Encode(players)
}

func ConfigureServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/player", ScrapePlayerHandler)
	mux.HandleFunc("/search", SearchHandler)
	return mux
}

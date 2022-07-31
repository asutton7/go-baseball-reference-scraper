package configure_server

import (
	s "baseball_scraper/internal/api/scrape_player"
	"encoding/json"
	"fmt"
	"net/http"
)

func ScrapePlayerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var requestJson struct{ Url string }

	err := json.NewDecoder(r.Body).Decode(&requestJson)

	if err != nil {
		panic(err)
	}
	stats := s.ScrapePlayer(requestJson.Url)
	fmt.Print(stats)
	json.NewEncoder(w).Encode(stats)
}

func ConfigureServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/player", ScrapePlayerHandler)
	return mux
}

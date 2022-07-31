package main

import (
	s "baseball_scraper/internal/api/scrape_player"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/player", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r, w)
		w.Header().Set("Content-Type", "application/json")
		var requestJson struct{ Url string }

		err := json.NewDecoder(r.Body).Decode(&requestJson)

		if err != nil {
			panic(err)
		}
		stats := s.ScrapePlayer(requestJson.Url)
		fmt.Print(stats)
		json.NewEncoder(w).Encode(stats)
	})

	http.ListenAndServe(":8080", mux)
}

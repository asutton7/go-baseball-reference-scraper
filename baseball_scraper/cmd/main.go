package main

import (
	cs "baseball_scraper/server"
	"net/http"
)

func main() {
	mux := cs.ConfigureServer()
	http.ListenAndServe(":8080", mux)
}

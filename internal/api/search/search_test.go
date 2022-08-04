package search

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
)

func TestSearch(t *testing.T) {
	received := search("Willia", []PlayerProfile{
		{Name: "Ted Williams", Link: "link-to-ted", YearsPlayed: "(1939-1960)"},
		{Name: "Dontrelle Willis", Link: "link-to-dontrelle", YearsPlayed: "(2003-2011)"},
	})

	expected := []PlayerProfile{{Name: "Ted Williams", Link: "link-to-ted", YearsPlayed: "(1939-1960)"}}

	if !reflect.DeepEqual(received, expected) {
		t.Errorf("expected %v but received %v", expected, received)
	}
}

func TestHandlerSearch(t *testing.T) {
	t.Run("scrapes web page and returns search results", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<!DOCTYPE html>
			<html>
			<head>
			<title>Testing Baseball Scraper</title>
			</head>
			<body>
			<div id="div_players_">
			<p><a href="link-to-ted">Ted Williams</a> (1939-1960)</p>
			<p><a href="link-to-dontrelle">Donetrelle Willis</a> (2003-2011)</p>
			</div>
			</body>
			</html>`))
		}))

		expected := []PlayerProfile{{Name: "Ted Williams", Link: "link-to-ted", YearsPlayed: "(1939-1960)"}}

		received := HandleSearch(*cache.New(time.Second, time.Second), server.URL, "Willia")

		if !reflect.DeepEqual(received, expected) {
			t.Errorf("expected %v but received %v", expected, received)
		}
	})

	t.Run("gets scraped player list from cache and does not attempt to scrape", func(t *testing.T) {
		c := cache.New(time.Second*5, time.Second*5)
		c.Add("w", []PlayerProfile{
			{Name: "Ted Williams", Link: "link-to-ted", YearsPlayed: "(1939-1960)"},
			{Name: "Dontrelle Willis", Link: "link-to-dontrelle", YearsPlayed: "(2003-2011)"},
		}, cache.DefaultExpiration)

		requestCount := 0
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestCount++
		}))

		expected := []PlayerProfile{{Name: "Ted Williams", Link: "link-to-ted", YearsPlayed: "(1939-1960)"}}

		received := HandleSearch(*c, server.URL, "Willia")

		if requestCount != 0 {
			t.Fatal("Server received request but it should not have")
		}

		if !reflect.DeepEqual(received, expected) {
			t.Errorf("expected %v but received %v", expected, received)
		}
	})
}

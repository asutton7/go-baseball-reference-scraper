package scrape_player

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gocolly/colly"
)

func TestScrapePlayer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!DOCTYPE html>
		<html>
		<head>
		<title>Testing Baseball Scraper</title>
		</head>
		<body>
		<table id="batting_standard">
		<thead>
		<tr>
		<th>year</th>
		<th>avg</th>
		</tr>
		</thead>
		<tbody>
		<tr>
		<th scope="row">2022</th>
		<td>.250</td>
		</tr>
		</tbody
		</table>
		</body>
		</html>`))
	}))
	received := ScrapePlayer(colly.NewCollector(), server.URL)

	expected := []map[string]string{{"year": "2022", "avg": ".250"}}

	if !reflect.DeepEqual(received, expected) {
		t.Errorf("expected %v but received %v", expected, received)
	}
}

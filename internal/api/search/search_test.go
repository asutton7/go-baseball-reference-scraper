package search

import (
	"reflect"
	"testing"
)

func playerScrapeMock(url string) []PlayerProfile {
	return []PlayerProfile{
		{Name: "Ted Williams", Link: "link-to-ted", YearsPlayed: "(1939 - 1960)"},
		{Name: "Dontrelle Willis", Link: "link-to-dontrelle", YearsPlayed: "(2003 - 2011)"},
	}
}

func TestSearch(t *testing.T) {
	received := Search("Willia", "baseurl", playerScrapeMock)

	expected := []PlayerProfile{{Name: "Ted Williams", Link: "link-to-ted", YearsPlayed: "(1939 - 1960)"}}

	if !reflect.DeepEqual(received, expected) {
		t.Errorf("expected %v but received %v", expected, received)
	}
}

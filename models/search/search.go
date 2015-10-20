package search

// Searcher defines a search
type Searcher interface {
	Perform()
}

// Search
type search struct {
	text     string
	channel  chan string
	category string
}

// constructor
// newSearch creates a new search
func newSearch(text string, channel chan string) *search {
	return &search{
		text:    text,
		channel: channel,
	}
}

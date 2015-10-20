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

func (search *search) likeText() string {
	return "%" + search.text + "%"
}

// constructor
// newSearch creates a new search
func newSearch(text string, channel chan string) *search {
	return &search{
		text:    text,
		channel: channel,
	}
}

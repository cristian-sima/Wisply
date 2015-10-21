package search

// ResultItem is one result from a search
// It has a url to go to
type ResultItem struct {
	Title       string
	URL         string
	Description string
	Icon        string
}

// Result encapsulates a list of results for a category
type Result struct {
	Results []ResultItem `json:"results"`
}

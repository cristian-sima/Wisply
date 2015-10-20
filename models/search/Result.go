package search

// ResultItem is one result from a search
// It has a url to go to
type ResultItem struct {
	Title string
	URL   string
	Data  interface{}
}

// Result encapsulates a list of results for a category
type Result struct {
	Results  []ResultItem `json:"results"`
	Category string
}

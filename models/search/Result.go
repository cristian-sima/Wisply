package search

// Result represents an item that makes the query
type Result struct {
	ID          int
	Title       string
	URL         string
	Description string
	Icon        string
	Category    string
	Data        interface{}
}

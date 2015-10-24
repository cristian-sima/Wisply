package search

// Response encapsulates a list of results for a query
type Response struct {
	Results []*Result `json:"results"`
	Query   string
	countID int
}

// AppendItem adds a new search result to the list
// It assigns an id
func (response *Response) AppendItem(result *Result) {
	result.ID = response.countID
	response.Results = append(response.Results, result)
	response.countID++
}

// NewResponse creates a reponse for the query
func NewResponse(query string) *Response {
	var results = []*Result{}
	return &Response{
		Results: results,
		Query:   query,
		countID: 1,
	}
}

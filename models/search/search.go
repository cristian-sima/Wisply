package search

import "github.com/cristian-sima/Wisply/models/database"

// Searcher defines the sets of methods which must be implemented by a search
type Searcher interface {
	Perform()
}

// Search
type search struct {
	query    string
	response *Response
	request  *Request
	options  database.SQLOptions
}

func (search *search) changeOptions(options database.SQLOptions) {
	search.options = options
}

func (search *search) likeQuery() string {
	return "%" + search.query + "%"
}

package search

import (
	"log"
	"time"

	"github.com/cristian-sima/Wisply/models/database"
)

// Request describes the operation which can be called to search something
type Request struct {
	Response  *Response
	query     string
	enquiries []Searcher
}

// SearchAnything searches for the query in all the
// - institutions
// It limits the number of results for each one to 5
func (request *Request) SearchAnything() {
	options, _ := database.NewSQLOptions(database.Temp{
		LimitMin: "0",
		Offset:   "5",
	})
	response := request.Response
	search := &search{
		query:    request.query,
		request:  request,
		response: response,
	}
	search.changeOptions(options)

	request.addEnquire(NewCurriculaSearch(search))
	request.addEnquire(NewInstitutionsSearch(search))
	request.addEnquire(NewRepositoriesSearch(search))
	request.addEnquire(NewCollectionsSearch(search))

	request.addEnquire(NewResourcesSearch(search))

	request.enquireData()
}

func (request *Request) addEnquire(enquire Searcher) {
	request.enquiries = append(request.enquiries, enquire)
}

func (request *Request) enquireData() {
	// TODO go channel
	for _, enquire := range request.enquiries {
		start := time.Now()

		enquire.Perform()

		elapsed := time.Since(start)
		log.Printf("Search has taken %s", elapsed)
	}
}

// NewRequest creates a search request
func NewRequest(query string) *Request {
	response := NewResponse(query)
	return &Request{
		query:    query,
		Response: response,
	}
}

package api

import "github.com/cristian-sima/Wisply/models/search"

// Search is the controller which manages the "search" operations
type Search struct {
	Controller
}

// SearchAnything searches for all
func (controller *Search) SearchAnything() {
	query := controller.Ctx.Input.Param(":query")
	if controller.IsAccountConnected() {
		controller.saveNotAccessedSearch(query)
	}
	request, err := search.NewRequest(query)
	if err != nil {
		controller.sendEmptyArray()
	} else {
		request.SearchAnything()
		controller.deliverResults(request.Response)
	}
}

// JustSaveAccountQuery saves the token
func (controller *Search) sendEmptyArray() {
	controller.Ctx.Output.Json(make([]int, 0), false, false)
}

// JustSaveAccountQuery saves the token
func (controller *Search) JustSaveAccountQuery() {
	query := controller.Ctx.Input.Param(":query")
	controller.saveAcessedSearch(query)
	controller.Ctx.Output.Json(true, false, false)
}

// SaveAccountQuery saves a search query that was accessed
func (controller *Search) saveAcessedSearch(query string) {
	controller.Account.GetSearches().InsertAccessed(query)
}

// SaveAccountQuery saves a search query that was not accessed
func (controller *Search) saveNotAccessedSearch(query string) {
	controller.Account.GetSearches().InsertNotAccessed(query)
}

func (controller *Search) deliverResults(response *search.Response) {
	results := response.Results
	controller.Ctx.Output.Json(results, false, false)
}

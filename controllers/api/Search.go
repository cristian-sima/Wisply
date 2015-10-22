package api

import "github.com/cristian-sima/Wisply/models/search"

// Search is the controller which manages the "search" operations
type Search struct {
	Controller
}

// SearchAnything searches for all
func (controller *Search) SearchAnything() {
	query := controller.Ctx.Input.Param(":text")
	request := search.NewRequest(query)

	request.SearchAnything()

	controller.deliverResults(request.Response)
}

func (controller *Search) deliverResults(response *search.Response) {
	results := response.Results
	controller.Ctx.Output.Json(results, false, false)
}

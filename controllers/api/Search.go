package api

import "github.com/cristian-sima/Wisply/models/search"

// Search is the controller which manages the "search" operations
type Search struct {
	Controller
}

// SearchAnything searches for all
func (controller *Search) SearchAnything() {
	text := controller.Ctx.Input.Param(":text")
	institutionSearch := search.NewInstitutionsSearch(text)
	results := institutionSearch.Perform()
	controller.Ctx.Output.Json(results.Results, false, false)
}

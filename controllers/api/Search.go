package api

import (
	"fmt"

	"github.com/cristian-sima/Wisply/models/search"
)

// Search is the controller which manages the "search" operations
type Search struct {
	Controller
	results []search.ResultItem
}

// Prepare allocates memory for result array
func (controller *Search) Prepare() {
	controller.results = make([]search.ResultItem, 0)
	//controller.Prepare()
}

// SearchAnything searches for all
func (controller *Search) SearchAnything() {
	text := controller.Ctx.Input.Param(":text")
	controller.getInstitutions(text)
	controller.getRepositories(text)
	controller.deliverResults()
}

func (controller *Search) getInstitutions(text string) {
	institutionsSearch := search.NewInstitutionsSearch(text)
	query := institutionsSearch.Perform()
	controller.add(query.Results)
}

func (controller *Search) getRepositories(text string) {
	repositoriesSearch := search.NewRepositoriesSearch(text)
	query := repositoriesSearch.Perform()
	controller.add(query.Results)
}

func (controller *Search) add(newResults []search.ResultItem) {
	controller.results = append(controller.results, newResults...)
}

func (controller *Search) deliverResults() {
	results := controller.results
	fmt.Println(results)
	controller.Ctx.Output.Json(results, false, false)
}

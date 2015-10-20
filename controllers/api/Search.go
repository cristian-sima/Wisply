package api

// Search is the controller which manages the "search" operations
type Search struct {
	Controller
}

// SearchAnything searches for all
func (controller *Search) SearchAnything() {
	controller.Ctx.Output.Json([]string{}, false, false)
}

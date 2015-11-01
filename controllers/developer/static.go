package developer

// Home shows the public HTML page for the API
type Home struct {
	Controller
}

// Display shows the static home page for API
func (controller *Home) Display() {
	controller.SetCustomTitle("Developers & Research")
	// Please use http://www.timestampgenerator.com/ for generating the timestamp
	controller.IndicateLastModification(1441987477)
	controller.LoadTemplate("home")
}

package harvest

// Home represents manages the home page for log area
type Home struct {
	Controller
}

// Display shows the home page for the log
func (controller *Home) Display() {
	controller.SetCustomTitle("Admin - Harvest Processes")
	controller.LoadTemplate("home")
}

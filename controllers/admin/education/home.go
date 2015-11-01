package education

// Home represents manages the home page for education
type Home struct {
	Controller
}

// Display shows the home page for the education
func (controller *Home) Display() {
	controller.SetCustomTitle("Admin - Education")
	controller.LoadTemplate("education")
}

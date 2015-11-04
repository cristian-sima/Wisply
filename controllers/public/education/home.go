package education

// Home manages the operations for education
type Home struct {
	Controller
}

// Display displays all the subjects
func (controller *Home) Display() {
	controller.SetCustomTitle("Wisply - Subject")
	controller.LoadTemplate("home")
}

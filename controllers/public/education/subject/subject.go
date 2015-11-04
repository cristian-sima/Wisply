package subject

// Subject manages the operations for curriculum
type Subject struct {
	Controller
}

// Display shows the public page for a subject of study
func (controller *Subject) Display() {
	controller.SetCustomTitle(controller.GetSubject().GetName())
	controller.LoadTemplate("home")
}

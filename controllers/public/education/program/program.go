package program

// Program manages the operations for curriculum
type Program struct {
	Controller
}

// Display shows the public page for a program of study
func (controller *Program) Display() {
	controller.SetCustomTitle(controller.GetProgram().GetName())
	controller.LoadTemplate("home")
}

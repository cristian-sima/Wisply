package program

// Program manages the operations for curriculum
type Program struct {
	Controller
}

// Display shows the dashboard for a program
func (controller *Program) Display() {
	controller.SetCustomTitle(controller.GetProgram().GetName())
	controller.TplNames = "site/public/curriculum/program/home.tpl"
}

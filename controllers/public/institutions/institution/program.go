package institution

import "github.com/cristian-sima/Wisply/models/repository"

// Program manages the operations with an program
type Program struct {
	Controller
	program *repository.Program
}

// Prepare loads the program
func (controller *Program) Prepare() {
	controller.Controller.Prepare()
	controller.loadProgram()
}

// GetProgram returns the reference to the program
func (controller *Program) GetProgram() *repository.Program {
	return controller.program
}

func (controller *Program) loadProgram() {
	ID := controller.Ctx.Input.Param(":program")
	program, err := repository.NewProgram(ID)
	if err == nil {
		controller.Data["program"] = program
		controller.program = program
	}
}

// Display shows the public page for a program
func (controller *Program) Display() {
	program := controller.GetProgram()
	controller.Data["modules"] = program.GetModules()
	controller.LoadTemplate("program")
}

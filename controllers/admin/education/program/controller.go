package curriculum

import (
	education "github.com/cristian-sima/Wisply/controllers/admin/education"
	model "github.com/cristian-sima/Wisply/models/education"
)

// Controller manages the operations for the controller
type Controller struct {
	education.Controller
	program *model.Program
}

// Prepare loads the program of study from the id of the request
func (controller *Controller) Prepare() {
	controller.loadProgram()
	controller.Controller.Prepare()
}

func (controller *Controller) loadProgram() {
	ID := controller.Ctx.Input.Param(":id")
	program, err := model.NewProgram(ID)
	if err != nil {
		controller.Abort("show-database-error")
	}
	controller.Data["program"] = program
	controller.program = program
	controller.SetCustomTitle("Admin - " + program.GetName())
}

package program

import (
	"github.com/cristian-sima/Wisply/controllers/public/education"
	model "github.com/cristian-sima/Wisply/models/education"
)

// Controller manages the operations with programs
type Controller struct {
	education.Controller
	program *model.Program
}

// Prepare loads the account
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("admin/accounts/account")
	controller.loadAccount()
}

// GetProgram returns the reference to the program
func (controller *Controller) GetProgram() *model.Program {
	return controller.program
}

func (controller *Controller) loadAccount() {
	ID := controller.Ctx.Input.Param(":program")
	program, err := model.NewProgram(ID)
	if err != nil {
		controller.Abort("show-database-error")
	}
	controller.Data["program"] = program
	controller.program = program
	controller.SetCustomTitle(program.GetName())
}

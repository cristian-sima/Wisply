package institution

import "github.com/cristian-sima/Wisply/models/repository"

// Module manages the operations with an module
type Module struct {
	Controller
	module repository.Module
}

// Prepare loads the module
func (controller *Module) Prepare() {
	controller.Controller.Prepare()
	controller.loadModule()
}

// GetModule returns the reference to the module
func (controller *Module) GetModule() repository.Module {
	return controller.module
}

func (controller *Module) loadModule() {
	ID := controller.Ctx.Input.Param(":module")
	module, err := repository.NewModule(ID)
	if err == nil {
		controller.Data["module"] = module
		controller.module = module
	}
}

// Display shows the public page for a module
func (controller *Module) Display() {
	controller.LoadTemplate("module")
}

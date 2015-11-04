package institution

import (
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/repository"
)

// Module manages the operations with an module
type Module struct {
	Program
	module *repository.Module
}

// Prepare loads the module
func (controller *Module) Prepare() {
	controller.Program.Prepare()
	controller.loadModule()
}

// GetModule returns the reference to the module
func (controller *Module) GetModule() *repository.Module {
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

// ShowInsertForm shows the form to add an module
func (controller *Module) ShowInsertForm() {
	controller.SetCustomTitle("Add Module")
	controller.showAddForm()
}

// CreateModule inserts an module in the database
func (controller *Module) CreateModule() {
	program := controller.GetProgram()
	data := make(map[string]interface{})
	data["module-title"] = strings.TrimSpace(controller.GetString("module-title"))
	data["module-content"] = strings.TrimSpace(controller.GetString("module-content"))
	data["module-code"] = strings.TrimSpace(controller.GetString("module-code"))
	data["module-credits"] = strings.TrimSpace(controller.GetString("module-credits"))
	data["module-year"] = strings.TrimSpace(controller.GetString("module-year"))
	data["module-program"] = program.GetID()

	problems, err := repository.CreateModule(data)
	if err != nil {
		controller.DisplayError(problems)
	} else {
		message := "The module has been inserted."
		goTo := "/admin/institutions/" + strconv.Itoa(controller.GetInstitution().ID) + "/program/" + strconv.Itoa(controller.GetProgram().GetID())
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// ShowModifyForm shows the form to modify a module's details
func (controller *Module) ShowModifyForm() {
	controller.showForm("Modify", "Modify this module")
}

// Modify updates an module in the database
func (controller *Module) Modify() {
	institution := controller.GetInstitution()
	module := controller.GetModule()
	program := controller.GetProgram()
	data := make(map[string]interface{})

	data["module-title"] = strings.TrimSpace(controller.GetString("module-title"))
	data["module-content"] = strings.TrimSpace(controller.GetString("module-content"))
	data["module-code"] = strings.TrimSpace(controller.GetString("module-code"))
	data["module-credits"] = strings.TrimSpace(controller.GetString("module-credits"))
	data["module-year"] = strings.TrimSpace(controller.GetString("module-year"))
	data["module-program"] = program.GetID()

	problems, err := module.Modify(data)
	if err != nil {
		controller.DisplayError(problems)
	} else {
		message := "The module has been modified."
		goTo := "/admin/institutions/" + strconv.Itoa(institution.ID) + "/program/" + strconv.Itoa(program.GetID())
		controller.DisplaySuccessMessage(message, goTo)
	}

}

// Delete deletes the module specified by parameter id
func (controller *Module) Delete() {
	module := controller.GetModule()
	err := module.Delete()
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		message := "The module [" + module.GetCode() + "] has been deleted."
		goTo := "/admin/modules/" + strconv.Itoa(module.GetID())
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// ShowAdvanceOptions displays the page with further options
// For instance, further options may be modify or delete
func (controller *Module) ShowAdvanceOptions() {
	controller.SetCustomTitle("Admin - Module - Advance options")
	controller.LoadTemplate("advance-options")

}

func (controller *Module) showAddForm() {
	controller.showForm("Add", "Add a new module")
}

func (controller *Module) showForm(action string, legend string) {
	controller.GenerateXSRF()
	controller.Data["action"] = action
	controller.Data["legend"] = legend
	controller.Data["actionURL"] = ""
	controller.Data["actionType"] = "POST"
	controller.LoadTemplate("form-module")
}

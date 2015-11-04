package institution

import (
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/repository"
)

// Program manages the operations with an program
type Program struct {
	Controller
	program repository.Program
}

// Prepare loads the program
func (controller *Program) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("admin/institutions/institution")
	controller.loadProgram()
}

// GetProgram returns the reference to the program
func (controller *Program) GetProgram() repository.Program {
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

// Display shows the administrative page for an program
func (controller *Program) Display() {
	program := controller.GetProgram()
	controller.Data["modules"] = program.GetModules()
	controller.LoadTemplate("program")
}

// ShowInsertForm shows the form to add an program
func (controller *Program) ShowInsertForm() {
	controller.SetCustomTitle("Add Program")
	controller.showAddForm()
}

// ShowAddModuleForm shows the form for adding a module for the program
// It shows the modules which are not included already
func (controller *Program) ShowAddModuleForm() {
	controller.SetCustomTitle(controller.GetProgram().GetCode() + " - Add Module")
	list := []repository.Module{}
	allModules := controller.GetInstitution().GetModules()
	currentModules := controller.GetProgram().GetModules()
	for _, institutionModule := range allModules {
		exists := false
		for _, programModule := range currentModules {
			if programModule.GetID() == institutionModule.GetID() {
				exists = true
			}
		}
		if !exists {
			list = append(list, institutionModule)
		}
	}
	controller.Data["modulesToAdd"] = list
	controller.GenerateXSRF()
	controller.LoadTemplate("add-module-form")
}

// AddModule adds a module
func (controller *Program) AddModule() {
	institution := controller.GetInstitution()
	program := controller.GetProgram()
	moduleID := strings.TrimSpace(controller.GetString("module-id"))
	err := program.AddModule(moduleID)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The program has been added."
		goTo := "/admin/institutions/" + strconv.Itoa(institution.ID) + "/program" + "/" + strconv.Itoa(program.GetID())
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// DeleteModule removes a module from a program of study
func (controller *Program) DeleteModule() {
	moduleID := controller.Ctx.Input.Param(":module")
	program := controller.GetProgram()
	err := program.DeleteModule(moduleID)
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		message := "The module has been deleted from the list of program " + program.GetTitle()
		goTo := "/admin/institutions/" + strconv.Itoa(program.GetID()) + "#programs"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// CreateProgram inserts an program in the database
func (controller *Program) CreateProgram() {
	institution := controller.GetInstitution()
	data := make(map[string]interface{})
	data["program-title"] = strings.TrimSpace(controller.GetString("program-title"))
	data["program-code"] = strings.TrimSpace(controller.GetString("program-code"))
	data["program-year"] = strings.TrimSpace(controller.GetString("program-year"))
	data["program-ucas-code"] = strings.TrimSpace(controller.GetString("program-ucas-code"))
	data["program-level"] = strings.TrimSpace(controller.GetString("program-level"))
	data["program-content"] = strings.TrimSpace(controller.GetString("program-content"))
	data["program-subject"] = strings.TrimSpace(controller.GetString("program-subject"))
	data["program-institution"] = institution.ID

	problems, err := repository.CreateProgram(data)
	if err != nil {
		controller.DisplayError(problems)
	} else {
		message := "The program has been inserted."
		goTo := "/admin/institutions/" + strconv.Itoa(institution.ID) + "#programs"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// ShowModifyForm shows the form to modify a program's details
func (controller *Program) ShowModifyForm() {
	controller.showForm("Modify", "Modify this program")
}

// Modify updates an program in the database
func (controller *Program) Modify() {
	program := controller.GetProgram()
	institution := controller.GetInstitution()
	data := make(map[string]interface{})
	data["program-title"] = strings.TrimSpace(controller.GetString("program-title"))
	data["program-code"] = strings.TrimSpace(controller.GetString("program-code"))
	data["program-year"] = strings.TrimSpace(controller.GetString("program-year"))
	data["program-ucas-code"] = strings.TrimSpace(controller.GetString("program-ucas-code"))
	data["program-level"] = strings.TrimSpace(controller.GetString("program-level"))
	data["program-content"] = strings.TrimSpace(controller.GetString("program-content"))
	data["program-subject"] = strings.TrimSpace(controller.GetString("program-subject"))
	data["program-institution"] = institution.ID

	problems, err := program.Modify(data)
	if err != nil {
		controller.DisplayError(problems)
	} else {
		message := "The program has been modified!"
		goTo := "/admin/institutions/" + strconv.Itoa(institution.ID) + "#programs"
		controller.DisplaySuccessMessage(message, goTo)
	}

}

// Delete deletes the program specified by parameter id
func (controller *Program) Delete() {
	program := controller.GetProgram()
	err := program.Delete()
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		message := "The program [" + program.GetCode() + "] has been deleted."
		goTo := "/admin/institutions/" + strconv.Itoa(program.GetID()) + "#programs"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// ShowAdvanceOptions displays the page with further options
// For instance, further options may be modify or delete
func (controller *Program) ShowAdvanceOptions() {
	controller.SetCustomTitle("Admin - Program - Advance options")
	controller.LoadTemplate("advance-options")

}

func (controller *Program) showAddForm() {
	controller.showForm("Add", "Add a new program")
}

func (controller *Program) showForm(action string, legend string) {
	controller.GenerateXSRF()
	controller.Data["action"] = action
	controller.Data["legend"] = legend
	controller.Data["actionURL"] = ""
	controller.Data["actionType"] = "POST"
	controller.LoadTemplate("form-program")
}

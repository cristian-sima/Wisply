package admin

import (
	"strings"

	"github.com/cristian-sima/Wisply/models/curriculum"
)

// Curriculum manages the operations for curriculum
type Curriculum struct {
	Controller
}

// ShowHomePage shows all the repositories
func (controller *Curriculum) ShowHomePage() {
	list := curriculum.GetAllPrograms()
	controller.Data["programs"] = list
	controller.SetCustomTitle("Admin - Curriculum")
	controller.TplNames = "site/admin/curriculum/list.tpl"
}

// ShowProgramAdvanceOptions shows the panel with the advance options for the
// program
func (controller *Curriculum) ShowProgramAdvanceOptions() {
	ID := controller.Ctx.Input.Param(":id")
	controller.loadProgramToTemplate(ID)
	controller.TplNames = "site/admin/curriculum/program/advance-options.tpl"
}

// ShowProgram shows the dashboard for a program
func (controller *Curriculum) ShowProgram() {
	ID := controller.Ctx.Input.Param(":id")
	controller.loadProgramToTemplate(ID)
	controller.TplNames = "site/admin/curriculum/program/home.tpl"
}

// ShowProgram shows the dashboard for a program
func (controller *Curriculum) loadProgramToTemplate(ID string) *curriculum.Program {
	program, err := curriculum.NewProgram(ID)
	if err != nil {
		controller.Abort("databaseError")
		return program
	} else {
		controller.Data["program"] = program
		// controller.Data["definitions"] = program.GetDefinitions()
		controller.SetCustomTitle("Admin - " + program.GetName())
		return program
	}
}

// ShowAddProgramForm shows the page with the form to add a program
func (controller *Curriculum) ShowAddProgramForm() {
	controller.showForm("Add")
}

// CreateProgram creates a new program
func (controller *Curriculum) CreateProgram() {
	name := strings.TrimSpace(controller.GetString("program-name"))
	err := curriculum.CreateProgram(name)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The program has been created!"
		goTo := "/admin/curriculum/"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

func (controller *Curriculum) showForm(action string) {
	controller.GenerateXSRF()
	controller.Data["action"] = action
	controller.TplNames = "site/admin/curriculum/form.tpl"
}

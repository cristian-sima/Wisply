package admin

import (
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/auth"
	"github.com/cristian-sima/Wisply/models/curriculum"
)

// Curriculum manages the operations for curriculum
type Curriculum struct {
	Controller
}

// ShowHomePage shows all the repositories
func (controller *Curriculum) ShowHomePage() {
	controller.SetCustomTitle("Admin - Curriculum")
	controller.TplNames = "site/admin/curriculum/list.tpl"
}

// ShowProgramAdvanceOptions shows the panel with the advance options for the
// program
func (controller *Curriculum) ShowProgramAdvanceOptions() {
	controller.GenerateXSRF()
	controller.loadProgramToTemplate()
	controller.TplNames = "site/admin/curriculum/program/advance-options.tpl"
}

// ShowModifyDescriptionForm displays the form to modify the static description
func (controller *Curriculum) ShowModifyDescriptionForm() {
	controller.GenerateXSRF()
	program := controller.loadProgramToTemplate()
	controller.Data["description"] = program.GetDescription()
	controller.TplNames = "site/admin/curriculum/program/form-description.tpl"
}

// UpdateProgramDescription changes the static description of the program
func (controller *Curriculum) UpdateProgramDescription() {
	program := controller.loadProgramToTemplate()
	description := strings.TrimSpace(controller.GetString("program-description"))
	err := program.SetDescription(description)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The description has been changed!"
		goTo := "/admin/curriculum/programs/" + strconv.Itoa(program.GetID()) + "/advance-options"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// ShowProgram shows the dashboard for a program
func (controller *Curriculum) ShowProgram() {
	controller.loadProgramToTemplate()
	controller.TplNames = "site/admin/curriculum/program/home.tpl"
}

// ShowAddProgramForm shows the page with the form to add a program
func (controller *Curriculum) ShowAddProgramForm() {
	controller.showForm("Add")
}

// ShowModifyProgramForm shows the page with the form modify the program
func (controller *Curriculum) ShowModifyProgramForm() {
	controller.loadProgramToTemplate()
	controller.showForm("Modify")
}

// UpdateProgram updates the details of the program
func (controller *Curriculum) UpdateProgram() {
	program := controller.loadProgramToTemplate()
	details := make(map[string]interface{})
	details["name"] = strings.TrimSpace(controller.GetString("program-name"))
	err := program.Modify(details)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The program has been updated!"
		goTo := "/admin/curriculum/programs/" + strconv.Itoa(program.GetID()) + "/advance-options"
		controller.DisplaySuccessMessage(message, goTo)
	}
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

// DeleteProgram deletes the entire program and data related to it
// It requires the admin password
func (controller *Curriculum) DeleteProgram() {
	password := strings.TrimSpace(controller.GetString("password"))
	isPasswordValid := auth.VerifyAccount(controller.Account, password)
	if isPasswordValid {
		program := controller.loadProgramToTemplate()
		err := program.Delete()
		if err != nil {
			controller.DisplaySimpleError(err.Error())
		} else {
			message := "Wisply deleted all the information related to [" + program.GetName() + "] !"
			goTo := "/admin/curriculum/"
			controller.DisplaySuccessMessage(message, goTo)
		}
	} else {
		controller.Redirect("/admin/curriculum", 404)
	}
}

func (controller *Curriculum) loadProgramToTemplate() *curriculum.Program {
	ID := controller.Ctx.Input.Param(":id")
	program, err := curriculum.NewProgram(ID)
	if err != nil {
		controller.Abort("databaseError")
		return program
	}
	controller.Data["program"] = program
	controller.SetCustomTitle("Admin - " + program.GetName())
	return program
}

func (controller *Curriculum) showForm(action string) {
	controller.GenerateXSRF()
	controller.Data["action"] = action
	controller.TplNames = "site/admin/curriculum/form.tpl"
}

package curriculum

import (
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/auth"
	model "github.com/cristian-sima/Wisply/models/education"
)

// Program is the controller which manages the operations for the program page
type Program struct {
	controller
}

// Display shows the dashboard for a program
func (controller *Program) Display() {
	controller.SetCustomTitle("Admin - " + controller.program.GetName())
	controller.LoadTemplate("home.tpl")
}

// ShowAdvanceOptions shows the page with the advance options for the program
func (controller *Program) ShowAdvanceOptions() {
	controller.GenerateXSRF()
	controller.TplNames = "site/admin/curriculum/program/advance-options.tpl"
}

// ShowModifyForm displays the form to modify the static description
func (controller *Program) ShowModifyForm() {
	controller.GenerateXSRF()
	controller.Data["description"] = controller.program.GetDescription()
	controller.TplNames = "site/admin/curriculum/program/form-description.tpl"
	controller.showForm("Modify")
}

// UpdateDescription modifies the static description
func (controller *Program) UpdateDescription() {
	description := strings.TrimSpace(controller.GetString("program-description"))
	err := controller.program.SetDescription(description)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The description has been modified."
		goTo := "/admin/curriculum/programs/" + strconv.Itoa(controller.program.GetID()) + "/advance-options"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// ShowAddProgramForm shows the page with the form to add a program
func (controller *Program) ShowAddProgramForm() {
	controller.showForm("Add")
}

// Update updates the details of the program
func (controller *Program) Update() {
	details := make(map[string]interface{})
	details["name"] = strings.TrimSpace(controller.GetString("program-name"))
	err := controller.program.Modify(details)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The program has been updated!"
		goTo := "/admin/curriculum/programs/" + strconv.Itoa(controller.program.GetID()) + "/advance-options"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// CreateProgram creates a new program
func (controller *Program) CreateProgram() {
	name := strings.TrimSpace(controller.GetString("program-name"))
	err := model.CreateProgram(name)
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
func (controller *Program) DeleteProgram() {
	password := strings.TrimSpace(controller.GetString("password"))
	isPasswordValid := auth.VerifyAccount(controller.Account, password)
	if isPasswordValid {
		err := controller.program.Delete()
		if err != nil {
			controller.DisplaySimpleError(err.Error())
		} else {
			message := "Wisply deleted all the information related to [" + controller.program.GetName() + "] !"
			goTo := "/admin/curriculum/"
			controller.DisplaySuccessMessage(message, goTo)
		}
	} else {
		controller.Redirect("/admin/curriculum", 404)
	}
}

func (controller *Program) showForm(action string) {
	controller.GenerateXSRF()
	controller.Data["action"] = action
	controller.TplNames = "site/admin/curriculum/form.tpl"
}

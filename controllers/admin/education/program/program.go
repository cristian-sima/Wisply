package program

import (
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/auth"
	model "github.com/cristian-sima/Wisply/models/education"
)

// Program is the controller which manages the operations for the program page
type Program struct {
	Controller
}

// Display shows the dashboard for a program
func (controller *Program) Display() {
	program := controller.program
	controller.SetCustomTitle("Admin - " + program.GetName())
	controller.LoadTemplate("home")
	controller.Data["definitions"] = program.GetDefinitions()
	controller.Data["KAs"] = program.GetKAs()
}

// ShowAdvanceOptions shows the page with the advance options for the program
func (controller *Program) ShowAdvanceOptions() {
	controller.GenerateXSRF()
	controller.LoadTemplate("advance-options")
}

// ShowModifyForm displays the form to modify the program details
func (controller *Program) ShowModifyForm() {
	controller.GenerateXSRF()
	controller.showForm("Modify")
}

// ShowModifyDescription displays the form to modify the static description
func (controller *Program) ShowModifyDescription() {
	controller.GenerateXSRF()
	controller.Data["description"] = controller.program.GetDescription()
	controller.LoadTemplate("form-description")
}

// UpdateDescription modifies the static description
func (controller *Program) UpdateDescription() {
	description := strings.TrimSpace(controller.GetString("program-description"))
	err := controller.program.SetDescription(description)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The description has been modified."
		goTo := "/admin/education/programs/" + strconv.Itoa(controller.program.GetID()) + "/advance-options"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// ShowAddForm shows the page with the form to add a program
func (controller *Program) ShowAddForm() {
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
		goTo := "/admin/education/programs/" + strconv.Itoa(controller.program.GetID()) + "/advance-options"
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
		goTo := "/admin/education/"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// Delete deletes the entire program and data related to it
// It requires the admin password
func (controller *Program) Delete() {
	password := strings.TrimSpace(controller.GetString("password"))
	isPasswordValid := auth.VerifyAccount(controller.Account, password)
	if isPasswordValid {
		err := controller.program.Delete()
		if err != nil {
			controller.DisplaySimpleError(err.Error())
		} else {
			message := "Wisply deleted all the information related to [" + controller.program.GetName() + "] !"
			goTo := "/admin/education/"
			controller.DisplaySuccessMessage(message, goTo)
		}
	} else {
		controller.Redirect("/admin/education", 404)
	}
}

func (controller *Program) showForm(action string) {
	controller.GenerateXSRF()
	controller.Data["action"] = action
	controller.LoadTemplate("form")
}

package program

import (
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/education"
)

// Definition is the controller which manages the formal definitions for
// a program of study
type Definition struct {
	Controller
	definition *education.Definition
}

// Prepare loads the definition
func (controller *Definition) Prepare() {
	controller.Controller.Prepare()
	controller.loadDefinition()
}

func (controller *Definition) loadDefinition() {
	ID := controller.Ctx.Input.Param(":definition")
	definition, err := education.NewDefinition(ID)
	if err == nil {
		controller.Data["definition"] = definition
		controller.definition = definition
	}
}

// GetDefinition returns the current defintion
func (controller *Definition) GetDefinition() *education.Definition {
	return controller.definition
}

// ShowModifyForm displays the form to modify the static description
func (controller *Definition) ShowModifyForm() {
	controller.GenerateXSRF()
	controller.Data["description"] = controller.program.GetDescription()
	controller.LoadTemplate("form-description")
	controller.showForm("Modify")
}

// UpdateDescription modifies the static description
func (controller *Definition) UpdateDescription() {
	description := strings.TrimSpace(controller.GetString("program-description"))
	err := controller.program.SetDescription(description)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The description has been modified."
		goTo := "/admin/education/programs/" + strconv.Itoa(controller.program.GetID())
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// ShowAddForm shows the page with the form to add a program
func (controller *Definition) ShowAddForm() {
	controller.showForm("Add")
}

// Update updates the details of the program
func (controller *Definition) Update() {
	source := strings.TrimSpace(controller.GetString("definition-source"))
	content := strings.TrimSpace(controller.GetString("definition-content"))
	data := make(map[string]interface{})
	data["definition-content"] = content
	data["definition-source"] = source
	data["definition-program"] = controller.GetProgram().GetID()
	err := controller.GetDefinition().Modify(data)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The definition has been updated!"
		goTo := "/admin/education/programs/" + strconv.Itoa(controller.program.GetID())
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// CreateDefinition creates a new program
func (controller *Definition) CreateDefinition() {
	source := strings.TrimSpace(controller.GetString("definition-source"))
	content := strings.TrimSpace(controller.GetString("definition-content"))
	data := make(map[string]interface{})
	data["definition-content"] = content
	data["definition-source"] = source
	data["definition-program"] = controller.GetProgram().GetID()
	err := education.CreateDefinition(data)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The definition has been inserted."
		goTo := "/admin/education/programs/" + strconv.Itoa(controller.program.GetID())
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// Delete deletes the entire program and data related to it
// It requires the admin password
func (controller *Definition) Delete() {
	err := controller.GetDefinition().Delete()
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The definition has been deleted."
		goTo := "/admin/education/programs/" + strconv.Itoa(controller.GetProgram().GetID())
		controller.DisplaySuccessMessage(message, goTo)
	}
}

func (controller *Definition) showForm(action string) {
	controller.GenerateXSRF()
	controller.Data["action"] = action
	controller.LoadTemplate("form-definition")
}

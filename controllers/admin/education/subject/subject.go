package subject

import (
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/auth"
	model "github.com/cristian-sima/Wisply/models/education"
)

// Subject is the controller which manages the operations for the subject page
type Subject struct {
	Controller
}

// Display shows the dashboard for a subject
func (controller *Subject) Display() {
	subject := controller.subject
	controller.SetCustomTitle("Admin - " + subject.GetName())
	controller.LoadTemplate("home")
	controller.Data["definitions"] = subject.GetDefinitions()
	controller.Data["KAs"] = subject.GetKAs()
}

// ShowAdvanceOptions shows the page with the advance options for the subject
func (controller *Subject) ShowAdvanceOptions() {
	controller.GenerateXSRF()
	controller.LoadTemplate("advance-options")
}

// ShowModifyForm displays the form to modify the subject details
func (controller *Subject) ShowModifyForm() {
	controller.GenerateXSRF()
	controller.showForm("Modify")
}

// ShowModifyDescription displays the form to modify the static description
func (controller *Subject) ShowModifyDescription() {
	controller.GenerateXSRF()
	controller.Data["description"] = controller.subject.GetDescription()
	controller.LoadTemplate("form-description")
}

// UpdateDescription modifies the static description
func (controller *Subject) UpdateDescription() {
	description := strings.TrimSpace(controller.GetString("subject-description"))
	err := controller.subject.SetDescription(description)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The description has been modified."
		goTo := "/admin/education/subjects/" + strconv.Itoa(controller.subject.GetID()) + "/advance-options"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// ShowAddForm shows the page with the form to add a subject
func (controller *Subject) ShowAddForm() {
	controller.showForm("Add")
}

// Update updates the details of the subject
func (controller *Subject) Update() {
	details := make(map[string]interface{})
	details["name"] = strings.TrimSpace(controller.GetString("subject-name"))
	err := controller.subject.Modify(details)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The subject has been updated!"
		goTo := "/admin/education/subjects/" + strconv.Itoa(controller.subject.GetID()) + "/advance-options"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// CreateSubject creates a new subject
func (controller *Subject) CreateSubject() {
	name := strings.TrimSpace(controller.GetString("subject-name"))
	err := model.CreateSubject(name)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The subject has been created!"
		goTo := "/admin/education/"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// Delete deletes the entire subject and data related to it
// It requires the admin password
func (controller *Subject) Delete() {
	password := strings.TrimSpace(controller.GetString("password"))
	isPasswordValid := auth.VerifyAccount(controller.Account, password)
	if isPasswordValid {
		err := controller.subject.Delete()
		if err != nil {
			controller.DisplaySimpleError(err.Error())
		} else {
			message := "Wisply deleted all the information related to [" + controller.subject.GetName() + "] !"
			goTo := "/admin/education/"
			controller.DisplaySuccessMessage(message, goTo)
		}
	} else {
		controller.Redirect("/admin/education", 404)
	}
}

func (controller *Subject) showForm(action string) {
	controller.GenerateXSRF()
	controller.Data["action"] = action
	controller.LoadTemplate("form")
}

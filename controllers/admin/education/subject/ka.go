package subject

import (
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/education"
)

// KA is the controller for Knowledge Area
type KA struct {
	Controller
	ka *education.KA
}

// Prepare loads the ka
func (controller *KA) Prepare() {
	controller.Controller.Prepare()
	controller.loadKA()
}

func (controller *KA) loadKA() {
	ID := controller.Ctx.Input.Param(":ka")
	ka, err := education.NewKA(ID)
	if err == nil {
		controller.Data["ka"] = ka
		controller.ka = ka
	}
}

// GetKA returns the current defintion
func (controller *KA) GetKA() *education.KA {
	return controller.ka
}

// ShowModifyForm displays the form to modify the static description
func (controller *KA) ShowModifyForm() {
	controller.GenerateXSRF()
	controller.Data["description"] = controller.subject.GetDescription()
	controller.LoadTemplate("form-description")
	controller.showForm("Modify")
}

// ShowAddForm shows the page with the form to add a subject
func (controller *KA) ShowAddForm() {
	controller.showForm("Add")
}

// Update updates the details of the subject
func (controller *KA) Update() {
	source := strings.TrimSpace(controller.GetString("ka-source"))
	content := strings.TrimSpace(controller.GetString("ka-content"))
	code := strings.TrimSpace(controller.GetString("ka-code"))
	title := strings.TrimSpace(controller.GetString("ka-title"))
	data := make(map[string]interface{})
	data["ka-content"] = content
	data["ka-source"] = source
	data["ka-subject"] = controller.GetSubject().GetID()
	data["ka-code"] = code
	data["ka-title"] = title
	err := controller.GetKA().Modify(data)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The knwoledge area has been updated!"
		goTo := "/admin/education/subjects/" + strconv.Itoa(controller.subject.GetID())
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// CreateKA creates a new subject
func (controller *KA) CreateKA() {
	source := strings.TrimSpace(controller.GetString("ka-source"))
	content := strings.TrimSpace(controller.GetString("ka-content"))
	code := strings.TrimSpace(controller.GetString("ka-code"))
	title := strings.TrimSpace(controller.GetString("ka-title"))
	data := make(map[string]interface{})
	data["ka-content"] = content
	data["ka-source"] = source
	data["ka-subject"] = controller.GetSubject().GetID()
	data["ka-code"] = code
	data["ka-title"] = title
	err := education.CreateKA(data)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The knwoledge area has been inserted."
		goTo := "/admin/education/subjects/" + strconv.Itoa(controller.GetSubject().GetID())
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// Delete deletes the entire subject and data related to it
// It requires the admin password
func (controller *KA) Delete() {
	err := controller.GetKA().Delete()
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		message := "The knwoledge area has been deleted."
		goTo := "/admin/education/subjects/" + strconv.Itoa(controller.GetSubject().GetID())
		controller.DisplaySuccessMessage(message, goTo)
	}
}

func (controller *KA) showForm(action string) {
	controller.GenerateXSRF()
	controller.Data["action"] = action
	controller.LoadTemplate("form-ka")
}

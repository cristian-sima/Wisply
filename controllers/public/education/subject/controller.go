package subject

import (
	"github.com/cristian-sima/Wisply/controllers/public/education"
	model "github.com/cristian-sima/Wisply/models/education"
)

// Controller manages the operations with subjects
type Controller struct {
	education.Controller
	subject *model.Subject
}

// Prepare loads the subject
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("public/education/subject")
	controller.loadSubject()
}

// GetSubject returns the reference to the subject
func (controller *Controller) GetSubject() *model.Subject {
	return controller.subject
}

func (controller *Controller) loadSubject() {
	ID := controller.Ctx.Input.Param(":subject")
	subject, err := model.NewSubject(ID)
	if err != nil {
		controller.Abort("show-database-error")
	}
	controller.Data["subject"] = subject
	controller.subject = subject
	controller.SetCustomTitle(subject.GetName())
}

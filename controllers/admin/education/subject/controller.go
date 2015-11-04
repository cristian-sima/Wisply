package subject

import (
	education "github.com/cristian-sima/Wisply/controllers/admin/education"
	model "github.com/cristian-sima/Wisply/models/education"
)

// Controller manages the operations for the controller
type Controller struct {
	education.Controller
	subject *model.Subject
}

// Prepare loads the subject of study from the id of the request
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("admin/education/subjects")
	controller.loadSubject()
}

// GetSubject returns the subject of the controller
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
	controller.SetCustomTitle("Admin - " + subject.GetName())
}

package institution

import (
	"github.com/cristian-sima/Wisply/controllers/public/education"
	"github.com/cristian-sima/Wisply/models/repository"
)

// Controller manages the operations with an institution
type Controller struct {
	education.Controller
	institution *repository.Institution
}

// Prepare loads the institution
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("public/institutions/institution")
	controller.loadInstitution()
}

func (controller *Controller) loadInstitution() {
	ID := controller.Ctx.Input.Param(":institution")
	institution, err := repository.NewInstitution(ID)
	if err != nil {
		controller.Abort("show-database-error")
	}
	controller.Data["institution"] = institution
	controller.institution = institution
	controller.SetCustomTitle(institution.Name)
}

// GetInstitution returns the reference to the institution
func (controller *Controller) GetInstitution() *repository.Institution {
	return controller.institution
}

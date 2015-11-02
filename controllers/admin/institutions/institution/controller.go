package institution

import (
	"github.com/cristian-sima/Wisply/controllers/admin/institutions"
	"github.com/cristian-sima/Wisply/models/repository"
)

// Controller manages the operations with the institutions
type Controller struct {
	institutions.Controller
	institution *repository.Institution
}

// Prepare loads the institution
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("admin/institutions/institution")
	controller.loadInstitution()
}

// GetInstitution returns the reference to the institution
func (controller *Controller) GetInstitution() *repository.Institution {
	return controller.institution
}

func (controller *Controller) loadInstitution() {
	ID := controller.Ctx.Input.Param(":institution")
	institution, err := repository.NewInstitution(ID)
	if err == nil {
		controller.Data["institution"] = institution
		controller.institution = institution
		controller.SetCustomTitle("Admin - " + institution.Name)
	}
}

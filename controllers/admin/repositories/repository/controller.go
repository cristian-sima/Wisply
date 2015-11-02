package repository

import (
	"github.com/cristian-sima/Wisply/controllers/admin/repositories"
	"github.com/cristian-sima/Wisply/models/repository"
)

// Controller manages the operations with the repositories
type Controller struct {
	repositories.Controller
	repository *repository.Repository
}

// Prepare loads the repository
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("admin/repositories/repository")
	controller.loadRepository()
}

// GetRepository returns the reference to the repository
func (controller *Controller) GetRepository() *repository.Repository {
	return controller.repository
}

func (controller *Controller) loadRepository() {
	ID := controller.Ctx.Input.Param(":repository")
	repository, err := repository.NewRepository(ID)
	if err == nil {
		controller.Data["repository"] = repository
		controller.repository = repository
		controller.SetCustomTitle("Admin - " + repository.Name)
	}
}

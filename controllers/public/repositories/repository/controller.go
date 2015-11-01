package repository

import (
	"github.com/cristian-sima/Wisply/controllers/public/education"
	"github.com/cristian-sima/Wisply/models/repository"
)

// Controller manages the operations with an repository
type Controller struct {
	education.Controller
	repository *repository.Repository
}

// Prepare loads the repository
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("public/repositories/repository")
	controller.loadRepository()
}

func (controller *Controller) loadRepository() {
	ID := controller.Ctx.Input.Param(":repository")
	repository, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("show-database-error")
	}
	controller.Data["repository"] = repository
	controller.repository = repository
	controller.SetCustomTitle(repository.Name)
}

// GetRepository returns the reference to the repository
func (controller *Controller) GetRepository() *repository.Repository {
	return controller.repository
}

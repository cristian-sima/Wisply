package public

import (
	"encoding/json"

	"github.com/cristian-sima/Wisply/models/harvest"
	"github.com/cristian-sima/Wisply/models/repository"
	"github.com/cristian-sima/Wisply/models/wisply"
)

// RepositoryController managers the operations for displaying repositories
type RepositoryController struct {
	Controller
	model repository.Model
}

// List shows all the institutions
func (controller *RepositoryController) List() {
	var exists bool
	list := controller.model.GetAllInstitutions()
	exists = (len(list) != 0)
	controller.Data["anything"] = exists
	controller.Data["institutions"] = list
	controller.Data["host"] = controller.Ctx.Request.Host
	controller.TplNames = "site/public/institution/list.tpl"
}

// ShowRepository shows the details regarding a repository
func (controller *RepositoryController) ShowRepository() {
	ID := controller.Ctx.Input.Param(":id")
	repo, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		controller.Data["repository"] = repo
		controller.SetCustomTitle(repo.Name)

		controller.Data["institution"] = repo.GetInstitution()
		controller.Data["identification"] = repo.GetIdentification()

		if repo.HasBeenProcessed() {
			process := harvest.NewProcess(repo.LastProcess)

			storage := wisply.NewStorage(repo)

			collections := storage.GetCollections()

			controller.Data["collections"] = collections

			collectionsJSON, _ := json.Marshal(collections)
			controller.Data["collectionsJSON"] = string(collectionsJSON)

			controller.Data["process"] = process
			controller.IndicateLastModification(process.Process.End)
		}

		controller.TplNames = "site/public/repository/repository.tpl"
	}
}

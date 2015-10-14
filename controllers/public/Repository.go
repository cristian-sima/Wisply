package public

import (
	"github.com/cristian-sima/Wisply/models/harvest"
	"github.com/cristian-sima/Wisply/models/harvest/wisply"
	"github.com/cristian-sima/Wisply/models/repository"
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
	controller.Layout = "site/public-layout.tpl"
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
			controller.Data["collections"] = wisply.GetCollections(repo.ID)
			controller.Data["process"] = harvest.NewProcess(repo.LastProcess)
		}

		// controller.Data["records"] = wisply.GetRecords(rep.ID, database.SQLOptions{
		// 	Limit: "0, 15",
		// })
		controller.Layout = "site/public-layout.tpl"
		controller.TplNames = "site/public/repository/repository.tpl"
	}
}

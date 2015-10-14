package api

import (
	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/harvest"
	"github.com/cristian-sima/Wisply/models/wisply"
	"github.com/cristian-sima/Wisply/models/repository"
)

// Repository contains
type Repository struct {
	Controller
}

// GetResources returns the resources for the repository
func (controller *Repository) GetResources() {

	ID := controller.Ctx.Input.Param(":id")

	min := controller.Ctx.Input.Param(":min")
	max := controller.Ctx.Input.Param(":max")

	repo, err := repository.NewRepository(ID)

	if err != nil {
		controller.Abort("databaseError")
	} else {
		options, err := database.NewSQLOptions(database.Temp{
			LimitMin: min,
			LimitMax: max,
			Limit:    100,
		})

		if err != nil {
			controller.Abort("databaseError")
		} else {

			controller.Data["repository"] = repo
			controller.SetCustomTitle(repo.Name)

			controller.Data["institution"] = repo.GetInstitution()
			controller.Data["identification"] = repo.GetIdentification()

			controller.Data["records"] = wisply.GetRecords(repo.ID, options)

			if repo.HasBeenProcessed() {
				controller.Data["collections"] = wisply.GetCollections(repo.ID)
				controller.Data["process"] = harvest.NewProcess(repo.LastProcess)
			}

			controller.TplNames = "site/api/html/repository.tpl"
		}
	}
}

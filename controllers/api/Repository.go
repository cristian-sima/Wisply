package api

import (
	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/harvest"
	"github.com/cristian-sima/Wisply/models/repository"
	"github.com/cristian-sima/Wisply/models/wisply"
)

// Repository contains
type Repository struct {
	Controller
}

// GetResources returns the resources for the repository
func (controller *Repository) GetResources() {

	ID := controller.Ctx.Input.Param(":id")

	min := controller.Ctx.Input.Param(":min")
	offset := controller.Ctx.Input.Param(":number")

	repo, err := repository.NewRepository(ID)

	if err != nil {
		controller.Abort("databaseError")
	} else {
		options, err := database.NewSQLOptions(database.Temp{
			LimitMin: min,
			Offset:   offset,
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
			controller.TplNames = "site/api/repository/resources/html.tpl"
		}
	}
}

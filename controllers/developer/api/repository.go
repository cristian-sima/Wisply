package api

import (
	"strings"

	"github.com/cristian-sima/Wisply/models/database"
	"github.com/cristian-sima/Wisply/models/repository"
	"github.com/cristian-sima/Wisply/models/wisply"
)

// Repository contains
type Repository struct {
	Controller
}

// Prepare removes the layout
func (controller *Repository) Prepare() {
	controller.Controller.Prepare()
	controller.RemoveLayout()
	controller.SetTemplatePath("developer/api/repository/resource")
}

// GetResources returns the resources for the repository
func (controller *Repository) GetResources() {

	ID := controller.Ctx.Input.Param(":id")
	min := controller.Ctx.Input.Param(":min")
	offset := controller.Ctx.Input.Param(":number")
	repo, err := repository.NewRepository(ID)
	collection := strings.TrimSpace(controller.GetString("collection"))
	orderBy := strings.TrimSpace(controller.GetString("orderBy"))

	if err != nil {
		controller.Abort("show-database-error")
	} else {
		options, err := database.NewSQLOptions(database.Temp{
			LimitMin: min,
			Offset:   offset,
			Limit:    100,
			OrderBy:  orderBy,
			Where: map[string]string{
				"collection": collection,
			},
		})

		if err != nil {
			controller.Abort("show-database-error")
		} else {
			records := wisply.GetRecords(repo.ID, options)
			switch strings.TrimSpace(controller.GetString("format")) {
			case "html":
				controller.Data["records"] = records
				controller.LoadTemplate("html")
				break
			case "json":
				controller.Ctx.Output.Json(records, false, false)
				break
			default:
				controller.ShowBlankPage()
				break
			}
		}
	}
}

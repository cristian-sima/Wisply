package api

import (
	"encoding/json"
	"strings"

	"github.com/cristian-sima/Wisply/models/database"
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
	collection := strings.TrimSpace(controller.GetString("collection"))

	if err != nil {
		controller.Abort("databaseError")
	} else {
		options, err := database.NewSQLOptions(database.Temp{
			LimitMin: min,
			Offset:   offset,
			Limit:    100,
			Where: map[string]string{
				"collection": collection,
			},
		})

		if err != nil {
			controller.Abort("databaseError")
		} else {
			records := wisply.GetRecords(repo.ID, options)

			switch strings.TrimSpace(controller.GetString("format")) {
			case "html":
				controller.Data["records"] = records
				controller.TplNames = "site/api/repository/resources/html.tpl"
				break
			case "json":

				jsonRecords, _ := json.Marshal(struct {
					Records []*wisply.Record `json:"Records"`
				}{
					Records: records,
				})
				controller.Data["jsonRecords"] = jsonRecords
				controller.TplNames = "site/api/repository/resources/json.tpl"
				break
			default:
				controller.TplNames = "site/api/problem.tpl"
				break
			}
		}
	}
}

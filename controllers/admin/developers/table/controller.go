package table

import (
	"github.com/cristian-sima/Wisply/controllers/admin/accounts"
	model "github.com/cristian-sima/Wisply/models/developer"
)

// Controller manages the operations with the accounts
type Controller struct {
	accounts.Controller
	table *model.Table
}

// Prepare loads the account
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("admin/developers/data")
	controller.loadTable()
}

// GetTable returns the reference to the table
func (controller *Controller) GetTable() *model.Table {
	return controller.table
}

func (controller *Controller) loadTable() {
	ID := controller.Ctx.Input.Param(":table")
	table, err := model.NewTable(ID)
	if err == nil {
		controller.Data["table"] = table
		controller.table = table
		controller.SetCustomTitle("Admin - " + table.Name)
	}
}

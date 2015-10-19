package admin

import (
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/api"
)

// APIController manages the operation for API
type APIController struct {
	Controller
}

// RemoveAllowedTable removes the table from the list
func (controller *APIController) RemoveAllowedTable() {
	IDString := strings.TrimSpace(controller.GetString("table-id"))
	ID, _ := strconv.Atoi(IDString)
	err := api.RemoveAllowedTable(ID)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		controller.DisplaySuccessMessage("The table has been removed from the list!", "/admin/api")
	}
}

// InsertNewTable inserts the new table name into the list
func (controller *APIController) InsertNewTable() {
	tableName := strings.TrimSpace(controller.GetString("table-name"))
	if !api.IsAllowedTable(tableName) {
		controller.DisplaySimpleError("This table name is restricted.")
	} else {
		err := api.InsertNewTable(tableName)
		if err != nil {
			controller.DisplaySimpleError(err.Error())
		} else {
			controller.DisplaySuccessMessage("The table is now available to be downloaded!", "/admin/api")
		}
	}
}

// ShowAddForm shows the form to add a table to the download list
func (controller *APIController) ShowAddForm() {
	controller.GenerateXSRF()
	controller.SetCustomTitle("Admin - API - Add table")
	controller.Data["tables"] = api.GetWisplyTablesNamesNotAllowed()
	controller.Data["action"] = "Allow table to be downloaded"
	controller.TplNames = "site/admin/api/form.tpl"
}

// ShowHomePage displays the home page
func (controller *APIController) ShowHomePage() {
	controller.GenerateXSRF()
	controller.Data["tables"] = api.GetAllowedTables()
	controller.SetCustomTitle("Admin - API")
	controller.TplNames = "site/admin/api/home.tpl"
}

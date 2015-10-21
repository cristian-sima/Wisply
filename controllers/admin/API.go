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
		message := "The table has been removed from the list!"
		goTo := "/admin/api"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// InsertNewTable inserts the new table name into the list
func (controller *APIController) InsertNewTable() {
	name := strings.TrimSpace(controller.GetString("table-name"))
	description := strings.TrimSpace(controller.GetString("table-description"))
	table := api.Table{
		Name:        name,
		Description: description,
	}
	if api.AreValidDetails(table) {
		controller.DisplaySimpleError("This table name is restricted.")
	} else {
		err := api.InsertNewTable(table)
		if err != nil {
			controller.DisplaySimpleError(err.Error())
		} else {
			message := "The table is now available to be downloaded!"
			goTo := "/admin/api"
			controller.DisplaySuccessMessage(message, goTo)
		}
	}
}

// ShowAddForm shows the form to add a table to the download list
func (controller *APIController) ShowAddForm() {
	controller.Data["type"] = "Add"
	controller.SetCustomTitle("Admin - API - Add table")
	controller.showForm()
}

// ShowModifyForm shows the form to modify a tanl
func (controller *APIController) ShowModifyForm() {
	controller.Data["type"] = "Modify"
	id := controller.Ctx.Input.Param(":id")
	table, _ := api.NewTable(id)
	controller.Data["currentTable"] = table
	controller.SetCustomTitle("Modify table")
	controller.showForm()
}

// ModifyTable changes the description
func (controller *APIController) ModifyTable() {
	id := strings.TrimSpace(controller.GetString("table-id"))
	description := strings.TrimSpace(controller.GetString("table-description"))
	table, err := api.NewTable(id)
	if err != nil {
		controller.DisplaySimpleError("This table is not found")
	} else {
		err := api.ModifyDetails(table, description)
		if err != nil {
			controller.DisplaySimpleError(err.Error())
		} else {
			message := "The table is now available to be downloaded!"
			goTo := "/admin/api"
			controller.DisplaySuccessMessage(message, goTo)
		}
	}
}

func (controller *APIController) showForm() {
	controller.GenerateXSRF()
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

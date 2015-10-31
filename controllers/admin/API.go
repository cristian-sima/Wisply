package admin

import (
	"strings"

	"github.com/cristian-sima/Wisply/models/api"
)

// Developers manages the operation for API
type Developers struct {
	Controller
}

// RemoveAllowedTable removes the table from the list
func (controller *Developers) RemoveAllowedTable() {
	ID := strings.TrimSpace(controller.GetString("table-id"))
	table, err := api.NewTable(ID)
	if err != nil {
		controller.DisplaySimpleError(err.Error())
	} else {
		err = table.Delete()
		if err != nil {
			controller.Abort("404")
		} else {
			message := "The table " + table.Name + " is no longer available to download."
			goTo := "/admin/api"
			controller.DisplaySuccessMessage(message, goTo)
		}
	}
}

// InsertNewTable inserts the new table name into the list
func (controller *Developers) InsertNewTable() {
	name := strings.TrimSpace(controller.GetString("table-name"))
	description := strings.TrimSpace(controller.GetString("table-description"))
	table := api.Table{
		Name:        name,
		Description: description,
	}
	if !api.AreValidDetails(table) {
		controller.DisplaySimpleError("This table name can't be inserted.")
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
func (controller *Developers) ShowAddForm() {
	controller.Data["type"] = "Add"
	controller.SetCustomTitle("Admin - API - Add table")
	controller.showForm()
}

// ShowModifyForm shows the form to modify a tanl
func (controller *Developers) ShowModifyForm() {
	controller.Data["type"] = "Modify"
	id := controller.Ctx.Input.Param(":id")
	table, _ := api.NewTable(id)
	controller.Data["currentTable"] = table
	controller.SetCustomTitle("Modify table")
	controller.showForm()
}

// ModifyTable changes the description
func (controller *Developers) ModifyTable() {
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

func (controller *Developers) showForm() {
	controller.GenerateXSRF()
	controller.Data["tables"] = api.GetWisplyTablesNamesNotAllowed()
	controller.Data["action"] = "Allow table to be downloaded"
	controller.TplNames = "site/admin/api/form.tpl"
}

// ShowHomePage displays the home page
func (controller *Developers) ShowHomePage() {
	controller.GenerateXSRF()
	controller.Data["tables"] = api.GetAllowedTables()
	controller.SetCustomTitle("Admin - API")
	controller.TplNames = "site/admin/api/home.tpl"
}

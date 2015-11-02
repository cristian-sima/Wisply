package table

import (
	"strings"

	"github.com/cristian-sima/Wisply/models/developer"
)

// Table manages the operations for a table
type Table struct {
	Controller
}

// RemoveAllowedTable removes the table from the list
func (controller *Table) RemoveAllowedTable() {
	table := controller.GetTable()
	err := table.Delete()
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		message := "The table " + table.Name + " is no longer available to download."
		goTo := "/admin/developers"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// ShowModifyForm shows the form to modify a tanl
func (controller *Table) ShowModifyForm() {
	controller.Data["type"] = "Modify"
	controller.SetCustomTitle("Modify table")
	controller.showForm()
}

// ModifyTable changes the description
func (controller *Table) ModifyTable() {
	id := strings.TrimSpace(controller.GetString("table-id"))
	description := strings.TrimSpace(controller.GetString("table-description"))
	table, err := developer.NewTable(id)
	if err != nil {
		controller.DisplaySimpleError("This table is not found")
	} else {
		err := developer.ModifyDetails(table, description)
		if err != nil {
			controller.DisplaySimpleError(err.Error())
		} else {
			message := "The table is now available to be downloaded!"
			goTo := "/admin/developers"
			controller.DisplaySuccessMessage(message, goTo)
		}
	}
}

// InsertNewTable inserts the new table name into the list
func (controller *Table) InsertNewTable() {
	name := strings.TrimSpace(controller.GetString("table-name"))
	description := strings.TrimSpace(controller.GetString("table-description"))
	table := developer.Table{
		Name:        name,
		Description: description,
	}
	if !developer.AreValidDetails(table) {
		controller.DisplaySimpleError("This table name can't be inserted. Check the description and make sure that it is allowed.")
	} else {
		err := developer.InsertNewTable(table)
		if err != nil {
			controller.DisplaySimpleError(err.Error())
		} else {
			message := "The table is now available to be downloaded!"
			goTo := "/admin/developers"
			controller.DisplaySuccessMessage(message, goTo)
		}
	}
}

// ShowAddForm shows the form to add a table to the download list
func (controller *Table) ShowAddForm() {
	controller.Data["type"] = "Add"
	controller.SetCustomTitle("Admin - developer - Add table")
	controller.showForm()
}

func (controller *Table) showForm() {
	controller.GenerateXSRF()
	controller.Data["tables"] = developer.GetWisplyTablesNamesNotAllowed()
	controller.Data["action"] = "Allow table to be downloaded"
	controller.LoadTemplate("form")
}

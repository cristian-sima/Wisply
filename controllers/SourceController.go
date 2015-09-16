package controllers

import (
	"strings"

	SourcesModel "github.com/cristian-sima/Wisply/models/sources"
)

// SourceController It manages the operations for sources (list, delete, add)
type SourceController struct {
	AdminController
	model SourcesModel.Model
}

// ListSources It shows all the sources
func (controller *SourceController) ListSources() {

	var exists bool

	list := controller.model.GetAll()

	exists = (len(list) != 0)

	controller.Data["anything"] = exists
	controller.Data["sources"] = list
	controller.TplNames = "site/source/list.tpl"
	controller.Layout = "site/admin.tpl"
}

// AddNewSource It shows the form to add a new source
func (controller *SourceController) AddNewSource() {
	controller.showAddForm()
}

// InsertSource It inserts a source in the database
func (controller *SourceController) InsertSource() {

	sourceDetails := make(map[string]interface{})
	sourceDetails["name"] = strings.TrimSpace(controller.GetString("source-name"))
	sourceDetails["description"] = strings.TrimSpace(controller.GetString("source-description"))
	sourceDetails["url"] = strings.TrimSpace(controller.GetString("source-URL"))

	problems, err := controller.model.InsertNewSource(sourceDetails)
	if err != nil {
		controller.DisplayError(problems)
	} else {
		controller.DisplaySuccessMessage("The source has been added!", "/admin/sources/")
	}
}

// Modify It shows the form to modify a source's details
func (controller *SourceController) Modify() {

	var ID string

	ID = controller.Ctx.Input.Param(":id")

	source, err := controller.model.NewSource(ID)

	if err != nil {
		controller.Abort("databaseError")
	} else {
		sourceDetails := map[string]string{
			"Name":        source.Name,
			"Description": source.Description,
			"Url":         source.Url,
		}
		controller.showModifyForm(sourceDetails)
	}
}

// Update It updates a source in the database
func (controller *SourceController) Update() {

	var ID string
	sourceDetails := make(map[string]interface{})

	ID = controller.Ctx.Input.Param(":id")

	sourceDetails["name"] = strings.TrimSpace(controller.GetString("source-name"))
	sourceDetails["description"] = strings.TrimSpace(controller.GetString("source-description"))
	sourceDetails["url"] = strings.TrimSpace(controller.GetString("source-URL"))

	source, err := controller.model.NewSource(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		problems, err := source.Modify(sourceDetails)
		if err != nil {
			controller.DisplayError(problems)
		} else {
			controller.DisplaySuccessMessage("The account has been modified!", "/admin/sources/")
		}
	}
}

// Delete It deletes the source specified by parameter id
func (controller *SourceController) Delete() {
	var ID string
	ID = controller.Ctx.Input.Param(":id")

	source, err := controller.model.NewSource(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		databaseError := source.Delete()
		if databaseError != nil {
			controller.Abort("databaseError")
		} else {
			controller.DisplaySuccessMessage("The source ["+source.Name+"] has been deleted. Well done!", "/admin/sources/")
		}
	}
}

func (controller *SourceController) showModifyForm(source map[string]string) {
	controller.Data["sourceName"] = source["Name"]
	controller.Data["sourceUrl"] = source["Url"]
	controller.Data["sourceDescription"] = source["Description"]
	controller.showForm("Modify", "Modify this source")
}

func (controller *SourceController) showAddForm() {
	controller.showForm("Add", "Add a new source")
}

func (controller *SourceController) showForm(action string, legend string) {
	controller.GenerateXsrf()
	controller.Data["action"] = action
	controller.Data["legend"] = legend
	controller.Data["actionURL"] = ""
	controller.Data["actionType"] = "POST"
	controller.Layout = "site/admin.tpl"
	controller.TplNames = "site/source/form.tpl"
}

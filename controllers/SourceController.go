package controllers

import (
	SourcesModel "github.com/cristian-sima/Wisply/models/sources"
	"strings"
)

type SourceController struct {
	AdminController
	model SourcesModel.Model
}

func (controller *SourceController) ListSources() {

	var exists bool = false

	list := controller.model.GetAll()

	exists = (len(list) != 0)

	controller.Data["anything"] = exists
	controller.Data["sources"] = list
	controller.TplNames = "site/source/list.tpl"
	controller.Layout = "site/admin.tpl"
}

func (controller *SourceController) AddNewSource() {
	controller.showAddForm()
}

func (controller *SourceController) InsertSource() {

	rawData := make(map[string]interface{})
	rawData["name"] = strings.TrimSpace(controller.GetString("source-name"))
	rawData["description"] = strings.TrimSpace(controller.GetString("source-description"))
	rawData["url"] = strings.TrimSpace(controller.GetString("source-URL"))

	problems, err := controller.model.ValidateSource(rawData)
	if err != nil {
		controller.DisplayErrorMessages(problems)
	} else {
		databaseError := controller.model.InsertNewSource(rawData)
		if databaseError != nil {
			controller.Abort("databaseError")
		} else {
			controller.DisplaySuccessMessage("The source has been added!", "/admin/sources/")
		}
	}
}

func (controller *SourceController) Modify() {

	var id string

	id = controller.Ctx.Input.Param(":id")

	source, err := controller.model.GetSourceById(id)

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

func (controller *SourceController) Update() {

	var sourceId string
	rawData := make(map[string]interface{})

	sourceId = controller.Ctx.Input.Param(":id")

	rawData["name"] = strings.TrimSpace(controller.GetString("source-name"))
	rawData["description"] = strings.TrimSpace(controller.GetString("source-description"))
	rawData["url"] = strings.TrimSpace(controller.GetString("source-URL"))

	_, err := controller.model.GetSourceById(sourceId)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		problems, err := controller.model.ValidateSource(rawData)
		if err != nil {
			controller.DisplayErrorMessages(problems)
		} else {
			databaseError := controller.model.UpdateSourceById(sourceId, rawData)
			if databaseError != nil {
				controller.Abort("databaseError")
			} else {
				controller.DisplaySuccessMessage("The source has been modified!", "/admin/sources/")
			}
		}
	}
}

func (controller *SourceController) Delete() {
	var id string
	id = controller.Ctx.Input.Param(":id")
	source, err := controller.model.GetSourceById(id)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		databaseError := controller.model.DeleteSourceById(id)
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

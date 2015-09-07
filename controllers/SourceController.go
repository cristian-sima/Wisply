package controllers

import (
    "strings"
    SourcesModel "github.com/cristian-sima/Wisply/models/sources"
)

type SourceController struct {
	DefaultController
    model SourcesModel.Model
}

func (c *SourceController) ListSources() {

    var  exists bool = false

    list := c.model.GetAll()

    exists = (len(list) != 0);

    c.Data["anything"] = exists
    c.Data["sources"] = list
    c.TplNames = "general/source/list.tpl"
    c.Layout = "general/admin.tpl"
}

func (c *SourceController) AddNewSource() {
	c.showAddForm()
}

func (c *SourceController) InsertSource() {

    rawData := make(map[string]interface{})
    rawData["name"] = strings.TrimSpace(c.GetString("source-name"))
    rawData["description"] = strings.TrimSpace(c.GetString("source-description"))
    rawData["url"] =  strings.TrimSpace(c.GetString("source-URL"))

    problems, err := c.model.ValidateSource(rawData)
    if err != nil {
        c.DisplayErrorMessage(problems)
    } else {
        databaseError := c.model.InsertNewSource(rawData)
        if databaseError != nil {
            c.Abort("databaseError");
        } else {
            c.DisplaySuccessMessage("The source has been added!", "/admin/sources/")
        }
    }
}

func (c *SourceController) Modify() {

	var id string

	id = c.Ctx.Input.Param(":id")

    source, err := c.model.GetSourceById(id)

    if err != nil {
		c.Abort("databaseError");
    } else {
		sourceDetails := map[string]string{
		"Name" : source.Name,
		"Description" : source.Description,
		"Url": source.Url,
		}
		c.showModifyForm(sourceDetails);
	}
}

func (c *SourceController) Update() {

	var sourceId string;
	rawData := make(map[string]interface{})

    sourceId = c.Ctx.Input.Param(":id")

    rawData["name"] = strings.TrimSpace(c.GetString("source-name"))
    rawData["description"] = strings.TrimSpace(c.GetString("source-description"))
    rawData["url"] =  strings.TrimSpace(c.GetString("source-URL"))

    _, err := c.model.GetSourceById(sourceId)
	if err != nil {
		c.Abort("databaseError");
	} else {
		problems, err := c.model.ValidateSource(rawData)
		if err != nil {
			c.DisplayErrorMessage(problems)
		} else {
			databaseError := c.model.UpdateSourceById(sourceId, rawData)
			if databaseError != nil {
				c.Abort("databaseError");
			} else {
				c.DisplaySuccessMessage("The source has been modified!", "/admin/sources/")
			}
		}
	}
}

func (c *SourceController) Delete () {
    var id string
	id = c.Ctx.Input.Param(":id")
	source, err := c.model.GetSourceById(id)
	if err != nil {
		c.Abort("databaseError");
	} else {
		databaseError := c.model.DeleteSourceById(id)
		if databaseError != nil {
			c.Abort("databaseError");
		} else {
			c.DisplaySuccessMessage("The source [" + source.Name + "] has been deleted. Well done!", "/admin/sources/")
		}
	}
}

func (c *SourceController) showModifyForm(source map[string]string) {
	c.Data["sourceName"] = source["Name"]
	c.Data["sourceUrl"] = source["Url"]
	c.Data["sourceDescription"] = source["Description"]
	c.showForm("Modify", "Modify this source");
}

func (c *SourceController) showAddForm() {
	c.showForm("Add", "Add a new source");
}

func (c *SourceController) showForm(action string, legend string) {
  c.GenerateXsrf();
	c.Data["action"] = action
	c.Data["legend"] = legend
	c.Data["actionURL"] = ""
	c.Data["actionType"] = "POST"
	c.Layout = "general/admin.tpl"
	c.TplNames = "general/source/form.tpl"
}

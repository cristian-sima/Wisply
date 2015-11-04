package institution

import (
	"strings"

	"github.com/cristian-sima/Wisply/models/repository"
)

// Institution manages the operations with an institution
type Institution struct {
	Controller
}

// Display shows the administrative page for an institution
func (controller *Institution) Display() {
	institution := controller.GetInstitution()
	controller.Data["repositories"] = institution.GetRepositories()
	controller.Data["institutionPrograms"] = institution.GetPrograms()
	controller.LoadTemplate("institution")
}

// ShowInsertForm shows the form to add an institution
func (controller *Institution) ShowInsertForm() {
	controller.SetCustomTitle("Add Institution")
	controller.showAddForm()
}

// Insert inserts an institution in the database
func (controller *Institution) Insert() {

	institutionDetails := make(map[string]interface{})
	institutionDetails["name"] = strings.TrimSpace(controller.GetString("institution-name"))
	description := controller.GetString("institution-description")

	// jquery has a problem with \r
	institutionDetails["description"] = strings.TrimSpace(strings.Replace(description, "\r", "", -1))

	institutionDetails["url"] = strings.TrimSpace(controller.GetString("institution-URL"))
	institutionDetails["logoURL"] = strings.TrimSpace(controller.GetString("institution-logoURL"))
	institutionDetails["wikiURL"] = strings.TrimSpace(controller.GetString("institution-wikiURL"))
	institutionDetails["wikiID"] = strings.TrimSpace(controller.GetString("institution-wikiID"))

	problems, err := repository.InsertNewInstitution(institutionDetails)
	if err != nil {
		controller.DisplayError(problems)
	} else {
		message := "The institution has been inserted."
		goTo := "/admin/institutions/"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// ShowModifyForm shows the form to modify a institution's details
func (controller *Institution) ShowModifyForm() {
	institution := controller.GetInstitution()
	wikiReceive := false
	if institution.WikiID == "" {
		institution.WikiID = "NULL"
	}
	if institution.WikiID == "NULL" {
		wikiReceive = false
	} else {
		wikiReceive = true
	}
	controller.Data["wikiID"] = institution.WikiID
	controller.Data["wikiReceive"] = wikiReceive
	controller.showModifyForm()
}

// Modify updates an institution in the database
func (controller *Institution) Modify() {
	institution := controller.GetInstitution()

	institutionDetails := make(map[string]interface{})

	institutionDetails["name"] = strings.TrimSpace(controller.GetString("institution-name"))
	description := controller.GetString("institution-description")

	// jquery has a problem with \r
	institutionDetails["description"] = strings.TrimSpace(strings.Replace(description, "\r", "", -1))

	institutionDetails["logoURL"] = strings.TrimSpace(controller.GetString("institution-logoURL"))
	institutionDetails["wikiURL"] = strings.TrimSpace(controller.GetString("institution-wikiURL"))
	institutionDetails["wikiID"] = strings.TrimSpace(controller.GetString("institution-wikiID"))

	problems, err := institution.Modify(institutionDetails)
	if err != nil {
		controller.DisplayError(problems)
	} else {
		message := "The account has been modified!"
		goTo := "/admin/institutions/"
		controller.DisplaySuccessMessage(message, goTo)
	}

}

// Delete deletes the institution specified by parameter id
func (controller *Institution) Delete() {
	institution := controller.GetInstitution()
	err := institution.Delete()
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		message := "The institution [" + institution.Name + "] has been deleted."
		goTo := "/admin/institutions/"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// ShowAdvanceOptions displays the page with further options
// For instance, further options may be modify or delete
func (controller *Institution) ShowAdvanceOptions() {
	controller.SetCustomTitle("Admin - Institution - Advance options")
	controller.LoadTemplate("advance-options")

}

func (controller *Institution) showModifyForm() {
	controller.showForm("Modify", "Modify this institution")
}

func (controller *Institution) showAddForm() {
	controller.showForm("Add", "Add a new institution")
	controller.Data["wikiID"] = "NULL"
	controller.Data["wikiReceive"] = false
}

func (controller *Institution) showForm(action string, legend string) {
	controller.GenerateXSRF()
	controller.Data["action"] = action
	controller.Data["legend"] = legend
	controller.Data["actionURL"] = ""
	controller.Data["actionType"] = "POST"
	controller.LoadTemplate("form")
}

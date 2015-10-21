package admin

import "strings"

import (
	"github.com/cristian-sima/Wisply/models/harvest"
	repository "github.com/cristian-sima/Wisply/models/repository"
)

// InstitutionController manages the operations for institutions
type InstitutionController struct {
	Controller
	model repository.Model
}

// DisplayAll shows all the institutions
func (controller *InstitutionController) DisplayAll() {
	var exists bool
	list := controller.model.GetAllInstitutions()
	exists = (len(list) != 0)
	controller.Data["anything"] = exists
	controller.Data["institutions"] = list
	controller.SetCustomTitle("Admin - Institutions")
	controller.TplNames = "site/admin/institution/list.tpl"
	controller.Layout = "site/admin-layout.tpl"
}

// Add shows the form to add an institution
func (controller *InstitutionController) Add() {
	controller.SetCustomTitle("Add Institution")
	controller.showAddForm()
}

// Insert inserts an institution in the database
func (controller *InstitutionController) Insert() {

	institutionDetails := make(map[string]interface{})
	institutionDetails["name"] = strings.TrimSpace(controller.GetString("institution-name"))
	description := controller.GetString("institution-description")

	// jquery has a problem with \r
	institutionDetails["description"] = strings.TrimSpace(strings.Replace(description, "\r", "", -1))

	institutionDetails["url"] = strings.TrimSpace(controller.GetString("institution-URL"))
	institutionDetails["logoURL"] = strings.TrimSpace(controller.GetString("institution-logoURL"))
	institutionDetails["wikiURL"] = strings.TrimSpace(controller.GetString("institution-wikiURL"))
	institutionDetails["wikiID"] = strings.TrimSpace(controller.GetString("institution-wikiID"))

	problems, err := controller.model.InsertNewInstitution(institutionDetails)
	if err != nil {
		controller.DisplayError(problems)
	} else {
		message := "The institution has been added!"
		goTo := "/admin/institutions/"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// Modify shows the form to modify a institution's details
func (controller *InstitutionController) Modify() {
	var ID string
	ID = controller.Ctx.Input.Param(":id")
	institution, err := repository.NewInstitution(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		controller.Data["institution"] = institution
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
}

// Update updates an institution in the database
func (controller *InstitutionController) Update() {
	institutionDetails := make(map[string]interface{})

	ID := controller.Ctx.Input.Param(":id")
	institutionDetails["name"] = strings.TrimSpace(controller.GetString("institution-name"))
	description := controller.GetString("institution-description")

	// jquery has a problem with \r
	institutionDetails["description"] = strings.TrimSpace(strings.Replace(description, "\r", "", -1))

	institutionDetails["logoURL"] = strings.TrimSpace(controller.GetString("institution-logoURL"))
	institutionDetails["wikiURL"] = strings.TrimSpace(controller.GetString("institution-wikiURL"))
	institutionDetails["wikiID"] = strings.TrimSpace(controller.GetString("institution-wikiID"))

	institution, err := repository.NewInstitution(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		problems, err := institution.Modify(institutionDetails)
		if err != nil {
			controller.DisplayError(problems)
		} else {
			message := "The account has been modified!"
			goTo := "/admin/institutions/"
			controller.DisplaySuccessMessage(message, goTo)
		}
	}
}

// Delete deletes the institution specified by parameter id
func (controller *InstitutionController) Delete() {
	var ID string
	ID = controller.Ctx.Input.Param(":id")

	institution, err := repository.NewInstitution(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		repositories := institution.GetRepositories()

		for _, repository := range repositories {
			processes := harvest.GetProcessesByRepository(repository.ID, 0)
			for _, process := range processes {
				process.Delete()
			}
		}

		databaseError := institution.Delete()
		if databaseError != nil {
			controller.Abort("databaseError")
		} else {
			message := "The institution [" + institution.Name + "] has been deleted. Well done!"
			goTo := "/admin/institutions/"
			controller.DisplaySuccessMessage(message, goTo)
		}
	}
}

// ShowInstitution shows the administrative page for an institution
func (controller *RepositoryController) ShowInstitution() {
	ID := controller.Ctx.Input.Param(":id")
	institution, err := repository.NewInstitution(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		controller.Data["institution"] = institution
		controller.Data["repositories"] = institution.GetRepositories()
		controller.TplNames = "site/admin/institution/institution.tpl"
	}
}

// ShowAdvanceInstitutionOptions displays the page with further options
// For instance, further options may be modify or delete
func (controller *RepositoryController) ShowAdvanceInstitutionOptions() {
	ID := controller.Ctx.Input.Param(":id")
	institution, err := repository.NewInstitution(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		controller.Data["institution"] = institution
		controller.TplNames = "site/admin/institution/advance-options.tpl"
	}
}

func (controller *InstitutionController) showModifyForm() {
	controller.showForm("Modify", "Modify this institution")
}

func (controller *InstitutionController) showAddForm() {
	controller.showForm("Add", "Add a new institution")
	controller.Data["wikiID"] = "NULL"
	controller.Data["wikiReceive"] = false
}

func (controller *InstitutionController) showForm(action string, legend string) {
	controller.GenerateXSRF()
	controller.Data["action"] = action
	controller.Data["legend"] = legend
	controller.Data["actionURL"] = ""
	controller.Data["actionType"] = "POST"
	controller.TplNames = "site/admin/institution/form.tpl"
}

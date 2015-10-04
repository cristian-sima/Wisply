package admin

import "strings"

import repository "github.com/cristian-sima/Wisply/models/repository"

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
	institutionDetails["description"] = strings.TrimSpace(controller.GetString("institution-description"))
	institutionDetails["url"] = strings.TrimSpace(controller.GetString("institution-URL"))

	problems, err := controller.model.InsertNewInstitution(institutionDetails)
	if err != nil {
		controller.DisplayError(problems)
	} else {
		controller.DisplaySuccessMessage("The institution has been added!", "/admin/institutions/")
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
		institutionDetails := map[string]string{
			"Name":        institution.Name,
			"Description": institution.Description,
		}
		controller.showModifyForm(institutionDetails)
	}
}

// Update updates an institution in the database
func (controller *InstitutionController) Update() {

	var ID string
	institutionDetails := make(map[string]interface{})

	ID = controller.Ctx.Input.Param(":id")

	institutionDetails["name"] = strings.TrimSpace(controller.GetString("institution-name"))
	institutionDetails["description"] = strings.TrimSpace(controller.GetString("institution-description"))

	institution, err := repository.NewInstitution(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		problems, err := institution.Modify(institutionDetails)
		if err != nil {
			controller.DisplayError(problems)
		} else {
			controller.DisplaySuccessMessage("The account has been modified!", "/admin/institutions/")
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
		databaseError := institution.Delete()
		if databaseError != nil {
			controller.Abort("databaseError")
		} else {
			controller.DisplaySuccessMessage("The institution ["+institution.Name+"] has been deleted. Well done!", "/admin/institutions/")
		}
	}
}

func (controller *InstitutionController) showModifyForm(institution map[string]string) {
	controller.Data["institutionName"] = institution["Name"]
	controller.Data["institutionUrl"] = institution["Url"]
	controller.Data["institutionDescription"] = institution["Description"]
	controller.showForm("Modify", "Modify this institution")
}

func (controller *InstitutionController) showAddForm() {
	controller.showForm("Add", "Add a new institution")
}

func (controller *InstitutionController) showForm(action string, legend string) {
	controller.GenerateXSRF()
	controller.Data["action"] = action
	controller.Data["legend"] = legend
	controller.Data["actionURL"] = ""
	controller.Data["actionType"] = "POST"
	controller.TplNames = "site/admin/institution/form.tpl"
}

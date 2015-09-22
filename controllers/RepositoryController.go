package controllers

import (
	"strings"

	RepositoryModel "github.com/cristian-sima/Wisply/models/repository"
)

// RepositoryController It manages the operations for repositories (list, delete, add)
type RepositoryController struct {
	AdminController
	model RepositoryModel.Model
}

// ListRepositories It shows all the repositories
func (controller *RepositoryController) ListRepositories() {

	var exists bool

	list := controller.model.GetAll()

	exists = (len(list) != 0)

	controller.Data["anything"] = exists
	controller.Data["repositories"] = list
	controller.TplNames = "site/repository/list.tpl"
	controller.Layout = "site/admin.tpl"
}

// AddNewRepository It shows the form to add a new repository
func (controller *RepositoryController) AddNewRepository() {
	controller.showAddForm()
}

// InsertRepository It inserts a repository in the database
func (controller *RepositoryController) InsertRepository() {

	repositoryDetails := make(map[string]interface{})
	repositoryDetails["name"] = strings.TrimSpace(controller.GetString("repository-name"))
	repositoryDetails["description"] = strings.TrimSpace(controller.GetString("repository-description"))
	repositoryDetails["url"] = strings.TrimSpace(controller.GetString("repository-URL"))

	problems, err := controller.model.InsertNewRepository(repositoryDetails)
	if err != nil {
		controller.DisplayError(problems)
	} else {
		controller.DisplaySuccessMessage("The repository has been added!", "/admin/repositories/")
	}
}

// Modify It shows the form to modify a repository's details
func (controller *RepositoryController) Modify() {

	var ID string

	ID = controller.Ctx.Input.Param(":id")

	repository, err := controller.model.NewRepository(ID)

	if err != nil {
		controller.Abort("databaseError")
	} else {
		repositoryDetails := map[string]string{
			"Name":        repository.Name,
			"Description": repository.Description,
		}
		controller.showModifyForm(repositoryDetails)
	}
}

// Update It updates a repository in the database
func (controller *RepositoryController) Update() {

	var ID string
	repositoryDetails := make(map[string]interface{})

	ID = controller.Ctx.Input.Param(":id")

	repositoryDetails["name"] = strings.TrimSpace(controller.GetString("repository-name"))
	repositoryDetails["description"] = strings.TrimSpace(controller.GetString("repository-description"))

	repository, err := controller.model.NewRepository(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		problems, err := repository.Modify(repositoryDetails)
		if err != nil {
			controller.DisplayError(problems)
		} else {
			controller.DisplaySuccessMessage("The account has been modified!", "/admin/repositories/")
		}
	}
}

// Delete It deletes the repository specified by parameter id
func (controller *RepositoryController) Delete() {
	var ID string
	ID = controller.Ctx.Input.Param(":id")

	repository, err := controller.model.NewRepository(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		databaseError := repository.Delete()
		if databaseError != nil {
			controller.Abort("databaseError")
		} else {
			controller.DisplaySuccessMessage("The repository ["+repository.Name+"] has been deleted. Well done!", "/admin/repositories/")
		}
	}
}

func (controller *RepositoryController) showModifyForm(repository map[string]string) {
	controller.Data["repositoryName"] = repository["Name"]
	controller.Data["repositoryUrl"] = repository["Url"]
	controller.Data["repositoryDescription"] = repository["Description"]
	controller.showForm("Modify", "Modify this repository")
}

func (controller *RepositoryController) showAddForm() {
	controller.showForm("Add", "Add a new repository")
}

func (controller *RepositoryController) showForm(action string, legend string) {
	controller.GenerateXSRF()
	controller.Data["action"] = action
	controller.Data["legend"] = legend
	controller.Data["actionURL"] = ""
	controller.Data["actionType"] = "POST"
	controller.Layout = "site/admin.tpl"
	controller.TplNames = "site/repository/form.tpl"
}

package admin

import (
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/harvest"
	repository "github.com/cristian-sima/Wisply/models/repository"
)

// RepositoryController manages the operations for repositories (list, delete, add)
type RepositoryController struct {
	Controller
	model repository.Model
}

// List shows all the repositories
func (controller *RepositoryController) List() {
	list := controller.model.GetAllRepositories()
	controller.Data["repositories"] = list
	controller.Data["host"] = controller.Ctx.Request.Host
	controller.SetCustomTitle("Admin - Repositories")
	controller.TplNames = "site/admin/repository/list.tpl"
}

// ShowTypes displays the types which are available
func (controller *RepositoryController) ShowTypes() {
	controller.Data["institution"] = strings.TrimSpace(controller.GetString("institution"))
	controller.TplNames = "site/admin/repository/add/category.tpl"
}

// Add shows the form to add a new repository
func (controller *RepositoryController) Add() {
	controller.Data["institutions"] = controller.model.GetAllInstitutions()
	selected, _ := strconv.Atoi(strings.TrimSpace(controller.GetString("institution")))
	controller.Data["selectedInstitution"] = selected
	controller.Data["category"] = strings.TrimSpace(controller.GetString("category"))
	controller.SetCustomTitle("Add Repository")
	controller.showAddForm()
}

// Insert inserts a repository in the database
func (controller *RepositoryController) Insert() {

	repositoryDetails := make(map[string]interface{})
	repositoryDetails["name"] = strings.TrimSpace(controller.GetString("repository-name"))
	repositoryDetails["description"] = strings.TrimSpace(controller.GetString("repository-description"))
	repositoryDetails["url"] = strings.TrimSpace(controller.GetString("repository-URL"))
	repositoryDetails["institution"] = strings.TrimSpace(controller.GetString("repository-institution"))
	repositoryDetails["public-url"] = strings.TrimSpace(controller.GetString("repository-public-url"))
	repositoryDetails["category"] = strings.TrimSpace(controller.GetString("repository-category"))

	problems, err := controller.model.InsertNewRepository(repositoryDetails)
	if err != nil {
		controller.DisplayError(problems)
	} else {
		controller.DisplaySuccessMessage("The repository has been added!", "/admin/repositories/")
	}
}

// Modify shows the form to modify a repository's details
func (controller *RepositoryController) Modify() {

	var ID string

	ID = controller.Ctx.Input.Param(":id")

	repository, err := repository.NewRepository(ID)

	if err != nil {
		controller.Abort("databaseError")
	} else {
		controller.Data["repository"] = repository
		controller.showModifyForm()
	}
}

// Update updates a repository in the database
func (controller *RepositoryController) Update() {

	var ID string
	repositoryDetails := make(map[string]interface{})

	ID = controller.Ctx.Input.Param(":id")

	repositoryDetails["name"] = strings.TrimSpace(controller.GetString("repository-name"))
	repositoryDetails["description"] = strings.TrimSpace(controller.GetString("repository-description"))
	repositoryDetails["institution"] = strings.TrimSpace(controller.GetString("repository-institution"))
	repositoryDetails["url"] = strings.TrimSpace(controller.GetString("repository-URL"))

	repository, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		problems, err := repository.Modify(repositoryDetails)
		if err != nil {
			controller.DisplayError(problems)
		} else {
			controller.DisplaySuccessMessage("The account has been modified!", "/admin/repositories/repository/"+strconv.Itoa(repository.ID))
		}
	}
}

// Delete deletes the repository specified by parameter id
func (controller *RepositoryController) Delete() {
	var ID string
	ID = controller.Ctx.Input.Param(":id")
	repository, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		processes := harvest.GetProcessesByRepository(repository.ID)
		for _, process := range processes {
			process.Delete()
		}
		databaseError := repository.Delete()
		if databaseError != nil {
			controller.Abort("databaseError")
		} else {
			controller.DisplaySuccessMessage("The repository ["+repository.Name+"] has been deleted. Well done!", "/admin/repositories/")
		}
	}
}

func (controller *RepositoryController) showModifyForm() {
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
	controller.TplNames = "site/admin/repository/form.tpl"
}

// ShowRepository shows the administrative information regarding a repository
func (controller *RepositoryController) ShowRepository() {
	ID := controller.Ctx.Input.Param(":id")
	repository, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		controller.Data["repository"] = repository
		controller.Data["institution"] = repository.GetInstitution()
		controller.Data["identification"] = repository.GetIdentification()
		controller.TplNames = "site/admin/repository/repository.tpl"
	}
}

// ShowAdvanceOptions displays the page with further options such as modify or delete
func (controller *RepositoryController) ShowAdvanceOptions() {
	ID := controller.Ctx.Input.Param(":id")
	repository, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("databaseError")
	} else {
		controller.Data["repository"] = repository
		controller.Data["institution"] = repository.GetInstitution()
		controller.Data["identification"] = repository.GetIdentification()
		controller.TplNames = "site/admin/repository/advance-options.tpl"
	}
}

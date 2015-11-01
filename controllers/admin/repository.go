package admin

import (
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/harvest"
	repository "github.com/cristian-sima/Wisply/models/repository"
	"github.com/cristian-sima/Wisply/models/wisply"
)

// Repository manages the operations for repositories
// For example: (list, delete, add)
type Repository struct {
	Controller
	model repository.Model
}

// List shows all the repositories
func (controller *Repository) List() {
	list := repository.GetAllRepositories()
	controller.Data["repositories"] = list
	controller.Data["host"] = controller.Ctx.Request.Host
	controller.SetCustomTitle("Admin - Repositories")
	controller.TplNames = "site/admin/repository/list.tpl"
}

// ShowTypes displays the types which are available
func (controller *Repository) ShowTypes() {
	controller.Data["institution"] = strings.TrimSpace(controller.GetString("institution"))
	controller.TplNames = "site/admin/repository/add/category.tpl"
	controller.SetCustomTitle("Add Repository")
}

// Add shows the form to add a new repository
func (controller *Repository) Add() {
	controller.Data["institutions"] = repository.GetAllInstitutions()
	selected, _ := strconv.Atoi(strings.TrimSpace(controller.GetString("institution")))
	controller.Data["selectedInstitution"] = selected
	controller.Data["category"] = strings.TrimSpace(controller.GetString("category"))
	controller.SetCustomTitle("Add Repository")
	controller.showAddForm()
}

// Insert inserts a repository in the database
func (controller *Repository) Insert() {

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
		message := "The repository has been added!"
		goTo := "/admin/repositories/"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// Modify shows the form to modify a repository's details
func (controller *Repository) Modify() {
	ID := controller.Ctx.Input.Param(":id")
	repository, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		controller.Data["repository"] = repository
		controller.showModifyForm()
	}
}

// ShowFilter shows the page for modifying the filter
func (controller *Repository) ShowFilter() {
	ID := controller.Ctx.Input.Param(":id")
	repository, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		controller.GenerateXSRF()
		controller.Data["repository"] = repository
		controller.TplNames = "site/admin/repository/filter.tpl"
	}
}

// ChangeFilter changes the filter
func (controller *Repository) ChangeFilter() {
	ID := controller.Ctx.Input.Param(":id")
	filter := strings.TrimSpace(controller.GetString("repository-filter"))
	repository, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		err := repository.SetFilter(filter)
		if err != nil {
			controller.Abort("show-database-error")
		} else {
			message := "The filter has been modified!"
			goTo := "/admin/repositories/repository/" + strconv.Itoa(repository.ID)
			controller.DisplaySuccessMessage(message, goTo)
		}
	}
}

// Update updates a repository in the database
func (controller *Repository) Update() {

	var ID string
	repositoryDetails := make(map[string]interface{})

	ID = controller.Ctx.Input.Param(":id")

	repositoryDetails["name"] = strings.TrimSpace(controller.GetString("repository-name"))
	repositoryDetails["description"] = strings.TrimSpace(controller.GetString("repository-description"))
	repositoryDetails["institution"] = strings.TrimSpace(controller.GetString("repository-institution"))
	repositoryDetails["url"] = strings.TrimSpace(controller.GetString("repository-URL"))

	repository, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		problems, err := repository.Modify(repositoryDetails)
		if err != nil {
			controller.DisplayError(problems)
		} else {
			message := "The account has been modified!"
			goTo := "/admin/repositories/repository/" + strconv.Itoa(repository.ID)
			controller.DisplaySuccessMessage(message, goTo)
		}
	}
}

// EmptyRepository deletes all records, formats, collections,
// emails and information about the repository
func (controller *Repository) EmptyRepository() {
	var ID string
	ID = controller.Ctx.Input.Param(":id")
	repository, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		processes := harvest.GetProcessesByRepository(repository.ID, 0)
		for _, process := range processes {
			process.Delete()
		}
		wisply.ClearRepository(repository.ID)
	}
	controller.TplNames = "site/admin/repository/blank.tpl"
}

// Delete deletes the repository specified by parameter id
func (controller *Repository) Delete() {
	var ID string
	ID = controller.Ctx.Input.Param(":id")
	repository, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		processes := harvest.GetProcessesByRepository(repository.ID, 0)
		for _, process := range processes {
			process.Delete()
		}
		databaseError := repository.Delete()
		if databaseError != nil {
			controller.Abort("show-database-error")
		} else {
			message := "The repository [" + repository.Name + "] has been deleted!"
			goTo := "/admin/repositories/"
			controller.DisplaySuccessMessage(message, goTo)
		}
	}
}

func (controller *Repository) showModifyForm() {
	controller.showForm("Modify", "Modify this repository")
}

func (controller *Repository) showAddForm() {
	controller.showForm("Add", "Add a new repository")
}

func (controller *Repository) showForm(action string, legend string) {
	controller.GenerateXSRF()
	controller.Data["action"] = action
	controller.Data["legend"] = legend
	controller.Data["actionURL"] = ""
	controller.Data["actionType"] = "POST"
	controller.TplNames = "site/admin/repository/form.tpl"
}

// ShowRepository shows the administrative information regarding a repository
func (controller *Repository) ShowRepository() {
	ID := controller.Ctx.Input.Param(":id")
	repository, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		controller.Data["processes"] = harvest.GetProcessesByRepository(repository.ID, 5)
		controller.Data["repository"] = repository
		controller.Data["institution"] = repository.GetInstitution()
		controller.Data["identification"] = repository.GetIdentification()
		controller.TplNames = "site/admin/repository/repository.tpl"
	}
}

// ShowAdvanceOptions displays the page with further options such as modify or delete
func (controller *Repository) ShowAdvanceOptions() {
	ID := controller.Ctx.Input.Param(":id")
	repository, err := repository.NewRepository(ID)
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		controller.Data["repository"] = repository
		controller.Data["institution"] = repository.GetInstitution()
		controller.Data["identification"] = repository.GetIdentification()
		controller.TplNames = "site/admin/repository/advance-options.tpl"
	}
}

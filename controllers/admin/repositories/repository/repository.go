package repository

import (
	"strconv"
	"strings"

	"github.com/cristian-sima/Wisply/models/harvest"
	model "github.com/cristian-sima/Wisply/models/repository"
	"github.com/cristian-sima/Wisply/models/wisply"
)

// Repository manages the operations with an repository
type Repository struct {
	Controller
}

// Display shows the default page for the repository
func (controller *Repository) Display() {
	repository := controller.GetRepository()
	controller.Data["processes"] = harvest.GetProcessesByRepository(repository.ID, 5)
	controller.Data["institution"] = repository.GetInstitution()
	controller.Data["identification"] = repository.GetIdentification()
	controller.LoadTemplate("home")
}

// ShowChooseCategory displays the types which are available
func (controller *Repository) ShowChooseCategory() {
	controller.Data["institution"] = strings.TrimSpace(controller.GetString("institution"))
	controller.LoadTemplate("add-repository-category")
	controller.SetCustomTitle("Admin - Add a new repository")
}

// ShowInsertForm shows the form to add a new repository
func (controller *Repository) ShowInsertForm() {
	controller.Data["institutions"] = model.GetAllInstitutions()
	selectedInstitution, _ := strconv.Atoi(strings.TrimSpace(controller.GetString("institution")))
	controller.Data["selectedInstitution"] = selectedInstitution
	controller.Data["category"] = strings.TrimSpace(controller.GetString("category"))
	controller.SetCustomTitle("Add a new repository")
	controller.showForm("Add", "Add a new repository")
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

	problems, err := model.InsertNewRepository(repositoryDetails)
	if err != nil {
		controller.DisplayError(problems)
	} else {
		message := "The repository has been added!"
		goTo := "/admin/repositories/"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// ShowModifyForm shows the form to modify a repository's details
func (controller *Repository) ShowModifyForm() {
	controller.showForm("Modify", "Modify this repository")
}

// ShowFilterForm shows the page for modifying the filter
func (controller *Repository) ShowFilterForm() {
	repository := controller.GetRepository()
	controller.GenerateXSRF()
	controller.Data["repository"] = repository
	controller.LoadTemplate("filter")
}

// ModifyFilter changes the filter
func (controller *Repository) ModifyFilter() {
	repository := controller.GetRepository()
	filter := strings.TrimSpace(controller.GetString("repository-filter"))
	err := repository.SetFilter(filter)
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		message := "The filter has been modified!"
		goTo := "/admin/repositories/" + strconv.Itoa(repository.ID)
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// Modify updates a repository in the database
func (controller *Repository) Modify() {
	repository := controller.GetRepository()
	repositoryDetails := make(map[string]interface{})

	repositoryDetails["name"] = strings.TrimSpace(controller.GetString("repository-name"))
	repositoryDetails["description"] = strings.TrimSpace(controller.GetString("repository-description"))
	repositoryDetails["institution"] = strings.TrimSpace(controller.GetString("repository-institution"))
	repositoryDetails["url"] = strings.TrimSpace(controller.GetString("repository-URL"))

	problems, err := repository.Modify(repositoryDetails)
	if err != nil {
		controller.DisplayError(problems)
	} else {
		message := "The account has been modified!"
		goTo := "/admin/repositories/" + strconv.Itoa(repository.ID)
		controller.DisplaySuccessMessage(message, goTo)
	}
}

// ClearRepository deletes all records, formats, collections,
// emails and information about the repository
func (controller *Repository) ClearRepository() {
	repository := controller.GetRepository()
	processes := harvest.GetProcessesByRepository(repository.ID, 0)
	for _, process := range processes {
		process.Delete()
	}
	wisply.ClearRepository(repository.ID)
	controller.RemoveLayout()
	controller.ShowBlankPage()
}

// Delete deletes the repository specified by parameter id
func (controller *Repository) Delete() {
	repository := controller.GetRepository()
	err := repository.Delete()
	if err != nil {
		controller.Abort("show-database-error")
	} else {
		message := "The repository [" + repository.Name + "] has been deleted!"
		goTo := "/admin/repositories/"
		controller.DisplaySuccessMessage(message, goTo)
	}
}

func (controller *Repository) showForm(action string, legend string) {
	controller.GenerateXSRF()
	controller.Data["action"] = action
	controller.Data["legend"] = legend
	controller.Data["actionURL"] = ""
	controller.Data["actionType"] = "POST"
	controller.LoadTemplate("form")
}

// ShowAdvanceOptions displays the page with further options such as modify or delete
func (controller *Repository) ShowAdvanceOptions() {
	repository := controller.GetRepository()
	controller.Data["repository"] = repository
	controller.Data["institution"] = repository.GetInstitution()
	controller.Data["identification"] = repository.GetIdentification()
	controller.LoadTemplate("advance-options")
}

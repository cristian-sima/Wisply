package subject

import "github.com/cristian-sima/Wisply/models/repository"

// Subject manages the operations for curriculum
type Subject struct {
	Controller
}

// Display shows the public page for a subject of study
func (controller *Subject) Display() {
	controller.SetCustomTitle(controller.GetSubject().GetName())
	controller.Data["institutions"] = repository.GetAllInstitutions()
	controller.LoadTemplate("home")
}

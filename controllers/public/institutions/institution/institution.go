package institution

import "github.com/cristian-sima/Wisply/models/analyse"

// Institution manages the operations for an institution
type Institution struct {
	Controller
}

// Display shows the public page for an institution
func (controller *Institution) Display() {
	institution := controller.GetInstitution()
	controller.SetCustomTitle(institution.Name)
	controller.Data["repositories"] = institution.GetRepositories()
	controller.Data["institutionPrograms"] = institution.GetPrograms()
	controller.Data["institutionAnalyses"] = analyse.GetInstitutionAnalysersByInstitution(institution.ID)
	controller.LoadTemplate("home")
}

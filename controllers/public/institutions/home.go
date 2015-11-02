package institutions

import repository "github.com/cristian-sima/Wisply/models/repository"

// Home manages the operations for education
type Home struct {
	Controller
}

// Display shows all the curricula
func (controller *Home) Display() {
	controller.SetCustomTitle("Wisply - Institutions")
	controller.Data["institutions"] = repository.GetAllInstitutions()
	controller.LoadTemplate("home")
}

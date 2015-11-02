package institutions

import "github.com/cristian-sima/Wisply/models/repository"

// Home manages the default page for institutions
type Home struct {
	Controller
}

// Display shows a table with all the institutions
func (controller *Home) Display() {
	controller.SetCustomTitle("Admin - Institutions")
	controller.Data["institutions"] = repository.GetAllInstitutions()
	controller.LoadTemplate("home")
}

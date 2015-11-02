package developers

import "github.com/cristian-sima/Wisply/models/developer"

// Home manages the default page for accounts
type Home struct {
	Controller
}

// Display shows a table with all the accounts
func (controller *Home) Display() {
	controller.GenerateXSRF()
	controller.Data["tables"] = developer.GetAllowedTables()
	controller.SetCustomTitle("Admin - Developers")
	controller.LoadTemplate("home")
}

package accounts

import (
	model "github.com/cristian-sima/Wisply/models/auth"
)

// Home manages the default page for accounts
type Home struct {
	Controller
}

// Display shows a table with all the accounts
func (controller *Home) Display() {
	controller.SetCustomTitle("Admin - Accounts")
	controller.Data["accounts"] = model.GetAll()
	controller.LoadTemplate("home")
}

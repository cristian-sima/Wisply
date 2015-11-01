package admin

import "github.com/cristian-sima/Wisply/controllers/wisply"

// Controller must be inherited by all the pages that are for administrators
// It ensures that an account is connected when accesing the page.
type Controller struct {
	wisply.Controller
}

// Prepare redirects to a login page in case the account is not connected,
// else it loads the page
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	if controller.isAdministratorPage() {
		controller.RedirectToLoginPage()
	} else {
		controller.initAdmin()
	}
}

func (controller *Controller) isAdministratorPage() bool {
	return !controller.AccountConnected || !controller.Account.IsAdministrator
}

func (controller *Controller) initAdmin() {
	controller.Data["isAdminPage"] = true
	controller.SetLayout("admin")
	controller.SetTemplatePath("admin")
}

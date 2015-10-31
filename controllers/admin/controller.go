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
	if !controller.AccountConnected || !controller.Account.IsAdministrator {
		controller.Controller.RedirectToLoginPage()
	} else {
		controller.initPage()
	}
	controller.loadLayout()
}

func (controller *Controller) loadLayout() {
	controller.Layout = "site/admin-layout.tpl"
}

// initPage is called when an administrator is connected
func (controller *Controller) initPage() {
	controller.Data["isAdminPage"] = true
}

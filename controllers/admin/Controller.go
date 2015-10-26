// Package admin contains the controllers used on administration pages
// The administration page is a restricted area and an account needs special priviledges
package admin

import general "github.com/cristian-sima/Wisply/controllers/general"

// Controller must be inherited by all the pages that are for administrators
// It ensures that an account is connected when accesing the page
type Controller struct {
	general.WisplyController
}

// Prepare redirects to a login page in case the account is not connected,
// else it loads the page
func (controller *Controller) Prepare() {
	controller.WisplyController.Prepare()
	if !controller.AccountConnected || !controller.Account.IsAdministrator {
		controller.WisplyController.RedirectToLoginPage()
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

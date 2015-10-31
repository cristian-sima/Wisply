// Package account contrains the controllers which manage the account activity
// A controller from this package is access only BY THE CONNECTED ACCOUNTS
package account

import (
	"github.com/cristian-sima/Wisply/controllers/general"
)

// Controller represents the basic Account controller
type Controller struct {
	general.WisplyController
}

// Prepare redirects to a login page in case the account is not connected,
// else it loads the page
func (controller *Controller) Prepare() {
	controller.WisplyController.Prepare()
	if !controller.AccountConnected {
		controller.WisplyController.RedirectToLoginPage()
	}
	controller.loadLayout()
}

func (controller *Controller) loadLayout() {
	controller.Layout = "site/account-layout.tpl"
}

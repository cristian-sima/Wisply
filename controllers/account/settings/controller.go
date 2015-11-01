package settings

import "github.com/cristian-sima/Wisply/controllers/account"

// Controller represents the basic Account controller
type Controller struct {
	account.Controller
}

// Prepare redirects to a login page in case the account is not connected,
// otherwise it loads the page
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetTemplatePath("account/settings")
}

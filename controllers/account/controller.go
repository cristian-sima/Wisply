package account

import "github.com/cristian-sima/Wisply/controllers/wisply"

// Controller represents the basic Account controller
type Controller struct {
	wisply.Controller
}

// Prepare redirects to a login page in case the account is not connected,
// otherwise it loads the page
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	if !controller.AccountConnected {
		controller.Controller.RedirectToLoginPage()
	} else {
		controller.SetLayout("account")
		controller.SetTemplatePath("account")
	}
}

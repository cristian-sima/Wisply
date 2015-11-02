package developer

import "github.com/cristian-sima/Wisply/controllers/wisply"

// Controller represents the basic API controller
type Controller struct {
	wisply.Controller
}

// Prepare redirects to a login page in case the account is not connected,
// else it loads the page
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.SetLayout("public")
	controller.SetTemplatePath("developer")
}

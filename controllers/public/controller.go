package public

import "github.com/cristian-sima/Wisply/controllers/wisply"

// Controller can be accsed by the users who are not connected
type Controller struct {
	wisply.Controller
}

// Prepare redirects to a login page in case the account is not connected,
// else it loads the page
func (controller *Controller) Prepare() {
	controller.Controller.Prepare()
	controller.loadLayout()
}

func (controller *Controller) loadLayout() {
	controller.Layout = "site/public-layout.tpl"
}

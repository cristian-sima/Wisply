// Package public manages all the public pages that do not require a connection
package public

import (
	general "github.com/cristian-sima/Wisply/controllers/general"
)

// Controller can be accsed by the users who are not connected
type Controller struct {
	general.WisplyController
}

// Prepare redirects to a login page in case the account is not connected,
// else it loads the page
func (controller *Controller) Prepare() {
	controller.WisplyController.Prepare()
	controller.loadLayout()
}

func (controller *Controller) loadLayout() {
	controller.Layout = "site/public-layout.tpl"
}
